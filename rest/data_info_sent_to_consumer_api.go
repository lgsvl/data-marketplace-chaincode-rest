package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) DataInfoSentToConsumerCreate(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-DataInfoSentToConsumerCreate")
	defer c.logger.Printf("ccHandler-exiting-DataInfoSentToConsumerCreate")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataInfoSentToConsumer := resources.DataInfoSentToConsumer{}
	c.invokeCreate(w, r, rest_resources.SetDataInfoSentToConsumerFunc, dataInfoSentToConsumer, rest_resources.CreateResourceSuccessMsgPattern, rest_resources.CreateResourceErrorMsgPattern)
}
