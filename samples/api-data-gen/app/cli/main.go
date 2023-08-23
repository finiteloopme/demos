package main

import (
	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/finiteloopme/goutils/pkg/log"
)

func main() {
	log.Info("In main")
	// Petstore e.g.: https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json
	// Open Banking e.g.: https://consumerdatastandardsaustralia.github.io/standards/includes/swagger/cds_banking.json
	openApi, err := util.LoadSwagger("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore-expanded.json")
	if err != nil {
		log.Fatal(err)
	}
	// Iterate through the API Endpoints
	for path := range openApi.Paths {
		log.Info(path)
		// Operations for each Endpoint
		pathItem := openApi.Paths.Find(path)
		// GET
		operation := pathItem.Get
		log.Info("\tGET: " + operation.OperationID)
		// POST
		operationPut := pathItem.Post
		log.Info("\tGET: " + operationPut.OperationID)
	}
}
