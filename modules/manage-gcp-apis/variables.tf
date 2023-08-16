variable "project_id" {
    description = "Project ID for the DB"
}

variable "project_apis" {
    type = list(string)
    description = "List of APIs"
}

variable "wait_time"{
    description = "Time to wait for API(s) call to 'activate'.  Default is 10s"
    default = "10s"
}

variable "disable_services_on_destroy" {
    description = "Disable services on resource destruction"
    default = true
}

variable "disable_dependent_services" {
    description = "How to handle dependent services, enabled by this API"
    default = true
}