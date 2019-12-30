set -e
export VAULT_ADDR=http://localhost:8200

TTL=1000
PARAM=$1
NAME=${PARAM:-"VaultKafkaPlugin"}
DATA=$(vault write -format=json pki/issue/kafka-clients common_name="$NAME" ttl=$TTL | jq -r .data)

echo "Creating a cert for $NAME"
echo $DATA
printf "%s" "$DATA" | jq -r .private_key > sample-app/private.key
printf "%s" "$DATA" | jq -r .issuing_ca > sample-app/ca.crt
printf "%s" "$DATA" | jq -r .certificate > sample-app/client.crt
echo "Done!"
