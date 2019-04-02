package rest

import (
	"net/http"

	rest_resources "github.com/lgsvl/data-marketplace-chaincode-rest/resources"
	"github.com/lgsvl/data-marketplace-chaincode-rest/utils"
	"github.com/lgsvl/data-marketplace-chaincode/resources"
)

func (c *CCHandler) QueryGetBusinesses(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetBusinesses")
	defer c.logger.Printf("ccHandler-exiting-QueryGetBusinesses")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c.invokeQuery(w, r, rest_resources.GetBusinessesFunc, rest_resources.QuerySuccessMsgPattern, []resources.Business{}, nil)
}

func (c *CCHandler) QueryGetBusinessesWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetBusinessesWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetBusinessesWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetBusinessesWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.Business{}, args)

}

func (c *CCHandler) QueryGetDataCategories(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataCategories")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataCategories")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c.invokeQuery(w, r, rest_resources.GetDataCategoriesFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataCategory{}, nil)
}

func (c *CCHandler) QueryGetDataCategoriesWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataCategoriesWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataCategoriesWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetDataCategoriesWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataCategory{}, args)
}

func (c *CCHandler) QueryGetDataContractTypes(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractTypes")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractTypes")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c.invokeQuery(w, r, rest_resources.GetDataContractTypesFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, nil)
}

func (c *CCHandler) QueryGetDataContractTypesAfterTimeStamp(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractTypesAfterTimeStamp")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractTypesAfterTimeStamp")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	timestamp := utils.ExtractParamsFromURL(r, "timestamp")

	args := []string{timestamp}

	c.invokeQuery(w, r, rest_resources.GetDataContractTypesAfterTimeStampFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, args)
}

func (c *CCHandler) QueryGetDataContractTypesWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractTypesWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractTypesWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetDataContractTypesWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, args)
}

func (c *CCHandler) QueryGetDataContractTypesByCategory(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractTypesByCategory")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractTypesByCategory")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	categoryID := utils.ExtractParamsFromURL(r, "categoryID")

	args := []string{categoryID}

	c.invokeQuery(w, r, rest_resources.GetDataContractTypesByCategoryFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, args)
}

func (c *CCHandler) QueryGetDataContractTypesByCategoryWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractTypesByCategoryWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractTypesByCategoryWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	categoryID := utils.ExtractParamsFromURL(r, "categoryID")
	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{categoryID, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetDataContractTypesByCategoryWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, args)
}

func (c *CCHandler) QueryGetDataContractTypesByProvider(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractTypesByProvider")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractTypesByProvider")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")

	args := []string{providerID}

	c.invokeQuery(w, r, rest_resources.GetDataContractTypesByProviderFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, args)
}

func (c *CCHandler) QueryGetDataContractTypesByProviderWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractTypesByProviderWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractTypesByProviderWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")
	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{providerID, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetDataContractTypesByProviderFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, args)
}

func (c *CCHandler) QueryGetDataContracts(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContracts")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContracts")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c.invokeQuery(w, r, rest_resources.GetDataContractsFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, nil)
}

