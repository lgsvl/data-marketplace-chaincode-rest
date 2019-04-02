package controller

import (
	"github.com/lgsvl/data-marketplace-chaincode-rest/resources"
)

//Controller to run invoke and queries on chaincode
//go:generate counterfeiter -o ../fakes/fake_controller.go . Controller
type Controller interface {
	InvokeCreateResource(resources.CreateResourceRequest) resources.CreateResourceResponse
	InvokeGetResource(resources.GetResourceRequest) resources.GetResourceResponse
	InvokeQuery(resources.QueryRequest) resources.QueryResponse
}
