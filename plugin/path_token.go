package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/errwrap"
	uuid "github.com/hashicorp/go-uuid"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

const (
	SecretTokenType = "token"
)

func pathToken(b *backend) *framework.Path {
	return &framework.Path{
		Pattern: "creds/" + framework.GenericNameRegex("role"),
		Fields: map[string]*framework.FieldSchema{
			"role": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Name of the role",
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation: b.pathTokenRead,
		},
	}
}

func (b *backend) pathTokenRead(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	role := d.Get("role").(string)
	log.Printf("[INFO] Reading role %s", role)

	log.Printf("[INFO] Getting Client ")
	c, userErr, intErr := client(ctx, req.Storage)
	log.Printf("[INFO] Got client %v", c)
	if intErr != nil {
		return nil, intErr
	}
	if userErr != nil {
		return logical.ErrorResponse(userErr.Error()), nil
	}

	log.Printf("[INFO] Retrieving role %s", role)
	entry, err := req.Storage.Get(ctx, "policy/"+role)
	if err != nil {
		return nil, errwrap.Wrapf("error retrieving role: {{err}}", err)
	}
	if entry == nil {
		return logical.ErrorResponse(fmt.Sprintf("role %q not found", role)), nil
	}

	var result roleConfig
	if err := entry.DecodeJSON(&result); err != nil {
		return nil, err
	}

	// Create a ACL, with the principal
	meUUID, err := uuid.GenerateUUID()
	maxLength := 64
	user := fmt.Sprintf("%s-%s", meUUID, role)
	if len(user) > maxLength {
		user = user[:maxLength-1]
	}

	p := result.RawPolicy
	p.ACL.Principal = "User:CN=" + user

	err = c.CreateACL(p)
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	pp, err := json.Marshal(p)
	fmt.Println("[INFO] pp" + string(pp))
	policy := result.Policy
	s := b.Secret(SecretTokenType).Response(map[string]interface{}{
		"policy": policy,
		"user":   user,
	}, map[string]interface{}{
		"user": user,
	})
	s.Secret.TTL = result.Lease

	return s, nil
}

func secretToken(b *backend) *framework.Secret {
	return &framework.Secret{
		Type: SecretTokenType,
		Fields: map[string]*framework.FieldSchema{
			"token": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Request token",
			},
		},

		Renew:  b.secretTokenRenew,
		Revoke: b.secretTokenRevoke,
	}
}

func (b *backend) secretTokenRevoke(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	c, userErr, intErr := client(ctx, req.Storage)
	if intErr != nil {
		return nil, intErr
	}
	if userErr != nil {
		//Returning logical.ErrorResponse from revocation function is risky
		return nil, userErr
	}

	for k, v := range req.Secret.InternalData {
		m := v.(string)

		fmt.Println("[WARN] InternalData|" + k + ": " + m)
	}

	// Get the kafka client
	//c, userErr, intErr := client(ctx, req.Storage)
	//if intErr != nil {
	//return nil, intErr
	//}
	//if userErr != nil {
	//return logical.ErrorResponse(userErr.Error()), nil
	//}

	acls, err := c.ListACLs()
	if err != nil {
		return nil, err
	}

	user := req.Secret.InternalData["user"].(string)
	for _, acl := range acls {
		found := false
		for _, a := range acl.Acls {
			fmt.Println(a.Principal)
			if a.Principal == ("User:CN=" + user) {
				fmt.Println("[WARN] Found an ACL for " + user)
				found = true
				break
			}

			if found {
				break
			}
		}
	}

	principal := "User:" + user
	fmt.Println("[INFO] Deleting " + principal)
	//err = c.DeleteACLPrincipal(principal)
	return nil, err
}

func (b *backend) secretTokenRenew(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	resp := &logical.Response{Secret: req.Secret}
	roleRaw, ok := req.Secret.InternalData["role"]
	if !ok || roleRaw == nil {
		return resp, nil
	}

	role, ok := roleRaw.(string)
	if !ok {
		return resp, nil
	}

	entry, err := req.Storage.Get(ctx, "policy/"+role)
	if err != nil {
		return nil, errwrap.Wrapf("error retrieving role: {{err}}", err)
	}
	if entry == nil {
		return logical.ErrorResponse(fmt.Sprintf("issuing role %q not found", role)), nil
	}

	var result roleConfig
	if err := entry.DecodeJSON(&result); err != nil {
		return nil, err
	}
	resp.Secret.TTL = result.Lease
	return resp, nil
}
