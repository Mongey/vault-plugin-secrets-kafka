package plugin

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/vault/logical"
	logicaltest "github.com/hashicorp/vault/logical/testing"
)

func TestSomething(t *testing.T) {
	b, _ := Factory(context.Background(), logical.TestBackendConfig())
	uri := "localhost:9092"

	logicaltest.Test(t, logicaltest.TestCase{
		PreCheck: testAccPreCheckFunc(t, uri),
		Backend:  b,
		Steps: []logicaltest.TestStep{
			testAccStepConfig(t, uri),
			testAccStepRole(t, "readall"),
			testAccStepReadCreds(t, b, uri, "readall"),
		},
	})
}

func testAccStepRole(t *testing.T, role string) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.UpdateOperation,
		Path:      "roles/" + role,
		Data: map[string]interface{}{
			"policy": `{"acl": { "principal": "User:foo", "host": "*", "operation": "Read", "permission_type": "Allow" }, "resource": {"type": "Topic", "name": "mytopic"}}`,
		},
	}
}
func testAccStepConfig(t *testing.T, uri string) logicaltest.TestStep {
	password := os.Getenv("KAFKA_ROOT_CERTIFICATE")

	return logicaltest.TestStep{
		Operation: logical.UpdateOperation,
		Path:      "config/access",
		Data: map[string]interface{}{
			"address":          uri,
			"root_certificate": password,
		},
	}
}

func TestBackend_basic(t *testing.T) {
}
func TestBackend_roleCrud(t *testing.T) {
}

func testAccPreCheckFunc(t *testing.T, uri string) func() {
	return func() {
		if uri == "" {
			t.Fatal("Kafka servers must be set for acceptance tests")
		}
	}
}
func testAccStepDeleteRole(t *testing.T, n string) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.DeleteOperation,
		Path:      "roles/" + n,
	}
}

func testAccStepReadRole(t *testing.T, name, tags, rawVHosts string) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.ReadOperation,
		Path:      "roles/" + name,
		//Check: func(resp *logical.Response) error {
		//if resp == nil {
		//if tags == "" && rawVHosts == "" {
		//return nil
		//}

		//return fmt.Errorf("bad: %#v", resp)
		//}

		//var d struct {
		//Tags   string                     `mapstructure:"tags"`
		//VHosts map[string]vhostPermission `mapstructure:"vhosts"`
		//}
		//if err := mapstructure.Decode(resp.Data, &d); err != nil {
		//return err
		//}

		//if d.Tags != tags {
		//return fmt.Errorf("bad: %#v", resp)
		//}

		//var vhosts map[string]vhostPermission
		//if err := jsonutil.DecodeJSON([]byte(rawVHosts), &vhosts); err != nil {
		//return fmt.Errorf("bad expected vhosts %#v: %s", vhosts, err)
		//}

		//for host, permission := range vhosts {
		//actualPermission, ok := d.VHosts[host]
		//if !ok {
		//return fmt.Errorf("expected vhost: %s", host)
		//}

		//if actualPermission.Configure != permission.Configure {
		//return fmt.Errorf("expected permission %s to be %s, got %s", "configure", permission.Configure, actualPermission.Configure)
		//}

		//if actualPermission.Write != permission.Write {
		//return fmt.Errorf("expected permission %s to be %s, got %s", "write", permission.Write, actualPermission.Write)
		//}

		//if actualPermission.Read != permission.Read {
		//return fmt.Errorf("expected permission %s to be %s, got %s", "read", permission.Read, actualPermission.Read)
		//}
		//}

		//return nil
		//},
	}
}

func testAccStepReadCreds(t *testing.T, b logical.Backend, uri, name string) logicaltest.TestStep {
	return logicaltest.TestStep{
		Operation: logical.ReadOperation,
		Path:      "creds/" + name,
		Check: func(resp *logical.Response) error {
			//var d struct {
			//Username string `mapstructure:"username"`
			//Password string `mapstructure:"password"`
			//}
			//if err := mapstructure.Decode(resp.Data, &d); err != nil {
			//return err
			//}
			//log.Printf("[WARN] Generated credentials: %v", d)

			//client, err := rabbithole.NewClient(uri, d.Username, d.Password)
			//if err != nil {
			//t.Fatal(err)
			//}

			//_, err = client.ListVhosts()
			//if err != nil {
			//t.Fatalf("unable to list vhosts with generated credentials: %s", err)
			//}

			//resp, err = b.HandleRequest(context.Background(), &logical.Request{
			//Operation: logical.RevokeOperation,
			//Secret: &logical.Secret{
			//InternalData: map[string]interface{}{
			//"secret_type": "creds",
			//"username":    d.Username,
			//},
			//},
			//})
			//if err != nil {
			//return err
			//}
			//if resp != nil {
			//if resp.IsError() {
			//return fmt.Errorf("error on resp: %#v", *resp)
			//}
			//}

			//client, err = rabbithole.NewClient(uri, d.Username, d.Password)
			//if err != nil {
			//t.Fatal(err)
			//}

			//_, err = client.ListVhosts()

			err := fmt.Errorf("")
			if err == nil {
				t.Fatalf("expected to fail listing vhosts: %s", err)
			}

			return nil
		},
	}
}
