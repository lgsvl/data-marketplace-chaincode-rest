package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) DataContractDeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *CCHandler) DataContractFindById(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("entering-get-data-data-contract")
	defer c.logger.Printf("exiting-get--data-contract")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractID := utils.ExtractVarsFromRequest(r, "id")
	c.invokeGet(w, r, rest_resources.GetDataContractFunc, dataContractID, rest_resources.GetResourceSuccessMsgPattern, rest_resources.GetResourceErrorMsgPattern, resources.DataContract{})
}

func (c *CCHandler) DataContractReplaceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
