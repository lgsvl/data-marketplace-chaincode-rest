package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) ReviewCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-ReviewCreate")
	defer c.logger.Printf("ccHandler-exiting-ReviewCreate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	review := resources.Review{}
	c.invokeCreate(w, r, rest_resources.SubmitReviewFunc, review, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}

func (c *CCHandler) ReviewDeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (c *CCHandler) ReviewFindById(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("entering-get-review")
	defer c.logger.Printf("exiting-get-review")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	reviewID := utils.ExtractVarsFromRequest(r, "id")
	c.invokeGet(w, r, rest_resources.GetReviewFunc, reviewID, rest_resources.GetResourceSuccessMsgPattern, rest_resources.GetResourceErrorMsgPattern, resources.Review{})
}

func (c *CCHandler) ReviewReplaceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
