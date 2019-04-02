package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) DataContractTypeCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-DataContractTypeCreate")
	defer c.logger.Printf("ccHandler-exiting-DataContractTypeCreate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractType := resources.DataContractType{}
	c.invokeCreate(w, r, rest_resources.CreateDataContractTypeFunc, dataContractType, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}

func (c *CCHandler) DataContractTypeDeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *CCHandler) DataContractTypeFindById(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("entering-get-data-data-type-contract")
	defer c.logger.Printf("exiting-get--data-contract-type")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractTypeID := utils.ExtractVarsFromRequest(r, "id")
	c.invokeGet(w, r, rest_resources.GetDataContractTypeFunc, dataContractTypeID, rest_resources.GetResourceSuccessMsgPattern, rest_resources.GetResourceErrorMsgPattern, resources.DataContractType{})
}

func (c *CCHandler) DataContractTypeReplaceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
