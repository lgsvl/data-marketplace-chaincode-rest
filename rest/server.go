package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	_ "github.com/lgsvl/data-marketplace-chaincode-rest/statik"
	"github.com/rakyll/statik/fs"
)

// CCServer represents the chaincode rest server
type CCServer struct {
	ctx       context.Context
	logger    *log.Logger
	ccHandler *CCHandler
	config    resources.ServerConfig
}

//NewDataStreamServer creates a new instance of the DataStreamServer
func NewCCServer(ct context.Context, l *log.Logger, c *CCHandler, conf resources.ServerConfig) *CCServer {
	return &CCServer{ctx: ct, logger: l, ccHandler: c, config: conf}
}

//Start starts the http server
func (c *CCServer) Start() error {
	router := c.NewRouter()
	http.Handle("/", router)
	c.printStartMsg()

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/swaggerui/", staticServer)
	http.Handle("/swaggerui/", sh)

	return http.ListenAndServe(fmt.Sprintf(":%d", c.config.Port), nil)
}
func (c *CCServer) printStartMsg() {
	c.logger.Printf(fmt.Sprintf("Starting data marketplace chaincode rest on port %d ....", c.config.Port))
	c.logger.Printf("CTL-C to exit/stop data marketplace chaincode service")
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (c *CCServer) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := c.getRoutes()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func (c *CCServer) getRoutes() Routes {
	return Routes{
		Route{
			"Index",
			"GET",
			"/api/",
			Index,
		},

		Route{
			"BusinessCreate",
			strings.ToUpper("Post"),
			"/api/Business",
			c.ccHandler.BusinessCreate,
		},

		Route{
			"BusinessDeleteById",
			strings.ToUpper("Delete"),
			"/api/Business/{id}",
			c.ccHandler.BusinessDeleteById,
		},

		Route{
			"BusinessFindById",
			strings.ToUpper("Get"),
			"/api/Business/{id}",
			c.ccHandler.BusinessFindById,
		},

		Route{
			"BusinessReplaceById",
			strings.ToUpper("Put"),
			"/api/Business/{id}",
			c.ccHandler.BusinessReplaceById,
		},

		Route{
			"DataCategoryCreate",
			strings.ToUpper("Post"),
			"/api/DataCategory",
			c.ccHandler.DataCategoryCreate,
		},

		Route{
			"DataCategoryDeleteById",
			strings.ToUpper("Delete"),
			"/api/DataCategory/{id}",
			c.ccHandler.DataCategoryDeleteById,
		},

		Route{
			"DataCategoryFindById",
			strings.ToUpper("Get"),
			"/api/DataCategory/{id}",
			c.ccHandler.DataCategoryFindById,
		},

		Route{
			"DataCategoryReplaceById",
			strings.ToUpper("Put"),
			"/api/DataCategory/{id}",
			c.ccHandler.DataCategoryReplaceById,
		},

		Route{
			"DataContractDeleteById",
			strings.ToUpper("Delete"),
			"/api/DataContract/{id}",
			c.ccHandler.DataContractDeleteById,
		},

		Route{
			"DataContractFindById",
			strings.ToUpper("Get"),
			"/api/DataContract/{id}",
			c.ccHandler.DataContractFindById,
		},

		Route{
			"DataContractReplaceById",
			strings.ToUpper("Put"),
			"/api/DataContract/{id}",
			c.ccHandler.DataContractReplaceById,
		},

		Route{
			"DataContractTypeCreate",
			strings.ToUpper("Post"),
			"/api/DataContractType",
			c.ccHandler.DataContractTypeCreate,
		},

		Route{
			"DataContractTypeDeleteById",
			strings.ToUpper("Delete"),
			"/api/DataContractType/{id}",
			c.ccHandler.DataContractTypeDeleteById,
		},

		Route{
			"DataContractTypeFindById",
			strings.ToUpper("Get"),
			"/api/DataContractType/{id}",
			c.ccHandler.DataContractTypeFindById,
		},

		Route{
			"DataContractTypeReplaceById",
			strings.ToUpper("Put"),
			"/api/DataContractType/{id}",
			c.ccHandler.DataContractTypeReplaceById,
		},

		Route{
			"DataInfoSentToConsumerCreate",
			strings.ToUpper("Post"),
			"/api/DataInfoSentToConsumer",
			c.ccHandler.DataInfoSentToConsumerCreate,
		},

		Route{
			"DataReceivedByConsumerCreate",
			strings.ToUpper("Post"),
			"/api/DataReceivedByConsumer",
			c.ccHandler.DataReceivedByConsumerCreate,
		},

		Route{
			"PersonCreate",
			strings.ToUpper("Post"),
			"/api/Person",
			c.ccHandler.PersonCreate,
		},

		Route{
			"PersonDeleteById",
			strings.ToUpper("Delete"),
			"/api/Person/{id}",
			c.ccHandler.PersonDeleteById,
		},

		Route{
			"PersonFindById",
			strings.ToUpper("Get"),
			"/api/Person/{id}",
			c.ccHandler.PersonFindById,
		},

		Route{
			"PersonReplaceById",
			strings.ToUpper("Put"),
			"/api/Person/{id}",
			c.ccHandler.PersonReplaceById,
		},

		Route{
			"QueryGetBusinesses",
			strings.ToUpper("Get"),
			"/api/queries/getBusinesses",
			c.ccHandler.QueryGetBusinesses,
		},

		Route{
			"QueryGetBusinessesWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getBusinessesWithPagination",
			c.ccHandler.QueryGetBusinessesWithPagination,
		},

		Route{
			"QueryGetDataCategories",
			strings.ToUpper("Get"),
			"/api/queries/getDataCategories",
			c.ccHandler.QueryGetDataCategories,
		},

		Route{
			"QueryGetDataCategoriesWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getDataCategoriesWithPagination",
			c.ccHandler.QueryGetDataCategoriesWithPagination,
		},

		Route{
			"QueryGetDataContractTypes",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractTypes",
			c.ccHandler.QueryGetDataContractTypes,
		},

		Route{
			"QueryGetDataContractTypesAfterTimeStamp",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractTypesAfterTimeStamp",
			c.ccHandler.QueryGetDataContractTypesAfterTimeStamp,
		},

		Route{
			"QueryGetDataContractTypesByCategory",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractTypesByCategory",
			c.ccHandler.QueryGetDataContractTypesByCategory,
		},

		Route{
			"QueryGetDataContractTypesByCategoryWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractTypesByCategoryWithPagination",
			c.ccHandler.QueryGetDataContractTypesByCategoryWithPagination,
		},

		Route{
			"QueryGetDataContractTypesByProvider",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractTypesByProvider",
			c.ccHandler.QueryGetDataContractTypesByProvider,
		},

		Route{
			"QueryGetDataContractTypesByProviderWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractTypesByProviderWithPagination",
			c.ccHandler.QueryGetDataContractTypesByProviderWithPagination,
		},

		Route{
			"QueryGetDataContractTypesWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractTypesWithPagination",
			c.ccHandler.QueryGetDataContractTypesWithPagination,
		},

		Route{
			"QueryGetDataContracts",
			strings.ToUpper("Get"),
			"/api/queries/getDataContracts",
			c.ccHandler.QueryGetDataContracts,
		},

		Route{
			"QueryGetDataContractsByConsumer",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractsByConsumer",
			c.ccHandler.QueryGetDataContractsByConsumer,
		},

		Route{
			"QueryGetDataContractsByConsumerWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractsByConsumerWithPagination",
			c.ccHandler.QueryGetDataContractsByConsumerWithPagination,
		},

		Route{
			"QueryGetDataContractsByProvider",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractsByProvider",
			c.ccHandler.QueryGetDataContractsByProvider,
		},

		Route{
			"QueryGetDataContractsByProviderWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractsByProviderWithPagination",
			c.ccHandler.QueryGetDataContractsByProviderWithPagination,
		},

		Route{
			"QueryGetDataContractsWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/getDataContractsWithPagination",
			c.ccHandler.QueryGetDataContractsWithPagination,
		},

		Route{
			"QueryGetPopularDataCategories",
			strings.ToUpper("Get"),
			"/api/queries/getPopularDataCategories",
			c.ccHandler.QueryGetPopularDataCategories,
		},

		Route{
			"QueryGetPopularDataContractTypes",
			strings.ToUpper("Get"),
			"/api/queries/getPopularDataContractTypes",
			c.ccHandler.QueryGetPopularDataContractTypes,
		},

		Route{
			"QueryGetRecommendedDataContractType",
			strings.ToUpper("Get"),
			"/api/queries/getRecommendedDataContractType",
			c.ccHandler.QueryGetRecommendedDataContractType,
		},

		Route{
			"QuerySelectBusinessDataSetsPurchasedDownloaded",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsPurchasedDownloaded",
			c.ccHandler.QuerySelectBusinessDataSetsPurchasedDownloaded,
		},

		Route{
			"QuerySelectBusinessDataSetsPurchasedDownloadedWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsPurchasedDownloadedWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsPurchasedDownloadedWithPagination,
		},

		Route{
			"QuerySelectBusinessDataSetsPurchasedNotUploaded",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsPurchasedNotUploaded",
			c.ccHandler.QuerySelectBusinessDataSetsPurchasedNotUploaded,
		},

		Route{
			"QuerySelectBusinessDataSetsPurchasedNotUploadedWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsPurchasedNotUploadedWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsPurchasedNotUploadedWithPagination,
		},

		Route{
			"QuerySelectBusinessDataSetsPurchasedUploadedNotDownloaded",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsPurchasedUploadedNotDownloaded",
			c.ccHandler.QuerySelectBusinessDataSetsPurchasedUploadedNotDownloaded,
		},

		Route{
			"QuerySelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination,
		},

		Route{
			"QuerySelectBusinessDataSetsSoldAndDownloaded",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsSoldAndDownloaded",
			c.ccHandler.QuerySelectBusinessDataSetsSoldAndDownloaded,
		},

		Route{
			"QuerySelectBusinessDataSetsSoldAndDownloadedWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsSoldAndDownloadedWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsSoldAndDownloadedWithPagination,
		},

		Route{
			"QuerySelectBusinessDataSetsSoldShippedNotDownloaded",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsSoldShippedNotDownloaded",
			c.ccHandler.QuerySelectBusinessDataSetsSoldShippedNotDownloaded,
		},

		Route{
			"QuerySelectBusinessDataSetsSoldShippedNotDownloadedWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsSoldShippedNotDownloadedWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsSoldShippedNotDownloadedWithPagination,
		},
		Route{
			"QuerySelectBusinessDataSetsByDataContractType",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsByDataContractType",
			c.ccHandler.QuerySelectBusinessDataSetsByDataContractType,
		},

		Route{
			"QuerySelectBusinessDataSetsByDataContractTypeWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsByDataContractTypeWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsByDataContractTypeWithPagination,
		},
		Route{
			"QuerySelectBusinessDataSetsToUploadByDataContractType",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsToUploadByDataContractType",
			c.ccHandler.QuerySelectBusinessDataSetsToUploadByDataContractType,
		},

		Route{
			"QuerySelectBusinessDataSetsToUploadByDataContractTypeWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsToUploadByDataContractTypeWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsToUploadByDataContractTypeWithPagination,
		},

		Route{
			"QuerySelectBusinessDataSetsToUpload",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsToUpload",
			c.ccHandler.QuerySelectBusinessDataSetsToUpload,
		},

		Route{
			"QuerySelectBusinessDataSetsToUploadWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectBusinessDataSetsToUploadWithPagination",
			c.ccHandler.QuerySelectBusinessDataSetsToUploadWithPagination,
		},

		Route{
			"QuerySelectDataSetContractsToUpload",
			strings.ToUpper("Get"),
			"/api/queries/selectDataSetContractsToUpload",
			c.ccHandler.QuerySelectDataSetContractsToUpload,
		},

		Route{
			"QuerySelectDataSetContractsToUploadWithPagination",
			strings.ToUpper("Get"),
			"/api/queries/selectDataSetContractsToUploadWithPagination",
			c.ccHandler.QuerySelectDataSetContractsToUploadWithPagination,
		},

		Route{
			"QuerySelectNumberOfBusinessDataSetsToUpload",
			strings.ToUpper("Get"),
			"/api/queries/selectNumberOfBusinessDataSetsToUpload",
			c.ccHandler.QuerySelectNumberOfBusinessDataSetsToUpload,
		},

		Route{
			"ReviewCreate",
			strings.ToUpper("Post"),
			"/api/Review",
			c.ccHandler.ReviewCreate,
		},

		Route{
			"ReviewDeleteById",
			strings.ToUpper("Delete"),
			"/api/Review/{id}",
			c.ccHandler.ReviewDeleteById,
		},

		Route{
			"ReviewFindById",
			strings.ToUpper("Get"),
			"/api/Review/{id}",
			c.ccHandler.ReviewFindById,
		},

		Route{
			"ReviewReplaceById",
			strings.ToUpper("Put"),
			"/api/Review/{id}",
			c.ccHandler.ReviewReplaceById,
		},

		Route{
			"SubmitDataContractProposalCreate",
			strings.ToUpper("Post"),
			"/api/SubmitDataContractProposal",
			c.ccHandler.SubmitDataContractProposalCreate,
		},
	}
}
