#!/bin/bash
set -e

export VAULT_ADDR=http://localhost:8200

TTL=1000
BROKER_NAME=broker0
NETWORK_ADDRESS=localhost
OUTPUT_FILE=server.keystore
PASSWORD=magicbeans
IP_ADDRESS=$(ipconfig getifaddr en0)

set +e
echo "Cleaning old server.keystore"
rm $OUTPUT_FILE
set -e

echo "Generating certificate"
DATA=$(vault write -format=json pki/issue/kafka-clients ip_sans="$IP_ADDRESS" common_name="$BROKER_NAME" alt_names="$NETWORK_ADDRESS" ttl="$TTL" | jq -r .data)
PRIVATE_KEY=$(printf "%s" "$DATA" | jq -r .private_key)
CA=$(printf "%s" "$DATA" | jq -r .issuing_ca)
CLIENT=$(printf "%s" "$DATA" | jq -r .certificate)

echo "Outputting ca"
echo "$CA" > ca.crt
echo "Outputting pk"
echo "$PRIVATE_KEY" > server.key
echo "Outputting client"
echo "$CLIENT" > server.crt

echo "=====>>>> Generating PKC12 file"
openssl pkcs12 \
  -export -in server.crt \
  -inkey server.key \
  -name $BROKER_NAME \
  -passin pass:"$PASSWORD" \
  -password pass:"$PASSWORD" \
  -CAfile ca.crt -caname root \
  -out server.p12

echo "=====>>>> Converting PKC12 file to Keystore"
keytool -importkeystore \
        -deststorepass "$PASSWORD" \
        -destkeypass "$PASSWORD" \
        -destkeystore $OUTPUT_FILE \
        -srckeystore server.p12 \
        -srcstoretype PKCS12 \
        -srcstorepass "$PASSWORD" \
        -alias $BROKER_NAME

rm ca.crt server.key server.crt server.p12

echo $PASSWORD > tmp/secrets/broker0_keystore_creds
mv $OUTPUT_FILE tmp/secrets/kafka.broker0.keystore.jks

echo "Done!"
