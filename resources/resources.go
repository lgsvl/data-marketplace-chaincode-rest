package resources

const (
	CreateBusinessFunc             string = "createBusiness"
	GetBusinessFunc                string = "getBusiness"
	CreateDataCategoryFunc         string = "createDataCategory"
	GetDataCategoryFunc            string = "getDataCategory"
	CreateDataContractTypeFunc     string = "createDataContractType"
	GetDataContractTypeFunc        string = "getDataContractType"
	SubmitDataContractProposalFunc string = "submitDataContractProposal"
	GetDataContractFunc            string = "getDataContract"
	SubmitReviewFunc               string = "submitReview"
	GetReviewFunc                  string = "getReview"
	AddPersonFunc                  string = "addPerson"
	GetPersonFunc                  string = "getPerson"
	SetDataInfoSentToConsumerFunc  string = "setDataInfoSentToConsumer"
	SetDataReceivedByConsumerFunc  string = "setDataReceivedByConsumer"
	// QUERIES
	GetBusinessesFunc                                                      string = "getBusinesses"
	GetBusinessesWithPaginationFunc                                        string = "getBusinessesWithPagination"
	GetDataCategoriesFunc                                                  string = "getDataCategories"
	GetDataCategoriesWithPaginationFunc                                    string = "getDataCategoriesWithPagination"
	GetPopularDataCategoriesFunc                                           string = "getPopularDataCategories"
	GetPopularDataContractTypesFunc                                        string = "getPopularDataContractTypes"
	GetRecommendedDataContractTypesFunc                                    string = "getRecommendedDataContractType"
	GetDataContractTypesFunc                                               string = "getDataContractTypes"
	GetDataContractTypesAfterTimeStampFunc                                 string = "getDataContractTypesAfterTimeStamp"
	GetDataContractTypesWithPaginationFunc                                 string = "getDataContractTypesWithPagination"
	GetDataContractTypesByCategoryFunc                                     string = "getDataContractTypesByCategory"
	GetDataContractTypesByCategoryWithPaginationFunc                       string = "getDataContractTypesByCategoryWithPagination"
	GetDataContractTypesByProviderFunc                                     string = "getDataContractTypesByProvider"
	GetDataContractTypesByProviderWithPaginationFunc                       string = "getDataContractTypesByProviderWithPagination"
	SelectNumberOfBusinessDataSetsToUploadFunc                             string = "selectNumberOfBusinessDataSetsToUpload"
	GetDataContractsFunc                                                   string = "getDataContracts"
	GetDataContractsWithPaginationFunc                                     string = "getDataContractsWithPagination"
	GetDataContractsByProviderFunc                                         string = "getDataContractsByProvider"
	GetDataContractsByProviderWithPaginationFunc                           string = "getDataContractsByProviderWithPagination"
	GetDataContractsByConsumerFunc                                         string = "getDataContractsByConsumer"
	GetDataContractsByConsumerWithPaginationFunc                           string = "getDataContractsByConsumerWithPagination"
	SelectDataSetContractsToUploadFunc                                     string = "selectDataSetContractsToUpload"
	SelectDataSetContractsToUploadWithPaginationFunc                       string = "selectDataSetContractsToUploadWithPagination"
	SelectBusinessDataSetsToUploadFunc                                     string = "selectBusinessDataSetsToUpload"
	SelectBusinessDataSetsToUploadWithPaginationFunc                       string = "selectBusinessDataSetsToUploadWithPagination"
	SelectBusinessDataSetsToUploadByDataContractTypeFunc                   string = "selectBusinessDataSetsToUploadByDataContractType"
	SelectDataContractsByDataContractTypeFunc                              string = "selectDataContractsByDataContractType"
	SelectDataContractsByDataContractTypeWithPaginationFunc                string = "selectDataContractsByDataContractTypeWithPagination"
	SelectBusinessDataSetsToUploadByDataContractTypeWithPaginationFunc     string = "selectBusinessDataSetsToUploadWithByDataContractTypePagination"
	SelectBusinessDataSetsSoldShippedNotDownloadedFunc                     string = "selectBusinessDataSetsSoldShippedNotDownloaded"
	SelectBusinessDataSetsSoldShippedNotDownloadedWithPaginationFunc       string = "selectBusinessDataSetsSoldShippedNotDownloadedWithPagination"
	SelectBusinessDataSetsSoldAndDownloadedFunc                            string = "selectBusinessDataSetsSoldAndDownloaded"
	SelectBusinessDataSetsSoldAndDownloadedWithPaginationFunc              string = "selectBusinessDataSetsSoldAndDownloadedWithPagination"
	SelectBusinessDataSetsPurchasedNotUploadedFunc                         string = "selectBusinessDataSetsPurchasedNotUploaded"
	SelectBusinessDataSetsPurchasedNotUploadedWithPaginationFunc           string = "selectBusinessDataSetsPurchasedNotUploadedWithPagination"
	SelectBusinessDataSetsPurchasedUploadedNotDownloadedFunc               string = "selectBusinessDataSetsPurchasedUploadedNotDownloaded"
	SelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPaginationFunc string = "selectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination"
	SelectBusinessDataSetsPurchasedDownloadedFunc                          string = "selectBusinessDataSetsPurchasedDownloaded"
	SelectBusinessDataSetsPurchasedDownloadedWithPaginationFunc            string = "selectBusinessDataSetsPurchasedDownloadedWithPagination"
	GetReviewsFunc                                                         string = "getReviews"
	GetReviewsWithPaginationFunc                                           string = "getReviewsWithPagination"
	GetPersonsFunc                                                         string = "getPersons"
	GetPersonsWithPaginationFunc                                           string = "getPersons"

	CreateResourceSuccessMsgPattern string = "Chaincode invoke successful. result: status:200"
	GetResourceSuccessMsgPattern    string = "Chaincode invoke successful. result: status:200 payload:\""
	QuerySuccessMsgPattern          string = "Chaincode invoke successful. result: status:200 payload:\""
	CreateResourceErrorMsgPattern   string = "response: status:500 message:"
	GetResourceErrorMsgPattern      string = "response: status:500 message:"
	QueryErrorMsgPattern            string = "response: status:500 message:"
	ChaincodePeerInvokeCommand      string = "peer chaincode invoke -o %s -C %s -n %s -c '%s'"
)

