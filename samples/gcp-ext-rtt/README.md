# MTR

**MTR** is a tool that allows you to test round trip latency (mtr) from any
Google Cloud region to an arbitrary endpoint.  

## Steps

1. Create a container for `mtr-script.sh`:

   ```bash
   make create-container
   ```

2. Deploy a _Cloud Run Job_ in a desired region:

   ```bash
   make deploy
   ```

3. Configure the `TARGET_PORT` & `TARGET_URL` as the environment variables in the `Makefile`
4. Execute the job

   ```bash
   make run
   ```

5. By default results are available at: `<date>--<GCP_REGION>.csv`
