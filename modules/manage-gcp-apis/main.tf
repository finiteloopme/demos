# Get the project
data "google_project" "project" {
    project_id = var.project_id
}

# Manage Project APIs
module "manage-project-apis" {
    source = "terraform-google-modules/project-factory/google//modules/project_services"
    project_id = data.google_project.project.project_id

    activate_apis = var.project_apis

    enable_apis = true

    # Disable services (APIs) when the resources are destroyed?
    disable_services_on_destroy = var.disable_services_on_destroy

    # Whether services that are enabled and which depend on this service should 
    # also be disabled when this service is destroyed.
    disable_dependent_services = var.disable_dependent_services
}

# We need to wait for services to "start"
resource "time_sleep" "wait_for_gcp_services" {
  create_duration = var.wait_time

  depends_on = [module.manage-project-apis]
}

