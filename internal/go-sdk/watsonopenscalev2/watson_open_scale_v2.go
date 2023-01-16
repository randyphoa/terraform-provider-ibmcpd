/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.55.1-b24c7487-20220831-201343
 */

// Package watsonopenscalev2 : Operations and models for the WatsonOpenScaleV2 service
package watsonopenscalev2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	common "terraform-provider-ibmcpd/internal/go-sdk/common"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// WatsonOpenScaleV2 : Watson OpenScale API Specification
//
// API Version: 2.0.0
// See: https://cloud.ibm.com/docs/services/ai-openscale/index.html
type WatsonOpenScaleV2 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ai_openscale"

const ParameterizedServiceURL = "https://aiopenscale.cloud.ibm.com/openscale/{serviceInstanceId}"

var defaultUrlVariables = map[string]string{
	"serviceInstanceId": "{service-instance-id}",
}

// WatsonOpenScaleV2Options : Service options
type WatsonOpenScaleV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewWatsonOpenScaleV2UsingExternalConfig : constructs an instance of WatsonOpenScaleV2 with passed in options and external configuration.
func NewWatsonOpenScaleV2UsingExternalConfig(options *WatsonOpenScaleV2Options) (watsonOpenScale *WatsonOpenScaleV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	watsonOpenScale, err = NewWatsonOpenScaleV2(options)
	if err != nil {
		return
	}

	err = watsonOpenScale.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = watsonOpenScale.Service.SetServiceURL(options.URL)
	}
	return
}

// NewWatsonOpenScaleV2 : constructs an instance of WatsonOpenScaleV2 with passed in options.
func NewWatsonOpenScaleV2(options *WatsonOpenScaleV2Options) (service *WatsonOpenScaleV2, err error) {
	serviceOptions := &core.ServiceOptions{
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &WatsonOpenScaleV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "watsonOpenScale" suitable for processing requests.
func (watsonOpenScale *WatsonOpenScaleV2) Clone() *WatsonOpenScaleV2 {
	if core.IsNil(watsonOpenScale) {
		return nil
	}
	clone := *watsonOpenScale
	clone.Service = watsonOpenScale.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (watsonOpenScale *WatsonOpenScaleV2) SetServiceURL(url string) error {
	return watsonOpenScale.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (watsonOpenScale *WatsonOpenScaleV2) GetServiceURL() string {
	return watsonOpenScale.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (watsonOpenScale *WatsonOpenScaleV2) SetDefaultHeaders(headers http.Header) {
	watsonOpenScale.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (watsonOpenScale *WatsonOpenScaleV2) SetEnableGzipCompression(enableGzip bool) {
	watsonOpenScale.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (watsonOpenScale *WatsonOpenScaleV2) GetEnableGzipCompression() bool {
	return watsonOpenScale.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (watsonOpenScale *WatsonOpenScaleV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	watsonOpenScale.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (watsonOpenScale *WatsonOpenScaleV2) DisableRetries() {
	watsonOpenScale.Service.DisableRetries()
}

// DataMartsList : List all data marts
// The method returns the data mart confugrations as an object.
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsList(dataMartsListOptions *DataMartsListOptions) (result *DataMartDatabaseResponseCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataMartsListWithContext(context.Background(), dataMartsListOptions)
}

// DataMartsListWithContext is an alternate form of the DataMartsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsListWithContext(ctx context.Context, dataMartsListOptions *DataMartsListOptions) (result *DataMartDatabaseResponseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(dataMartsListOptions, "dataMartsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_marts`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataMartsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataMartsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataMartDatabaseResponseCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DataMartsAdd : Create a new data mart
// Create a new data mart with the given database connection.
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsAdd(dataMartsAddOptions *DataMartsAddOptions) (result *DataMartDatabaseResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataMartsAddWithContext(context.Background(), dataMartsAddOptions)
}

// DataMartsAddWithContext is an alternate form of the DataMartsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsAddWithContext(ctx context.Context, dataMartsAddOptions *DataMartsAddOptions) (result *DataMartDatabaseResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataMartsAddOptions, "dataMartsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataMartsAddOptions, "dataMartsAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_marts`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataMartsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataMartsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if dataMartsAddOptions.Force != nil {
		builder.AddQuery("force", fmt.Sprint(*dataMartsAddOptions.Force))
	}

	body := make(map[string]interface{})
	if dataMartsAddOptions.DatabaseConfiguration != nil {
		body["database_configuration"] = dataMartsAddOptions.DatabaseConfiguration
	}
	if dataMartsAddOptions.DatabaseDiscovery != nil {
		body["database_discovery"] = dataMartsAddOptions.DatabaseDiscovery
	}
	if dataMartsAddOptions.Description != nil {
		body["description"] = dataMartsAddOptions.Description
	}
	if dataMartsAddOptions.InternalDatabase != nil {
		body["internal_database"] = dataMartsAddOptions.InternalDatabase
	}
	if dataMartsAddOptions.Name != nil {
		body["name"] = dataMartsAddOptions.Name
	}
	if dataMartsAddOptions.ServiceInstanceCrn != nil {
		body["service_instance_crn"] = dataMartsAddOptions.ServiceInstanceCrn
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataMartDatabaseResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DataMartsDelete : Delete a data mart
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsDelete(dataMartsDeleteOptions *DataMartsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataMartsDeleteWithContext(context.Background(), dataMartsDeleteOptions)
}

// DataMartsDeleteWithContext is an alternate form of the DataMartsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsDeleteWithContext(ctx context.Context, dataMartsDeleteOptions *DataMartsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataMartsDeleteOptions, "dataMartsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataMartsDeleteOptions, "dataMartsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_mart_id": *dataMartsDeleteOptions.DataMartID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_marts/{data_mart_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataMartsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataMartsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if dataMartsDeleteOptions.Force != nil {
		builder.AddQuery("force", fmt.Sprint(*dataMartsDeleteOptions.Force))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// DataMartsGet : Get data mart with the given id
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsGet(dataMartsGetOptions *DataMartsGetOptions) (result *DataMartDatabaseResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataMartsGetWithContext(context.Background(), dataMartsGetOptions)
}

// DataMartsGetWithContext is an alternate form of the DataMartsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsGetWithContext(ctx context.Context, dataMartsGetOptions *DataMartsGetOptions) (result *DataMartDatabaseResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataMartsGetOptions, "dataMartsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataMartsGetOptions, "dataMartsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_mart_id": *dataMartsGetOptions.DataMartID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_marts/{data_mart_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataMartsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataMartsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataMartDatabaseResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DataMartsPatch : Update a data mart
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsPatch(dataMartsPatchOptions *DataMartsPatchOptions) (result *DataMartDatabaseResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataMartsPatchWithContext(context.Background(), dataMartsPatchOptions)
}

// DataMartsPatchWithContext is an alternate form of the DataMartsPatch method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataMartsPatchWithContext(ctx context.Context, dataMartsPatchOptions *DataMartsPatchOptions) (result *DataMartDatabaseResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataMartsPatchOptions, "dataMartsPatchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataMartsPatchOptions, "dataMartsPatchOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_mart_id": *dataMartsPatchOptions.DataMartID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_marts/{data_mart_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataMartsPatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataMartsPatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(dataMartsPatchOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataMartDatabaseResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ServiceProvidersList : List service providers
// List assosiated Machine Learning service instances.
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersList(serviceProvidersListOptions *ServiceProvidersListOptions) (result *ServiceProviderResponseCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.ServiceProvidersListWithContext(context.Background(), serviceProvidersListOptions)
}

// ServiceProvidersListWithContext is an alternate form of the ServiceProvidersList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersListWithContext(ctx context.Context, serviceProvidersListOptions *ServiceProvidersListOptions) (result *ServiceProviderResponseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(serviceProvidersListOptions, "serviceProvidersListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/service_providers`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range serviceProvidersListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ServiceProvidersList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if serviceProvidersListOptions.ShowDeleted != nil {
		builder.AddQuery("show_deleted", fmt.Sprint(*serviceProvidersListOptions.ShowDeleted))
	}
	if serviceProvidersListOptions.ServiceType != nil {
		builder.AddQuery("service_type", fmt.Sprint(*serviceProvidersListOptions.ServiceType))
	}
	if serviceProvidersListOptions.InstanceID != nil {
		builder.AddQuery("instance_id", fmt.Sprint(*serviceProvidersListOptions.InstanceID))
	}
	if serviceProvidersListOptions.OperationalSpaceID != nil {
		builder.AddQuery("operational_space_id", fmt.Sprint(*serviceProvidersListOptions.OperationalSpaceID))
	}
	if serviceProvidersListOptions.DeploymentSpaceID != nil {
		builder.AddQuery("deployment_space_id", fmt.Sprint(*serviceProvidersListOptions.DeploymentSpaceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceProviderResponseCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ServiceProvidersAdd : Add service provider
// Assosiate external Machine Learning service instance with the OpenScale DataMart.
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersAdd(serviceProvidersAddOptions *ServiceProvidersAddOptions) (result *ServiceProviderResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.ServiceProvidersAddWithContext(context.Background(), serviceProvidersAddOptions)
}

// ServiceProvidersAddWithContext is an alternate form of the ServiceProvidersAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersAddWithContext(ctx context.Context, serviceProvidersAddOptions *ServiceProvidersAddOptions) (result *ServiceProviderResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(serviceProvidersAddOptions, "serviceProvidersAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(serviceProvidersAddOptions, "serviceProvidersAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/service_providers`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range serviceProvidersAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ServiceProvidersAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if serviceProvidersAddOptions.Credentials != nil {
		body["credentials"] = serviceProvidersAddOptions.Credentials
	}
	if serviceProvidersAddOptions.Name != nil {
		body["name"] = serviceProvidersAddOptions.Name
	}
	if serviceProvidersAddOptions.ServiceType != nil {
		body["service_type"] = serviceProvidersAddOptions.ServiceType
	}
	if serviceProvidersAddOptions.DeploymentSpaceID != nil {
		body["deployment_space_id"] = serviceProvidersAddOptions.DeploymentSpaceID
	}
	if serviceProvidersAddOptions.Description != nil {
		body["description"] = serviceProvidersAddOptions.Description
	}
	if serviceProvidersAddOptions.OperationalSpaceID != nil {
		body["operational_space_id"] = serviceProvidersAddOptions.OperationalSpaceID
	}
	if serviceProvidersAddOptions.RequestHeaders != nil {
		body["request_headers"] = serviceProvidersAddOptions.RequestHeaders
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceProviderResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ServiceProvidersDelete : Delete a service provider
// Detach Machine Learning service provider.
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersDelete(serviceProvidersDeleteOptions *ServiceProvidersDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.ServiceProvidersDeleteWithContext(context.Background(), serviceProvidersDeleteOptions)
}

// ServiceProvidersDeleteWithContext is an alternate form of the ServiceProvidersDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersDeleteWithContext(ctx context.Context, serviceProvidersDeleteOptions *ServiceProvidersDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(serviceProvidersDeleteOptions, "serviceProvidersDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(serviceProvidersDeleteOptions, "serviceProvidersDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"service_provider_id": *serviceProvidersDeleteOptions.ServiceProviderID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/service_providers/{service_provider_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range serviceProvidersDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ServiceProvidersDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if serviceProvidersDeleteOptions.Force != nil {
		builder.AddQuery("force", fmt.Sprint(*serviceProvidersDeleteOptions.Force))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// ServiceProvidersGet : Get a specific service provider
// Get the assosiated Machine Learning service provider details.
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersGet(serviceProvidersGetOptions *ServiceProvidersGetOptions) (result *ServiceProviderResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.ServiceProvidersGetWithContext(context.Background(), serviceProvidersGetOptions)
}

// ServiceProvidersGetWithContext is an alternate form of the ServiceProvidersGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersGetWithContext(ctx context.Context, serviceProvidersGetOptions *ServiceProvidersGetOptions) (result *ServiceProviderResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(serviceProvidersGetOptions, "serviceProvidersGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(serviceProvidersGetOptions, "serviceProvidersGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"service_provider_id": *serviceProvidersGetOptions.ServiceProviderID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/service_providers/{service_provider_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range serviceProvidersGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ServiceProvidersGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceProviderResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ServiceProvidersUpdate : Update a service provider
// Update existing service provider.
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersUpdate(serviceProvidersUpdateOptions *ServiceProvidersUpdateOptions) (result *ServiceProviderResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.ServiceProvidersUpdateWithContext(context.Background(), serviceProvidersUpdateOptions)
}

// ServiceProvidersUpdateWithContext is an alternate form of the ServiceProvidersUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ServiceProvidersUpdateWithContext(ctx context.Context, serviceProvidersUpdateOptions *ServiceProvidersUpdateOptions) (result *ServiceProviderResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(serviceProvidersUpdateOptions, "serviceProvidersUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(serviceProvidersUpdateOptions, "serviceProvidersUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"service_provider_id": *serviceProvidersUpdateOptions.ServiceProviderID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/service_providers/{service_provider_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range serviceProvidersUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ServiceProvidersUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(serviceProvidersUpdateOptions.PatchDocument)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceProviderResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SubscriptionsList : List subscriptions
// List subscriptions.
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsList(subscriptionsListOptions *SubscriptionsListOptions) (result *SubscriptionResponseCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.SubscriptionsListWithContext(context.Background(), subscriptionsListOptions)
}

// SubscriptionsListWithContext is an alternate form of the SubscriptionsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsListWithContext(ctx context.Context, subscriptionsListOptions *SubscriptionsListOptions) (result *SubscriptionResponseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(subscriptionsListOptions, "subscriptionsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/subscriptions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range subscriptionsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "SubscriptionsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if subscriptionsListOptions.DataMartID != nil {
		builder.AddQuery("data_mart_id", fmt.Sprint(*subscriptionsListOptions.DataMartID))
	}
	if subscriptionsListOptions.ServiceProviderID != nil {
		builder.AddQuery("service_provider_id", fmt.Sprint(*subscriptionsListOptions.ServiceProviderID))
	}
	if subscriptionsListOptions.AssetAssetID != nil {
		builder.AddQuery("asset.asset_id", fmt.Sprint(*subscriptionsListOptions.AssetAssetID))
	}
	if subscriptionsListOptions.DeploymentDeploymentID != nil {
		builder.AddQuery("deployment.deployment_id", fmt.Sprint(*subscriptionsListOptions.DeploymentDeploymentID))
	}
	if subscriptionsListOptions.DeploymentDeploymentType != nil {
		builder.AddQuery("deployment.deployment_type", fmt.Sprint(*subscriptionsListOptions.DeploymentDeploymentType))
	}
	if subscriptionsListOptions.IntegrationReferenceIntegratedSystemID != nil {
		builder.AddQuery("integration_reference.integrated_system_id", fmt.Sprint(*subscriptionsListOptions.IntegrationReferenceIntegratedSystemID))
	}
	if subscriptionsListOptions.IntegrationReferenceExternalID != nil {
		builder.AddQuery("integration_reference.external_id", fmt.Sprint(*subscriptionsListOptions.IntegrationReferenceExternalID))
	}
	if subscriptionsListOptions.RiskEvaluationStatusState != nil {
		builder.AddQuery("risk_evaluation_status.state", fmt.Sprint(*subscriptionsListOptions.RiskEvaluationStatusState))
	}
	if subscriptionsListOptions.ServiceProviderOperationalSpaceID != nil {
		builder.AddQuery("service_provider.operational_space_id", fmt.Sprint(*subscriptionsListOptions.ServiceProviderOperationalSpaceID))
	}
	if subscriptionsListOptions.PreProductionReferenceID != nil {
		builder.AddQuery("pre_production_reference_id", fmt.Sprint(*subscriptionsListOptions.PreProductionReferenceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSubscriptionResponseCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SubscriptionsAdd : Add a new subscription
// Add a new subscription to the model deployment.
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsAdd(subscriptionsAddOptions *SubscriptionsAddOptions) (result *SubscriptionResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.SubscriptionsAddWithContext(context.Background(), subscriptionsAddOptions)
}

// SubscriptionsAddWithContext is an alternate form of the SubscriptionsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsAddWithContext(ctx context.Context, subscriptionsAddOptions *SubscriptionsAddOptions) (result *SubscriptionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(subscriptionsAddOptions, "subscriptionsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(subscriptionsAddOptions, "subscriptionsAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/subscriptions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range subscriptionsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "SubscriptionsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if subscriptionsAddOptions.Asset != nil {
		body["asset"] = subscriptionsAddOptions.Asset
	}
	if subscriptionsAddOptions.DataMartID != nil {
		body["data_mart_id"] = subscriptionsAddOptions.DataMartID
	}
	if subscriptionsAddOptions.Deployment != nil {
		body["deployment"] = subscriptionsAddOptions.Deployment
	}
	if subscriptionsAddOptions.ServiceProviderID != nil {
		body["service_provider_id"] = subscriptionsAddOptions.ServiceProviderID
	}
	if subscriptionsAddOptions.AnalyticsEngine != nil {
		body["analytics_engine"] = subscriptionsAddOptions.AnalyticsEngine
	}
	if subscriptionsAddOptions.AssetProperties != nil {
		body["asset_properties"] = subscriptionsAddOptions.AssetProperties
	}
	if subscriptionsAddOptions.DataSources != nil {
		body["data_sources"] = subscriptionsAddOptions.DataSources
	}
	if subscriptionsAddOptions.RiskEvaluationStatus != nil {
		body["risk_evaluation_status"] = subscriptionsAddOptions.RiskEvaluationStatus
	}

	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)

	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSubscriptionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SubscriptionsDelete : Delete a subscription
// Delete a subscription.
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsDelete(subscriptionsDeleteOptions *SubscriptionsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.SubscriptionsDeleteWithContext(context.Background(), subscriptionsDeleteOptions)
}

// SubscriptionsDeleteWithContext is an alternate form of the SubscriptionsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsDeleteWithContext(ctx context.Context, subscriptionsDeleteOptions *SubscriptionsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(subscriptionsDeleteOptions, "subscriptionsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(subscriptionsDeleteOptions, "subscriptionsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"subscription_id": *subscriptionsDeleteOptions.SubscriptionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/subscriptions/{subscription_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range subscriptionsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "SubscriptionsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if subscriptionsDeleteOptions.Force != nil {
		builder.AddQuery("force", fmt.Sprint(*subscriptionsDeleteOptions.Force))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// SubscriptionsGet : Get a specific subscription
// Get a specific subscription.
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsGet(subscriptionsGetOptions *SubscriptionsGetOptions) (result *SubscriptionResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.SubscriptionsGetWithContext(context.Background(), subscriptionsGetOptions)
}

// SubscriptionsGetWithContext is an alternate form of the SubscriptionsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsGetWithContext(ctx context.Context, subscriptionsGetOptions *SubscriptionsGetOptions) (result *SubscriptionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(subscriptionsGetOptions, "subscriptionsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(subscriptionsGetOptions, "subscriptionsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"subscription_id": *subscriptionsGetOptions.SubscriptionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/subscriptions/{subscription_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range subscriptionsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "SubscriptionsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSubscriptionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SubscriptionsUpdate : Update a subscription
// Update existing asset (from ML service instance) subscription.
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsUpdate(subscriptionsUpdateOptions *SubscriptionsUpdateOptions) (result *SubscriptionResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.SubscriptionsUpdateWithContext(context.Background(), subscriptionsUpdateOptions)
}

// SubscriptionsUpdateWithContext is an alternate form of the SubscriptionsUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsUpdateWithContext(ctx context.Context, subscriptionsUpdateOptions *SubscriptionsUpdateOptions) (result *SubscriptionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(subscriptionsUpdateOptions, "subscriptionsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(subscriptionsUpdateOptions, "subscriptionsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"subscription_id": *subscriptionsUpdateOptions.SubscriptionID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/subscriptions/{subscription_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range subscriptionsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "SubscriptionsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(subscriptionsUpdateOptions.PatchDocument)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSubscriptionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SubscriptionsScore : Computes the bias mitigation/remediation for the specified model
// Computes the bias mitigation/remediation for the specified model. The fairness monitoring debias request payload
// details must be valid.
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsScore(subscriptionsScoreOptions *SubscriptionsScoreOptions) (result *FairnessMonitoringRemediation, response *core.DetailedResponse, err error) {
	return watsonOpenScale.SubscriptionsScoreWithContext(context.Background(), subscriptionsScoreOptions)
}

// SubscriptionsScoreWithContext is an alternate form of the SubscriptionsScore method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsScoreWithContext(ctx context.Context, subscriptionsScoreOptions *SubscriptionsScoreOptions) (result *FairnessMonitoringRemediation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(subscriptionsScoreOptions, "subscriptionsScoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(subscriptionsScoreOptions, "subscriptionsScoreOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"subscription_id": *subscriptionsScoreOptions.SubscriptionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/subscriptions/{subscription_id}/predictions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range subscriptionsScoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "SubscriptionsScore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if subscriptionsScoreOptions.Values != nil {
		body["values"] = subscriptionsScoreOptions.Values
	}
	if subscriptionsScoreOptions.Fields != nil {
		body["fields"] = subscriptionsScoreOptions.Fields
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFairnessMonitoringRemediation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SubscriptionsSchemas : Derive model schemas from the training data
// Derive model schemas from the training data. Only "structured" input data type is supported. If the input_data_type
// field in the subscription (subscription -> entity -> asset -> input_data_type) is not "structured", an error will be
// returned.
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsSchemas(subscriptionsSchemasOptions *SubscriptionsSchemasOptions) (result *SchemaInferenceResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.SubscriptionsSchemasWithContext(context.Background(), subscriptionsSchemasOptions)
}

// SubscriptionsSchemasWithContext is an alternate form of the SubscriptionsSchemas method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) SubscriptionsSchemasWithContext(ctx context.Context, subscriptionsSchemasOptions *SubscriptionsSchemasOptions) (result *SchemaInferenceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(subscriptionsSchemasOptions, "subscriptionsSchemasOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(subscriptionsSchemasOptions, "subscriptionsSchemasOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"subscription_id": *subscriptionsSchemasOptions.SubscriptionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/subscriptions/{subscription_id}/schemas`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range subscriptionsSchemasOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "SubscriptionsSchemas")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if subscriptionsSchemasOptions.InputData != nil {
		body["input_data"] = subscriptionsSchemasOptions.InputData
	}
	if subscriptionsSchemasOptions.TrainingDataReference != nil {
		body["training_data_reference"] = subscriptionsSchemasOptions.TrainingDataReference
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSchemaInferenceResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DataSetsList : List all data sets specified by the parameters
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsList(dataSetsListOptions *DataSetsListOptions) (result *DataSetResponseCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataSetsListWithContext(context.Background(), dataSetsListOptions)
}

// DataSetsListWithContext is an alternate form of the DataSetsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsListWithContext(ctx context.Context, dataSetsListOptions *DataSetsListOptions) (result *DataSetResponseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(dataSetsListOptions, "dataSetsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataSetsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataSetsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if dataSetsListOptions.TargetTargetID != nil {
		builder.AddQuery("target.target_id", fmt.Sprint(*dataSetsListOptions.TargetTargetID))
	}
	if dataSetsListOptions.TargetTargetType != nil {
		builder.AddQuery("target.target_type", fmt.Sprint(*dataSetsListOptions.TargetTargetType))
	}
	if dataSetsListOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*dataSetsListOptions.Type))
	}
	if dataSetsListOptions.ManagedBy != nil {
		builder.AddQuery("managed_by", fmt.Sprint(*dataSetsListOptions.ManagedBy))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataSetResponseCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DataSetsAdd : Create a new data set
// Create a new data set.
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsAdd(dataSetsAddOptions *DataSetsAddOptions) (result *DataSetResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataSetsAddWithContext(context.Background(), dataSetsAddOptions)
}

// DataSetsAddWithContext is an alternate form of the DataSetsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsAddWithContext(ctx context.Context, dataSetsAddOptions *DataSetsAddOptions) (result *DataSetResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataSetsAddOptions, "dataSetsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataSetsAddOptions, "dataSetsAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataSetsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataSetsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if dataSetsAddOptions.DataMartID != nil {
		body["data_mart_id"] = dataSetsAddOptions.DataMartID
	}
	if dataSetsAddOptions.DataSchema != nil {
		body["data_schema"] = dataSetsAddOptions.DataSchema
	}
	if dataSetsAddOptions.Name != nil {
		body["name"] = dataSetsAddOptions.Name
	}
	if dataSetsAddOptions.Target != nil {
		body["target"] = dataSetsAddOptions.Target
	}
	if dataSetsAddOptions.Type != nil {
		body["type"] = dataSetsAddOptions.Type
	}
	if dataSetsAddOptions.Description != nil {
		body["description"] = dataSetsAddOptions.Description
	}
	if dataSetsAddOptions.Location != nil {
		body["location"] = dataSetsAddOptions.Location
	}
	if dataSetsAddOptions.ManagedBy != nil {
		body["managed_by"] = dataSetsAddOptions.ManagedBy
	}
	if dataSetsAddOptions.SchemaUpdateMode != nil {
		body["schema_update_mode"] = dataSetsAddOptions.SchemaUpdateMode
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataSetResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DataSetsDelete : Delete a data set
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsDelete(dataSetsDeleteOptions *DataSetsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataSetsDeleteWithContext(context.Background(), dataSetsDeleteOptions)
}

// DataSetsDeleteWithContext is an alternate form of the DataSetsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsDeleteWithContext(ctx context.Context, dataSetsDeleteOptions *DataSetsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataSetsDeleteOptions, "dataSetsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataSetsDeleteOptions, "dataSetsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *dataSetsDeleteOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataSetsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataSetsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// DataSetsGet : Get data set with the given id
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsGet(dataSetsGetOptions *DataSetsGetOptions) (result *DataSetResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataSetsGetWithContext(context.Background(), dataSetsGetOptions)
}

// DataSetsGetWithContext is an alternate form of the DataSetsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsGetWithContext(ctx context.Context, dataSetsGetOptions *DataSetsGetOptions) (result *DataSetResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataSetsGetOptions, "dataSetsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataSetsGetOptions, "dataSetsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *dataSetsGetOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataSetsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataSetsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataSetResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DataSetsUpdate : Update a data set
// Update the data set.
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsUpdate(dataSetsUpdateOptions *DataSetsUpdateOptions) (result *DataSetResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DataSetsUpdateWithContext(context.Background(), dataSetsUpdateOptions)
}

// DataSetsUpdateWithContext is an alternate form of the DataSetsUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DataSetsUpdateWithContext(ctx context.Context, dataSetsUpdateOptions *DataSetsUpdateOptions) (result *DataSetResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(dataSetsUpdateOptions, "dataSetsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(dataSetsUpdateOptions, "dataSetsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *dataSetsUpdateOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range dataSetsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DataSetsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(dataSetsUpdateOptions.PatchDocument)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataSetResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RecordsQuery : Get data set records using record_id or transaction_id
// Get data set records with specific record_id or transaction_id.
func (watsonOpenScale *WatsonOpenScaleV2) RecordsQuery(recordsQueryOptions *RecordsQueryOptions) (result *DataSetRecords, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RecordsQueryWithContext(context.Background(), recordsQueryOptions)
}

// RecordsQueryWithContext is an alternate form of the RecordsQuery method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RecordsQueryWithContext(ctx context.Context, recordsQueryOptions *RecordsQueryOptions) (result *DataSetRecords, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recordsQueryOptions, "recordsQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recordsQueryOptions, "recordsQueryOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_set_records`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range recordsQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RecordsQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("data_set_type", fmt.Sprint(*recordsQueryOptions.DataSetType))
	if recordsQueryOptions.RecordID != nil {
		builder.AddQuery("record_id", strings.Join(recordsQueryOptions.RecordID, ","))
	}
	if recordsQueryOptions.TransactionID != nil {
		builder.AddQuery("transaction_id", strings.Join(recordsQueryOptions.TransactionID, ","))
	}
	if recordsQueryOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*recordsQueryOptions.Start))
	}
	if recordsQueryOptions.End != nil {
		builder.AddQuery("end", fmt.Sprint(*recordsQueryOptions.End))
	}
	if recordsQueryOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*recordsQueryOptions.Offset))
	}
	if recordsQueryOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*recordsQueryOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataSetRecords)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RecordsList : List data set records
// List data set records.
func (watsonOpenScale *WatsonOpenScaleV2) RecordsList(recordsListOptions *RecordsListOptions) (result RecordsListResponseIntf, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RecordsListWithContext(context.Background(), recordsListOptions)
}

// RecordsListWithContext is an alternate form of the RecordsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RecordsListWithContext(ctx context.Context, recordsListOptions *RecordsListOptions) (result RecordsListResponseIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recordsListOptions, "recordsListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recordsListOptions, "recordsListOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *recordsListOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range recordsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RecordsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if recordsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*recordsListOptions.Start))
	}
	if recordsListOptions.End != nil {
		builder.AddQuery("end", fmt.Sprint(*recordsListOptions.End))
	}
	if recordsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*recordsListOptions.Limit))
	}
	if recordsListOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*recordsListOptions.Offset))
	}
	if recordsListOptions.Annotations != nil {
		builder.AddQuery("annotations", strings.Join(recordsListOptions.Annotations, ","))
	}
	if recordsListOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*recordsListOptions.Filter))
	}
	if recordsListOptions.IncludeTotalCount != nil {
		builder.AddQuery("include_total_count", fmt.Sprint(*recordsListOptions.IncludeTotalCount))
	}
	if recordsListOptions.Format != nil {
		builder.AddQuery("format", fmt.Sprint(*recordsListOptions.Format))
	}
	if recordsListOptions.BinaryFormat != nil {
		builder.AddQuery("binary_format", fmt.Sprint(*recordsListOptions.BinaryFormat))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRecordsListResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RecordsPatch : Update data set records
// Update data set records.
func (watsonOpenScale *WatsonOpenScaleV2) RecordsPatch(recordsPatchOptions *RecordsPatchOptions) (result *Status, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RecordsPatchWithContext(context.Background(), recordsPatchOptions)
}

// RecordsPatchWithContext is an alternate form of the RecordsPatch method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RecordsPatchWithContext(ctx context.Context, recordsPatchOptions *RecordsPatchOptions) (result *Status, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recordsPatchOptions, "recordsPatchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recordsPatchOptions, "recordsPatchOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *recordsPatchOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range recordsPatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RecordsPatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(recordsPatchOptions.PatchDocument)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RecordsAdd : Add new data set records
// Add new data set records.
func (watsonOpenScale *WatsonOpenScaleV2) RecordsAdd(recordsAddOptions *RecordsAddOptions) (result *Status, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RecordsAddWithContext(context.Background(), recordsAddOptions)
}

// RecordsAddWithContext is an alternate form of the RecordsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RecordsAddWithContext(ctx context.Context, recordsAddOptions *RecordsAddOptions) (result *Status, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recordsAddOptions, "recordsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recordsAddOptions, "recordsAddOptions")
	if err != nil {
		return
	}

	if recordsAddOptions.DatasetRecordsPayloadItem != nil && recordsAddOptions.ContentType == nil {
		recordsAddOptions.SetContentType("application/json")
	}

	pathParamsMap := map[string]string{
		"data_set_id": *recordsAddOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range recordsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RecordsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if recordsAddOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*recordsAddOptions.ContentType))
	}

	if recordsAddOptions.Header != nil {
		builder.AddQuery("header", fmt.Sprint(*recordsAddOptions.Header))
	}
	if recordsAddOptions.Skip != nil {
		builder.AddQuery("skip", fmt.Sprint(*recordsAddOptions.Skip))
	}
	if recordsAddOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*recordsAddOptions.Limit))
	}
	if recordsAddOptions.Delimiter != nil {
		builder.AddQuery("delimiter", fmt.Sprint(*recordsAddOptions.Delimiter))
	}
	if recordsAddOptions.OnError != nil {
		builder.AddQuery("on_error", fmt.Sprint(*recordsAddOptions.OnError))
	}
	if recordsAddOptions.CsvMaxLineLength != nil {
		builder.AddQuery("csv_max_line_length", fmt.Sprint(*recordsAddOptions.CsvMaxLineLength))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(recordsAddOptions.ContentType), recordsAddOptions.DatasetRecordsPayloadItem, nil, recordsAddOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RecordsGet : Get a specific data set record with the given id
// Get a specific record in a data set.
func (watsonOpenScale *WatsonOpenScaleV2) RecordsGet(recordsGetOptions *RecordsGetOptions) (result *DataRecordResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RecordsGetWithContext(context.Background(), recordsGetOptions)
}

// RecordsGetWithContext is an alternate form of the RecordsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RecordsGetWithContext(ctx context.Context, recordsGetOptions *RecordsGetOptions) (result *DataRecordResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recordsGetOptions, "recordsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recordsGetOptions, "recordsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *recordsGetOptions.DataSetID,
		"record_id":   *recordsGetOptions.RecordID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/records/{record_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range recordsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RecordsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if recordsGetOptions.BinaryFormat != nil {
		builder.AddQuery("binary_format", fmt.Sprint(*recordsGetOptions.BinaryFormat))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataRecordResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RecordsUpdate : Update a specific record in a data set
// Update a specific record in a data set.
func (watsonOpenScale *WatsonOpenScaleV2) RecordsUpdate(recordsUpdateOptions *RecordsUpdateOptions) (result *DataRecordResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RecordsUpdateWithContext(context.Background(), recordsUpdateOptions)
}

// RecordsUpdateWithContext is an alternate form of the RecordsUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RecordsUpdateWithContext(ctx context.Context, recordsUpdateOptions *RecordsUpdateOptions) (result *DataRecordResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recordsUpdateOptions, "recordsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recordsUpdateOptions, "recordsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *recordsUpdateOptions.DataSetID,
		"record_id":   *recordsUpdateOptions.RecordID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/records/{record_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range recordsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RecordsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if recordsUpdateOptions.BinaryFormat != nil {
		builder.AddQuery("binary_format", fmt.Sprint(*recordsUpdateOptions.BinaryFormat))
	}

	_, err = builder.SetBodyContentJSON(recordsUpdateOptions.PatchDocument)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataRecordResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RecordsField : Get value of a field in a given record
// Get value of a field in a given record.
func (watsonOpenScale *WatsonOpenScaleV2) RecordsField(recordsFieldOptions *RecordsFieldOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.RecordsFieldWithContext(context.Background(), recordsFieldOptions)
}

// RecordsFieldWithContext is an alternate form of the RecordsField method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RecordsFieldWithContext(ctx context.Context, recordsFieldOptions *RecordsFieldOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(recordsFieldOptions, "recordsFieldOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(recordsFieldOptions, "recordsFieldOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *recordsFieldOptions.DataSetID,
		"record_id":   *recordsFieldOptions.RecordID,
		"field_name":  *recordsFieldOptions.FieldName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/records/{record_id}/{field_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range recordsFieldOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RecordsField")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// RequestsGet : Get status of a specific request
// Get status of a specific request.
func (watsonOpenScale *WatsonOpenScaleV2) RequestsGet(requestsGetOptions *RequestsGetOptions) (result *Status, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RequestsGetWithContext(context.Background(), requestsGetOptions)
}

// RequestsGetWithContext is an alternate form of the RequestsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RequestsGetWithContext(ctx context.Context, requestsGetOptions *RequestsGetOptions) (result *Status, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(requestsGetOptions, "requestsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(requestsGetOptions, "requestsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *requestsGetOptions.DataSetID,
		"request_id":  *requestsGetOptions.RequestID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/requests/{request_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range requestsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RequestsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DistributionsDelete : Delete data distributions
// Delete data distribution.
func (watsonOpenScale *WatsonOpenScaleV2) DistributionsDelete(distributionsDeleteOptions *DistributionsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.DistributionsDeleteWithContext(context.Background(), distributionsDeleteOptions)
}

// DistributionsDeleteWithContext is an alternate form of the DistributionsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DistributionsDeleteWithContext(ctx context.Context, distributionsDeleteOptions *DistributionsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(distributionsDeleteOptions, "distributionsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(distributionsDeleteOptions, "distributionsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *distributionsDeleteOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/distributions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range distributionsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DistributionsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// DistributionsAdd : add new data disbributions
// add new data disbributions.
func (watsonOpenScale *WatsonOpenScaleV2) DistributionsAdd(distributionsAddOptions *DistributionsAddOptions) (result *DataDistributionResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DistributionsAddWithContext(context.Background(), distributionsAddOptions)
}

// DistributionsAddWithContext is an alternate form of the DistributionsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DistributionsAddWithContext(ctx context.Context, distributionsAddOptions *DistributionsAddOptions) (result *DataDistributionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(distributionsAddOptions, "distributionsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(distributionsAddOptions, "distributionsAddOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id": *distributionsAddOptions.DataSetID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/distributions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range distributionsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DistributionsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if distributionsAddOptions.Nocache != nil {
		builder.AddQuery("nocache", fmt.Sprint(*distributionsAddOptions.Nocache))
	}

	body := make(map[string]interface{})
	if distributionsAddOptions.Dataset != nil {
		body["dataset"] = distributionsAddOptions.Dataset
	}
	if distributionsAddOptions.End != nil {
		body["end"] = distributionsAddOptions.End
	}
	if distributionsAddOptions.Group != nil {
		body["group"] = distributionsAddOptions.Group
	}
	if distributionsAddOptions.Start != nil {
		body["start"] = distributionsAddOptions.Start
	}
	if distributionsAddOptions.Agg != nil {
		body["agg"] = distributionsAddOptions.Agg
	}
	if distributionsAddOptions.Filter != nil {
		body["filter"] = distributionsAddOptions.Filter
	}
	if distributionsAddOptions.Limit != nil {
		body["limit"] = distributionsAddOptions.Limit
	}
	if distributionsAddOptions.MaxBins != nil {
		body["max_bins"] = distributionsAddOptions.MaxBins
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataDistributionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DistributionsGet : Get a specific data distribution
// Get a specific data distribution.
func (watsonOpenScale *WatsonOpenScaleV2) DistributionsGet(distributionsGetOptions *DistributionsGetOptions) (result *DataDistributionResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.DistributionsGetWithContext(context.Background(), distributionsGetOptions)
}

// DistributionsGetWithContext is an alternate form of the DistributionsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DistributionsGetWithContext(ctx context.Context, distributionsGetOptions *DistributionsGetOptions) (result *DataDistributionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(distributionsGetOptions, "distributionsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(distributionsGetOptions, "distributionsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_set_id":          *distributionsGetOptions.DataSetID,
		"data_distribution_id": *distributionsGetOptions.DataDistributionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/data_sets/{data_set_id}/distributions/{data_distribution_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range distributionsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DistributionsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataDistributionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MonitorsList : List available monitors
// List available monitors.
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsList(monitorsListOptions *MonitorsListOptions) (result *MonitorCollections, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MonitorsListWithContext(context.Background(), monitorsListOptions)
}

// MonitorsListWithContext is an alternate form of the MonitorsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsListWithContext(ctx context.Context, monitorsListOptions *MonitorsListOptions) (result *MonitorCollections, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(monitorsListOptions, "monitorsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range monitorsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MonitorsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if monitorsListOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*monitorsListOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorCollections)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MonitorsAdd : Add custom monitor
// Add custom monitor.
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsAdd(monitorsAddOptions *MonitorsAddOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MonitorsAddWithContext(context.Background(), monitorsAddOptions)
}

// MonitorsAddWithContext is an alternate form of the MonitorsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsAddWithContext(ctx context.Context, monitorsAddOptions *MonitorsAddOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(monitorsAddOptions, "monitorsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(monitorsAddOptions, "monitorsAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range monitorsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MonitorsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if monitorsAddOptions.Metrics != nil {
		body["metrics"] = monitorsAddOptions.Metrics
	}
	if monitorsAddOptions.Name != nil {
		body["name"] = monitorsAddOptions.Name
	}
	if monitorsAddOptions.Tags != nil {
		body["tags"] = monitorsAddOptions.Tags
	}
	if monitorsAddOptions.AppliesTo != nil {
		body["applies_to"] = monitorsAddOptions.AppliesTo
	}
	if monitorsAddOptions.Description != nil {
		body["description"] = monitorsAddOptions.Description
	}
	if monitorsAddOptions.ManagedBy != nil {
		body["managed_by"] = monitorsAddOptions.ManagedBy
	}
	if monitorsAddOptions.ParametersSchema != nil {
		body["parameters_schema"] = monitorsAddOptions.ParametersSchema
	}
	if monitorsAddOptions.Schedule != nil {
		body["schedule"] = monitorsAddOptions.Schedule
	}
	if monitorsAddOptions.Schedules != nil {
		body["schedules"] = monitorsAddOptions.Schedules
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorDisplayForm)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MonitorsDelete : Delete a monitor definition
// Delete a monitor definition.
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsDelete(monitorsDeleteOptions *MonitorsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.MonitorsDeleteWithContext(context.Background(), monitorsDeleteOptions)
}

// MonitorsDeleteWithContext is an alternate form of the MonitorsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsDeleteWithContext(ctx context.Context, monitorsDeleteOptions *MonitorsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(monitorsDeleteOptions, "monitorsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(monitorsDeleteOptions, "monitorsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_definition_id": *monitorsDeleteOptions.MonitorDefinitionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_definitions/{monitor_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range monitorsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MonitorsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// MonitorsGet : Get a specific monitor definition
// Get a specific monitor definition.
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsGet(monitorsGetOptions *MonitorsGetOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MonitorsGetWithContext(context.Background(), monitorsGetOptions)
}

// MonitorsGetWithContext is an alternate form of the MonitorsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsGetWithContext(ctx context.Context, monitorsGetOptions *MonitorsGetOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(monitorsGetOptions, "monitorsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(monitorsGetOptions, "monitorsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_definition_id": *monitorsGetOptions.MonitorDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_definitions/{monitor_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range monitorsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MonitorsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorDisplayForm)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MonitorsPatch : Update a monitor definition
// Update a monitor definition.
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsPatch(monitorsPatchOptions *MonitorsPatchOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MonitorsPatchWithContext(context.Background(), monitorsPatchOptions)
}

// MonitorsPatchWithContext is an alternate form of the MonitorsPatch method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsPatchWithContext(ctx context.Context, monitorsPatchOptions *MonitorsPatchOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(monitorsPatchOptions, "monitorsPatchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(monitorsPatchOptions, "monitorsPatchOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_definition_id": *monitorsPatchOptions.MonitorDefinitionID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_definitions/{monitor_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range monitorsPatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MonitorsPatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(monitorsPatchOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorDisplayForm)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MonitorsUpdate : Update the monitor definition
// Update a monitor definition.
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsUpdate(monitorsUpdateOptions *MonitorsUpdateOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MonitorsUpdateWithContext(context.Background(), monitorsUpdateOptions)
}

// MonitorsUpdateWithContext is an alternate form of the MonitorsUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MonitorsUpdateWithContext(ctx context.Context, monitorsUpdateOptions *MonitorsUpdateOptions) (result *MonitorDisplayForm, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(monitorsUpdateOptions, "monitorsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(monitorsUpdateOptions, "monitorsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_definition_id": *monitorsUpdateOptions.MonitorDefinitionID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_definitions/{monitor_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range monitorsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MonitorsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if monitorsUpdateOptions.Metrics != nil {
		body["metrics"] = monitorsUpdateOptions.Metrics
	}
	if monitorsUpdateOptions.Name != nil {
		body["name"] = monitorsUpdateOptions.Name
	}
	if monitorsUpdateOptions.Tags != nil {
		body["tags"] = monitorsUpdateOptions.Tags
	}
	if monitorsUpdateOptions.AppliesTo != nil {
		body["applies_to"] = monitorsUpdateOptions.AppliesTo
	}
	if monitorsUpdateOptions.Description != nil {
		body["description"] = monitorsUpdateOptions.Description
	}
	if monitorsUpdateOptions.ManagedBy != nil {
		body["managed_by"] = monitorsUpdateOptions.ManagedBy
	}
	if monitorsUpdateOptions.ParametersSchema != nil {
		body["parameters_schema"] = monitorsUpdateOptions.ParametersSchema
	}
	if monitorsUpdateOptions.Schedule != nil {
		body["schedule"] = monitorsUpdateOptions.Schedule
	}
	if monitorsUpdateOptions.Schedules != nil {
		body["schedules"] = monitorsUpdateOptions.Schedules
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorDisplayForm)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InstancesList : List monitor instances
// List monitor instances.
func (watsonOpenScale *WatsonOpenScaleV2) InstancesList(instancesListOptions *InstancesListOptions) (result *MonitorInstanceCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.InstancesListWithContext(context.Background(), instancesListOptions)
}

// InstancesListWithContext is an alternate form of the InstancesList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) InstancesListWithContext(ctx context.Context, instancesListOptions *InstancesListOptions) (result *MonitorInstanceCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(instancesListOptions, "instancesListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range instancesListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "InstancesList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if instancesListOptions.DataMartID != nil {
		builder.AddQuery("data_mart_id", fmt.Sprint(*instancesListOptions.DataMartID))
	}
	if instancesListOptions.MonitorDefinitionID != nil {
		builder.AddQuery("monitor_definition_id", fmt.Sprint(*instancesListOptions.MonitorDefinitionID))
	}
	if instancesListOptions.TargetTargetID != nil {
		builder.AddQuery("target.target_id", fmt.Sprint(*instancesListOptions.TargetTargetID))
	}
	if instancesListOptions.TargetTargetType != nil {
		builder.AddQuery("target.target_type", fmt.Sprint(*instancesListOptions.TargetTargetType))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorInstanceCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InstancesAdd : Create a new monitor instance
// Create a new monitor instance.
func (watsonOpenScale *WatsonOpenScaleV2) InstancesAdd(instancesAddOptions *InstancesAddOptions) (result *MonitorInstanceResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.InstancesAddWithContext(context.Background(), instancesAddOptions)
}

// InstancesAddWithContext is an alternate form of the InstancesAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) InstancesAddWithContext(ctx context.Context, instancesAddOptions *InstancesAddOptions) (result *MonitorInstanceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(instancesAddOptions, "instancesAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(instancesAddOptions, "instancesAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range instancesAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "InstancesAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if instancesAddOptions.SkipScheduler != nil {
		builder.AddQuery("skip_scheduler", fmt.Sprint(*instancesAddOptions.SkipScheduler))
	}

	body := make(map[string]interface{})
	if instancesAddOptions.DataMartID != nil {
		body["data_mart_id"] = instancesAddOptions.DataMartID
	}
	if instancesAddOptions.MonitorDefinitionID != nil {
		body["monitor_definition_id"] = instancesAddOptions.MonitorDefinitionID
	}
	if instancesAddOptions.Target != nil {
		body["target"] = instancesAddOptions.Target
	}
	if instancesAddOptions.ManagedBy != nil {
		body["managed_by"] = instancesAddOptions.ManagedBy
	}
	if instancesAddOptions.Parameters != nil {
		body["parameters"] = instancesAddOptions.Parameters
	}
	if instancesAddOptions.Schedule != nil {
		body["schedule"] = instancesAddOptions.Schedule
	}
	if instancesAddOptions.ScheduleID != nil {
		body["schedule_id"] = instancesAddOptions.ScheduleID
	}
	if instancesAddOptions.Schedules != nil {
		body["schedules"] = instancesAddOptions.Schedules
	}
	if instancesAddOptions.Thresholds != nil {
		body["thresholds"] = instancesAddOptions.Thresholds
	}
	if instancesAddOptions.TotalRecords != nil {
		body["total_records"] = instancesAddOptions.TotalRecords
	}
	if instancesAddOptions.UnprocessedRecords != nil {
		body["unprocessed_records"] = instancesAddOptions.UnprocessedRecords
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorInstanceResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InstancesDelete : Delete a monitor instance
// Delete a monitor instance.
func (watsonOpenScale *WatsonOpenScaleV2) InstancesDelete(instancesDeleteOptions *InstancesDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.InstancesDeleteWithContext(context.Background(), instancesDeleteOptions)
}

// InstancesDeleteWithContext is an alternate form of the InstancesDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) InstancesDeleteWithContext(ctx context.Context, instancesDeleteOptions *InstancesDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(instancesDeleteOptions, "instancesDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(instancesDeleteOptions, "instancesDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *instancesDeleteOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range instancesDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "InstancesDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// InstancesGet : Get monitor instance details
// Get monitor instance details.
func (watsonOpenScale *WatsonOpenScaleV2) InstancesGet(instancesGetOptions *InstancesGetOptions) (result *MonitorInstanceResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.InstancesGetWithContext(context.Background(), instancesGetOptions)
}

// InstancesGetWithContext is an alternate form of the InstancesGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) InstancesGetWithContext(ctx context.Context, instancesGetOptions *InstancesGetOptions) (result *MonitorInstanceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(instancesGetOptions, "instancesGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(instancesGetOptions, "instancesGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *instancesGetOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range instancesGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "InstancesGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if instancesGetOptions.Expand != nil {
		builder.AddQuery("expand", fmt.Sprint(*instancesGetOptions.Expand))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorInstanceResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InstancesUpdate : Update a monitor instance
// Update a monitor instance.
func (watsonOpenScale *WatsonOpenScaleV2) InstancesUpdate(instancesUpdateOptions *InstancesUpdateOptions) (result *MonitorInstanceResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.InstancesUpdateWithContext(context.Background(), instancesUpdateOptions)
}

// InstancesUpdateWithContext is an alternate form of the InstancesUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) InstancesUpdateWithContext(ctx context.Context, instancesUpdateOptions *InstancesUpdateOptions) (result *MonitorInstanceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(instancesUpdateOptions, "instancesUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(instancesUpdateOptions, "instancesUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *instancesUpdateOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range instancesUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "InstancesUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if instancesUpdateOptions.UpdateMetadataOnly != nil {
		builder.AddQuery("update_metadata_only", fmt.Sprint(*instancesUpdateOptions.UpdateMetadataOnly))
	}

	_, err = builder.SetBodyContentJSON(instancesUpdateOptions.PatchDocument)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorInstanceResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunsList : List monitoring runs
// List monitoring runs.
func (watsonOpenScale *WatsonOpenScaleV2) RunsList(runsListOptions *RunsListOptions) (result *MonitoringRunCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RunsListWithContext(context.Background(), runsListOptions)
}

// RunsListWithContext is an alternate form of the RunsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RunsListWithContext(ctx context.Context, runsListOptions *RunsListOptions) (result *MonitoringRunCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runsListOptions, "runsListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runsListOptions, "runsListOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *runsListOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/runs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RunsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if runsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*runsListOptions.Start))
	}
	if runsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*runsListOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitoringRunCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunsAdd : Trigger monitoring run
// Trigger monitoring run.
func (watsonOpenScale *WatsonOpenScaleV2) RunsAdd(runsAddOptions *RunsAddOptions) (result *MonitoringRun, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RunsAddWithContext(context.Background(), runsAddOptions)
}

// RunsAddWithContext is an alternate form of the RunsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RunsAddWithContext(ctx context.Context, runsAddOptions *RunsAddOptions) (result *MonitoringRun, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runsAddOptions, "runsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runsAddOptions, "runsAddOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *runsAddOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/runs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RunsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if runsAddOptions.BusinessMetricContext != nil {
		body["business_metric_context"] = runsAddOptions.BusinessMetricContext
	}
	if runsAddOptions.ExpirationDate != nil {
		body["expiration_date"] = runsAddOptions.ExpirationDate
	}
	if runsAddOptions.Parameters != nil {
		body["parameters"] = runsAddOptions.Parameters
	}
	if runsAddOptions.TriggeredBy != nil {
		body["triggered_by"] = runsAddOptions.TriggeredBy
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitoringRun)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunsGet : Get monitoring run details
// Get monitoring run details.
func (watsonOpenScale *WatsonOpenScaleV2) RunsGet(runsGetOptions *RunsGetOptions) (result *MonitoringRun, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RunsGetWithContext(context.Background(), runsGetOptions)
}

// RunsGetWithContext is an alternate form of the RunsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RunsGetWithContext(ctx context.Context, runsGetOptions *RunsGetOptions) (result *MonitoringRun, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runsGetOptions, "runsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runsGetOptions, "runsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *runsGetOptions.MonitorInstanceID,
		"monitoring_run_id":   *runsGetOptions.MonitoringRunID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/runs/{monitoring_run_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RunsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitoringRun)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RunsUpdate : Update existing monitoring run details
// Update existing monitoring run details.
func (watsonOpenScale *WatsonOpenScaleV2) RunsUpdate(runsUpdateOptions *RunsUpdateOptions) (result *MonitoringRun, response *core.DetailedResponse, err error) {
	return watsonOpenScale.RunsUpdateWithContext(context.Background(), runsUpdateOptions)
}

// RunsUpdateWithContext is an alternate form of the RunsUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) RunsUpdateWithContext(ctx context.Context, runsUpdateOptions *RunsUpdateOptions) (result *MonitoringRun, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runsUpdateOptions, "runsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runsUpdateOptions, "runsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *runsUpdateOptions.MonitorInstanceID,
		"monitoring_run_id":   *runsUpdateOptions.MonitoringRunID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/runs/{monitoring_run_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range runsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "RunsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(runsUpdateOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitoringRun)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MeasurementsQuery : Query for the recent measurement
// Query for the recent measurement grouped by the monitoring target (subscription or business application).
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsQuery(measurementsQueryOptions *MeasurementsQueryOptions) (result *MeasurementsResponseCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MeasurementsQueryWithContext(context.Background(), measurementsQueryOptions)
}

// MeasurementsQueryWithContext is an alternate form of the MeasurementsQuery method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsQueryWithContext(ctx context.Context, measurementsQueryOptions *MeasurementsQueryOptions) (result *MeasurementsResponseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(measurementsQueryOptions, "measurementsQueryOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/measurements`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range measurementsQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MeasurementsQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if measurementsQueryOptions.TargetID != nil {
		builder.AddQuery("target_id", fmt.Sprint(*measurementsQueryOptions.TargetID))
	}
	if measurementsQueryOptions.TargetType != nil {
		builder.AddQuery("target_type", fmt.Sprint(*measurementsQueryOptions.TargetType))
	}
	if measurementsQueryOptions.MonitorDefinitionID != nil {
		builder.AddQuery("monitor_definition_id", fmt.Sprint(*measurementsQueryOptions.MonitorDefinitionID))
	}
	if measurementsQueryOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*measurementsQueryOptions.Filter))
	}
	if measurementsQueryOptions.RecentCount != nil {
		builder.AddQuery("recent_count", fmt.Sprint(*measurementsQueryOptions.RecentCount))
	}
	if measurementsQueryOptions.Format != nil {
		builder.AddQuery("format", fmt.Sprint(*measurementsQueryOptions.Format))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMeasurementsResponseCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MeasurementsDelete : Delete measeurements for a given monitor instance
// Delete measeurements for a given monitor_instance_id.
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsDelete(measurementsDeleteOptions *MeasurementsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.MeasurementsDeleteWithContext(context.Background(), measurementsDeleteOptions)
}

// MeasurementsDeleteWithContext is an alternate form of the MeasurementsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsDeleteWithContext(ctx context.Context, measurementsDeleteOptions *MeasurementsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(measurementsDeleteOptions, "measurementsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(measurementsDeleteOptions, "measurementsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *measurementsDeleteOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/measurements`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range measurementsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MeasurementsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// MeasurementsList : Query measeurements from OpenScale DataMart
// Query measeurements from OpenScale DataMart. It is required to either provide a `start end` or `run_id` parameter.
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsList(measurementsListOptions *MeasurementsListOptions) (result *MonitorMeasurementResponseCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MeasurementsListWithContext(context.Background(), measurementsListOptions)
}

// MeasurementsListWithContext is an alternate form of the MeasurementsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsListWithContext(ctx context.Context, measurementsListOptions *MeasurementsListOptions) (result *MonitorMeasurementResponseCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(measurementsListOptions, "measurementsListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(measurementsListOptions, "measurementsListOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *measurementsListOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/measurements`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range measurementsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MeasurementsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("start", fmt.Sprint(*measurementsListOptions.Start))
	builder.AddQuery("end", fmt.Sprint(*measurementsListOptions.End))
	builder.AddQuery("run_id", fmt.Sprint(*measurementsListOptions.RunID))
	if measurementsListOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*measurementsListOptions.Filter))
	}
	if measurementsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*measurementsListOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorMeasurementResponseCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MeasurementsAdd : Publish measurement data to OpenScale
// Publish measurement data to OpenScale.
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsAdd(measurementsAddOptions *MeasurementsAddOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.MeasurementsAddWithContext(context.Background(), measurementsAddOptions)
}

// MeasurementsAddWithContext is an alternate form of the MeasurementsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsAddWithContext(ctx context.Context, measurementsAddOptions *MeasurementsAddOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(measurementsAddOptions, "measurementsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(measurementsAddOptions, "measurementsAddOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *measurementsAddOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/measurements`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range measurementsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MeasurementsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(measurementsAddOptions.MonitorMeasurementRequest)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// MeasurementsGet : Get measeurement data from OpenScale DataMart
// Get measeurement data from OpenScale DataMart.
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsGet(measurementsGetOptions *MeasurementsGetOptions) (result *MonitorMeasurementResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MeasurementsGetWithContext(context.Background(), measurementsGetOptions)
}

// MeasurementsGetWithContext is an alternate form of the MeasurementsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MeasurementsGetWithContext(ctx context.Context, measurementsGetOptions *MeasurementsGetOptions) (result *MonitorMeasurementResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(measurementsGetOptions, "measurementsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(measurementsGetOptions, "measurementsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *measurementsGetOptions.MonitorInstanceID,
		"measurement_id":      *measurementsGetOptions.MeasurementID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/measurements/{measurement_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range measurementsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MeasurementsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitorMeasurementResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MetricsList : Query monitor instance metrics from OpenScale DataMart
// Query monitor instance metrics from OpenScale DataMart. See <a
// href="https://github.ibm.com/aiopenscale/aios-datamart-service-api/wiki/1.3.-Metrics-Query-Language">Metrics Query
// Language documentation</a>.
func (watsonOpenScale *WatsonOpenScaleV2) MetricsList(metricsListOptions *MetricsListOptions) (result *DataMartGetMonitorInstanceMetrics, response *core.DetailedResponse, err error) {
	return watsonOpenScale.MetricsListWithContext(context.Background(), metricsListOptions)
}

// MetricsListWithContext is an alternate form of the MetricsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) MetricsListWithContext(ctx context.Context, metricsListOptions *MetricsListOptions) (result *DataMartGetMonitorInstanceMetrics, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(metricsListOptions, "metricsListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(metricsListOptions, "metricsListOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *metricsListOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitor_instances/{monitor_instance_id}/metrics`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range metricsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "MetricsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("start", fmt.Sprint(*metricsListOptions.Start))
	builder.AddQuery("end", fmt.Sprint(*metricsListOptions.End))
	builder.AddQuery("agg", fmt.Sprint(*metricsListOptions.Agg))
	if metricsListOptions.Interval != nil {
		builder.AddQuery("interval", fmt.Sprint(*metricsListOptions.Interval))
	}
	if metricsListOptions.Filter != nil {
		builder.AddQuery("filter", fmt.Sprint(*metricsListOptions.Filter))
	}
	if metricsListOptions.Group != nil {
		builder.AddQuery("group", fmt.Sprint(*metricsListOptions.Group))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDataMartGetMonitorInstanceMetrics)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BusinessApplicationsList : List business applications
// List configured business applications.
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsList(businessApplicationsListOptions *BusinessApplicationsListOptions) (result *BusinessApplicationsCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.BusinessApplicationsListWithContext(context.Background(), businessApplicationsListOptions)
}

// BusinessApplicationsListWithContext is an alternate form of the BusinessApplicationsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsListWithContext(ctx context.Context, businessApplicationsListOptions *BusinessApplicationsListOptions) (result *BusinessApplicationsCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(businessApplicationsListOptions, "businessApplicationsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/business_applications`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range businessApplicationsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "BusinessApplicationsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBusinessApplicationsCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BusinessApplicationsAdd : Add business application
// Add the new business application.
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsAdd(businessApplicationsAddOptions *BusinessApplicationsAddOptions) (result *BusinessApplicationResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.BusinessApplicationsAddWithContext(context.Background(), businessApplicationsAddOptions)
}

// BusinessApplicationsAddWithContext is an alternate form of the BusinessApplicationsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsAddWithContext(ctx context.Context, businessApplicationsAddOptions *BusinessApplicationsAddOptions) (result *BusinessApplicationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(businessApplicationsAddOptions, "businessApplicationsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(businessApplicationsAddOptions, "businessApplicationsAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/business_applications`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range businessApplicationsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "BusinessApplicationsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if businessApplicationsAddOptions.BusinessMetrics != nil {
		body["business_metrics"] = businessApplicationsAddOptions.BusinessMetrics
	}
	if businessApplicationsAddOptions.Description != nil {
		body["description"] = businessApplicationsAddOptions.Description
	}
	if businessApplicationsAddOptions.Name != nil {
		body["name"] = businessApplicationsAddOptions.Name
	}
	if businessApplicationsAddOptions.PayloadFields != nil {
		body["payload_fields"] = businessApplicationsAddOptions.PayloadFields
	}
	if businessApplicationsAddOptions.BusinessMetricsMonitorDefinitionID != nil {
		body["business_metrics_monitor_definition_id"] = businessApplicationsAddOptions.BusinessMetricsMonitorDefinitionID
	}
	if businessApplicationsAddOptions.BusinessMetricsMonitorInstanceID != nil {
		body["business_metrics_monitor_instance_id"] = businessApplicationsAddOptions.BusinessMetricsMonitorInstanceID
	}
	if businessApplicationsAddOptions.BusinessPayloadDataSetID != nil {
		body["business_payload_data_set_id"] = businessApplicationsAddOptions.BusinessPayloadDataSetID
	}
	if businessApplicationsAddOptions.CorrelationMonitorInstanceID != nil {
		body["correlation_monitor_instance_id"] = businessApplicationsAddOptions.CorrelationMonitorInstanceID
	}
	if businessApplicationsAddOptions.SubscriptionIds != nil {
		body["subscription_ids"] = businessApplicationsAddOptions.SubscriptionIds
	}
	if businessApplicationsAddOptions.TransactionBatchesDataSetID != nil {
		body["transaction_batches_data_set_id"] = businessApplicationsAddOptions.TransactionBatchesDataSetID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBusinessApplicationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BusinessApplicationsDelete : Delete business application
// Delete business application.
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsDelete(businessApplicationsDeleteOptions *BusinessApplicationsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.BusinessApplicationsDeleteWithContext(context.Background(), businessApplicationsDeleteOptions)
}

// BusinessApplicationsDeleteWithContext is an alternate form of the BusinessApplicationsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsDeleteWithContext(ctx context.Context, businessApplicationsDeleteOptions *BusinessApplicationsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(businessApplicationsDeleteOptions, "businessApplicationsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(businessApplicationsDeleteOptions, "businessApplicationsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"business_application_id": *businessApplicationsDeleteOptions.BusinessApplicationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/business_applications/{business_application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range businessApplicationsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "BusinessApplicationsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// BusinessApplicationsGet : Get application details
// Get business application details.
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsGet(businessApplicationsGetOptions *BusinessApplicationsGetOptions) (result *BusinessApplicationResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.BusinessApplicationsGetWithContext(context.Background(), businessApplicationsGetOptions)
}

// BusinessApplicationsGetWithContext is an alternate form of the BusinessApplicationsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsGetWithContext(ctx context.Context, businessApplicationsGetOptions *BusinessApplicationsGetOptions) (result *BusinessApplicationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(businessApplicationsGetOptions, "businessApplicationsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(businessApplicationsGetOptions, "businessApplicationsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"business_application_id": *businessApplicationsGetOptions.BusinessApplicationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/business_applications/{business_application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range businessApplicationsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "BusinessApplicationsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBusinessApplicationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BusinessApplicationsUpdate : Update business application
// Update business application.
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsUpdate(businessApplicationsUpdateOptions *BusinessApplicationsUpdateOptions) (result *BusinessApplicationResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.BusinessApplicationsUpdateWithContext(context.Background(), businessApplicationsUpdateOptions)
}

// BusinessApplicationsUpdateWithContext is an alternate form of the BusinessApplicationsUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) BusinessApplicationsUpdateWithContext(ctx context.Context, businessApplicationsUpdateOptions *BusinessApplicationsUpdateOptions) (result *BusinessApplicationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(businessApplicationsUpdateOptions, "businessApplicationsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(businessApplicationsUpdateOptions, "businessApplicationsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"business_application_id": *businessApplicationsUpdateOptions.BusinessApplicationID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/business_applications/{business_application_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range businessApplicationsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "BusinessApplicationsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(businessApplicationsUpdateOptions.PatchDocument)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBusinessApplicationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// IntegratedSystemsList : List integrated systems
// List integrated systems.
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsList(integratedSystemsListOptions *IntegratedSystemsListOptions) (result *IntegratedSystemCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.IntegratedSystemsListWithContext(context.Background(), integratedSystemsListOptions)
}

// IntegratedSystemsListWithContext is an alternate form of the IntegratedSystemsList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsListWithContext(ctx context.Context, integratedSystemsListOptions *IntegratedSystemsListOptions) (result *IntegratedSystemCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(integratedSystemsListOptions, "integratedSystemsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/integrated_systems`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range integratedSystemsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "IntegratedSystemsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntegratedSystemCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// IntegratedSystemsAdd : Create a new integrated system
// Create a new integrated system.
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsAdd(integratedSystemsAddOptions *IntegratedSystemsAddOptions) (result *IntegratedSystemResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.IntegratedSystemsAddWithContext(context.Background(), integratedSystemsAddOptions)
}

// IntegratedSystemsAddWithContext is an alternate form of the IntegratedSystemsAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsAddWithContext(ctx context.Context, integratedSystemsAddOptions *IntegratedSystemsAddOptions) (result *IntegratedSystemResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(integratedSystemsAddOptions, "integratedSystemsAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(integratedSystemsAddOptions, "integratedSystemsAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/integrated_systems`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range integratedSystemsAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "IntegratedSystemsAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if integratedSystemsAddOptions.Credentials != nil {
		body["credentials"] = integratedSystemsAddOptions.Credentials
	}
	if integratedSystemsAddOptions.Description != nil {
		body["description"] = integratedSystemsAddOptions.Description
	}
	if integratedSystemsAddOptions.Name != nil {
		body["name"] = integratedSystemsAddOptions.Name
	}
	if integratedSystemsAddOptions.Type != nil {
		body["type"] = integratedSystemsAddOptions.Type
	}
	if integratedSystemsAddOptions.Connection != nil {
		body["connection"] = integratedSystemsAddOptions.Connection
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntegratedSystemResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// IntegratedSystemsDelete : Delete an integrated system
// Delete an integrated system.
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsDelete(integratedSystemsDeleteOptions *IntegratedSystemsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.IntegratedSystemsDeleteWithContext(context.Background(), integratedSystemsDeleteOptions)
}

// IntegratedSystemsDeleteWithContext is an alternate form of the IntegratedSystemsDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsDeleteWithContext(ctx context.Context, integratedSystemsDeleteOptions *IntegratedSystemsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(integratedSystemsDeleteOptions, "integratedSystemsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(integratedSystemsDeleteOptions, "integratedSystemsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"integrated_system_id": *integratedSystemsDeleteOptions.IntegratedSystemID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/integrated_systems/{integrated_system_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range integratedSystemsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "IntegratedSystemsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// IntegratedSystemsGet : Get a specific integrated system
// Get a specific integrated system.
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsGet(integratedSystemsGetOptions *IntegratedSystemsGetOptions) (result *IntegratedSystemResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.IntegratedSystemsGetWithContext(context.Background(), integratedSystemsGetOptions)
}

// IntegratedSystemsGetWithContext is an alternate form of the IntegratedSystemsGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsGetWithContext(ctx context.Context, integratedSystemsGetOptions *IntegratedSystemsGetOptions) (result *IntegratedSystemResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(integratedSystemsGetOptions, "integratedSystemsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(integratedSystemsGetOptions, "integratedSystemsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"integrated_system_id": *integratedSystemsGetOptions.IntegratedSystemID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/integrated_systems/{integrated_system_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range integratedSystemsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "IntegratedSystemsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntegratedSystemResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// IntegratedSystemsUpdate : Update an integrated system
// Update an integrated system.
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsUpdate(integratedSystemsUpdateOptions *IntegratedSystemsUpdateOptions) (result *IntegratedSystemResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.IntegratedSystemsUpdateWithContext(context.Background(), integratedSystemsUpdateOptions)
}

// IntegratedSystemsUpdateWithContext is an alternate form of the IntegratedSystemsUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) IntegratedSystemsUpdateWithContext(ctx context.Context, integratedSystemsUpdateOptions *IntegratedSystemsUpdateOptions) (result *IntegratedSystemResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(integratedSystemsUpdateOptions, "integratedSystemsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(integratedSystemsUpdateOptions, "integratedSystemsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"integrated_system_id": *integratedSystemsUpdateOptions.IntegratedSystemID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/integrated_systems/{integrated_system_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range integratedSystemsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "IntegratedSystemsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(integratedSystemsUpdateOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalIntegratedSystemResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// OperationalSpacesList : List Operational Spaces
// List Operational Spaces.
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesList(operationalSpacesListOptions *OperationalSpacesListOptions) (result *OperationalSpaceCollection, response *core.DetailedResponse, err error) {
	return watsonOpenScale.OperationalSpacesListWithContext(context.Background(), operationalSpacesListOptions)
}

// OperationalSpacesListWithContext is an alternate form of the OperationalSpacesList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesListWithContext(ctx context.Context, operationalSpacesListOptions *OperationalSpacesListOptions) (result *OperationalSpaceCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(operationalSpacesListOptions, "operationalSpacesListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/operational_spaces`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range operationalSpacesListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "OperationalSpacesList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOperationalSpaceCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// OperationalSpacesAdd : Create an operational space
// Create an operational space.
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesAdd(operationalSpacesAddOptions *OperationalSpacesAddOptions) (result *OperationalSpaceResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.OperationalSpacesAddWithContext(context.Background(), operationalSpacesAddOptions)
}

// OperationalSpacesAddWithContext is an alternate form of the OperationalSpacesAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesAddWithContext(ctx context.Context, operationalSpacesAddOptions *OperationalSpacesAddOptions) (result *OperationalSpaceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(operationalSpacesAddOptions, "operationalSpacesAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(operationalSpacesAddOptions, "operationalSpacesAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/operational_spaces`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range operationalSpacesAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "OperationalSpacesAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if operationalSpacesAddOptions.Name != nil {
		body["name"] = operationalSpacesAddOptions.Name
	}
	if operationalSpacesAddOptions.Description != nil {
		body["description"] = operationalSpacesAddOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOperationalSpaceResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// OperationalSpacesDelete : Delete an operational space
// Delete an operational space.
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesDelete(operationalSpacesDeleteOptions *OperationalSpacesDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.OperationalSpacesDeleteWithContext(context.Background(), operationalSpacesDeleteOptions)
}

// OperationalSpacesDeleteWithContext is an alternate form of the OperationalSpacesDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesDeleteWithContext(ctx context.Context, operationalSpacesDeleteOptions *OperationalSpacesDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(operationalSpacesDeleteOptions, "operationalSpacesDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(operationalSpacesDeleteOptions, "operationalSpacesDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"operational_space_id": *operationalSpacesDeleteOptions.OperationalSpaceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/operational_spaces/{operational_space_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range operationalSpacesDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "OperationalSpacesDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// OperationalSpacesGet : Get an operational space
// Get an operational space.
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesGet(operationalSpacesGetOptions *OperationalSpacesGetOptions) (result *OperationalSpaceResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.OperationalSpacesGetWithContext(context.Background(), operationalSpacesGetOptions)
}

// OperationalSpacesGetWithContext is an alternate form of the OperationalSpacesGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesGetWithContext(ctx context.Context, operationalSpacesGetOptions *OperationalSpacesGetOptions) (result *OperationalSpaceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(operationalSpacesGetOptions, "operationalSpacesGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(operationalSpacesGetOptions, "operationalSpacesGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"operational_space_id": *operationalSpacesGetOptions.OperationalSpaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/operational_spaces/{operational_space_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range operationalSpacesGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "OperationalSpacesGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOperationalSpaceResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// OperationalSpacesUpdate : Update an operational space
// Update an operational space.
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesUpdate(operationalSpacesUpdateOptions *OperationalSpacesUpdateOptions) (result *OperationalSpaceResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.OperationalSpacesUpdateWithContext(context.Background(), operationalSpacesUpdateOptions)
}

// OperationalSpacesUpdateWithContext is an alternate form of the OperationalSpacesUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) OperationalSpacesUpdateWithContext(ctx context.Context, operationalSpacesUpdateOptions *OperationalSpacesUpdateOptions) (result *OperationalSpaceResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(operationalSpacesUpdateOptions, "operationalSpacesUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(operationalSpacesUpdateOptions, "operationalSpacesUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"operational_space_id": *operationalSpacesUpdateOptions.OperationalSpaceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/operational_spaces/{operational_space_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range operationalSpacesUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "OperationalSpacesUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(operationalSpacesUpdateOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOperationalSpaceResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UserPreferencesList : Get User Preferences
// Get User Preferences.
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesList(userPreferencesListOptions *UserPreferencesListOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.UserPreferencesListWithContext(context.Background(), userPreferencesListOptions)
}

// UserPreferencesListWithContext is an alternate form of the UserPreferencesList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesListWithContext(ctx context.Context, userPreferencesListOptions *UserPreferencesListOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(userPreferencesListOptions, "userPreferencesListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/user_preferences`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range userPreferencesListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "UserPreferencesList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// UserPreferencesPatch : Update User Preferences
// Update User Preferences.
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesPatch(userPreferencesPatchOptions *UserPreferencesPatchOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.UserPreferencesPatchWithContext(context.Background(), userPreferencesPatchOptions)
}

// UserPreferencesPatchWithContext is an alternate form of the UserPreferencesPatch method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesPatchWithContext(ctx context.Context, userPreferencesPatchOptions *UserPreferencesPatchOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(userPreferencesPatchOptions, "userPreferencesPatchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(userPreferencesPatchOptions, "userPreferencesPatchOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/user_preferences`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range userPreferencesPatchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "UserPreferencesPatch")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(userPreferencesPatchOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// UserPreferencesDelete : Delete the user preference
// Delete the user preference.
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesDelete(userPreferencesDeleteOptions *UserPreferencesDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.UserPreferencesDeleteWithContext(context.Background(), userPreferencesDeleteOptions)
}

// UserPreferencesDeleteWithContext is an alternate form of the UserPreferencesDelete method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesDeleteWithContext(ctx context.Context, userPreferencesDeleteOptions *UserPreferencesDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(userPreferencesDeleteOptions, "userPreferencesDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(userPreferencesDeleteOptions, "userPreferencesDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"user_preference_key": *userPreferencesDeleteOptions.UserPreferenceKey,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/user_preferences/{user_preference_key}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range userPreferencesDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "UserPreferencesDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// UserPreferencesGet : Get a specific user prefrence
// Get a specific user preference.
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesGet(userPreferencesGetOptions *UserPreferencesGetOptions) (result UserPreferencesGetResponseIntf, response *core.DetailedResponse, err error) {
	return watsonOpenScale.UserPreferencesGetWithContext(context.Background(), userPreferencesGetOptions)
}

// UserPreferencesGetWithContext is an alternate form of the UserPreferencesGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesGetWithContext(ctx context.Context, userPreferencesGetOptions *UserPreferencesGetOptions) (result UserPreferencesGetResponseIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(userPreferencesGetOptions, "userPreferencesGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(userPreferencesGetOptions, "userPreferencesGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"user_preference_key": *userPreferencesGetOptions.UserPreferenceKey,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/user_preferences/{user_preference_key}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range userPreferencesGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "UserPreferencesGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUserPreferencesGetResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UserPreferencesUpdate : Update the user preference
// Update the user preference.
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesUpdate(userPreferencesUpdateOptions *UserPreferencesUpdateOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.UserPreferencesUpdateWithContext(context.Background(), userPreferencesUpdateOptions)
}

// UserPreferencesUpdateWithContext is an alternate form of the UserPreferencesUpdate method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) UserPreferencesUpdateWithContext(ctx context.Context, userPreferencesUpdateOptions *UserPreferencesUpdateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(userPreferencesUpdateOptions, "userPreferencesUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(userPreferencesUpdateOptions, "userPreferencesUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"user_preference_key": *userPreferencesUpdateOptions.UserPreferenceKey,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/user_preferences/{user_preference_key}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range userPreferencesUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "UserPreferencesUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(userPreferencesUpdateOptions.UserPreferencesUpdateRequest)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// ExplanationTasksList : List all explanations
// List of all the computed explanations.
func (watsonOpenScale *WatsonOpenScaleV2) ExplanationTasksList(explanationTasksListOptions *ExplanationTasksListOptions) (result *GetExplanationTasksResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.ExplanationTasksListWithContext(context.Background(), explanationTasksListOptions)
}

// ExplanationTasksListWithContext is an alternate form of the ExplanationTasksList method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ExplanationTasksListWithContext(ctx context.Context, explanationTasksListOptions *ExplanationTasksListOptions) (result *GetExplanationTasksResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(explanationTasksListOptions, "explanationTasksListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/explanation_tasks`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range explanationTasksListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ExplanationTasksList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if explanationTasksListOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*explanationTasksListOptions.Offset))
	}
	if explanationTasksListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*explanationTasksListOptions.Limit))
	}
	if explanationTasksListOptions.SubscriptionID != nil {
		builder.AddQuery("subscription_id", fmt.Sprint(*explanationTasksListOptions.SubscriptionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetExplanationTasksResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExplanationTasksAdd : Compute explanations
// Submit tasks for computing explanation of predictions.
func (watsonOpenScale *WatsonOpenScaleV2) ExplanationTasksAdd(explanationTasksAddOptions *ExplanationTasksAddOptions) (result *PostExplanationTaskResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.ExplanationTasksAddWithContext(context.Background(), explanationTasksAddOptions)
}

// ExplanationTasksAddWithContext is an alternate form of the ExplanationTasksAdd method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ExplanationTasksAddWithContext(ctx context.Context, explanationTasksAddOptions *ExplanationTasksAddOptions) (result *PostExplanationTaskResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(explanationTasksAddOptions, "explanationTasksAddOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(explanationTasksAddOptions, "explanationTasksAddOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/explanation_tasks`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range explanationTasksAddOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ExplanationTasksAdd")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if explanationTasksAddOptions.ScoringIds != nil {
		body["scoring_ids"] = explanationTasksAddOptions.ScoringIds
	}
	if explanationTasksAddOptions.ExplanationTypes != nil {
		body["explanation_types"] = explanationTasksAddOptions.ExplanationTypes
	}
	if explanationTasksAddOptions.InputRows != nil {
		body["input_rows"] = explanationTasksAddOptions.InputRows
	}
	if explanationTasksAddOptions.SubscriptionIds != nil {
		body["subscription_ids"] = explanationTasksAddOptions.SubscriptionIds
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPostExplanationTaskResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExplanationTasksGet : Get explanation
// Get explanation for the given explanation task id.
func (watsonOpenScale *WatsonOpenScaleV2) ExplanationTasksGet(explanationTasksGetOptions *ExplanationTasksGetOptions) (result *GetExplanationTaskResponse, response *core.DetailedResponse, err error) {
	return watsonOpenScale.ExplanationTasksGetWithContext(context.Background(), explanationTasksGetOptions)
}

// ExplanationTasksGetWithContext is an alternate form of the ExplanationTasksGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) ExplanationTasksGetWithContext(ctx context.Context, explanationTasksGetOptions *ExplanationTasksGetOptions) (result *GetExplanationTaskResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(explanationTasksGetOptions, "explanationTasksGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(explanationTasksGetOptions, "explanationTasksGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"explanation_task_id": *explanationTasksGetOptions.ExplanationTaskID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/explanation_tasks/{explanation_task_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range explanationTasksGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "ExplanationTasksGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if explanationTasksGetOptions.SubscriptionID != nil {
		builder.AddQuery("subscription_id", fmt.Sprint(*explanationTasksGetOptions.SubscriptionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonOpenScale.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetExplanationTaskResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DriftArchiveHead : Retrieves the Drift archive metadata
// API to retrieve the Drift archive metadata.
func (watsonOpenScale *WatsonOpenScaleV2) DriftArchiveHead(driftArchiveHeadOptions *DriftArchiveHeadOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.DriftArchiveHeadWithContext(context.Background(), driftArchiveHeadOptions)
}

// DriftArchiveHeadWithContext is an alternate form of the DriftArchiveHead method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DriftArchiveHeadWithContext(ctx context.Context, driftArchiveHeadOptions *DriftArchiveHeadOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(driftArchiveHeadOptions, "driftArchiveHeadOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(driftArchiveHeadOptions, "driftArchiveHeadOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *driftArchiveHeadOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.HEAD)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitoring_services/drift/monitor_instances/{monitor_instance_id}/archives`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range driftArchiveHeadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DriftArchiveHead")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// DriftArchivePost : Upload Drift archives
// API to upload drift archive such as the Drift Detection Model.
func (watsonOpenScale *WatsonOpenScaleV2) DriftArchivePost(driftArchivePostOptions *DriftArchivePostOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.DriftArchivePostWithContext(context.Background(), driftArchivePostOptions)
}

// DriftArchivePostWithContext is an alternate form of the DriftArchivePost method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DriftArchivePostWithContext(ctx context.Context, driftArchivePostOptions *DriftArchivePostOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(driftArchivePostOptions, "driftArchivePostOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(driftArchivePostOptions, "driftArchivePostOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"data_mart_id":    *driftArchivePostOptions.DataMartID,
		"subscription_id": *driftArchivePostOptions.SubscriptionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitoring_services/drift/data_marts/{data_mart_id}/subscriptions/{subscription_id}/archives`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range driftArchivePostOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DriftArchivePost")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/octet-stream")

	if driftArchivePostOptions.ArchiveName != nil {
		builder.AddQuery("archive_name", fmt.Sprint(*driftArchivePostOptions.ArchiveName))
	}
	if driftArchivePostOptions.EnableDataDrift != nil {
		builder.AddQuery("enable_data_drift", fmt.Sprint(*driftArchivePostOptions.EnableDataDrift))
	}
	if driftArchivePostOptions.EnableModelDrift != nil {
		builder.AddQuery("enable_model_drift", fmt.Sprint(*driftArchivePostOptions.EnableModelDrift))
	}

	_, err = builder.SetBodyContent("application/octet-stream", nil, nil, driftArchivePostOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// DriftArchiveGet : Retrieves the Drift archives
// API to retrieve the Drift archives.
func (watsonOpenScale *WatsonOpenScaleV2) DriftArchiveGet(driftArchiveGetOptions *DriftArchiveGetOptions) (response *core.DetailedResponse, err error) {
	return watsonOpenScale.DriftArchiveGetWithContext(context.Background(), driftArchiveGetOptions)
}

// DriftArchiveGetWithContext is an alternate form of the DriftArchiveGet method which supports a Context parameter
func (watsonOpenScale *WatsonOpenScaleV2) DriftArchiveGetWithContext(ctx context.Context, driftArchiveGetOptions *DriftArchiveGetOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(driftArchiveGetOptions, "driftArchiveGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(driftArchiveGetOptions, "driftArchiveGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"monitor_instance_id": *driftArchiveGetOptions.MonitorInstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonOpenScale.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonOpenScale.Service.Options.URL, `/v2/monitoring_services/drift/monitor_instances/{monitor_instance_id}/archives`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range driftArchiveGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ai_openscale", "V2", "DriftArchiveGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/octet-stream")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonOpenScale.Service.Request(request, nil)

	return
}

// AnalyticsEngine : AnalyticsEngine struct
type AnalyticsEngine struct {
	// Credentials to override credentials in integration_reference.
	Credentials interface{} `json:"credentials,omitempty"`

	// id of the Integrated System.
	IntegratedSystemID *string `json:"integrated_system_id,omitempty"`

	// Additional parameters (e.g. max_num_executors, min_num_executors, executor_cores, executor_memory, driver_cores,
	// driver_memory).
	Parameters interface{} `json:"parameters,omitempty"`

	// Type of analytics engine. e.g. spark.
	Type *string `json:"type,omitempty"`
}

// UnmarshalAnalyticsEngine unmarshals an instance of AnalyticsEngine from the specified map of raw messages.
func UnmarshalAnalyticsEngine(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticsEngine)
	err = core.UnmarshalPrimitive(m, "credentials", &obj.Credentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "integrated_system_id", &obj.IntegratedSystemID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicabilitySelection : ApplicabilitySelection struct
type ApplicabilitySelection struct {
	InputDataType []string `json:"input_data_type,omitempty"`

	ProblemType []string `json:"problem_type,omitempty"`

	TargetType []string `json:"target_type,omitempty"`
}

// Constants associated with the ApplicabilitySelection.InputDataType property.
const (
	ApplicabilitySelection_InputDataType_Structured        = "structured"
	ApplicabilitySelection_InputDataType_UnstructuredAudio = "unstructured_audio"
	ApplicabilitySelection_InputDataType_UnstructuredImage = "unstructured_image"
	ApplicabilitySelection_InputDataType_UnstructuredText  = "unstructured_text"
	ApplicabilitySelection_InputDataType_UnstructuredVideo = "unstructured_video"
)

// Constants associated with the ApplicabilitySelection.ProblemType property.
const (
	ApplicabilitySelection_ProblemType_Binary     = "binary"
	ApplicabilitySelection_ProblemType_Multiclass = "multiclass"
	ApplicabilitySelection_ProblemType_Regression = "regression"
)

// Constants associated with the ApplicabilitySelection.TargetType property.
// Type of the target (e.g. subscription, business application, ...).
const (
	ApplicabilitySelection_TargetType_BusinessApplication = "business_application"
	ApplicabilitySelection_TargetType_DataMart            = "data_mart"
	ApplicabilitySelection_TargetType_Instance            = "instance"
	ApplicabilitySelection_TargetType_Subscription        = "subscription"
)

// UnmarshalApplicabilitySelection unmarshals an instance of ApplicabilitySelection from the specified map of raw messages.
func UnmarshalApplicabilitySelection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicabilitySelection)
	err = core.UnmarshalPrimitive(m, "input_data_type", &obj.InputDataType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "problem_type", &obj.ProblemType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_type", &obj.TargetType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Asset : Asset struct
type Asset struct {
	AssetID *string `json:"asset_id" validate:"required"`

	// Asset Resource Name (used for integration with 3rd party ML engines).
	AssetRn *string `json:"asset_rn,omitempty"`

	AssetType *string `json:"asset_type" validate:"required"`

	CreatedAt *string `json:"created_at,omitempty"`

	InputDataType *string `json:"input_data_type,omitempty"`

	ModelType *string `json:"model_type,omitempty"`

	Name *string `json:"name,omitempty"`

	ProblemType *string `json:"problem_type,omitempty"`

	RuntimeEnvironment *string `json:"runtime_environment,omitempty"`

	URL *string `json:"url" validate:"required"`
}

// Constants associated with the Asset.AssetType property.
const (
	Asset_AssetType_Function = "function"
	Asset_AssetType_Model    = "model"
)

// Constants associated with the Asset.InputDataType property.
const (
	Asset_InputDataType_Structured        = "structured"
	Asset_InputDataType_UnstructuredAudio = "unstructured_audio"
	Asset_InputDataType_UnstructuredImage = "unstructured_image"
	Asset_InputDataType_UnstructuredText  = "unstructured_text"
	Asset_InputDataType_UnstructuredVideo = "unstructured_video"
)

// Constants associated with the Asset.ProblemType property.
const (
	Asset_ProblemType_Binary     = "binary"
	Asset_ProblemType_Multiclass = "multiclass"
	Asset_ProblemType_Regression = "regression"
)

// NewAsset : Instantiate Asset (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewAsset(assetID string, assetType string, url string) (_model *Asset, err error) {
	_model = &Asset{
		AssetID:   core.StringPtr(assetID),
		AssetType: core.StringPtr(assetType),
		URL:       core.StringPtr(url),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAsset unmarshals an instance of Asset from the specified map of raw messages.
func UnmarshalAsset(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Asset)
	err = core.UnmarshalPrimitive(m, "asset_id", &obj.AssetID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_rn", &obj.AssetRn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_type", &obj.AssetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_data_type", &obj.InputDataType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_type", &obj.ModelType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "problem_type", &obj.ProblemType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "runtime_environment", &obj.RuntimeEnvironment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetDeployment : AssetDeployment struct
type AssetDeployment struct {
	CreatedAt *string `json:"created_at,omitempty"`

	DeploymentID *string `json:"deployment_id,omitempty"`

	// Deployment Resource Name (used for integration with 3rd party ML engines).
	DeploymentRn *string `json:"deployment_rn,omitempty"`

	// Deployment type, e.g. online, batch.
	DeploymentType *string `json:"deployment_type,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// Definition of scoring endpoint in custom_machine_learning.
	ScoringEndpoint *ScoringEndpoint `json:"scoring_endpoint,omitempty"`

	URL *string `json:"url,omitempty"`
}

// UnmarshalAssetDeployment unmarshals an instance of AssetDeployment from the specified map of raw messages.
func UnmarshalAssetDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetDeployment)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_rn", &obj.DeploymentRn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scoring_endpoint", &obj.ScoringEndpoint, UnmarshalScoringEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetDeploymentRequest : AssetDeploymentRequest struct
type AssetDeploymentRequest struct {
	CreatedAt *string `json:"created_at,omitempty"`

	DeploymentID *string `json:"deployment_id" validate:"required"`

	// Deployment Resource Name (used for integration with 3rd party ML engines).
	DeploymentRn *string `json:"deployment_rn,omitempty"`

	// Deployment type, e.g. online, batch.
	DeploymentType *string `json:"deployment_type" validate:"required"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name" validate:"required"`

	// Definition of scoring endpoint in custom_machine_learning.
	ScoringEndpoint *ScoringEndpointRequest `json:"scoring_endpoint,omitempty"`

	URL *string `json:"url,omitempty"`
}

// NewAssetDeploymentRequest : Instantiate AssetDeploymentRequest (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewAssetDeploymentRequest(deploymentID string, deploymentType string, name string) (_model *AssetDeploymentRequest, err error) {
	_model = &AssetDeploymentRequest{
		DeploymentID:   core.StringPtr(deploymentID),
		DeploymentType: core.StringPtr(deploymentType),
		Name:           core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAssetDeploymentRequest unmarshals an instance of AssetDeploymentRequest from the specified map of raw messages.
func UnmarshalAssetDeploymentRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetDeploymentRequest)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_rn", &obj.DeploymentRn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_type", &obj.DeploymentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scoring_endpoint", &obj.ScoringEndpoint, UnmarshalScoringEndpointRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetProperties : Additional asset properties (subject of discovery if not provided when creating the subscription).
type AssetProperties struct {
	// Fields to be given metadata `measure` of value `discrete`.
	CategoricalFields []string `json:"categorical_fields,omitempty"`

	DashboardConfiguration map[string]interface{} `json:"dashboard_configuration,omitempty"`

	// Fields to be given modeling_role feature.
	FeatureFields []string `json:"feature_fields,omitempty"`

	InputDataSchema *SparkStruct `json:"input_data_schema,omitempty"`

	InputDataType *string `json:"input_data_type,omitempty"`

	LabelColumn *string `json:"label_column,omitempty"`

	ModelType *string `json:"model_type,omitempty"`

	OutputDataSchema *SparkStruct `json:"output_data_schema,omitempty"`

	// Field with this name will be given modeling_role `decoded-target`.
	PredictedTargetField *string `json:"predicted_target_field,omitempty"`

	// Field with this name will be given modeling_role `prediction`.
	PredictionField *string `json:"prediction_field,omitempty"`

	// Fields to be given modeling_role `class_probability` (for columns of double data type) or `probability` (for column
	// of array data type).
	ProbabilityFields []string `json:"probability_fields,omitempty"`

	ProblemType *string `json:"problem_type,omitempty"`

	RuntimeEnvironment *string `json:"runtime_environment,omitempty"`

	TrainingDataReference *SecretCleaned `json:"training_data_reference,omitempty"`

	TrainingDataSchema *SparkStruct `json:"training_data_schema,omitempty"`

	// Field with this name will have `transaction_id_key` metadata set to true.
	TransactionIdField *string `json:"transaction_id_field,omitempty"`
}

// Constants associated with the AssetProperties.InputDataType property.
const (
	AssetProperties_InputDataType_Structured        = "structured"
	AssetProperties_InputDataType_UnstructuredAudio = "unstructured_audio"
	AssetProperties_InputDataType_UnstructuredImage = "unstructured_image"
	AssetProperties_InputDataType_UnstructuredText  = "unstructured_text"
	AssetProperties_InputDataType_UnstructuredVideo = "unstructured_video"
)

// Constants associated with the AssetProperties.ProblemType property.
const (
	AssetProperties_ProblemType_Binary     = "binary"
	AssetProperties_ProblemType_Multiclass = "multiclass"
	AssetProperties_ProblemType_Regression = "regression"
)

// UnmarshalAssetProperties unmarshals an instance of AssetProperties from the specified map of raw messages.
func UnmarshalAssetProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetProperties)
	err = core.UnmarshalPrimitive(m, "categorical_fields", &obj.CategoricalFields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dashboard_configuration", &obj.DashboardConfiguration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "feature_fields", &obj.FeatureFields)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data_schema", &obj.InputDataSchema, UnmarshalSparkStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_data_type", &obj.InputDataType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_column", &obj.LabelColumn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_type", &obj.ModelType)
	if err != nil {
		return
	}
	// err = core.UnmarshalModel(m, "output_data_schema", &obj.OutputDataSchema, UnmarshalSparkStruct)
	// if err != nil {
	// 	return
	// }
	err = core.UnmarshalPrimitive(m, "predicted_target_field", &obj.PredictedTargetField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prediction_field", &obj.PredictionField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "probability_fields", &obj.ProbabilityFields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "problem_type", &obj.ProblemType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "runtime_environment", &obj.RuntimeEnvironment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data_reference", &obj.TrainingDataReference, UnmarshalSecretCleaned)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data_schema", &obj.TrainingDataSchema, UnmarshalSparkStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transaction_id_field", &obj.TransactionIdField)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssetPropertiesRequest : Additional asset properties (subject of discovery if not provided when creating the subscription).
type AssetPropertiesRequest struct {
	// Fields to be given metadata `measure` of value `discrete`.
	CategoricalFields []string `json:"categorical_fields,omitempty"`

	DashboardConfiguration map[string]interface{} `json:"dashboard_configuration,omitempty"`

	// Fields to be given modeling_role feature.
	FeatureFields []string `json:"feature_fields,omitempty"`

	InputDataSchema *SparkStruct `json:"input_data_schema,omitempty"`

	LabelColumn *string `json:"label_column,omitempty"`

	Labels []string `json:"labels,omitempty"`

	OutputDataSchema *SparkStruct `json:"output_data_schema,omitempty"`

	// Field with this name will be given modeling_role `decoded-target`.
	PredictedTargetField *string `json:"predicted_target_field,omitempty"`

	// Field with this name will be given modeling_role `prediction`.
	PredictionField *string `json:"prediction_field,omitempty"`

	// Fields to be given modeling_role `class_probability` (for columns of double data type) or `probability` (for column
	// of array data type).
	ProbabilityFields []string `json:"probability_fields,omitempty"`

	TrainingDataReference *TrainingDataReference `json:"training_data_reference,omitempty"`

	TrainingDataSchema *SparkStruct `json:"training_data_schema,omitempty"`

	// Field with this name will have `transaction_id_key` metadata set to true.
	TransactionIdField *string `json:"transaction_id_field,omitempty"`
}

// UnmarshalAssetPropertiesRequest unmarshals an instance of AssetPropertiesRequest from the specified map of raw messages.
func UnmarshalAssetPropertiesRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssetPropertiesRequest)
	err = core.UnmarshalPrimitive(m, "categorical_fields", &obj.CategoricalFields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dashboard_configuration", &obj.DashboardConfiguration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "feature_fields", &obj.FeatureFields)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data_schema", &obj.InputDataSchema, UnmarshalSparkStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_column", &obj.LabelColumn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output_data_schema", &obj.OutputDataSchema, UnmarshalSparkStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "predicted_target_field", &obj.PredictedTargetField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prediction_field", &obj.PredictionField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "probability_fields", &obj.ProbabilityFields)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data_reference", &obj.TrainingDataReference, UnmarshalTrainingDataReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data_schema", &obj.TrainingDataSchema, UnmarshalSparkStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transaction_id_field", &obj.TransactionIdField)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AzureWorkspaceCredentials : AzureWorkspaceCredentials struct
type AzureWorkspaceCredentials struct {
	Token *string `json:"token" validate:"required"`

	WorkspaceID *string `json:"workspace_id" validate:"required"`
}

// NewAzureWorkspaceCredentials : Instantiate AzureWorkspaceCredentials (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewAzureWorkspaceCredentials(token string, workspaceID string) (_model *AzureWorkspaceCredentials, err error) {
	_model = &AzureWorkspaceCredentials{
		Token:       core.StringPtr(token),
		WorkspaceID: core.StringPtr(workspaceID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAzureWorkspaceCredentials unmarshals an instance of AzureWorkspaceCredentials from the specified map of raw messages.
func UnmarshalAzureWorkspaceCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AzureWorkspaceCredentials)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BusinessApplicationResponse : BusinessApplicationResponse struct
type BusinessApplicationResponse struct {
	// business application payload.
	Entity *BusinessApplicationResponseEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalBusinessApplicationResponse unmarshals an instance of BusinessApplicationResponse from the specified map of raw messages.
func UnmarshalBusinessApplicationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BusinessApplicationResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalBusinessApplicationResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BusinessApplicationResponseEntity : business application payload.
type BusinessApplicationResponseEntity struct {
	BusinessMetrics []BusinessMetric `json:"business_metrics" validate:"required"`

	BusinessMetricsMonitorDefinitionID *string `json:"business_metrics_monitor_definition_id,omitempty"`

	BusinessMetricsMonitorInstanceID *string `json:"business_metrics_monitor_instance_id,omitempty"`

	// Unique identifier of the data set (like scoring, feedback or business payload).
	BusinessPayloadDataSetID *string `json:"business_payload_data_set_id,omitempty"`

	CorrelationMonitorInstanceID *string `json:"correlation_monitor_instance_id,omitempty"`

	// description fo the business application.
	Description *string `json:"description" validate:"required"`

	// name fo the business application.
	Name *string `json:"name" validate:"required"`

	PayloadFields []PayloadField `json:"payload_fields" validate:"required"`

	SubscriptionIds []string `json:"subscription_ids,omitempty"`

	// Unique identifier of the data set (like scoring, feedback or business payload).
	TransactionBatchesDataSetID *string `json:"transaction_batches_data_set_id,omitempty"`

	Status *Status `json:"status,omitempty"`
}

// UnmarshalBusinessApplicationResponseEntity unmarshals an instance of BusinessApplicationResponseEntity from the specified map of raw messages.
func UnmarshalBusinessApplicationResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BusinessApplicationResponseEntity)
	err = core.UnmarshalModel(m, "business_metrics", &obj.BusinessMetrics, UnmarshalBusinessMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "business_metrics_monitor_definition_id", &obj.BusinessMetricsMonitorDefinitionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "business_metrics_monitor_instance_id", &obj.BusinessMetricsMonitorInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "business_payload_data_set_id", &obj.BusinessPayloadDataSetID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "correlation_monitor_instance_id", &obj.CorrelationMonitorInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "payload_fields", &obj.PayloadFields, UnmarshalPayloadField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subscription_ids", &obj.SubscriptionIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transaction_batches_data_set_id", &obj.TransactionBatchesDataSetID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BusinessApplicationsAddOptions : The BusinessApplicationsAdd options.
type BusinessApplicationsAddOptions struct {
	BusinessMetrics []BusinessMetric `json:"business_metrics" validate:"required"`

	// description fo the business application.
	Description *string `json:"description" validate:"required"`

	// name fo the business application.
	Name *string `json:"name" validate:"required"`

	PayloadFields []PayloadField `json:"payload_fields" validate:"required"`

	BusinessMetricsMonitorDefinitionID *string `json:"business_metrics_monitor_definition_id,omitempty"`

	BusinessMetricsMonitorInstanceID *string `json:"business_metrics_monitor_instance_id,omitempty"`

	// Unique identifier of the data set (like scoring, feedback or business payload).
	BusinessPayloadDataSetID *string `json:"business_payload_data_set_id,omitempty"`

	CorrelationMonitorInstanceID *string `json:"correlation_monitor_instance_id,omitempty"`

	SubscriptionIds []string `json:"subscription_ids,omitempty"`

	// Unique identifier of the data set (like scoring, feedback or business payload).
	TransactionBatchesDataSetID *string `json:"transaction_batches_data_set_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewBusinessApplicationsAddOptions : Instantiate BusinessApplicationsAddOptions
func (*WatsonOpenScaleV2) NewBusinessApplicationsAddOptions(businessMetrics []BusinessMetric, description string, name string, payloadFields []PayloadField) *BusinessApplicationsAddOptions {
	return &BusinessApplicationsAddOptions{
		BusinessMetrics: businessMetrics,
		Description:     core.StringPtr(description),
		Name:            core.StringPtr(name),
		PayloadFields:   payloadFields,
	}
}

// SetBusinessMetrics : Allow user to set BusinessMetrics
func (_options *BusinessApplicationsAddOptions) SetBusinessMetrics(businessMetrics []BusinessMetric) *BusinessApplicationsAddOptions {
	_options.BusinessMetrics = businessMetrics
	return _options
}

// SetDescription : Allow user to set Description
func (_options *BusinessApplicationsAddOptions) SetDescription(description string) *BusinessApplicationsAddOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetName : Allow user to set Name
func (_options *BusinessApplicationsAddOptions) SetName(name string) *BusinessApplicationsAddOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetPayloadFields : Allow user to set PayloadFields
func (_options *BusinessApplicationsAddOptions) SetPayloadFields(payloadFields []PayloadField) *BusinessApplicationsAddOptions {
	_options.PayloadFields = payloadFields
	return _options
}

// SetBusinessMetricsMonitorDefinitionID : Allow user to set BusinessMetricsMonitorDefinitionID
func (_options *BusinessApplicationsAddOptions) SetBusinessMetricsMonitorDefinitionID(businessMetricsMonitorDefinitionID string) *BusinessApplicationsAddOptions {
	_options.BusinessMetricsMonitorDefinitionID = core.StringPtr(businessMetricsMonitorDefinitionID)
	return _options
}

// SetBusinessMetricsMonitorInstanceID : Allow user to set BusinessMetricsMonitorInstanceID
func (_options *BusinessApplicationsAddOptions) SetBusinessMetricsMonitorInstanceID(businessMetricsMonitorInstanceID string) *BusinessApplicationsAddOptions {
	_options.BusinessMetricsMonitorInstanceID = core.StringPtr(businessMetricsMonitorInstanceID)
	return _options
}

// SetBusinessPayloadDataSetID : Allow user to set BusinessPayloadDataSetID
func (_options *BusinessApplicationsAddOptions) SetBusinessPayloadDataSetID(businessPayloadDataSetID string) *BusinessApplicationsAddOptions {
	_options.BusinessPayloadDataSetID = core.StringPtr(businessPayloadDataSetID)
	return _options
}

// SetCorrelationMonitorInstanceID : Allow user to set CorrelationMonitorInstanceID
func (_options *BusinessApplicationsAddOptions) SetCorrelationMonitorInstanceID(correlationMonitorInstanceID string) *BusinessApplicationsAddOptions {
	_options.CorrelationMonitorInstanceID = core.StringPtr(correlationMonitorInstanceID)
	return _options
}

// SetSubscriptionIds : Allow user to set SubscriptionIds
func (_options *BusinessApplicationsAddOptions) SetSubscriptionIds(subscriptionIds []string) *BusinessApplicationsAddOptions {
	_options.SubscriptionIds = subscriptionIds
	return _options
}

// SetTransactionBatchesDataSetID : Allow user to set TransactionBatchesDataSetID
func (_options *BusinessApplicationsAddOptions) SetTransactionBatchesDataSetID(transactionBatchesDataSetID string) *BusinessApplicationsAddOptions {
	_options.TransactionBatchesDataSetID = core.StringPtr(transactionBatchesDataSetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *BusinessApplicationsAddOptions) SetHeaders(param map[string]string) *BusinessApplicationsAddOptions {
	options.Headers = param
	return options
}

// BusinessApplicationsCollection : BusinessApplicationsCollection struct
type BusinessApplicationsCollection struct {
	BusinessApplications []BusinessApplicationResponse `json:"business_applications" validate:"required"`
}

// UnmarshalBusinessApplicationsCollection unmarshals an instance of BusinessApplicationsCollection from the specified map of raw messages.
func UnmarshalBusinessApplicationsCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BusinessApplicationsCollection)
	err = core.UnmarshalModel(m, "business_applications", &obj.BusinessApplications, UnmarshalBusinessApplicationResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BusinessApplicationsDeleteOptions : The BusinessApplicationsDelete options.
type BusinessApplicationsDeleteOptions struct {
	// ID of the business application.
	BusinessApplicationID *string `json:"business_application_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewBusinessApplicationsDeleteOptions : Instantiate BusinessApplicationsDeleteOptions
func (*WatsonOpenScaleV2) NewBusinessApplicationsDeleteOptions(businessApplicationID string) *BusinessApplicationsDeleteOptions {
	return &BusinessApplicationsDeleteOptions{
		BusinessApplicationID: core.StringPtr(businessApplicationID),
	}
}

// SetBusinessApplicationID : Allow user to set BusinessApplicationID
func (_options *BusinessApplicationsDeleteOptions) SetBusinessApplicationID(businessApplicationID string) *BusinessApplicationsDeleteOptions {
	_options.BusinessApplicationID = core.StringPtr(businessApplicationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *BusinessApplicationsDeleteOptions) SetHeaders(param map[string]string) *BusinessApplicationsDeleteOptions {
	options.Headers = param
	return options
}

// BusinessApplicationsGetOptions : The BusinessApplicationsGet options.
type BusinessApplicationsGetOptions struct {
	// ID of the business application.
	BusinessApplicationID *string `json:"business_application_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewBusinessApplicationsGetOptions : Instantiate BusinessApplicationsGetOptions
func (*WatsonOpenScaleV2) NewBusinessApplicationsGetOptions(businessApplicationID string) *BusinessApplicationsGetOptions {
	return &BusinessApplicationsGetOptions{
		BusinessApplicationID: core.StringPtr(businessApplicationID),
	}
}

// SetBusinessApplicationID : Allow user to set BusinessApplicationID
func (_options *BusinessApplicationsGetOptions) SetBusinessApplicationID(businessApplicationID string) *BusinessApplicationsGetOptions {
	_options.BusinessApplicationID = core.StringPtr(businessApplicationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *BusinessApplicationsGetOptions) SetHeaders(param map[string]string) *BusinessApplicationsGetOptions {
	options.Headers = param
	return options
}

// BusinessApplicationsListOptions : The BusinessApplicationsList options.
type BusinessApplicationsListOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewBusinessApplicationsListOptions : Instantiate BusinessApplicationsListOptions
func (*WatsonOpenScaleV2) NewBusinessApplicationsListOptions() *BusinessApplicationsListOptions {
	return &BusinessApplicationsListOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *BusinessApplicationsListOptions) SetHeaders(param map[string]string) *BusinessApplicationsListOptions {
	options.Headers = param
	return options
}

// BusinessApplicationsUpdateOptions : The BusinessApplicationsUpdate options.
type BusinessApplicationsUpdateOptions struct {
	// ID of the business application.
	BusinessApplicationID *string `json:"business_application_id" validate:"required,ne="`

	PatchDocument []PatchDocument `json:"PatchDocument" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewBusinessApplicationsUpdateOptions : Instantiate BusinessApplicationsUpdateOptions
func (*WatsonOpenScaleV2) NewBusinessApplicationsUpdateOptions(businessApplicationID string, patchDocument []PatchDocument) *BusinessApplicationsUpdateOptions {
	return &BusinessApplicationsUpdateOptions{
		BusinessApplicationID: core.StringPtr(businessApplicationID),
		PatchDocument:         patchDocument,
	}
}

// SetBusinessApplicationID : Allow user to set BusinessApplicationID
func (_options *BusinessApplicationsUpdateOptions) SetBusinessApplicationID(businessApplicationID string) *BusinessApplicationsUpdateOptions {
	_options.BusinessApplicationID = core.StringPtr(businessApplicationID)
	return _options
}

// SetPatchDocument : Allow user to set PatchDocument
func (_options *BusinessApplicationsUpdateOptions) SetPatchDocument(patchDocument []PatchDocument) *BusinessApplicationsUpdateOptions {
	_options.PatchDocument = patchDocument
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *BusinessApplicationsUpdateOptions) SetHeaders(param map[string]string) *BusinessApplicationsUpdateOptions {
	options.Headers = param
	return options
}

// BusinessMetric : BusinessMetric struct
type BusinessMetric struct {
	CalculationMetadata *CalculationMeta `json:"calculation_metadata" validate:"required"`

	// Description of the metrics.
	Description *string `json:"description,omitempty"`

	// the indicator specifying the expected direction of the monotonic metric values.
	ExpectedDirection *string `json:"expected_direction,omitempty"`

	ID *string `json:"id,omitempty"`

	// unique name used by UI instead of id (must be unique in scope of the monitor defition across both metrics and tags).
	Name *string `json:"name" validate:"required"`

	Required *bool `json:"required,omitempty"`

	Thresholds []MetricThreshold `json:"thresholds,omitempty"`
}

// Constants associated with the BusinessMetric.ExpectedDirection property.
// the indicator specifying the expected direction of the monotonic metric values.
const (
	BusinessMetric_ExpectedDirection_Decreasing = "decreasing"
	BusinessMetric_ExpectedDirection_Increasing = "increasing"
	BusinessMetric_ExpectedDirection_Unknown    = "unknown"
)

// NewBusinessMetric : Instantiate BusinessMetric (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewBusinessMetric(calculationMetadata *CalculationMeta, name string) (_model *BusinessMetric, err error) {
	_model = &BusinessMetric{
		CalculationMetadata: calculationMetadata,
		Name:                core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalBusinessMetric unmarshals an instance of BusinessMetric from the specified map of raw messages.
func UnmarshalBusinessMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BusinessMetric)
	err = core.UnmarshalModel(m, "calculation_metadata", &obj.CalculationMetadata, UnmarshalCalculationMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_direction", &obj.ExpectedDirection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "thresholds", &obj.Thresholds, UnmarshalMetricThreshold)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CalculationMeta : CalculationMeta struct
type CalculationMeta struct {
	Aggregation *string `json:"aggregation,omitempty"`

	FieldName *string `json:"field_name" validate:"required"`

	TimeFrame *TimeFrame `json:"time_frame,omitempty"`
}

// Constants associated with the CalculationMeta.Aggregation property.
const (
	CalculationMeta_Aggregation_Avg = "avg"
	CalculationMeta_Aggregation_Max = "max"
	CalculationMeta_Aggregation_Min = "min"
	CalculationMeta_Aggregation_Sum = "sum"
)

// NewCalculationMeta : Instantiate CalculationMeta (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewCalculationMeta(fieldName string) (_model *CalculationMeta, err error) {
	_model = &CalculationMeta{
		FieldName: core.StringPtr(fieldName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalCalculationMeta unmarshals an instance of CalculationMeta from the specified map of raw messages.
func UnmarshalCalculationMeta(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CalculationMeta)
	err = core.UnmarshalPrimitive(m, "aggregation", &obj.Aggregation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field_name", &obj.FieldName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "time_frame", &obj.TimeFrame, UnmarshalTimeFrame)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CollectionUrlModel : CollectionUrlModel struct
type CollectionUrlModel struct {
	// URI of a resource.
	URL *string `json:"url" validate:"required"`
}

// UnmarshalCollectionUrlModel unmarshals an instance of CollectionUrlModel from the specified map of raw messages.
func UnmarshalCollectionUrlModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CollectionUrlModel)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataDistributionResponse : Data distribution details response.
type DataDistributionResponse struct {
	// The status information for the monitoring run.
	Entity *DataDistributionResponseEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalDataDistributionResponse unmarshals an instance of DataDistributionResponse from the specified map of raw messages.
func UnmarshalDataDistributionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataDistributionResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalDataDistributionResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataDistributionResponseEntity : The status information for the monitoring run.
type DataDistributionResponseEntity struct {
	// Definition of aggregations, by default 'count'.
	//
	// Aggregations can be one of:
	// * count
	// * <column_name>:sum
	// * <column_name>:min
	// * <column_name>:max
	// * <column_name>:avg
	// * <column_name>:stddev.
	Agg []string `json:"agg,omitempty"`

	// type of a data set.
	Dataset *string `json:"dataset" validate:"required"`

	// end datetime in ISO format.
	End *string `json:"end" validate:"required"`

	// Filters defined by user in format: {field_name}:{op}:{value}. Partly compatible with filters in "filter" parameter
	// of GET /v2/data_sets/{data_set_id}/records.
	//
	// Possible filter operators:
	// * eq - equals (numeric, string)
	// * gt - greater than (numeric)
	// * gte - greater than or equal (numeric)
	// * lt - lower than (numeric)
	// * lte - lower than or equal (numeric)
	// * in - value in a set (numeric, string)
	// * field:null (a no-argument filter) - value is null (any nullable)
	// * field:exists (a no-argument filter) - value is not null (any column).
	Filter *string `json:"filter,omitempty"`

	// names of columns to be grouped.
	Group []string `json:"group" validate:"required"`

	// limit for number of rows, by default it is 50,000 (max possible limit is 50,000).
	Limit *float64 `json:"limit,omitempty"`

	// max number of bins which will be generated for data.
	MaxBins *float64 `json:"max_bins,omitempty"`

	// start datetime in ISO format.
	Start *string `json:"start" validate:"required"`

	Distribution *DataDistributionResponseEntityDistribution `json:"distribution,omitempty"`

	// was the limit used on data.
	LimitedData *bool `json:"limited_data,omitempty"`

	// number of processed records.
	ProcessedRecords *float64 `json:"processed_records,omitempty"`

	// The status information for the monitoring run.
	Status *DataDistributionStatus `json:"status,omitempty"`
}

// Constants associated with the DataDistributionResponseEntity.Dataset property.
// type of a data set.
const (
	DataDistributionResponseEntity_Dataset_BusinessPayload    = "business_payload"
	DataDistributionResponseEntity_Dataset_Custom             = "custom"
	DataDistributionResponseEntity_Dataset_Explanations       = "explanations"
	DataDistributionResponseEntity_Dataset_ExplanationsWhatif = "explanations_whatif"
	DataDistributionResponseEntity_Dataset_Feedback           = "feedback"
	DataDistributionResponseEntity_Dataset_ManualLabeling     = "manual_labeling"
	DataDistributionResponseEntity_Dataset_PayloadLogging     = "payload_logging"
	DataDistributionResponseEntity_Dataset_Training           = "training"
)

// UnmarshalDataDistributionResponseEntity unmarshals an instance of DataDistributionResponseEntity from the specified map of raw messages.
func UnmarshalDataDistributionResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataDistributionResponseEntity)
	err = core.UnmarshalPrimitive(m, "agg", &obj.Agg)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dataset", &obj.Dataset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "filter", &obj.Filter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "group", &obj.Group)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_bins", &obj.MaxBins)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "distribution", &obj.Distribution, UnmarshalDataDistributionResponseEntityDistribution)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limited_data", &obj.LimitedData)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "processed_records", &obj.ProcessedRecords)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalDataDistributionStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataDistributionResponseEntityDistribution : DataDistributionResponseEntityDistribution struct
type DataDistributionResponseEntityDistribution struct {
	// names of the data distribution fields.
	Fields []string `json:"fields" validate:"required"`

	// data distribution rows.
	Values []interface{} `json:"values" validate:"required"`
}

// UnmarshalDataDistributionResponseEntityDistribution unmarshals an instance of DataDistributionResponseEntityDistribution from the specified map of raw messages.
func UnmarshalDataDistributionResponseEntityDistribution(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataDistributionResponseEntityDistribution)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataDistributionStatus : The status information for the monitoring run.
type DataDistributionStatus struct {
	// The timestamp when the monitoring run finished running (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	Failure *GenericErrorResponse `json:"failure,omitempty"`

	// The timestamp when the monitoring run was started running (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	State *string `json:"state,omitempty"`

	// The timestamp when the monitoring run was last updated (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`
}

// Constants associated with the DataDistributionStatus.State property.
const (
	DataDistributionStatus_State_Completed    = "completed"
	DataDistributionStatus_State_Failed       = "failed"
	DataDistributionStatus_State_Initializing = "initializing"
	DataDistributionStatus_State_Running      = "running"
)

// UnmarshalDataDistributionStatus unmarshals an instance of DataDistributionStatus from the specified map of raw messages.
func UnmarshalDataDistributionStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataDistributionStatus)
	err = core.UnmarshalPrimitive(m, "completed_at", &obj.CompletedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalGenericErrorResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_at", &obj.StartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartDatabaseResponse : DataMartDatabaseResponse struct
type DataMartDatabaseResponse struct {
	Entity *DataMartDatabaseResponseEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalDataMartDatabaseResponse unmarshals an instance of DataMartDatabaseResponse from the specified map of raw messages.
func UnmarshalDataMartDatabaseResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartDatabaseResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalDataMartDatabaseResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonOpenScaleV2) NewDataMartDatabaseResponsePatch(dataMartDatabaseResponse *DataMartDatabaseResponse) (_patch []JSONPatchOperation) {
	if dataMartDatabaseResponse.Entity != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/entity"),
			Value: dataMartDatabaseResponse.Entity,
		})
	}
	if dataMartDatabaseResponse.Metadata != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/metadata"),
			Value: dataMartDatabaseResponse.Metadata,
		})
	}
	return
}

// DataMartDatabaseResponseCollection : DataMartDatabaseResponseCollection struct
type DataMartDatabaseResponseCollection struct {
	DataMarts []DataMartDatabaseResponse `json:"data_marts" validate:"required"`
}

// UnmarshalDataMartDatabaseResponseCollection unmarshals an instance of DataMartDatabaseResponseCollection from the specified map of raw messages.
func UnmarshalDataMartDatabaseResponseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartDatabaseResponseCollection)
	err = core.UnmarshalModel(m, "data_marts", &obj.DataMarts, UnmarshalDataMartDatabaseResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartDatabaseResponseEntity : DataMartDatabaseResponseEntity struct
type DataMartDatabaseResponseEntity struct {
	// Database configuration ignored if internal database is requested (`internal_database` is `true`).
	DatabaseConfiguration *DatabaseConfiguration `json:"database_configuration,omitempty"`

	// Used by UI to check if database discovery was automatic or manual.
	DatabaseDiscovery *string `json:"database_discovery,omitempty"`

	Description *string `json:"description,omitempty"`

	// If `true` the internal database managed by AI OpenScale is provided for the user.
	InternalDatabase *bool `json:"internal_database,omitempty"`

	Name *string `json:"name,omitempty"`

	ServiceInstanceCrn *string `json:"service_instance_crn,omitempty"`

	Status *Status `json:"status,omitempty"`
}

// Constants associated with the DataMartDatabaseResponseEntity.DatabaseDiscovery property.
// Used by UI to check if database discovery was automatic or manual.
const (
	DataMartDatabaseResponseEntity_DatabaseDiscovery_Automatic = "automatic"
	DataMartDatabaseResponseEntity_DatabaseDiscovery_Manual    = "manual"
)

// UnmarshalDataMartDatabaseResponseEntity unmarshals an instance of DataMartDatabaseResponseEntity from the specified map of raw messages.
func UnmarshalDataMartDatabaseResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartDatabaseResponseEntity)
	err = core.UnmarshalModel(m, "database_configuration", &obj.DatabaseConfiguration, UnmarshalDatabaseConfiguration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_discovery", &obj.DatabaseDiscovery)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "internal_database", &obj.InternalDatabase)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_instance_crn", &obj.ServiceInstanceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartGetMonitorInstanceMetrics : DataMartGetMonitorInstanceMetrics struct
type DataMartGetMonitorInstanceMetrics struct {
	// Ceiled to full interval.
	End *strfmt.DateTime `json:"end,omitempty"`

	Groups []DataMartGetMonitorInstanceMetricsGroupsItem `json:"groups,omitempty"`

	Interval *string `json:"interval,omitempty"`

	MonitorDefinitionID *string `json:"monitor_definition_id,omitempty"`

	// Floored to full interval.
	Start *strfmt.DateTime `json:"start,omitempty"`
}

// UnmarshalDataMartGetMonitorInstanceMetrics unmarshals an instance of DataMartGetMonitorInstanceMetrics from the specified map of raw messages.
func UnmarshalDataMartGetMonitorInstanceMetrics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartGetMonitorInstanceMetrics)
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalDataMartGetMonitorInstanceMetricsGroupsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor_definition_id", &obj.MonitorDefinitionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartGetMonitorInstanceMetricsGroupsItem : DataMartGetMonitorInstanceMetricsGroupsItem struct
type DataMartGetMonitorInstanceMetricsGroupsItem struct {
	Metrics []DataMartGetMonitorInstanceMetricsGroupsItemMetricsItem `json:"metrics,omitempty"`

	Tags []DataMartGetMonitorInstanceMetricsGroupsItemTagsItem `json:"tags,omitempty"`
}

// UnmarshalDataMartGetMonitorInstanceMetricsGroupsItem unmarshals an instance of DataMartGetMonitorInstanceMetricsGroupsItem from the specified map of raw messages.
func UnmarshalDataMartGetMonitorInstanceMetricsGroupsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartGetMonitorInstanceMetricsGroupsItem)
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemMetricsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemTagsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartGetMonitorInstanceMetricsGroupsItemMetricsItem : DataMartGetMonitorInstanceMetricsGroupsItemMetricsItem struct
type DataMartGetMonitorInstanceMetricsGroupsItemMetricsItem struct {
	ID *string `json:"id" validate:"required"`

	LowerLimit *float64 `json:"lower_limit,omitempty"`

	Min *DataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin `json:"min,omitempty"`

	UpperLimit *float64 `json:"upper_limit,omitempty"`
}

// UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemMetricsItem unmarshals an instance of DataMartGetMonitorInstanceMetricsGroupsItemMetricsItem from the specified map of raw messages.
func UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemMetricsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartGetMonitorInstanceMetricsGroupsItemMetricsItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "lower_limit", &obj.LowerLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "min", &obj.Min, UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "upper_limit", &obj.UpperLimit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin : DataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin struct
type DataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin struct {
	MeasurementID []string `json:"measurement_id,omitempty"`

	Value []float64 `json:"value,omitempty"`
}

// UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin unmarshals an instance of DataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin from the specified map of raw messages.
func UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartGetMonitorInstanceMetricsGroupsItemMetricsItemMin)
	err = core.UnmarshalPrimitive(m, "measurement_id", &obj.MeasurementID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartGetMonitorInstanceMetricsGroupsItemTagsItem : DataMartGetMonitorInstanceMetricsGroupsItemTagsItem struct
type DataMartGetMonitorInstanceMetricsGroupsItemTagsItem struct {
	ID *string `json:"id" validate:"required"`

	Value *string `json:"value" validate:"required"`
}

// UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemTagsItem unmarshals an instance of DataMartGetMonitorInstanceMetricsGroupsItemTagsItem from the specified map of raw messages.
func UnmarshalDataMartGetMonitorInstanceMetricsGroupsItemTagsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataMartGetMonitorInstanceMetricsGroupsItemTagsItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataMartsAddOptions : The DataMartsAdd options.
type DataMartsAddOptions struct {
	// Database configuration ignored if internal database is requested (`internal_database` is `true`).
	DatabaseConfiguration *DatabaseConfigurationRequest `json:"database_configuration,omitempty"`

	// Indicates if the database was discovered automatically or manually added by user through UI.
	DatabaseDiscovery *string `json:"database_discovery,omitempty"`

	// Description of the data mart.
	Description *string `json:"description,omitempty"`

	// If `true` the internal database managed by AI OpenScale is provided for the user.
	InternalDatabase *bool `json:"internal_database,omitempty"`

	// Name of the data mart.
	Name *string `json:"name,omitempty"`

	// Can be omitted if user token is used for authorization.
	ServiceInstanceCrn *string `json:"service_instance_crn,omitempty"`

	// force update of metadata and db credentials (assumption is that the new database is already prepared and populated).
	Force *bool `json:"force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DataMartsAddOptions.DatabaseDiscovery property.
// Indicates if the database was discovered automatically or manually added by user through UI.
const (
	DataMartsAddOptions_DatabaseDiscovery_Automatic = "automatic"
	DataMartsAddOptions_DatabaseDiscovery_Manual    = "manual"
)

// NewDataMartsAddOptions : Instantiate DataMartsAddOptions
func (*WatsonOpenScaleV2) NewDataMartsAddOptions() *DataMartsAddOptions {
	return &DataMartsAddOptions{}
}

// SetDatabaseConfiguration : Allow user to set DatabaseConfiguration
func (_options *DataMartsAddOptions) SetDatabaseConfiguration(databaseConfiguration *DatabaseConfigurationRequest) *DataMartsAddOptions {
	_options.DatabaseConfiguration = databaseConfiguration
	return _options
}

// SetDatabaseDiscovery : Allow user to set DatabaseDiscovery
func (_options *DataMartsAddOptions) SetDatabaseDiscovery(databaseDiscovery string) *DataMartsAddOptions {
	_options.DatabaseDiscovery = core.StringPtr(databaseDiscovery)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *DataMartsAddOptions) SetDescription(description string) *DataMartsAddOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetInternalDatabase : Allow user to set InternalDatabase
func (_options *DataMartsAddOptions) SetInternalDatabase(internalDatabase bool) *DataMartsAddOptions {
	_options.InternalDatabase = core.BoolPtr(internalDatabase)
	return _options
}

// SetName : Allow user to set Name
func (_options *DataMartsAddOptions) SetName(name string) *DataMartsAddOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetServiceInstanceCrn : Allow user to set ServiceInstanceCrn
func (_options *DataMartsAddOptions) SetServiceInstanceCrn(serviceInstanceCrn string) *DataMartsAddOptions {
	_options.ServiceInstanceCrn = core.StringPtr(serviceInstanceCrn)
	return _options
}

// SetForce : Allow user to set Force
func (_options *DataMartsAddOptions) SetForce(force bool) *DataMartsAddOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataMartsAddOptions) SetHeaders(param map[string]string) *DataMartsAddOptions {
	options.Headers = param
	return options
}

// DataMartsDeleteOptions : The DataMartsDelete options.
type DataMartsDeleteOptions struct {
	// ID of the data mart.
	DataMartID *string `json:"data_mart_id" validate:"required,ne="`

	// force hard delete.
	Force *bool `json:"force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDataMartsDeleteOptions : Instantiate DataMartsDeleteOptions
func (*WatsonOpenScaleV2) NewDataMartsDeleteOptions(dataMartID string) *DataMartsDeleteOptions {
	return &DataMartsDeleteOptions{
		DataMartID: core.StringPtr(dataMartID),
	}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *DataMartsDeleteOptions) SetDataMartID(dataMartID string) *DataMartsDeleteOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetForce : Allow user to set Force
func (_options *DataMartsDeleteOptions) SetForce(force bool) *DataMartsDeleteOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataMartsDeleteOptions) SetHeaders(param map[string]string) *DataMartsDeleteOptions {
	options.Headers = param
	return options
}

// DataMartsGetOptions : The DataMartsGet options.
type DataMartsGetOptions struct {
	// ID of the data mart.
	DataMartID *string `json:"data_mart_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDataMartsGetOptions : Instantiate DataMartsGetOptions
func (*WatsonOpenScaleV2) NewDataMartsGetOptions(dataMartID string) *DataMartsGetOptions {
	return &DataMartsGetOptions{
		DataMartID: core.StringPtr(dataMartID),
	}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *DataMartsGetOptions) SetDataMartID(dataMartID string) *DataMartsGetOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataMartsGetOptions) SetHeaders(param map[string]string) *DataMartsGetOptions {
	options.Headers = param
	return options
}

// DataMartsListOptions : The DataMartsList options.
type DataMartsListOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDataMartsListOptions : Instantiate DataMartsListOptions
func (*WatsonOpenScaleV2) NewDataMartsListOptions() *DataMartsListOptions {
	return &DataMartsListOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *DataMartsListOptions) SetHeaders(param map[string]string) *DataMartsListOptions {
	options.Headers = param
	return options
}

// DataMartsPatchOptions : The DataMartsPatch options.
type DataMartsPatchOptions struct {
	// ID of the data mart.
	DataMartID *string `json:"data_mart_id" validate:"required,ne="`

	// Array of patch operations as defined in RFC 6902.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDataMartsPatchOptions : Instantiate DataMartsPatchOptions
func (*WatsonOpenScaleV2) NewDataMartsPatchOptions(dataMartID string, jsonPatchOperation []JSONPatchOperation) *DataMartsPatchOptions {
	return &DataMartsPatchOptions{
		DataMartID:         core.StringPtr(dataMartID),
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *DataMartsPatchOptions) SetDataMartID(dataMartID string) *DataMartsPatchOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *DataMartsPatchOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *DataMartsPatchOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataMartsPatchOptions) SetHeaders(param map[string]string) *DataMartsPatchOptions {
	options.Headers = param
	return options
}

// DataRecord : DataRecord struct
type DataRecord struct {
	// Any JSON object representing annotations.
	Annotations map[string]interface{} `json:"annotations,omitempty"`

	// Fields and values of the entity matches JSON object's fields and values.
	Values *JsDictElem `json:"values" validate:"required"`
}

// UnmarshalDataRecord unmarshals an instance of DataRecord from the specified map of raw messages.
func UnmarshalDataRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataRecord)
	err = core.UnmarshalPrimitive(m, "annotations", &obj.Annotations)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalJsDictElem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataRecordResponse : DataRecordResponse struct
type DataRecordResponse struct {
	Entity *DataRecord `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalDataRecordResponse unmarshals an instance of DataRecordResponse from the specified map of raw messages.
func UnmarshalDataRecordResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataRecordResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalDataRecord)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataRecordResponseList : DataRecordResponseList struct
type DataRecordResponseList struct {
	// List of annotations objects.
	Annotations []map[string]interface{} `json:"annotations,omitempty"`

	// Fields' names are listed in order in 'fields'.
	Fields []string `json:"fields" validate:"required"`

	// Rows organized per fields as described in 'fields'.
	Values [][]interface{} `json:"values" validate:"required"`
}

// UnmarshalDataRecordResponseList unmarshals an instance of DataRecordResponseList from the specified map of raw messages.
func UnmarshalDataRecordResponseList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataRecordResponseList)
	err = core.UnmarshalPrimitive(m, "annotations", &obj.Annotations)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetObject : DataSetObject struct
type DataSetObject struct {
	DataMartID *string `json:"data_mart_id" validate:"required"`

	DataSchema *SparkStruct `json:"data_schema" validate:"required"`

	Description *string `json:"description,omitempty"`

	Location *LocationTableName `json:"location,omitempty"`

	ManagedBy *string `json:"managed_by,omitempty"`

	Name *string `json:"name" validate:"required"`

	SchemaUpdateMode *string `json:"schema_update_mode,omitempty"`

	Status *Status `json:"status" validate:"required"`

	Target *Target `json:"target" validate:"required"`

	// type of a data set.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the DataSetObject.SchemaUpdateMode property.
const (
	DataSetObject_SchemaUpdateMode_Auto = "auto"
	DataSetObject_SchemaUpdateMode_None = "none"
)

// Constants associated with the DataSetObject.Type property.
// type of a data set.
const (
	DataSetObject_Type_BusinessPayload    = "business_payload"
	DataSetObject_Type_Custom             = "custom"
	DataSetObject_Type_Explanations       = "explanations"
	DataSetObject_Type_ExplanationsWhatif = "explanations_whatif"
	DataSetObject_Type_Feedback           = "feedback"
	DataSetObject_Type_ManualLabeling     = "manual_labeling"
	DataSetObject_Type_PayloadLogging     = "payload_logging"
	DataSetObject_Type_Training           = "training"
)

// UnmarshalDataSetObject unmarshals an instance of DataSetObject from the specified map of raw messages.
func UnmarshalDataSetObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetObject)
	err = core.UnmarshalPrimitive(m, "data_mart_id", &obj.DataMartID)
	if err != nil {
		return
	}
	// err = core.UnmarshalModel(m, "data_schema", &obj.DataSchema, UnmarshalSparkStruct)
	// if err != nil {
	// 	return
	// }
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocationTableName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schema_update_mode", &obj.SchemaUpdateMode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetRecords : DataSetRecords struct
type DataSetRecords struct {
	DataSetRecords []DataSetRecordsDataSetRecordsItem `json:"data_set_records" validate:"required"`
}

// UnmarshalDataSetRecords unmarshals an instance of DataSetRecords from the specified map of raw messages.
func UnmarshalDataSetRecords(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetRecords)
	err = core.UnmarshalModel(m, "data_set_records", &obj.DataSetRecords, UnmarshalDataSetRecordsDataSetRecordsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetRecordsDataSetRecordsItem : DataSetRecordsDataSetRecordsItem struct
type DataSetRecordsDataSetRecordsItem struct {
	DataSet *DataSetRecordsDataSetRecordsItemDataSet `json:"data_set" validate:"required"`

	Records *DataSetRecordsDataSetRecordsItemRecords `json:"records" validate:"required"`
}

// UnmarshalDataSetRecordsDataSetRecordsItem unmarshals an instance of DataSetRecordsDataSetRecordsItem from the specified map of raw messages.
func UnmarshalDataSetRecordsDataSetRecordsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetRecordsDataSetRecordsItem)
	err = core.UnmarshalModel(m, "data_set", &obj.DataSet, UnmarshalDataSetRecordsDataSetRecordsItemDataSet)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "records", &obj.Records, UnmarshalDataSetRecordsDataSetRecordsItemRecords)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetRecordsDataSetRecordsItemDataSet : DataSetRecordsDataSetRecordsItemDataSet struct
type DataSetRecordsDataSetRecordsItemDataSet struct {
	DataMartID *string `json:"data_mart_id" validate:"required"`

	DataSchema *SparkStruct `json:"data_schema" validate:"required"`

	Target *Target `json:"target" validate:"required"`

	// type of a data set.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the DataSetRecordsDataSetRecordsItemDataSet.Type property.
// type of a data set.
const (
	DataSetRecordsDataSetRecordsItemDataSet_Type_BusinessPayload    = "business_payload"
	DataSetRecordsDataSetRecordsItemDataSet_Type_Custom             = "custom"
	DataSetRecordsDataSetRecordsItemDataSet_Type_Explanations       = "explanations"
	DataSetRecordsDataSetRecordsItemDataSet_Type_ExplanationsWhatif = "explanations_whatif"
	DataSetRecordsDataSetRecordsItemDataSet_Type_Feedback           = "feedback"
	DataSetRecordsDataSetRecordsItemDataSet_Type_ManualLabeling     = "manual_labeling"
	DataSetRecordsDataSetRecordsItemDataSet_Type_PayloadLogging     = "payload_logging"
	DataSetRecordsDataSetRecordsItemDataSet_Type_Training           = "training"
)

// UnmarshalDataSetRecordsDataSetRecordsItemDataSet unmarshals an instance of DataSetRecordsDataSetRecordsItemDataSet from the specified map of raw messages.
func UnmarshalDataSetRecordsDataSetRecordsItemDataSet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetRecordsDataSetRecordsItemDataSet)
	err = core.UnmarshalPrimitive(m, "data_mart_id", &obj.DataMartID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_schema", &obj.DataSchema, UnmarshalSparkStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetRecordsDataSetRecordsItemRecords : DataSetRecordsDataSetRecordsItemRecords struct
type DataSetRecordsDataSetRecordsItemRecords struct {
	Entity *DataSetRecordsDataSetRecordsItemRecordsEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalDataSetRecordsDataSetRecordsItemRecords unmarshals an instance of DataSetRecordsDataSetRecordsItemRecords from the specified map of raw messages.
func UnmarshalDataSetRecordsDataSetRecordsItemRecords(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetRecordsDataSetRecordsItemRecords)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalDataSetRecordsDataSetRecordsItemRecordsEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetRecordsDataSetRecordsItemRecordsEntity : DataSetRecordsDataSetRecordsItemRecordsEntity struct
type DataSetRecordsDataSetRecordsItemRecordsEntity struct {
	// Fields and values of the entity matches JSON object's fields and values.
	Values *JsDictElem `json:"values" validate:"required"`
}

// UnmarshalDataSetRecordsDataSetRecordsItemRecordsEntity unmarshals an instance of DataSetRecordsDataSetRecordsItemRecordsEntity from the specified map of raw messages.
func UnmarshalDataSetRecordsDataSetRecordsItemRecordsEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetRecordsDataSetRecordsItemRecordsEntity)
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalJsDictElem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetResponse : DataSetResponse struct
type DataSetResponse struct {
	Entity *DataSetObject `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalDataSetResponse unmarshals an instance of DataSetResponse from the specified map of raw messages.
func UnmarshalDataSetResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalDataSetObject)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetResponseCollection : DataSetResponseCollection struct
type DataSetResponseCollection struct {
	DataSets []DataSetResponse `json:"data_sets" validate:"required"`
}

// UnmarshalDataSetResponseCollection unmarshals an instance of DataSetResponseCollection from the specified map of raw messages.
func UnmarshalDataSetResponseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSetResponseCollection)
	err = core.UnmarshalModel(m, "data_sets", &obj.DataSets, UnmarshalDataSetResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSetsAddOptions : The DataSetsAdd options.
type DataSetsAddOptions struct {
	DataMartID *string `json:"data_mart_id" validate:"required"`

	DataSchema *SparkStruct `json:"data_schema" validate:"required"`

	Name *string `json:"name" validate:"required"`

	Target *Target `json:"target" validate:"required"`

	// type of a data set.
	Type *string `json:"type" validate:"required"`

	Description *string `json:"description,omitempty"`

	Location *LocationTableName `json:"location,omitempty"`

	ManagedBy *string `json:"managed_by,omitempty"`

	SchemaUpdateMode *string `json:"schema_update_mode,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DataSetsAddOptions.Type property.
// type of a data set.
const (
	DataSetsAddOptions_Type_BusinessPayload    = "business_payload"
	DataSetsAddOptions_Type_Custom             = "custom"
	DataSetsAddOptions_Type_Explanations       = "explanations"
	DataSetsAddOptions_Type_ExplanationsWhatif = "explanations_whatif"
	DataSetsAddOptions_Type_Feedback           = "feedback"
	DataSetsAddOptions_Type_ManualLabeling     = "manual_labeling"
	DataSetsAddOptions_Type_PayloadLogging     = "payload_logging"
	DataSetsAddOptions_Type_Training           = "training"
)

// Constants associated with the DataSetsAddOptions.SchemaUpdateMode property.
const (
	DataSetsAddOptions_SchemaUpdateMode_Auto = "auto"
	DataSetsAddOptions_SchemaUpdateMode_None = "none"
)

// NewDataSetsAddOptions : Instantiate DataSetsAddOptions
func (*WatsonOpenScaleV2) NewDataSetsAddOptions(dataMartID string, dataSchema *SparkStruct, name string, target *Target, typeVar string) *DataSetsAddOptions {
	return &DataSetsAddOptions{
		DataMartID: core.StringPtr(dataMartID),
		DataSchema: dataSchema,
		Name:       core.StringPtr(name),
		Target:     target,
		Type:       core.StringPtr(typeVar),
	}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *DataSetsAddOptions) SetDataMartID(dataMartID string) *DataSetsAddOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetDataSchema : Allow user to set DataSchema
func (_options *DataSetsAddOptions) SetDataSchema(dataSchema *SparkStruct) *DataSetsAddOptions {
	_options.DataSchema = dataSchema
	return _options
}

// SetName : Allow user to set Name
func (_options *DataSetsAddOptions) SetName(name string) *DataSetsAddOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *DataSetsAddOptions) SetTarget(target *Target) *DataSetsAddOptions {
	_options.Target = target
	return _options
}

// SetType : Allow user to set Type
func (_options *DataSetsAddOptions) SetType(typeVar string) *DataSetsAddOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *DataSetsAddOptions) SetDescription(description string) *DataSetsAddOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetLocation : Allow user to set Location
func (_options *DataSetsAddOptions) SetLocation(location *LocationTableName) *DataSetsAddOptions {
	_options.Location = location
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *DataSetsAddOptions) SetManagedBy(managedBy string) *DataSetsAddOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetSchemaUpdateMode : Allow user to set SchemaUpdateMode
func (_options *DataSetsAddOptions) SetSchemaUpdateMode(schemaUpdateMode string) *DataSetsAddOptions {
	_options.SchemaUpdateMode = core.StringPtr(schemaUpdateMode)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataSetsAddOptions) SetHeaders(param map[string]string) *DataSetsAddOptions {
	options.Headers = param
	return options
}

// DataSetsDeleteOptions : The DataSetsDelete options.
type DataSetsDeleteOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDataSetsDeleteOptions : Instantiate DataSetsDeleteOptions
func (*WatsonOpenScaleV2) NewDataSetsDeleteOptions(dataSetID string) *DataSetsDeleteOptions {
	return &DataSetsDeleteOptions{
		DataSetID: core.StringPtr(dataSetID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *DataSetsDeleteOptions) SetDataSetID(dataSetID string) *DataSetsDeleteOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataSetsDeleteOptions) SetHeaders(param map[string]string) *DataSetsDeleteOptions {
	options.Headers = param
	return options
}

// DataSetsGetOptions : The DataSetsGet options.
type DataSetsGetOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDataSetsGetOptions : Instantiate DataSetsGetOptions
func (*WatsonOpenScaleV2) NewDataSetsGetOptions(dataSetID string) *DataSetsGetOptions {
	return &DataSetsGetOptions{
		DataSetID: core.StringPtr(dataSetID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *DataSetsGetOptions) SetDataSetID(dataSetID string) *DataSetsGetOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataSetsGetOptions) SetHeaders(param map[string]string) *DataSetsGetOptions {
	options.Headers = param
	return options
}

// DataSetsListOptions : The DataSetsList options.
type DataSetsListOptions struct {
	// ID of the data set target (e.g. subscription ID, business application ID).
	TargetTargetID *string `json:"target.target_id,omitempty"`

	// type of the target.
	TargetTargetType *string `json:"target.target_type,omitempty"`

	// type of the data set.
	Type *string `json:"type,omitempty"`

	// ID of the managing entity (e.g. business application ID).
	ManagedBy *string `json:"managed_by,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DataSetsListOptions.TargetTargetType property.
// type of the target.
const (
	DataSetsListOptions_TargetTargetType_BusinessApplication = "business_application"
	DataSetsListOptions_TargetTargetType_DataMart            = "data_mart"
	DataSetsListOptions_TargetTargetType_Instance            = "instance"
	DataSetsListOptions_TargetTargetType_Subscription        = "subscription"
)

// Constants associated with the DataSetsListOptions.Type property.
// type of the data set.
const (
	DataSetsListOptions_Type_BusinessPayload    = "business_payload"
	DataSetsListOptions_Type_Custom             = "custom"
	DataSetsListOptions_Type_Explanations       = "explanations"
	DataSetsListOptions_Type_ExplanationsWhatif = "explanations_whatif"
	DataSetsListOptions_Type_Feedback           = "feedback"
	DataSetsListOptions_Type_ManualLabeling     = "manual_labeling"
	DataSetsListOptions_Type_PayloadLogging     = "payload_logging"
	DataSetsListOptions_Type_Training           = "training"
)

// NewDataSetsListOptions : Instantiate DataSetsListOptions
func (*WatsonOpenScaleV2) NewDataSetsListOptions() *DataSetsListOptions {
	return &DataSetsListOptions{}
}

// SetTargetTargetID : Allow user to set TargetTargetID
func (_options *DataSetsListOptions) SetTargetTargetID(targetTargetID string) *DataSetsListOptions {
	_options.TargetTargetID = core.StringPtr(targetTargetID)
	return _options
}

// SetTargetTargetType : Allow user to set TargetTargetType
func (_options *DataSetsListOptions) SetTargetTargetType(targetTargetType string) *DataSetsListOptions {
	_options.TargetTargetType = core.StringPtr(targetTargetType)
	return _options
}

// SetType : Allow user to set Type
func (_options *DataSetsListOptions) SetType(typeVar string) *DataSetsListOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *DataSetsListOptions) SetManagedBy(managedBy string) *DataSetsListOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataSetsListOptions) SetHeaders(param map[string]string) *DataSetsListOptions {
	options.Headers = param
	return options
}

// DataSetsUpdateOptions : The DataSetsUpdate options.
type DataSetsUpdateOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	PatchDocument []PatchDocument `json:"PatchDocument" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDataSetsUpdateOptions : Instantiate DataSetsUpdateOptions
func (*WatsonOpenScaleV2) NewDataSetsUpdateOptions(dataSetID string, patchDocument []PatchDocument) *DataSetsUpdateOptions {
	return &DataSetsUpdateOptions{
		DataSetID:     core.StringPtr(dataSetID),
		PatchDocument: patchDocument,
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *DataSetsUpdateOptions) SetDataSetID(dataSetID string) *DataSetsUpdateOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetPatchDocument : Allow user to set PatchDocument
func (_options *DataSetsUpdateOptions) SetPatchDocument(patchDocument []PatchDocument) *DataSetsUpdateOptions {
	_options.PatchDocument = patchDocument
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DataSetsUpdateOptions) SetHeaders(param map[string]string) *DataSetsUpdateOptions {
	options.Headers = param
	return options
}

// DataSource : DataSource struct
type DataSource struct {
	Connection *DataSourceConnection `json:"connection,omitempty"`

	// database name.
	DatabaseName *string `json:"database_name,omitempty"`

	Endpoint *DataSourceEndpoint `json:"endpoint,omitempty"`

	// Additional parameters.
	Parameters interface{} `json:"parameters,omitempty"`

	// schema name.
	SchemaName *string `json:"schema_name,omitempty"`

	// table name.
	TableName *string `json:"table_name,omitempty"`

	// Type of data source. e.g. payload, feedback, drift,explain.
	Type *string `json:"type,omitempty"`
}

// UnmarshalDataSource unmarshals an instance of DataSource from the specified map of raw messages.
func UnmarshalDataSource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSource)
	err = core.UnmarshalModel(m, "connection", &obj.Connection, UnmarshalDataSourceConnection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "endpoint", &obj.Endpoint, UnmarshalDataSourceEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schema_name", &obj.SchemaName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "table_name", &obj.TableName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSourceConnection : DataSourceConnection struct
type DataSourceConnection struct {
	// id of the integrated system.
	IntegratedSystemID *string `json:"integrated_system_id,omitempty"`

	// Additional parameters.
	Parameters interface{} `json:"parameters,omitempty"`

	// Type of integrated system, e.g. hive.
	Type *string `json:"type,omitempty"`
}

// UnmarshalDataSourceConnection unmarshals an instance of DataSourceConnection from the specified map of raw messages.
func UnmarshalDataSourceConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSourceConnection)
	err = core.UnmarshalPrimitive(m, "integrated_system_id", &obj.IntegratedSystemID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSourceEndpoint : DataSourceEndpoint struct
type DataSourceEndpoint struct {
	// Credentials for the endpoint.
	Credentials interface{} `json:"credentials,omitempty"`

	// Url of the endpoint.
	URL *string `json:"url" validate:"required"`
}

// NewDataSourceEndpoint : Instantiate DataSourceEndpoint (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewDataSourceEndpoint(url string) (_model *DataSourceEndpoint, err error) {
	_model = &DataSourceEndpoint{
		URL: core.StringPtr(url),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDataSourceEndpoint unmarshals an instance of DataSourceEndpoint from the specified map of raw messages.
func UnmarshalDataSourceEndpoint(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSourceEndpoint)
	err = core.UnmarshalPrimitive(m, "credentials", &obj.Credentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseConfiguration : Database configuration ignored if internal database is requested (`internal_database` is `true`).
type DatabaseConfiguration struct {
	Credentials *SecretCleaned `json:"credentials" validate:"required"`

	DatabaseType *string `json:"database_type" validate:"required"`

	InstanceID *string `json:"instance_id,omitempty"`

	Location *LocationSchemaName `json:"location,omitempty"`

	Name *string `json:"name,omitempty"`
}

// Constants associated with the DatabaseConfiguration.DatabaseType property.
const (
	DatabaseConfiguration_DatabaseType_Db2        = "db2"
	DatabaseConfiguration_DatabaseType_Postgresql = "postgresql"
)

// UnmarshalDatabaseConfiguration unmarshals an instance of DatabaseConfiguration from the specified map of raw messages.
func UnmarshalDatabaseConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseConfiguration)
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalSecretCleaned)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_type", &obj.DatabaseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocationSchemaName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatabaseConfigurationRequest : Database configuration ignored if internal database is requested (`internal_database` is `true`).
type DatabaseConfigurationRequest struct {
	Credentials PrimaryStorageCredentialsIntf `json:"credentials" validate:"required"`

	DatabaseType *string `json:"database_type" validate:"required"`

	InstanceID *string `json:"instance_id,omitempty"`

	Location *LocationSchemaName `json:"location,omitempty"`

	Name *string `json:"name,omitempty"`
}

// Constants associated with the DatabaseConfigurationRequest.DatabaseType property.
const (
	DatabaseConfigurationRequest_DatabaseType_Db2        = "db2"
	DatabaseConfigurationRequest_DatabaseType_Postgresql = "postgresql"
)

// NewDatabaseConfigurationRequest : Instantiate DatabaseConfigurationRequest (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewDatabaseConfigurationRequest(credentials PrimaryStorageCredentialsIntf, databaseType string) (_model *DatabaseConfigurationRequest, err error) {
	_model = &DatabaseConfigurationRequest{
		Credentials:  credentials,
		DatabaseType: core.StringPtr(databaseType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDatabaseConfigurationRequest unmarshals an instance of DatabaseConfigurationRequest from the specified map of raw messages.
func UnmarshalDatabaseConfigurationRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatabaseConfigurationRequest)
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalPrimaryStorageCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_type", &obj.DatabaseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalLocationSchemaName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatasetRecordsPayloadItem : DatasetRecordsPayloadItem struct
// Models which "extend" this model:
// - DatasetRecordsPayloadItemJsDictElem
// - DatasetRecordsPayloadItemJsList
// - DatasetRecordsPayloadItemScoringPayloadRequest
type DatasetRecordsPayloadItem struct {
	// Fields' names are listed in order in 'fields'.
	Fields []string `json:"fields,omitempty"`

	// Rows organized per fields as described in 'fields'.
	Values [][]interface{} `json:"values,omitempty"`

	AssetRevision *string `json:"asset_revision,omitempty"`

	// This property is important for payload logging for unstructered data. There will be one record created in payload
	// logging table if this property is set to `false`. The `request.values` field and `response.values` field are treated
	// as one scoring request and one scoring response in such case. If this field is set to `true` then it might be
	// created more than one row in the payload logging table. The first dimension of the `request.values` and
	// `response.values` correponds to the separate entry in the payload logging table in such case.
	MultipleRecords *bool `json:"multiple_records,omitempty"`

	Request *ScoringPayloadRequestRequest `json:"request,omitempty"`

	Response *ScoringPayloadRequestResponse `json:"response,omitempty"`

	ResponseTime *float64 `json:"response_time,omitempty"`

	ScoringID *string `json:"scoring_id,omitempty"`

	ScoringTimestamp *string `json:"scoring_timestamp,omitempty"`
}

func (*DatasetRecordsPayloadItem) isaDatasetRecordsPayloadItem() bool {
	return true
}

type DatasetRecordsPayloadItemIntf interface {
	isaDatasetRecordsPayloadItem() bool
}

// UnmarshalDatasetRecordsPayloadItem unmarshals an instance of DatasetRecordsPayloadItem from the specified map of raw messages.
func UnmarshalDatasetRecordsPayloadItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatasetRecordsPayloadItem)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_revision", &obj.AssetRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "multiple_records", &obj.MultipleRecords)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "request", &obj.Request, UnmarshalScoringPayloadRequestRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalScoringPayloadRequestResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "response_time", &obj.ResponseTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scoring_id", &obj.ScoringID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scoring_timestamp", &obj.ScoringTimestamp)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DistributionsAddOptions : The DistributionsAdd options.
type DistributionsAddOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// type of a data set.
	Dataset *string `json:"dataset" validate:"required"`

	// end datetime in ISO format.
	End *string `json:"end" validate:"required"`

	// names of columns to be grouped.
	Group []string `json:"group" validate:"required"`

	// start datetime in ISO format.
	Start *string `json:"start" validate:"required"`

	// Definition of aggregations, by default 'count'.
	//
	// Aggregations can be one of:
	// * count
	// * <column_name>:sum
	// * <column_name>:min
	// * <column_name>:max
	// * <column_name>:avg
	// * <column_name>:stddev.
	Agg []string `json:"agg,omitempty"`

	// Filters defined by user in format: {field_name}:{op}:{value}. Partly compatible with filters in "filter" parameter
	// of GET /v2/data_sets/{data_set_id}/records.
	//
	// Possible filter operators:
	// * eq - equals (numeric, string)
	// * gt - greater than (numeric)
	// * gte - greater than or equal (numeric)
	// * lt - lower than (numeric)
	// * lte - lower than or equal (numeric)
	// * in - value in a set (numeric, string)
	// * field:null (a no-argument filter) - value is null (any nullable)
	// * field:exists (a no-argument filter) - value is not null (any column).
	Filter *string `json:"filter,omitempty"`

	// limit for number of rows, by default it is 50,000 (max possible limit is 50,000).
	Limit *float64 `json:"limit,omitempty"`

	// max number of bins which will be generated for data.
	MaxBins *float64 `json:"max_bins,omitempty"`

	// force columns data refresh.
	Nocache *bool `json:"nocache,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DistributionsAddOptions.Dataset property.
// type of a data set.
const (
	DistributionsAddOptions_Dataset_BusinessPayload    = "business_payload"
	DistributionsAddOptions_Dataset_Custom             = "custom"
	DistributionsAddOptions_Dataset_Explanations       = "explanations"
	DistributionsAddOptions_Dataset_ExplanationsWhatif = "explanations_whatif"
	DistributionsAddOptions_Dataset_Feedback           = "feedback"
	DistributionsAddOptions_Dataset_ManualLabeling     = "manual_labeling"
	DistributionsAddOptions_Dataset_PayloadLogging     = "payload_logging"
	DistributionsAddOptions_Dataset_Training           = "training"
)

// NewDistributionsAddOptions : Instantiate DistributionsAddOptions
func (*WatsonOpenScaleV2) NewDistributionsAddOptions(dataSetID string, dataset string, end string, group []string, start string) *DistributionsAddOptions {
	return &DistributionsAddOptions{
		DataSetID: core.StringPtr(dataSetID),
		Dataset:   core.StringPtr(dataset),
		End:       core.StringPtr(end),
		Group:     group,
		Start:     core.StringPtr(start),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *DistributionsAddOptions) SetDataSetID(dataSetID string) *DistributionsAddOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetDataset : Allow user to set Dataset
func (_options *DistributionsAddOptions) SetDataset(dataset string) *DistributionsAddOptions {
	_options.Dataset = core.StringPtr(dataset)
	return _options
}

// SetEnd : Allow user to set End
func (_options *DistributionsAddOptions) SetEnd(end string) *DistributionsAddOptions {
	_options.End = core.StringPtr(end)
	return _options
}

// SetGroup : Allow user to set Group
func (_options *DistributionsAddOptions) SetGroup(group []string) *DistributionsAddOptions {
	_options.Group = group
	return _options
}

// SetStart : Allow user to set Start
func (_options *DistributionsAddOptions) SetStart(start string) *DistributionsAddOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetAgg : Allow user to set Agg
func (_options *DistributionsAddOptions) SetAgg(agg []string) *DistributionsAddOptions {
	_options.Agg = agg
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *DistributionsAddOptions) SetFilter(filter string) *DistributionsAddOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *DistributionsAddOptions) SetLimit(limit float64) *DistributionsAddOptions {
	_options.Limit = core.Float64Ptr(limit)
	return _options
}

// SetMaxBins : Allow user to set MaxBins
func (_options *DistributionsAddOptions) SetMaxBins(maxBins float64) *DistributionsAddOptions {
	_options.MaxBins = core.Float64Ptr(maxBins)
	return _options
}

// SetNocache : Allow user to set Nocache
func (_options *DistributionsAddOptions) SetNocache(nocache bool) *DistributionsAddOptions {
	_options.Nocache = core.BoolPtr(nocache)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DistributionsAddOptions) SetHeaders(param map[string]string) *DistributionsAddOptions {
	options.Headers = param
	return options
}

// DistributionsDeleteOptions : The DistributionsDelete options.
type DistributionsDeleteOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDistributionsDeleteOptions : Instantiate DistributionsDeleteOptions
func (*WatsonOpenScaleV2) NewDistributionsDeleteOptions(dataSetID string) *DistributionsDeleteOptions {
	return &DistributionsDeleteOptions{
		DataSetID: core.StringPtr(dataSetID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *DistributionsDeleteOptions) SetDataSetID(dataSetID string) *DistributionsDeleteOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DistributionsDeleteOptions) SetHeaders(param map[string]string) *DistributionsDeleteOptions {
	options.Headers = param
	return options
}

// DistributionsGetOptions : The DistributionsGet options.
type DistributionsGetOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// ID of the data distribution requested to be calculated.
	DataDistributionID *string `json:"data_distribution_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDistributionsGetOptions : Instantiate DistributionsGetOptions
func (*WatsonOpenScaleV2) NewDistributionsGetOptions(dataSetID string, dataDistributionID string) *DistributionsGetOptions {
	return &DistributionsGetOptions{
		DataSetID:          core.StringPtr(dataSetID),
		DataDistributionID: core.StringPtr(dataDistributionID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *DistributionsGetOptions) SetDataSetID(dataSetID string) *DistributionsGetOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetDataDistributionID : Allow user to set DataDistributionID
func (_options *DistributionsGetOptions) SetDataDistributionID(dataDistributionID string) *DistributionsGetOptions {
	_options.DataDistributionID = core.StringPtr(dataDistributionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DistributionsGetOptions) SetHeaders(param map[string]string) *DistributionsGetOptions {
	options.Headers = param
	return options
}

// DriftArchiveGetOptions : The DriftArchiveGet options.
type DriftArchiveGetOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDriftArchiveGetOptions : Instantiate DriftArchiveGetOptions
func (*WatsonOpenScaleV2) NewDriftArchiveGetOptions(monitorInstanceID string) *DriftArchiveGetOptions {
	return &DriftArchiveGetOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *DriftArchiveGetOptions) SetMonitorInstanceID(monitorInstanceID string) *DriftArchiveGetOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DriftArchiveGetOptions) SetHeaders(param map[string]string) *DriftArchiveGetOptions {
	options.Headers = param
	return options
}

// DriftArchiveHeadOptions : The DriftArchiveHead options.
type DriftArchiveHeadOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDriftArchiveHeadOptions : Instantiate DriftArchiveHeadOptions
func (*WatsonOpenScaleV2) NewDriftArchiveHeadOptions(monitorInstanceID string) *DriftArchiveHeadOptions {
	return &DriftArchiveHeadOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *DriftArchiveHeadOptions) SetMonitorInstanceID(monitorInstanceID string) *DriftArchiveHeadOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DriftArchiveHeadOptions) SetHeaders(param map[string]string) *DriftArchiveHeadOptions {
	options.Headers = param
	return options
}

// DriftArchivePostOptions : The DriftArchivePost options.
type DriftArchivePostOptions struct {
	// ID of the data mart.
	DataMartID *string `json:"data_mart_id" validate:"required,ne="`

	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id" validate:"required,ne="`

	Body io.ReadCloser `json:"body" validate:"required"`

	// The name of the archive being uploaded.
	ArchiveName *string `json:"archive_name,omitempty"`

	// Flag to enable/disable data drift.
	EnableDataDrift *bool `json:"enable_data_drift,omitempty"`

	// Flag to enable/disable model drift.
	EnableModelDrift *bool `json:"enable_model_drift,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDriftArchivePostOptions : Instantiate DriftArchivePostOptions
func (*WatsonOpenScaleV2) NewDriftArchivePostOptions(dataMartID string, subscriptionID string, body io.ReadCloser) *DriftArchivePostOptions {
	return &DriftArchivePostOptions{
		DataMartID:     core.StringPtr(dataMartID),
		SubscriptionID: core.StringPtr(subscriptionID),
		Body:           body,
	}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *DriftArchivePostOptions) SetDataMartID(dataMartID string) *DriftArchivePostOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *DriftArchivePostOptions) SetSubscriptionID(subscriptionID string) *DriftArchivePostOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *DriftArchivePostOptions) SetBody(body io.ReadCloser) *DriftArchivePostOptions {
	_options.Body = body
	return _options
}

// SetArchiveName : Allow user to set ArchiveName
func (_options *DriftArchivePostOptions) SetArchiveName(archiveName string) *DriftArchivePostOptions {
	_options.ArchiveName = core.StringPtr(archiveName)
	return _options
}

// SetEnableDataDrift : Allow user to set EnableDataDrift
func (_options *DriftArchivePostOptions) SetEnableDataDrift(enableDataDrift bool) *DriftArchivePostOptions {
	_options.EnableDataDrift = core.BoolPtr(enableDataDrift)
	return _options
}

// SetEnableModelDrift : Allow user to set EnableModelDrift
func (_options *DriftArchivePostOptions) SetEnableModelDrift(enableModelDrift bool) *DriftArchivePostOptions {
	_options.EnableModelDrift = core.BoolPtr(enableModelDrift)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DriftArchivePostOptions) SetHeaders(param map[string]string) *DriftArchivePostOptions {
	options.Headers = param
	return options
}

// ExplanationFeature : Explanation feature details.
type ExplanationFeature struct {
	// Name of the feature column.
	FeatureName *string `json:"feature_name,omitempty"`

	// Range of feature values.
	FeatureRange *ExplanationFeatureFeatureRange `json:"feature_range,omitempty"`

	// Value of the feature column.
	FeatureValue *string `json:"feature_value,omitempty"`

	// Contributing importance to the output explanation for the given feature.
	Importance *string `json:"importance,omitempty"`

	// Contributing weight to the output explanation for the given feature.
	Weight *float64 `json:"weight,omitempty"`
}

// UnmarshalExplanationFeature unmarshals an instance of ExplanationFeature from the specified map of raw messages.
func UnmarshalExplanationFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExplanationFeature)
	err = core.UnmarshalPrimitive(m, "feature_name", &obj.FeatureName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "feature_range", &obj.FeatureRange, UnmarshalExplanationFeatureFeatureRange)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "feature_value", &obj.FeatureValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "importance", &obj.Importance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "weight", &obj.Weight)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExplanationFeatureFeatureRange : Range of feature values.
type ExplanationFeatureFeatureRange struct {
	// Maximum possible value for given feature.
	Max *string `json:"max,omitempty"`

	// Identifies if the maximum value is inclusive or not.
	MaxInclusive *bool `json:"max_inclusive,omitempty"`

	// Minimum possible value for given feature.
	Min *string `json:"min,omitempty"`

	// Identifies if the minimum value is inclusive or not.
	MinInclusive *bool `json:"min_inclusive,omitempty"`
}

// UnmarshalExplanationFeatureFeatureRange unmarshals an instance of ExplanationFeatureFeatureRange from the specified map of raw messages.
func UnmarshalExplanationFeatureFeatureRange(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExplanationFeatureFeatureRange)
	err = core.UnmarshalPrimitive(m, "max", &obj.Max)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_inclusive", &obj.MaxInclusive)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min", &obj.Min)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min_inclusive", &obj.MinInclusive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExplanationTaskResponseEntityAsset : Asset details in get explanation task response.
type ExplanationTaskResponseEntityAsset struct {
	// Asset deployment details.
	Deployment *ExplanationTaskResponseEntityAssetDeployment `json:"deployment,omitempty"`

	// Identifier for the asset.
	ID *string `json:"id,omitempty"`

	// Type of the input data.
	InputDataType *string `json:"input_data_type,omitempty"`

	// Name of the asset.
	Name *string `json:"name,omitempty"`

	// Problem type.
	ProblemType *string `json:"problem_type,omitempty"`
}

// Constants associated with the ExplanationTaskResponseEntityAsset.InputDataType property.
// Type of the input data.
const (
	ExplanationTaskResponseEntityAsset_InputDataType_Structured        = "structured"
	ExplanationTaskResponseEntityAsset_InputDataType_UnstructuredImage = "unstructured_image"
	ExplanationTaskResponseEntityAsset_InputDataType_UnstructuredText  = "unstructured_text"
)

// Constants associated with the ExplanationTaskResponseEntityAsset.ProblemType property.
// Problem type.
const (
	ExplanationTaskResponseEntityAsset_ProblemType_Binary     = "binary"
	ExplanationTaskResponseEntityAsset_ProblemType_Multiclass = "multiclass"
	ExplanationTaskResponseEntityAsset_ProblemType_Regression = "regression"
)

// UnmarshalExplanationTaskResponseEntityAsset unmarshals an instance of ExplanationTaskResponseEntityAsset from the specified map of raw messages.
func UnmarshalExplanationTaskResponseEntityAsset(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExplanationTaskResponseEntityAsset)
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalExplanationTaskResponseEntityAssetDeployment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_data_type", &obj.InputDataType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "problem_type", &obj.ProblemType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExplanationTaskResponseEntityAssetDeployment : Asset deployment details.
type ExplanationTaskResponseEntityAssetDeployment struct {
	// Identifier for the asset deployment.
	ID *string `json:"id,omitempty"`

	// Name of the asset deployment.
	Name *string `json:"name,omitempty"`
}

// UnmarshalExplanationTaskResponseEntityAssetDeployment unmarshals an instance of ExplanationTaskResponseEntityAssetDeployment from the specified map of raw messages.
func UnmarshalExplanationTaskResponseEntityAssetDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExplanationTaskResponseEntityAssetDeployment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExplanationTaskResponseEntityInputFeature : Input feature details in get explanation task response.
type ExplanationTaskResponseEntityInputFeature struct {
	// Identifies the type of feature column.
	FeatureType *string `json:"feature_type,omitempty"`

	// Name of the feature column.
	Name *string `json:"name,omitempty"`

	// Value of the feature column.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the ExplanationTaskResponseEntityInputFeature.FeatureType property.
// Identifies the type of feature column.
const (
	ExplanationTaskResponseEntityInputFeature_FeatureType_Categorical = "categorical"
	ExplanationTaskResponseEntityInputFeature_FeatureType_Numerical   = "numerical"
)

// UnmarshalExplanationTaskResponseEntityInputFeature unmarshals an instance of ExplanationTaskResponseEntityInputFeature from the specified map of raw messages.
func UnmarshalExplanationTaskResponseEntityInputFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExplanationTaskResponseEntityInputFeature)
	err = core.UnmarshalPrimitive(m, "feature_type", &obj.FeatureType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExplanationTasksAddOptions : The ExplanationTasksAdd options.
type ExplanationTasksAddOptions struct {
	// IDs of the scoring transaction.
	ScoringIds []string `json:"scoring_ids" validate:"required"`

	// Types of explanations to generate.
	ExplanationTypes []string `json:"explanation_types,omitempty"`

	// List of scoring transactions.
	InputRows []map[string]interface{} `json:"input_rows,omitempty"`

	// List of subscription ids.
	SubscriptionIds []string `json:"subscription_ids,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ExplanationTasksAddOptions.ExplanationTypes property.
const (
	ExplanationTasksAddOptions_ExplanationTypes_Contrastive = "contrastive"
	ExplanationTasksAddOptions_ExplanationTypes_Lime        = "lime"
)

// NewExplanationTasksAddOptions : Instantiate ExplanationTasksAddOptions
func (*WatsonOpenScaleV2) NewExplanationTasksAddOptions(scoringIds []string) *ExplanationTasksAddOptions {
	return &ExplanationTasksAddOptions{
		ScoringIds: scoringIds,
	}
}

// SetScoringIds : Allow user to set ScoringIds
func (_options *ExplanationTasksAddOptions) SetScoringIds(scoringIds []string) *ExplanationTasksAddOptions {
	_options.ScoringIds = scoringIds
	return _options
}

// SetExplanationTypes : Allow user to set ExplanationTypes
func (_options *ExplanationTasksAddOptions) SetExplanationTypes(explanationTypes []string) *ExplanationTasksAddOptions {
	_options.ExplanationTypes = explanationTypes
	return _options
}

// SetInputRows : Allow user to set InputRows
func (_options *ExplanationTasksAddOptions) SetInputRows(inputRows []map[string]interface{}) *ExplanationTasksAddOptions {
	_options.InputRows = inputRows
	return _options
}

// SetSubscriptionIds : Allow user to set SubscriptionIds
func (_options *ExplanationTasksAddOptions) SetSubscriptionIds(subscriptionIds []string) *ExplanationTasksAddOptions {
	_options.SubscriptionIds = subscriptionIds
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExplanationTasksAddOptions) SetHeaders(param map[string]string) *ExplanationTasksAddOptions {
	options.Headers = param
	return options
}

// ExplanationTasksGetOptions : The ExplanationTasksGet options.
type ExplanationTasksGetOptions struct {
	// ID of the explanation task.
	ExplanationTaskID *string `json:"explanation_task_id" validate:"required,ne="`

	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExplanationTasksGetOptions : Instantiate ExplanationTasksGetOptions
func (*WatsonOpenScaleV2) NewExplanationTasksGetOptions(explanationTaskID string) *ExplanationTasksGetOptions {
	return &ExplanationTasksGetOptions{
		ExplanationTaskID: core.StringPtr(explanationTaskID),
	}
}

// SetExplanationTaskID : Allow user to set ExplanationTaskID
func (_options *ExplanationTasksGetOptions) SetExplanationTaskID(explanationTaskID string) *ExplanationTasksGetOptions {
	_options.ExplanationTaskID = core.StringPtr(explanationTaskID)
	return _options
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *ExplanationTasksGetOptions) SetSubscriptionID(subscriptionID string) *ExplanationTasksGetOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExplanationTasksGetOptions) SetHeaders(param map[string]string) *ExplanationTasksGetOptions {
	options.Headers = param
	return options
}

// ExplanationTasksListOptions : The ExplanationTasksList options.
type ExplanationTasksListOptions struct {
	// offset of the explanations to return.
	Offset *int64 `json:"offset,omitempty"`

	// Maximum number of explanations to return.
	Limit *int64 `json:"limit,omitempty"`

	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExplanationTasksListOptions : Instantiate ExplanationTasksListOptions
func (*WatsonOpenScaleV2) NewExplanationTasksListOptions() *ExplanationTasksListOptions {
	return &ExplanationTasksListOptions{}
}

// SetOffset : Allow user to set Offset
func (_options *ExplanationTasksListOptions) SetOffset(offset int64) *ExplanationTasksListOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ExplanationTasksListOptions) SetLimit(limit int64) *ExplanationTasksListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *ExplanationTasksListOptions) SetSubscriptionID(subscriptionID string) *ExplanationTasksListOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExplanationTasksListOptions) SetHeaders(param map[string]string) *ExplanationTasksListOptions {
	options.Headers = param
	return options
}

// FairnessMonitoringRemediation : FairnessMonitoringRemediation struct
type FairnessMonitoringRemediation struct {
	// The fields of the model processed debias scoring.
	Fields []string `json:"fields" validate:"required"`

	// The values associated to the fields.
	Values []interface{} `json:"values" validate:"required"`
}

// UnmarshalFairnessMonitoringRemediation unmarshals an instance of FairnessMonitoringRemediation from the specified map of raw messages.
func UnmarshalFairnessMonitoringRemediation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FairnessMonitoringRemediation)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FileAssetMetadata : File data asset metadata.
type FileAssetMetadata struct {
	// File data asset reference.
	AssetHref *string `json:"asset_href" validate:"required"`

	// File data asset id.
	AssetID *string `json:"asset_id,omitempty"`

	// File data asset name.
	AssetName *string `json:"asset_name,omitempty"`

	// additional options for different types of training data references.
	Meta *FileTrainingDataReferenceOptions `json:"meta,omitempty"`

	// Project id.
	ProjectID *string `json:"project_id,omitempty"`

	// Project name.
	ProjectName *string `json:"project_name,omitempty"`
}

// UnmarshalFileAssetMetadata unmarshals an instance of FileAssetMetadata from the specified map of raw messages.
func UnmarshalFileAssetMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FileAssetMetadata)
	err = core.UnmarshalPrimitive(m, "asset_href", &obj.AssetHref)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_id", &obj.AssetID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_name", &obj.AssetName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "meta", &obj.Meta, UnmarshalFileTrainingDataReferenceOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_name", &obj.ProjectName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FileTrainingDataReferenceOptions : additional options for different types of training data references.
type FileTrainingDataReferenceOptions struct {
	// maximum length of single line in bytes (default 10000000).
	CsvMaxLineLength *float64 `json:"csv_max_line_length,omitempty"`

	// delimiter character for data provided as csv.
	Delimiter *string `json:"delimiter,omitempty"`

	// File format.
	FileFormat *string `json:"file_format,omitempty"`

	// file name.
	FileName *string `json:"file_name,omitempty"`

	// if not provided service will attempt to automatically detect header in the first line (for data provided as csv).
	FirstLineIsHeader *bool `json:"first_line_is_header,omitempty"`

	// Expected behaviour on error while reading a csv file. Default behaviour is "stop".
	OnError *string `json:"on_error,omitempty"`
}

// Constants associated with the FileTrainingDataReferenceOptions.OnError property.
// Expected behaviour on error while reading a csv file. Default behaviour is "stop".
const (
	FileTrainingDataReferenceOptions_OnError_Continue = "continue"
	FileTrainingDataReferenceOptions_OnError_Stop     = "stop"
)

// UnmarshalFileTrainingDataReferenceOptions unmarshals an instance of FileTrainingDataReferenceOptions from the specified map of raw messages.
func UnmarshalFileTrainingDataReferenceOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FileTrainingDataReferenceOptions)
	err = core.UnmarshalPrimitive(m, "csv_max_line_length", &obj.CsvMaxLineLength)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "delimiter", &obj.Delimiter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_format", &obj.FileFormat)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_name", &obj.FileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first_line_is_header", &obj.FirstLineIsHeader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "on_error", &obj.OnError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericErrorResponse : GenericErrorResponse struct
type GenericErrorResponse struct {
	Errors []GenericErrorResponseErrorsItem `json:"errors" validate:"required"`

	// ID of the original request.
	Trace *string `json:"trace" validate:"required"`
}

// NewGenericErrorResponse : Instantiate GenericErrorResponse (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewGenericErrorResponse(errors []GenericErrorResponseErrorsItem, trace string) (_model *GenericErrorResponse, err error) {
	_model = &GenericErrorResponse{
		Errors: errors,
		Trace:  core.StringPtr(trace),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalGenericErrorResponse unmarshals an instance of GenericErrorResponse from the specified map of raw messages.
func UnmarshalGenericErrorResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericErrorResponse)
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalGenericErrorResponseErrorsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GenericErrorResponseErrorsItem : GenericErrorResponseErrorsItem struct
type GenericErrorResponseErrorsItem struct {
	// Error code.
	Code *string `json:"code" validate:"required"`

	// Error message.
	Message *string `json:"message" validate:"required"`

	// Error message parameters.
	Parameters []string `json:"parameters,omitempty"`
}

// NewGenericErrorResponseErrorsItem : Instantiate GenericErrorResponseErrorsItem (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewGenericErrorResponseErrorsItem(code string, message string) (_model *GenericErrorResponseErrorsItem, err error) {
	_model = &GenericErrorResponseErrorsItem{
		Code:    core.StringPtr(code),
		Message: core.StringPtr(message),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalGenericErrorResponseErrorsItem unmarshals an instance of GenericErrorResponseErrorsItem from the specified map of raw messages.
func UnmarshalGenericErrorResponseErrorsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GenericErrorResponseErrorsItem)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponse : Get explanation task response.
type GetExplanationTaskResponse struct {
	// Entity of get explanation task response.
	Entity *GetExplanationTaskResponseEntity `json:"entity" validate:"required"`

	// Metadata of get explanation task response.
	Metadata *GetExplanationTaskResponseMetadata `json:"metadata" validate:"required"`
}

// UnmarshalGetExplanationTaskResponse unmarshals an instance of GetExplanationTaskResponse from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalGetExplanationTaskResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalGetExplanationTaskResponseMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntity : Entity of get explanation task response.
type GetExplanationTaskResponseEntity struct {
	// Asset details in get explanation task response.
	Asset *ExplanationTaskResponseEntityAsset `json:"asset,omitempty"`

	// List of generated explanations.
	Explanations []GetExplanationTaskResponseEntityExplanationsItemIntf `json:"explanations,omitempty"`

	// List of input features.
	InputFeatures []ExplanationTaskResponseEntityInputFeature `json:"input_features,omitempty"`

	// Indicate whether the transaction is perturbed or not.
	Perturbed *string `json:"perturbed,omitempty"`

	// Status of the explanation task.
	Status *GetExplanationTaskResponseEntityStatus `json:"status" validate:"required"`
}

// Constants associated with the GetExplanationTaskResponseEntity.Perturbed property.
// Indicate whether the transaction is perturbed or not.
const (
	GetExplanationTaskResponseEntity_Perturbed_False = "false"
	GetExplanationTaskResponseEntity_Perturbed_True  = "true"
)

// UnmarshalGetExplanationTaskResponseEntity unmarshals an instance of GetExplanationTaskResponseEntity from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntity)
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalExplanationTaskResponseEntityAsset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "explanations", &obj.Explanations, UnmarshalGetExplanationTaskResponseEntityExplanationsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_features", &obj.InputFeatures, UnmarshalExplanationTaskResponseEntityInputFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "perturbed", &obj.Perturbed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalGetExplanationTaskResponseEntityStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityExplanationsItem : GetExplanationTaskResponseEntityExplanationsItem struct
// Models which "extend" this model:
// - GetExplanationTaskResponseEntityExplanationsItemLimeExplanation
// - GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation
type GetExplanationTaskResponseEntityExplanationsItem struct {
	// Type of the explanation.
	ExplanationType *string `json:"explanation_type,omitempty"`

	// Lime explanations of predictions.
	Predictions []LimeExplanationPrediction `json:"predictions,omitempty"`

	// These factors, if added, would cause the classification to change.
	PertinentNegative *GetExplanationTaskResponseEntityExplanationsItemPertinentNegative `json:"pertinent_negative,omitempty"`

	// These factors are sufficient evidence in themselves to yeild the given classification.
	PertinentPositive *GetExplanationTaskResponseEntityExplanationsItemPertinentPositive `json:"pertinent_positive,omitempty"`
}

func (*GetExplanationTaskResponseEntityExplanationsItem) isaGetExplanationTaskResponseEntityExplanationsItem() bool {
	return true
}

type GetExplanationTaskResponseEntityExplanationsItemIntf interface {
	isaGetExplanationTaskResponseEntityExplanationsItem() bool
}

// UnmarshalGetExplanationTaskResponseEntityExplanationsItem unmarshals an instance of GetExplanationTaskResponseEntityExplanationsItem from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityExplanationsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityExplanationsItem)
	err = core.UnmarshalPrimitive(m, "explanation_type", &obj.ExplanationType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "predictions", &obj.Predictions, UnmarshalLimeExplanationPrediction)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pertinent_negative", &obj.PertinentNegative, UnmarshalGetExplanationTaskResponseEntityExplanationsItemPertinentNegative)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pertinent_positive", &obj.PertinentPositive, UnmarshalGetExplanationTaskResponseEntityExplanationsItemPertinentPositive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative : These factors, if added, would cause the classification to change.
type GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative struct {
	// List of feature names, values and their importance.
	Features []ExplanationFeature `json:"features,omitempty"`

	// Classification of pertinent negative features.
	Prediction *string `json:"prediction,omitempty"`

	// Probability of pertinent negative features.
	Probability *float64 `json:"probability,omitempty"`
}

// UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative unmarshals an instance of GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative)
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalExplanationFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prediction", &obj.Prediction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "probability", &obj.Probability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive : These factors are sufficient evidence in themselves to yeild the given classification.
type GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive struct {
	// List of feature names, values and their importance.
	Features []ExplanationFeature `json:"features,omitempty"`
}

// UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive unmarshals an instance of GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive)
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalExplanationFeature)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityExplanationsItemPertinentNegative : These factors, if added, would cause the classification to change.
type GetExplanationTaskResponseEntityExplanationsItemPertinentNegative struct {
	// List of feature names, values and their importance.
	Features []ExplanationFeature `json:"features,omitempty"`

	// Classification of pertinent negative features.
	Prediction *string `json:"prediction,omitempty"`

	// Probability of pertinent negative features.
	Probability *float64 `json:"probability,omitempty"`
}

// UnmarshalGetExplanationTaskResponseEntityExplanationsItemPertinentNegative unmarshals an instance of GetExplanationTaskResponseEntityExplanationsItemPertinentNegative from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityExplanationsItemPertinentNegative(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityExplanationsItemPertinentNegative)
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalExplanationFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prediction", &obj.Prediction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "probability", &obj.Probability)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityExplanationsItemPertinentPositive : These factors are sufficient evidence in themselves to yeild the given classification.
type GetExplanationTaskResponseEntityExplanationsItemPertinentPositive struct {
	// List of feature names, values and their importance.
	Features []ExplanationFeature `json:"features,omitempty"`
}

// UnmarshalGetExplanationTaskResponseEntityExplanationsItemPertinentPositive unmarshals an instance of GetExplanationTaskResponseEntityExplanationsItemPertinentPositive from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityExplanationsItemPertinentPositive(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityExplanationsItemPertinentPositive)
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalExplanationFeature)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityStatus : Status of the explanation task.
type GetExplanationTaskResponseEntityStatus struct {
	// Overall status of the explanation task.
	State *string `json:"state,omitempty"`
}

// Constants associated with the GetExplanationTaskResponseEntityStatus.State property.
// Overall status of the explanation task.
const (
	GetExplanationTaskResponseEntityStatus_State_Finished   = "finished"
	GetExplanationTaskResponseEntityStatus_State_InProgress = "in_progress"
)

// UnmarshalGetExplanationTaskResponseEntityStatus unmarshals an instance of GetExplanationTaskResponseEntityStatus from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityStatus)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseMetadata : Metadata of get explanation task response.
type GetExplanationTaskResponseMetadata struct {
	// Time when the explanation task was initiated.
	CreatedAt *string `json:"created_at" validate:"required"`

	// ID of the user creating explanation task.
	CreatedBy *string `json:"created_by" validate:"required"`

	// Identifier for tracking explanation task.
	ExplanationTaskID *string `json:"explanation_task_id" validate:"required"`

	// Time when the explanation task was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// UnmarshalGetExplanationTaskResponseMetadata unmarshals an instance of GetExplanationTaskResponseMetadata from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseMetadata)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "explanation_task_id", &obj.ExplanationTaskID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTasksResponse : List all explanations response.
type GetExplanationTasksResponse struct {
	// The list of explanation fields.
	ExplanationFields []string `json:"explanation_fields" validate:"required"`

	// The list of explanation values.
	ExplanationValues []string `json:"explanation_values" validate:"required"`

	// Maximum number of returned explanations.
	Limit *int64 `json:"limit" validate:"required"`

	// Offset of returned explanations.
	Offset *int64 `json:"offset" validate:"required"`

	// Total number of computed explanations.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalGetExplanationTasksResponse unmarshals an instance of GetExplanationTasksResponse from the specified map of raw messages.
func UnmarshalGetExplanationTasksResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTasksResponse)
	err = core.UnmarshalPrimitive(m, "explanation_fields", &obj.ExplanationFields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "explanation_values", &obj.ExplanationValues)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InputDataReference : InputDataReference is the same as TrainingDataReference except that neither location nor connection is required. This
// is needed for the Schemas API and to avoid updating existing APIs.
type InputDataReference struct {
	// training data set connection credentials.
	Connection TrainingDataReferenceConnectionIntf `json:"connection,omitempty"`

	// training data set location.
	Location TrainingDataReferenceLocationIntf `json:"location,omitempty"`

	Name *string `json:"name,omitempty"`

	// Type of the storage.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the InputDataReference.Type property.
// Type of the storage.
const (
	InputDataReference_Type_Cos       = "cos"
	InputDataReference_Type_Dataset   = "dataset"
	InputDataReference_Type_Db2       = "db2"
	InputDataReference_Type_FileAsset = "file_asset"
)

// NewInputDataReference : Instantiate InputDataReference (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewInputDataReference(typeVar string) (_model *InputDataReference, err error) {
	_model = &InputDataReference{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalInputDataReference unmarshals an instance of InputDataReference from the specified map of raw messages.
func UnmarshalInputDataReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InputDataReference)
	err = core.UnmarshalModel(m, "connection", &obj.Connection, UnmarshalTrainingDataReferenceConnection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTrainingDataReferenceLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstancesAddOptions : The InstancesAdd options.
type InstancesAddOptions struct {
	DataMartID *string `json:"data_mart_id" validate:"required"`

	MonitorDefinitionID *string `json:"monitor_definition_id" validate:"required"`

	Target *Target `json:"target" validate:"required"`

	ManagedBy *string `json:"managed_by,omitempty"`

	// Monitoring parameters consistent with the `parameters_schema` from the monitor definition.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
	// Defaults to once every hour if not specified.
	Schedule *MonitorInstanceSchedule `json:"schedule,omitempty"`

	ScheduleID *string `json:"schedule_id,omitempty"`

	// A set of schedules used to control how frequently the target is monitored for online and batch deployment type.
	Schedules *MonitorInstanceScheduleCollection `json:"schedules,omitempty"`

	Thresholds []MetricThresholdOverride `json:"thresholds,omitempty"`

	// Summary about records count.
	TotalRecords *RecordsCountSummary `json:"total_records,omitempty"`

	// Summary about records count.
	UnprocessedRecords *RecordsCountSummary `json:"unprocessed_records,omitempty"`

	// prevent schedule creation for this monitor instnace.
	SkipScheduler *bool `json:"skip_scheduler,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstancesAddOptions : Instantiate InstancesAddOptions
func (*WatsonOpenScaleV2) NewInstancesAddOptions(dataMartID string, monitorDefinitionID string, target *Target) *InstancesAddOptions {
	return &InstancesAddOptions{
		DataMartID:          core.StringPtr(dataMartID),
		MonitorDefinitionID: core.StringPtr(monitorDefinitionID),
		Target:              target,
	}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *InstancesAddOptions) SetDataMartID(dataMartID string) *InstancesAddOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetMonitorDefinitionID : Allow user to set MonitorDefinitionID
func (_options *InstancesAddOptions) SetMonitorDefinitionID(monitorDefinitionID string) *InstancesAddOptions {
	_options.MonitorDefinitionID = core.StringPtr(monitorDefinitionID)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *InstancesAddOptions) SetTarget(target *Target) *InstancesAddOptions {
	_options.Target = target
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *InstancesAddOptions) SetManagedBy(managedBy string) *InstancesAddOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetParameters : Allow user to set Parameters
func (_options *InstancesAddOptions) SetParameters(parameters map[string]interface{}) *InstancesAddOptions {
	_options.Parameters = parameters
	return _options
}

// SetSchedule : Allow user to set Schedule
func (_options *InstancesAddOptions) SetSchedule(schedule *MonitorInstanceSchedule) *InstancesAddOptions {
	_options.Schedule = schedule
	return _options
}

// SetScheduleID : Allow user to set ScheduleID
func (_options *InstancesAddOptions) SetScheduleID(scheduleID string) *InstancesAddOptions {
	_options.ScheduleID = core.StringPtr(scheduleID)
	return _options
}

// SetSchedules : Allow user to set Schedules
func (_options *InstancesAddOptions) SetSchedules(schedules *MonitorInstanceScheduleCollection) *InstancesAddOptions {
	_options.Schedules = schedules
	return _options
}

// SetThresholds : Allow user to set Thresholds
func (_options *InstancesAddOptions) SetThresholds(thresholds []MetricThresholdOverride) *InstancesAddOptions {
	_options.Thresholds = thresholds
	return _options
}

// SetTotalRecords : Allow user to set TotalRecords
func (_options *InstancesAddOptions) SetTotalRecords(totalRecords *RecordsCountSummary) *InstancesAddOptions {
	_options.TotalRecords = totalRecords
	return _options
}

// SetUnprocessedRecords : Allow user to set UnprocessedRecords
func (_options *InstancesAddOptions) SetUnprocessedRecords(unprocessedRecords *RecordsCountSummary) *InstancesAddOptions {
	_options.UnprocessedRecords = unprocessedRecords
	return _options
}

// SetSkipScheduler : Allow user to set SkipScheduler
func (_options *InstancesAddOptions) SetSkipScheduler(skipScheduler bool) *InstancesAddOptions {
	_options.SkipScheduler = core.BoolPtr(skipScheduler)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstancesAddOptions) SetHeaders(param map[string]string) *InstancesAddOptions {
	options.Headers = param
	return options
}

// InstancesDeleteOptions : The InstancesDelete options.
type InstancesDeleteOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstancesDeleteOptions : Instantiate InstancesDeleteOptions
func (*WatsonOpenScaleV2) NewInstancesDeleteOptions(monitorInstanceID string) *InstancesDeleteOptions {
	return &InstancesDeleteOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *InstancesDeleteOptions) SetMonitorInstanceID(monitorInstanceID string) *InstancesDeleteOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstancesDeleteOptions) SetHeaders(param map[string]string) *InstancesDeleteOptions {
	options.Headers = param
	return options
}

// InstancesGetOptions : The InstancesGet options.
type InstancesGetOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// comma-separeted list of fields (supported fields are unprocessed_records and total_records).
	Expand *string `json:"expand,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstancesGetOptions : Instantiate InstancesGetOptions
func (*WatsonOpenScaleV2) NewInstancesGetOptions(monitorInstanceID string) *InstancesGetOptions {
	return &InstancesGetOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *InstancesGetOptions) SetMonitorInstanceID(monitorInstanceID string) *InstancesGetOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetExpand : Allow user to set Expand
func (_options *InstancesGetOptions) SetExpand(expand string) *InstancesGetOptions {
	_options.Expand = core.StringPtr(expand)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstancesGetOptions) SetHeaders(param map[string]string) *InstancesGetOptions {
	options.Headers = param
	return options
}

// InstancesListOptions : The InstancesList options.
type InstancesListOptions struct {
	// comma-separeted list of IDs.
	DataMartID *string `json:"data_mart_id,omitempty"`

	// comma-separeted list of IDs.
	MonitorDefinitionID *string `json:"monitor_definition_id,omitempty"`

	// comma-separeted list of IDs.
	TargetTargetID *string `json:"target.target_id,omitempty"`

	// comma-separeted list of types.
	TargetTargetType *string `json:"target.target_type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstancesListOptions : Instantiate InstancesListOptions
func (*WatsonOpenScaleV2) NewInstancesListOptions() *InstancesListOptions {
	return &InstancesListOptions{}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *InstancesListOptions) SetDataMartID(dataMartID string) *InstancesListOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetMonitorDefinitionID : Allow user to set MonitorDefinitionID
func (_options *InstancesListOptions) SetMonitorDefinitionID(monitorDefinitionID string) *InstancesListOptions {
	_options.MonitorDefinitionID = core.StringPtr(monitorDefinitionID)
	return _options
}

// SetTargetTargetID : Allow user to set TargetTargetID
func (_options *InstancesListOptions) SetTargetTargetID(targetTargetID string) *InstancesListOptions {
	_options.TargetTargetID = core.StringPtr(targetTargetID)
	return _options
}

// SetTargetTargetType : Allow user to set TargetTargetType
func (_options *InstancesListOptions) SetTargetTargetType(targetTargetType string) *InstancesListOptions {
	_options.TargetTargetType = core.StringPtr(targetTargetType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstancesListOptions) SetHeaders(param map[string]string) *InstancesListOptions {
	options.Headers = param
	return options
}

// InstancesUpdateOptions : The InstancesUpdate options.
type InstancesUpdateOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	PatchDocument []PatchDocument `json:"PatchDocument" validate:"required"`

	// Flag that allows to control if the underlaying actions related to the monitor reconfiguration should be triggered.
	UpdateMetadataOnly *bool `json:"update_metadata_only,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstancesUpdateOptions : Instantiate InstancesUpdateOptions
func (*WatsonOpenScaleV2) NewInstancesUpdateOptions(monitorInstanceID string, patchDocument []PatchDocument) *InstancesUpdateOptions {
	return &InstancesUpdateOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
		PatchDocument:     patchDocument,
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *InstancesUpdateOptions) SetMonitorInstanceID(monitorInstanceID string) *InstancesUpdateOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetPatchDocument : Allow user to set PatchDocument
func (_options *InstancesUpdateOptions) SetPatchDocument(patchDocument []PatchDocument) *InstancesUpdateOptions {
	_options.PatchDocument = patchDocument
	return _options
}

// SetUpdateMetadataOnly : Allow user to set UpdateMetadataOnly
func (_options *InstancesUpdateOptions) SetUpdateMetadataOnly(updateMetadataOnly bool) *InstancesUpdateOptions {
	_options.UpdateMetadataOnly = core.BoolPtr(updateMetadataOnly)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstancesUpdateOptions) SetHeaders(param map[string]string) *InstancesUpdateOptions {
	options.Headers = param
	return options
}

// IntegratedSystem : Integrated System definition.
type IntegratedSystem struct {
	// The additional connection information for the Integrated System.
	Connection interface{} `json:"connection,omitempty"`

	// The credentials for the Integrated System.
	Credentials map[string]interface{} `json:"credentials" validate:"required"`

	// The description of the Integrated System.
	Description *string `json:"description" validate:"required"`

	// The name of the Integrated System.
	Name *string `json:"name" validate:"required"`

	Type *string `json:"type" validate:"required"`
}

// Constants associated with the IntegratedSystem.Type property.
const (
	IntegratedSystem_Type_Hive              = "hive"
	IntegratedSystem_Type_OpenPages         = "open_pages"
	IntegratedSystem_Type_Slack             = "slack"
	IntegratedSystem_Type_Spark             = "spark"
	IntegratedSystem_Type_WatsonDataCatalog = "watson_data_catalog"
)

// NewIntegratedSystem : Instantiate IntegratedSystem (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewIntegratedSystem(credentials map[string]interface{}, description string, name string, typeVar string) (_model *IntegratedSystem, err error) {
	_model = &IntegratedSystem{
		Credentials: credentials,
		Description: core.StringPtr(description),
		Name:        core.StringPtr(name),
		Type:        core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalIntegratedSystem unmarshals an instance of IntegratedSystem from the specified map of raw messages.
func UnmarshalIntegratedSystem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntegratedSystem)
	err = core.UnmarshalPrimitive(m, "connection", &obj.Connection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credentials", &obj.Credentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IntegratedSystemCollection : IntegratedSystemCollection struct
type IntegratedSystemCollection struct {
	IntegratedSystems []IntegratedSystemResponse `json:"integrated_systems" validate:"required"`
}

// UnmarshalIntegratedSystemCollection unmarshals an instance of IntegratedSystemCollection from the specified map of raw messages.
func UnmarshalIntegratedSystemCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntegratedSystemCollection)
	err = core.UnmarshalModel(m, "integrated_systems", &obj.IntegratedSystems, UnmarshalIntegratedSystemResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IntegratedSystemReference : Integrated System reference.
type IntegratedSystemReference struct {
	// id of the resource in the Integrated System.
	ExternalID *string `json:"external_id" validate:"required"`

	// id of the Integrated System.
	IntegratedSystemID *string `json:"integrated_system_id" validate:"required"`
}

// UnmarshalIntegratedSystemReference unmarshals an instance of IntegratedSystemReference from the specified map of raw messages.
func UnmarshalIntegratedSystemReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntegratedSystemReference)
	err = core.UnmarshalPrimitive(m, "external_id", &obj.ExternalID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "integrated_system_id", &obj.IntegratedSystemID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IntegratedSystemResponse : IntegratedSystemResponse struct
type IntegratedSystemResponse struct {
	// Integrated System definition.
	Entity *IntegratedSystem `json:"entity,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`
}

// UnmarshalIntegratedSystemResponse unmarshals an instance of IntegratedSystemResponse from the specified map of raw messages.
func UnmarshalIntegratedSystemResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntegratedSystemResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalIntegratedSystem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonOpenScaleV2) NewIntegratedSystemResponsePatch(integratedSystemResponse *IntegratedSystemResponse) (_patch []JSONPatchOperation) {
	if integratedSystemResponse.Entity != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/entity"),
			Value: integratedSystemResponse.Entity,
		})
	}
	if integratedSystemResponse.Metadata != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/metadata"),
			Value: integratedSystemResponse.Metadata,
		})
	}
	return
}

// IntegratedSystemsAddOptions : The IntegratedSystemsAdd options.
type IntegratedSystemsAddOptions struct {
	// The credentials for the Integrated System.
	Credentials map[string]interface{} `json:"credentials" validate:"required"`

	// The description of the Integrated System.
	Description *string `json:"description" validate:"required"`

	// The name of the Integrated System.
	Name *string `json:"name" validate:"required"`

	Type *string `json:"type" validate:"required"`

	// The additional connection information for the Integrated System.
	Connection interface{} `json:"connection,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the IntegratedSystemsAddOptions.Type property.
const (
	IntegratedSystemsAddOptions_Type_Hive              = "hive"
	IntegratedSystemsAddOptions_Type_OpenPages         = "open_pages"
	IntegratedSystemsAddOptions_Type_Slack             = "slack"
	IntegratedSystemsAddOptions_Type_Spark             = "spark"
	IntegratedSystemsAddOptions_Type_WatsonDataCatalog = "watson_data_catalog"
)

// NewIntegratedSystemsAddOptions : Instantiate IntegratedSystemsAddOptions
func (*WatsonOpenScaleV2) NewIntegratedSystemsAddOptions(credentials map[string]interface{}, description string, name string, typeVar string) *IntegratedSystemsAddOptions {
	return &IntegratedSystemsAddOptions{
		Credentials: credentials,
		Description: core.StringPtr(description),
		Name:        core.StringPtr(name),
		Type:        core.StringPtr(typeVar),
	}
}

// SetCredentials : Allow user to set Credentials
func (_options *IntegratedSystemsAddOptions) SetCredentials(credentials map[string]interface{}) *IntegratedSystemsAddOptions {
	_options.Credentials = credentials
	return _options
}

// SetDescription : Allow user to set Description
func (_options *IntegratedSystemsAddOptions) SetDescription(description string) *IntegratedSystemsAddOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetName : Allow user to set Name
func (_options *IntegratedSystemsAddOptions) SetName(name string) *IntegratedSystemsAddOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetType : Allow user to set Type
func (_options *IntegratedSystemsAddOptions) SetType(typeVar string) *IntegratedSystemsAddOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetConnection : Allow user to set Connection
func (_options *IntegratedSystemsAddOptions) SetConnection(connection interface{}) *IntegratedSystemsAddOptions {
	_options.Connection = connection
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *IntegratedSystemsAddOptions) SetHeaders(param map[string]string) *IntegratedSystemsAddOptions {
	options.Headers = param
	return options
}

// IntegratedSystemsDeleteOptions : The IntegratedSystemsDelete options.
type IntegratedSystemsDeleteOptions struct {
	// Unique integrated system ID.
	IntegratedSystemID *string `json:"integrated_system_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewIntegratedSystemsDeleteOptions : Instantiate IntegratedSystemsDeleteOptions
func (*WatsonOpenScaleV2) NewIntegratedSystemsDeleteOptions(integratedSystemID string) *IntegratedSystemsDeleteOptions {
	return &IntegratedSystemsDeleteOptions{
		IntegratedSystemID: core.StringPtr(integratedSystemID),
	}
}

// SetIntegratedSystemID : Allow user to set IntegratedSystemID
func (_options *IntegratedSystemsDeleteOptions) SetIntegratedSystemID(integratedSystemID string) *IntegratedSystemsDeleteOptions {
	_options.IntegratedSystemID = core.StringPtr(integratedSystemID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *IntegratedSystemsDeleteOptions) SetHeaders(param map[string]string) *IntegratedSystemsDeleteOptions {
	options.Headers = param
	return options
}

// IntegratedSystemsGetOptions : The IntegratedSystemsGet options.
type IntegratedSystemsGetOptions struct {
	// Unique integrated system ID.
	IntegratedSystemID *string `json:"integrated_system_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewIntegratedSystemsGetOptions : Instantiate IntegratedSystemsGetOptions
func (*WatsonOpenScaleV2) NewIntegratedSystemsGetOptions(integratedSystemID string) *IntegratedSystemsGetOptions {
	return &IntegratedSystemsGetOptions{
		IntegratedSystemID: core.StringPtr(integratedSystemID),
	}
}

// SetIntegratedSystemID : Allow user to set IntegratedSystemID
func (_options *IntegratedSystemsGetOptions) SetIntegratedSystemID(integratedSystemID string) *IntegratedSystemsGetOptions {
	_options.IntegratedSystemID = core.StringPtr(integratedSystemID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *IntegratedSystemsGetOptions) SetHeaders(param map[string]string) *IntegratedSystemsGetOptions {
	options.Headers = param
	return options
}

// IntegratedSystemsListOptions : The IntegratedSystemsList options.
type IntegratedSystemsListOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewIntegratedSystemsListOptions : Instantiate IntegratedSystemsListOptions
func (*WatsonOpenScaleV2) NewIntegratedSystemsListOptions() *IntegratedSystemsListOptions {
	return &IntegratedSystemsListOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *IntegratedSystemsListOptions) SetHeaders(param map[string]string) *IntegratedSystemsListOptions {
	options.Headers = param
	return options
}

// IntegratedSystemsUpdateOptions : The IntegratedSystemsUpdate options.
type IntegratedSystemsUpdateOptions struct {
	// Unique integrated system ID.
	IntegratedSystemID *string `json:"integrated_system_id" validate:"required,ne="`

	// Array of patch operations as defined in RFC 6902.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewIntegratedSystemsUpdateOptions : Instantiate IntegratedSystemsUpdateOptions
func (*WatsonOpenScaleV2) NewIntegratedSystemsUpdateOptions(integratedSystemID string, jsonPatchOperation []JSONPatchOperation) *IntegratedSystemsUpdateOptions {
	return &IntegratedSystemsUpdateOptions{
		IntegratedSystemID: core.StringPtr(integratedSystemID),
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetIntegratedSystemID : Allow user to set IntegratedSystemID
func (_options *IntegratedSystemsUpdateOptions) SetIntegratedSystemID(integratedSystemID string) *IntegratedSystemsUpdateOptions {
	_options.IntegratedSystemID = core.StringPtr(integratedSystemID)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *IntegratedSystemsUpdateOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *IntegratedSystemsUpdateOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *IntegratedSystemsUpdateOptions) SetHeaders(param map[string]string) *IntegratedSystemsUpdateOptions {
	options.Headers = param
	return options
}

// JsDictElem : Fields and values of the entity matches JSON object's fields and values.
type JsDictElem struct {
}

// UnmarshalJsDictElem unmarshals an instance of JsDictElem from the specified map of raw messages.
func UnmarshalJsDictElem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JsDictElem)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on a JSON document, as defined by RFC 6902.
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The JSON Pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The JSON Pointer that identifies the field that is the source of the operation.
	From *string `json:"from,omitempty"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add     = "add"
	JSONPatchOperation_Op_Copy    = "copy"
	JSONPatchOperation_Op_Move    = "move"
	JSONPatchOperation_Op_Remove  = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test    = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op:   core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LimeExplanationPrediction : Lime explanation prediction details.
type LimeExplanationPrediction struct {
	// List of features and their contribution in prediction.
	ExplanationFeatures []ExplanationFeature `json:"explanation_features,omitempty"`

	// Signifies probability of this particular prediction.
	Probability *float64 `json:"probability,omitempty"`

	// Value of the output field in this particular prediction.
	Value *string `json:"value,omitempty"`
}

// UnmarshalLimeExplanationPrediction unmarshals an instance of LimeExplanationPrediction from the specified map of raw messages.
func UnmarshalLimeExplanationPrediction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LimeExplanationPrediction)
	err = core.UnmarshalModel(m, "explanation_features", &obj.ExplanationFeatures, UnmarshalExplanationFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "probability", &obj.Probability)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LocationSchemaName : LocationSchemaName struct
type LocationSchemaName struct {
	// Database schema name (for PostgreSQL default is a public schema).
	SchemaName *string `json:"schema_name,omitempty"`
}

// UnmarshalLocationSchemaName unmarshals an instance of LocationSchemaName from the specified map of raw messages.
func UnmarshalLocationSchemaName(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LocationSchemaName)
	err = core.UnmarshalPrimitive(m, "schema_name", &obj.SchemaName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LocationTableName : LocationTableName struct
type LocationTableName struct {
	// Database table name.
	TableName *string `json:"table_name,omitempty"`
}

// UnmarshalLocationTableName unmarshals an instance of LocationTableName from the specified map of raw messages.
func UnmarshalLocationTableName(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LocationTableName)
	err = core.UnmarshalPrimitive(m, "table_name", &obj.TableName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MLCredentials : MLCredentials struct
// Models which "extend" this model:
// - SageMakerCredentials
// - AzureCredentials
// - CustomCredentials
// - WMLCredentialsCloud
// - WMLCredentialsCP4D
// - SPSSCredentials
// - UnknownCredentials
type MLCredentials struct {
	AccessKeyID *string `json:"access_key_id,omitempty"`

	Region *string `json:"region,omitempty"`

	SecretAccessKey *string `json:"secret_access_key,omitempty"`

	ClientID *string `json:"client_id,omitempty"`

	ClientSecret *string `json:"client_secret,omitempty"`

	Password *string `json:"password,omitempty"`

	SubscriptionID *string `json:"subscription_id,omitempty"`

	Tenant *string `json:"tenant,omitempty"`

	Token *string `json:"token,omitempty"`

	Username *string `json:"username,omitempty"`

	Workspaces []AzureWorkspaceCredentials `json:"workspaces,omitempty"`

	URL *string `json:"url,omitempty"`

	Apikey *string `json:"apikey,omitempty"`

	InstanceID *string `json:"instance_id,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

func (*MLCredentials) isaMLCredentials() bool {
	return true
}

type MLCredentialsIntf interface {
	isaMLCredentials() bool
	SetProperty(key string, value interface{})
	SetProperties(m map[string]interface{})
	GetProperty(key string) interface{}
	GetProperties() map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of MLCredentials
func (o *MLCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of MLCredentials
func (o *MLCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of MLCredentials
func (o *MLCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of MLCredentials
func (o *MLCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of MLCredentials
func (o *MLCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.AccessKeyID != nil {
		m["access_key_id"] = o.AccessKeyID
	}
	if o.Region != nil {
		m["region"] = o.Region
	}
	if o.SecretAccessKey != nil {
		m["secret_access_key"] = o.SecretAccessKey
	}
	if o.ClientID != nil {
		m["client_id"] = o.ClientID
	}
	if o.ClientSecret != nil {
		m["client_secret"] = o.ClientSecret
	}
	if o.Password != nil {
		m["password"] = o.Password
	}
	if o.SubscriptionID != nil {
		m["subscription_id"] = o.SubscriptionID
	}
	if o.Tenant != nil {
		m["tenant"] = o.Tenant
	}
	if o.Token != nil {
		m["token"] = o.Token
	}
	if o.Username != nil {
		m["username"] = o.Username
	}
	if o.Workspaces != nil {
		m["workspaces"] = o.Workspaces
	}
	if o.URL != nil {
		m["url"] = o.URL
	}
	if o.Apikey != nil {
		m["apikey"] = o.Apikey
	}
	if o.InstanceID != nil {
		m["instance_id"] = o.InstanceID
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalMLCredentials unmarshals an instance of MLCredentials from the specified map of raw messages.
func UnmarshalMLCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MLCredentials)
	err = core.UnmarshalPrimitive(m, "access_key_id", &obj.AccessKeyID)
	if err != nil {
		return
	}
	delete(m, "access_key_id")
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	delete(m, "region")
	err = core.UnmarshalPrimitive(m, "secret_access_key", &obj.SecretAccessKey)
	if err != nil {
		return
	}
	delete(m, "secret_access_key")
	err = core.UnmarshalPrimitive(m, "client_id", &obj.ClientID)
	if err != nil {
		return
	}
	delete(m, "client_id")
	err = core.UnmarshalPrimitive(m, "client_secret", &obj.ClientSecret)
	if err != nil {
		return
	}
	delete(m, "client_secret")
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	delete(m, "password")
	err = core.UnmarshalPrimitive(m, "subscription_id", &obj.SubscriptionID)
	if err != nil {
		return
	}
	delete(m, "subscription_id")
	err = core.UnmarshalPrimitive(m, "tenant", &obj.Tenant)
	if err != nil {
		return
	}
	delete(m, "tenant")
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	delete(m, "token")
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	delete(m, "username")
	err = core.UnmarshalModel(m, "workspaces", &obj.Workspaces, UnmarshalAzureWorkspaceCredentials)
	if err != nil {
		return
	}
	delete(m, "workspaces")
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	delete(m, "url")
	err = core.UnmarshalPrimitive(m, "apikey", &obj.Apikey)
	if err != nil {
		return
	}
	delete(m, "apikey")
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	delete(m, "instance_id")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MeasurementEntity : MeasurementEntity struct
type MeasurementEntity struct {
	// Revision number of the ML model or function used by the monitor.
	AssetRevision *string `json:"asset_revision,omitempty"`

	// Number of the metrics with issues, which exceeded limits.
	IssueCount *int64 `json:"issue_count" validate:"required"`

	MonitorDefinitionID *string `json:"monitor_definition_id,omitempty"`

	MonitorInstanceID *string `json:"monitor_instance_id,omitempty"`

	// ID of the monitoring run which produced the measurement.
	RunID *string `json:"run_id,omitempty"`

	Target *Target `json:"target,omitempty"`

	Timestamp *strfmt.DateTime `json:"timestamp" validate:"required"`

	// Metrics grouped for a single measurement.
	Values []MonitorMeasurementValue `json:"values" validate:"required"`
}

// UnmarshalMeasurementEntity unmarshals an instance of MeasurementEntity from the specified map of raw messages.
func UnmarshalMeasurementEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MeasurementEntity)
	err = core.UnmarshalPrimitive(m, "asset_revision", &obj.AssetRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "issue_count", &obj.IssueCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor_definition_id", &obj.MonitorDefinitionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor_instance_id", &obj.MonitorInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_id", &obj.RunID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalMonitorMeasurementValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MeasurementsAddOptions : The MeasurementsAdd options.
type MeasurementsAddOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	MonitorMeasurementRequest []MonitorMeasurementRequest `json:"MonitorMeasurementRequest" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMeasurementsAddOptions : Instantiate MeasurementsAddOptions
func (*WatsonOpenScaleV2) NewMeasurementsAddOptions(monitorInstanceID string, monitorMeasurementRequest []MonitorMeasurementRequest) *MeasurementsAddOptions {
	return &MeasurementsAddOptions{
		MonitorInstanceID:         core.StringPtr(monitorInstanceID),
		MonitorMeasurementRequest: monitorMeasurementRequest,
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *MeasurementsAddOptions) SetMonitorInstanceID(monitorInstanceID string) *MeasurementsAddOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetMonitorMeasurementRequest : Allow user to set MonitorMeasurementRequest
func (_options *MeasurementsAddOptions) SetMonitorMeasurementRequest(monitorMeasurementRequest []MonitorMeasurementRequest) *MeasurementsAddOptions {
	_options.MonitorMeasurementRequest = monitorMeasurementRequest
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MeasurementsAddOptions) SetHeaders(param map[string]string) *MeasurementsAddOptions {
	options.Headers = param
	return options
}

// MeasurementsDeleteOptions : The MeasurementsDelete options.
type MeasurementsDeleteOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMeasurementsDeleteOptions : Instantiate MeasurementsDeleteOptions
func (*WatsonOpenScaleV2) NewMeasurementsDeleteOptions(monitorInstanceID string) *MeasurementsDeleteOptions {
	return &MeasurementsDeleteOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *MeasurementsDeleteOptions) SetMonitorInstanceID(monitorInstanceID string) *MeasurementsDeleteOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MeasurementsDeleteOptions) SetHeaders(param map[string]string) *MeasurementsDeleteOptions {
	options.Headers = param
	return options
}

// MeasurementsGetOptions : The MeasurementsGet options.
type MeasurementsGetOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Uniq monitoring run ID.
	MeasurementID *string `json:"measurement_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMeasurementsGetOptions : Instantiate MeasurementsGetOptions
func (*WatsonOpenScaleV2) NewMeasurementsGetOptions(monitorInstanceID string, measurementID string) *MeasurementsGetOptions {
	return &MeasurementsGetOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
		MeasurementID:     core.StringPtr(measurementID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *MeasurementsGetOptions) SetMonitorInstanceID(monitorInstanceID string) *MeasurementsGetOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetMeasurementID : Allow user to set MeasurementID
func (_options *MeasurementsGetOptions) SetMeasurementID(measurementID string) *MeasurementsGetOptions {
	_options.MeasurementID = core.StringPtr(measurementID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MeasurementsGetOptions) SetHeaders(param map[string]string) *MeasurementsGetOptions {
	options.Headers = param
	return options
}

// MeasurementsListOptions : The MeasurementsList options.
type MeasurementsListOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Beginning of the time range.
	Start *strfmt.DateTime `json:"start" validate:"required"`

	// End of the time range.
	End *strfmt.DateTime `json:"end" validate:"required"`

	// Comma delimited list of measurement run_id.
	RunID *string `json:"run_id" validate:"required"`

	// Filter expression can consist of any metric tag or a common column of string type followed by filter name and
	// optionally a value, all delimited by colon. Supported filters are: `in`, `eq`, `null` and `exists`. Sample filters
	// are: `filter=region:in:[us,pl],segment:eq:sales` or `filter=region:null,segment:exists`.
	Filter *string `json:"filter,omitempty"`

	// Maximum number of elements returned.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMeasurementsListOptions : Instantiate MeasurementsListOptions
func (*WatsonOpenScaleV2) NewMeasurementsListOptions(monitorInstanceID string, start *strfmt.DateTime, end *strfmt.DateTime, runID string) *MeasurementsListOptions {
	return &MeasurementsListOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
		Start:             start,
		End:               end,
		RunID:             core.StringPtr(runID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *MeasurementsListOptions) SetMonitorInstanceID(monitorInstanceID string) *MeasurementsListOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *MeasurementsListOptions) SetStart(start *strfmt.DateTime) *MeasurementsListOptions {
	_options.Start = start
	return _options
}

// SetEnd : Allow user to set End
func (_options *MeasurementsListOptions) SetEnd(end *strfmt.DateTime) *MeasurementsListOptions {
	_options.End = end
	return _options
}

// SetRunID : Allow user to set RunID
func (_options *MeasurementsListOptions) SetRunID(runID string) *MeasurementsListOptions {
	_options.RunID = core.StringPtr(runID)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *MeasurementsListOptions) SetFilter(filter string) *MeasurementsListOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *MeasurementsListOptions) SetLimit(limit int64) *MeasurementsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MeasurementsListOptions) SetHeaders(param map[string]string) *MeasurementsListOptions {
	options.Headers = param
	return options
}

// MeasurementsQueryOptions : The MeasurementsQuery options.
type MeasurementsQueryOptions struct {
	// Comma separated ID of the monitoring target (subscription or business application).
	TargetID *string `json:"target_id,omitempty"`

	// Type of the monitoring target (subscription or business application).
	TargetType *string `json:"target_type,omitempty"`

	// Comma separated ID of the monitor definition.
	MonitorDefinitionID *string `json:"monitor_definition_id,omitempty"`

	// Filter expression can consist of any metric tag or a common column of string type followed by filter name and
	// optionally a value, all delimited by colon and prepended with `monitor_definition_id.` string. Supported filters
	// are: `in`, `eq`, `null` and `exists`. Sample filters are:
	// `monitor_definition_id.filter=region:in:[us,pl],monitor_definition_id.segment:eq:sales` or
	// `filter=monitor_definition_id.region:null,monitor_definition_id.segment:exists`. Every monitor_definition_id can
	// have own set of filters.
	Filter *string `json:"filter,omitempty"`

	// Number of measurements (per target) to be returned.
	RecentCount *int64 `json:"recent_count,omitempty"`

	// Format of the returned data. `full` format compared to `compact` is additive and contains `sources` part.
	Format *string `json:"format,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the MeasurementsQueryOptions.TargetType property.
// Type of the monitoring target (subscription or business application).
const (
	MeasurementsQueryOptions_TargetType_BusinessApplication = "business_application"
	MeasurementsQueryOptions_TargetType_DataMart            = "data_mart"
	MeasurementsQueryOptions_TargetType_Instance            = "instance"
	MeasurementsQueryOptions_TargetType_Subscription        = "subscription"
)

// Constants associated with the MeasurementsQueryOptions.Format property.
// Format of the returned data. `full` format compared to `compact` is additive and contains `sources` part.
const (
	MeasurementsQueryOptions_Format_Compact = "compact"
	MeasurementsQueryOptions_Format_Full    = "full"
)

// NewMeasurementsQueryOptions : Instantiate MeasurementsQueryOptions
func (*WatsonOpenScaleV2) NewMeasurementsQueryOptions() *MeasurementsQueryOptions {
	return &MeasurementsQueryOptions{}
}

// SetTargetID : Allow user to set TargetID
func (_options *MeasurementsQueryOptions) SetTargetID(targetID string) *MeasurementsQueryOptions {
	_options.TargetID = core.StringPtr(targetID)
	return _options
}

// SetTargetType : Allow user to set TargetType
func (_options *MeasurementsQueryOptions) SetTargetType(targetType string) *MeasurementsQueryOptions {
	_options.TargetType = core.StringPtr(targetType)
	return _options
}

// SetMonitorDefinitionID : Allow user to set MonitorDefinitionID
func (_options *MeasurementsQueryOptions) SetMonitorDefinitionID(monitorDefinitionID string) *MeasurementsQueryOptions {
	_options.MonitorDefinitionID = core.StringPtr(monitorDefinitionID)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *MeasurementsQueryOptions) SetFilter(filter string) *MeasurementsQueryOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetRecentCount : Allow user to set RecentCount
func (_options *MeasurementsQueryOptions) SetRecentCount(recentCount int64) *MeasurementsQueryOptions {
	_options.RecentCount = core.Int64Ptr(recentCount)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *MeasurementsQueryOptions) SetFormat(format string) *MeasurementsQueryOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MeasurementsQueryOptions) SetHeaders(param map[string]string) *MeasurementsQueryOptions {
	options.Headers = param
	return options
}

// MeasurementsResponseCollection : MeasurementsResponseCollection struct
type MeasurementsResponseCollection struct {
	Measurements []MeasurementsResponseCollectionMeasurementsItem `json:"measurements" validate:"required"`
}

// UnmarshalMeasurementsResponseCollection unmarshals an instance of MeasurementsResponseCollection from the specified map of raw messages.
func UnmarshalMeasurementsResponseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MeasurementsResponseCollection)
	err = core.UnmarshalModel(m, "measurements", &obj.Measurements, UnmarshalMeasurementsResponseCollectionMeasurementsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MeasurementsResponseCollectionMeasurementsItem : MeasurementsResponseCollectionMeasurementsItem struct
type MeasurementsResponseCollectionMeasurementsItem struct {
	Entity *MeasurementEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalMeasurementsResponseCollectionMeasurementsItem unmarshals an instance of MeasurementsResponseCollectionMeasurementsItem from the specified map of raw messages.
func UnmarshalMeasurementsResponseCollectionMeasurementsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MeasurementsResponseCollectionMeasurementsItem)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalMeasurementEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Metadata : Metadata struct
type Metadata struct {
	// The timestamp when the resource was first created In format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ,
	// matching the date-time format as specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The IAM ID of the user who created the resource.
	CreatedBy *string `json:"created_by,omitempty"`

	// Cloud Resource Name (CRN) uniquely identify IBM Cloud resource (https://console.bluemix.net/docs/overview/crn.html).
	Crn *string `json:"crn,omitempty"`

	// The ID (typically a GUID) which uniquely identifies the resource.
	ID *string `json:"id" validate:"required"`

	// The timestamp when the resource was first created In format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ,
	// matching the date-time format as specified by RFC 3339.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The IAM ID of the user who last modified the resource.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// The URL which can be used to uniquely refer to the resource Typically a GET on this url would return details of the
	// resource, a DELETE would delete it and a PUT/PATCH would update it.
	URL *string `json:"url" validate:"required"`
}

// UnmarshalMetadata unmarshals an instance of Metadata from the specified map of raw messages.
func UnmarshalMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Metadata)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricSpecificThresholdObject : MetricSpecificThresholdObject struct
type MetricSpecificThresholdObject struct {
	AppliesTo []ThresholdConditionObject `json:"applies_to" validate:"required"`

	// default value of threshold.
	Default *float64 `json:"default" validate:"required"`

	ID *string `json:"id" validate:"required"`

	Recommendation *string `json:"recommendation,omitempty"`
}

// NewMetricSpecificThresholdObject : Instantiate MetricSpecificThresholdObject (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMetricSpecificThresholdObject(appliesTo []ThresholdConditionObject, defaultVar float64, id string) (_model *MetricSpecificThresholdObject, err error) {
	_model = &MetricSpecificThresholdObject{
		AppliesTo: appliesTo,
		Default:   core.Float64Ptr(defaultVar),
		ID:        core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMetricSpecificThresholdObject unmarshals an instance of MetricSpecificThresholdObject from the specified map of raw messages.
func UnmarshalMetricSpecificThresholdObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricSpecificThresholdObject)
	err = core.UnmarshalModel(m, "applies_to", &obj.AppliesTo, UnmarshalThresholdConditionObject)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "recommendation", &obj.Recommendation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricSpecificThresholdShortObject : MetricSpecificThresholdShortObject struct
type MetricSpecificThresholdShortObject struct {
	AppliesTo []ThresholdConditionObject `json:"applies_to" validate:"required"`

	// value of threshold.
	Value *float64 `json:"value" validate:"required"`
}

// NewMetricSpecificThresholdShortObject : Instantiate MetricSpecificThresholdShortObject (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMetricSpecificThresholdShortObject(appliesTo []ThresholdConditionObject, value float64) (_model *MetricSpecificThresholdShortObject, err error) {
	_model = &MetricSpecificThresholdShortObject{
		AppliesTo: appliesTo,
		Value:     core.Float64Ptr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMetricSpecificThresholdShortObject unmarshals an instance of MetricSpecificThresholdShortObject from the specified map of raw messages.
func UnmarshalMetricSpecificThresholdShortObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricSpecificThresholdShortObject)
	err = core.UnmarshalModel(m, "applies_to", &obj.AppliesTo, UnmarshalThresholdConditionObject)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricThreshold : MetricThreshold struct
type MetricThreshold struct {
	// default value of threshold.
	Default *float64 `json:"default,omitempty"`

	DefaultRecommendation *string `json:"default_recommendation,omitempty"`

	SpecificValues []MetricSpecificThresholdObject `json:"specific_values,omitempty"`

	Type *string `json:"type" validate:"required"`
}

// Constants associated with the MetricThreshold.Type property.
const (
	MetricThreshold_Type_LowerLimit = "lower_limit"
	MetricThreshold_Type_UpperLimit = "upper_limit"
)

// NewMetricThreshold : Instantiate MetricThreshold (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMetricThreshold(typeVar string) (_model *MetricThreshold, err error) {
	_model = &MetricThreshold{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMetricThreshold unmarshals an instance of MetricThreshold from the specified map of raw messages.
func UnmarshalMetricThreshold(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricThreshold)
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_recommendation", &obj.DefaultRecommendation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "specific_values", &obj.SpecificValues, UnmarshalMetricSpecificThresholdObject)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricThresholdOverride : MetricThresholdOverride struct
type MetricThresholdOverride struct {
	MetricID *string `json:"metric_id" validate:"required"`

	SpecificValues []MetricSpecificThresholdShortObject `json:"specific_values,omitempty"`

	Type *string `json:"type" validate:"required"`

	// value of the threshold.
	Value *float64 `json:"value,omitempty"`
}

// Constants associated with the MetricThresholdOverride.Type property.
const (
	MetricThresholdOverride_Type_LowerLimit = "lower_limit"
	MetricThresholdOverride_Type_UpperLimit = "upper_limit"
)

// NewMetricThresholdOverride : Instantiate MetricThresholdOverride (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMetricThresholdOverride(metricID string, typeVar string) (_model *MetricThresholdOverride, err error) {
	_model = &MetricThresholdOverride{
		MetricID: core.StringPtr(metricID),
		Type:     core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMetricThresholdOverride unmarshals an instance of MetricThresholdOverride from the specified map of raw messages.
func UnmarshalMetricThresholdOverride(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricThresholdOverride)
	err = core.UnmarshalPrimitive(m, "metric_id", &obj.MetricID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "specific_values", &obj.SpecificValues, UnmarshalMetricSpecificThresholdShortObject)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricsListOptions : The MetricsList options.
type MetricsListOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Calculations **inclusive**, internally floored to achieve full interval. If interval is vulnerable to time zone, the
	// calculated value depends on a backend db engine: PostgreSQL respects time zone and DB2 use UTC time. Calculated
	// value is returned in response.
	Start *strfmt.DateTime `json:"start" validate:"required"`

	// Calculations **exclusive**, internally ceiled to achieve full interval. If interval is vulnerable to time zone, the
	// calculated value depends on a backend db engine: PostgreSQL respects time zone and DB2 use UTC time. Calculated
	// value is returned in response.
	End *strfmt.DateTime `json:"end" validate:"required"`

	// Comma delimited function list constructed from metric name and function, e.g. `agg=metric_name:count,:last` that
	// defines aggregations.
	Agg *string `json:"agg" validate:"required"`

	// Time unit in which metrics are grouped and aggregated, interval by interval.
	Interval *string `json:"interval,omitempty"`

	// Filter expression can consist of any metric tag or a common column of string type followed by filter name and
	// optionally a value, all delimited by colon. Supported filters are: `in`, `eq`, `null` and `exists`. Sample filters
	// are: `filter=region:in:[us,pl],segment:eq:sales` or `filter=region:null,segment:exists`.
	Filter *string `json:"filter,omitempty"`

	// Comma delimited list constructed from metric tags, e.g. `group=region,segment` to group metrics before aggregations.
	Group *string `json:"group,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the MetricsListOptions.Agg property.
// Comma delimited function list constructed from metric name and function, e.g. `agg=metric_name:count,:last` that
// defines aggregations.
const (
	MetricsListOptions_Agg_Avg    = "avg"
	MetricsListOptions_Agg_Count  = "count"
	MetricsListOptions_Agg_First  = "first"
	MetricsListOptions_Agg_Last   = "last"
	MetricsListOptions_Agg_Max    = "max"
	MetricsListOptions_Agg_Min    = "min"
	MetricsListOptions_Agg_Stddev = "stddev"
	MetricsListOptions_Agg_Sum    = "sum"
)

// Constants associated with the MetricsListOptions.Interval property.
// Time unit in which metrics are grouped and aggregated, interval by interval.
const (
	MetricsListOptions_Interval_Day    = "day"
	MetricsListOptions_Interval_Hour   = "hour"
	MetricsListOptions_Interval_Minute = "minute"
	MetricsListOptions_Interval_Month  = "month"
	MetricsListOptions_Interval_Week   = "week"
	MetricsListOptions_Interval_Year   = "year"
)

// NewMetricsListOptions : Instantiate MetricsListOptions
func (*WatsonOpenScaleV2) NewMetricsListOptions(monitorInstanceID string, start *strfmt.DateTime, end *strfmt.DateTime, agg string) *MetricsListOptions {
	return &MetricsListOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
		Start:             start,
		End:               end,
		Agg:               core.StringPtr(agg),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *MetricsListOptions) SetMonitorInstanceID(monitorInstanceID string) *MetricsListOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *MetricsListOptions) SetStart(start *strfmt.DateTime) *MetricsListOptions {
	_options.Start = start
	return _options
}

// SetEnd : Allow user to set End
func (_options *MetricsListOptions) SetEnd(end *strfmt.DateTime) *MetricsListOptions {
	_options.End = end
	return _options
}

// SetAgg : Allow user to set Agg
func (_options *MetricsListOptions) SetAgg(agg string) *MetricsListOptions {
	_options.Agg = core.StringPtr(agg)
	return _options
}

// SetInterval : Allow user to set Interval
func (_options *MetricsListOptions) SetInterval(interval string) *MetricsListOptions {
	_options.Interval = core.StringPtr(interval)
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *MetricsListOptions) SetFilter(filter string) *MetricsListOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetGroup : Allow user to set Group
func (_options *MetricsListOptions) SetGroup(group string) *MetricsListOptions {
	_options.Group = core.StringPtr(group)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MetricsListOptions) SetHeaders(param map[string]string) *MetricsListOptions {
	options.Headers = param
	return options
}

// Monitor : Monitor struct
type Monitor struct {
	AppliesTo *ApplicabilitySelection `json:"applies_to,omitempty"`

	// Long monitoring description presented in monitoring catalog.
	Description *string `json:"description,omitempty"`

	// translated resources.
	Dictionary map[string]interface{} `json:"dictionary,omitempty"`

	ManagedBy *string `json:"managed_by,omitempty"`

	// A list of metric definition.
	Metrics []MonitorMetric `json:"metrics" validate:"required"`

	MonitorRuntime *MonitorRuntime `json:"monitor_runtime,omitempty"`

	// Monitor UI label (must be unique).
	Name *string `json:"name" validate:"required"`

	// JSON schema that will be used to validate monitoring parameters when enabled.
	ParametersSchema map[string]interface{} `json:"parameters_schema,omitempty"`

	// The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
	// Defaults to once every hour if not specified.
	Schedule *MonitorInstanceSchedule `json:"schedule,omitempty"`

	// A set of schedules used to control how frequently the target is monitored for online and batch deployment type.
	Schedules *MonitorInstanceScheduleCollection `json:"schedules,omitempty"`

	// Available tags.
	Tags []MonitorTag `json:"tags" validate:"required"`
}

// UnmarshalMonitor unmarshals an instance of Monitor from the specified map of raw messages.
func UnmarshalMonitor(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Monitor)
	err = core.UnmarshalModel(m, "applies_to", &obj.AppliesTo, UnmarshalApplicabilitySelection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dictionary", &obj.Dictionary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMonitorMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "monitor_runtime", &obj.MonitorRuntime, UnmarshalMonitorRuntime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters_schema", &obj.ParametersSchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schedule", &obj.Schedule, UnmarshalMonitorInstanceSchedule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schedules", &obj.Schedules, UnmarshalMonitorInstanceScheduleCollection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalMonitorTag)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorCollections : MonitorCollections struct
type MonitorCollections struct {
	MonitorDefinitions []MonitorDisplayForm `json:"monitor_definitions" validate:"required"`
}

// UnmarshalMonitorCollections unmarshals an instance of MonitorCollections from the specified map of raw messages.
func UnmarshalMonitorCollections(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorCollections)
	err = core.UnmarshalModel(m, "monitor_definitions", &obj.MonitorDefinitions, UnmarshalMonitorDisplayForm)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorDisplayForm : MonitorDisplayForm struct
type MonitorDisplayForm struct {
	Entity *Monitor `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalMonitorDisplayForm unmarshals an instance of MonitorDisplayForm from the specified map of raw messages.
func UnmarshalMonitorDisplayForm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorDisplayForm)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalMonitor)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonOpenScaleV2) NewMonitorDisplayFormPatch(monitorDisplayForm *MonitorDisplayForm) (_patch []JSONPatchOperation) {
	if monitorDisplayForm.Entity != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/entity"),
			Value: monitorDisplayForm.Entity,
		})
	}
	if monitorDisplayForm.Metadata != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/metadata"),
			Value: monitorDisplayForm.Metadata,
		})
	}
	return
}

// MonitorInstanceCollection : MonitorInstanceCollection struct
type MonitorInstanceCollection struct {
	MonitorInstances []MonitorInstanceResponse `json:"monitor_instances" validate:"required"`
}

// UnmarshalMonitorInstanceCollection unmarshals an instance of MonitorInstanceCollection from the specified map of raw messages.
func UnmarshalMonitorInstanceCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorInstanceCollection)
	err = core.UnmarshalModel(m, "monitor_instances", &obj.MonitorInstances, UnmarshalMonitorInstanceResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorInstanceResponse : MonitorInstanceResponse struct
type MonitorInstanceResponse struct {
	Entity *MonitorInstanceResponseEntity `json:"entity,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`
}

// UnmarshalMonitorInstanceResponse unmarshals an instance of MonitorInstanceResponse from the specified map of raw messages.
func UnmarshalMonitorInstanceResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorInstanceResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalMonitorInstanceResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorInstanceResponseEntity : MonitorInstanceResponseEntity struct
type MonitorInstanceResponseEntity struct {
	DataMartID *string `json:"data_mart_id" validate:"required"`

	ManagedBy *string `json:"managed_by,omitempty"`

	MonitorDefinitionID *string `json:"monitor_definition_id" validate:"required"`

	// Monitoring parameters consistent with the `parameters_schema` from the monitor definition.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
	// Defaults to once every hour if not specified.
	Schedule *MonitorInstanceSchedule `json:"schedule,omitempty"`

	ScheduleID *string `json:"schedule_id,omitempty"`

	// A set of schedules used to control how frequently the target is monitored for online and batch deployment type.
	Schedules *MonitorInstanceScheduleCollection `json:"schedules,omitempty"`

	Target *Target `json:"target" validate:"required"`

	Thresholds []MetricThresholdOverride `json:"thresholds,omitempty"`

	// Summary about records count.
	TotalRecords *RecordsCountSummary `json:"total_records,omitempty"`

	// Summary about records count.
	UnprocessedRecords *RecordsCountSummary `json:"unprocessed_records,omitempty"`

	Status *MonitorInstanceResponseEntityStatus `json:"status" validate:"required"`
}

// UnmarshalMonitorInstanceResponseEntity unmarshals an instance of MonitorInstanceResponseEntity from the specified map of raw messages.
func UnmarshalMonitorInstanceResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorInstanceResponseEntity)
	err = core.UnmarshalPrimitive(m, "data_mart_id", &obj.DataMartID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "managed_by", &obj.ManagedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor_definition_id", &obj.MonitorDefinitionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schedule", &obj.Schedule, UnmarshalMonitorInstanceSchedule)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schedule_id", &obj.ScheduleID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schedules", &obj.Schedules, UnmarshalMonitorInstanceScheduleCollection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "thresholds", &obj.Thresholds, UnmarshalMetricThresholdOverride)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "total_records", &obj.TotalRecords, UnmarshalRecordsCountSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "unprocessed_records", &obj.UnprocessedRecords, UnmarshalRecordsCountSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalMonitorInstanceResponseEntityStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorInstanceResponseEntityStatus : MonitorInstanceResponseEntityStatus struct
type MonitorInstanceResponseEntityStatus struct {
	ActivityStatus *MonitorInstanceResponseEntityStatusActivityStatus `json:"activity_status,omitempty"`

	Failure *GenericErrorResponse `json:"failure,omitempty"`

	State *string `json:"state" validate:"required"`
}

// Constants associated with the MonitorInstanceResponseEntityStatus.State property.
const (
	MonitorInstanceResponseEntityStatus_State_Active        = "active"
	MonitorInstanceResponseEntityStatus_State_Deleteing     = "deleteing"
	MonitorInstanceResponseEntityStatus_State_Failed        = "failed"
	MonitorInstanceResponseEntityStatus_State_PendingDelete = "pending_delete"
	MonitorInstanceResponseEntityStatus_State_Preparing     = "preparing"
)

// UnmarshalMonitorInstanceResponseEntityStatus unmarshals an instance of MonitorInstanceResponseEntityStatus from the specified map of raw messages.
func UnmarshalMonitorInstanceResponseEntityStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorInstanceResponseEntityStatus)
	err = core.UnmarshalModel(m, "activity_status", &obj.ActivityStatus, UnmarshalMonitorInstanceResponseEntityStatusActivityStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalGenericErrorResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorInstanceResponseEntityStatusActivityStatus : MonitorInstanceResponseEntityStatusActivityStatus struct
type MonitorInstanceResponseEntityStatusActivityStatus struct {
	ID *string `json:"id,omitempty"`

	URL *string `json:"url,omitempty"`
}

// UnmarshalMonitorInstanceResponseEntityStatusActivityStatus unmarshals an instance of MonitorInstanceResponseEntityStatusActivityStatus from the specified map of raw messages.
func UnmarshalMonitorInstanceResponseEntityStatusActivityStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorInstanceResponseEntityStatusActivityStatus)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorInstanceSchedule : The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
// Defaults to once every hour if not specified.
type MonitorInstanceSchedule struct {
	// The interval to monitor the target.
	RepeatInterval *int64 `json:"repeat_interval" validate:"required"`

	// The type of interval to monitor the target.
	RepeatType *string `json:"repeat_type,omitempty"`

	// The type of interval to monitor the target.
	RepeatUnit *string `json:"repeat_unit" validate:"required"`

	// Defintion of first run time for scheduled activity; either absolute or relative the the moment of activation.
	StartTime *ScheduleStartTime `json:"start_time,omitempty"`
}

// Constants associated with the MonitorInstanceSchedule.RepeatUnit property.
// The type of interval to monitor the target.
const (
	MonitorInstanceSchedule_RepeatUnit_Day    = "day"
	MonitorInstanceSchedule_RepeatUnit_Hour   = "hour"
	MonitorInstanceSchedule_RepeatUnit_Minute = "minute"
	MonitorInstanceSchedule_RepeatUnit_Month  = "month"
	MonitorInstanceSchedule_RepeatUnit_Week   = "week"
	MonitorInstanceSchedule_RepeatUnit_Year   = "year"
)

// NewMonitorInstanceSchedule : Instantiate MonitorInstanceSchedule (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMonitorInstanceSchedule(repeatInterval int64, repeatUnit string) (_model *MonitorInstanceSchedule, err error) {
	_model = &MonitorInstanceSchedule{
		RepeatInterval: core.Int64Ptr(repeatInterval),
		RepeatUnit:     core.StringPtr(repeatUnit),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMonitorInstanceSchedule unmarshals an instance of MonitorInstanceSchedule from the specified map of raw messages.
func UnmarshalMonitorInstanceSchedule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorInstanceSchedule)
	err = core.UnmarshalPrimitive(m, "repeat_interval", &obj.RepeatInterval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "repeat_type", &obj.RepeatType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "repeat_unit", &obj.RepeatUnit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "start_time", &obj.StartTime, UnmarshalScheduleStartTime)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorInstanceScheduleCollection : A set of schedules used to control how frequently the target is monitored for online and batch deployment type.
type MonitorInstanceScheduleCollection struct {
	// The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
	// Defaults to once every hour if not specified.
	Batch *MonitorInstanceSchedule `json:"batch,omitempty"`

	// The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
	// Defaults to once every hour if not specified.
	Online *MonitorInstanceSchedule `json:"online,omitempty"`
}

// UnmarshalMonitorInstanceScheduleCollection unmarshals an instance of MonitorInstanceScheduleCollection from the specified map of raw messages.
func UnmarshalMonitorInstanceScheduleCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorInstanceScheduleCollection)
	err = core.UnmarshalModel(m, "batch", &obj.Batch, UnmarshalMonitorInstanceSchedule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "online", &obj.Online, UnmarshalMonitorInstanceSchedule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementMetric : Value and limits for the metrics.
type MonitorMeasurementMetric struct {
	ID *string `json:"id" validate:"required"`

	LowerLimit *float64 `json:"lower_limit,omitempty"`

	UpperLimit *float64 `json:"upper_limit,omitempty"`

	Value *float64 `json:"value" validate:"required"`
}

// UnmarshalMonitorMeasurementMetric unmarshals an instance of MonitorMeasurementMetric from the specified map of raw messages.
func UnmarshalMonitorMeasurementMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementMetric)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "lower_limit", &obj.LowerLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "upper_limit", &obj.UpperLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementRequest : MonitorMeasurementRequest struct
type MonitorMeasurementRequest struct {
	// Revision number of the ML model or function used by the monitor.
	AssetRevision *string `json:"asset_revision,omitempty"`

	// Metrics grouped for a single measurement.
	Metrics []map[string]interface{} `json:"metrics" validate:"required"`

	// ID of the monitoring run which produced the measurement.
	RunID *string `json:"run_id,omitempty"`

	// The sources of the metric.
	Sources []Source `json:"sources,omitempty"`

	Timestamp *strfmt.DateTime `json:"timestamp" validate:"required"`
}

// NewMonitorMeasurementRequest : Instantiate MonitorMeasurementRequest (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMonitorMeasurementRequest(metrics []map[string]interface{}, timestamp *strfmt.DateTime) (_model *MonitorMeasurementRequest, err error) {
	_model = &MonitorMeasurementRequest{
		Metrics:   metrics,
		Timestamp: timestamp,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMonitorMeasurementRequest unmarshals an instance of MonitorMeasurementRequest from the specified map of raw messages.
func UnmarshalMonitorMeasurementRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementRequest)
	err = core.UnmarshalPrimitive(m, "asset_revision", &obj.AssetRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metrics", &obj.Metrics)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_id", &obj.RunID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sources", &obj.Sources, UnmarshalSource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementResponse : MonitorMeasurementResponse struct
type MonitorMeasurementResponse struct {
	Entity *MonitorMeasurementResponseEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalMonitorMeasurementResponse unmarshals an instance of MonitorMeasurementResponse from the specified map of raw messages.
func UnmarshalMonitorMeasurementResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalMonitorMeasurementResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementResponseCollection : MonitorMeasurementResponseCollection struct
type MonitorMeasurementResponseCollection struct {
	// End of the time range.
	End *strfmt.DateTime `json:"end,omitempty"`

	Measurements []MonitorMeasurementResponseCollectionMeasurementsItem `json:"measurements" validate:"required"`

	// Beginning of the time range.
	Start *strfmt.DateTime `json:"start,omitempty"`
}

// UnmarshalMonitorMeasurementResponseCollection unmarshals an instance of MonitorMeasurementResponseCollection from the specified map of raw messages.
func UnmarshalMonitorMeasurementResponseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementResponseCollection)
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "measurements", &obj.Measurements, UnmarshalMonitorMeasurementResponseCollectionMeasurementsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementResponseCollectionMeasurementsItem : MonitorMeasurementResponseCollectionMeasurementsItem struct
type MonitorMeasurementResponseCollectionMeasurementsItem struct {
	Entity *MeasurementEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalMonitorMeasurementResponseCollectionMeasurementsItem unmarshals an instance of MonitorMeasurementResponseCollectionMeasurementsItem from the specified map of raw messages.
func UnmarshalMonitorMeasurementResponseCollectionMeasurementsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementResponseCollectionMeasurementsItem)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalMeasurementEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementResponseEntity : MonitorMeasurementResponseEntity struct
type MonitorMeasurementResponseEntity struct {
	// Revision number of the ML model or function used by the monitor.
	AssetRevision *string `json:"asset_revision,omitempty"`

	// Number of the metrics with issues, which exceeded limits.
	IssueCount *int64 `json:"issue_count" validate:"required"`

	MonitorDefinitionID *string `json:"monitor_definition_id,omitempty"`

	MonitorInstanceID *string `json:"monitor_instance_id,omitempty"`

	// ID of the monitoring run which produced the measurement.
	RunID *string `json:"run_id,omitempty"`

	Target *Target `json:"target,omitempty"`

	Timestamp *strfmt.DateTime `json:"timestamp" validate:"required"`

	// Metrics grouped for a single measurement.
	Values []MonitorMeasurementValue `json:"values" validate:"required"`

	// The sources of the metric.
	Sources []Source `json:"sources,omitempty"`
}

// UnmarshalMonitorMeasurementResponseEntity unmarshals an instance of MonitorMeasurementResponseEntity from the specified map of raw messages.
func UnmarshalMonitorMeasurementResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementResponseEntity)
	err = core.UnmarshalPrimitive(m, "asset_revision", &obj.AssetRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "issue_count", &obj.IssueCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor_definition_id", &obj.MonitorDefinitionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor_instance_id", &obj.MonitorInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_id", &obj.RunID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalMonitorMeasurementValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sources", &obj.Sources, UnmarshalSource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementTag : Tag related to the metrics.
type MonitorMeasurementTag struct {
	ID *string `json:"id" validate:"required"`

	Value *string `json:"value" validate:"required"`
}

// UnmarshalMonitorMeasurementTag unmarshals an instance of MonitorMeasurementTag from the specified map of raw messages.
func UnmarshalMonitorMeasurementTag(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementTag)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMeasurementValue : Measurment metrics and tags.
type MonitorMeasurementValue struct {
	// Metrics related to the measurement.
	Metrics []MonitorMeasurementMetric `json:"metrics" validate:"required"`

	// Tags related to the measurement.
	Tags []MonitorMeasurementTag `json:"tags" validate:"required"`
}

// UnmarshalMonitorMeasurementValue unmarshals an instance of MonitorMeasurementValue from the specified map of raw messages.
func UnmarshalMonitorMeasurementValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMeasurementValue)
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMonitorMeasurementMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalMonitorMeasurementTag)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMetric : MonitorMetric struct
type MonitorMetric struct {
	AppliesTo *ApplicabilitySelection `json:"applies_to,omitempty"`

	DefaultAggregation *string `json:"default_aggregation,omitempty"`

	// Description of the metrics.
	Description *string `json:"description,omitempty"`

	// the indicator specifying the expected direction of the monotonic metric values.
	ExpectedDirection *string `json:"expected_direction,omitempty"`

	// unique name used by UI instead of id (must be unique in scope of the monitor defition across both metrics and tags).
	Name *string `json:"name" validate:"required"`

	Required *bool `json:"required,omitempty"`

	Thresholds []MetricThreshold `json:"thresholds,omitempty"`

	ID *string `json:"id" validate:"required"`
}

// Constants associated with the MonitorMetric.ExpectedDirection property.
// the indicator specifying the expected direction of the monotonic metric values.
const (
	MonitorMetric_ExpectedDirection_Decreasing = "decreasing"
	MonitorMetric_ExpectedDirection_Increasing = "increasing"
	MonitorMetric_ExpectedDirection_Unknown    = "unknown"
)

// UnmarshalMonitorMetric unmarshals an instance of MonitorMetric from the specified map of raw messages.
func UnmarshalMonitorMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMetric)
	err = core.UnmarshalModel(m, "applies_to", &obj.AppliesTo, UnmarshalApplicabilitySelection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_aggregation", &obj.DefaultAggregation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_direction", &obj.ExpectedDirection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "thresholds", &obj.Thresholds, UnmarshalMetricThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorMetricRequest : MonitorMetricRequest struct
type MonitorMetricRequest struct {
	AppliesTo *ApplicabilitySelection `json:"applies_to,omitempty"`

	DefaultAggregation *string `json:"default_aggregation,omitempty"`

	// Description of the metrics.
	Description *string `json:"description,omitempty"`

	// the indicator specifying the expected direction of the monotonic metric values.
	ExpectedDirection *string `json:"expected_direction,omitempty"`

	// unique name used by UI instead of id (must be unique in scope of the monitor defition across both metrics and tags).
	Name *string `json:"name" validate:"required"`

	Required *bool `json:"required,omitempty"`

	Thresholds []MetricThreshold `json:"thresholds,omitempty"`
}

// Constants associated with the MonitorMetricRequest.ExpectedDirection property.
// the indicator specifying the expected direction of the monotonic metric values.
const (
	MonitorMetricRequest_ExpectedDirection_Decreasing = "decreasing"
	MonitorMetricRequest_ExpectedDirection_Increasing = "increasing"
	MonitorMetricRequest_ExpectedDirection_Unknown    = "unknown"
)

// NewMonitorMetricRequest : Instantiate MonitorMetricRequest (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMonitorMetricRequest(name string) (_model *MonitorMetricRequest, err error) {
	_model = &MonitorMetricRequest{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMonitorMetricRequest unmarshals an instance of MonitorMetricRequest from the specified map of raw messages.
func UnmarshalMonitorMetricRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorMetricRequest)
	err = core.UnmarshalModel(m, "applies_to", &obj.AppliesTo, UnmarshalApplicabilitySelection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_aggregation", &obj.DefaultAggregation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_direction", &obj.ExpectedDirection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "thresholds", &obj.Thresholds, UnmarshalMetricThreshold)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorRuntime : MonitorRuntime struct
type MonitorRuntime struct {
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the MonitorRuntime.Type property.
const (
	MonitorRuntime_Type_None    = "none"
	MonitorRuntime_Type_Service = "service"
)

// UnmarshalMonitorRuntime unmarshals an instance of MonitorRuntime from the specified map of raw messages.
func UnmarshalMonitorRuntime(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorRuntime)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorTag : MonitorTag struct
type MonitorTag struct {
	// Description of the tag.
	Description *string `json:"description,omitempty"`

	// unique name used by UI instead of id (must be unique in scope of the monitor defition across both metrics and tags).
	Name *string `json:"name" validate:"required"`

	Required *bool `json:"required,omitempty"`

	ID *string `json:"id" validate:"required"`
}

// UnmarshalMonitorTag unmarshals an instance of MonitorTag from the specified map of raw messages.
func UnmarshalMonitorTag(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorTag)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorTagRequest : MonitorTagRequest struct
type MonitorTagRequest struct {
	// Description of the tag.
	Description *string `json:"description,omitempty"`

	// unique name used by UI instead of id (must be unique in scope of the monitor defition across both metrics and tags).
	Name *string `json:"name" validate:"required"`

	Required *bool `json:"required,omitempty"`
}

// NewMonitorTagRequest : Instantiate MonitorTagRequest (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMonitorTagRequest(name string) (_model *MonitorTagRequest, err error) {
	_model = &MonitorTagRequest{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMonitorTagRequest unmarshals an instance of MonitorTagRequest from the specified map of raw messages.
func UnmarshalMonitorTagRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitorTagRequest)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitoringRun : MonitoringRun struct
type MonitoringRun struct {
	Entity *MonitoringRunEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalMonitoringRun unmarshals an instance of MonitoringRun from the specified map of raw messages.
func UnmarshalMonitoringRun(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitoringRun)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalMonitoringRunEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonOpenScaleV2) NewMonitoringRunPatch(monitoringRun *MonitoringRun) (_patch []JSONPatchOperation) {
	if monitoringRun.Entity != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/entity"),
			Value: monitoringRun.Entity,
		})
	}
	if monitoringRun.Metadata != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/metadata"),
			Value: monitoringRun.Metadata,
		})
	}
	return
}

// MonitoringRunBusinessMetricContext : Properties defining the business metric context in the triggered run of AI metric calculation.
type MonitoringRunBusinessMetricContext struct {
	// Unique identifier of the business application, which models are monitored by OpenScale.
	BusinessApplicationID *string `json:"business_application_id" validate:"required"`

	MetricID *string `json:"metric_id" validate:"required"`

	TransactionBatchID *string `json:"transaction_batch_id" validate:"required"`

	// Unique identifier of the data set (like scoring, feedback or business payload).
	TransactionDataSetID *string `json:"transaction_data_set_id" validate:"required"`
}

// NewMonitoringRunBusinessMetricContext : Instantiate MonitoringRunBusinessMetricContext (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewMonitoringRunBusinessMetricContext(businessApplicationID string, metricID string, transactionBatchID string, transactionDataSetID string) (_model *MonitoringRunBusinessMetricContext, err error) {
	_model = &MonitoringRunBusinessMetricContext{
		BusinessApplicationID: core.StringPtr(businessApplicationID),
		MetricID:              core.StringPtr(metricID),
		TransactionBatchID:    core.StringPtr(transactionBatchID),
		TransactionDataSetID:  core.StringPtr(transactionDataSetID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMonitoringRunBusinessMetricContext unmarshals an instance of MonitoringRunBusinessMetricContext from the specified map of raw messages.
func UnmarshalMonitoringRunBusinessMetricContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitoringRunBusinessMetricContext)
	err = core.UnmarshalPrimitive(m, "business_application_id", &obj.BusinessApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metric_id", &obj.MetricID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transaction_batch_id", &obj.TransactionBatchID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transaction_data_set_id", &obj.TransactionDataSetID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitoringRunCollection : A page from a collection of monitoring runs.
type MonitoringRunCollection struct {
	First *CollectionUrlModel `json:"first,omitempty"`

	Last *CollectionUrlModel `json:"last,omitempty"`

	// The number of monitoring runs requested to be returned.
	Limit *int64 `json:"limit,omitempty"`

	Next *CollectionUrlModel `json:"next,omitempty"`

	Prev *CollectionUrlModel `json:"prev,omitempty"`

	// A page from a collection of monitoring runs.
	Runs []MonitoringRun `json:"runs" validate:"required"`

	// The total number of monitoring runs available.
	TotalCount *int64 `json:"total_count,omitempty"`
}

// UnmarshalMonitoringRunCollection unmarshals an instance of MonitoringRunCollection from the specified map of raw messages.
func UnmarshalMonitoringRunCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitoringRunCollection)
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalCollectionUrlModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalCollectionUrlModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalCollectionUrlModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "prev", &obj.Prev, UnmarshalCollectionUrlModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "runs", &obj.Runs, UnmarshalMonitoringRun)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitoringRunEntity : MonitoringRunEntity struct
type MonitoringRunEntity struct {
	// Monitoring parameters consistent with the `parameters_schema` from the monitor definition.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The status information for the monitoring run.
	Status *MonitoringRunStatus `json:"status" validate:"required"`

	// An identifier representing the source that triggered the run request (optional). One of: event, scheduler, user,
	// webhook.
	TriggeredBy *string `json:"triggered_by,omitempty"`
}

// Constants associated with the MonitoringRunEntity.TriggeredBy property.
// An identifier representing the source that triggered the run request (optional). One of: event, scheduler, user,
// webhook.
const (
	MonitoringRunEntity_TriggeredBy_BkpiManager = "bkpi_manager"
	MonitoringRunEntity_TriggeredBy_Event       = "event"
	MonitoringRunEntity_TriggeredBy_Scheduler   = "scheduler"
	MonitoringRunEntity_TriggeredBy_User        = "user"
	MonitoringRunEntity_TriggeredBy_Webhook     = "webhook"
)

// UnmarshalMonitoringRunEntity unmarshals an instance of MonitoringRunEntity from the specified map of raw messages.
func UnmarshalMonitoringRunEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitoringRunEntity)
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalMonitoringRunStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "triggered_by", &obj.TriggeredBy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitoringRunOperator : MonitoringRunOperator struct
type MonitoringRunOperator struct {
	ID *string `json:"id,omitempty"`

	// Result produced by the operator, if any.
	Result map[string]interface{} `json:"result,omitempty"`

	Status *MonitoringRunOperatorStatus `json:"status,omitempty"`
}

// UnmarshalMonitoringRunOperator unmarshals an instance of MonitoringRunOperator from the specified map of raw messages.
func UnmarshalMonitoringRunOperator(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitoringRunOperator)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalMonitoringRunOperatorStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitoringRunOperatorStatus : MonitoringRunOperatorStatus struct
type MonitoringRunOperatorStatus struct {
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	Failure *GenericErrorResponse `json:"failure,omitempty"`

	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	State *string `json:"state,omitempty"`
}

// Constants associated with the MonitoringRunOperatorStatus.State property.
const (
	MonitoringRunOperatorStatus_State_Error    = "error"
	MonitoringRunOperatorStatus_State_Finished = "finished"
	MonitoringRunOperatorStatus_State_Queued   = "queued"
	MonitoringRunOperatorStatus_State_Running  = "running"
)

// UnmarshalMonitoringRunOperatorStatus unmarshals an instance of MonitoringRunOperatorStatus from the specified map of raw messages.
func UnmarshalMonitoringRunOperatorStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitoringRunOperatorStatus)
	err = core.UnmarshalPrimitive(m, "completed_at", &obj.CompletedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalGenericErrorResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_at", &obj.StartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitoringRunStatus : The status information for the monitoring run.
type MonitoringRunStatus struct {
	// The timestamp when the monitoring run finished running (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	Failure *GenericErrorResponse `json:"failure,omitempty"`

	// Any message associated with the monitoring run.
	Message *string `json:"message,omitempty"`

	Operators []MonitoringRunOperator `json:"operators,omitempty"`

	// The timestamp when the monitoring run was queued to be run (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	QueuedAt *strfmt.DateTime `json:"queued_at,omitempty"`

	// The timestamp when the monitoring run was started running (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	State *string `json:"state,omitempty"`

	// The timestamp when the monitoring run was last updated (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`
}

// Constants associated with the MonitoringRunStatus.State property.
const (
	MonitoringRunStatus_State_Error    = "error"
	MonitoringRunStatus_State_Finished = "finished"
	MonitoringRunStatus_State_Queued   = "queued"
	MonitoringRunStatus_State_Running  = "running"
)

// UnmarshalMonitoringRunStatus unmarshals an instance of MonitoringRunStatus from the specified map of raw messages.
func UnmarshalMonitoringRunStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MonitoringRunStatus)
	err = core.UnmarshalPrimitive(m, "completed_at", &obj.CompletedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalGenericErrorResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "operators", &obj.Operators, UnmarshalMonitoringRunOperator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "queued_at", &obj.QueuedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_at", &obj.StartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorsAddOptions : The MonitorsAdd options.
type MonitorsAddOptions struct {
	// A list of metric definition.
	Metrics []MonitorMetricRequest `json:"metrics" validate:"required"`

	// Monitor UI label (must be unique).
	Name *string `json:"name" validate:"required"`

	// Available tags.
	Tags []MonitorTagRequest `json:"tags" validate:"required"`

	AppliesTo *ApplicabilitySelection `json:"applies_to,omitempty"`

	// Long monitoring description presented in monitoring catalog.
	Description *string `json:"description,omitempty"`

	ManagedBy *string `json:"managed_by,omitempty"`

	// JSON schema that will be used to validate monitoring parameters when enabled.
	ParametersSchema map[string]interface{} `json:"parameters_schema,omitempty"`

	// The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
	// Defaults to once every hour if not specified.
	Schedule *MonitorInstanceSchedule `json:"schedule,omitempty"`

	// A set of schedules used to control how frequently the target is monitored for online and batch deployment type.
	Schedules *MonitorInstanceScheduleCollection `json:"schedules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMonitorsAddOptions : Instantiate MonitorsAddOptions
func (*WatsonOpenScaleV2) NewMonitorsAddOptions(metrics []MonitorMetricRequest, name string, tags []MonitorTagRequest) *MonitorsAddOptions {
	return &MonitorsAddOptions{
		Metrics: metrics,
		Name:    core.StringPtr(name),
		Tags:    tags,
	}
}

// SetMetrics : Allow user to set Metrics
func (_options *MonitorsAddOptions) SetMetrics(metrics []MonitorMetricRequest) *MonitorsAddOptions {
	_options.Metrics = metrics
	return _options
}

// SetName : Allow user to set Name
func (_options *MonitorsAddOptions) SetName(name string) *MonitorsAddOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *MonitorsAddOptions) SetTags(tags []MonitorTagRequest) *MonitorsAddOptions {
	_options.Tags = tags
	return _options
}

// SetAppliesTo : Allow user to set AppliesTo
func (_options *MonitorsAddOptions) SetAppliesTo(appliesTo *ApplicabilitySelection) *MonitorsAddOptions {
	_options.AppliesTo = appliesTo
	return _options
}

// SetDescription : Allow user to set Description
func (_options *MonitorsAddOptions) SetDescription(description string) *MonitorsAddOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *MonitorsAddOptions) SetManagedBy(managedBy string) *MonitorsAddOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetParametersSchema : Allow user to set ParametersSchema
func (_options *MonitorsAddOptions) SetParametersSchema(parametersSchema map[string]interface{}) *MonitorsAddOptions {
	_options.ParametersSchema = parametersSchema
	return _options
}

// SetSchedule : Allow user to set Schedule
func (_options *MonitorsAddOptions) SetSchedule(schedule *MonitorInstanceSchedule) *MonitorsAddOptions {
	_options.Schedule = schedule
	return _options
}

// SetSchedules : Allow user to set Schedules
func (_options *MonitorsAddOptions) SetSchedules(schedules *MonitorInstanceScheduleCollection) *MonitorsAddOptions {
	_options.Schedules = schedules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MonitorsAddOptions) SetHeaders(param map[string]string) *MonitorsAddOptions {
	options.Headers = param
	return options
}

// MonitorsDeleteOptions : The MonitorsDelete options.
type MonitorsDeleteOptions struct {
	// Unique monitor definition ID.
	MonitorDefinitionID *string `json:"monitor_definition_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMonitorsDeleteOptions : Instantiate MonitorsDeleteOptions
func (*WatsonOpenScaleV2) NewMonitorsDeleteOptions(monitorDefinitionID string) *MonitorsDeleteOptions {
	return &MonitorsDeleteOptions{
		MonitorDefinitionID: core.StringPtr(monitorDefinitionID),
	}
}

// SetMonitorDefinitionID : Allow user to set MonitorDefinitionID
func (_options *MonitorsDeleteOptions) SetMonitorDefinitionID(monitorDefinitionID string) *MonitorsDeleteOptions {
	_options.MonitorDefinitionID = core.StringPtr(monitorDefinitionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MonitorsDeleteOptions) SetHeaders(param map[string]string) *MonitorsDeleteOptions {
	options.Headers = param
	return options
}

// MonitorsGetOptions : The MonitorsGet options.
type MonitorsGetOptions struct {
	// Unique monitor definition ID.
	MonitorDefinitionID *string `json:"monitor_definition_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMonitorsGetOptions : Instantiate MonitorsGetOptions
func (*WatsonOpenScaleV2) NewMonitorsGetOptions(monitorDefinitionID string) *MonitorsGetOptions {
	return &MonitorsGetOptions{
		MonitorDefinitionID: core.StringPtr(monitorDefinitionID),
	}
}

// SetMonitorDefinitionID : Allow user to set MonitorDefinitionID
func (_options *MonitorsGetOptions) SetMonitorDefinitionID(monitorDefinitionID string) *MonitorsGetOptions {
	_options.MonitorDefinitionID = core.StringPtr(monitorDefinitionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MonitorsGetOptions) SetHeaders(param map[string]string) *MonitorsGetOptions {
	options.Headers = param
	return options
}

// MonitorsListOptions : The MonitorsList options.
type MonitorsListOptions struct {
	// comma-separeted list of names.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMonitorsListOptions : Instantiate MonitorsListOptions
func (*WatsonOpenScaleV2) NewMonitorsListOptions() *MonitorsListOptions {
	return &MonitorsListOptions{}
}

// SetName : Allow user to set Name
func (_options *MonitorsListOptions) SetName(name string) *MonitorsListOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MonitorsListOptions) SetHeaders(param map[string]string) *MonitorsListOptions {
	options.Headers = param
	return options
}

// MonitorsPatchOptions : The MonitorsPatch options.
type MonitorsPatchOptions struct {
	// Unique monitor definition ID.
	MonitorDefinitionID *string `json:"monitor_definition_id" validate:"required,ne="`

	// Array of patch operations as defined in RFC 6902.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMonitorsPatchOptions : Instantiate MonitorsPatchOptions
func (*WatsonOpenScaleV2) NewMonitorsPatchOptions(monitorDefinitionID string, jsonPatchOperation []JSONPatchOperation) *MonitorsPatchOptions {
	return &MonitorsPatchOptions{
		MonitorDefinitionID: core.StringPtr(monitorDefinitionID),
		JSONPatchOperation:  jsonPatchOperation,
	}
}

// SetMonitorDefinitionID : Allow user to set MonitorDefinitionID
func (_options *MonitorsPatchOptions) SetMonitorDefinitionID(monitorDefinitionID string) *MonitorsPatchOptions {
	_options.MonitorDefinitionID = core.StringPtr(monitorDefinitionID)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *MonitorsPatchOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *MonitorsPatchOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MonitorsPatchOptions) SetHeaders(param map[string]string) *MonitorsPatchOptions {
	options.Headers = param
	return options
}

// MonitorsUpdateOptions : The MonitorsUpdate options.
type MonitorsUpdateOptions struct {
	// Unique monitor definition ID.
	MonitorDefinitionID *string `json:"monitor_definition_id" validate:"required,ne="`

	// A list of metric definition.
	Metrics []MonitorMetricRequest `json:"metrics" validate:"required"`

	// Monitor UI label (must be unique).
	Name *string `json:"name" validate:"required"`

	// Available tags.
	Tags []MonitorTagRequest `json:"tags" validate:"required"`

	AppliesTo *ApplicabilitySelection `json:"applies_to,omitempty"`

	// Long monitoring description presented in monitoring catalog.
	Description *string `json:"description,omitempty"`

	ManagedBy *string `json:"managed_by,omitempty"`

	// JSON schema that will be used to validate monitoring parameters when enabled.
	ParametersSchema map[string]interface{} `json:"parameters_schema,omitempty"`

	// The schedule used to control how frequently the target is monitored. The maximum frequency is once every 30 minutes.
	// Defaults to once every hour if not specified.
	Schedule *MonitorInstanceSchedule `json:"schedule,omitempty"`

	// A set of schedules used to control how frequently the target is monitored for online and batch deployment type.
	Schedules *MonitorInstanceScheduleCollection `json:"schedules,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMonitorsUpdateOptions : Instantiate MonitorsUpdateOptions
func (*WatsonOpenScaleV2) NewMonitorsUpdateOptions(monitorDefinitionID string, metrics []MonitorMetricRequest, name string, tags []MonitorTagRequest) *MonitorsUpdateOptions {
	return &MonitorsUpdateOptions{
		MonitorDefinitionID: core.StringPtr(monitorDefinitionID),
		Metrics:             metrics,
		Name:                core.StringPtr(name),
		Tags:                tags,
	}
}

// SetMonitorDefinitionID : Allow user to set MonitorDefinitionID
func (_options *MonitorsUpdateOptions) SetMonitorDefinitionID(monitorDefinitionID string) *MonitorsUpdateOptions {
	_options.MonitorDefinitionID = core.StringPtr(monitorDefinitionID)
	return _options
}

// SetMetrics : Allow user to set Metrics
func (_options *MonitorsUpdateOptions) SetMetrics(metrics []MonitorMetricRequest) *MonitorsUpdateOptions {
	_options.Metrics = metrics
	return _options
}

// SetName : Allow user to set Name
func (_options *MonitorsUpdateOptions) SetName(name string) *MonitorsUpdateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *MonitorsUpdateOptions) SetTags(tags []MonitorTagRequest) *MonitorsUpdateOptions {
	_options.Tags = tags
	return _options
}

// SetAppliesTo : Allow user to set AppliesTo
func (_options *MonitorsUpdateOptions) SetAppliesTo(appliesTo *ApplicabilitySelection) *MonitorsUpdateOptions {
	_options.AppliesTo = appliesTo
	return _options
}

// SetDescription : Allow user to set Description
func (_options *MonitorsUpdateOptions) SetDescription(description string) *MonitorsUpdateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetManagedBy : Allow user to set ManagedBy
func (_options *MonitorsUpdateOptions) SetManagedBy(managedBy string) *MonitorsUpdateOptions {
	_options.ManagedBy = core.StringPtr(managedBy)
	return _options
}

// SetParametersSchema : Allow user to set ParametersSchema
func (_options *MonitorsUpdateOptions) SetParametersSchema(parametersSchema map[string]interface{}) *MonitorsUpdateOptions {
	_options.ParametersSchema = parametersSchema
	return _options
}

// SetSchedule : Allow user to set Schedule
func (_options *MonitorsUpdateOptions) SetSchedule(schedule *MonitorInstanceSchedule) *MonitorsUpdateOptions {
	_options.Schedule = schedule
	return _options
}

// SetSchedules : Allow user to set Schedules
func (_options *MonitorsUpdateOptions) SetSchedules(schedules *MonitorInstanceScheduleCollection) *MonitorsUpdateOptions {
	_options.Schedules = schedules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MonitorsUpdateOptions) SetHeaders(param map[string]string) *MonitorsUpdateOptions {
	options.Headers = param
	return options
}

// OperationalSpace : Operational Space definition.
type OperationalSpace struct {
	// The description of the Operational Space.
	Description *string `json:"description,omitempty"`

	// The name of the Operational Space.
	Name *string `json:"name" validate:"required"`
}

// NewOperationalSpace : Instantiate OperationalSpace (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewOperationalSpace(name string) (_model *OperationalSpace, err error) {
	_model = &OperationalSpace{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalOperationalSpace unmarshals an instance of OperationalSpace from the specified map of raw messages.
func UnmarshalOperationalSpace(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OperationalSpace)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OperationalSpaceCollection : OperationalSpaceCollection struct
type OperationalSpaceCollection struct {
	OperationalSpaces []OperationalSpace `json:"operational_spaces" validate:"required"`
}

// UnmarshalOperationalSpaceCollection unmarshals an instance of OperationalSpaceCollection from the specified map of raw messages.
func UnmarshalOperationalSpaceCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OperationalSpaceCollection)
	err = core.UnmarshalModel(m, "operational_spaces", &obj.OperationalSpaces, UnmarshalOperationalSpace)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OperationalSpaceResponse : OperationalSpaceResponse struct
type OperationalSpaceResponse struct {
	// Operational Space definition.
	Entity *OperationalSpace `json:"entity,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`
}

// UnmarshalOperationalSpaceResponse unmarshals an instance of OperationalSpaceResponse from the specified map of raw messages.
func UnmarshalOperationalSpaceResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OperationalSpaceResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalOperationalSpace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonOpenScaleV2) NewOperationalSpaceResponsePatch(operationalSpaceResponse *OperationalSpaceResponse) (_patch []JSONPatchOperation) {
	if operationalSpaceResponse.Entity != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/entity"),
			Value: operationalSpaceResponse.Entity,
		})
	}
	if operationalSpaceResponse.Metadata != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/metadata"),
			Value: operationalSpaceResponse.Metadata,
		})
	}
	return
}

// OperationalSpacesAddOptions : The OperationalSpacesAdd options.
type OperationalSpacesAddOptions struct {
	// The name of the Operational Space.
	Name *string `json:"name" validate:"required"`

	// The description of the Operational Space.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewOperationalSpacesAddOptions : Instantiate OperationalSpacesAddOptions
func (*WatsonOpenScaleV2) NewOperationalSpacesAddOptions(name string) *OperationalSpacesAddOptions {
	return &OperationalSpacesAddOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *OperationalSpacesAddOptions) SetName(name string) *OperationalSpacesAddOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *OperationalSpacesAddOptions) SetDescription(description string) *OperationalSpacesAddOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *OperationalSpacesAddOptions) SetHeaders(param map[string]string) *OperationalSpacesAddOptions {
	options.Headers = param
	return options
}

// OperationalSpacesDeleteOptions : The OperationalSpacesDelete options.
type OperationalSpacesDeleteOptions struct {
	// Unique Operational Space ID.
	OperationalSpaceID *string `json:"operational_space_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewOperationalSpacesDeleteOptions : Instantiate OperationalSpacesDeleteOptions
func (*WatsonOpenScaleV2) NewOperationalSpacesDeleteOptions(operationalSpaceID string) *OperationalSpacesDeleteOptions {
	return &OperationalSpacesDeleteOptions{
		OperationalSpaceID: core.StringPtr(operationalSpaceID),
	}
}

// SetOperationalSpaceID : Allow user to set OperationalSpaceID
func (_options *OperationalSpacesDeleteOptions) SetOperationalSpaceID(operationalSpaceID string) *OperationalSpacesDeleteOptions {
	_options.OperationalSpaceID = core.StringPtr(operationalSpaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *OperationalSpacesDeleteOptions) SetHeaders(param map[string]string) *OperationalSpacesDeleteOptions {
	options.Headers = param
	return options
}

// OperationalSpacesGetOptions : The OperationalSpacesGet options.
type OperationalSpacesGetOptions struct {
	// Unique Operational Space ID.
	OperationalSpaceID *string `json:"operational_space_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewOperationalSpacesGetOptions : Instantiate OperationalSpacesGetOptions
func (*WatsonOpenScaleV2) NewOperationalSpacesGetOptions(operationalSpaceID string) *OperationalSpacesGetOptions {
	return &OperationalSpacesGetOptions{
		OperationalSpaceID: core.StringPtr(operationalSpaceID),
	}
}

// SetOperationalSpaceID : Allow user to set OperationalSpaceID
func (_options *OperationalSpacesGetOptions) SetOperationalSpaceID(operationalSpaceID string) *OperationalSpacesGetOptions {
	_options.OperationalSpaceID = core.StringPtr(operationalSpaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *OperationalSpacesGetOptions) SetHeaders(param map[string]string) *OperationalSpacesGetOptions {
	options.Headers = param
	return options
}

// OperationalSpacesListOptions : The OperationalSpacesList options.
type OperationalSpacesListOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewOperationalSpacesListOptions : Instantiate OperationalSpacesListOptions
func (*WatsonOpenScaleV2) NewOperationalSpacesListOptions() *OperationalSpacesListOptions {
	return &OperationalSpacesListOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *OperationalSpacesListOptions) SetHeaders(param map[string]string) *OperationalSpacesListOptions {
	options.Headers = param
	return options
}

// OperationalSpacesUpdateOptions : The OperationalSpacesUpdate options.
type OperationalSpacesUpdateOptions struct {
	// Unique Operational Space ID.
	OperationalSpaceID *string `json:"operational_space_id" validate:"required,ne="`

	// Array of patch operations as defined in RFC 6902.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewOperationalSpacesUpdateOptions : Instantiate OperationalSpacesUpdateOptions
func (*WatsonOpenScaleV2) NewOperationalSpacesUpdateOptions(operationalSpaceID string, jsonPatchOperation []JSONPatchOperation) *OperationalSpacesUpdateOptions {
	return &OperationalSpacesUpdateOptions{
		OperationalSpaceID: core.StringPtr(operationalSpaceID),
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetOperationalSpaceID : Allow user to set OperationalSpaceID
func (_options *OperationalSpacesUpdateOptions) SetOperationalSpaceID(operationalSpaceID string) *OperationalSpacesUpdateOptions {
	_options.OperationalSpaceID = core.StringPtr(operationalSpaceID)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *OperationalSpacesUpdateOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *OperationalSpacesUpdateOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *OperationalSpacesUpdateOptions) SetHeaders(param map[string]string) *OperationalSpacesUpdateOptions {
	options.Headers = param
	return options
}

// PatchDocument : A JSONPatch document as defined by RFC 6902.
type PatchDocument struct {
	// A string containing a JSON Pointer value.
	From *string `json:"from,omitempty"`

	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// A JSON-Pointer.
	Path *string `json:"path" validate:"required"`

	// The value to be used within the operations.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the PatchDocument.Op property.
// The operation to be performed.
const (
	PatchDocument_Op_Add     = "add"
	PatchDocument_Op_Copy    = "copy"
	PatchDocument_Op_Move    = "move"
	PatchDocument_Op_Remove  = "remove"
	PatchDocument_Op_Replace = "replace"
	PatchDocument_Op_Test    = "test"
)

// NewPatchDocument : Instantiate PatchDocument (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewPatchDocument(op string, path string) (_model *PatchDocument, err error) {
	_model = &PatchDocument{
		Op:   core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPatchDocument unmarshals an instance of PatchDocument from the specified map of raw messages.
func UnmarshalPatchDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PatchDocument)
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PayloadField : payload field.
type PayloadField struct {
	// field description.
	Description *string `json:"description,omitempty"`

	// field name.
	Name *string `json:"name" validate:"required"`

	// field type.
	Type *string `json:"type" validate:"required"`

	// field unit.
	Unit *string `json:"unit,omitempty"`
}

// Constants associated with the PayloadField.Type property.
// field type.
const (
	PayloadField_Type_Number = "number"
	PayloadField_Type_String = "string"
)

// NewPayloadField : Instantiate PayloadField (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewPayloadField(name string, typeVar string) (_model *PayloadField, err error) {
	_model = &PayloadField{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPayloadField unmarshals an instance of PayloadField from the specified map of raw messages.
func UnmarshalPayloadField(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PayloadField)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostExplanationTaskResponse : Post explanation tasks response.
type PostExplanationTaskResponse struct {
	// Metadata of post explanation tasks response.
	Metadata *PostExplanationTaskResponseMetadata `json:"metadata" validate:"required"`
}

// UnmarshalPostExplanationTaskResponse unmarshals an instance of PostExplanationTaskResponse from the specified map of raw messages.
func UnmarshalPostExplanationTaskResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostExplanationTaskResponse)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalPostExplanationTaskResponseMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostExplanationTaskResponseMetadata : Metadata of post explanation tasks response.
type PostExplanationTaskResponseMetadata struct {
	// Time when the explanation task was initiated.
	CreatedAt *string `json:"created_at" validate:"required"`

	// ID of the user creating explanation task.
	CreatedBy *string `json:"created_by" validate:"required"`

	// List of identifiers for tracking explanation tasks.
	ExplanationTaskIds []string `json:"explanation_task_ids" validate:"required"`
}

// UnmarshalPostExplanationTaskResponseMetadata unmarshals an instance of PostExplanationTaskResponseMetadata from the specified map of raw messages.
func UnmarshalPostExplanationTaskResponseMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostExplanationTaskResponseMetadata)
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "explanation_task_ids", &obj.ExplanationTaskIds)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrimaryStorageCredentials : PrimaryStorageCredentials struct
// Models which "extend" this model:
// - PrimaryStorageCredentialsShort
// - PrimaryStorageCredentialsLong
type PrimaryStorageCredentials struct {
	URI *string `json:"uri,omitempty"`

	// any additional properties to be included in connection url.
	AdditionalProperties map[string]interface{} `json:"additional_properties,omitempty"`

	// DER-encoded certificate in Base64 encoding. The decoded content must be bound at the beginning by -----BEGIN
	// CERTIFICATE----- and at the end by -----END CERTIFICATE-----.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`

	Db *string `json:"db,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Password *string `json:"password,omitempty"`

	Port *int64 `json:"port,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	// (postgresql only).
	Sslmode *string `json:"sslmode,omitempty"`

	Username *string `json:"username,omitempty"`
}

func (*PrimaryStorageCredentials) isaPrimaryStorageCredentials() bool {
	return true
}

type PrimaryStorageCredentialsIntf interface {
	isaPrimaryStorageCredentials() bool
}

// UnmarshalPrimaryStorageCredentials unmarshals an instance of PrimaryStorageCredentials from the specified map of raw messages.
func UnmarshalPrimaryStorageCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrimaryStorageCredentials)
	err = core.UnmarshalPrimitive(m, "uri", &obj.URI)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "additional_properties", &obj.AdditionalProperties)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "db", &obj.Db)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sslmode", &obj.Sslmode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsAddOptions : The RecordsAdd options.
type RecordsAddOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	DatasetRecordsPayloadItem []DatasetRecordsPayloadItemIntf `json:"DatasetRecordsPayloadItem,omitempty"`

	Body *string `json:"body,omitempty"`

	// The type of the input. A character encoding can be specified by including a `charset` parameter. For example,
	// 'text/csv;charset=utf-8'.
	ContentType *string `json:"Content-Type,omitempty"`

	// if not provided service will attempt to automatically detect header in the first line.
	Header *bool `json:"header,omitempty"`

	// skip number of rows from input.
	Skip *int64 `json:"skip,omitempty"`

	// limit for number of processed input rows.
	Limit *int64 `json:"limit,omitempty"`

	// delimiter character for data provided as csv.
	Delimiter *string `json:"delimiter,omitempty"`

	// expected behaviour on error.
	OnError *string `json:"on_error,omitempty"`

	// maximum length of single line in bytes (default 10MB).
	CsvMaxLineLength *float64 `json:"csv_max_line_length,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RecordsAddOptions.OnError property.
// expected behaviour on error.
const (
	RecordsAddOptions_OnError_Continue = "continue"
	RecordsAddOptions_OnError_Stop     = "stop"
)

// NewRecordsAddOptions : Instantiate RecordsAddOptions
func (*WatsonOpenScaleV2) NewRecordsAddOptions(dataSetID string) *RecordsAddOptions {
	return &RecordsAddOptions{
		DataSetID: core.StringPtr(dataSetID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *RecordsAddOptions) SetDataSetID(dataSetID string) *RecordsAddOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetDatasetRecordsPayloadItem : Allow user to set DatasetRecordsPayloadItem
func (_options *RecordsAddOptions) SetDatasetRecordsPayloadItem(datasetRecordsPayloadItem []DatasetRecordsPayloadItemIntf) *RecordsAddOptions {
	_options.DatasetRecordsPayloadItem = datasetRecordsPayloadItem
	return _options
}

// SetBody : Allow user to set Body
func (_options *RecordsAddOptions) SetBody(body string) *RecordsAddOptions {
	_options.Body = core.StringPtr(body)
	return _options
}

// SetContentType : Allow user to set ContentType
func (_options *RecordsAddOptions) SetContentType(contentType string) *RecordsAddOptions {
	_options.ContentType = core.StringPtr(contentType)
	return _options
}

// SetHeader : Allow user to set Header
func (_options *RecordsAddOptions) SetHeader(header bool) *RecordsAddOptions {
	_options.Header = core.BoolPtr(header)
	return _options
}

// SetSkip : Allow user to set Skip
func (_options *RecordsAddOptions) SetSkip(skip int64) *RecordsAddOptions {
	_options.Skip = core.Int64Ptr(skip)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *RecordsAddOptions) SetLimit(limit int64) *RecordsAddOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetDelimiter : Allow user to set Delimiter
func (_options *RecordsAddOptions) SetDelimiter(delimiter string) *RecordsAddOptions {
	_options.Delimiter = core.StringPtr(delimiter)
	return _options
}

// SetOnError : Allow user to set OnError
func (_options *RecordsAddOptions) SetOnError(onError string) *RecordsAddOptions {
	_options.OnError = core.StringPtr(onError)
	return _options
}

// SetCsvMaxLineLength : Allow user to set CsvMaxLineLength
func (_options *RecordsAddOptions) SetCsvMaxLineLength(csvMaxLineLength float64) *RecordsAddOptions {
	_options.CsvMaxLineLength = core.Float64Ptr(csvMaxLineLength)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecordsAddOptions) SetHeaders(param map[string]string) *RecordsAddOptions {
	options.Headers = param
	return options
}

// RecordsCountSummary : Summary about records count.
type RecordsCountSummary struct {
	Count *int64 `json:"count" validate:"required"`

	Failure *GenericErrorResponse `json:"failure,omitempty"`

	// timestamp of last consumed record (only for unprocessed_records).
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// The type of records time.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the RecordsCountSummary.Type property.
// The type of records time.
const (
	RecordsCountSummary_Type_Feedback       = "feedback"
	RecordsCountSummary_Type_PayloadLogging = "payload_logging"
)

// NewRecordsCountSummary : Instantiate RecordsCountSummary (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewRecordsCountSummary(count int64, typeVar string) (_model *RecordsCountSummary, err error) {
	_model = &RecordsCountSummary{
		Count: core.Int64Ptr(count),
		Type:  core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRecordsCountSummary unmarshals an instance of RecordsCountSummary from the specified map of raw messages.
func UnmarshalRecordsCountSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordsCountSummary)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalGenericErrorResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsFieldOptions : The RecordsField options.
type RecordsFieldOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// ID of the record.
	RecordID *string `json:"record_id" validate:"required,ne="`

	// field_name should map to db column name which value is to be retrieved.
	FieldName *string `json:"field_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRecordsFieldOptions : Instantiate RecordsFieldOptions
func (*WatsonOpenScaleV2) NewRecordsFieldOptions(dataSetID string, recordID string, fieldName string) *RecordsFieldOptions {
	return &RecordsFieldOptions{
		DataSetID: core.StringPtr(dataSetID),
		RecordID:  core.StringPtr(recordID),
		FieldName: core.StringPtr(fieldName),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *RecordsFieldOptions) SetDataSetID(dataSetID string) *RecordsFieldOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetRecordID : Allow user to set RecordID
func (_options *RecordsFieldOptions) SetRecordID(recordID string) *RecordsFieldOptions {
	_options.RecordID = core.StringPtr(recordID)
	return _options
}

// SetFieldName : Allow user to set FieldName
func (_options *RecordsFieldOptions) SetFieldName(fieldName string) *RecordsFieldOptions {
	_options.FieldName = core.StringPtr(fieldName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecordsFieldOptions) SetHeaders(param map[string]string) *RecordsFieldOptions {
	options.Headers = param
	return options
}

// RecordsGetOptions : The RecordsGet options.
type RecordsGetOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// ID of the record.
	RecordID *string `json:"record_id" validate:"required,ne="`

	// Binary data presentation format. By default, the binary field value is encoded to base64 string. If _reference_ is
	// chosen, every binary field is moved to the _references_ section with value set to an uri to the particular field
	// within the record that can be GET in a separate request.
	BinaryFormat *string `json:"binary_format,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RecordsGetOptions.BinaryFormat property.
// Binary data presentation format. By default, the binary field value is encoded to base64 string. If _reference_ is
// chosen, every binary field is moved to the _references_ section with value set to an uri to the particular field
// within the record that can be GET in a separate request.
const (
	RecordsGetOptions_BinaryFormat_Reference = "reference"
)

// NewRecordsGetOptions : Instantiate RecordsGetOptions
func (*WatsonOpenScaleV2) NewRecordsGetOptions(dataSetID string, recordID string) *RecordsGetOptions {
	return &RecordsGetOptions{
		DataSetID: core.StringPtr(dataSetID),
		RecordID:  core.StringPtr(recordID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *RecordsGetOptions) SetDataSetID(dataSetID string) *RecordsGetOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetRecordID : Allow user to set RecordID
func (_options *RecordsGetOptions) SetRecordID(recordID string) *RecordsGetOptions {
	_options.RecordID = core.StringPtr(recordID)
	return _options
}

// SetBinaryFormat : Allow user to set BinaryFormat
func (_options *RecordsGetOptions) SetBinaryFormat(binaryFormat string) *RecordsGetOptions {
	_options.BinaryFormat = core.StringPtr(binaryFormat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecordsGetOptions) SetHeaders(param map[string]string) *RecordsGetOptions {
	options.Headers = param
	return options
}

// RecordsListOptions : The RecordsList options.
type RecordsListOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// return records with timestamp grather then or equal to `start` parameter.
	Start *strfmt.DateTime `json:"start,omitempty"`

	// return records with timestamp lower then `end` parameter.
	End *strfmt.DateTime `json:"end,omitempty"`

	// limit for number of returned records. If the value is greater than 1000 than it will be truncated.
	Limit *int64 `json:"limit,omitempty"`

	// offset of returned records.
	Offset *int64 `json:"offset,omitempty"`

	// return record annotations with given names.
	Annotations []string `json:"annotations,omitempty"`

	// Only return records that match given filters. There are two types of filters, separated by commas:
	//   * normal filter (multiple are possible), {field_name}:{op}:{value} —
	//       filter records directly
	//   * joining filter (only a single one is possible), {data_set_id}.{field_name}:{op}:{value} —
	//       join a data set by transaction_id (the user must ensure it's provided!)
	//       and filter by this data set's records' field.
	//       Will fail if the user hasn't provided transaction_id for both data sets' records. Filters of different types
	// can be mixed. They are partly compatible with the ones in POST /v2/data_sets/{data_set_id}/distributions.
	//
	// Available operators:
	//   | op   |  meaning                    |     example        |  code equivalent         |
	//   |:----:|:---------------------------:|:------------------:|:------------------------:|
	//   | eq   |  equality                   | field:eq:value     |  field == value          |
	//   | gt   |  greater than               | field:gt:value     |  field > value           |
	//   | gte  |  greater or equal           | field:gte:value    |  field >= value          |
	//   | lt   |  less than                  | field:lt:value     |  field < value           |
	//   | lte  |  less or equal              | field:lte:value    |  field <= value          |
	//   | like |  matching a simple pattern* | field:like:pattern |  pattern.match(field)    |
	//   | in   |  is contained in list       | field:in:a;b;c     |  [a,b,c].contains(field) |
	//
	// * - "%" means "one or more character", "_" means "any single character", other characters have their usual,
	//     literal meaning (e.g. "|" means character "|").
	Filter *string `json:"filter,omitempty"`

	// If total_count should be included. It can have performance impact if total_count is calculated.
	IncludeTotalCount *bool `json:"include_total_count,omitempty"`

	// What JSON format to use on output.
	Format *string `json:"format,omitempty"`

	// Binary data presentation format. By default, the binary field value is encoded to base64 string. If _reference_ is
	// chosen, every binary field is moved to the _references_ section with value set to an uri to the particular field
	// within the record that can be GET in a separate request.
	BinaryFormat *string `json:"binary_format,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RecordsListOptions.Format property.
// What JSON format to use on output.
const (
	RecordsListOptions_Format_Dict = "dict"
	RecordsListOptions_Format_List = "list"
)

// Constants associated with the RecordsListOptions.BinaryFormat property.
// Binary data presentation format. By default, the binary field value is encoded to base64 string. If _reference_ is
// chosen, every binary field is moved to the _references_ section with value set to an uri to the particular field
// within the record that can be GET in a separate request.
const (
	RecordsListOptions_BinaryFormat_Reference = "reference"
)

// NewRecordsListOptions : Instantiate RecordsListOptions
func (*WatsonOpenScaleV2) NewRecordsListOptions(dataSetID string) *RecordsListOptions {
	return &RecordsListOptions{
		DataSetID: core.StringPtr(dataSetID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *RecordsListOptions) SetDataSetID(dataSetID string) *RecordsListOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *RecordsListOptions) SetStart(start *strfmt.DateTime) *RecordsListOptions {
	_options.Start = start
	return _options
}

// SetEnd : Allow user to set End
func (_options *RecordsListOptions) SetEnd(end *strfmt.DateTime) *RecordsListOptions {
	_options.End = end
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *RecordsListOptions) SetLimit(limit int64) *RecordsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *RecordsListOptions) SetOffset(offset int64) *RecordsListOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetAnnotations : Allow user to set Annotations
func (_options *RecordsListOptions) SetAnnotations(annotations []string) *RecordsListOptions {
	_options.Annotations = annotations
	return _options
}

// SetFilter : Allow user to set Filter
func (_options *RecordsListOptions) SetFilter(filter string) *RecordsListOptions {
	_options.Filter = core.StringPtr(filter)
	return _options
}

// SetIncludeTotalCount : Allow user to set IncludeTotalCount
func (_options *RecordsListOptions) SetIncludeTotalCount(includeTotalCount bool) *RecordsListOptions {
	_options.IncludeTotalCount = core.BoolPtr(includeTotalCount)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *RecordsListOptions) SetFormat(format string) *RecordsListOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetBinaryFormat : Allow user to set BinaryFormat
func (_options *RecordsListOptions) SetBinaryFormat(binaryFormat string) *RecordsListOptions {
	_options.BinaryFormat = core.StringPtr(binaryFormat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecordsListOptions) SetHeaders(param map[string]string) *RecordsListOptions {
	options.Headers = param
	return options
}

// RecordsListResponse : RecordsListResponse struct
// Models which "extend" this model:
// - RecordsListResponseDataRecordsResponseCollectionDict
// - RecordsListResponseDataRecordsResponseCollectionList
type RecordsListResponse struct {
	Records []DataRecordResponse `json:"records,omitempty"`

	// Number of all rows which satisfy the query. It is calculated and returned when include_total_count query param is
	// set to `true`.
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (*RecordsListResponse) isaRecordsListResponse() bool {
	return true
}

type RecordsListResponseIntf interface {
	isaRecordsListResponse() bool
}

// UnmarshalRecordsListResponse unmarshals an instance of RecordsListResponse from the specified map of raw messages.
func UnmarshalRecordsListResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordsListResponse)
	err = core.UnmarshalModel(m, "records", &obj.Records, UnmarshalDataRecordResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsPatchOptions : The RecordsPatch options.
type RecordsPatchOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	PatchDocument []PatchDocument `json:"PatchDocument" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRecordsPatchOptions : Instantiate RecordsPatchOptions
func (*WatsonOpenScaleV2) NewRecordsPatchOptions(dataSetID string, patchDocument []PatchDocument) *RecordsPatchOptions {
	return &RecordsPatchOptions{
		DataSetID:     core.StringPtr(dataSetID),
		PatchDocument: patchDocument,
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *RecordsPatchOptions) SetDataSetID(dataSetID string) *RecordsPatchOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetPatchDocument : Allow user to set PatchDocument
func (_options *RecordsPatchOptions) SetPatchDocument(patchDocument []PatchDocument) *RecordsPatchOptions {
	_options.PatchDocument = patchDocument
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecordsPatchOptions) SetHeaders(param map[string]string) *RecordsPatchOptions {
	options.Headers = param
	return options
}

// RecordsQueryOptions : The RecordsQuery options.
type RecordsQueryOptions struct {
	// a (single) data set type.
	DataSetType *string `json:"data_set_type" validate:"required"`

	// one or more record id values that should be matched.
	RecordID []string `json:"record_id,omitempty"`

	// one or more transaction id values that should be matched.
	TransactionID []string `json:"transaction_id,omitempty"`

	// beginning of the time range.
	Start *strfmt.DateTime `json:"start,omitempty"`

	// end of the time range.
	End *strfmt.DateTime `json:"end,omitempty"`

	// offset of returned explanations.
	Offset *int64 `json:"offset,omitempty"`

	// Maximum number of elements returned.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RecordsQueryOptions.DataSetType property.
// a (single) data set type.
const (
	RecordsQueryOptions_DataSetType_BusinessPayload    = "business_payload"
	RecordsQueryOptions_DataSetType_Custom             = "custom"
	RecordsQueryOptions_DataSetType_Explanations       = "explanations"
	RecordsQueryOptions_DataSetType_ExplanationsWhatif = "explanations_whatif"
	RecordsQueryOptions_DataSetType_Feedback           = "feedback"
	RecordsQueryOptions_DataSetType_ManualLabeling     = "manual_labeling"
	RecordsQueryOptions_DataSetType_PayloadLogging     = "payload_logging"
	RecordsQueryOptions_DataSetType_Training           = "training"
)

// NewRecordsQueryOptions : Instantiate RecordsQueryOptions
func (*WatsonOpenScaleV2) NewRecordsQueryOptions(dataSetType string) *RecordsQueryOptions {
	return &RecordsQueryOptions{
		DataSetType: core.StringPtr(dataSetType),
	}
}

// SetDataSetType : Allow user to set DataSetType
func (_options *RecordsQueryOptions) SetDataSetType(dataSetType string) *RecordsQueryOptions {
	_options.DataSetType = core.StringPtr(dataSetType)
	return _options
}

// SetRecordID : Allow user to set RecordID
func (_options *RecordsQueryOptions) SetRecordID(recordID []string) *RecordsQueryOptions {
	_options.RecordID = recordID
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *RecordsQueryOptions) SetTransactionID(transactionID []string) *RecordsQueryOptions {
	_options.TransactionID = transactionID
	return _options
}

// SetStart : Allow user to set Start
func (_options *RecordsQueryOptions) SetStart(start *strfmt.DateTime) *RecordsQueryOptions {
	_options.Start = start
	return _options
}

// SetEnd : Allow user to set End
func (_options *RecordsQueryOptions) SetEnd(end *strfmt.DateTime) *RecordsQueryOptions {
	_options.End = end
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *RecordsQueryOptions) SetOffset(offset int64) *RecordsQueryOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *RecordsQueryOptions) SetLimit(limit int64) *RecordsQueryOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecordsQueryOptions) SetHeaders(param map[string]string) *RecordsQueryOptions {
	options.Headers = param
	return options
}

// RecordsUpdateOptions : The RecordsUpdate options.
type RecordsUpdateOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// ID of the record.
	RecordID *string `json:"record_id" validate:"required,ne="`

	PatchDocument []PatchDocument `json:"PatchDocument" validate:"required"`

	// Binary data presentation format. By default, the binary field value is encoded to base64 string. If _reference_ is
	// chosen, every binary field is moved to the _references_ section with value set to an uri to the particular field
	// within the record that can be GET in a separate request.
	BinaryFormat *string `json:"binary_format,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RecordsUpdateOptions.BinaryFormat property.
// Binary data presentation format. By default, the binary field value is encoded to base64 string. If _reference_ is
// chosen, every binary field is moved to the _references_ section with value set to an uri to the particular field
// within the record that can be GET in a separate request.
const (
	RecordsUpdateOptions_BinaryFormat_Reference = "reference"
)

// NewRecordsUpdateOptions : Instantiate RecordsUpdateOptions
func (*WatsonOpenScaleV2) NewRecordsUpdateOptions(dataSetID string, recordID string, patchDocument []PatchDocument) *RecordsUpdateOptions {
	return &RecordsUpdateOptions{
		DataSetID:     core.StringPtr(dataSetID),
		RecordID:      core.StringPtr(recordID),
		PatchDocument: patchDocument,
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *RecordsUpdateOptions) SetDataSetID(dataSetID string) *RecordsUpdateOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetRecordID : Allow user to set RecordID
func (_options *RecordsUpdateOptions) SetRecordID(recordID string) *RecordsUpdateOptions {
	_options.RecordID = core.StringPtr(recordID)
	return _options
}

// SetPatchDocument : Allow user to set PatchDocument
func (_options *RecordsUpdateOptions) SetPatchDocument(patchDocument []PatchDocument) *RecordsUpdateOptions {
	_options.PatchDocument = patchDocument
	return _options
}

// SetBinaryFormat : Allow user to set BinaryFormat
func (_options *RecordsUpdateOptions) SetBinaryFormat(binaryFormat string) *RecordsUpdateOptions {
	_options.BinaryFormat = core.StringPtr(binaryFormat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RecordsUpdateOptions) SetHeaders(param map[string]string) *RecordsUpdateOptions {
	options.Headers = param
	return options
}

// RequestsGetOptions : The RequestsGet options.
type RequestsGetOptions struct {
	// ID of the data set.
	DataSetID *string `json:"data_set_id" validate:"required,ne="`

	// ID of the request.
	RequestID *string `json:"request_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRequestsGetOptions : Instantiate RequestsGetOptions
func (*WatsonOpenScaleV2) NewRequestsGetOptions(dataSetID string, requestID string) *RequestsGetOptions {
	return &RequestsGetOptions{
		DataSetID: core.StringPtr(dataSetID),
		RequestID: core.StringPtr(requestID),
	}
}

// SetDataSetID : Allow user to set DataSetID
func (_options *RequestsGetOptions) SetDataSetID(dataSetID string) *RequestsGetOptions {
	_options.DataSetID = core.StringPtr(dataSetID)
	return _options
}

// SetRequestID : Allow user to set RequestID
func (_options *RequestsGetOptions) SetRequestID(requestID string) *RequestsGetOptions {
	_options.RequestID = core.StringPtr(requestID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RequestsGetOptions) SetHeaders(param map[string]string) *RequestsGetOptions {
	options.Headers = param
	return options
}

// RiskEvaluationStatus : RiskEvaluationStatus struct
type RiskEvaluationStatus struct {
	// Optional comment to the evaluation.
	Comment *string `json:"comment,omitempty"`

	// Time of the evaluation.
	EvaluatedAt *string `json:"evaluated_at,omitempty"`

	// Author of the evaluation.
	EvaluatedBy *string `json:"evaluated_by,omitempty"`

	State *string `json:"state" validate:"required"`
}

// Constants associated with the RiskEvaluationStatus.State property.
const (
	RiskEvaluationStatus_State_Approved          = "approved"
	RiskEvaluationStatus_State_PendingEvaluation = "pending_evaluation"
	RiskEvaluationStatus_State_Rejected          = "rejected"
)

// NewRiskEvaluationStatus : Instantiate RiskEvaluationStatus (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewRiskEvaluationStatus(state string) (_model *RiskEvaluationStatus, err error) {
	_model = &RiskEvaluationStatus{
		State: core.StringPtr(state),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRiskEvaluationStatus unmarshals an instance of RiskEvaluationStatus from the specified map of raw messages.
func UnmarshalRiskEvaluationStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RiskEvaluationStatus)
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "evaluated_at", &obj.EvaluatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "evaluated_by", &obj.EvaluatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RunsAddOptions : The RunsAdd options.
type RunsAddOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Properties defining the business metric context in the triggered run of AI metric calculation.
	BusinessMetricContext *MonitoringRunBusinessMetricContext `json:"business_metric_context,omitempty"`

	// The timestamp when the monitoring run was created with expiry date (in the format YYYY-MM-DDTHH:mm:ssZ or
	// YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339).
	ExpirationDate *strfmt.DateTime `json:"expiration_date,omitempty"`

	// Monitoring parameters consistent with the `parameters_schema` from the monitor definition.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// An identifier representing the source that triggered the run request (optional). One of: event, scheduler, user,
	// webhook.
	TriggeredBy *string `json:"triggered_by,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the RunsAddOptions.TriggeredBy property.
// An identifier representing the source that triggered the run request (optional). One of: event, scheduler, user,
// webhook.
const (
	RunsAddOptions_TriggeredBy_BkpiManager = "bkpi_manager"
	RunsAddOptions_TriggeredBy_Event       = "event"
	RunsAddOptions_TriggeredBy_Scheduler   = "scheduler"
	RunsAddOptions_TriggeredBy_User        = "user"
	RunsAddOptions_TriggeredBy_Webhook     = "webhook"
)

// NewRunsAddOptions : Instantiate RunsAddOptions
func (*WatsonOpenScaleV2) NewRunsAddOptions(monitorInstanceID string) *RunsAddOptions {
	return &RunsAddOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *RunsAddOptions) SetMonitorInstanceID(monitorInstanceID string) *RunsAddOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetBusinessMetricContext : Allow user to set BusinessMetricContext
func (_options *RunsAddOptions) SetBusinessMetricContext(businessMetricContext *MonitoringRunBusinessMetricContext) *RunsAddOptions {
	_options.BusinessMetricContext = businessMetricContext
	return _options
}

// SetExpirationDate : Allow user to set ExpirationDate
func (_options *RunsAddOptions) SetExpirationDate(expirationDate *strfmt.DateTime) *RunsAddOptions {
	_options.ExpirationDate = expirationDate
	return _options
}

// SetParameters : Allow user to set Parameters
func (_options *RunsAddOptions) SetParameters(parameters map[string]interface{}) *RunsAddOptions {
	_options.Parameters = parameters
	return _options
}

// SetTriggeredBy : Allow user to set TriggeredBy
func (_options *RunsAddOptions) SetTriggeredBy(triggeredBy string) *RunsAddOptions {
	_options.TriggeredBy = core.StringPtr(triggeredBy)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunsAddOptions) SetHeaders(param map[string]string) *RunsAddOptions {
	options.Headers = param
	return options
}

// RunsGetOptions : The RunsGet options.
type RunsGetOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Unique monitoring run ID.
	MonitoringRunID *string `json:"monitoring_run_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRunsGetOptions : Instantiate RunsGetOptions
func (*WatsonOpenScaleV2) NewRunsGetOptions(monitorInstanceID string, monitoringRunID string) *RunsGetOptions {
	return &RunsGetOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
		MonitoringRunID:   core.StringPtr(monitoringRunID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *RunsGetOptions) SetMonitorInstanceID(monitorInstanceID string) *RunsGetOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetMonitoringRunID : Allow user to set MonitoringRunID
func (_options *RunsGetOptions) SetMonitoringRunID(monitoringRunID string) *RunsGetOptions {
	_options.MonitoringRunID = core.StringPtr(monitoringRunID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunsGetOptions) SetHeaders(param map[string]string) *RunsGetOptions {
	options.Headers = param
	return options
}

// RunsListOptions : The RunsList options.
type RunsListOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// The page token indicating where to start paging from.
	Start *string `json:"start,omitempty"`

	// The limit of the number of items to return, for example limit=50. If not specified a default of 100 will be  used.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRunsListOptions : Instantiate RunsListOptions
func (*WatsonOpenScaleV2) NewRunsListOptions(monitorInstanceID string) *RunsListOptions {
	return &RunsListOptions{
		MonitorInstanceID: core.StringPtr(monitorInstanceID),
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *RunsListOptions) SetMonitorInstanceID(monitorInstanceID string) *RunsListOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *RunsListOptions) SetStart(start string) *RunsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *RunsListOptions) SetLimit(limit int64) *RunsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunsListOptions) SetHeaders(param map[string]string) *RunsListOptions {
	options.Headers = param
	return options
}

// RunsUpdateOptions : The RunsUpdate options.
type RunsUpdateOptions struct {
	// Unique monitor instance ID.
	MonitorInstanceID *string `json:"monitor_instance_id" validate:"required,ne="`

	// Unique monitoring run ID.
	MonitoringRunID *string `json:"monitoring_run_id" validate:"required,ne="`

	// Array of patch operations as defined in RFC 6902.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRunsUpdateOptions : Instantiate RunsUpdateOptions
func (*WatsonOpenScaleV2) NewRunsUpdateOptions(monitorInstanceID string, monitoringRunID string, jsonPatchOperation []JSONPatchOperation) *RunsUpdateOptions {
	return &RunsUpdateOptions{
		MonitorInstanceID:  core.StringPtr(monitorInstanceID),
		MonitoringRunID:    core.StringPtr(monitoringRunID),
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetMonitorInstanceID : Allow user to set MonitorInstanceID
func (_options *RunsUpdateOptions) SetMonitorInstanceID(monitorInstanceID string) *RunsUpdateOptions {
	_options.MonitorInstanceID = core.StringPtr(monitorInstanceID)
	return _options
}

// SetMonitoringRunID : Allow user to set MonitoringRunID
func (_options *RunsUpdateOptions) SetMonitoringRunID(monitoringRunID string) *RunsUpdateOptions {
	_options.MonitoringRunID = core.StringPtr(monitoringRunID)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *RunsUpdateOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *RunsUpdateOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RunsUpdateOptions) SetHeaders(param map[string]string) *RunsUpdateOptions {
	options.Headers = param
	return options
}

// ScheduleStartTime : Defintion of first run time for scheduled activity; either absolute or relative the the moment of activation.
type ScheduleStartTime struct {
	// must be set if type is `relative`.
	Delay *int64 `json:"delay,omitempty"`

	// must be set if type is `relative`.
	DelayUnit *string `json:"delay_unit,omitempty"`

	// must be set if type is `absolute`.
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// The type of start time.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the ScheduleStartTime.DelayUnit property.
// must be set if type is `relative`.
const (
	ScheduleStartTime_DelayUnit_Day    = "day"
	ScheduleStartTime_DelayUnit_Hour   = "hour"
	ScheduleStartTime_DelayUnit_Minute = "minute"
	ScheduleStartTime_DelayUnit_Month  = "month"
	ScheduleStartTime_DelayUnit_Week   = "week"
	ScheduleStartTime_DelayUnit_Year   = "year"
)

// Constants associated with the ScheduleStartTime.Type property.
// The type of start time.
const (
	ScheduleStartTime_Type_Absolute = "absolute"
	ScheduleStartTime_Type_Relative = "relative"
)

// NewScheduleStartTime : Instantiate ScheduleStartTime (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewScheduleStartTime(typeVar string) (_model *ScheduleStartTime, err error) {
	_model = &ScheduleStartTime{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalScheduleStartTime unmarshals an instance of ScheduleStartTime from the specified map of raw messages.
func UnmarshalScheduleStartTime(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScheduleStartTime)
	err = core.UnmarshalPrimitive(m, "delay", &obj.Delay)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "delay_unit", &obj.DelayUnit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchemaInferenceResponse : Schema inference response.
type SchemaInferenceResponse struct {
	// File data asset metadata.
	FileAssetMetadata *FileAssetMetadata `json:"file_asset_metadata,omitempty"`

	Subscription *SubscriptionResponse `json:"subscription" validate:"required"`
}

// UnmarshalSchemaInferenceResponse unmarshals an instance of SchemaInferenceResponse from the specified map of raw messages.
func UnmarshalSchemaInferenceResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchemaInferenceResponse)
	err = core.UnmarshalModel(m, "file_asset_metadata", &obj.FileAssetMetadata, UnmarshalFileAssetMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subscription", &obj.Subscription, UnmarshalSubscriptionResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoreData : Score data object.
type ScoreData struct {
	// Score fields.
	Fields []string `json:"fields,omitempty"`

	// Discriminates the data for multi input data situation. For example in cases where multiple tensors are expected.
	ID *string `json:"id,omitempty"`

	// Score value records.
	Values []interface{} `json:"values" validate:"required"`
}

// NewScoreData : Instantiate ScoreData (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewScoreData(values []interface{}) (_model *ScoreData, err error) {
	_model = &ScoreData{
		Values: values,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalScoreData unmarshals an instance of ScoreData from the specified map of raw messages.
func UnmarshalScoreData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoreData)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringEndpoint : Definition of scoring endpoint in custom_machine_learning.
type ScoringEndpoint struct {
	Credentials *SecretCleaned `json:"credentials,omitempty"`

	// map header name to header value.
	RequestHeaders map[string]interface{} `json:"request_headers,omitempty"`

	URL *string `json:"url,omitempty"`
}

// UnmarshalScoringEndpoint unmarshals an instance of ScoringEndpoint from the specified map of raw messages.
func UnmarshalScoringEndpoint(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringEndpoint)
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalSecretCleaned)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "request_headers", &obj.RequestHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringEndpointCredentials : ScoringEndpointCredentials struct
// Models which "extend" this model:
// - ScoringEndpointCredentialsAzureScoringEndpointCredentials
type ScoringEndpointCredentials struct {
	Token *string `json:"token,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

func (*ScoringEndpointCredentials) isaScoringEndpointCredentials() bool {
	return true
}

type ScoringEndpointCredentialsIntf interface {
	isaScoringEndpointCredentials() bool
	SetProperty(key string, value interface{})
	SetProperties(m map[string]interface{})
	GetProperty(key string) interface{}
	GetProperties() map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of ScoringEndpointCredentials
func (o *ScoringEndpointCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of ScoringEndpointCredentials
func (o *ScoringEndpointCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of ScoringEndpointCredentials
func (o *ScoringEndpointCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of ScoringEndpointCredentials
func (o *ScoringEndpointCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of ScoringEndpointCredentials
func (o *ScoringEndpointCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Token != nil {
		m["token"] = o.Token
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalScoringEndpointCredentials unmarshals an instance of ScoringEndpointCredentials from the specified map of raw messages.
func UnmarshalScoringEndpointCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringEndpointCredentials)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	delete(m, "token")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringEndpointRequest : Definition of scoring endpoint in custom_machine_learning.
type ScoringEndpointRequest struct {
	Credentials ScoringEndpointCredentialsIntf `json:"credentials,omitempty"`

	// map header name to header value.
	RequestHeaders map[string]interface{} `json:"request_headers,omitempty"`

	URL *string `json:"url,omitempty"`
}

// UnmarshalScoringEndpointRequest unmarshals an instance of ScoringEndpointRequest from the specified map of raw messages.
func UnmarshalScoringEndpointRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringEndpointRequest)
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalScoringEndpointCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "request_headers", &obj.RequestHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringPayloadRequestRequest : ScoringPayloadRequestRequest struct
type ScoringPayloadRequestRequest struct {
	// The data field names of model's or python function's input data. This property is mandatory for Spark based models.
	// It might not be required for other frameworks.
	Fields []string `json:"fields,omitempty"`

	Meta *ScoringPayloadRequestRequestMeta `json:"meta,omitempty"`

	// The scoring input data rows.
	Values [][]interface{} `json:"values" validate:"required"`
}

// NewScoringPayloadRequestRequest : Instantiate ScoringPayloadRequestRequest (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewScoringPayloadRequestRequest(values [][]interface{}) (_model *ScoringPayloadRequestRequest, err error) {
	_model = &ScoringPayloadRequestRequest{
		Values: values,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalScoringPayloadRequestRequest unmarshals an instance of ScoringPayloadRequestRequest from the specified map of raw messages.
func UnmarshalScoringPayloadRequestRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringPayloadRequestRequest)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "meta", &obj.Meta, UnmarshalScoringPayloadRequestRequestMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringPayloadRequestRequestMeta : ScoringPayloadRequestRequestMeta struct
type ScoringPayloadRequestRequestMeta struct {
	// The names of the additional columns which will be created in the payload logging table.
	Fields []string `json:"fields" validate:"required"`

	// Values for the additional columns which will be logged in the payload logging table.
	Values [][]interface{} `json:"values" validate:"required"`
}

// NewScoringPayloadRequestRequestMeta : Instantiate ScoringPayloadRequestRequestMeta (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewScoringPayloadRequestRequestMeta(fields []string, values [][]interface{}) (_model *ScoringPayloadRequestRequestMeta, err error) {
	_model = &ScoringPayloadRequestRequestMeta{
		Fields: fields,
		Values: values,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalScoringPayloadRequestRequestMeta unmarshals an instance of ScoringPayloadRequestRequestMeta from the specified map of raw messages.
func UnmarshalScoringPayloadRequestRequestMeta(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringPayloadRequestRequestMeta)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringPayloadRequestResponse : ScoringPayloadRequestResponse struct
type ScoringPayloadRequestResponse struct {
	// Names of the data fields.
	Fields []string `json:"fields,omitempty"`

	// The scoring output data rows.
	Values []interface{} `json:"values" validate:"required"`
}

// NewScoringPayloadRequestResponse : Instantiate ScoringPayloadRequestResponse (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewScoringPayloadRequestResponse(values []interface{}) (_model *ScoringPayloadRequestResponse, err error) {
	_model = &ScoringPayloadRequestResponse{
		Values: values,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalScoringPayloadRequestResponse unmarshals an instance of ScoringPayloadRequestResponse from the specified map of raw messages.
func UnmarshalScoringPayloadRequestResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringPayloadRequestResponse)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SecretCleaned : SecretCleaned struct
type SecretCleaned struct {
	// Generated id which identifies credentials.
	SecretID *string `json:"secret_id" validate:"required"`
}

// UnmarshalSecretCleaned unmarshals an instance of SecretCleaned from the specified map of raw messages.
func UnmarshalSecretCleaned(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SecretCleaned)
	err = core.UnmarshalPrimitive(m, "secret_id", &obj.SecretID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceProviderResponse : ServiceProviderResponse struct
type ServiceProviderResponse struct {
	Entity *ServiceProviderResponseEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalServiceProviderResponse unmarshals an instance of ServiceProviderResponse from the specified map of raw messages.
func UnmarshalServiceProviderResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceProviderResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalServiceProviderResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceProviderResponseCollection : ServiceProviderResponseCollection struct
type ServiceProviderResponseCollection struct {
	ServiceProviders []ServiceProviderResponse `json:"service_providers" validate:"required"`
}

// UnmarshalServiceProviderResponseCollection unmarshals an instance of ServiceProviderResponseCollection from the specified map of raw messages.
func UnmarshalServiceProviderResponseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceProviderResponseCollection)
	err = core.UnmarshalModel(m, "service_providers", &obj.ServiceProviders, UnmarshalServiceProviderResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceProviderResponseEntity : ServiceProviderResponseEntity struct
type ServiceProviderResponseEntity struct {
	Credentials *SecretCleaned `json:"credentials,omitempty"`

	// Reference to V2 Space ID.
	DeploymentSpaceID *string `json:"deployment_space_id,omitempty"`

	// ID of the ML service instance (required for Watson Machine Learning).
	InstanceID *string `json:"instance_id,omitempty"`

	// Name of the ML service instance.
	Name *string `json:"name" validate:"required"`

	// Reference to Operational Space.
	OperationalSpaceID *string `json:"operational_space_id,omitempty"`

	// Additional headers passed to the ML engine API (for example when scoring).
	RequestHeaders map[string]interface{} `json:"request_headers,omitempty"`

	ServiceType *string `json:"service_type" validate:"required"`

	Status *Status `json:"status,omitempty"`
}

// Constants associated with the ServiceProviderResponseEntity.ServiceType property.
const (
	ServiceProviderResponseEntity_ServiceType_AmazonSagemaker                        = "amazon_sagemaker"
	ServiceProviderResponseEntity_ServiceType_AzureMachineLearning                   = "azure_machine_learning"
	ServiceProviderResponseEntity_ServiceType_CustomMachineLearning                  = "custom_machine_learning"
	ServiceProviderResponseEntity_ServiceType_SpssCollaborationAndDeploymentServices = "spss_collaboration_and_deployment_services"
	ServiceProviderResponseEntity_ServiceType_WatsonMachineLearning                  = "watson_machine_learning"
)

// UnmarshalServiceProviderResponseEntity unmarshals an instance of ServiceProviderResponseEntity from the specified map of raw messages.
func UnmarshalServiceProviderResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceProviderResponseEntity)
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalSecretCleaned)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_space_id", &obj.DeploymentSpaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operational_space_id", &obj.OperationalSpaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "request_headers", &obj.RequestHeaders)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_type", &obj.ServiceType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceProvidersAddOptions : The ServiceProvidersAdd options.
type ServiceProvidersAddOptions struct {
	Credentials MLCredentialsIntf `json:"credentials" validate:"required"`

	// Name of the ML service instance.
	Name *string `json:"name" validate:"required"`

	// machine learning service type (azure_machine_learning_studio is a prefered alias for azure_machine_learning and
	// should be used in new service bindings).
	ServiceType *string `json:"service_type" validate:"required"`

	// Reference to V2 Space ID.
	DeploymentSpaceID *string `json:"deployment_space_id,omitempty"`

	Description *string `json:"description,omitempty"`

	// Reference to Operational Space.
	OperationalSpaceID *string `json:"operational_space_id,omitempty"`

	// map header name to header value.
	RequestHeaders map[string]interface{} `json:"request_headers,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ServiceProvidersAddOptions.ServiceType property.
// machine learning service type (azure_machine_learning_studio is a prefered alias for azure_machine_learning and
// should be used in new service bindings).
const (
	ServiceProvidersAddOptions_ServiceType_AmazonSagemaker                        = "amazon_sagemaker"
	ServiceProvidersAddOptions_ServiceType_AzureMachineLearning                   = "azure_machine_learning"
	ServiceProvidersAddOptions_ServiceType_AzureMachineLearningService            = "azure_machine_learning_service"
	ServiceProvidersAddOptions_ServiceType_AzureMachineLearningStudio             = "azure_machine_learning_studio"
	ServiceProvidersAddOptions_ServiceType_CustomMachineLearning                  = "custom_machine_learning"
	ServiceProvidersAddOptions_ServiceType_SpssCollaborationAndDeploymentServices = "spss_collaboration_and_deployment_services"
	ServiceProvidersAddOptions_ServiceType_WatsonMachineLearning                  = "watson_machine_learning"
)

// NewServiceProvidersAddOptions : Instantiate ServiceProvidersAddOptions
func (*WatsonOpenScaleV2) NewServiceProvidersAddOptions(credentials MLCredentialsIntf, name string, serviceType string) *ServiceProvidersAddOptions {
	return &ServiceProvidersAddOptions{
		Credentials: credentials,
		Name:        core.StringPtr(name),
		ServiceType: core.StringPtr(serviceType),
	}
}

// SetCredentials : Allow user to set Credentials
func (_options *ServiceProvidersAddOptions) SetCredentials(credentials MLCredentialsIntf) *ServiceProvidersAddOptions {
	_options.Credentials = credentials
	return _options
}

// SetName : Allow user to set Name
func (_options *ServiceProvidersAddOptions) SetName(name string) *ServiceProvidersAddOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetServiceType : Allow user to set ServiceType
func (_options *ServiceProvidersAddOptions) SetServiceType(serviceType string) *ServiceProvidersAddOptions {
	_options.ServiceType = core.StringPtr(serviceType)
	return _options
}

// SetDeploymentSpaceID : Allow user to set DeploymentSpaceID
func (_options *ServiceProvidersAddOptions) SetDeploymentSpaceID(deploymentSpaceID string) *ServiceProvidersAddOptions {
	_options.DeploymentSpaceID = core.StringPtr(deploymentSpaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ServiceProvidersAddOptions) SetDescription(description string) *ServiceProvidersAddOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetOperationalSpaceID : Allow user to set OperationalSpaceID
func (_options *ServiceProvidersAddOptions) SetOperationalSpaceID(operationalSpaceID string) *ServiceProvidersAddOptions {
	_options.OperationalSpaceID = core.StringPtr(operationalSpaceID)
	return _options
}

// SetRequestHeaders : Allow user to set RequestHeaders
func (_options *ServiceProvidersAddOptions) SetRequestHeaders(requestHeaders map[string]interface{}) *ServiceProvidersAddOptions {
	_options.RequestHeaders = requestHeaders
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ServiceProvidersAddOptions) SetHeaders(param map[string]string) *ServiceProvidersAddOptions {
	options.Headers = param
	return options
}

// ServiceProvidersDeleteOptions : The ServiceProvidersDelete options.
type ServiceProvidersDeleteOptions struct {
	// ID of the ML service provider.
	ServiceProviderID *string `json:"service_provider_id" validate:"required,ne="`

	// force hard delete.
	Force *bool `json:"force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewServiceProvidersDeleteOptions : Instantiate ServiceProvidersDeleteOptions
func (*WatsonOpenScaleV2) NewServiceProvidersDeleteOptions(serviceProviderID string) *ServiceProvidersDeleteOptions {
	return &ServiceProvidersDeleteOptions{
		ServiceProviderID: core.StringPtr(serviceProviderID),
	}
}

// SetServiceProviderID : Allow user to set ServiceProviderID
func (_options *ServiceProvidersDeleteOptions) SetServiceProviderID(serviceProviderID string) *ServiceProvidersDeleteOptions {
	_options.ServiceProviderID = core.StringPtr(serviceProviderID)
	return _options
}

// SetForce : Allow user to set Force
func (_options *ServiceProvidersDeleteOptions) SetForce(force bool) *ServiceProvidersDeleteOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ServiceProvidersDeleteOptions) SetHeaders(param map[string]string) *ServiceProvidersDeleteOptions {
	options.Headers = param
	return options
}

// ServiceProvidersGetOptions : The ServiceProvidersGet options.
type ServiceProvidersGetOptions struct {
	// ID of the ML service provider.
	ServiceProviderID *string `json:"service_provider_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewServiceProvidersGetOptions : Instantiate ServiceProvidersGetOptions
func (*WatsonOpenScaleV2) NewServiceProvidersGetOptions(serviceProviderID string) *ServiceProvidersGetOptions {
	return &ServiceProvidersGetOptions{
		ServiceProviderID: core.StringPtr(serviceProviderID),
	}
}

// SetServiceProviderID : Allow user to set ServiceProviderID
func (_options *ServiceProvidersGetOptions) SetServiceProviderID(serviceProviderID string) *ServiceProvidersGetOptions {
	_options.ServiceProviderID = core.StringPtr(serviceProviderID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ServiceProvidersGetOptions) SetHeaders(param map[string]string) *ServiceProvidersGetOptions {
	options.Headers = param
	return options
}

// ServiceProvidersListOptions : The ServiceProvidersList options.
type ServiceProvidersListOptions struct {
	// show also resources pending delete.
	ShowDeleted *bool `json:"show_deleted,omitempty"`

	// Type of service.
	ServiceType *string `json:"service_type,omitempty"`

	// comma-separeted list of IDs.
	InstanceID *string `json:"instance_id,omitempty"`

	// comma-separeted list of IDs.
	OperationalSpaceID *string `json:"operational_space_id,omitempty"`

	// comma-separeted list of IDs.
	DeploymentSpaceID *string `json:"deployment_space_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewServiceProvidersListOptions : Instantiate ServiceProvidersListOptions
func (*WatsonOpenScaleV2) NewServiceProvidersListOptions() *ServiceProvidersListOptions {
	return &ServiceProvidersListOptions{}
}

// SetShowDeleted : Allow user to set ShowDeleted
func (_options *ServiceProvidersListOptions) SetShowDeleted(showDeleted bool) *ServiceProvidersListOptions {
	_options.ShowDeleted = core.BoolPtr(showDeleted)
	return _options
}

// SetServiceType : Allow user to set ServiceType
func (_options *ServiceProvidersListOptions) SetServiceType(serviceType string) *ServiceProvidersListOptions {
	_options.ServiceType = core.StringPtr(serviceType)
	return _options
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ServiceProvidersListOptions) SetInstanceID(instanceID string) *ServiceProvidersListOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetOperationalSpaceID : Allow user to set OperationalSpaceID
func (_options *ServiceProvidersListOptions) SetOperationalSpaceID(operationalSpaceID string) *ServiceProvidersListOptions {
	_options.OperationalSpaceID = core.StringPtr(operationalSpaceID)
	return _options
}

// SetDeploymentSpaceID : Allow user to set DeploymentSpaceID
func (_options *ServiceProvidersListOptions) SetDeploymentSpaceID(deploymentSpaceID string) *ServiceProvidersListOptions {
	_options.DeploymentSpaceID = core.StringPtr(deploymentSpaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ServiceProvidersListOptions) SetHeaders(param map[string]string) *ServiceProvidersListOptions {
	options.Headers = param
	return options
}

// ServiceProvidersUpdateOptions : The ServiceProvidersUpdate options.
type ServiceProvidersUpdateOptions struct {
	// ID of the ML service provider.
	ServiceProviderID *string `json:"service_provider_id" validate:"required,ne="`

	PatchDocument []PatchDocument `json:"PatchDocument" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewServiceProvidersUpdateOptions : Instantiate ServiceProvidersUpdateOptions
func (*WatsonOpenScaleV2) NewServiceProvidersUpdateOptions(serviceProviderID string, patchDocument []PatchDocument) *ServiceProvidersUpdateOptions {
	return &ServiceProvidersUpdateOptions{
		ServiceProviderID: core.StringPtr(serviceProviderID),
		PatchDocument:     patchDocument,
	}
}

// SetServiceProviderID : Allow user to set ServiceProviderID
func (_options *ServiceProvidersUpdateOptions) SetServiceProviderID(serviceProviderID string) *ServiceProvidersUpdateOptions {
	_options.ServiceProviderID = core.StringPtr(serviceProviderID)
	return _options
}

// SetPatchDocument : Allow user to set PatchDocument
func (_options *ServiceProvidersUpdateOptions) SetPatchDocument(patchDocument []PatchDocument) *ServiceProvidersUpdateOptions {
	_options.PatchDocument = patchDocument
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ServiceProvidersUpdateOptions) SetHeaders(param map[string]string) *ServiceProvidersUpdateOptions {
	options.Headers = param
	return options
}

// Source : Source struct
type Source struct {
	// Data representing the source. It can be any value - object, string, number, boolean or array.
	Data interface{} `json:"data" validate:"required"`

	// id of the source.
	ID *string `json:"id" validate:"required"`

	// a selection of metrics that the source applies to (if not provided the source applies to all metrics).
	MetricIds []string `json:"metric_ids,omitempty"`

	// type of the source.
	Type *string `json:"type" validate:"required"`
}

// NewSource : Instantiate Source (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewSource(data interface{}, id string, typeVar string) (_model *Source, err error) {
	_model = &Source{
		Data: data,
		ID:   core.StringPtr(id),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSource unmarshals an instance of Source from the specified map of raw messages.
func UnmarshalSource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Source)
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metric_ids", &obj.MetricIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkStruct : SparkStruct struct
type SparkStruct struct {
	Fields []SparkStructField `json:"fields" validate:"required"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type" validate:"required"`
}

// NewSparkStruct : Instantiate SparkStruct (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewSparkStruct(fields []SparkStructField, typeVar string) (_model *SparkStruct, err error) {
	_model = &SparkStruct{
		Fields: fields,
		Type:   core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSparkStruct unmarshals an instance of SparkStruct from the specified map of raw messages.
func UnmarshalSparkStruct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkStruct)
	err = core.UnmarshalModel(m, "fields", &obj.Fields, UnmarshalSparkStructField)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SparkStructField : Spark struct field.
type SparkStructField struct {
	Name     *string `json:"name" validate:"required"`
	Type     *string `json:"type" validate:"required"`
	Nullable *bool   `json:"nullable" validate:"required"`
}

// UnmarshalSparkStructField unmarshals an instance of SparkStructField from the specified map of raw messages.
func UnmarshalSparkStructField(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SparkStructField)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nullable", &obj.Nullable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Status : Status struct
type Status struct {
	DeletedAt *string `json:"deleted_at,omitempty"`

	Failure *GenericErrorResponse `json:"failure,omitempty"`

	State *string `json:"state" validate:"required"`
}

// Constants associated with the Status.State property.
const (
	Status_State_Active        = "active"
	Status_State_Deleteing     = "deleteing"
	Status_State_Disabled      = "disabled"
	Status_State_Error         = "error"
	Status_State_PendingDelete = "pending_delete"
	Status_State_Preparing     = "preparing"
)

// UnmarshalStatus unmarshals an instance of Status from the specified map of raw messages.
func UnmarshalStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Status)
	err = core.UnmarshalPrimitive(m, "deleted_at", &obj.DeletedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalGenericErrorResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubscriptionResponse : SubscriptionResponse struct
type SubscriptionResponse struct {
	Entity *SubscriptionResponseEntity `json:"entity" validate:"required"`

	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalSubscriptionResponse unmarshals an instance of SubscriptionResponse from the specified map of raw messages.
func UnmarshalSubscriptionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SubscriptionResponse)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalSubscriptionResponseEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubscriptionResponseCollection : SubscriptionResponseCollection struct
type SubscriptionResponseCollection struct {
	Subscriptions []SubscriptionResponse `json:"subscriptions" validate:"required"`
}

// UnmarshalSubscriptionResponseCollection unmarshals an instance of SubscriptionResponseCollection from the specified map of raw messages.
func UnmarshalSubscriptionResponseCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SubscriptionResponseCollection)
	err = core.UnmarshalModel(m, "subscriptions", &obj.Subscriptions, UnmarshalSubscriptionResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubscriptionResponseEntity : SubscriptionResponseEntity struct
type SubscriptionResponseEntity struct {
	AnalyticsEngine *AnalyticsEngine `json:"analytics_engine,omitempty"`

	Asset *Asset `json:"asset" validate:"required"`

	// Additional asset properties (subject of discovery if not provided when creating the subscription).
	AssetProperties *AssetProperties `json:"asset_properties,omitempty"`

	DataMartID *string `json:"data_mart_id" validate:"required"`

	DataSources []DataSource `json:"data_sources,omitempty"`

	Deployment *AssetDeployment `json:"deployment" validate:"required"`

	// Integrated System reference.
	IntegrationReference *IntegratedSystemReference `json:"integration_reference,omitempty"`

	RiskEvaluationStatus *RiskEvaluationStatus `json:"risk_evaluation_status,omitempty"`

	ServiceProviderID *string `json:"service_provider_id" validate:"required"`

	Status *Status `json:"status,omitempty"`
}

// UnmarshalSubscriptionResponseEntity unmarshals an instance of SubscriptionResponseEntity from the specified map of raw messages.
func UnmarshalSubscriptionResponseEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SubscriptionResponseEntity)
	err = core.UnmarshalModel(m, "analytics_engine", &obj.AnalyticsEngine, UnmarshalAnalyticsEngine)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalAsset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "asset_properties", &obj.AssetProperties, UnmarshalAssetProperties)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_mart_id", &obj.DataMartID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_sources", &obj.DataSources, UnmarshalDataSource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalAssetDeployment)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "integration_reference", &obj.IntegrationReference, UnmarshalIntegratedSystemReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "risk_evaluation_status", &obj.RiskEvaluationStatus, UnmarshalRiskEvaluationStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_provider_id", &obj.ServiceProviderID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubscriptionsAddOptions : The SubscriptionsAdd options.
type SubscriptionsAddOptions struct {
	Asset *Asset `json:"asset" validate:"required"`

	DataMartID *string `json:"data_mart_id" validate:"required"`

	Deployment *AssetDeploymentRequest `json:"deployment" validate:"required"`

	ServiceProviderID *string `json:"service_provider_id" validate:"required"`

	AnalyticsEngine *AnalyticsEngine `json:"analytics_engine,omitempty"`

	// Additional asset properties (subject of discovery if not provided when creating the subscription).
	AssetProperties *AssetPropertiesRequest `json:"asset_properties,omitempty"`

	DataSources []DataSource `json:"data_sources,omitempty"`

	RiskEvaluationStatus *RiskEvaluationStatus `json:"risk_evaluation_status,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubscriptionsAddOptions : Instantiate SubscriptionsAddOptions
func (*WatsonOpenScaleV2) NewSubscriptionsAddOptions(asset *Asset, dataMartID string, deployment *AssetDeploymentRequest, serviceProviderID string) *SubscriptionsAddOptions {
	return &SubscriptionsAddOptions{
		Asset:             asset,
		DataMartID:        core.StringPtr(dataMartID),
		Deployment:        deployment,
		ServiceProviderID: core.StringPtr(serviceProviderID),
	}
}

// SetAsset : Allow user to set Asset
func (_options *SubscriptionsAddOptions) SetAsset(asset *Asset) *SubscriptionsAddOptions {
	_options.Asset = asset
	return _options
}

// SetDataMartID : Allow user to set DataMartID
func (_options *SubscriptionsAddOptions) SetDataMartID(dataMartID string) *SubscriptionsAddOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *SubscriptionsAddOptions) SetDeployment(deployment *AssetDeploymentRequest) *SubscriptionsAddOptions {
	_options.Deployment = deployment
	return _options
}

// SetServiceProviderID : Allow user to set ServiceProviderID
func (_options *SubscriptionsAddOptions) SetServiceProviderID(serviceProviderID string) *SubscriptionsAddOptions {
	_options.ServiceProviderID = core.StringPtr(serviceProviderID)
	return _options
}

// SetAnalyticsEngine : Allow user to set AnalyticsEngine
func (_options *SubscriptionsAddOptions) SetAnalyticsEngine(analyticsEngine *AnalyticsEngine) *SubscriptionsAddOptions {
	_options.AnalyticsEngine = analyticsEngine
	return _options
}

// SetAssetProperties : Allow user to set AssetProperties
func (_options *SubscriptionsAddOptions) SetAssetProperties(assetProperties *AssetPropertiesRequest) *SubscriptionsAddOptions {
	_options.AssetProperties = assetProperties
	return _options
}

// SetDataSources : Allow user to set DataSources
func (_options *SubscriptionsAddOptions) SetDataSources(dataSources []DataSource) *SubscriptionsAddOptions {
	_options.DataSources = dataSources
	return _options
}

// SetRiskEvaluationStatus : Allow user to set RiskEvaluationStatus
func (_options *SubscriptionsAddOptions) SetRiskEvaluationStatus(riskEvaluationStatus *RiskEvaluationStatus) *SubscriptionsAddOptions {
	_options.RiskEvaluationStatus = riskEvaluationStatus
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SubscriptionsAddOptions) SetHeaders(param map[string]string) *SubscriptionsAddOptions {
	options.Headers = param
	return options
}

// SubscriptionsDeleteOptions : The SubscriptionsDelete options.
type SubscriptionsDeleteOptions struct {
	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id" validate:"required,ne="`

	// force hard delete.
	Force *bool `json:"force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubscriptionsDeleteOptions : Instantiate SubscriptionsDeleteOptions
func (*WatsonOpenScaleV2) NewSubscriptionsDeleteOptions(subscriptionID string) *SubscriptionsDeleteOptions {
	return &SubscriptionsDeleteOptions{
		SubscriptionID: core.StringPtr(subscriptionID),
	}
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *SubscriptionsDeleteOptions) SetSubscriptionID(subscriptionID string) *SubscriptionsDeleteOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetForce : Allow user to set Force
func (_options *SubscriptionsDeleteOptions) SetForce(force bool) *SubscriptionsDeleteOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SubscriptionsDeleteOptions) SetHeaders(param map[string]string) *SubscriptionsDeleteOptions {
	options.Headers = param
	return options
}

// SubscriptionsGetOptions : The SubscriptionsGet options.
type SubscriptionsGetOptions struct {
	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubscriptionsGetOptions : Instantiate SubscriptionsGetOptions
func (*WatsonOpenScaleV2) NewSubscriptionsGetOptions(subscriptionID string) *SubscriptionsGetOptions {
	return &SubscriptionsGetOptions{
		SubscriptionID: core.StringPtr(subscriptionID),
	}
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *SubscriptionsGetOptions) SetSubscriptionID(subscriptionID string) *SubscriptionsGetOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SubscriptionsGetOptions) SetHeaders(param map[string]string) *SubscriptionsGetOptions {
	options.Headers = param
	return options
}

// SubscriptionsListOptions : The SubscriptionsList options.
type SubscriptionsListOptions struct {
	// comma-separeted list of IDs.
	DataMartID *string `json:"data_mart_id,omitempty"`

	// comma-separeted list of IDs.
	ServiceProviderID *string `json:"service_provider_id,omitempty"`

	// comma-separeted list of IDs.
	AssetAssetID *string `json:"asset.asset_id,omitempty"`

	// comma-separeted list of IDs.
	DeploymentDeploymentID *string `json:"deployment.deployment_id,omitempty"`

	// comma-separeted list of types.
	DeploymentDeploymentType *string `json:"deployment.deployment_type,omitempty"`

	// comma-separeted list of IDs.
	IntegrationReferenceIntegratedSystemID *string `json:"integration_reference.integrated_system_id,omitempty"`

	// comma-separeted list of IDs.
	IntegrationReferenceExternalID *string `json:"integration_reference.external_id,omitempty"`

	// comma-separeted list of states.
	RiskEvaluationStatusState *string `json:"risk_evaluation_status.state,omitempty"`

	// comma-separeted list of operational space ids (property of service provider object).
	ServiceProviderOperationalSpaceID *string `json:"service_provider.operational_space_id,omitempty"`

	// comma-separeted list of IDs.
	PreProductionReferenceID *string `json:"pre_production_reference_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubscriptionsListOptions : Instantiate SubscriptionsListOptions
func (*WatsonOpenScaleV2) NewSubscriptionsListOptions() *SubscriptionsListOptions {
	return &SubscriptionsListOptions{}
}

// SetDataMartID : Allow user to set DataMartID
func (_options *SubscriptionsListOptions) SetDataMartID(dataMartID string) *SubscriptionsListOptions {
	_options.DataMartID = core.StringPtr(dataMartID)
	return _options
}

// SetServiceProviderID : Allow user to set ServiceProviderID
func (_options *SubscriptionsListOptions) SetServiceProviderID(serviceProviderID string) *SubscriptionsListOptions {
	_options.ServiceProviderID = core.StringPtr(serviceProviderID)
	return _options
}

// SetAssetAssetID : Allow user to set AssetAssetID
func (_options *SubscriptionsListOptions) SetAssetAssetID(assetAssetID string) *SubscriptionsListOptions {
	_options.AssetAssetID = core.StringPtr(assetAssetID)
	return _options
}

// SetDeploymentDeploymentID : Allow user to set DeploymentDeploymentID
func (_options *SubscriptionsListOptions) SetDeploymentDeploymentID(deploymentDeploymentID string) *SubscriptionsListOptions {
	_options.DeploymentDeploymentID = core.StringPtr(deploymentDeploymentID)
	return _options
}

// SetDeploymentDeploymentType : Allow user to set DeploymentDeploymentType
func (_options *SubscriptionsListOptions) SetDeploymentDeploymentType(deploymentDeploymentType string) *SubscriptionsListOptions {
	_options.DeploymentDeploymentType = core.StringPtr(deploymentDeploymentType)
	return _options
}

// SetIntegrationReferenceIntegratedSystemID : Allow user to set IntegrationReferenceIntegratedSystemID
func (_options *SubscriptionsListOptions) SetIntegrationReferenceIntegratedSystemID(integrationReferenceIntegratedSystemID string) *SubscriptionsListOptions {
	_options.IntegrationReferenceIntegratedSystemID = core.StringPtr(integrationReferenceIntegratedSystemID)
	return _options
}

// SetIntegrationReferenceExternalID : Allow user to set IntegrationReferenceExternalID
func (_options *SubscriptionsListOptions) SetIntegrationReferenceExternalID(integrationReferenceExternalID string) *SubscriptionsListOptions {
	_options.IntegrationReferenceExternalID = core.StringPtr(integrationReferenceExternalID)
	return _options
}

// SetRiskEvaluationStatusState : Allow user to set RiskEvaluationStatusState
func (_options *SubscriptionsListOptions) SetRiskEvaluationStatusState(riskEvaluationStatusState string) *SubscriptionsListOptions {
	_options.RiskEvaluationStatusState = core.StringPtr(riskEvaluationStatusState)
	return _options
}

// SetServiceProviderOperationalSpaceID : Allow user to set ServiceProviderOperationalSpaceID
func (_options *SubscriptionsListOptions) SetServiceProviderOperationalSpaceID(serviceProviderOperationalSpaceID string) *SubscriptionsListOptions {
	_options.ServiceProviderOperationalSpaceID = core.StringPtr(serviceProviderOperationalSpaceID)
	return _options
}

// SetPreProductionReferenceID : Allow user to set PreProductionReferenceID
func (_options *SubscriptionsListOptions) SetPreProductionReferenceID(preProductionReferenceID string) *SubscriptionsListOptions {
	_options.PreProductionReferenceID = core.StringPtr(preProductionReferenceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SubscriptionsListOptions) SetHeaders(param map[string]string) *SubscriptionsListOptions {
	options.Headers = param
	return options
}

// SubscriptionsSchemasOptions : The SubscriptionsSchemas options.
type SubscriptionsSchemasOptions struct {
	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id" validate:"required,ne="`

	// Array of score data object. If multiple score data objects are included, the "fields" array (if any) for score
	// purposes will always be taken from the first score data object.
	InputData []ScoreData `json:"input_data,omitempty"`

	// InputDataReference is the same as TrainingDataReference except that neither location nor connection is required.
	// This is needed for the Schemas API and to avoid updating existing APIs.
	TrainingDataReference *InputDataReference `json:"training_data_reference,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubscriptionsSchemasOptions : Instantiate SubscriptionsSchemasOptions
func (*WatsonOpenScaleV2) NewSubscriptionsSchemasOptions(subscriptionID string) *SubscriptionsSchemasOptions {
	return &SubscriptionsSchemasOptions{
		SubscriptionID: core.StringPtr(subscriptionID),
	}
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *SubscriptionsSchemasOptions) SetSubscriptionID(subscriptionID string) *SubscriptionsSchemasOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetInputData : Allow user to set InputData
func (_options *SubscriptionsSchemasOptions) SetInputData(inputData []ScoreData) *SubscriptionsSchemasOptions {
	_options.InputData = inputData
	return _options
}

// SetTrainingDataReference : Allow user to set TrainingDataReference
func (_options *SubscriptionsSchemasOptions) SetTrainingDataReference(trainingDataReference *InputDataReference) *SubscriptionsSchemasOptions {
	_options.TrainingDataReference = trainingDataReference
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SubscriptionsSchemasOptions) SetHeaders(param map[string]string) *SubscriptionsSchemasOptions {
	options.Headers = param
	return options
}

// SubscriptionsScoreOptions : The SubscriptionsScore options.
type SubscriptionsScoreOptions struct {
	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id" validate:"required,ne="`

	// The values associated to the fields.
	Values []string `json:"values" validate:"required"`

	// The fields to process debias scoring.
	Fields []string `json:"fields,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubscriptionsScoreOptions : Instantiate SubscriptionsScoreOptions
func (*WatsonOpenScaleV2) NewSubscriptionsScoreOptions(subscriptionID string, values []string) *SubscriptionsScoreOptions {
	return &SubscriptionsScoreOptions{
		SubscriptionID: core.StringPtr(subscriptionID),
		Values:         values,
	}
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *SubscriptionsScoreOptions) SetSubscriptionID(subscriptionID string) *SubscriptionsScoreOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetValues : Allow user to set Values
func (_options *SubscriptionsScoreOptions) SetValues(values []string) *SubscriptionsScoreOptions {
	_options.Values = values
	return _options
}

// SetFields : Allow user to set Fields
func (_options *SubscriptionsScoreOptions) SetFields(fields []string) *SubscriptionsScoreOptions {
	_options.Fields = fields
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SubscriptionsScoreOptions) SetHeaders(param map[string]string) *SubscriptionsScoreOptions {
	options.Headers = param
	return options
}

// SubscriptionsUpdateOptions : The SubscriptionsUpdate options.
type SubscriptionsUpdateOptions struct {
	// Unique subscription ID.
	SubscriptionID *string `json:"subscription_id" validate:"required,ne="`

	PatchDocument []PatchDocument `json:"PatchDocument" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSubscriptionsUpdateOptions : Instantiate SubscriptionsUpdateOptions
func (*WatsonOpenScaleV2) NewSubscriptionsUpdateOptions(subscriptionID string, patchDocument []PatchDocument) *SubscriptionsUpdateOptions {
	return &SubscriptionsUpdateOptions{
		SubscriptionID: core.StringPtr(subscriptionID),
		PatchDocument:  patchDocument,
	}
}

// SetSubscriptionID : Allow user to set SubscriptionID
func (_options *SubscriptionsUpdateOptions) SetSubscriptionID(subscriptionID string) *SubscriptionsUpdateOptions {
	_options.SubscriptionID = core.StringPtr(subscriptionID)
	return _options
}

// SetPatchDocument : Allow user to set PatchDocument
func (_options *SubscriptionsUpdateOptions) SetPatchDocument(patchDocument []PatchDocument) *SubscriptionsUpdateOptions {
	_options.PatchDocument = patchDocument
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SubscriptionsUpdateOptions) SetHeaders(param map[string]string) *SubscriptionsUpdateOptions {
	options.Headers = param
	return options
}

// Target : Target struct
type Target struct {
	// ID of the data set target (e.g. subscription ID, business application ID, ...).
	TargetID *string `json:"target_id" validate:"required"`

	// Type of the target (e.g. subscription, business application, ...).
	TargetType *string `json:"target_type" validate:"required"`
}

// Constants associated with the Target.TargetType property.
// Type of the target (e.g. subscription, business application, ...).
const (
	Target_TargetType_BusinessApplication = "business_application"
	Target_TargetType_DataMart            = "data_mart"
	Target_TargetType_Instance            = "instance"
	Target_TargetType_Subscription        = "subscription"
)

// NewTarget : Instantiate Target (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewTarget(targetID string, targetType string) (_model *Target, err error) {
	_model = &Target{
		TargetID:   core.StringPtr(targetID),
		TargetType: core.StringPtr(targetType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTarget unmarshals an instance of Target from the specified map of raw messages.
func UnmarshalTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Target)
	err = core.UnmarshalPrimitive(m, "target_id", &obj.TargetID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_type", &obj.TargetType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ThresholdConditionObject : ThresholdConditionObject struct
type ThresholdConditionObject struct {
	Key *string `json:"key" validate:"required"`

	Type *string `json:"type" validate:"required"`

	Value *string `json:"value" validate:"required"`
}

// Constants associated with the ThresholdConditionObject.Type property.
const (
	ThresholdConditionObject_Type_Tag = "tag"
)

// NewThresholdConditionObject : Instantiate ThresholdConditionObject (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewThresholdConditionObject(key string, typeVar string, value string) (_model *ThresholdConditionObject, err error) {
	_model = &ThresholdConditionObject{
		Key:   core.StringPtr(key),
		Type:  core.StringPtr(typeVar),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalThresholdConditionObject unmarshals an instance of ThresholdConditionObject from the specified map of raw messages.
func UnmarshalThresholdConditionObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ThresholdConditionObject)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TimeFrame : TimeFrame struct
type TimeFrame struct {
	Count *int64 `json:"count" validate:"required"`

	Unit *string `json:"unit" validate:"required"`
}

// Constants associated with the TimeFrame.Unit property.
const (
	TimeFrame_Unit_Day    = "day"
	TimeFrame_Unit_Hour   = "hour"
	TimeFrame_Unit_Minute = "minute"
	TimeFrame_Unit_Month  = "month"
	TimeFrame_Unit_Week   = "week"
)

// NewTimeFrame : Instantiate TimeFrame (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewTimeFrame(count int64, unit string) (_model *TimeFrame, err error) {
	_model = &TimeFrame{
		Count: core.Int64Ptr(count),
		Unit:  core.StringPtr(unit),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTimeFrame unmarshals an instance of TimeFrame from the specified map of raw messages.
func UnmarshalTimeFrame(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TimeFrame)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDataReference : TrainingDataReference struct
type TrainingDataReference struct {
	// training data set connection credentials.
	Connection TrainingDataReferenceConnectionIntf `json:"connection" validate:"required"`

	// training data set location.
	Location TrainingDataReferenceLocationIntf `json:"location" validate:"required"`

	Name *string `json:"name,omitempty"`

	// Type of the storage.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the TrainingDataReference.Type property.
// Type of the storage.
const (
	TrainingDataReference_Type_Cos       = "cos"
	TrainingDataReference_Type_Dataset   = "dataset"
	TrainingDataReference_Type_Db2       = "db2"
	TrainingDataReference_Type_FileAsset = "file_asset"
)

// NewTrainingDataReference : Instantiate TrainingDataReference (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewTrainingDataReference(connection TrainingDataReferenceConnectionIntf, location TrainingDataReferenceLocationIntf, typeVar string) (_model *TrainingDataReference, err error) {
	_model = &TrainingDataReference{
		Connection: connection,
		Location:   location,
		Type:       core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTrainingDataReference unmarshals an instance of TrainingDataReference from the specified map of raw messages.
func UnmarshalTrainingDataReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDataReference)
	err = core.UnmarshalModel(m, "connection", &obj.Connection, UnmarshalTrainingDataReferenceConnection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalTrainingDataReferenceLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDataReferenceConnection : training data set connection credentials.
// Models which "extend" this model:
// - DB2TrainingDataReferenceConnection
// - COSTrainingDataReferenceConnection
type TrainingDataReferenceConnection struct {
	// DER-encoded certificate in Base64 encoding. The decoded content must be bound at the beginning by -----BEGIN
	// CERTIFICATE----- and at the end by -----END CERTIFICATE-----.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`

	ConnectionString *string `json:"connection_string,omitempty"`

	DatabaseName *string `json:"database_name,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Password *string `json:"password,omitempty"`

	Port *int64 `json:"port,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	Username *string `json:"username,omitempty"`

	ApiKey *string `json:"api_key,omitempty"`

	IamURL *string `json:"iam_url,omitempty"`

	ResourceInstanceID *string `json:"resource_instance_id,omitempty"`

	URL *string `json:"url,omitempty"`
}

func (*TrainingDataReferenceConnection) isaTrainingDataReferenceConnection() bool {
	return true
}

type TrainingDataReferenceConnectionIntf interface {
	isaTrainingDataReferenceConnection() bool
}

// UnmarshalTrainingDataReferenceConnection unmarshals an instance of TrainingDataReferenceConnection from the specified map of raw messages.
func UnmarshalTrainingDataReferenceConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDataReferenceConnection)
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iam_url", &obj.IamURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_instance_id", &obj.ResourceInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDataReferenceLocation : training data set location.
// Models which "extend" this model:
// - DB2TrainingDataReferenceLocation
// - COSTrainingDataReferenceLocation
// - DatasetTrainingDataReferenceLocation
// - FileAssetTrainingDataReferenceLocation
type TrainingDataReferenceLocation struct {
	// name of the schema.
	SchemaName *string `json:"schema_name,omitempty"`

	// name of the table.
	TableName *string `json:"table_name,omitempty"`

	Bucket *string `json:"bucket,omitempty"`

	FileFormat *string `json:"file_format,omitempty"`

	FileName *string `json:"file_name,omitempty"`

	Firstlineheader *bool `json:"firstlineheader,omitempty"`

	InferSchema *string `json:"infer_schema,omitempty"`

	// Dataset id.
	DatasetID *string `json:"dataset_id,omitempty"`

	// additional options for different types of training data references.
	Meta *FileTrainingDataReferenceOptions `json:"meta,omitempty"`

	// File data asset reference.
	AssetHref *string `json:"asset_href,omitempty"`

	// File data asset id.
	AssetID *string `json:"asset_id,omitempty"`

	// Project id.
	ProjectID *string `json:"project_id,omitempty"`
}

// Constants associated with the TrainingDataReferenceLocation.FileFormat property.
const (
	TrainingDataReferenceLocation_FileFormat_Csv = "csv"
)

func (*TrainingDataReferenceLocation) isaTrainingDataReferenceLocation() bool {
	return true
}

type TrainingDataReferenceLocationIntf interface {
	isaTrainingDataReferenceLocation() bool
}

// UnmarshalTrainingDataReferenceLocation unmarshals an instance of TrainingDataReferenceLocation from the specified map of raw messages.
func UnmarshalTrainingDataReferenceLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDataReferenceLocation)
	err = core.UnmarshalPrimitive(m, "schema_name", &obj.SchemaName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "table_name", &obj.TableName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_format", &obj.FileFormat)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_name", &obj.FileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "firstlineheader", &obj.Firstlineheader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "infer_schema", &obj.InferSchema)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dataset_id", &obj.DatasetID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "meta", &obj.Meta, UnmarshalFileTrainingDataReferenceOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_href", &obj.AssetHref)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_id", &obj.AssetID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserPreferencesDeleteOptions : The UserPreferencesDelete options.
type UserPreferencesDeleteOptions struct {
	// key in user preferences.
	UserPreferenceKey *string `json:"user_preference_key" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserPreferencesDeleteOptions : Instantiate UserPreferencesDeleteOptions
func (*WatsonOpenScaleV2) NewUserPreferencesDeleteOptions(userPreferenceKey string) *UserPreferencesDeleteOptions {
	return &UserPreferencesDeleteOptions{
		UserPreferenceKey: core.StringPtr(userPreferenceKey),
	}
}

// SetUserPreferenceKey : Allow user to set UserPreferenceKey
func (_options *UserPreferencesDeleteOptions) SetUserPreferenceKey(userPreferenceKey string) *UserPreferencesDeleteOptions {
	_options.UserPreferenceKey = core.StringPtr(userPreferenceKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UserPreferencesDeleteOptions) SetHeaders(param map[string]string) *UserPreferencesDeleteOptions {
	options.Headers = param
	return options
}

// UserPreferencesGetOptions : The UserPreferencesGet options.
type UserPreferencesGetOptions struct {
	// key in user preferences.
	UserPreferenceKey *string `json:"user_preference_key" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserPreferencesGetOptions : Instantiate UserPreferencesGetOptions
func (*WatsonOpenScaleV2) NewUserPreferencesGetOptions(userPreferenceKey string) *UserPreferencesGetOptions {
	return &UserPreferencesGetOptions{
		UserPreferenceKey: core.StringPtr(userPreferenceKey),
	}
}

// SetUserPreferenceKey : Allow user to set UserPreferenceKey
func (_options *UserPreferencesGetOptions) SetUserPreferenceKey(userPreferenceKey string) *UserPreferencesGetOptions {
	_options.UserPreferenceKey = core.StringPtr(userPreferenceKey)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UserPreferencesGetOptions) SetHeaders(param map[string]string) *UserPreferencesGetOptions {
	options.Headers = param
	return options
}

// UserPreferencesGetResponse : UserPreferencesGetResponse struct
// Models which "extend" this model:
// - UserPreferencesGetResponseUserPreferenceValueObject
// - UserPreferencesGetResponseUserPreferenceValueString
type UserPreferencesGetResponse struct {
	// response string value.
	Value *string `json:"value,omitempty"`
}

func (*UserPreferencesGetResponse) isaUserPreferencesGetResponse() bool {
	return true
}

type UserPreferencesGetResponseIntf interface {
	isaUserPreferencesGetResponse() bool
}

// UnmarshalUserPreferencesGetResponse unmarshals an instance of UserPreferencesGetResponse from the specified map of raw messages.
func UnmarshalUserPreferencesGetResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserPreferencesGetResponse)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserPreferencesListOptions : The UserPreferencesList options.
type UserPreferencesListOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserPreferencesListOptions : Instantiate UserPreferencesListOptions
func (*WatsonOpenScaleV2) NewUserPreferencesListOptions() *UserPreferencesListOptions {
	return &UserPreferencesListOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *UserPreferencesListOptions) SetHeaders(param map[string]string) *UserPreferencesListOptions {
	options.Headers = param
	return options
}

// UserPreferencesPatchOptions : The UserPreferencesPatch options.
type UserPreferencesPatchOptions struct {
	// Array of patch operations as defined in RFC 6902.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserPreferencesPatchOptions : Instantiate UserPreferencesPatchOptions
func (*WatsonOpenScaleV2) NewUserPreferencesPatchOptions(jsonPatchOperation []JSONPatchOperation) *UserPreferencesPatchOptions {
	return &UserPreferencesPatchOptions{
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *UserPreferencesPatchOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *UserPreferencesPatchOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UserPreferencesPatchOptions) SetHeaders(param map[string]string) *UserPreferencesPatchOptions {
	options.Headers = param
	return options
}

// UserPreferencesUpdateOptions : The UserPreferencesUpdate options.
type UserPreferencesUpdateOptions struct {
	// key in user preferences.
	UserPreferenceKey *string `json:"user_preference_key" validate:"required,ne="`

	UserPreferencesUpdateRequest UserPreferencesUpdateRequestIntf `json:"UserPreferencesUpdateRequest" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserPreferencesUpdateOptions : Instantiate UserPreferencesUpdateOptions
func (*WatsonOpenScaleV2) NewUserPreferencesUpdateOptions(userPreferenceKey string, userPreferencesUpdateRequest UserPreferencesUpdateRequestIntf) *UserPreferencesUpdateOptions {
	return &UserPreferencesUpdateOptions{
		UserPreferenceKey:            core.StringPtr(userPreferenceKey),
		UserPreferencesUpdateRequest: userPreferencesUpdateRequest,
	}
}

// SetUserPreferenceKey : Allow user to set UserPreferenceKey
func (_options *UserPreferencesUpdateOptions) SetUserPreferenceKey(userPreferenceKey string) *UserPreferencesUpdateOptions {
	_options.UserPreferenceKey = core.StringPtr(userPreferenceKey)
	return _options
}

// SetUserPreferencesUpdateRequest : Allow user to set UserPreferencesUpdateRequest
func (_options *UserPreferencesUpdateOptions) SetUserPreferencesUpdateRequest(userPreferencesUpdateRequest UserPreferencesUpdateRequestIntf) *UserPreferencesUpdateOptions {
	_options.UserPreferencesUpdateRequest = userPreferencesUpdateRequest
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UserPreferencesUpdateOptions) SetHeaders(param map[string]string) *UserPreferencesUpdateOptions {
	options.Headers = param
	return options
}

// UserPreferencesUpdateRequest : UserPreferencesUpdateRequest struct
// Models which "extend" this model:
// - UserPreferencesUpdateRequestUserPreferenceValueObject
// - UserPreferencesUpdateRequestUserPreferenceValueString
type UserPreferencesUpdateRequest struct {
	// response string value.
	Value *string `json:"value,omitempty"`
}

func (*UserPreferencesUpdateRequest) isaUserPreferencesUpdateRequest() bool {
	return true
}

type UserPreferencesUpdateRequestIntf interface {
	isaUserPreferencesUpdateRequest() bool
}

// UnmarshalUserPreferencesUpdateRequest unmarshals an instance of UserPreferencesUpdateRequest from the specified map of raw messages.
func UnmarshalUserPreferencesUpdateRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserPreferencesUpdateRequest)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AzureCredentials : AzureCredentials struct
// This model "extends" MLCredentials
type AzureCredentials struct {
	ClientID *string `json:"client_id,omitempty"`

	ClientSecret *string `json:"client_secret,omitempty"`

	Password *string `json:"password,omitempty"`

	SubscriptionID *string `json:"subscription_id,omitempty"`

	Tenant *string `json:"tenant,omitempty"`

	Token *string `json:"token,omitempty"`

	Username *string `json:"username,omitempty"`

	Workspaces []AzureWorkspaceCredentials `json:"workspaces,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

func (*AzureCredentials) isaMLCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of AzureCredentials
func (o *AzureCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of AzureCredentials
func (o *AzureCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of AzureCredentials
func (o *AzureCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of AzureCredentials
func (o *AzureCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of AzureCredentials
func (o *AzureCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.ClientID != nil {
		m["client_id"] = o.ClientID
	}
	if o.ClientSecret != nil {
		m["client_secret"] = o.ClientSecret
	}
	if o.Password != nil {
		m["password"] = o.Password
	}
	if o.SubscriptionID != nil {
		m["subscription_id"] = o.SubscriptionID
	}
	if o.Tenant != nil {
		m["tenant"] = o.Tenant
	}
	if o.Token != nil {
		m["token"] = o.Token
	}
	if o.Username != nil {
		m["username"] = o.Username
	}
	if o.Workspaces != nil {
		m["workspaces"] = o.Workspaces
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalAzureCredentials unmarshals an instance of AzureCredentials from the specified map of raw messages.
func UnmarshalAzureCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AzureCredentials)
	err = core.UnmarshalPrimitive(m, "client_id", &obj.ClientID)
	if err != nil {
		return
	}
	delete(m, "client_id")
	err = core.UnmarshalPrimitive(m, "client_secret", &obj.ClientSecret)
	if err != nil {
		return
	}
	delete(m, "client_secret")
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	delete(m, "password")
	err = core.UnmarshalPrimitive(m, "subscription_id", &obj.SubscriptionID)
	if err != nil {
		return
	}
	delete(m, "subscription_id")
	err = core.UnmarshalPrimitive(m, "tenant", &obj.Tenant)
	if err != nil {
		return
	}
	delete(m, "tenant")
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	delete(m, "token")
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	delete(m, "username")
	err = core.UnmarshalModel(m, "workspaces", &obj.Workspaces, UnmarshalAzureWorkspaceCredentials)
	if err != nil {
		return
	}
	delete(m, "workspaces")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// COSTrainingDataReferenceConnection : COSTrainingDataReferenceConnection struct
// This model "extends" TrainingDataReferenceConnection
type COSTrainingDataReferenceConnection struct {
	ApiKey *string `json:"api_key" validate:"required"`

	IamURL *string `json:"iam_url,omitempty"`

	ResourceInstanceID *string `json:"resource_instance_id" validate:"required"`

	URL *string `json:"url" validate:"required"`
}

// NewCOSTrainingDataReferenceConnection : Instantiate COSTrainingDataReferenceConnection (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewCOSTrainingDataReferenceConnection(apiKey string, resourceInstanceID string, url string) (_model *COSTrainingDataReferenceConnection, err error) {
	_model = &COSTrainingDataReferenceConnection{
		ApiKey:             core.StringPtr(apiKey),
		ResourceInstanceID: core.StringPtr(resourceInstanceID),
		URL:                core.StringPtr(url),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*COSTrainingDataReferenceConnection) isaTrainingDataReferenceConnection() bool {
	return true
}

// UnmarshalCOSTrainingDataReferenceConnection unmarshals an instance of COSTrainingDataReferenceConnection from the specified map of raw messages.
func UnmarshalCOSTrainingDataReferenceConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(COSTrainingDataReferenceConnection)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iam_url", &obj.IamURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_instance_id", &obj.ResourceInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// COSTrainingDataReferenceLocation : COS file location.
// This model "extends" TrainingDataReferenceLocation
type COSTrainingDataReferenceLocation struct {
	Bucket *string `json:"bucket" validate:"required"`

	FileFormat *string `json:"file_format,omitempty"`

	FileName *string `json:"file_name" validate:"required"`

	Firstlineheader *bool `json:"firstlineheader,omitempty"`

	InferSchema *string `json:"infer_schema,omitempty"`
}

// Constants associated with the COSTrainingDataReferenceLocation.FileFormat property.
const (
	COSTrainingDataReferenceLocation_FileFormat_Csv = "csv"
)

// NewCOSTrainingDataReferenceLocation : Instantiate COSTrainingDataReferenceLocation (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewCOSTrainingDataReferenceLocation(bucket string, fileName string) (_model *COSTrainingDataReferenceLocation, err error) {
	_model = &COSTrainingDataReferenceLocation{
		Bucket:   core.StringPtr(bucket),
		FileName: core.StringPtr(fileName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*COSTrainingDataReferenceLocation) isaTrainingDataReferenceLocation() bool {
	return true
}

// UnmarshalCOSTrainingDataReferenceLocation unmarshals an instance of COSTrainingDataReferenceLocation from the specified map of raw messages.
func UnmarshalCOSTrainingDataReferenceLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(COSTrainingDataReferenceLocation)
	err = core.UnmarshalPrimitive(m, "bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_format", &obj.FileFormat)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_name", &obj.FileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "firstlineheader", &obj.Firstlineheader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "infer_schema", &obj.InferSchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomCredentials : CustomCredentials struct
// This model "extends" MLCredentials
type CustomCredentials struct {
	Password *string `json:"password,omitempty"`

	URL *string `json:"url,omitempty"`

	Username *string `json:"username,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

func (*CustomCredentials) isaMLCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of CustomCredentials
func (o *CustomCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of CustomCredentials
func (o *CustomCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of CustomCredentials
func (o *CustomCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of CustomCredentials
func (o *CustomCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of CustomCredentials
func (o *CustomCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Password != nil {
		m["password"] = o.Password
	}
	if o.URL != nil {
		m["url"] = o.URL
	}
	if o.Username != nil {
		m["username"] = o.Username
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalCustomCredentials unmarshals an instance of CustomCredentials from the specified map of raw messages.
func UnmarshalCustomCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomCredentials)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	delete(m, "password")
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	delete(m, "url")
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	delete(m, "username")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DB2TrainingDataReferenceConnection : DB2TrainingDataReferenceConnection struct
// This model "extends" TrainingDataReferenceConnection
type DB2TrainingDataReferenceConnection struct {
	// DER-encoded certificate in Base64 encoding. The decoded content must be bound at the beginning by -----BEGIN
	// CERTIFICATE----- and at the end by -----END CERTIFICATE-----.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`

	ConnectionString *string `json:"connection_string,omitempty"`

	DatabaseName *string `json:"database_name" validate:"required"`

	Hostname *string `json:"hostname" validate:"required"`

	Password *string `json:"password" validate:"required"`

	Port *int64 `json:"port,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	Username *string `json:"username" validate:"required"`
}

// NewDB2TrainingDataReferenceConnection : Instantiate DB2TrainingDataReferenceConnection (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewDB2TrainingDataReferenceConnection(databaseName string, hostname string, password string, username string) (_model *DB2TrainingDataReferenceConnection, err error) {
	_model = &DB2TrainingDataReferenceConnection{
		DatabaseName: core.StringPtr(databaseName),
		Hostname:     core.StringPtr(hostname),
		Password:     core.StringPtr(password),
		Username:     core.StringPtr(username),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DB2TrainingDataReferenceConnection) isaTrainingDataReferenceConnection() bool {
	return true
}

// UnmarshalDB2TrainingDataReferenceConnection unmarshals an instance of DB2TrainingDataReferenceConnection from the specified map of raw messages.
func UnmarshalDB2TrainingDataReferenceConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DB2TrainingDataReferenceConnection)
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection_string", &obj.ConnectionString)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "database_name", &obj.DatabaseName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DB2TrainingDataReferenceLocation : DB2 table name.
// This model "extends" TrainingDataReferenceLocation
type DB2TrainingDataReferenceLocation struct {
	// name of the schema.
	SchemaName *string `json:"schema_name,omitempty"`

	// name of the table.
	TableName *string `json:"table_name" validate:"required"`
}

// NewDB2TrainingDataReferenceLocation : Instantiate DB2TrainingDataReferenceLocation (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewDB2TrainingDataReferenceLocation(tableName string) (_model *DB2TrainingDataReferenceLocation, err error) {
	_model = &DB2TrainingDataReferenceLocation{
		TableName: core.StringPtr(tableName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DB2TrainingDataReferenceLocation) isaTrainingDataReferenceLocation() bool {
	return true
}

// UnmarshalDB2TrainingDataReferenceLocation unmarshals an instance of DB2TrainingDataReferenceLocation from the specified map of raw messages.
func UnmarshalDB2TrainingDataReferenceLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DB2TrainingDataReferenceLocation)
	err = core.UnmarshalPrimitive(m, "schema_name", &obj.SchemaName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "table_name", &obj.TableName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatasetRecordsPayloadItemJsDictElem : Fields and values of the entity matches JSON object's fields and values.
// This model "extends" DatasetRecordsPayloadItem
type DatasetRecordsPayloadItemJsDictElem struct {
}

func (*DatasetRecordsPayloadItemJsDictElem) isaDatasetRecordsPayloadItem() bool {
	return true
}

// UnmarshalDatasetRecordsPayloadItemJsDictElem unmarshals an instance of DatasetRecordsPayloadItemJsDictElem from the specified map of raw messages.
func UnmarshalDatasetRecordsPayloadItemJsDictElem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatasetRecordsPayloadItemJsDictElem)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatasetRecordsPayloadItemJsList : JSON list (condensed) format.
// This model "extends" DatasetRecordsPayloadItem
type DatasetRecordsPayloadItemJsList struct {
	// Fields' names are listed in order in 'fields'.
	Fields []string `json:"fields" validate:"required"`

	// Rows organized per fields as described in 'fields'.
	Values [][]interface{} `json:"values" validate:"required"`
}

// NewDatasetRecordsPayloadItemJsList : Instantiate DatasetRecordsPayloadItemJsList (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewDatasetRecordsPayloadItemJsList(fields []string, values [][]interface{}) (_model *DatasetRecordsPayloadItemJsList, err error) {
	_model = &DatasetRecordsPayloadItemJsList{
		Fields: fields,
		Values: values,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DatasetRecordsPayloadItemJsList) isaDatasetRecordsPayloadItem() bool {
	return true
}

// UnmarshalDatasetRecordsPayloadItemJsList unmarshals an instance of DatasetRecordsPayloadItemJsList from the specified map of raw messages.
func UnmarshalDatasetRecordsPayloadItemJsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatasetRecordsPayloadItemJsList)
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatasetRecordsPayloadItemScoringPayloadRequest : DatasetRecordsPayloadItemScoringPayloadRequest struct
// This model "extends" DatasetRecordsPayloadItem
type DatasetRecordsPayloadItemScoringPayloadRequest struct {
	AssetRevision *string `json:"asset_revision,omitempty"`

	// This property is important for payload logging for unstructered data. There will be one record created in payload
	// logging table if this property is set to `false`. The `request.values` field and `response.values` field are treated
	// as one scoring request and one scoring response in such case. If this field is set to `true` then it might be
	// created more than one row in the payload logging table. The first dimension of the `request.values` and
	// `response.values` correponds to the separate entry in the payload logging table in such case.
	MultipleRecords *bool `json:"multiple_records,omitempty"`

	Request *ScoringPayloadRequestRequest `json:"request" validate:"required"`

	Response *ScoringPayloadRequestResponse `json:"response" validate:"required"`

	ResponseTime *float64 `json:"response_time,omitempty"`

	ScoringID *string `json:"scoring_id,omitempty"`

	ScoringTimestamp *string `json:"scoring_timestamp,omitempty"`
}

// NewDatasetRecordsPayloadItemScoringPayloadRequest : Instantiate DatasetRecordsPayloadItemScoringPayloadRequest (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewDatasetRecordsPayloadItemScoringPayloadRequest(request *ScoringPayloadRequestRequest, response *ScoringPayloadRequestResponse) (_model *DatasetRecordsPayloadItemScoringPayloadRequest, err error) {
	_model = &DatasetRecordsPayloadItemScoringPayloadRequest{
		Request:  request,
		Response: response,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DatasetRecordsPayloadItemScoringPayloadRequest) isaDatasetRecordsPayloadItem() bool {
	return true
}

// UnmarshalDatasetRecordsPayloadItemScoringPayloadRequest unmarshals an instance of DatasetRecordsPayloadItemScoringPayloadRequest from the specified map of raw messages.
func UnmarshalDatasetRecordsPayloadItemScoringPayloadRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatasetRecordsPayloadItemScoringPayloadRequest)
	err = core.UnmarshalPrimitive(m, "asset_revision", &obj.AssetRevision)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "multiple_records", &obj.MultipleRecords)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "request", &obj.Request, UnmarshalScoringPayloadRequestRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "response", &obj.Response, UnmarshalScoringPayloadRequestResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "response_time", &obj.ResponseTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scoring_id", &obj.ScoringID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scoring_timestamp", &obj.ScoringTimestamp)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DatasetTrainingDataReferenceLocation : WOS dataset id.
// This model "extends" TrainingDataReferenceLocation
type DatasetTrainingDataReferenceLocation struct {
	// Dataset id.
	DatasetID *string `json:"dataset_id" validate:"required"`

	// additional options for different types of training data references.
	Meta *FileTrainingDataReferenceOptions `json:"meta,omitempty"`
}

// NewDatasetTrainingDataReferenceLocation : Instantiate DatasetTrainingDataReferenceLocation (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewDatasetTrainingDataReferenceLocation(datasetID string) (_model *DatasetTrainingDataReferenceLocation, err error) {
	_model = &DatasetTrainingDataReferenceLocation{
		DatasetID: core.StringPtr(datasetID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*DatasetTrainingDataReferenceLocation) isaTrainingDataReferenceLocation() bool {
	return true
}

// UnmarshalDatasetTrainingDataReferenceLocation unmarshals an instance of DatasetTrainingDataReferenceLocation from the specified map of raw messages.
func UnmarshalDatasetTrainingDataReferenceLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DatasetTrainingDataReferenceLocation)
	err = core.UnmarshalPrimitive(m, "dataset_id", &obj.DatasetID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "meta", &obj.Meta, UnmarshalFileTrainingDataReferenceOptions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FileAssetTrainingDataReferenceLocation : File data asset reference location.
// This model "extends" TrainingDataReferenceLocation
type FileAssetTrainingDataReferenceLocation struct {
	// File data asset reference.
	AssetHref *string `json:"asset_href" validate:"required"`

	// File data asset id.
	AssetID *string `json:"asset_id,omitempty"`

	// additional options for different types of training data references.
	Meta *FileTrainingDataReferenceOptions `json:"meta,omitempty"`

	// Project id.
	ProjectID *string `json:"project_id,omitempty"`
}

// NewFileAssetTrainingDataReferenceLocation : Instantiate FileAssetTrainingDataReferenceLocation (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewFileAssetTrainingDataReferenceLocation(assetHref string) (_model *FileAssetTrainingDataReferenceLocation, err error) {
	_model = &FileAssetTrainingDataReferenceLocation{
		AssetHref: core.StringPtr(assetHref),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*FileAssetTrainingDataReferenceLocation) isaTrainingDataReferenceLocation() bool {
	return true
}

// UnmarshalFileAssetTrainingDataReferenceLocation unmarshals an instance of FileAssetTrainingDataReferenceLocation from the specified map of raw messages.
func UnmarshalFileAssetTrainingDataReferenceLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FileAssetTrainingDataReferenceLocation)
	err = core.UnmarshalPrimitive(m, "asset_href", &obj.AssetHref)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_id", &obj.AssetID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "meta", &obj.Meta, UnmarshalFileTrainingDataReferenceOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation : Contrastive explanation details.
// This model "extends" GetExplanationTaskResponseEntityExplanationsItem
type GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation struct {
	// Type of explanation.
	ExplanationType *string `json:"explanation_type,omitempty"`

	// These factors, if added, would cause the classification to change.
	PertinentNegative *GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative `json:"pertinent_negative,omitempty"`

	// These factors are sufficient evidence in themselves to yeild the given classification.
	PertinentPositive *GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive `json:"pertinent_positive,omitempty"`
}

func (*GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation) isaGetExplanationTaskResponseEntityExplanationsItem() bool {
	return true
}

// UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation unmarshals an instance of GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityExplanationsItemContrastiveExplanation)
	err = core.UnmarshalPrimitive(m, "explanation_type", &obj.ExplanationType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pertinent_negative", &obj.PertinentNegative, UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentNegative)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pertinent_positive", &obj.PertinentPositive, UnmarshalGetExplanationTaskResponseEntityExplanationsItemContrastiveExplanationPertinentPositive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetExplanationTaskResponseEntityExplanationsItemLimeExplanation : Lime explanation details.
// This model "extends" GetExplanationTaskResponseEntityExplanationsItem
type GetExplanationTaskResponseEntityExplanationsItemLimeExplanation struct {
	// Type of the explanation.
	ExplanationType *string `json:"explanation_type,omitempty"`

	// Lime explanations of predictions.
	Predictions []LimeExplanationPrediction `json:"predictions,omitempty"`
}

func (*GetExplanationTaskResponseEntityExplanationsItemLimeExplanation) isaGetExplanationTaskResponseEntityExplanationsItem() bool {
	return true
}

// UnmarshalGetExplanationTaskResponseEntityExplanationsItemLimeExplanation unmarshals an instance of GetExplanationTaskResponseEntityExplanationsItemLimeExplanation from the specified map of raw messages.
func UnmarshalGetExplanationTaskResponseEntityExplanationsItemLimeExplanation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetExplanationTaskResponseEntityExplanationsItemLimeExplanation)
	err = core.UnmarshalPrimitive(m, "explanation_type", &obj.ExplanationType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "predictions", &obj.Predictions, UnmarshalLimeExplanationPrediction)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrimaryStorageCredentialsLong : PrimaryStorageCredentialsLong struct
// This model "extends" PrimaryStorageCredentials
type PrimaryStorageCredentialsLong struct {
	// any additional properties to be included in connection url.
	AdditionalProperties map[string]interface{} `json:"additional_properties,omitempty"`

	// DER-encoded certificate in Base64 encoding. The decoded content must be bound at the beginning by -----BEGIN
	// CERTIFICATE----- and at the end by -----END CERTIFICATE-----.
	CertificateBase64 *string `json:"certificate_base64,omitempty"`

	Db *string `json:"db" validate:"required"`

	Hostname *string `json:"hostname" validate:"required"`

	Password *string `json:"password" validate:"required"`

	Port *int64 `json:"port,omitempty"`

	Ssl *bool `json:"ssl,omitempty"`

	// (postgresql only).
	Sslmode *string `json:"sslmode,omitempty"`

	Username *string `json:"username" validate:"required"`
}

// NewPrimaryStorageCredentialsLong : Instantiate PrimaryStorageCredentialsLong (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewPrimaryStorageCredentialsLong(db string, hostname string, password string, username string) (_model *PrimaryStorageCredentialsLong, err error) {
	_model = &PrimaryStorageCredentialsLong{
		Db:       core.StringPtr(db),
		Hostname: core.StringPtr(hostname),
		Password: core.StringPtr(password),
		Username: core.StringPtr(username),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*PrimaryStorageCredentialsLong) isaPrimaryStorageCredentials() bool {
	return true
}

// UnmarshalPrimaryStorageCredentialsLong unmarshals an instance of PrimaryStorageCredentialsLong from the specified map of raw messages.
func UnmarshalPrimaryStorageCredentialsLong(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrimaryStorageCredentialsLong)
	err = core.UnmarshalPrimitive(m, "additional_properties", &obj.AdditionalProperties)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "certificate_base64", &obj.CertificateBase64)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "db", &obj.Db)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hostname", &obj.Hostname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ssl", &obj.Ssl)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sslmode", &obj.Sslmode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PrimaryStorageCredentialsShort : PrimaryStorageCredentialsShort struct
// This model "extends" PrimaryStorageCredentials
type PrimaryStorageCredentialsShort struct {
	URI *string `json:"uri" validate:"required"`
}

// NewPrimaryStorageCredentialsShort : Instantiate PrimaryStorageCredentialsShort (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewPrimaryStorageCredentialsShort(uri string) (_model *PrimaryStorageCredentialsShort, err error) {
	_model = &PrimaryStorageCredentialsShort{
		URI: core.StringPtr(uri),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*PrimaryStorageCredentialsShort) isaPrimaryStorageCredentials() bool {
	return true
}

// UnmarshalPrimaryStorageCredentialsShort unmarshals an instance of PrimaryStorageCredentialsShort from the specified map of raw messages.
func UnmarshalPrimaryStorageCredentialsShort(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PrimaryStorageCredentialsShort)
	err = core.UnmarshalPrimitive(m, "uri", &obj.URI)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsListResponseDataRecordsResponseCollectionDict : dict format.
// This model "extends" RecordsListResponse
type RecordsListResponseDataRecordsResponseCollectionDict struct {
	Records []DataRecordResponse `json:"records" validate:"required"`

	// Number of all rows which satisfy the query. It is calculated and returned when include_total_count query param is
	// set to `true`.
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (*RecordsListResponseDataRecordsResponseCollectionDict) isaRecordsListResponse() bool {
	return true
}

// UnmarshalRecordsListResponseDataRecordsResponseCollectionDict unmarshals an instance of RecordsListResponseDataRecordsResponseCollectionDict from the specified map of raw messages.
func UnmarshalRecordsListResponseDataRecordsResponseCollectionDict(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordsListResponseDataRecordsResponseCollectionDict)
	err = core.UnmarshalModel(m, "records", &obj.Records, UnmarshalDataRecordResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsListResponseDataRecordsResponseCollectionList : list format.
// This model "extends" RecordsListResponse
type RecordsListResponseDataRecordsResponseCollectionList struct {
	Records []DataRecordResponseList `json:"records" validate:"required"`

	// Number of all rows which satisfy the query. It is calculated and returned when include_total_count query param is
	// set to `true`.
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (*RecordsListResponseDataRecordsResponseCollectionList) isaRecordsListResponse() bool {
	return true
}

// UnmarshalRecordsListResponseDataRecordsResponseCollectionList unmarshals an instance of RecordsListResponseDataRecordsResponseCollectionList from the specified map of raw messages.
func UnmarshalRecordsListResponseDataRecordsResponseCollectionList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordsListResponseDataRecordsResponseCollectionList)
	err = core.UnmarshalModel(m, "records", &obj.Records, UnmarshalDataRecordResponseList)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SPSSCredentials : SPSSCredentials struct
// This model "extends" MLCredentials
type SPSSCredentials struct {
	Password *string `json:"password" validate:"required"`

	URL *string `json:"url" validate:"required"`

	Username *string `json:"username" validate:"required"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewSPSSCredentials : Instantiate SPSSCredentials (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewSPSSCredentials(password string, url string, username string) (_model *SPSSCredentials, err error) {
	_model = &SPSSCredentials{
		Password: core.StringPtr(password),
		URL:      core.StringPtr(url),
		Username: core.StringPtr(username),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*SPSSCredentials) isaMLCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of SPSSCredentials
func (o *SPSSCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of SPSSCredentials
func (o *SPSSCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of SPSSCredentials
func (o *SPSSCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of SPSSCredentials
func (o *SPSSCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of SPSSCredentials
func (o *SPSSCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Password != nil {
		m["password"] = o.Password
	}
	if o.URL != nil {
		m["url"] = o.URL
	}
	if o.Username != nil {
		m["username"] = o.Username
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalSPSSCredentials unmarshals an instance of SPSSCredentials from the specified map of raw messages.
func UnmarshalSPSSCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SPSSCredentials)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	delete(m, "password")
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	delete(m, "url")
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	delete(m, "username")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SageMakerCredentials : SageMakerCredentials struct
// This model "extends" MLCredentials
type SageMakerCredentials struct {
	AccessKeyID *string `json:"access_key_id" validate:"required"`

	Region *string `json:"region,omitempty"`

	SecretAccessKey *string `json:"secret_access_key" validate:"required"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewSageMakerCredentials : Instantiate SageMakerCredentials (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewSageMakerCredentials(accessKeyID string, secretAccessKey string) (_model *SageMakerCredentials, err error) {
	_model = &SageMakerCredentials{
		AccessKeyID:     core.StringPtr(accessKeyID),
		SecretAccessKey: core.StringPtr(secretAccessKey),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*SageMakerCredentials) isaMLCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of SageMakerCredentials
func (o *SageMakerCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of SageMakerCredentials
func (o *SageMakerCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of SageMakerCredentials
func (o *SageMakerCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of SageMakerCredentials
func (o *SageMakerCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of SageMakerCredentials
func (o *SageMakerCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.AccessKeyID != nil {
		m["access_key_id"] = o.AccessKeyID
	}
	if o.Region != nil {
		m["region"] = o.Region
	}
	if o.SecretAccessKey != nil {
		m["secret_access_key"] = o.SecretAccessKey
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalSageMakerCredentials unmarshals an instance of SageMakerCredentials from the specified map of raw messages.
func UnmarshalSageMakerCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SageMakerCredentials)
	err = core.UnmarshalPrimitive(m, "access_key_id", &obj.AccessKeyID)
	if err != nil {
		return
	}
	delete(m, "access_key_id")
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	delete(m, "region")
	err = core.UnmarshalPrimitive(m, "secret_access_key", &obj.SecretAccessKey)
	if err != nil {
		return
	}
	delete(m, "secret_access_key")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringEndpointCredentialsAzureScoringEndpointCredentials : ScoringEndpointCredentialsAzureScoringEndpointCredentials struct
// This model "extends" ScoringEndpointCredentials
type ScoringEndpointCredentialsAzureScoringEndpointCredentials struct {
	Token *string `json:"token" validate:"required"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewScoringEndpointCredentialsAzureScoringEndpointCredentials : Instantiate ScoringEndpointCredentialsAzureScoringEndpointCredentials (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewScoringEndpointCredentialsAzureScoringEndpointCredentials(token string) (_model *ScoringEndpointCredentialsAzureScoringEndpointCredentials, err error) {
	_model = &ScoringEndpointCredentialsAzureScoringEndpointCredentials{
		Token: core.StringPtr(token),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ScoringEndpointCredentialsAzureScoringEndpointCredentials) isaScoringEndpointCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of ScoringEndpointCredentialsAzureScoringEndpointCredentials
func (o *ScoringEndpointCredentialsAzureScoringEndpointCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of ScoringEndpointCredentialsAzureScoringEndpointCredentials
func (o *ScoringEndpointCredentialsAzureScoringEndpointCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of ScoringEndpointCredentialsAzureScoringEndpointCredentials
func (o *ScoringEndpointCredentialsAzureScoringEndpointCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of ScoringEndpointCredentialsAzureScoringEndpointCredentials
func (o *ScoringEndpointCredentialsAzureScoringEndpointCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of ScoringEndpointCredentialsAzureScoringEndpointCredentials
func (o *ScoringEndpointCredentialsAzureScoringEndpointCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Token != nil {
		m["token"] = o.Token
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalScoringEndpointCredentialsAzureScoringEndpointCredentials unmarshals an instance of ScoringEndpointCredentialsAzureScoringEndpointCredentials from the specified map of raw messages.
func UnmarshalScoringEndpointCredentialsAzureScoringEndpointCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringEndpointCredentialsAzureScoringEndpointCredentials)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	delete(m, "token")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UnknownCredentials : Unknown service provider credentials.
// This model "extends" MLCredentials
type UnknownCredentials struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

func (*UnknownCredentials) isaMLCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of UnknownCredentials
func (o *UnknownCredentials) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of UnknownCredentials
func (o *UnknownCredentials) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of UnknownCredentials
func (o *UnknownCredentials) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of UnknownCredentials
func (o *UnknownCredentials) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of UnknownCredentials
func (o *UnknownCredentials) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalUnknownCredentials unmarshals an instance of UnknownCredentials from the specified map of raw messages.
func UnmarshalUnknownCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UnknownCredentials)
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserPreferencesGetResponseUserPreferenceValueObject : user preference object value.
// This model "extends" UserPreferencesGetResponse
type UserPreferencesGetResponseUserPreferenceValueObject struct {
}

func (*UserPreferencesGetResponseUserPreferenceValueObject) isaUserPreferencesGetResponse() bool {
	return true
}

// UnmarshalUserPreferencesGetResponseUserPreferenceValueObject unmarshals an instance of UserPreferencesGetResponseUserPreferenceValueObject from the specified map of raw messages.
func UnmarshalUserPreferencesGetResponseUserPreferenceValueObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserPreferencesGetResponseUserPreferenceValueObject)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserPreferencesGetResponseUserPreferenceValueString : user preference string value - response string wrapped in an object.
// This model "extends" UserPreferencesGetResponse
type UserPreferencesGetResponseUserPreferenceValueString struct {
	// response string value.
	Value *string `json:"value,omitempty"`
}

func (*UserPreferencesGetResponseUserPreferenceValueString) isaUserPreferencesGetResponse() bool {
	return true
}

// UnmarshalUserPreferencesGetResponseUserPreferenceValueString unmarshals an instance of UserPreferencesGetResponseUserPreferenceValueString from the specified map of raw messages.
func UnmarshalUserPreferencesGetResponseUserPreferenceValueString(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserPreferencesGetResponseUserPreferenceValueString)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserPreferencesUpdateRequestUserPreferenceValueObject : user preference object value.
// This model "extends" UserPreferencesUpdateRequest
type UserPreferencesUpdateRequestUserPreferenceValueObject struct {
}

func (*UserPreferencesUpdateRequestUserPreferenceValueObject) isaUserPreferencesUpdateRequest() bool {
	return true
}

// UnmarshalUserPreferencesUpdateRequestUserPreferenceValueObject unmarshals an instance of UserPreferencesUpdateRequestUserPreferenceValueObject from the specified map of raw messages.
func UnmarshalUserPreferencesUpdateRequestUserPreferenceValueObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserPreferencesUpdateRequestUserPreferenceValueObject)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserPreferencesUpdateRequestUserPreferenceValueString : user preference string value - response string wrapped in an object.
// This model "extends" UserPreferencesUpdateRequest
type UserPreferencesUpdateRequestUserPreferenceValueString struct {
	// response string value.
	Value *string `json:"value,omitempty"`
}

func (*UserPreferencesUpdateRequestUserPreferenceValueString) isaUserPreferencesUpdateRequest() bool {
	return true
}

// UnmarshalUserPreferencesUpdateRequestUserPreferenceValueString unmarshals an instance of UserPreferencesUpdateRequestUserPreferenceValueString from the specified map of raw messages.
func UnmarshalUserPreferencesUpdateRequestUserPreferenceValueString(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserPreferencesUpdateRequestUserPreferenceValueString)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WMLCredentialsCP4D : WMLCredentialsCP4D struct
// This model "extends" MLCredentials
type WMLCredentialsCP4D struct {
	InstanceID *string `json:"instance_id,omitempty"`

	Password *string `json:"password,omitempty"`

	Token *string `json:"token,omitempty"`

	URL *string `json:"url,omitempty"`

	Username *string `json:"username,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

func (*WMLCredentialsCP4D) isaMLCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of WMLCredentialsCP4D
func (o *WMLCredentialsCP4D) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of WMLCredentialsCP4D
func (o *WMLCredentialsCP4D) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of WMLCredentialsCP4D
func (o *WMLCredentialsCP4D) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of WMLCredentialsCP4D
func (o *WMLCredentialsCP4D) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of WMLCredentialsCP4D
func (o *WMLCredentialsCP4D) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.InstanceID != nil {
		m["instance_id"] = o.InstanceID
	}
	if o.Password != nil {
		m["password"] = o.Password
	}
	if o.Token != nil {
		m["token"] = o.Token
	}
	if o.URL != nil {
		m["url"] = o.URL
	}
	if o.Username != nil {
		m["username"] = o.Username
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalWMLCredentialsCP4D unmarshals an instance of WMLCredentialsCP4D from the specified map of raw messages.
func UnmarshalWMLCredentialsCP4D(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WMLCredentialsCP4D)
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	delete(m, "instance_id")
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	delete(m, "password")
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	delete(m, "token")
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	delete(m, "url")
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	delete(m, "username")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WMLCredentialsCloud : WMLCredentialsCloud struct
// This model "extends" MLCredentials
type WMLCredentialsCloud struct {
	Apikey *string `json:"apikey,omitempty"`

	InstanceID *string `json:"instance_id" validate:"required"`

	Token *string `json:"token,omitempty"`

	URL *string `json:"url" validate:"required"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewWMLCredentialsCloud : Instantiate WMLCredentialsCloud (Generic Model Constructor)
func (*WatsonOpenScaleV2) NewWMLCredentialsCloud(instanceID string, url string) (_model *WMLCredentialsCloud, err error) {
	_model = &WMLCredentialsCloud{
		InstanceID: core.StringPtr(instanceID),
		URL:        core.StringPtr(url),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*WMLCredentialsCloud) isaMLCredentials() bool {
	return true
}

// SetProperty allows the user to set an arbitrary property on an instance of WMLCredentialsCloud
func (o *WMLCredentialsCloud) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of WMLCredentialsCloud
func (o *WMLCredentialsCloud) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of WMLCredentialsCloud
func (o *WMLCredentialsCloud) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of WMLCredentialsCloud
func (o *WMLCredentialsCloud) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of WMLCredentialsCloud
func (o *WMLCredentialsCloud) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Apikey != nil {
		m["apikey"] = o.Apikey
	}
	if o.InstanceID != nil {
		m["instance_id"] = o.InstanceID
	}
	if o.Token != nil {
		m["token"] = o.Token
	}
	if o.URL != nil {
		m["url"] = o.URL
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalWMLCredentialsCloud unmarshals an instance of WMLCredentialsCloud from the specified map of raw messages.
func UnmarshalWMLCredentialsCloud(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WMLCredentialsCloud)
	err = core.UnmarshalPrimitive(m, "apikey", &obj.Apikey)
	if err != nil {
		return
	}
	delete(m, "apikey")
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	delete(m, "instance_id")
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	delete(m, "token")
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	delete(m, "url")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
