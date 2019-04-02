package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/lgsvl/data-marketplace-chaincode-rest/controller"
	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
)

//CCHandler is a handler for chaincode
type CCHandler struct {
	ctx        context.Context
	logger     *log.Logger
	controller controller.Controller
}

//NewCCHandler returns a new CCHandler
func NewCCHandler(c context.Context, lg *log.Logger, cont controller.Controller) *CCHandler {
	return &CCHandler{ctx: c, logger: lg, controller: cont}
}

func (c *CCHandler) invokeCreate(w http.ResponseWriter, r *http.Request, function string, resource interface{}, successStringPattern string, errorStringPattern string) {
	resourceBytes, err := utils.UnmarshalDataFromRequest(r, &resource)
	if err != nil {
		c.logger.Printf(err.Error())
		utils.WriteResponse(w, 500, err)
		return
	}

	request := rest_resources.CreateResourceRequest{
		AuthorizationToken:   utils.GetFromHeader(r, "authorization"),
		ResourceString:       string(resourceBytes[:]),
		Function:             function,
		SuccessStringPattern: successStringPattern,
		ErrorStringPattern:   errorStringPattern,
	}
	response := c.controller.InvokeCreateResource(request)
	if response.Error != "" {
		c.logger.Printf(response.Error)
		utils.WriteResponse(w, 500, response)
		return
	}

	utils.WriteResponse(w, http.StatusOK, resource)

}

func (c *CCHandler) invokeGet(w http.ResponseWriter, r *http.Request, function string, ID string, successStringPattern string, errorStringPattern string, container interface{}) {

	request := rest_resources.GetResourceRequest{
		AuthorizationToken:   utils.GetFromHeader(r, "authorization"),
		ID:                   ID,
		Function:             function,
		SuccessStringPattern: successStringPattern,
		ErrorStringPattern:   errorStringPattern,
		Resource:             container,
	}

	response := c.controller.InvokeGetResource(request)
	if response.Error != "" {
		c.logger.Printf(response.Error)
		utils.WriteResponse(w, 500, response)
		return
	}
	c.logger.Printf("resource-retrieved-%#v", response.Response)
	utils.WriteResponse(w, http.StatusOK, response)
}

func (c *CCHandler) invokeQuery(w http.ResponseWriter, r *http.Request, function string, successStringPattern string, container interface{}, args []string) {

	request := rest_resources.QueryRequest{
		AuthorizationToken:   utils.GetFromHeader(r, "authorization"),
		Function:             function,
		SuccessStringPattern: successStringPattern,
		ResponseContainer:    container,
		Args:                 args,
	}

	response := c.controller.InvokeQuery(request)
	if response.Error != "" {
		c.logger.Printf(response.Error)
		utils.WriteResponse(w, 500, response)
		return
	}
	c.logger.Printf("resources-retrieved-%#v", response.Response)
	utils.WriteResponse(w, http.StatusOK, response)
}
