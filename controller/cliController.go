package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
)

type PeerController struct {
	ctx         context.Context
	logger      *log.Logger
	executor    utils.Executor
	fabricSetup FabricSetup
}

func NewPeerController(c context.Context, l *log.Logger, setup FabricSetup) Controller {
	return &PeerController{
		ctx:         c,
		logger:      l,
		executor:    utils.NewExecutor(c, l),
		fabricSetup: setup,
	}
}

func NewPeerControllerWithExecutor(c context.Context, l *log.Logger, setup FabricSetup, exec utils.Executor) Controller {
	return &PeerController{
		ctx:         c,
		logger:      l,
		executor:    exec,
		fabricSetup: setup,
	}
}

func (p *PeerController) InvokeCreateResource(request rest_resources.CreateResourceRequest) rest_resources.CreateResourceResponse {
	p.logger.Printf("peercontroller-entering-InvokeCreateResource-%s", request.Function)
	defer p.logger.Printf("peercontroller-exiting-InvokeCreateResource-%s", request.Function)

	queryArgs := []string{request.Function, request.ResourceString, request.AuthorizationToken}
	query := rest_resources.Query{Args: queryArgs}
	stringResponse, err := p.invoke(query)
	if err != nil {
		errorMsg := fmt.Sprintf("error-to-execute-cmd-create-resource-on-blockchain-%s", p.getErrorMsgFromRawResponse(stringResponse, request.ErrorStringPattern))
		p.logger.Printf(errorMsg)
		return rest_resources.CreateResourceResponse{Error: errorMsg}
	}
	if !strings.Contains(stringResponse, request.SuccessStringPattern) {
		errorMsg := fmt.Sprintf("error-to-create-resource-on-blockchain")
		p.logger.Printf(errorMsg)
		return rest_resources.CreateResourceResponse{Response: stringResponse, Error: errorMsg}
	}

	p.logger.Printf("created-resource-%#v", stringResponse)
	return rest_resources.CreateResourceResponse{Response: request.ResourceString}
}

func (p *PeerController) InvokeGetResource(request rest_resources.GetResourceRequest) rest_resources.GetResourceResponse {
	p.logger.Printf("peercontroller-entering-InvokeGetResource-%s", request.Function)
	defer p.logger.Printf("peercontroller-exiting-InvokeGetResource-%s", request.Function)
	queryArgs := []string{request.Function, request.ID, request.AuthorizationToken}
	query := rest_resources.Query{Args: queryArgs}
	stringResponse, err := p.invoke(query)
	if err != nil {
		errorMsg := fmt.Sprintf("error-to-execute-cmd-get-response-from-blockchain-%s", p.getErrorMsgFromRawResponse(stringResponse, request.ErrorStringPattern))
		p.logger.Printf(errorMsg)
		return rest_resources.GetResourceResponse{Error: errorMsg}
	}
	if !strings.Contains(stringResponse, request.SuccessStringPattern) {
		errorMsg := fmt.Sprintf("error-to-get-resource-from-blockchain")
		p.logger.Printf(errorMsg)
		return rest_resources.GetResourceResponse{Response: stringResponse, Error: errorMsg}
	}

	resourceString := p.getJsonResponseFromRawResponse(stringResponse, request.SuccessStringPattern)

	err = json.Unmarshal([]byte(resourceString), &request.Resource)
	if err != nil {
		p.logger.Printf(err.Error())
		return rest_resources.GetResourceResponse{Error: err.Error()}
	}
	p.logger.Printf("returned-resource-%#v", request.Resource)
	return rest_resources.GetResourceResponse{Response: request.Resource}
}

