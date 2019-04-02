package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) PersonCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-PersonCreate")
	defer c.logger.Printf("ccHandler-exiting-PersonCreate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	person := resources.Person{}
	c.invokeCreate(w, r, rest_resources.AddPersonFunc, person, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}

func (c *CCHandler) PersonDeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *CCHandler) PersonFindById(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("entering-get-person")
	defer c.logger.Printf("exiting-get-person")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	personID := utils.ExtractVarsFromRequest(r, "id")
	c.invokeGet(w, r, rest_resources.GetPersonFunc, personID, rest_resources.GetResourceSuccessMsgPattern, rest_resources.GetResourceErrorMsgPattern, resources.Person{})
}

func (c *CCHandler) PersonReplaceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
