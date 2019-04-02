package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) DataReceivedByConsumerCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-DataReceivedByConsumerCreate")
	defer c.logger.Printf("ccHandler-exiting-DataReceivedByConsumerCreate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataReceivedByConsumer := resources.DataReceivedByConsumer{}
	c.invokeCreate(w, r, rest_resources.SetDataReceivedByConsumerFunc, dataReceivedByConsumer, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}
