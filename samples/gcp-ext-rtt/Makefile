SERVICE_NAME=latency-test
GCP_REGION=europe-west9
TARGET_PORT=80
# TARGET_URL=ap1.binance.com
TARGET_URL=rbx.proof.ovh.net
IMAGE_TAG=us-central1-docker.pkg.dev/gcda-apac-sc/container-demo/${SERVICE_NAME}
TARGET_GCS_BUCKET=gcda-apac-sc/latency
GCP_PROJECT=gcda-apac-sc

create-container:
	gcloud builds submit . \
		--tag ${IMAGE_TAG} \
	    --project=${GCP_PROJECT}
deploy:
	gcloud run jobs deploy ${SERVICE_NAME} \
		--image ${IMAGE_TAG} \
		--tasks 1 \
		--set-env-vars TARGET_PORT=${TARGET_PORT} \
		--set-env-vars TARGET_URL=${TARGET_URL} \
		--set-env-vars TARGET_GCS_BUCKET=${TARGET_GCS_BUCKET} \
		--max-retries 5 \
		--region=${GCP_REGION} \
		--project=${GCP_PROJECT}

run:
	gcloud run jobs execute ${SERVICE_NAME} \
		--region=${GCP_REGION} \
		--project=${GCP_PROJECT}