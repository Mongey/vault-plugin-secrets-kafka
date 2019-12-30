module github.com/Mongey/vault-plugin-secrets-kafka

go 1.13

replace github.com/Mongey/terraform-provider-kafka v0.2.3 => ../terraform-provider-kafka/

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Mongey/terraform-provider-kafka v0.2.3-0.20191229235334-7090c7cff2aa
	github.com/Shopify/sarama v1.24.1
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/containerd/continuity v0.0.0-20191214063359-1097c8bae83b // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20191128021309-1d7a30a10f73 // indirect
	github.com/gocql/gocql v0.0.0-20191224103530-b7facab47581 // indirect
	github.com/hashicorp/consul/api v1.3.0 // indirect
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-memdb v1.0.4 // indirect
	github.com/hashicorp/go-uuid v1.0.2-0.20191001231223-f32f5fe8d6a8
	github.com/hashicorp/memberlist v0.1.5 // indirect
	github.com/hashicorp/serf v0.8.5 // indirect
	github.com/hashicorp/vault v1.3.1
	github.com/hashicorp/vault/api v1.0.5-0.20191216174727-9d51b36f3ae4
	github.com/hashicorp/vault/sdk v0.1.14-0.20191218020134-06959d23b502
	github.com/jefferai/jsonx v1.0.1 // indirect
	github.com/ory/dockertest v3.3.5+incompatible // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
)
