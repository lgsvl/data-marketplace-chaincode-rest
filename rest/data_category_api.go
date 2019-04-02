package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) DataCategoryCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-DataCategoryCreate")
	defer c.logger.Printf("ccHandler-exiting-DataCategoryCreate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataCategory := resources.DataCategory{}
	c.invokeCreate(w, r, rest_resources.CreateDataCategoryFunc, dataCategory, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}

func (c *CCHandler) DataCategoryDeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *CCHandler) DataCategoryFindById(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("entering-get-data-data-category")
	defer c.logger.Printf("exiting-get--data-category")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataCategoryID := utils.ExtractVarsFromRequest(r, "id")
	c.invokeGet(w, r, rest_resources.GetDataCategoryFunc, dataCategoryID, rest_resources.GetResourceSuccessMsgPattern, rest_resources.GetResourceErrorMsgPattern, resources.DataCategory{})
}

func (c *CCHandler) DataCategoryReplaceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
