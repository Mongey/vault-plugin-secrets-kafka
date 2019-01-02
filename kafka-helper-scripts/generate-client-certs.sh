set -e
export VAULT_ADDR=http://localhost:8200

TTL=1000
PARAM=$1
NAME=${PARAM:-"unknown client"}
DATA=$(vault write -format=json pki/issue/kafka-clients common_name="$NAME" ttl=$TTL | jq -r .data)

echo "Hello $NAME"
printf "%s" "$DATA" | jq -r .private_key > private.key
printf "%s" "$DATA" | jq -r .issuing_ca > ca.cert
printf "%s" "$DATA" | jq -r .certificate > client.cert
echo "Done!"