func (p *PeerController) InvokeQuery(request rest_resources.QueryRequest) rest_resources.QueryResponse {
	p.logger.Printf("peercontroller-entering-InvokeQuery-%s", request.Function)
	defer p.logger.Printf("peercontroller-exiting-InvokeQuery-%s", request.Function)

	query := rest_resources.Query{}
	var queryArgs []string
	if request.Args == nil || len(request.Args) == 0 {
		queryArgs = []string{request.Function}

	} else {
		queryArgs = append([]string{request.Function}, request.Args...)
	}

	queryArgs = append(queryArgs, request.AuthorizationToken)
	query = rest_resources.Query{Args: queryArgs}

	stringResponse, err := p.invoke(query)
	if err != nil {
		errorMsg := fmt.Sprintf("error-to-execute-cmd-get-response-from-blockchain-%s", p.getErrorMsgFromRawResponse(stringResponse, request.ErrorStringPattern))
		p.logger.Printf(errorMsg)
		return rest_resources.QueryResponse{Error: errorMsg}
	}
	if !strings.Contains(stringResponse, request.SuccessStringPattern) {
		errorMsg := fmt.Sprintf("error-to-get-resource-from-blockchain")
		p.logger.Printf(errorMsg)
		return rest_resources.QueryResponse{Response: stringResponse, Error: errorMsg}
	}

	var resourceString, metadataString string
	var metadata []rest_resources.ResponseMetadata

	if strings.Contains(request.Function, "WithPagination") {
		resourceString, metadataString = p.getJsonResponseFromRawResponseWithPagination(stringResponse, request.SuccessStringPattern)
		metadata = []rest_resources.ResponseMetadata{}
		err = json.Unmarshal([]byte(metadataString), &metadata)
		if err != nil {
			p.logger.Printf(err.Error())
			return rest_resources.QueryResponse{Error: err.Error()}
		}

	} else {
		resourceString = p.getJsonResponseFromRawResponse(stringResponse, request.SuccessStringPattern)
	}

	err = json.Unmarshal([]byte(resourceString), &request.ResponseContainer)
	if err != nil {
		p.logger.Printf(err.Error())
		return rest_resources.QueryResponse{Error: err.Error()}
	}

	p.logger.Printf("returned-resources-%#v", request.ResponseContainer)
	return rest_resources.QueryResponse{Response: request.ResponseContainer, ResponseMetadata: metadata}
}

func (p *PeerController) invoke(query rest_resources.Query) (string, error) {
	p.logger.Printf("peercontroller-entering-invoke")
	defer p.logger.Printf("peercontroller-exiting-invoke")
	queryBytes, err := json.Marshal(query)
	if err != nil {
		p.logger.Printf(err.Error())
		return "", err
	}
	queryString := fmt.Sprintf(rest_resources.ChaincodePeerInvokeCommand, p.fabricSetup.OrdererURL, p.fabricSetup.ChannelID, p.fabricSetup.ChainCodeID, queryBytes)
	args := []string{"-c", queryString}
	response, err := p.executor.Execute("sh", args)
	if err != nil {
		errorMsg := fmt.Sprintf("error-executing-cmd-%s", err.Error())
		p.logger.Printf(errorMsg)
		return string(response[:]), fmt.Errorf(errorMsg)
	}
	return string(response[:]), nil
}

func (p *PeerController) getJsonResponseFromRawResponse(rawResponse string, successStringPattern string) string {
	p.logger.Printf("entering-getJsonResponseFromRawResponse")
	defer p.logger.Printf("peercontroller-exiting-getJsonResponseFromRawResponse")

	respArray := strings.Split(rawResponse, successStringPattern)
	jsonResponse := respArray[1]
	oldLen := 0
	newLen := len(jsonResponse)
	for oldLen != newLen {
		oldLen = newLen
		jsonResponse = strings.TrimRight(jsonResponse, "\n")
		jsonResponse = strings.TrimRight(jsonResponse, "\t")
		jsonResponse = strings.TrimRight(jsonResponse, " ")
		jsonResponse = strings.TrimRight(jsonResponse, "\"")
		newLen = len(jsonResponse)
	}
	return strings.Replace(jsonResponse, "\\\"", "\"", -1)
}

func (p *PeerController) getJsonResponseFromRawResponseWithPagination(rawResponse string, successStringPattern string) (string, string) {
	p.logger.Printf("entering-getJsonResponseFromRawResponse")
	defer p.logger.Printf("peercontroller-exiting-getJsonResponseFromRawResponse")

	jsonResponse := p.getJsonResponseFromRawResponse(rawResponse, successStringPattern)
	p.logger.Printf("jsonResponse:%s", jsonResponse)
	// res := strings.Split(jsonResponse, "}][{")
	res := strings.Split(jsonResponse, "][")
	// return fmt.Sprintf("%s}]", res[0]), fmt.Sprintf("[{%s", res[1])
	return fmt.Sprintf("%s]", res[0]), fmt.Sprintf("[%s", res[1])
}
func (p *PeerController) getErrorMsgFromRawResponse(rawResponse string, errorStringPattern string) string {
	p.logger.Printf("entering-getErrorMsgFromRawResponse")
	defer p.logger.Printf("peercontroller-exiting-getErrorMsgFromRawResponse")

	respArray := strings.Split(rawResponse, errorStringPattern)
	jsonResponse := respArray[1]
	oldLen := 0
	newLen := len(jsonResponse)
	for oldLen != newLen {
		oldLen = newLen
		jsonResponse = strings.TrimRight(jsonResponse, "\n")
		jsonResponse = strings.TrimRight(jsonResponse, "\t")
		jsonResponse = strings.TrimRight(jsonResponse, " ")
		jsonResponse = strings.TrimRight(jsonResponse, "\"")
		newLen = len(jsonResponse)
	}
	return jsonResponse
}
