package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathListRoles(b *backend) *framework.Path {
	return &framework.Path{
		Pattern: "roles/?$",

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ListOperation: b.pathRoleList,
		},
	}
}
func pathRoles() *framework.Path {
	return &framework.Path{
		Pattern: "roles/" + framework.GenericNameRegex("name"),
		Fields: map[string]*framework.FieldSchema{
			"name": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Name of the role",
			},

			"policy": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: `Policy document. Required`,
			},

			"lease": &framework.FieldSchema{
				Type:        framework.TypeDurationSecond,
				Description: "Lease time of the role.",
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation:   pathRolesRead,
			logical.UpdateOperation: pathRolesWrite,
			logical.DeleteOperation: pathRolesDelete,
		},
	}
}

func (b *backend) pathRoleList(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	entries, err := req.Storage.List(ctx, "policy/")
	if err != nil {
		return nil, err
	}

	return logical.ListResponse(entries), nil
}

func pathRolesRead(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Searching for policy %s", name)
	entry, err := req.Storage.Get(ctx, "policy/"+name)
	if err != nil {
		log.Printf("[DEBUG] Searching for policy %s: %s", name, err)
		return nil, err
	}
	if entry == nil {
		log.Printf("[DEBUG] policy %s: nil entry", name)
		return nil, fmt.Errorf("Could not find a policy for %s", name)
	}

	var result roleConfig
	if err := entry.DecodeJSON(&result); err != nil {
		return nil, err
	}

	// Generate the response
	resp := &logical.Response{
		Data: map[string]interface{}{
			"lease": int64(result.Lease.Seconds()),
		},
	}
	if result.Policy != "" {
		resp.Data["policy"] = []byte(result.Policy)
		resp.Data["raw_policy"] = result.RawPolicy
	}
	return resp, nil
}

func pathRolesWrite(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	// do some validation here
	name := d.Get("name").(string)
	policy := d.Get("policy").(string)
	var err error

	if policy == "" {
		return logical.ErrorResponse("policy cannot be empty"), nil
	}

	realPolicy := kafka.StringlyTypedACL{}
	err = json.Unmarshal([]byte(policy), &realPolicy)
	if err != nil {
		return logical.ErrorResponse(fmt.Sprintf("Error unparshalling policy: %s", err)), nil
	}
	realPolicy.ACL.Principal = ""

	var lease time.Duration
	leaseParamRaw, ok := d.GetOk("lease")
	if ok {
		lease = time.Second * time.Duration(leaseParamRaw.(int))
	}

	mj, err := json.Marshal(realPolicy)
	if err != nil {
		return logical.ErrorResponse(fmt.Sprintf("Error marshalling policy: %s", err)), nil
	}

	entry, err := logical.StorageEntryJSON("policy/"+name, roleConfig{
		Policy:    string(mj),
		RawPolicy: realPolicy,
		Lease:     lease,
	})
	if err != nil {
		return nil, err
	}

	if err := req.Storage.Put(ctx, entry); err != nil {
		return nil, err
	}

	return nil, nil
}

func pathRolesDelete(ctx context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
	name := d.Get("name").(string)
	if err := req.Storage.Delete(ctx, "policy/"+name); err != nil {
		return nil, err
	}
	return nil, nil
}

type roleConfig struct {
	Policy    string                 `json:"policy"`
	RawPolicy kafka.StringlyTypedACL `json:"rawpolicy"`
	Lease     time.Duration          `json:"lease"`
}
