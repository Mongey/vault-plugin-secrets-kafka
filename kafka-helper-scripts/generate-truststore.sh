set -e

PASSWORD=helloworld
TMP_FILE=issuing_ca.cer
OUTPUT_FILE=kafka.server.truststore.jks
export VAULT_ADDR=http://localhost:8200

echo "Cleaning up"
set +e
rm $OUTPUT_FILE
set -e

vault write -format=json \
  pki/issue/kafka-clients \
  common_name=go-client ttl=1 | jq -r .data.issuing_ca > $TMP_FILE

keytool -keystore $OUTPUT_FILE  \
  -alias CARoot \
  -import -file $TMP_FILE -noprompt \
  -storepass $PASSWORD

echo "Cleaning up"
set +e
rm issuing_ca.cer
set -e

