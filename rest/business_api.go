package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) BusinessCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-BusinessCreate")
	defer c.logger.Printf("ccHandler-exiting-create-business")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	business := resources.Business{}
	c.invokeCreate(w, r, rest_resources.CreateBusinessFunc, business, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}

func (c *CCHandler) BusinessDeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *CCHandler) BusinessFindById(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-get-business")
	defer c.logger.Printf("ccHandler-exiting-get-business")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	businessID := utils.ExtractVarsFromRequest(r, "id")
	c.invokeGet(w, r, rest_resources.GetBusinessFunc, businessID, rest_resources.GetResourceSuccessMsgPattern, rest_resources.GetResourceErrorMsgPattern, resources.Business{})
}

func (c *CCHandler) BusinessReplaceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
