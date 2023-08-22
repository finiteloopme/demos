module "default-network" {
    source  = "github.com/finiteloopme/demos//modules/network"
    project_id = module.manage-project-apis.project_id
}
