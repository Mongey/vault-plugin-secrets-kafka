set -e
export VAULT_ADDR=http://localhost:8200

TTL=1000
PARAM=$1
NAME=${PARAM:-"unknown client"}
DATA=$(vault write -format=json pki/issue/kafka-clients common_name="$NAME" ttl=$TTL | jq -r .data)

echo "Hello $NAME"
printf "%s" "$DATA" | jq -r .private_key > sample-app/private.key
printf "%s" "$DATA" | jq -r .issuing_ca > sample-app/ca.cert
printf "%s" "$DATA" | jq -r .certificate > sample-app/client.cert
echo "Done!"