func (c *CCHandler) QueryGetDataContractsWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractsWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractsWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetDataContractsWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QueryGetDataContractsByConsumer(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractsByConsumer")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractsByConsumer")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")

	args := []string{consumerID}

	c.invokeQuery(w, r, rest_resources.GetDataContractsByConsumerFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QueryGetDataContractsByConsumerWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractsByConsumerWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractsByConsumerWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")
	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{consumerID, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetDataContractsByConsumerWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QueryGetDataContractsByProvider(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractsByProvider")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractsByProvider")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")

	args := []string{providerID}

	c.invokeQuery(w, r, rest_resources.GetDataContractsByProviderFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QueryGetDataContractsByProviderWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetDataContractsByProviderWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QueryGetDataContractsByProviderWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")
	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{providerID, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.GetDataContractsByProviderWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QueryGetPopularDataCategories(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetPopularDataCategories")
	defer c.logger.Printf("ccHandler-exiting-QueryGetPopularDataCategories")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	size := utils.ExtractParamsFromURL(r, "size")

	args := []string{size}

	c.invokeQuery(w, r, rest_resources.GetPopularDataCategoriesFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataCategory{}, args)
}

func (c *CCHandler) QueryGetPopularDataContractTypes(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetPopularDataContractTypes")
	defer c.logger.Printf("ccHandler-exiting-QueryGetPopularDataContractTypes")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	size := utils.ExtractParamsFromURL(r, "size")

	args := []string{size}

	c.invokeQuery(w, r, rest_resources.GetPopularDataContractTypesFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, args)
}

func (c *CCHandler) QueryGetRecommendedDataContractType(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QueryGetRecommendedDataContractTypes")
	defer c.logger.Printf("ccHandler-exiting-QueryGetRecommendedDataContractTypes")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c.invokeQuery(w, r, rest_resources.GetRecommendedDataContractTypesFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContractType{}, nil)
}

func (c *CCHandler) QuerySelectBusinessDataSetsPurchasedDownloaded(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsPurchasedDownloaded")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsPurchasedDownloaded")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")
	today := utils.ExtractParamsFromURL(r, "today")

	args := []string{consumerID, today}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsPurchasedDownloadedFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsPurchasedDownloadedWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsPurchasedDownloadedWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsPurchasedDownloadedWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")
	today := utils.ExtractParamsFromURL(r, "today")
	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{consumerID, today, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsPurchasedDownloadedWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsPurchasedNotUploaded(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsPurchasedNotUploaded")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsPurchasedNotUploaded")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")

	args := []string{consumerID}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsPurchasedNotUploadedFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsPurchasedNotUploadedWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsPurchasedNotUploadedWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsPurchasedNotUploadedWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")
	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{consumerID, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsPurchasedNotUploadedWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsPurchasedUploadedNotDownloaded(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsPurchasedUploadedNotDownloaded")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsPurchasedUploadedNotDownloaded")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")
	today := utils.ExtractParamsFromURL(r, "today")

	args := []string{consumerID, today}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsPurchasedUploadedNotDownloadedFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	consumerID := utils.ExtractParamsFromURL(r, "consumerID")
	today := utils.ExtractParamsFromURL(r, "today")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{consumerID, today, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsPurchasedUploadedNotDownloadedWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsSoldAndDownloaded(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsSoldAndDownloaded")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsSoldAndDownloaded")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")
	today := utils.ExtractParamsFromURL(r, "today")

	args := []string{providerID, today}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsSoldAndDownloadedFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsSoldAndDownloadedWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsSoldAndDownloadedWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsSoldAndDownloadedWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")
	today := utils.ExtractParamsFromURL(r, "today")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{providerID, today, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsSoldAndDownloadedWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsSoldShippedNotDownloaded(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsSoldShippedNotDownloaded")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsSoldShippedNotDownloaded")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")
	today := utils.ExtractParamsFromURL(r, "today")

	args := []string{providerID, today}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsSoldShippedNotDownloadedFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsSoldShippedNotDownloadedWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsSoldShippedNotDownloadedWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsSoldShippedNotDownloadedWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")
	today := utils.ExtractParamsFromURL(r, "today")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{providerID, today, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsSoldShippedNotDownloadedWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsToUpload(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsToUpload")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsToUpload")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")

	args := []string{providerID}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsToUploadFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsToUploadWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsToUploadWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsToUploadWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{providerID, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsToUploadWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsByDataContractType(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsByDataContractType")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsByDataContractType")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractType := utils.ExtractParamsFromURL(r, "dataContractType")

	args := []string{dataContractType}

	c.invokeQuery(w, r, rest_resources.SelectDataContractsByDataContractTypeFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsByDataContractTypeWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsByDataContractTypeWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsByDataContractTypeWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractType := utils.ExtractParamsFromURL(r, "dataContractType")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{dataContractType, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectDataContractsByDataContractTypeWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsToUploadByDataContractType(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsToUploadByDataContractType")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsToUploadByDataContractType")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractType := utils.ExtractParamsFromURL(r, "dataContractType")

	args := []string{dataContractType}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsToUploadByDataContractTypeFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectBusinessDataSetsToUploadByDataContractTypeWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectBusinessDataSetsToUploadByDataContractTypeWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectBusinessDataSetsToUploadByDataContractTypeWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractType := utils.ExtractParamsFromURL(r, "dataContractType")

	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{dataContractType, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectBusinessDataSetsToUploadByDataContractTypeWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectDataSetContractsToUpload(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectDataSetContractsToUpload")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectDataSetContractsToUpload")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractTypeID := utils.ExtractParamsFromURL(r, "dataContractTypeID")

	args := []string{dataContractTypeID}

	c.invokeQuery(w, r, rest_resources.SelectDataSetContractsToUploadFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectDataSetContractsToUploadWithPagination(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectDataSetContractsToUploadWithPagination")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectDataSetContractsToUploadWithPagination")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dataContractTypeID := utils.ExtractParamsFromURL(r, "dataContractTypeID")
	pageSize := utils.ExtractParamsFromURL(r, "pageSize")
	bookmark := utils.ExtractParamsFromURL(r, "bookmark")

	args := []string{dataContractTypeID, pageSize, bookmark}

	c.invokeQuery(w, r, rest_resources.SelectDataSetContractsToUploadWithPaginationFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}

func (c *CCHandler) QuerySelectNumberOfBusinessDataSetsToUpload(w http.ResponseWriter, r *http.Request) {
	c.logger.Printf("ccHandler-entering-QuerySelectNumberOfBusinessDataSetsToUpload")
	defer c.logger.Printf("ccHandler-exiting-QuerySelectNumberOfBusinessDataSetsToUpload")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	providerID := utils.ExtractParamsFromURL(r, "providerID")

	args := []string{providerID}

	c.invokeQuery(w, r, rest_resources.SelectNumberOfBusinessDataSetsToUploadFunc, rest_resources.QuerySuccessMsgPattern, []resources.DataContract{}, args)
}
