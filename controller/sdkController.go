package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
	"github.com/pkg/errors"
)

// FabricSetup implementation
type FabricSetup struct {
	ctx             context.Context
	logger          *log.Logger
	ConfigFile      string
	OrgID           string
	OrdererID       string
	OrdererURL      string
	ChannelID       string
	ChainCodeID     string
	initialized     bool
	ChannelConfig   string
	ChaincodeGoPath string
	ChaincodePath   string
	OrgAdmin        string
	OrgName         string
	UserName        string
	client          *channel.Client
	admin           *resmgmt.Client
	sdk             *fabsdk.FabricSDK
	event           *event.Client
}

// Initialize reads the configuration file and sets up the client, chain and event hub
func (setup *FabricSetup) Initialize() error {

	// Add parameters for the initialization
	if setup.initialized {
		return errors.New("sdk already initialized")
	}

	// Initialize the SDK with the configuration file
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK")
	}
	setup.sdk = sdk
	fmt.Println("SDK created")

	// The resource management client is responsible for managing channels (create/update channel)
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))
	if err != nil {
		return errors.WithMessage(err, "failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create channel management client from Admin identity")
	}
	setup.admin = resMgmtClient
	fmt.Println("Ressource management client created")

	// The MSP client allow us to retrieve user information from their identity, like its signing identity which we will need to save the channel
	// mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(setup.OrgName))
	// if err != nil {
	// 	return errors.WithMessage(err, "failed to create MSP client")
	// }
	// adminIdentity, err := mspClient.GetSigningIdentity(setup.OrgAdmin)
	// if err != nil {
	// 	return errors.WithMessage(err, "failed to get admin signing identity")
	// }

	// _ = resmgmt.SaveChannelRequest{ChannelID: setup.ChannelID, ChannelConfigPath: setup.ChannelConfig, SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	// // txID, err := setup.admin.SaveChannel(req, resmgmt.WithOrdererEndpoint(setup.OrdererID))
	// // if err != nil || txID.TransactionID == "" {
	// 	return errors.WithMessage(err, "failed to save channel")
	// }
	// fmt.Println("Channel created")

	//Make admin user join the previously created channel
	// if err = setup.admin.JoinChannel(setup.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(setup.OrdererID)); err != nil {
	// 	return errors.WithMessage(err, "failed to make admin join channel")
	// }
	// fmt.Println("Channel joined")
	// Channel client is used to query and execute transactions
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	setup.client, err = channel.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new channel client")
	}
	fmt.Println("Channel client created")

	// Creation of the client which will enables access to our channel events
	setup.event, err = event.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new event client")
	}

	fmt.Println("Event client created")

	fmt.Println("Initialization Successful")
	setup.initialized = true
	return nil
}

func (setup *FabricSetup) InstallAndInstantiateCC() error {

	// Create the chaincode package that will be sent to the peers
	ccPkg, err := packager.NewCCPackage(setup.ChaincodePath, setup.ChaincodeGoPath)
	if err != nil {
		return errors.WithMessage(err, "failed to create chaincode package")
	}
	fmt.Println("ccPkg created")

	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodePath, Version: "0", Package: ccPkg}
	_, err = setup.admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return errors.WithMessage(err, "failed to install chaincode")
	}
	fmt.Println("Chaincode installed")

	// Set up chaincode policy
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1.example.com"})

	resp, err := setup.admin.InstantiateCC(setup.ChannelID, resmgmt.InstantiateCCRequest{Name: setup.ChainCodeID, Path: setup.ChaincodeGoPath, Version: "0", Args: [][]byte{[]byte("init")}, Policy: ccPolicy})
	if err != nil || resp.TransactionID == "" {
		return errors.WithMessage(err, "failed to instantiate the chaincode")
	}
	fmt.Println("Chaincode instantiated")

	// Channel client is used to query and execute transactions
	clientContext := setup.sdk.ChannelContext(setup.ChannelID, fabsdk.WithUser(setup.UserName))
	setup.client, err = channel.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new channel client")
	}
	fmt.Println("Channel client created")

	// Creation of the client which will enables access to our channel events
	setup.event, err = event.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "failed to create new event client")
	}
	fmt.Println("Event client created")

	fmt.Println("Chaincode Installation & Instantiation Successful")
	return nil
}

func (setup *FabricSetup) InvokeGetBusiness(businessID string) (resources.Business, error) {
	var args []string
	// args = append(args, "invoke")
	args = append(args, "getBusiness")
	args = append(args, businessID)

	eventID := "getBusinessInvoke"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in InvokeGetBusiness invoke")

	reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return resources.Business{}, err
	}
	defer setup.event.Unregister(reg)

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}, TransientMap: transientDataMap})
	if err != nil {
		return resources.Business{}, err
	}

	// Wait for the result of the submission
	select {
	case ccEvent := <-notifier:
		fmt.Printf("getbusiness-cc-received-event-%v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return resources.Business{}, fmt.Errorf("did-NOT-receive-CC-event-for-eventId(%s)", eventID)
	}
	business := resources.Business{}
	json.Unmarshal(response.Payload, &business)

	return business, nil
}

//CloseSDK Closes the SDK
func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}
