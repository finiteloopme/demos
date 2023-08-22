
# Demo Overview


## Deployment

```bash
# clone the git repo
git clone https://github.com/finiteloopme/demos
cd demos/samples/api-data-gen
# Deploy Infrastructure in GCP
# We assume the you have a GCP project
# And you have created a GCS bucket: ${GCP_PROJECT_ID}-tf-state
make infra-init
make infra-plan
make infra-deploy
```

## Clean up

```bash
make infra-undeploy
```