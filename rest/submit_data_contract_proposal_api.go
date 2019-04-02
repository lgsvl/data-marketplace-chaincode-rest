package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) SubmitDataContractProposalCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-SubmitDataContractProposalCreate")
	defer c.logger.Printf("ccHandler-exiting-SubmitDataContractProposalCreate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractProposal := resources.DataContractProposal{}
	c.invokeCreate(w, r, rest_resources.SubmitDataContractProposalFunc, dataContractProposal, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}
