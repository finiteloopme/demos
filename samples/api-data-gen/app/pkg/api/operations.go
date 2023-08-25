package api

import "github.com/finiteloopme/goutils/pkg/log"

func (apiSpec *ApiSpec) LoadAllOperations() {
	for _, path := range apiSpec.Paths {
		apiSpec.LoadOperations(path.Path)
	}
}

func (apiSpec *ApiSpec) LoadOperations(pathName string) {
	if apiSpec.Paths == nil {
		apiSpec.LoadPaths()
	}

	path, err := apiSpec.GetPath(pathName)
	if err == nil {
		pathItem := apiSpec.Api.Paths.Find(path.Path)
		if pathItem.Get != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        GET.String(),
				OperationID: pathItem.Get.OperationID,
				// Used for prompt engineering?
				Description: pathItem.Get.Description,
				Summary:     pathItem.Get.Summary,
			})
		}
		if pathItem.Post != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        POST.String(),
				OperationID: pathItem.Post.OperationID,
				Description: pathItem.Post.Description,
				Summary:     pathItem.Post.Summary,
			})
		}
		if pathItem.Put != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        PUT.String(),
				OperationID: pathItem.Put.OperationID,
				Description: pathItem.Put.Description,
				Summary:     pathItem.Put.Summary,
			})
		}
		if pathItem.Delete != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        DELETE.String(),
				OperationID: pathItem.Delete.OperationID,
				Description: pathItem.Delete.Description,
				Summary:     pathItem.Delete.Summary,
			})
		}
		if pathItem.Patch != nil {
			path.Operations = append(path.Operations, &Operation{
				Type:        PATCH.String(),
				OperationID: pathItem.Patch.OperationID,
				Description: pathItem.Patch.Description,
				Summary:     pathItem.Patch.Summary,
			})
		}
	}
}

func (apiSpec *ApiSpec) PrintAllOperations() {
	for _, path := range apiSpec.Paths {
		apiSpec.PrintOperations(path.Path)
	}
}

func (apiSpec *ApiSpec) PrintOperations(pathName string) {
	path, _ := apiSpec.GetPath(pathName)
	apiSpec.PrintPath(pathName)
	for _, ops := range path.Operations {
		log.Info("	Method: " + ops.Type)
		log.Info("	OperationID: " + ops.OperationID)
		log.Info("	Summary: " + ops.Summary)
		log.Info("	Description: " + ops.Description)
	}
}