//ServerConfig contains the configuration of the server
type ServerConfig struct {
	Port int `json:"port"`
}

type GenericError struct {
	Error error `json:"error"`
}

type Query struct {
	Args []string `json:"Args"`
}

type CreateResourceRequest struct {
	AuthorizationToken   string `json:"authorization"`
	ResourceString       string `json:"resource"`
	Function             string `json:"function"`
	ErrorStringPattern   string `json:"pattern"`
	SuccessStringPattern string `json:"pattern"`
}

type CreateResourceResponse struct {
	Response interface{} `json:"response"`
	Error    string      `json:"error"`
}

type GetResourceRequest struct {
	AuthorizationToken   string      `json:"authorization"`
	ID                   string      `json:"ID"`
	Function             string      `json:"function"`
	SuccessStringPattern string      `json:"pattern"`
	ErrorStringPattern   string      `json:"pattern"`
	Resource             interface{} `json:"resource"`
}

type GetResourceResponse struct {
	Response interface{} `json:"response"`
	Error    string      `json:"error"`
}

type QueryRequest struct {
	AuthorizationToken   string      `json:"authorization"`
	Function             string      `json:"function"`
	SuccessStringPattern string      `json:"pattern"`
	ErrorStringPattern   string      `json:"pattern"`
	ResponseContainer    interface{} `json:"resource"`
	Args                 []string    `json:"args"`
}

type QueryResponse struct {
	Response         interface{}        `json:"response"`
	ResponseMetadata []ResponseMetadata `json:"ResponseMetadata`
	Error            string             `json:"error"`
}

type GenericResponse struct {
	Response interface{} `json:"response"`
	Error    string      `json:"error"`
}

type ResponseMetadata struct {
	ResponseMetadata Metadata `json:"ResponseMetadata"`
}

type Metadata struct {
	RecordsCount string `json:"RecordsCount"`
	Bookmark     string `json:"Bookmark"`
}
