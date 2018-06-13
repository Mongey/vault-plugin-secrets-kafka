# `[WIP] vault-plugin-kafka-secret`

A Vault plugin for generating credentials for Apache Kafka clients.

Generates a dynamic username and ACL that can be used to create a uniq SSL
certificate for a Kafka client.

Use this in combination with the vault pki backend.

## 🔌 Installation

* Download the plugin to Vault's plugin directory.
* Register the plugin with Vault
  * ```sh
    vault write sys/plugins/catalog/vault-plugin-secrets-kafka \
      sha_256="$SHASUM" \
      command="vault-plugin-secrets-kafka"
    ```
* Enable the plugin mount
  * ```sh
    vault secrets enable -path=kafka -plugin-name=vault-plugin-secrets-kafka plugin
    ```

## 🛠 Configure
* Configure the plugin
  * ```sh
    vault write kafka/config/access address="localhost:9092" ca_certificate="$CA" client_certificate="$CERT" client_key="$PRIVATE_KEY"
    ```
  * The client must be capable of writing creating and deleting ACLs.

* Write a policy
  * ```json
    {
      "acl": {
        "host": "*",
        "operation": "Read",
        "permission_type": "Allow"
      },
      "resource": {
        "type": "Topic",
        "name": "*",
        "pattern_type_filter": "any"
      }
    }
    ```
* Write the role
  * `vault write kafka/roles/read-all-topics policy=$(cat bin/policy.json)`
* Read the credentials, pick the username
  * `vault read kafka/creds/read-all-topics`
* Generate a SSL certificate for this client
  * ```sh
      NAME=$(vault read -field=user kafka/creds/read-all-topics)
      DATA=$(vault write -format=json pki/issue/kafka-clients common_name="$NAME" ttl=$TTL | jq -r .data)
      printf "%s" "$DATA" | jq -r .private_key > private.key
      printf "%s" "$DATA" | jq -r .certificate > client.cert
      printf "%s" "$DATA" | jq -r .issuing_ca  > ca.cert
    ```
