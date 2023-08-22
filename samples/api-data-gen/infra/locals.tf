locals {

  # API services
  project_apis = [
        "compute.googleapis.com",
       # Enabling the ServiceUsage API allows the new project to be quota checked from now on.
        "serviceusage.googleapis.com",
        "logging.googleapis.com",
  ]
}