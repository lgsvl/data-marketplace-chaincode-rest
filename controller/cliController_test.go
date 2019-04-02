package controller_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lgsvl/data-marketplace-chaincode-rest/controller"
	"github.com/lgsvl/data-marketplace-chaincode-rest/fakes"
	"github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cli-controller", func() {
	var (
		logger         *log.Logger
		fakeExecutor   *fakes.FakeExecutor
		ctx            context.Context
		fabricSetup    controller.FabricSetup
		peerController controller.Controller
	)
	BeforeEach(func() {
		logger = log.New(os.Stdout, "data-stream-delivery-controller-test: ", log.Lshortfile|log.LstdFlags)
		ctx = context.Background()
		fakeExecutor = new(fakes.FakeExecutor)

		fabricSetup = controller.FabricSetup{
			// Network parameters
			OrdererURL: "fake-orderer:1111",
			// Channel parameters
			ChannelID: "fake-channel",
			// Chaincode parameters
			ChainCodeID: "dmp",
		}

		peerController = controller.NewPeerControllerWithExecutor(ctx, logger, fabricSetup, fakeExecutor)

	})

	Context(".InvokeCreateResource", func() {
		It("should fail when executor fails to execute command", func() {
			request := resources.CreateResourceRequest{
				Function:           "fakeFunc",
				ResourceString:     "{}",
				ErrorStringPattern: "error",
			}

			fakeExecutor.ExecuteReturns([]byte("fake-error-exec"), fmt.Errorf("fake-error-exec"))
			createResourceResponse := peerController.InvokeCreateResource(request)
			Expect(createResourceResponse.Error).NotTo(Equal(""))
		})

	})
})
