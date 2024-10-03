#!/bin/bash
STAMP=$(date '+%Y-%m-%d')
REGION=$(curl -H "Metadata-Flavor: Google" "http://metadata.google.internal/computeMetadata/v1/instance/region" | cut -d'/' -f 4)
FILE_NAME=${STAMP}--${REGION}.csv
echo "Latency report will be written to: ${FILE_NAME}"
mtr -T -o "L BAWV MIX" -z -P ${TARGET_PORT} ${TARGET_URL} --report --csv -c50 > ${FILE_NAME}
gcloud storage cp ${FILE_NAME} gs://${TARGET_GCS_BUCKET}