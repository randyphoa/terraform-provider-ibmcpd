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

// Package watsonmachinelearningv4 : Operations and models for the WatsonMachineLearningV4 service
package watsonmachinelearningv4

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

// WatsonMachineLearningV4 : ## Introduction
//
// Step by step instructions on how to use IBM Watson Machine Learning (WML) can be found
// [here](https://dataplatform.cloud.ibm.com/docs/content/wsj/analyze-data/ml-overview.html?context=analytics).
//
// ## Authentication
//
// In order to authenticate in WML, you need to generate an `IAM token`.
//
// To start working with the API, create an IAM token as described below
// * [Watson Machine Learning
// authentication](https://dataplatform.cloud.ibm.com/docs/content/wsj/analyze-data/ml-authentication.html)
// * [Authenticating with IAM tokens](https://cloud.ibm.com/docs/services/watson?topic=watson-iam#iam)
// * [Create an IAM access token for a user or service
// ID](https://cloud.ibm.com/apidocs/iam-identity-token-api#create-an-iam-access-token-for-a-user-or-service-i)
//
// The obtained `IAM token` needs to be prepended with the word `Bearer`, and it needs to be passed in the
// `Authorization` header for API calls.
//
// Example of an API request with a 'Bearer' access token:
//
// > curl https://us-south.ml.cloud.ibm.com/ml/v4/models
// >   -H "Authorization: Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hb..."
//
// ## Versioning
//
// API requests require a version parameter that takes a date in the format `version=YYYY-MM-DD`. When there is a change
// to the API in a [backwards-incompatible way](https://github.com/watson-developer-cloud/api-guidelines/#versioning),
// there will be a new version date published.
//
// The service uses the API version for the date you specify or the most recent version before that date. It is not
// recommended to default to the current date. Instead, specify a date that matches a version that is compatible with
// your app and do not change it until your app is ready for a later version.
//
// ## Error handling
//
// The IBM Watson Machine Learning uses standard HTTP response codes to indicate indicate if a method completed
// successfully. A `200`, `201` or `204` HTTP response always indicates success. HTTP response codes with the format
// `4xx` indicate a failure. A `500` HTTP response code usually indicates an internal system error that cannot be
// resolved by the user.
//
// Note that all un-recognized query parameters to an API call may be silently ignored.
//
// API Version: 4.0.0
type WatsonMachineLearningV4 struct {
	Service *core.BaseService

	// The version date for the API of the form `YYYY-MM-DD`.
	Version *strfmt.Date
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.ml.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "watson_machine_learning"

// WatsonMachineLearningV4Options : Service options
type WatsonMachineLearningV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// The version date for the API of the form `YYYY-MM-DD`.
	Version *strfmt.Date `validate:"required"`
}

// NewWatsonMachineLearningV4UsingExternalConfig : constructs an instance of WatsonMachineLearningV4 with passed in options and external configuration.
func NewWatsonMachineLearningV4UsingExternalConfig(options *WatsonMachineLearningV4Options) (watsonMachineLearning *WatsonMachineLearningV4, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	watsonMachineLearning, err = NewWatsonMachineLearningV4(options)
	if err != nil {
		return
	}

	err = watsonMachineLearning.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = watsonMachineLearning.Service.SetServiceURL(options.URL)
	}
	return
}

// NewWatsonMachineLearningV4 : constructs an instance of WatsonMachineLearningV4 with passed in options.
func NewWatsonMachineLearningV4(options *WatsonMachineLearningV4Options) (service *WatsonMachineLearningV4, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
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

	service = &WatsonMachineLearningV4{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "watsonMachineLearning" suitable for processing requests.
func (watsonMachineLearning *WatsonMachineLearningV4) Clone() *WatsonMachineLearningV4 {
	if core.IsNil(watsonMachineLearning) {
		return nil
	}
	clone := *watsonMachineLearning
	clone.Service = watsonMachineLearning.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (watsonMachineLearning *WatsonMachineLearningV4) SetServiceURL(url string) error {
	return watsonMachineLearning.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (watsonMachineLearning *WatsonMachineLearningV4) GetServiceURL() string {
	return watsonMachineLearning.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (watsonMachineLearning *WatsonMachineLearningV4) SetDefaultHeaders(headers http.Header) {
	watsonMachineLearning.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (watsonMachineLearning *WatsonMachineLearningV4) SetEnableGzipCompression(enableGzip bool) {
	watsonMachineLearning.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (watsonMachineLearning *WatsonMachineLearningV4) GetEnableGzipCompression() bool {
	return watsonMachineLearning.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (watsonMachineLearning *WatsonMachineLearningV4) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	watsonMachineLearning.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (watsonMachineLearning *WatsonMachineLearningV4) DisableRetries() {
	watsonMachineLearning.Service.DisableRetries()
}

// DeploymentsCreate : Create a new WML deployment
// Create a new deployment, the parameters specifying the deployment type are `online`, `r_shiny` and `batch`. These
// parameters are mutually exclusive, specify only one of these when creating a deployment.
//
// Use `hybrid_pipeline_hardware_specs` only when creating a `batch` deployment job of a hybrid pipeline in order to
// specify compute configuration for each pipeline node. For all other `batch` deployment cases use `hardware_spec` to
// specify compute configuration. The `hybrid_pipeline_hardware_specs` and
// `hardware_spec` are mutually exclusive, specify only one of these when creating a deployment.
//
// For `batch` deployments, `hardware_spec.num_nodes` parameter is not currently supported.
//
// For `online` deployments, `hardware_spec` cannot be specified at the time of creation,
// `hardware_spec.num_nodes` parameter is not supported as part of
// `POST /ml/v4/deployments` API request, but it can be updated using `PATCH /ml/v4/deployments/<deployment id>`.
//
// For `online` and `r_shiny` deployments, `serving_name` can be provided in
// `online.parameters` or `r_shiny.parameters`. The serving name can have the characters `[a-z,0-9,_]` and must not be
// more than 36 characters. The `serving_name` can be used in the prediction URL in place of the `deployment_id`.
//
// See the documentation [supported
// frameworks](https://www.ibm.com/docs/en/cloud-paks/cp-data/4.0?topic=specifications-supported-deployment-frameworks)
// for details about which frameworks can be used in a deployment.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsCreate(deploymentsCreateOptions *DeploymentsCreateOptions) (result *DeploymentResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentsCreateWithContext(context.Background(), deploymentsCreateOptions)
}

// DeploymentsCreateWithContext is an alternate form of the DeploymentsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsCreateWithContext(ctx context.Context, deploymentsCreateOptions *DeploymentsCreateOptions) (result *DeploymentResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentsCreateOptions, "deploymentsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentsCreateOptions, "deploymentsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if deploymentsCreateOptions.SpaceID != nil {
		body["space_id"] = deploymentsCreateOptions.SpaceID
	}
	if deploymentsCreateOptions.Tags != nil {
		body["tags"] = deploymentsCreateOptions.Tags
	}
	if deploymentsCreateOptions.Name != nil {
		body["name"] = deploymentsCreateOptions.Name
	}
	if deploymentsCreateOptions.Description != nil {
		body["description"] = deploymentsCreateOptions.Description
	}
	if deploymentsCreateOptions.Custom != nil {
		body["custom"] = deploymentsCreateOptions.Custom
	}
	if deploymentsCreateOptions.Asset != nil {
		body["asset"] = deploymentsCreateOptions.Asset
	}
	if deploymentsCreateOptions.HardwareSpec != nil {
		body["hardware_spec"] = deploymentsCreateOptions.HardwareSpec
	}
	if deploymentsCreateOptions.HybridPipelineHardwareSpecs != nil {
		body["hybrid_pipeline_hardware_specs"] = deploymentsCreateOptions.HybridPipelineHardwareSpecs
	}
	if deploymentsCreateOptions.Online != nil {
		body["online"] = deploymentsCreateOptions.Online
	}
	if deploymentsCreateOptions.Batch != nil {
		body["batch"] = deploymentsCreateOptions.Batch
	}
	if deploymentsCreateOptions.RShiny != nil {
		body["r_shiny"] = deploymentsCreateOptions.RShiny
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeploymentResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentsList : Retrieve the deployments
// Retrieve the list of deployments for the specified space.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsList(deploymentsListOptions *DeploymentsListOptions) (result *DeploymentResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentsListWithContext(context.Background(), deploymentsListOptions)
}

// DeploymentsListWithContext is an alternate form of the DeploymentsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsListWithContext(ctx context.Context, deploymentsListOptions *DeploymentsListOptions) (result *DeploymentResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deploymentsListOptions, "deploymentsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if deploymentsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*deploymentsListOptions.SpaceID))
	}
	if deploymentsListOptions.ServingName != nil {
		builder.AddQuery("serving_name", fmt.Sprint(*deploymentsListOptions.ServingName))
	}
	if deploymentsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*deploymentsListOptions.TagValue))
	}
	if deploymentsListOptions.AssetID != nil {
		builder.AddQuery("asset_id", fmt.Sprint(*deploymentsListOptions.AssetID))
	}
	if deploymentsListOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*deploymentsListOptions.Name))
	}
	if deploymentsListOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*deploymentsListOptions.Type))
	}
	if deploymentsListOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*deploymentsListOptions.State))
	}
	if deploymentsListOptions.Stats != nil {
		builder.AddQuery("stats", fmt.Sprint(*deploymentsListOptions.Stats))
	}
	if deploymentsListOptions.Conflict != nil {
		builder.AddQuery("conflict", fmt.Sprint(*deploymentsListOptions.Conflict))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeploymentResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentsGet : Retrieve the deployment details
// Retrieve the deployment details with the specified identifier.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsGet(deploymentsGetOptions *DeploymentsGetOptions) (result *DeploymentResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentsGetWithContext(context.Background(), deploymentsGetOptions)
}

// DeploymentsGetWithContext is an alternate form of the DeploymentsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsGetWithContext(ctx context.Context, deploymentsGetOptions *DeploymentsGetOptions) (result *DeploymentResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentsGetOptions, "deploymentsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentsGetOptions, "deploymentsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"deployment_id": *deploymentsGetOptions.DeploymentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployments/{deployment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentsGetOptions.SpaceID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeploymentResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentsDelete : Delete the deployment
// Delete the deployment with the specified identifier.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsDelete(deploymentsDeleteOptions *DeploymentsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentsDeleteWithContext(context.Background(), deploymentsDeleteOptions)
}

// DeploymentsDeleteWithContext is an alternate form of the DeploymentsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsDeleteWithContext(ctx context.Context, deploymentsDeleteOptions *DeploymentsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentsDeleteOptions, "deploymentsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentsDeleteOptions, "deploymentsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"deployment_id": *deploymentsDeleteOptions.DeploymentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployments/{deployment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentsDeleteOptions.SpaceID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// DeploymentsUpdate : Update the deployment metadata
// Update the deployment metadata. The following parameters of deployment metadata are supported for the patch
// operation.
//
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`
// - `/hardware_spec`
// - `/hybrid_pipeline_hardware_specs`
// - `/asset`
// - `/online/parameters`
// - `/r_shiny/authentication`
// - `/r_shiny/parameters/code_package` (only the `path` field in `code_package`)
//
// In case of online deployments, using PATCH operation of `/ml/v4/deployments`, users can update the number of copies
// of an online deployment. Users can specify the desired value of number of copies in `hardware_spec.num_nodes`
// parameter. As `hardware_spec.name` or `hardware_spec.id` is mandatory for `hardware_spec` schema, a valid value such
// as `XS`, `S` must be specified for `hardware_spec.name` parameter as part of PATCH request. Alternatively, users can
// also specify a valid ID of a hardware specification in `hardware_spec.id` parameter. However, changes related to
// `hardware_spec.name` or `hardware_spec.id` specified in PATCH operation will not be applied for online deployments.
// <br /> In case of batch deployments, using PATCH operation of `/ml/v4/deployments`, users can update the hardware
// specification so that subsequent batch deployment jobs can make use of the updated compute configurations. To update
// the compute configuration, users must specify a valid value for either `hardware_spec.name` or `hardware_spec.id` of
// the hardware specification that suits their requirement. In the batch deployment context, `hardware_spec.num_nodes`
// parameter is not currently supported.
// <br /> When 'asset' is patched with id/rev:
//
// - Deployment with the patched id/rev is started.
// - If the deployment is asynchronous, 202 response code will be returned and one can poll the deployment status
// thereafter.
// - If the deployment is synchronous, 200 response code will be returned with patched deployment response.
// - If any failures, deployment will be reverted back to the previous id/rev and the failure message will be captured
// in 'failures' field in the response.
//
// In the case of an online deployment, the PATCH operation with path specified as `/online/parameters` can be used to
// update the `serving_name`. In the case of a Shiny deployment, the PATCH operation with path specified as
// `/r_shiny/parameters` can be used to update the `serving_name`.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsUpdate(deploymentsUpdateOptions *DeploymentsUpdateOptions) (result *DeploymentResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentsUpdateWithContext(context.Background(), deploymentsUpdateOptions)
}

// DeploymentsUpdateWithContext is an alternate form of the DeploymentsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsUpdateWithContext(ctx context.Context, deploymentsUpdateOptions *DeploymentsUpdateOptions) (result *DeploymentResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentsUpdateOptions, "deploymentsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentsUpdateOptions, "deploymentsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"deployment_id": *deploymentsUpdateOptions.DeploymentID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployments/{deployment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentsUpdateOptions.SpaceID))

	_, err = builder.SetBodyContentJSON(deploymentsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeploymentResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentsComputePredictions : Execute a synchronous deployment prediction
// Execute a synchronous prediction for the deployment with the specified identifier. If a `serving_name` is used then
// it must match the `serving_name` that is returned in the `serving_urls`.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsComputePredictions(deploymentsComputePredictionsOptions *DeploymentsComputePredictionsOptions) (result *SyncScoringDataResults, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentsComputePredictionsWithContext(context.Background(), deploymentsComputePredictionsOptions)
}

// DeploymentsComputePredictionsWithContext is an alternate form of the DeploymentsComputePredictions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentsComputePredictionsWithContext(ctx context.Context, deploymentsComputePredictionsOptions *DeploymentsComputePredictionsOptions) (result *SyncScoringDataResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentsComputePredictionsOptions, "deploymentsComputePredictionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentsComputePredictionsOptions, "deploymentsComputePredictionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"deployment_id": *deploymentsComputePredictionsOptions.DeploymentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployments/{deployment_id}/predictions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentsComputePredictionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentsComputePredictions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if deploymentsComputePredictionsOptions.InputData != nil {
		body["input_data"] = deploymentsComputePredictionsOptions.InputData
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSyncScoringDataResults)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobsList : Retrieve the deployment jobs
// Retrieve the status of the current jobs. The system will apply a max limit of jobs retained by the system as we
// cannot accumulate an infinite number of jobs. Only most recent 300 jobs (system configurable) will be preserved.
// Older jobs will be purged by the system.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsList(deploymentJobsListOptions *DeploymentJobsListOptions) (result *JobsResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobsListWithContext(context.Background(), deploymentJobsListOptions)
}

// DeploymentJobsListWithContext is an alternate form of the DeploymentJobsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsListWithContext(ctx context.Context, deploymentJobsListOptions *DeploymentJobsListOptions) (result *JobsResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobsListOptions, "deploymentJobsListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobsListOptions, "deploymentJobsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_jobs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobsListOptions.SpaceID))
	if deploymentJobsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*deploymentJobsListOptions.TagValue))
	}
	if deploymentJobsListOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*deploymentJobsListOptions.State))
	}
	if deploymentJobsListOptions.DeploymentID != nil {
		builder.AddQuery("deployment_id", fmt.Sprint(*deploymentJobsListOptions.DeploymentID))
	}
	if deploymentJobsListOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*deploymentJobsListOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobsResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobsCreate : Start an asynchronous deployment job
// Start a deployment job asynchronously. This can perform batch scoring, streaming, or other types of batch operations,
// such as solving a Decision Optimization problem. Depending on the `version` date passed, the `platform_jobs` section
// in the response may or may not be populated. Use the GET call to retrieve the deployment job, this GET call will
// eventually populate the `platform_jobs` section. Refer to the `version date` description for more details.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsCreate(deploymentJobsCreateOptions *DeploymentJobsCreateOptions) (result *JobsResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobsCreateWithContext(context.Background(), deploymentJobsCreateOptions)
}

// DeploymentJobsCreateWithContext is an alternate form of the DeploymentJobsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsCreateWithContext(ctx context.Context, deploymentJobsCreateOptions *DeploymentJobsCreateOptions) (result *JobsResource, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deploymentJobsCreateOptions, "deploymentJobsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_jobs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if deploymentJobsCreateOptions.Retention != nil {
		builder.AddQuery("retention", fmt.Sprint(*deploymentJobsCreateOptions.Retention))
	}

	body := make(map[string]interface{})
	if deploymentJobsCreateOptions.SpaceID != nil {
		body["space_id"] = deploymentJobsCreateOptions.SpaceID
	}
	if deploymentJobsCreateOptions.Name != nil {
		body["name"] = deploymentJobsCreateOptions.Name
	}
	if deploymentJobsCreateOptions.Description != nil {
		body["description"] = deploymentJobsCreateOptions.Description
	}
	if deploymentJobsCreateOptions.Tags != nil {
		body["tags"] = deploymentJobsCreateOptions.Tags
	}
	if deploymentJobsCreateOptions.Deployment != nil {
		body["deployment"] = deploymentJobsCreateOptions.Deployment
	}
	if deploymentJobsCreateOptions.Custom != nil {
		body["custom"] = deploymentJobsCreateOptions.Custom
	}
	if deploymentJobsCreateOptions.HardwareSpec != nil {
		body["hardware_spec"] = deploymentJobsCreateOptions.HardwareSpec
	}
	if deploymentJobsCreateOptions.HybridPipelineHardwareSpecs != nil {
		body["hybrid_pipeline_hardware_specs"] = deploymentJobsCreateOptions.HybridPipelineHardwareSpecs
	}
	if deploymentJobsCreateOptions.Scoring != nil {
		body["scoring"] = deploymentJobsCreateOptions.Scoring
	}
	if deploymentJobsCreateOptions.DecisionOptimization != nil {
		body["decision_optimization"] = deploymentJobsCreateOptions.DecisionOptimization
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobsResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobsDelete : Cancel the deployment job
// Cancel the specified deployment job.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsDelete(deploymentJobsDeleteOptions *DeploymentJobsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobsDeleteWithContext(context.Background(), deploymentJobsDeleteOptions)
}

// DeploymentJobsDeleteWithContext is an alternate form of the DeploymentJobsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsDeleteWithContext(ctx context.Context, deploymentJobsDeleteOptions *DeploymentJobsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobsDeleteOptions, "deploymentJobsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobsDeleteOptions, "deploymentJobsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_id": *deploymentJobsDeleteOptions.JobID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_jobs/{job_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobsDeleteOptions.SpaceID))
	if deploymentJobsDeleteOptions.HardDelete != nil {
		builder.AddQuery("hard_delete", fmt.Sprint(*deploymentJobsDeleteOptions.HardDelete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// DeploymentJobsGet : Retrieve the deployment job
// Retrieve the deployment job. The predicted data bound to this job_id is going to be kept around for a limited time
// based on the service configuration.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsGet(deploymentJobsGetOptions *DeploymentJobsGetOptions) (result *JobsResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobsGetWithContext(context.Background(), deploymentJobsGetOptions)
}

// DeploymentJobsGetWithContext is an alternate form of the DeploymentJobsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobsGetWithContext(ctx context.Context, deploymentJobsGetOptions *DeploymentJobsGetOptions) (result *JobsResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobsGetOptions, "deploymentJobsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobsGetOptions, "deploymentJobsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_id": *deploymentJobsGetOptions.JobID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_jobs/{job_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobsGetOptions.SpaceID))
	if deploymentJobsGetOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*deploymentJobsGetOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobsResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobDefinitionsCreate : Create a new deployment job definition
// Create a new deployment job definition with the given payload. A deployment job definition represents the deployment
// metadata information in order to create a batch job in WML. This contains the same metadata used by the
// /ml/v4/deployment_jobs endpoint. This means that when submitting batch deployment jobs the user can either provide
// the job definition inline or reference a job definition in a query parameter.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsCreate(deploymentJobDefinitionsCreateOptions *DeploymentJobDefinitionsCreateOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobDefinitionsCreateWithContext(context.Background(), deploymentJobDefinitionsCreateOptions)
}

// DeploymentJobDefinitionsCreateWithContext is an alternate form of the DeploymentJobDefinitionsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsCreateWithContext(ctx context.Context, deploymentJobDefinitionsCreateOptions *DeploymentJobDefinitionsCreateOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobDefinitionsCreateOptions, "deploymentJobDefinitionsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobDefinitionsCreateOptions, "deploymentJobDefinitionsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_job_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobDefinitionsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobDefinitionsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if deploymentJobDefinitionsCreateOptions.SpaceID != nil {
		body["space_id"] = deploymentJobDefinitionsCreateOptions.SpaceID
	}
	if deploymentJobDefinitionsCreateOptions.Name != nil {
		body["name"] = deploymentJobDefinitionsCreateOptions.Name
	}
	if deploymentJobDefinitionsCreateOptions.Deployment != nil {
		body["deployment"] = deploymentJobDefinitionsCreateOptions.Deployment
	}
	if deploymentJobDefinitionsCreateOptions.Description != nil {
		body["description"] = deploymentJobDefinitionsCreateOptions.Description
	}
	if deploymentJobDefinitionsCreateOptions.Tags != nil {
		body["tags"] = deploymentJobDefinitionsCreateOptions.Tags
	}
	if deploymentJobDefinitionsCreateOptions.Custom != nil {
		body["custom"] = deploymentJobDefinitionsCreateOptions.Custom
	}
	if deploymentJobDefinitionsCreateOptions.HardwareSpec != nil {
		body["hardware_spec"] = deploymentJobDefinitionsCreateOptions.HardwareSpec
	}
	if deploymentJobDefinitionsCreateOptions.HybridPipelineHardwareSpecs != nil {
		body["hybrid_pipeline_hardware_specs"] = deploymentJobDefinitionsCreateOptions.HybridPipelineHardwareSpecs
	}
	if deploymentJobDefinitionsCreateOptions.Scoring != nil {
		body["scoring"] = deploymentJobDefinitionsCreateOptions.Scoring
	}
	if deploymentJobDefinitionsCreateOptions.DecisionOptimization != nil {
		body["decision_optimization"] = deploymentJobDefinitionsCreateOptions.DecisionOptimization
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobDefinitionsList : Retrieve the deployment job definitions
// Retrieve the deployment job definitions for the specified space.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsList(deploymentJobDefinitionsListOptions *DeploymentJobDefinitionsListOptions) (result *JobResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobDefinitionsListWithContext(context.Background(), deploymentJobDefinitionsListOptions)
}

// DeploymentJobDefinitionsListWithContext is an alternate form of the DeploymentJobDefinitionsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsListWithContext(ctx context.Context, deploymentJobDefinitionsListOptions *DeploymentJobDefinitionsListOptions) (result *JobResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobDefinitionsListOptions, "deploymentJobDefinitionsListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobDefinitionsListOptions, "deploymentJobDefinitionsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_job_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobDefinitionsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobDefinitionsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobDefinitionsListOptions.SpaceID))
	if deploymentJobDefinitionsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*deploymentJobDefinitionsListOptions.Start))
	}
	if deploymentJobDefinitionsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*deploymentJobDefinitionsListOptions.Limit))
	}
	if deploymentJobDefinitionsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*deploymentJobDefinitionsListOptions.TagValue))
	}
	if deploymentJobDefinitionsListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*deploymentJobDefinitionsListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobDefinitionsGet : Retrieve the deployment job definition
// Retrieve the deployment job definition with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsGet(deploymentJobDefinitionsGetOptions *DeploymentJobDefinitionsGetOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobDefinitionsGetWithContext(context.Background(), deploymentJobDefinitionsGetOptions)
}

// DeploymentJobDefinitionsGetWithContext is an alternate form of the DeploymentJobDefinitionsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsGetWithContext(ctx context.Context, deploymentJobDefinitionsGetOptions *DeploymentJobDefinitionsGetOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobDefinitionsGetOptions, "deploymentJobDefinitionsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobDefinitionsGetOptions, "deploymentJobDefinitionsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_definition_id": *deploymentJobDefinitionsGetOptions.JobDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_job_definitions/{job_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobDefinitionsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobDefinitionsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobDefinitionsGetOptions.SpaceID))
	if deploymentJobDefinitionsGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*deploymentJobDefinitionsGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobDefinitionsUpdate : Update the deployment job definition
// Update the deployment job definition with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`
// - `/deployment`.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsUpdate(deploymentJobDefinitionsUpdateOptions *DeploymentJobDefinitionsUpdateOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobDefinitionsUpdateWithContext(context.Background(), deploymentJobDefinitionsUpdateOptions)
}

// DeploymentJobDefinitionsUpdateWithContext is an alternate form of the DeploymentJobDefinitionsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsUpdateWithContext(ctx context.Context, deploymentJobDefinitionsUpdateOptions *DeploymentJobDefinitionsUpdateOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobDefinitionsUpdateOptions, "deploymentJobDefinitionsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobDefinitionsUpdateOptions, "deploymentJobDefinitionsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_definition_id": *deploymentJobDefinitionsUpdateOptions.JobDefinitionID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_job_definitions/{job_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobDefinitionsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobDefinitionsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobDefinitionsUpdateOptions.SpaceID))

	_, err = builder.SetBodyContentJSON(deploymentJobDefinitionsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobDefinitionsDelete : Delete the deployment job definition
// Delete the deployment job definition with the specified identifier. This will delete all revisions of this deployment
// job definition as well. For each revision all attachments will also be deleted.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsDelete(deploymentJobDefinitionsDeleteOptions *DeploymentJobDefinitionsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobDefinitionsDeleteWithContext(context.Background(), deploymentJobDefinitionsDeleteOptions)
}

// DeploymentJobDefinitionsDeleteWithContext is an alternate form of the DeploymentJobDefinitionsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsDeleteWithContext(ctx context.Context, deploymentJobDefinitionsDeleteOptions *DeploymentJobDefinitionsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobDefinitionsDeleteOptions, "deploymentJobDefinitionsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobDefinitionsDeleteOptions, "deploymentJobDefinitionsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_definition_id": *deploymentJobDefinitionsDeleteOptions.JobDefinitionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_job_definitions/{job_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobDefinitionsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobDefinitionsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobDefinitionsDeleteOptions.SpaceID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// DeploymentJobDefinitionsCreateRevision : Create a new deployment job definition revision
// Create a new deployment job definition revision. The current metadata and content for job_definition_id will be taken
// and a new revision created. This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsCreateRevision(deploymentJobDefinitionsCreateRevisionOptions *DeploymentJobDefinitionsCreateRevisionOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobDefinitionsCreateRevisionWithContext(context.Background(), deploymentJobDefinitionsCreateRevisionOptions)
}

// DeploymentJobDefinitionsCreateRevisionWithContext is an alternate form of the DeploymentJobDefinitionsCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsCreateRevisionWithContext(ctx context.Context, deploymentJobDefinitionsCreateRevisionOptions *DeploymentJobDefinitionsCreateRevisionOptions) (result *JobResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobDefinitionsCreateRevisionOptions, "deploymentJobDefinitionsCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobDefinitionsCreateRevisionOptions, "deploymentJobDefinitionsCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_definition_id": *deploymentJobDefinitionsCreateRevisionOptions.JobDefinitionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_job_definitions/{job_definition_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobDefinitionsCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobDefinitionsCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if deploymentJobDefinitionsCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = deploymentJobDefinitionsCreateRevisionOptions.SpaceID
	}
	if deploymentJobDefinitionsCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = deploymentJobDefinitionsCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeploymentJobDefinitionsListRevisions : Retrieve the deployment job definition revisions
// Retrieve the deployment job definition revisions. This command is supported starting with release 3.5 of Cloud Pak
// for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsListRevisions(deploymentJobDefinitionsListRevisionsOptions *DeploymentJobDefinitionsListRevisionsOptions) (result *JobResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.DeploymentJobDefinitionsListRevisionsWithContext(context.Background(), deploymentJobDefinitionsListRevisionsOptions)
}

// DeploymentJobDefinitionsListRevisionsWithContext is an alternate form of the DeploymentJobDefinitionsListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) DeploymentJobDefinitionsListRevisionsWithContext(ctx context.Context, deploymentJobDefinitionsListRevisionsOptions *DeploymentJobDefinitionsListRevisionsOptions) (result *JobResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deploymentJobDefinitionsListRevisionsOptions, "deploymentJobDefinitionsListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deploymentJobDefinitionsListRevisionsOptions, "deploymentJobDefinitionsListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"job_definition_id": *deploymentJobDefinitionsListRevisionsOptions.JobDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/deployment_job_definitions/{job_definition_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deploymentJobDefinitionsListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "DeploymentJobDefinitionsListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("space_id", fmt.Sprint(*deploymentJobDefinitionsListRevisionsOptions.SpaceID))
	if deploymentJobDefinitionsListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*deploymentJobDefinitionsListRevisionsOptions.Start))
	}
	if deploymentJobDefinitionsListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*deploymentJobDefinitionsListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalJobResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExperimentsCreate : Create a new experiment
// Create a new experiment with the given payload. An experiment represents an asset that captures a set of `pipeline`
// or `model definition` assets that will be trained at the same time on the same data set.
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsCreate(experimentsCreateOptions *ExperimentsCreateOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ExperimentsCreateWithContext(context.Background(), experimentsCreateOptions)
}

// ExperimentsCreateWithContext is an alternate form of the ExperimentsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsCreateWithContext(ctx context.Context, experimentsCreateOptions *ExperimentsCreateOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(experimentsCreateOptions, "experimentsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(experimentsCreateOptions, "experimentsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/experiments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range experimentsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ExperimentsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if experimentsCreateOptions.Name != nil {
		body["name"] = experimentsCreateOptions.Name
	}
	if experimentsCreateOptions.ProjectID != nil {
		body["project_id"] = experimentsCreateOptions.ProjectID
	}
	if experimentsCreateOptions.SpaceID != nil {
		body["space_id"] = experimentsCreateOptions.SpaceID
	}
	if experimentsCreateOptions.Description != nil {
		body["description"] = experimentsCreateOptions.Description
	}
	if experimentsCreateOptions.Tags != nil {
		body["tags"] = experimentsCreateOptions.Tags
	}
	if experimentsCreateOptions.LabelColumn != nil {
		body["label_column"] = experimentsCreateOptions.LabelColumn
	}
	if experimentsCreateOptions.EvaluationDefinition != nil {
		body["evaluation_definition"] = experimentsCreateOptions.EvaluationDefinition
	}
	if experimentsCreateOptions.TrainingReferences != nil {
		body["training_references"] = experimentsCreateOptions.TrainingReferences
	}
	if experimentsCreateOptions.Custom != nil {
		body["custom"] = experimentsCreateOptions.Custom
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExperimentResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExperimentsList : Retrieve the experiments
// Retrieve the experiments for the specified space or project.
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsList(experimentsListOptions *ExperimentsListOptions) (result *ExperimentResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ExperimentsListWithContext(context.Background(), experimentsListOptions)
}

// ExperimentsListWithContext is an alternate form of the ExperimentsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsListWithContext(ctx context.Context, experimentsListOptions *ExperimentsListOptions) (result *ExperimentResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(experimentsListOptions, "experimentsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/experiments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range experimentsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ExperimentsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if experimentsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*experimentsListOptions.SpaceID))
	}
	if experimentsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*experimentsListOptions.ProjectID))
	}
	if experimentsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*experimentsListOptions.Start))
	}
	if experimentsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*experimentsListOptions.Limit))
	}
	if experimentsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*experimentsListOptions.TagValue))
	}
	if experimentsListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*experimentsListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExperimentResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExperimentsGet : Retrieve the experiment
// Retrieve the experiment with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsGet(experimentsGetOptions *ExperimentsGetOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ExperimentsGetWithContext(context.Background(), experimentsGetOptions)
}

// ExperimentsGetWithContext is an alternate form of the ExperimentsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsGetWithContext(ctx context.Context, experimentsGetOptions *ExperimentsGetOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(experimentsGetOptions, "experimentsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(experimentsGetOptions, "experimentsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"experiment_id": *experimentsGetOptions.ExperimentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/experiments/{experiment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range experimentsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ExperimentsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if experimentsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*experimentsGetOptions.SpaceID))
	}
	if experimentsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*experimentsGetOptions.ProjectID))
	}
	if experimentsGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*experimentsGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExperimentResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExperimentsUpdate : Update the experiment
// Update the experiment with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`.
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsUpdate(experimentsUpdateOptions *ExperimentsUpdateOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ExperimentsUpdateWithContext(context.Background(), experimentsUpdateOptions)
}

// ExperimentsUpdateWithContext is an alternate form of the ExperimentsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsUpdateWithContext(ctx context.Context, experimentsUpdateOptions *ExperimentsUpdateOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(experimentsUpdateOptions, "experimentsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(experimentsUpdateOptions, "experimentsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"experiment_id": *experimentsUpdateOptions.ExperimentID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/experiments/{experiment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range experimentsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ExperimentsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if experimentsUpdateOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*experimentsUpdateOptions.SpaceID))
	}
	if experimentsUpdateOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*experimentsUpdateOptions.ProjectID))
	}

	_, err = builder.SetBodyContentJSON(experimentsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExperimentResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExperimentsDelete : Delete the experiment
// Delete the experiment with the specified identifier. This will delete all revisions of this experiment as well. For
// each revision all attachments will also be deleted.
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsDelete(experimentsDeleteOptions *ExperimentsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ExperimentsDeleteWithContext(context.Background(), experimentsDeleteOptions)
}

// ExperimentsDeleteWithContext is an alternate form of the ExperimentsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsDeleteWithContext(ctx context.Context, experimentsDeleteOptions *ExperimentsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(experimentsDeleteOptions, "experimentsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(experimentsDeleteOptions, "experimentsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"experiment_id": *experimentsDeleteOptions.ExperimentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/experiments/{experiment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range experimentsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ExperimentsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if experimentsDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*experimentsDeleteOptions.SpaceID))
	}
	if experimentsDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*experimentsDeleteOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// ExperimentsCreateRevision : Create a new experiment revision
// Create a new experiment revision. The current metadata and content for experiment_id will be taken and a new revision
// created. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsCreateRevision(experimentsCreateRevisionOptions *ExperimentsCreateRevisionOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ExperimentsCreateRevisionWithContext(context.Background(), experimentsCreateRevisionOptions)
}

// ExperimentsCreateRevisionWithContext is an alternate form of the ExperimentsCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsCreateRevisionWithContext(ctx context.Context, experimentsCreateRevisionOptions *ExperimentsCreateRevisionOptions) (result *ExperimentResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(experimentsCreateRevisionOptions, "experimentsCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(experimentsCreateRevisionOptions, "experimentsCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"experiment_id": *experimentsCreateRevisionOptions.ExperimentID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/experiments/{experiment_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range experimentsCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ExperimentsCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if experimentsCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = experimentsCreateRevisionOptions.SpaceID
	}
	if experimentsCreateRevisionOptions.ProjectID != nil {
		body["project_id"] = experimentsCreateRevisionOptions.ProjectID
	}
	if experimentsCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = experimentsCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExperimentResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExperimentsListRevisions : Retrieve the experiment revisions
// Retrieve the experiment revisions.
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsListRevisions(experimentsListRevisionsOptions *ExperimentsListRevisionsOptions) (result *ExperimentResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ExperimentsListRevisionsWithContext(context.Background(), experimentsListRevisionsOptions)
}

// ExperimentsListRevisionsWithContext is an alternate form of the ExperimentsListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ExperimentsListRevisionsWithContext(ctx context.Context, experimentsListRevisionsOptions *ExperimentsListRevisionsOptions) (result *ExperimentResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(experimentsListRevisionsOptions, "experimentsListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(experimentsListRevisionsOptions, "experimentsListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"experiment_id": *experimentsListRevisionsOptions.ExperimentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/experiments/{experiment_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range experimentsListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ExperimentsListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if experimentsListRevisionsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*experimentsListRevisionsOptions.SpaceID))
	}
	if experimentsListRevisionsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*experimentsListRevisionsOptions.ProjectID))
	}
	if experimentsListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*experimentsListRevisionsOptions.Start))
	}
	if experimentsListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*experimentsListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExperimentResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsCreate : Create a new function
// Create a new function with the given payload. A function is some code that can be deployed as an online, or batch
// deployment.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsCreate(functionsCreateOptions *FunctionsCreateOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsCreateWithContext(context.Background(), functionsCreateOptions)
}

// FunctionsCreateWithContext is an alternate form of the FunctionsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsCreateWithContext(ctx context.Context, functionsCreateOptions *FunctionsCreateOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsCreateOptions, "functionsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsCreateOptions, "functionsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if functionsCreateOptions.Name != nil {
		body["name"] = functionsCreateOptions.Name
	}
	if functionsCreateOptions.SoftwareSpec != nil {
		body["software_spec"] = functionsCreateOptions.SoftwareSpec
	}
	if functionsCreateOptions.ProjectID != nil {
		body["project_id"] = functionsCreateOptions.ProjectID
	}
	if functionsCreateOptions.SpaceID != nil {
		body["space_id"] = functionsCreateOptions.SpaceID
	}
	if functionsCreateOptions.Description != nil {
		body["description"] = functionsCreateOptions.Description
	}
	if functionsCreateOptions.Tags != nil {
		body["tags"] = functionsCreateOptions.Tags
	}
	if functionsCreateOptions.Type != nil {
		body["type"] = functionsCreateOptions.Type
	}
	if functionsCreateOptions.SampleScoringInput != nil {
		body["sample_scoring_input"] = functionsCreateOptions.SampleScoringInput
	}
	if functionsCreateOptions.Schemas != nil {
		body["schemas"] = functionsCreateOptions.Schemas
	}
	if functionsCreateOptions.Custom != nil {
		body["custom"] = functionsCreateOptions.Custom
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFunctionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsList : Retrieve the functions
// Retrieve the functions for the specified space or project.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsList(functionsListOptions *FunctionsListOptions) (result *FunctionResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsListWithContext(context.Background(), functionsListOptions)
}

// FunctionsListWithContext is an alternate form of the FunctionsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsListWithContext(ctx context.Context, functionsListOptions *FunctionsListOptions) (result *FunctionResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(functionsListOptions, "functionsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if functionsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*functionsListOptions.SpaceID))
	}
	if functionsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*functionsListOptions.ProjectID))
	}
	if functionsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*functionsListOptions.Start))
	}
	if functionsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*functionsListOptions.Limit))
	}
	if functionsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*functionsListOptions.TagValue))
	}
	if functionsListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*functionsListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFunctionResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsGet : Retrieve the function
// Retrieve the function with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsGet(functionsGetOptions *FunctionsGetOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsGetWithContext(context.Background(), functionsGetOptions)
}

// FunctionsGetWithContext is an alternate form of the FunctionsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsGetWithContext(ctx context.Context, functionsGetOptions *FunctionsGetOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsGetOptions, "functionsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsGetOptions, "functionsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"function_id": *functionsGetOptions.FunctionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions/{function_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if functionsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*functionsGetOptions.SpaceID))
	}
	if functionsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*functionsGetOptions.ProjectID))
	}
	if functionsGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*functionsGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFunctionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsUpdate : Update the function
// Update the function with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsUpdate(functionsUpdateOptions *FunctionsUpdateOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsUpdateWithContext(context.Background(), functionsUpdateOptions)
}

// FunctionsUpdateWithContext is an alternate form of the FunctionsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsUpdateWithContext(ctx context.Context, functionsUpdateOptions *FunctionsUpdateOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsUpdateOptions, "functionsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsUpdateOptions, "functionsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"function_id": *functionsUpdateOptions.FunctionID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions/{function_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if functionsUpdateOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*functionsUpdateOptions.SpaceID))
	}
	if functionsUpdateOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*functionsUpdateOptions.ProjectID))
	}

	_, err = builder.SetBodyContentJSON(functionsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFunctionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsDelete : Delete the function
// Delete the function with the specified identifier. This will delete all revisions of this function as well. For each
// revision all attachments will also be deleted.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsDelete(functionsDeleteOptions *FunctionsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsDeleteWithContext(context.Background(), functionsDeleteOptions)
}

// FunctionsDeleteWithContext is an alternate form of the FunctionsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsDeleteWithContext(ctx context.Context, functionsDeleteOptions *FunctionsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsDeleteOptions, "functionsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsDeleteOptions, "functionsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"function_id": *functionsDeleteOptions.FunctionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions/{function_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if functionsDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*functionsDeleteOptions.SpaceID))
	}
	if functionsDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*functionsDeleteOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// FunctionsCreateRevision : Create a new function revision
// Create a new function revision. The current metadata and content for function_id will be taken and a new revision
// created. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsCreateRevision(functionsCreateRevisionOptions *FunctionsCreateRevisionOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsCreateRevisionWithContext(context.Background(), functionsCreateRevisionOptions)
}

// FunctionsCreateRevisionWithContext is an alternate form of the FunctionsCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsCreateRevisionWithContext(ctx context.Context, functionsCreateRevisionOptions *FunctionsCreateRevisionOptions) (result *FunctionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsCreateRevisionOptions, "functionsCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsCreateRevisionOptions, "functionsCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"function_id": *functionsCreateRevisionOptions.FunctionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions/{function_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if functionsCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = functionsCreateRevisionOptions.SpaceID
	}
	if functionsCreateRevisionOptions.ProjectID != nil {
		body["project_id"] = functionsCreateRevisionOptions.ProjectID
	}
	if functionsCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = functionsCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFunctionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsListRevisions : Retrieve the function revisions
// Retrieve the function revisions.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsListRevisions(functionsListRevisionsOptions *FunctionsListRevisionsOptions) (result *FunctionResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsListRevisionsWithContext(context.Background(), functionsListRevisionsOptions)
}

// FunctionsListRevisionsWithContext is an alternate form of the FunctionsListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsListRevisionsWithContext(ctx context.Context, functionsListRevisionsOptions *FunctionsListRevisionsOptions) (result *FunctionResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsListRevisionsOptions, "functionsListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsListRevisionsOptions, "functionsListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"function_id": *functionsListRevisionsOptions.FunctionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions/{function_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if functionsListRevisionsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*functionsListRevisionsOptions.SpaceID))
	}
	if functionsListRevisionsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*functionsListRevisionsOptions.ProjectID))
	}
	if functionsListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*functionsListRevisionsOptions.Start))
	}
	if functionsListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*functionsListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFunctionResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsUploadCode : Upload the function code
// Upload the function code. Functions expect a zip file that contains a python file     that make up the function.
// Python functions specify what needs to be run when     the function is deployed and what needs to be run when the
// scoring function is     called. In other words, you are able to customize what preparation WML does in     the
// environment when you deploy the function, as well as what steps WML takes to     generate the output when you call
// the API later on. The function consists of the     outer function (any place that is not within the score function)
// and the inner     score function. The code that sits in the outer function runs when the function     is deployed,
// and the environment is then frozen and ready to be used whenever     the online scoring or batch inline job
// processing API is called. The code that     sits in the inner score function runs when the online scoring or batch
// inline     job processing API is called, in the environment customized by the outer function.     See [Deploying
// Python
// function](https://dataplatform.cloud.ibm.com/docs/content/wsj/analyze-data/ml-deploy-py-function.html?context=cpdaas${content_description}audience=wdp)
// for more details.         This is illustrated in the example below:        <pre> <br />     ...python code used to
// set up the environment... <br />     <br />     def score(payload): <br />         df_payload =
// pd.DataFrame(payload[values]) <br />         df_payload.columns = payload[fields] <br />         ... <br />
// output = {result : res} <br />         return output <br />     <br />     return score <br />     </pre>.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsUploadCode(functionsUploadCodeOptions *FunctionsUploadCodeOptions) (result *ContentMetadata, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsUploadCodeWithContext(context.Background(), functionsUploadCodeOptions)
}

// FunctionsUploadCodeWithContext is an alternate form of the FunctionsUploadCode method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsUploadCodeWithContext(ctx context.Context, functionsUploadCodeOptions *FunctionsUploadCodeOptions) (result *ContentMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsUploadCodeOptions, "functionsUploadCodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsUploadCodeOptions, "functionsUploadCodeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"function_id": *functionsUploadCodeOptions.FunctionID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions/{function_id}/code`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsUploadCodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsUploadCode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/gzip")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if functionsUploadCodeOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*functionsUploadCodeOptions.SpaceID))
	}
	if functionsUploadCodeOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*functionsUploadCodeOptions.ProjectID))
	}

	_, err = builder.SetBodyContent("application/gzip", nil, nil, functionsUploadCodeOptions.UploadCode)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContentMetadata)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// FunctionsDownloadCode : Download the function code
// Download the function code. It is possible to get the `code` for a given revision of the `function`. Functions expect
// a zip file that contains a python file     that make up the function. Python functions specify what needs to be run
// when     the function is deployed and what needs to be run when the scoring function is     called. In other words,
// you are able to customize what preparation WML does in     the environment when you deploy the function, as well as
// what steps WML takes to     generate the output when you call the API later on. The function consists of the
// outer function (any place that is not within the score function) and the inner     score function. The code that sits
// in the outer function runs when the function     is deployed, and the environment is then frozen and ready to be used
// whenever     the online scoring or batch inline job processing API is called. The code that     sits in the inner
// score function runs when the online scoring or batch inline     job processing API is called, in the environment
// customized by the outer function.     See [Deploying Python
// function](https://dataplatform.cloud.ibm.com/docs/content/wsj/analyze-data/ml-deploy-py-function.html?context=cpdaas${content_description}audience=wdp)
// for more details.         This is illustrated in the example below:        <pre> <br />     ...python code used to
// set up the environment... <br />     <br />     def score(payload): <br />         df_payload =
// pd.DataFrame(payload[values]) <br />         df_payload.columns = payload[fields] <br />         ... <br />
// output = {result : res} <br />         return output <br />     <br />     return score <br />     </pre>.
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsDownloadCode(functionsDownloadCodeOptions *FunctionsDownloadCodeOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.FunctionsDownloadCodeWithContext(context.Background(), functionsDownloadCodeOptions)
}

// FunctionsDownloadCodeWithContext is an alternate form of the FunctionsDownloadCode method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) FunctionsDownloadCodeWithContext(ctx context.Context, functionsDownloadCodeOptions *FunctionsDownloadCodeOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(functionsDownloadCodeOptions, "functionsDownloadCodeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(functionsDownloadCodeOptions, "functionsDownloadCodeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"function_id": *functionsDownloadCodeOptions.FunctionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/functions/{function_id}/code`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range functionsDownloadCodeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "FunctionsDownloadCode")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/zip")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if functionsDownloadCodeOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*functionsDownloadCodeOptions.SpaceID))
	}
	if functionsDownloadCodeOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*functionsDownloadCodeOptions.ProjectID))
	}
	if functionsDownloadCodeOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*functionsDownloadCodeOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, &result)

	return
}

// InstancesList : Retrieve the service instances
// Retrieve the service instances. Either `space_id` or `project_id` query parameter is mandatory but both can be
// provided.
func (watsonMachineLearning *WatsonMachineLearningV4) InstancesList(instancesListOptions *InstancesListOptions) (result *InstanceResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.InstancesListWithContext(context.Background(), instancesListOptions)
}

// InstancesListWithContext is an alternate form of the InstancesList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) InstancesListWithContext(ctx context.Context, instancesListOptions *InstancesListOptions) (result *InstanceResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(instancesListOptions, "instancesListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/instances`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range instancesListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "InstancesList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if instancesListOptions.SpaceID != nil {
		builder.AddQuery("space_id", strings.Join(instancesListOptions.SpaceID, ","))
	}
	if instancesListOptions.ProjectID != nil {
		builder.AddQuery("project_id", strings.Join(instancesListOptions.ProjectID, ","))
	}
	if instancesListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*instancesListOptions.Start))
	}
	if instancesListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*instancesListOptions.Limit))
	}
	if instancesListOptions.ConsumptionDetails != nil {
		builder.AddQuery("consumption_details", fmt.Sprint(*instancesListOptions.ConsumptionDetails))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InstancesGet : Retrieve the service instance
// Retrieve the service instance details. Instances with `plan.version` set to `2` will have complete data. Instances
// with `plan.version` set to `1` will not have the `consumption` section and one should refer to the v3 API usage
// section.
func (watsonMachineLearning *WatsonMachineLearningV4) InstancesGet(instancesGetOptions *InstancesGetOptions) (result *InstanceResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.InstancesGetWithContext(context.Background(), instancesGetOptions)
}

// InstancesGetWithContext is an alternate form of the InstancesGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) InstancesGetWithContext(ctx context.Context, instancesGetOptions *InstancesGetOptions) (result *InstanceResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(instancesGetOptions, "instancesGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(instancesGetOptions, "instancesGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *instancesGetOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range instancesGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "InstancesGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if instancesGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*instancesGetOptions.SpaceID))
	}
	if instancesGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*instancesGetOptions.ProjectID))
	}
	if instancesGetOptions.ConsumptionDetails != nil {
		builder.AddQuery("consumption_details", fmt.Sprint(*instancesGetOptions.ConsumptionDetails))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstanceResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsCreate : Create a new model
// Create a new model with the given payload. A model represents a machine learning model asset. If a `202` status is
// returned then the model will be ready when the
// `content_import_state` in the model status (/ml/v4/models/{model_id}) is `completed`. If `content_import_state` is
// not used then a `201` status is returned.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsCreate(modelsCreateOptions *ModelsCreateOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsCreateWithContext(context.Background(), modelsCreateOptions)
}

// ModelsCreateWithContext is an alternate form of the ModelsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsCreateWithContext(ctx context.Context, modelsCreateOptions *ModelsCreateOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsCreateOptions, "modelsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsCreateOptions, "modelsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if modelsCreateOptions.Name != nil {
		body["name"] = modelsCreateOptions.Name
	}
	if modelsCreateOptions.Type != nil {
		body["type"] = modelsCreateOptions.Type
	}
	if modelsCreateOptions.SoftwareSpec != nil {
		body["software_spec"] = modelsCreateOptions.SoftwareSpec
	}
	if modelsCreateOptions.ProjectID != nil {
		body["project_id"] = modelsCreateOptions.ProjectID
	}
	if modelsCreateOptions.SpaceID != nil {
		body["space_id"] = modelsCreateOptions.SpaceID
	}
	if modelsCreateOptions.Description != nil {
		body["description"] = modelsCreateOptions.Description
	}
	if modelsCreateOptions.Tags != nil {
		body["tags"] = modelsCreateOptions.Tags
	}
	if modelsCreateOptions.Pipeline != nil {
		body["pipeline"] = modelsCreateOptions.Pipeline
	}
	if modelsCreateOptions.ModelDefinition != nil {
		body["model_definition"] = modelsCreateOptions.ModelDefinition
	}
	if modelsCreateOptions.HyperParameters != nil {
		body["hyper_parameters"] = modelsCreateOptions.HyperParameters
	}
	if modelsCreateOptions.Domain != nil {
		body["domain"] = modelsCreateOptions.Domain
	}
	if modelsCreateOptions.TrainingDataReferences != nil {
		body["training_data_references"] = modelsCreateOptions.TrainingDataReferences
	}
	if modelsCreateOptions.TestDataReferences != nil {
		body["test_data_references"] = modelsCreateOptions.TestDataReferences
	}
	if modelsCreateOptions.Schemas != nil {
		body["schemas"] = modelsCreateOptions.Schemas
	}
	if modelsCreateOptions.LabelColumn != nil {
		body["label_column"] = modelsCreateOptions.LabelColumn
	}
	if modelsCreateOptions.TransformedLabelColumn != nil {
		body["transformed_label_column"] = modelsCreateOptions.TransformedLabelColumn
	}
	if modelsCreateOptions.Size != nil {
		body["size"] = modelsCreateOptions.Size
	}
	if modelsCreateOptions.Metrics != nil {
		body["metrics"] = modelsCreateOptions.Metrics
	}
	if modelsCreateOptions.Custom != nil {
		body["custom"] = modelsCreateOptions.Custom
	}
	if modelsCreateOptions.UserDefinedObjects != nil {
		body["user_defined_objects"] = modelsCreateOptions.UserDefinedObjects
	}
	if modelsCreateOptions.ContentLocation != nil {
		body["content_location"] = modelsCreateOptions.ContentLocation
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsList : Retrieve the models
// Retrieve the models for the specified space or project.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsList(modelsListOptions *ModelsListOptions) (result *ModelResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsListWithContext(context.Background(), modelsListOptions)
}

// ModelsListWithContext is an alternate form of the ModelsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsListWithContext(ctx context.Context, modelsListOptions *ModelsListOptions) (result *ModelResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(modelsListOptions, "modelsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsListOptions.SpaceID))
	}
	if modelsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsListOptions.ProjectID))
	}
	if modelsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*modelsListOptions.Start))
	}
	if modelsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*modelsListOptions.Limit))
	}
	if modelsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*modelsListOptions.TagValue))
	}
	if modelsListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*modelsListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsGet : Retrieve the model
// Retrieve the model with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsGet(modelsGetOptions *ModelsGetOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsGetWithContext(context.Background(), modelsGetOptions)
}

// ModelsGetWithContext is an alternate form of the ModelsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsGetWithContext(ctx context.Context, modelsGetOptions *ModelsGetOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsGetOptions, "modelsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsGetOptions, "modelsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsGetOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsGetOptions.SpaceID))
	}
	if modelsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsGetOptions.ProjectID))
	}
	if modelsGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*modelsGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsUpdate : Update the model
// Update the model with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`
// - `/software_spec` (operation 'replace' only).
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsUpdate(modelsUpdateOptions *ModelsUpdateOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsUpdateWithContext(context.Background(), modelsUpdateOptions)
}

// ModelsUpdateWithContext is an alternate form of the ModelsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsUpdateWithContext(ctx context.Context, modelsUpdateOptions *ModelsUpdateOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsUpdateOptions, "modelsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsUpdateOptions, "modelsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsUpdateOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsUpdateOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsUpdateOptions.SpaceID))
	}
	if modelsUpdateOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsUpdateOptions.ProjectID))
	}

	_, err = builder.SetBodyContentJSON(modelsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsDelete : Delete the model
// Delete the model with the specified identifier. This will delete all revisions of this model as well. For each
// revision all attachments will also be deleted.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsDelete(modelsDeleteOptions *ModelsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsDeleteWithContext(context.Background(), modelsDeleteOptions)
}

// ModelsDeleteWithContext is an alternate form of the ModelsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsDeleteWithContext(ctx context.Context, modelsDeleteOptions *ModelsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsDeleteOptions, "modelsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsDeleteOptions, "modelsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsDeleteOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsDeleteOptions.SpaceID))
	}
	if modelsDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsDeleteOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// ModelsCreateRevision : Create a new model revision
// Create a new model revision. The current metadata and content for model_id will be taken and a new revision created.
// Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsCreateRevision(modelsCreateRevisionOptions *ModelsCreateRevisionOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsCreateRevisionWithContext(context.Background(), modelsCreateRevisionOptions)
}

// ModelsCreateRevisionWithContext is an alternate form of the ModelsCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsCreateRevisionWithContext(ctx context.Context, modelsCreateRevisionOptions *ModelsCreateRevisionOptions) (result *ModelResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsCreateRevisionOptions, "modelsCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsCreateRevisionOptions, "modelsCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsCreateRevisionOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if modelsCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = modelsCreateRevisionOptions.SpaceID
	}
	if modelsCreateRevisionOptions.ProjectID != nil {
		body["project_id"] = modelsCreateRevisionOptions.ProjectID
	}
	if modelsCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = modelsCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsListRevisions : Retrieve the model revisions
// Retrieve the model revisions.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsListRevisions(modelsListRevisionsOptions *ModelsListRevisionsOptions) (result *ModelResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsListRevisionsWithContext(context.Background(), modelsListRevisionsOptions)
}

// ModelsListRevisionsWithContext is an alternate form of the ModelsListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsListRevisionsWithContext(ctx context.Context, modelsListRevisionsOptions *ModelsListRevisionsOptions) (result *ModelResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsListRevisionsOptions, "modelsListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsListRevisionsOptions, "modelsListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsListRevisionsOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsListRevisionsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsListRevisionsOptions.SpaceID))
	}
	if modelsListRevisionsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsListRevisionsOptions.ProjectID))
	}
	if modelsListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*modelsListRevisionsOptions.Start))
	}
	if modelsListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*modelsListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsListAttachments : Retrieve the model content metadata list
// Retrieve the content metadata list for the specified model attachments. This command is supported starting with
// release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsListAttachments(modelsListAttachmentsOptions *ModelsListAttachmentsOptions) (result *AllContentMetadata, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsListAttachmentsWithContext(context.Background(), modelsListAttachmentsOptions)
}

// ModelsListAttachmentsWithContext is an alternate form of the ModelsListAttachments method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsListAttachmentsWithContext(ctx context.Context, modelsListAttachmentsOptions *ModelsListAttachmentsOptions) (result *AllContentMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsListAttachmentsOptions, "modelsListAttachmentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsListAttachmentsOptions, "modelsListAttachmentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsListAttachmentsOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}/content`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsListAttachmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsListAttachments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsListAttachmentsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsListAttachmentsOptions.SpaceID))
	}
	if modelsListAttachmentsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsListAttachmentsOptions.ProjectID))
	}
	if modelsListAttachmentsOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*modelsListAttachmentsOptions.Rev))
	}
	if modelsListAttachmentsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*modelsListAttachmentsOptions.Name))
	}
	if modelsListAttachmentsOptions.ContentFormat != nil {
		builder.AddQuery("content_format", fmt.Sprint(*modelsListAttachmentsOptions.ContentFormat))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAllContentMetadata)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsUploadContent : Upload the model content
// Upload the content for the specified model.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsUploadContent(modelsUploadContentOptions *ModelsUploadContentOptions) (result *ContentMetadata, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsUploadContentWithContext(context.Background(), modelsUploadContentOptions)
}

// ModelsUploadContentWithContext is an alternate form of the ModelsUploadContent method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsUploadContentWithContext(ctx context.Context, modelsUploadContentOptions *ModelsUploadContentOptions) (result *ContentMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsUploadContentOptions, "modelsUploadContentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsUploadContentOptions, "modelsUploadContentOptions")
	if err != nil {
		return
	}

	if modelsUploadContentOptions.UploadContent != nil && modelsUploadContentOptions.ContentType == nil {
		modelsUploadContentOptions.SetContentType("application/json")
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsUploadContentOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}/content`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsUploadContentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsUploadContent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if modelsUploadContentOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*modelsUploadContentOptions.ContentType))
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	builder.AddQuery("content_format", fmt.Sprint(*modelsUploadContentOptions.ContentFormat))
	if modelsUploadContentOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsUploadContentOptions.SpaceID))
	}
	if modelsUploadContentOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsUploadContentOptions.ProjectID))
	}
	if modelsUploadContentOptions.PipelineNodeID != nil {
		builder.AddQuery("pipeline_node_id", fmt.Sprint(*modelsUploadContentOptions.PipelineNodeID))
	}
	if modelsUploadContentOptions.DeploymentID != nil {
		builder.AddQuery("deployment_id", fmt.Sprint(*modelsUploadContentOptions.DeploymentID))
	}
	if modelsUploadContentOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*modelsUploadContentOptions.Name))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(modelsUploadContentOptions.ContentType), modelsUploadContentOptions.UploadContent, nil, modelsUploadContentOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContentMetadata)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelsDownloadContent : Download the model content
// Download the model content.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsDownloadContent(modelsDownloadContentOptions *ModelsDownloadContentOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsDownloadContentWithContext(context.Background(), modelsDownloadContentOptions)
}

// ModelsDownloadContentWithContext is an alternate form of the ModelsDownloadContent method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsDownloadContentWithContext(ctx context.Context, modelsDownloadContentOptions *ModelsDownloadContentOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsDownloadContentOptions, "modelsDownloadContentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsDownloadContentOptions, "modelsDownloadContentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id":      *modelsDownloadContentOptions.ModelID,
		"attachment_id": *modelsDownloadContentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}/content/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsDownloadContentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsDownloadContent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/zip")
	if modelsDownloadContentOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*modelsDownloadContentOptions.Accept))
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsDownloadContentOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsDownloadContentOptions.SpaceID))
	}
	if modelsDownloadContentOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsDownloadContentOptions.ProjectID))
	}
	if modelsDownloadContentOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*modelsDownloadContentOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, &result)

	return
}

// ModelsDeleteContent : Delete the model content
// Delete the content for the specified model. This command is supported starting with release 3.5 of Cloud Pak for
// Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsDeleteContent(modelsDeleteContentOptions *ModelsDeleteContentOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsDeleteContentWithContext(context.Background(), modelsDeleteContentOptions)
}

// ModelsDeleteContentWithContext is an alternate form of the ModelsDeleteContent method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsDeleteContentWithContext(ctx context.Context, modelsDeleteContentOptions *ModelsDeleteContentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsDeleteContentOptions, "modelsDeleteContentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsDeleteContentOptions, "modelsDeleteContentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id":      *modelsDeleteContentOptions.ModelID,
		"attachment_id": *modelsDeleteContentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}/content/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsDeleteContentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsDeleteContent")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsDeleteContentOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsDeleteContentOptions.SpaceID))
	}
	if modelsDeleteContentOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsDeleteContentOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// ModelsFilteredDownload : Download the model content that matches a certain criteria
// Download the model content identified by the provided criteria. If more than one attachment is found for the given
// filter then a
// `400` error is returned. If there are no attachments that match the filter then a `404` error is returned. If there
// are no filters then, if there is a single attachment, the attachment content will be returned otherwise a `400` or
// `404` error will be returned as described above. This method is just a shortcut for getting the attachment metadata
// and then downloading by attachment id. This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsFilteredDownload(modelsFilteredDownloadOptions *ModelsFilteredDownloadOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelsFilteredDownloadWithContext(context.Background(), modelsFilteredDownloadOptions)
}

// ModelsFilteredDownloadWithContext is an alternate form of the ModelsFilteredDownload method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelsFilteredDownloadWithContext(ctx context.Context, modelsFilteredDownloadOptions *ModelsFilteredDownloadOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelsFilteredDownloadOptions, "modelsFilteredDownloadOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelsFilteredDownloadOptions, "modelsFilteredDownloadOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_id": *modelsFilteredDownloadOptions.ModelID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/models/{model_id}/download`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelsFilteredDownloadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelsFilteredDownload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/zip")
	if modelsFilteredDownloadOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*modelsFilteredDownloadOptions.Accept))
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelsFilteredDownloadOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelsFilteredDownloadOptions.SpaceID))
	}
	if modelsFilteredDownloadOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelsFilteredDownloadOptions.ProjectID))
	}
	if modelsFilteredDownloadOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*modelsFilteredDownloadOptions.Rev))
	}
	if modelsFilteredDownloadOptions.PipelineNodeID != nil {
		builder.AddQuery("pipeline_node_id", fmt.Sprint(*modelsFilteredDownloadOptions.PipelineNodeID))
	}
	if modelsFilteredDownloadOptions.DeploymentID != nil {
		builder.AddQuery("deployment_id", fmt.Sprint(*modelsFilteredDownloadOptions.DeploymentID))
	}
	if modelsFilteredDownloadOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*modelsFilteredDownloadOptions.Name))
	}
	if modelsFilteredDownloadOptions.ContentFormat != nil {
		builder.AddQuery("content_format", fmt.Sprint(*modelsFilteredDownloadOptions.ContentFormat))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, &result)

	return
}

// ModelDefinitionsCreate : Create a new model definition
// Create a new model definition with the given payload. A model definition represents the code that is used to train
// one or more models. This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsCreate(modelDefinitionsCreateOptions *ModelDefinitionsCreateOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsCreateWithContext(context.Background(), modelDefinitionsCreateOptions)
}

// ModelDefinitionsCreateWithContext is an alternate form of the ModelDefinitionsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsCreateWithContext(ctx context.Context, modelDefinitionsCreateOptions *ModelDefinitionsCreateOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsCreateOptions, "modelDefinitionsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsCreateOptions, "modelDefinitionsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if modelDefinitionsCreateOptions.Name != nil {
		body["name"] = modelDefinitionsCreateOptions.Name
	}
	if modelDefinitionsCreateOptions.Version != nil {
		body["version"] = modelDefinitionsCreateOptions.Version
	}
	if modelDefinitionsCreateOptions.Platform != nil {
		body["platform"] = modelDefinitionsCreateOptions.Platform
	}
	if modelDefinitionsCreateOptions.ProjectID != nil {
		body["project_id"] = modelDefinitionsCreateOptions.ProjectID
	}
	if modelDefinitionsCreateOptions.SpaceID != nil {
		body["space_id"] = modelDefinitionsCreateOptions.SpaceID
	}
	if modelDefinitionsCreateOptions.Description != nil {
		body["description"] = modelDefinitionsCreateOptions.Description
	}
	if modelDefinitionsCreateOptions.Tags != nil {
		body["tags"] = modelDefinitionsCreateOptions.Tags
	}
	if modelDefinitionsCreateOptions.Command != nil {
		body["command"] = modelDefinitionsCreateOptions.Command
	}
	if modelDefinitionsCreateOptions.Custom != nil {
		body["custom"] = modelDefinitionsCreateOptions.Custom
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelDefinitionsList : Retrieve the model definitions
// Retrieve the model definitions for the specified space or project. This command is supported starting with release
// 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsList(modelDefinitionsListOptions *ModelDefinitionsListOptions) (result *ModelDefinitionResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsListWithContext(context.Background(), modelDefinitionsListOptions)
}

// ModelDefinitionsListWithContext is an alternate form of the ModelDefinitionsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsListWithContext(ctx context.Context, modelDefinitionsListOptions *ModelDefinitionsListOptions) (result *ModelDefinitionResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(modelDefinitionsListOptions, "modelDefinitionsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelDefinitionsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelDefinitionsListOptions.SpaceID))
	}
	if modelDefinitionsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelDefinitionsListOptions.ProjectID))
	}
	if modelDefinitionsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*modelDefinitionsListOptions.Start))
	}
	if modelDefinitionsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*modelDefinitionsListOptions.Limit))
	}
	if modelDefinitionsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*modelDefinitionsListOptions.TagValue))
	}
	if modelDefinitionsListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*modelDefinitionsListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelDefinitionResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelDefinitionsGet : Retrieve the model definition
// Retrieve the model definition with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record. Either `space_id` or `project_id` has to be provided and is mandatory. This command is supported starting
// with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsGet(modelDefinitionsGetOptions *ModelDefinitionsGetOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsGetWithContext(context.Background(), modelDefinitionsGetOptions)
}

// ModelDefinitionsGetWithContext is an alternate form of the ModelDefinitionsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsGetWithContext(ctx context.Context, modelDefinitionsGetOptions *ModelDefinitionsGetOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsGetOptions, "modelDefinitionsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsGetOptions, "modelDefinitionsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_definition_id": *modelDefinitionsGetOptions.ModelDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions/{model_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelDefinitionsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelDefinitionsGetOptions.SpaceID))
	}
	if modelDefinitionsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelDefinitionsGetOptions.ProjectID))
	}
	if modelDefinitionsGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*modelDefinitionsGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelDefinitionsUpdate : Update the model definition
// Update the model definition with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom` This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsUpdate(modelDefinitionsUpdateOptions *ModelDefinitionsUpdateOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsUpdateWithContext(context.Background(), modelDefinitionsUpdateOptions)
}

// ModelDefinitionsUpdateWithContext is an alternate form of the ModelDefinitionsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsUpdateWithContext(ctx context.Context, modelDefinitionsUpdateOptions *ModelDefinitionsUpdateOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsUpdateOptions, "modelDefinitionsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsUpdateOptions, "modelDefinitionsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_definition_id": *modelDefinitionsUpdateOptions.ModelDefinitionID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions/{model_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelDefinitionsUpdateOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelDefinitionsUpdateOptions.SpaceID))
	}
	if modelDefinitionsUpdateOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelDefinitionsUpdateOptions.ProjectID))
	}

	_, err = builder.SetBodyContentJSON(modelDefinitionsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelDefinitionsDelete : Delete the model definition
// Delete the model definition with the specified identifier. This will delete all revisions of this model definition as
// well. For each revision all attachments will also be deleted. This command is supported starting with release 3.5 of
// Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsDelete(modelDefinitionsDeleteOptions *ModelDefinitionsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsDeleteWithContext(context.Background(), modelDefinitionsDeleteOptions)
}

// ModelDefinitionsDeleteWithContext is an alternate form of the ModelDefinitionsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsDeleteWithContext(ctx context.Context, modelDefinitionsDeleteOptions *ModelDefinitionsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsDeleteOptions, "modelDefinitionsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsDeleteOptions, "modelDefinitionsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_definition_id": *modelDefinitionsDeleteOptions.ModelDefinitionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions/{model_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelDefinitionsDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelDefinitionsDeleteOptions.SpaceID))
	}
	if modelDefinitionsDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelDefinitionsDeleteOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// ModelDefinitionsCreateRevision : Create a new model definition revision
// Create a new model definition revision. The current metadata and content for model_definition_id will be taken and a
// new revision created. Either `space_id` or `project_id` has to be provided and is mandatory. This command is
// supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsCreateRevision(modelDefinitionsCreateRevisionOptions *ModelDefinitionsCreateRevisionOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsCreateRevisionWithContext(context.Background(), modelDefinitionsCreateRevisionOptions)
}

// ModelDefinitionsCreateRevisionWithContext is an alternate form of the ModelDefinitionsCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsCreateRevisionWithContext(ctx context.Context, modelDefinitionsCreateRevisionOptions *ModelDefinitionsCreateRevisionOptions) (result *ModelDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsCreateRevisionOptions, "modelDefinitionsCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsCreateRevisionOptions, "modelDefinitionsCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_definition_id": *modelDefinitionsCreateRevisionOptions.ModelDefinitionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions/{model_definition_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if modelDefinitionsCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = modelDefinitionsCreateRevisionOptions.SpaceID
	}
	if modelDefinitionsCreateRevisionOptions.ProjectID != nil {
		body["project_id"] = modelDefinitionsCreateRevisionOptions.ProjectID
	}
	if modelDefinitionsCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = modelDefinitionsCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelDefinitionsListRevisions : Retrieve the model definition revisions
// Retrieve the model definition revisions. This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsListRevisions(modelDefinitionsListRevisionsOptions *ModelDefinitionsListRevisionsOptions) (result *ModelDefinitionResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsListRevisionsWithContext(context.Background(), modelDefinitionsListRevisionsOptions)
}

// ModelDefinitionsListRevisionsWithContext is an alternate form of the ModelDefinitionsListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsListRevisionsWithContext(ctx context.Context, modelDefinitionsListRevisionsOptions *ModelDefinitionsListRevisionsOptions) (result *ModelDefinitionResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsListRevisionsOptions, "modelDefinitionsListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsListRevisionsOptions, "modelDefinitionsListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_definition_id": *modelDefinitionsListRevisionsOptions.ModelDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions/{model_definition_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelDefinitionsListRevisionsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelDefinitionsListRevisionsOptions.SpaceID))
	}
	if modelDefinitionsListRevisionsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelDefinitionsListRevisionsOptions.ProjectID))
	}
	if modelDefinitionsListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*modelDefinitionsListRevisionsOptions.Start))
	}
	if modelDefinitionsListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*modelDefinitionsListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalModelDefinitionResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelDefinitionsUploadModel : Upload the model definition model
// Upload the model definition model. Model definitions for Deep Learning accept a zip file that contains one or more
//
//	python files organized in any directory structure. This command is supported starting with release 3.5 of Cloud Pak
//
// for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsUploadModel(modelDefinitionsUploadModelOptions *ModelDefinitionsUploadModelOptions) (result *ContentMetadata, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsUploadModelWithContext(context.Background(), modelDefinitionsUploadModelOptions)
}

// ModelDefinitionsUploadModelWithContext is an alternate form of the ModelDefinitionsUploadModel method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsUploadModelWithContext(ctx context.Context, modelDefinitionsUploadModelOptions *ModelDefinitionsUploadModelOptions) (result *ContentMetadata, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsUploadModelOptions, "modelDefinitionsUploadModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsUploadModelOptions, "modelDefinitionsUploadModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_definition_id": *modelDefinitionsUploadModelOptions.ModelDefinitionID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions/{model_definition_id}/model`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsUploadModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsUploadModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/gzip")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelDefinitionsUploadModelOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelDefinitionsUploadModelOptions.SpaceID))
	}
	if modelDefinitionsUploadModelOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelDefinitionsUploadModelOptions.ProjectID))
	}

	_, err = builder.SetBodyContent("application/gzip", nil, nil, modelDefinitionsUploadModelOptions.UploadModel)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalContentMetadata)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ModelDefinitionsDownloadModel : Download the model definition model
// Download the model definition model. It is possible to get the `model` for a given revision of the `model
// definition`. Model definitions for Deep Learning accept a zip file that contains one or more     python files
// organized in any directory structure. This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsDownloadModel(modelDefinitionsDownloadModelOptions *ModelDefinitionsDownloadModelOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.ModelDefinitionsDownloadModelWithContext(context.Background(), modelDefinitionsDownloadModelOptions)
}

// ModelDefinitionsDownloadModelWithContext is an alternate form of the ModelDefinitionsDownloadModel method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) ModelDefinitionsDownloadModelWithContext(ctx context.Context, modelDefinitionsDownloadModelOptions *ModelDefinitionsDownloadModelOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(modelDefinitionsDownloadModelOptions, "modelDefinitionsDownloadModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(modelDefinitionsDownloadModelOptions, "modelDefinitionsDownloadModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"model_definition_id": *modelDefinitionsDownloadModelOptions.ModelDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/model_definitions/{model_definition_id}/model`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range modelDefinitionsDownloadModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "ModelDefinitionsDownloadModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/zip")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if modelDefinitionsDownloadModelOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*modelDefinitionsDownloadModelOptions.SpaceID))
	}
	if modelDefinitionsDownloadModelOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*modelDefinitionsDownloadModelOptions.ProjectID))
	}
	if modelDefinitionsDownloadModelOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*modelDefinitionsDownloadModelOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, &result)

	return
}

// PipelinesCreate : Create a new pipeline
// Create a new pipeline with the given payload. A pipeline represents a hybrid-pipeline, as a JSON document, that is
// used to train one or more models.
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesCreate(pipelinesCreateOptions *PipelinesCreateOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.PipelinesCreateWithContext(context.Background(), pipelinesCreateOptions)
}

// PipelinesCreateWithContext is an alternate form of the PipelinesCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesCreateWithContext(ctx context.Context, pipelinesCreateOptions *PipelinesCreateOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pipelinesCreateOptions, "pipelinesCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(pipelinesCreateOptions, "pipelinesCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/pipelines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range pipelinesCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "PipelinesCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if pipelinesCreateOptions.Name != nil {
		body["name"] = pipelinesCreateOptions.Name
	}
	if pipelinesCreateOptions.Document != nil {
		body["document"] = pipelinesCreateOptions.Document
	}
	if pipelinesCreateOptions.ProjectID != nil {
		body["project_id"] = pipelinesCreateOptions.ProjectID
	}
	if pipelinesCreateOptions.SpaceID != nil {
		body["space_id"] = pipelinesCreateOptions.SpaceID
	}
	if pipelinesCreateOptions.Description != nil {
		body["description"] = pipelinesCreateOptions.Description
	}
	if pipelinesCreateOptions.Tags != nil {
		body["tags"] = pipelinesCreateOptions.Tags
	}
	if pipelinesCreateOptions.Custom != nil {
		body["custom"] = pipelinesCreateOptions.Custom
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPipelineResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PipelinesList : Retrieve the pipelines
// Retrieve the pipelines for the specified space or project.
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesList(pipelinesListOptions *PipelinesListOptions) (result *PipelineResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.PipelinesListWithContext(context.Background(), pipelinesListOptions)
}

// PipelinesListWithContext is an alternate form of the PipelinesList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesListWithContext(ctx context.Context, pipelinesListOptions *PipelinesListOptions) (result *PipelineResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(pipelinesListOptions, "pipelinesListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/pipelines`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range pipelinesListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "PipelinesList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if pipelinesListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*pipelinesListOptions.SpaceID))
	}
	if pipelinesListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*pipelinesListOptions.ProjectID))
	}
	if pipelinesListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*pipelinesListOptions.Start))
	}
	if pipelinesListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*pipelinesListOptions.Limit))
	}
	if pipelinesListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*pipelinesListOptions.TagValue))
	}
	if pipelinesListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*pipelinesListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPipelineResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PipelinesGet : Retrieve the pipeline
// Retrieve the pipeline with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesGet(pipelinesGetOptions *PipelinesGetOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.PipelinesGetWithContext(context.Background(), pipelinesGetOptions)
}

// PipelinesGetWithContext is an alternate form of the PipelinesGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesGetWithContext(ctx context.Context, pipelinesGetOptions *PipelinesGetOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pipelinesGetOptions, "pipelinesGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(pipelinesGetOptions, "pipelinesGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"pipeline_id": *pipelinesGetOptions.PipelineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/pipelines/{pipeline_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range pipelinesGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "PipelinesGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if pipelinesGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*pipelinesGetOptions.SpaceID))
	}
	if pipelinesGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*pipelinesGetOptions.ProjectID))
	}
	if pipelinesGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*pipelinesGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPipelineResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PipelinesUpdate : Update the pipeline
// Update the pipeline with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`.
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesUpdate(pipelinesUpdateOptions *PipelinesUpdateOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.PipelinesUpdateWithContext(context.Background(), pipelinesUpdateOptions)
}

// PipelinesUpdateWithContext is an alternate form of the PipelinesUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesUpdateWithContext(ctx context.Context, pipelinesUpdateOptions *PipelinesUpdateOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pipelinesUpdateOptions, "pipelinesUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(pipelinesUpdateOptions, "pipelinesUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"pipeline_id": *pipelinesUpdateOptions.PipelineID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/pipelines/{pipeline_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range pipelinesUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "PipelinesUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if pipelinesUpdateOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*pipelinesUpdateOptions.SpaceID))
	}
	if pipelinesUpdateOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*pipelinesUpdateOptions.ProjectID))
	}

	_, err = builder.SetBodyContentJSON(pipelinesUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPipelineResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PipelinesDelete : Delete the pipeline
// Delete the pipeline with the specified identifier. This will delete all revisions of this pipeline as well. For each
// revision all attachments will also be deleted.
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesDelete(pipelinesDeleteOptions *PipelinesDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.PipelinesDeleteWithContext(context.Background(), pipelinesDeleteOptions)
}

// PipelinesDeleteWithContext is an alternate form of the PipelinesDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesDeleteWithContext(ctx context.Context, pipelinesDeleteOptions *PipelinesDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pipelinesDeleteOptions, "pipelinesDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(pipelinesDeleteOptions, "pipelinesDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"pipeline_id": *pipelinesDeleteOptions.PipelineID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/pipelines/{pipeline_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range pipelinesDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "PipelinesDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if pipelinesDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*pipelinesDeleteOptions.SpaceID))
	}
	if pipelinesDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*pipelinesDeleteOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// PipelinesCreateRevision : Create a new pipeline revision
// Create a new pipeline revision. The current metadata and content for pipeline_id will be taken and a new revision
// created. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesCreateRevision(pipelinesCreateRevisionOptions *PipelinesCreateRevisionOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.PipelinesCreateRevisionWithContext(context.Background(), pipelinesCreateRevisionOptions)
}

// PipelinesCreateRevisionWithContext is an alternate form of the PipelinesCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesCreateRevisionWithContext(ctx context.Context, pipelinesCreateRevisionOptions *PipelinesCreateRevisionOptions) (result *PipelineResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pipelinesCreateRevisionOptions, "pipelinesCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(pipelinesCreateRevisionOptions, "pipelinesCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"pipeline_id": *pipelinesCreateRevisionOptions.PipelineID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/pipelines/{pipeline_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range pipelinesCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "PipelinesCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if pipelinesCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = pipelinesCreateRevisionOptions.SpaceID
	}
	if pipelinesCreateRevisionOptions.ProjectID != nil {
		body["project_id"] = pipelinesCreateRevisionOptions.ProjectID
	}
	if pipelinesCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = pipelinesCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPipelineResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PipelinesListRevisions : Retrieve the pipeline revisions
// Retrieve the pipeline revisions.
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesListRevisions(pipelinesListRevisionsOptions *PipelinesListRevisionsOptions) (result *PipelineResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.PipelinesListRevisionsWithContext(context.Background(), pipelinesListRevisionsOptions)
}

// PipelinesListRevisionsWithContext is an alternate form of the PipelinesListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) PipelinesListRevisionsWithContext(ctx context.Context, pipelinesListRevisionsOptions *PipelinesListRevisionsOptions) (result *PipelineResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(pipelinesListRevisionsOptions, "pipelinesListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(pipelinesListRevisionsOptions, "pipelinesListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"pipeline_id": *pipelinesListRevisionsOptions.PipelineID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/pipelines/{pipeline_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range pipelinesListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "PipelinesListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if pipelinesListRevisionsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*pipelinesListRevisionsOptions.SpaceID))
	}
	if pipelinesListRevisionsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*pipelinesListRevisionsOptions.ProjectID))
	}
	if pipelinesListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*pipelinesListRevisionsOptions.Start))
	}
	if pipelinesListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*pipelinesListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPipelineResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoteTrainingSystemsCreate : Create a new remote training system
// Create a new remote training system with the given payload. A remote training system represents a remote system used
// by Federated Learning. This definition is necessary to control who can register to a federated learning job.
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsCreate(remoteTrainingSystemsCreateOptions *RemoteTrainingSystemsCreateOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.RemoteTrainingSystemsCreateWithContext(context.Background(), remoteTrainingSystemsCreateOptions)
}

// RemoteTrainingSystemsCreateWithContext is an alternate form of the RemoteTrainingSystemsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsCreateWithContext(ctx context.Context, remoteTrainingSystemsCreateOptions *RemoteTrainingSystemsCreateOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(remoteTrainingSystemsCreateOptions, "remoteTrainingSystemsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(remoteTrainingSystemsCreateOptions, "remoteTrainingSystemsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/remote_training_systems`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range remoteTrainingSystemsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "RemoteTrainingSystemsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if remoteTrainingSystemsCreateOptions.Name != nil {
		body["name"] = remoteTrainingSystemsCreateOptions.Name
	}
	if remoteTrainingSystemsCreateOptions.AllowedIdentities != nil {
		body["allowed_identities"] = remoteTrainingSystemsCreateOptions.AllowedIdentities
	}
	if remoteTrainingSystemsCreateOptions.ProjectID != nil {
		body["project_id"] = remoteTrainingSystemsCreateOptions.ProjectID
	}
	if remoteTrainingSystemsCreateOptions.SpaceID != nil {
		body["space_id"] = remoteTrainingSystemsCreateOptions.SpaceID
	}
	if remoteTrainingSystemsCreateOptions.Description != nil {
		body["description"] = remoteTrainingSystemsCreateOptions.Description
	}
	if remoteTrainingSystemsCreateOptions.Tags != nil {
		body["tags"] = remoteTrainingSystemsCreateOptions.Tags
	}
	if remoteTrainingSystemsCreateOptions.Organization != nil {
		body["organization"] = remoteTrainingSystemsCreateOptions.Organization
	}
	if remoteTrainingSystemsCreateOptions.RemoteAdmin != nil {
		body["remote_admin"] = remoteTrainingSystemsCreateOptions.RemoteAdmin
	}
	if remoteTrainingSystemsCreateOptions.Custom != nil {
		body["custom"] = remoteTrainingSystemsCreateOptions.Custom
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRemoteTrainingSystemResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoteTrainingSystemsList : Retrieve the remote training systems
// Retrieve the remote training systems for the specified space or project.
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsList(remoteTrainingSystemsListOptions *RemoteTrainingSystemsListOptions) (result *RemoteTrainingSystemResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.RemoteTrainingSystemsListWithContext(context.Background(), remoteTrainingSystemsListOptions)
}

// RemoteTrainingSystemsListWithContext is an alternate form of the RemoteTrainingSystemsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsListWithContext(ctx context.Context, remoteTrainingSystemsListOptions *RemoteTrainingSystemsListOptions) (result *RemoteTrainingSystemResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(remoteTrainingSystemsListOptions, "remoteTrainingSystemsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/remote_training_systems`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range remoteTrainingSystemsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "RemoteTrainingSystemsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if remoteTrainingSystemsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*remoteTrainingSystemsListOptions.SpaceID))
	}
	if remoteTrainingSystemsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*remoteTrainingSystemsListOptions.ProjectID))
	}
	if remoteTrainingSystemsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*remoteTrainingSystemsListOptions.Start))
	}
	if remoteTrainingSystemsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*remoteTrainingSystemsListOptions.Limit))
	}
	if remoteTrainingSystemsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*remoteTrainingSystemsListOptions.TagValue))
	}
	if remoteTrainingSystemsListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*remoteTrainingSystemsListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRemoteTrainingSystemResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoteTrainingSystemsGet : Retrieve the remote training system
// Retrieve the remote training system with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsGet(remoteTrainingSystemsGetOptions *RemoteTrainingSystemsGetOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.RemoteTrainingSystemsGetWithContext(context.Background(), remoteTrainingSystemsGetOptions)
}

// RemoteTrainingSystemsGetWithContext is an alternate form of the RemoteTrainingSystemsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsGetWithContext(ctx context.Context, remoteTrainingSystemsGetOptions *RemoteTrainingSystemsGetOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(remoteTrainingSystemsGetOptions, "remoteTrainingSystemsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(remoteTrainingSystemsGetOptions, "remoteTrainingSystemsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"remote_training_system_id": *remoteTrainingSystemsGetOptions.RemoteTrainingSystemID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/remote_training_systems/{remote_training_system_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range remoteTrainingSystemsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "RemoteTrainingSystemsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if remoteTrainingSystemsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*remoteTrainingSystemsGetOptions.SpaceID))
	}
	if remoteTrainingSystemsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*remoteTrainingSystemsGetOptions.ProjectID))
	}
	if remoteTrainingSystemsGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*remoteTrainingSystemsGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRemoteTrainingSystemResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoteTrainingSystemsUpdate : Update the remote training system
// Update the remote training system with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`
// - `/organization`
// - `/allowed_identities`
// - `/remote_admin`.
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsUpdate(remoteTrainingSystemsUpdateOptions *RemoteTrainingSystemsUpdateOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.RemoteTrainingSystemsUpdateWithContext(context.Background(), remoteTrainingSystemsUpdateOptions)
}

// RemoteTrainingSystemsUpdateWithContext is an alternate form of the RemoteTrainingSystemsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsUpdateWithContext(ctx context.Context, remoteTrainingSystemsUpdateOptions *RemoteTrainingSystemsUpdateOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(remoteTrainingSystemsUpdateOptions, "remoteTrainingSystemsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(remoteTrainingSystemsUpdateOptions, "remoteTrainingSystemsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"remote_training_system_id": *remoteTrainingSystemsUpdateOptions.RemoteTrainingSystemID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/remote_training_systems/{remote_training_system_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range remoteTrainingSystemsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "RemoteTrainingSystemsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if remoteTrainingSystemsUpdateOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*remoteTrainingSystemsUpdateOptions.SpaceID))
	}
	if remoteTrainingSystemsUpdateOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*remoteTrainingSystemsUpdateOptions.ProjectID))
	}

	_, err = builder.SetBodyContentJSON(remoteTrainingSystemsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRemoteTrainingSystemResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoteTrainingSystemsDelete : Delete the remote training system
// Delete the remote training system with the specified identifier. This will delete all revisions of this remote
// training system as well. For each revision all attachments will also be deleted.
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsDelete(remoteTrainingSystemsDeleteOptions *RemoteTrainingSystemsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.RemoteTrainingSystemsDeleteWithContext(context.Background(), remoteTrainingSystemsDeleteOptions)
}

// RemoteTrainingSystemsDeleteWithContext is an alternate form of the RemoteTrainingSystemsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsDeleteWithContext(ctx context.Context, remoteTrainingSystemsDeleteOptions *RemoteTrainingSystemsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(remoteTrainingSystemsDeleteOptions, "remoteTrainingSystemsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(remoteTrainingSystemsDeleteOptions, "remoteTrainingSystemsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"remote_training_system_id": *remoteTrainingSystemsDeleteOptions.RemoteTrainingSystemID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/remote_training_systems/{remote_training_system_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range remoteTrainingSystemsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "RemoteTrainingSystemsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if remoteTrainingSystemsDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*remoteTrainingSystemsDeleteOptions.SpaceID))
	}
	if remoteTrainingSystemsDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*remoteTrainingSystemsDeleteOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// RemoteTrainingSystemsCreateRevision : Create a new remote training system revision
// Create a new remote training system revision. The current metadata and content for remote_training_system_id will be
// taken and a new revision created. Either `space_id` or `project_id` has to be provided and is mandatory.
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsCreateRevision(remoteTrainingSystemsCreateRevisionOptions *RemoteTrainingSystemsCreateRevisionOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.RemoteTrainingSystemsCreateRevisionWithContext(context.Background(), remoteTrainingSystemsCreateRevisionOptions)
}

// RemoteTrainingSystemsCreateRevisionWithContext is an alternate form of the RemoteTrainingSystemsCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsCreateRevisionWithContext(ctx context.Context, remoteTrainingSystemsCreateRevisionOptions *RemoteTrainingSystemsCreateRevisionOptions) (result *RemoteTrainingSystemResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(remoteTrainingSystemsCreateRevisionOptions, "remoteTrainingSystemsCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(remoteTrainingSystemsCreateRevisionOptions, "remoteTrainingSystemsCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"remote_training_system_id": *remoteTrainingSystemsCreateRevisionOptions.RemoteTrainingSystemID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/remote_training_systems/{remote_training_system_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range remoteTrainingSystemsCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "RemoteTrainingSystemsCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if remoteTrainingSystemsCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = remoteTrainingSystemsCreateRevisionOptions.SpaceID
	}
	if remoteTrainingSystemsCreateRevisionOptions.ProjectID != nil {
		body["project_id"] = remoteTrainingSystemsCreateRevisionOptions.ProjectID
	}
	if remoteTrainingSystemsCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = remoteTrainingSystemsCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRemoteTrainingSystemResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoteTrainingSystemsListRevisions : Retrieve the remote training system revisions
// Retrieve the remote training system revisions.
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsListRevisions(remoteTrainingSystemsListRevisionsOptions *RemoteTrainingSystemsListRevisionsOptions) (result *RemoteTrainingSystemResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.RemoteTrainingSystemsListRevisionsWithContext(context.Background(), remoteTrainingSystemsListRevisionsOptions)
}

// RemoteTrainingSystemsListRevisionsWithContext is an alternate form of the RemoteTrainingSystemsListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) RemoteTrainingSystemsListRevisionsWithContext(ctx context.Context, remoteTrainingSystemsListRevisionsOptions *RemoteTrainingSystemsListRevisionsOptions) (result *RemoteTrainingSystemResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(remoteTrainingSystemsListRevisionsOptions, "remoteTrainingSystemsListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(remoteTrainingSystemsListRevisionsOptions, "remoteTrainingSystemsListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"remote_training_system_id": *remoteTrainingSystemsListRevisionsOptions.RemoteTrainingSystemID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/remote_training_systems/{remote_training_system_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range remoteTrainingSystemsListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "RemoteTrainingSystemsListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if remoteTrainingSystemsListRevisionsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*remoteTrainingSystemsListRevisionsOptions.SpaceID))
	}
	if remoteTrainingSystemsListRevisionsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*remoteTrainingSystemsListRevisionsOptions.ProjectID))
	}
	if remoteTrainingSystemsListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*remoteTrainingSystemsListRevisionsOptions.Start))
	}
	if remoteTrainingSystemsListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*remoteTrainingSystemsListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRemoteTrainingSystemResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingsCreate : Create a new WML training
// Create a new WML training.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsCreate(trainingsCreateOptions *TrainingsCreateOptions) (result *TrainingResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingsCreateWithContext(context.Background(), trainingsCreateOptions)
}

// TrainingsCreateWithContext is an alternate form of the TrainingsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsCreateWithContext(ctx context.Context, trainingsCreateOptions *TrainingsCreateOptions) (result *TrainingResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingsCreateOptions, "trainingsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingsCreateOptions, "trainingsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/trainings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if trainingsCreateOptions.ResultsReference != nil {
		body["results_reference"] = trainingsCreateOptions.ResultsReference
	}
	if trainingsCreateOptions.Experiment != nil {
		body["experiment"] = trainingsCreateOptions.Experiment
	}
	if trainingsCreateOptions.Pipeline != nil {
		body["pipeline"] = trainingsCreateOptions.Pipeline
	}
	if trainingsCreateOptions.ModelDefinition != nil {
		body["model_definition"] = trainingsCreateOptions.ModelDefinition
	}
	if trainingsCreateOptions.FederatedLearning != nil {
		body["federated_learning"] = trainingsCreateOptions.FederatedLearning
	}
	if trainingsCreateOptions.TrainingDataReferences != nil {
		body["training_data_references"] = trainingsCreateOptions.TrainingDataReferences
	}
	if trainingsCreateOptions.TestDataReferences != nil {
		body["test_data_references"] = trainingsCreateOptions.TestDataReferences
	}
	if trainingsCreateOptions.Custom != nil {
		body["custom"] = trainingsCreateOptions.Custom
	}
	if trainingsCreateOptions.Tags != nil {
		body["tags"] = trainingsCreateOptions.Tags
	}
	if trainingsCreateOptions.Name != nil {
		body["name"] = trainingsCreateOptions.Name
	}
	if trainingsCreateOptions.Description != nil {
		body["description"] = trainingsCreateOptions.Description
	}
	if trainingsCreateOptions.SpaceID != nil {
		body["space_id"] = trainingsCreateOptions.SpaceID
	}
	if trainingsCreateOptions.ProjectID != nil {
		body["project_id"] = trainingsCreateOptions.ProjectID
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingsList : Retrieve the list of trainings
// Retrieve the list of trainings for the specified space or project.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsList(trainingsListOptions *TrainingsListOptions) (result *TrainingResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingsListWithContext(context.Background(), trainingsListOptions)
}

// TrainingsListWithContext is an alternate form of the TrainingsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsListWithContext(ctx context.Context, trainingsListOptions *TrainingsListOptions) (result *TrainingResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(trainingsListOptions, "trainingsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/trainings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*trainingsListOptions.Start))
	}
	if trainingsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*trainingsListOptions.Limit))
	}
	if trainingsListOptions.TotalCount != nil {
		builder.AddQuery("total_count", fmt.Sprint(*trainingsListOptions.TotalCount))
	}
	if trainingsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*trainingsListOptions.TagValue))
	}
	if trainingsListOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*trainingsListOptions.Type))
	}
	if trainingsListOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*trainingsListOptions.State))
	}
	if trainingsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingsListOptions.SpaceID))
	}
	if trainingsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingsListOptions.ProjectID))
	}
	if trainingsListOptions.ParentID != nil {
		builder.AddQuery("parent_id", fmt.Sprint(*trainingsListOptions.ParentID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingsGet : Retrieve the training
// Retrieve the training with the specified identifier. This call supports Web-Socket upgrade. However in order to
// preserve bandwidth, web-socket messages are not context complete. Meaning that a single web-socket message only
// reflects a message or metric happening in the context of a training job or sub-job (in case of experiment trainings
// or HPO/AutoML trainings). Hence the metadata property of a web-socket message contains a parent with the href
// information of the parent job that triggered this particular job. Also the metrics will be provided as they arrive
// from the backend runtime, and not as a cumulative list. In order to get the full view of the running training job the
// caller should do a regular GET call.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsGet(trainingsGetOptions *TrainingsGetOptions) (result *TrainingResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingsGetWithContext(context.Background(), trainingsGetOptions)
}

// TrainingsGetWithContext is an alternate form of the TrainingsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsGetWithContext(ctx context.Context, trainingsGetOptions *TrainingsGetOptions) (result *TrainingResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingsGetOptions, "trainingsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingsGetOptions, "trainingsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"training_id": *trainingsGetOptions.TrainingID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/trainings/{training_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingsGetOptions.SpaceID))
	}
	if trainingsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingsGetOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingsDelete : Cancel the training
// Cancel the specified training and remove it.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsDelete(trainingsDeleteOptions *TrainingsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingsDeleteWithContext(context.Background(), trainingsDeleteOptions)
}

// TrainingsDeleteWithContext is an alternate form of the TrainingsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingsDeleteWithContext(ctx context.Context, trainingsDeleteOptions *TrainingsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingsDeleteOptions, "trainingsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingsDeleteOptions, "trainingsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"training_id": *trainingsDeleteOptions.TrainingID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/trainings/{training_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingsDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingsDeleteOptions.SpaceID))
	}
	if trainingsDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingsDeleteOptions.ProjectID))
	}
	if trainingsDeleteOptions.HardDelete != nil {
		builder.AddQuery("hard_delete", fmt.Sprint(*trainingsDeleteOptions.HardDelete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// TrainingDefinitionsCreate : Create a new training definition
// Create a new training definition with the given payload. A training definition represents the training meta-data
// necessary to start a training job. This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsCreate(trainingDefinitionsCreateOptions *TrainingDefinitionsCreateOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingDefinitionsCreateWithContext(context.Background(), trainingDefinitionsCreateOptions)
}

// TrainingDefinitionsCreateWithContext is an alternate form of the TrainingDefinitionsCreate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsCreateWithContext(ctx context.Context, trainingDefinitionsCreateOptions *TrainingDefinitionsCreateOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingDefinitionsCreateOptions, "trainingDefinitionsCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingDefinitionsCreateOptions, "trainingDefinitionsCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/training_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingDefinitionsCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingDefinitionsCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if trainingDefinitionsCreateOptions.Name != nil {
		body["name"] = trainingDefinitionsCreateOptions.Name
	}
	if trainingDefinitionsCreateOptions.ResultsReference != nil {
		body["results_reference"] = trainingDefinitionsCreateOptions.ResultsReference
	}
	if trainingDefinitionsCreateOptions.ProjectID != nil {
		body["project_id"] = trainingDefinitionsCreateOptions.ProjectID
	}
	if trainingDefinitionsCreateOptions.SpaceID != nil {
		body["space_id"] = trainingDefinitionsCreateOptions.SpaceID
	}
	if trainingDefinitionsCreateOptions.Description != nil {
		body["description"] = trainingDefinitionsCreateOptions.Description
	}
	if trainingDefinitionsCreateOptions.Tags != nil {
		body["tags"] = trainingDefinitionsCreateOptions.Tags
	}
	if trainingDefinitionsCreateOptions.Experiment != nil {
		body["experiment"] = trainingDefinitionsCreateOptions.Experiment
	}
	if trainingDefinitionsCreateOptions.Pipeline != nil {
		body["pipeline"] = trainingDefinitionsCreateOptions.Pipeline
	}
	if trainingDefinitionsCreateOptions.ModelDefinition != nil {
		body["model_definition"] = trainingDefinitionsCreateOptions.ModelDefinition
	}
	if trainingDefinitionsCreateOptions.FederatedLearning != nil {
		body["federated_learning"] = trainingDefinitionsCreateOptions.FederatedLearning
	}
	if trainingDefinitionsCreateOptions.TrainingDataReferences != nil {
		body["training_data_references"] = trainingDefinitionsCreateOptions.TrainingDataReferences
	}
	if trainingDefinitionsCreateOptions.TestDataReferences != nil {
		body["test_data_references"] = trainingDefinitionsCreateOptions.TestDataReferences
	}
	if trainingDefinitionsCreateOptions.Custom != nil {
		body["custom"] = trainingDefinitionsCreateOptions.Custom
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingDefinitionsList : Retrieve the training definitions
// Retrieve the training definitions for the specified space or project. This command is supported starting with release
// 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsList(trainingDefinitionsListOptions *TrainingDefinitionsListOptions) (result *TrainingDefinitionResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingDefinitionsListWithContext(context.Background(), trainingDefinitionsListOptions)
}

// TrainingDefinitionsListWithContext is an alternate form of the TrainingDefinitionsList method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsListWithContext(ctx context.Context, trainingDefinitionsListOptions *TrainingDefinitionsListOptions) (result *TrainingDefinitionResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(trainingDefinitionsListOptions, "trainingDefinitionsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/training_definitions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingDefinitionsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingDefinitionsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingDefinitionsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingDefinitionsListOptions.SpaceID))
	}
	if trainingDefinitionsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingDefinitionsListOptions.ProjectID))
	}
	if trainingDefinitionsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*trainingDefinitionsListOptions.Start))
	}
	if trainingDefinitionsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*trainingDefinitionsListOptions.Limit))
	}
	if trainingDefinitionsListOptions.TagValue != nil {
		builder.AddQuery("tag.value", fmt.Sprint(*trainingDefinitionsListOptions.TagValue))
	}
	if trainingDefinitionsListOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*trainingDefinitionsListOptions.Search))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDefinitionResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingDefinitionsGet : Retrieve the training definition
// Retrieve the training definition with the specified identifier. If `rev` query parameter is provided,
// `rev=latest` will fetch the latest revision. A call with `rev={revision_number}` will fetch the given revision_number
// record. Either `space_id` or `project_id` has to be provided and is mandatory. This command is supported starting
// with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsGet(trainingDefinitionsGetOptions *TrainingDefinitionsGetOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingDefinitionsGetWithContext(context.Background(), trainingDefinitionsGetOptions)
}

// TrainingDefinitionsGetWithContext is an alternate form of the TrainingDefinitionsGet method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsGetWithContext(ctx context.Context, trainingDefinitionsGetOptions *TrainingDefinitionsGetOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingDefinitionsGetOptions, "trainingDefinitionsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingDefinitionsGetOptions, "trainingDefinitionsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"training_definition_id": *trainingDefinitionsGetOptions.TrainingDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/training_definitions/{training_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingDefinitionsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingDefinitionsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingDefinitionsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingDefinitionsGetOptions.SpaceID))
	}
	if trainingDefinitionsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingDefinitionsGetOptions.ProjectID))
	}
	if trainingDefinitionsGetOptions.Rev != nil {
		builder.AddQuery("rev", fmt.Sprint(*trainingDefinitionsGetOptions.Rev))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingDefinitionsUpdate : Update the training definition
// Update the training definition with the provided patch data. The following fields can be patched:
// - `/tags`
// - `/name`
// - `/description`
// - `/custom`
// - `/federated_learning` This command is supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsUpdate(trainingDefinitionsUpdateOptions *TrainingDefinitionsUpdateOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingDefinitionsUpdateWithContext(context.Background(), trainingDefinitionsUpdateOptions)
}

// TrainingDefinitionsUpdateWithContext is an alternate form of the TrainingDefinitionsUpdate method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsUpdateWithContext(ctx context.Context, trainingDefinitionsUpdateOptions *TrainingDefinitionsUpdateOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingDefinitionsUpdateOptions, "trainingDefinitionsUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingDefinitionsUpdateOptions, "trainingDefinitionsUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"training_definition_id": *trainingDefinitionsUpdateOptions.TrainingDefinitionID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/training_definitions/{training_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingDefinitionsUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingDefinitionsUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingDefinitionsUpdateOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingDefinitionsUpdateOptions.SpaceID))
	}
	if trainingDefinitionsUpdateOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingDefinitionsUpdateOptions.ProjectID))
	}

	_, err = builder.SetBodyContentJSON(trainingDefinitionsUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingDefinitionsDelete : Delete the training definition
// Delete the training definition with the specified identifier. This will delete all revisions of this training
// definition as well. For each revision all attachments will also be deleted. This command is supported starting with
// release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsDelete(trainingDefinitionsDeleteOptions *TrainingDefinitionsDeleteOptions) (response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingDefinitionsDeleteWithContext(context.Background(), trainingDefinitionsDeleteOptions)
}

// TrainingDefinitionsDeleteWithContext is an alternate form of the TrainingDefinitionsDelete method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsDeleteWithContext(ctx context.Context, trainingDefinitionsDeleteOptions *TrainingDefinitionsDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingDefinitionsDeleteOptions, "trainingDefinitionsDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingDefinitionsDeleteOptions, "trainingDefinitionsDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"training_definition_id": *trainingDefinitionsDeleteOptions.TrainingDefinitionID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/training_definitions/{training_definition_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingDefinitionsDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingDefinitionsDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingDefinitionsDeleteOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingDefinitionsDeleteOptions.SpaceID))
	}
	if trainingDefinitionsDeleteOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingDefinitionsDeleteOptions.ProjectID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = watsonMachineLearning.Service.Request(request, nil)

	return
}

// TrainingDefinitionsCreateRevision : Create a new training definition revision
// Create a new training definition revision. The current metadata and content for training_definition_id will be taken
// and a new revision created. Either `space_id` or `project_id` has to be provided and is mandatory. This command is
// supported starting with release 3.5 of Cloud Pak for Data.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsCreateRevision(trainingDefinitionsCreateRevisionOptions *TrainingDefinitionsCreateRevisionOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingDefinitionsCreateRevisionWithContext(context.Background(), trainingDefinitionsCreateRevisionOptions)
}

// TrainingDefinitionsCreateRevisionWithContext is an alternate form of the TrainingDefinitionsCreateRevision method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsCreateRevisionWithContext(ctx context.Context, trainingDefinitionsCreateRevisionOptions *TrainingDefinitionsCreateRevisionOptions) (result *TrainingDefinitionResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingDefinitionsCreateRevisionOptions, "trainingDefinitionsCreateRevisionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingDefinitionsCreateRevisionOptions, "trainingDefinitionsCreateRevisionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"training_definition_id": *trainingDefinitionsCreateRevisionOptions.TrainingDefinitionID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/training_definitions/{training_definition_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingDefinitionsCreateRevisionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingDefinitionsCreateRevision")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))

	body := make(map[string]interface{})
	if trainingDefinitionsCreateRevisionOptions.SpaceID != nil {
		body["space_id"] = trainingDefinitionsCreateRevisionOptions.SpaceID
	}
	if trainingDefinitionsCreateRevisionOptions.ProjectID != nil {
		body["project_id"] = trainingDefinitionsCreateRevisionOptions.ProjectID
	}
	if trainingDefinitionsCreateRevisionOptions.CommitMessage != nil {
		body["commit_message"] = trainingDefinitionsCreateRevisionOptions.CommitMessage
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
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDefinitionResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// TrainingDefinitionsListRevisions : Retrieve the training definition revisions
// Retrieve the training definition revisions. This command is supported starting with release 3.5 of Cloud Pak for
// Data.
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsListRevisions(trainingDefinitionsListRevisionsOptions *TrainingDefinitionsListRevisionsOptions) (result *TrainingDefinitionResources, response *core.DetailedResponse, err error) {
	return watsonMachineLearning.TrainingDefinitionsListRevisionsWithContext(context.Background(), trainingDefinitionsListRevisionsOptions)
}

// TrainingDefinitionsListRevisionsWithContext is an alternate form of the TrainingDefinitionsListRevisions method which supports a Context parameter
func (watsonMachineLearning *WatsonMachineLearningV4) TrainingDefinitionsListRevisionsWithContext(ctx context.Context, trainingDefinitionsListRevisionsOptions *TrainingDefinitionsListRevisionsOptions) (result *TrainingDefinitionResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(trainingDefinitionsListRevisionsOptions, "trainingDefinitionsListRevisionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(trainingDefinitionsListRevisionsOptions, "trainingDefinitionsListRevisionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"training_definition_id": *trainingDefinitionsListRevisionsOptions.TrainingDefinitionID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = watsonMachineLearning.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(watsonMachineLearning.Service.Options.URL, `/v4/training_definitions/{training_definition_id}/revisions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range trainingDefinitionsListRevisionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("watson_machine_learning", "V4", "TrainingDefinitionsListRevisions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*watsonMachineLearning.Version))
	if trainingDefinitionsListRevisionsOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*trainingDefinitionsListRevisionsOptions.SpaceID))
	}
	if trainingDefinitionsListRevisionsOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*trainingDefinitionsListRevisionsOptions.ProjectID))
	}
	if trainingDefinitionsListRevisionsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*trainingDefinitionsListRevisionsOptions.Start))
	}
	if trainingDefinitionsListRevisionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*trainingDefinitionsListRevisionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = watsonMachineLearning.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTrainingDefinitionResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ConsumptionCapacityUnitHours : ConsumptionCapacityUnitHours struct
type ConsumptionCapacityUnitHours struct {
	// The expiration date of the instance limit.
	ExpirationDate *strfmt.Date `json:"expiration_date,omitempty"`

	// The current total computation time (in capacity unit hours). It is a sum of both, reserved and already send to BSS,
	// units.
	Current *float64 `json:"current,omitempty"`

	// The maximal computation time (in capacity unit hours).
	Limit *float64 `json:"limit,omitempty"`
}

// UnmarshalConsumptionCapacityUnitHours unmarshals an instance of ConsumptionCapacityUnitHours from the specified map of raw messages.
func UnmarshalConsumptionCapacityUnitHours(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConsumptionCapacityUnitHours)
	err = core.UnmarshalPrimitive(m, "expiration_date", &obj.ExpirationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "current", &obj.Current)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConsumptionDeploymentJobCount : Limit for deployment jobs.
type ConsumptionDeploymentJobCount struct {
	// The maximal number of deployment jobs.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalConsumptionDeploymentJobCount unmarshals an instance of ConsumptionDeploymentJobCount from the specified map of raw messages.
func UnmarshalConsumptionDeploymentJobCount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConsumptionDeploymentJobCount)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConsumptionDoJobCount : ConsumptionDoJobCount struct
type ConsumptionDoJobCount struct {
	// The maximal number of parallel DO jobs.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalConsumptionDoJobCount unmarshals an instance of ConsumptionDoJobCount from the specified map of raw messages.
func UnmarshalConsumptionDoJobCount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConsumptionDoJobCount)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConsumptionGpuCount : ConsumptionGpuCount struct
type ConsumptionGpuCount struct {
	// The current number of reserved GPUs.
	Current *int64 `json:"current,omitempty"`

	// The maximal number of reserved GPUs.
	Limit *int64 `json:"limit,omitempty"`
}

// UnmarshalConsumptionGpuCount unmarshals an instance of ConsumptionGpuCount from the specified map of raw messages.
func UnmarshalConsumptionGpuCount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConsumptionGpuCount)
	err = core.UnmarshalPrimitive(m, "current", &obj.Current)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityRequestBatch : Indicates that this is a batch deployment. An empty object has to be specified. More properties will be added later
// on to setup the batch deployment.
type DeploymentEntityRequestBatch struct {
	// A set of key-value pairs where `key` is the parameter name.
	Parameters interface{} `json:"parameters,omitempty"`
}

// UnmarshalDeploymentEntityRequestBatch unmarshals an instance of DeploymentEntityRequestBatch from the specified map of raw messages.
func UnmarshalDeploymentEntityRequestBatch(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityRequestBatch)
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityRequestOnline : Indicates that this is an online deployment. An empty object has to be specified. More properties will be added later
// on to setup the online deployment. The `serving_name` can be provided in the `online.parameters`. The serving name
// can only have the characters [a-z,0-9,_]  and the length should not be more than 36 characters. The `serving_name`
// can be used in the prediction URL in place of the `deployment_id`.
type DeploymentEntityRequestOnline struct {
	// A set of key-value pairs where `key` is the parameter name.
	Parameters interface{} `json:"parameters,omitempty"`
}

// UnmarshalDeploymentEntityRequestOnline unmarshals an instance of DeploymentEntityRequestOnline from the specified map of raw messages.
func UnmarshalDeploymentEntityRequestOnline(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityRequestOnline)
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityRequestRShiny : Indicates that this is a Shiny application deployment.
type DeploymentEntityRequestRShiny struct {
	// Specifies the type of users who can access the Shiny application.
	Authentication *string `json:"authentication" validate:"required"`

	// A set of parameters that specify details about the Shiny deployment.
	Parameters *DeploymentEntityRequestRShinyParameters `json:"parameters,omitempty"`
}

// Constants associated with the DeploymentEntityRequestRShiny.Authentication property.
// Specifies the type of users who can access the Shiny application.
const (
	DeploymentEntityRequestRShiny_Authentication_AnyValidUser             = "any_valid_user"
	DeploymentEntityRequestRShiny_Authentication_AnyoneWithURL            = "anyone_with_url"
	DeploymentEntityRequestRShiny_Authentication_MembersOfDeploymentSpace = "members_of_deployment_space"
)

// NewDeploymentEntityRequestRShiny : Instantiate DeploymentEntityRequestRShiny (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewDeploymentEntityRequestRShiny(authentication string) (_model *DeploymentEntityRequestRShiny, err error) {
	_model = &DeploymentEntityRequestRShiny{
		Authentication: core.StringPtr(authentication),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDeploymentEntityRequestRShiny unmarshals an instance of DeploymentEntityRequestRShiny from the specified map of raw messages.
func UnmarshalDeploymentEntityRequestRShiny(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityRequestRShiny)
	err = core.UnmarshalPrimitive(m, "authentication", &obj.Authentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalDeploymentEntityRequestRShinyParameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityRequestRShinyParameters : A set of parameters that specify details about the Shiny deployment.
type DeploymentEntityRequestRShinyParameters struct {
	// Specifies the unique `serving_name` that will be used to create a URL for the application. The creation of the
	// deployment will fail if the `serving_name` is not unique.
	ServingName *string `json:"serving_name,omitempty"`

	// Specifies the details when deploying with a code_package asset.
	CodePackage *DeploymentEntityRequestRShinyParametersCodePackage `json:"code_package,omitempty"`
}

// UnmarshalDeploymentEntityRequestRShinyParameters unmarshals an instance of DeploymentEntityRequestRShinyParameters from the specified map of raw messages.
func UnmarshalDeploymentEntityRequestRShinyParameters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityRequestRShinyParameters)
	err = core.UnmarshalPrimitive(m, "serving_name", &obj.ServingName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "code_package", &obj.CodePackage, UnmarshalDeploymentEntityRequestRShinyParametersCodePackage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityRequestRShinyParametersCodePackage : Specifies the details when deploying with a code_package asset.
type DeploymentEntityRequestRShinyParametersCodePackage struct {
	// The path to the application files in the code package.
	Path *string `json:"path" validate:"required"`
}

// NewDeploymentEntityRequestRShinyParametersCodePackage : Instantiate DeploymentEntityRequestRShinyParametersCodePackage (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewDeploymentEntityRequestRShinyParametersCodePackage(path string) (_model *DeploymentEntityRequestRShinyParametersCodePackage, err error) {
	_model = &DeploymentEntityRequestRShinyParametersCodePackage{
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDeploymentEntityRequestRShinyParametersCodePackage unmarshals an instance of DeploymentEntityRequestRShinyParametersCodePackage from the specified map of raw messages.
func UnmarshalDeploymentEntityRequestRShinyParametersCodePackage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityRequestRShinyParametersCodePackage)
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityStatus : Specifies the current status, additional information about the deployment and any failure messages in case of
// deployment failures.
type DeploymentEntityStatus struct {
	// Specifies the current state of the deployment.
	State *string `json:"state,omitempty"`

	// Optional messages related to the deployment.
	Message *Message `json:"message,omitempty"`

	// The data returned when an error is encountered.
	Failure *Error `json:"failure,omitempty"`

	// The URL's that can be used to submit online prediction API requests. These URL's will contain the
	// `deployment_id` and the `serving_name`, if the `serving_name` was set.
	ServingUrls []string `json:"serving_urls,omitempty"`

	// The model download URLs in the context of virtual deployments.
	VirtualDeploymentDownloads []DeploymentEntityStatusVirtualDeploymentDownloadsItem `json:"virtual_deployment_downloads,omitempty"`

	// Status information related to the state of the scaling, if scaling is in progress or has completed.
	Scaling *DeploymentEntityStatusScaling `json:"scaling,omitempty"`
}

// Constants associated with the DeploymentEntityStatus.State property.
// Specifies the current state of the deployment.
const (
	DeploymentEntityStatus_State_Failed       = "failed"
	DeploymentEntityStatus_State_Initializing = "initializing"
	DeploymentEntityStatus_State_Ready        = "ready"
	DeploymentEntityStatus_State_Updating     = "updating"
)

// UnmarshalDeploymentEntityStatus unmarshals an instance of DeploymentEntityStatus from the specified map of raw messages.
func UnmarshalDeploymentEntityStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityStatus)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "message", &obj.Message, UnmarshalMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "serving_urls", &obj.ServingUrls)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "virtual_deployment_downloads", &obj.VirtualDeploymentDownloads, UnmarshalDeploymentEntityStatusVirtualDeploymentDownloadsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scaling", &obj.Scaling, UnmarshalDeploymentEntityStatusScaling)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityStatusScaling : Status information related to the state of the scaling, if scaling is in progress or has completed.
type DeploymentEntityStatusScaling struct {
	// The state of the scaling request.
	State *string `json:"state" validate:"required"`

	// The time when the scaling was attempted.
	AttemptedAt *strfmt.DateTime `json:"attempted_at,omitempty"`

	// The number of requested replicas.
	RequestedReplicas *int64 `json:"requested_replicas,omitempty"`

	// The number of replicas currently deployed.
	DeployedReplicas *int64 `json:"deployed_replicas,omitempty"`

	// Optional messages related to the deployment.
	Message *Message `json:"message,omitempty"`
}

// Constants associated with the DeploymentEntityStatusScaling.State property.
// The state of the scaling request.
const (
	DeploymentEntityStatusScaling_State_Failed     = "failed"
	DeploymentEntityStatusScaling_State_InProgress = "in_progress"
	DeploymentEntityStatusScaling_State_Success    = "success"
)

// UnmarshalDeploymentEntityStatusScaling unmarshals an instance of DeploymentEntityStatusScaling from the specified map of raw messages.
func UnmarshalDeploymentEntityStatusScaling(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityStatusScaling)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attempted_at", &obj.AttemptedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requested_replicas", &obj.RequestedReplicas)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployed_replicas", &obj.DeployedReplicas)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "message", &obj.Message, UnmarshalMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntityStatusVirtualDeploymentDownloadsItem : The model download URL details.
type DeploymentEntityStatusVirtualDeploymentDownloadsItem struct {
	// The model download URL.
	URL *string `json:"url,omitempty"`
}

// UnmarshalDeploymentEntityStatusVirtualDeploymentDownloadsItem unmarshals an instance of DeploymentEntityStatusVirtualDeploymentDownloadsItem from the specified map of raw messages.
func UnmarshalDeploymentEntityStatusVirtualDeploymentDownloadsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntityStatusVirtualDeploymentDownloadsItem)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentJobDefinitionsCreateOptions : The DeploymentJobDefinitionsCreate options.
type DeploymentJobDefinitionsCreateOptions struct {
	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// A reference to a resource.
	Deployment *SimpleRel `json:"deployment" validate:"required"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// User defined properties.
	Custom interface{} `json:"custom,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityRequestHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`

	// Details about the input/output data and other properties to be used for a batch deployment job of a model,
	// Python Function or a Python Scripts.
	//
	// Use `input_data` property to specify the input data for batch processing as part of the job's payload.
	// The `input_data` property is mutually exclusive with `input_data_references` property, only use one of these.
	// When `input_data` is specified, the processed output of batch deployment job will be available in
	// `scoring.predictions`
	// parameter in the deployment job response. `input_data` property is not supported for batch deployment of Python
	// Scripts.
	//
	// Use `input_data_references` property to specify the details pertaining to the remote source where the input data for
	// batch deployment job is available. The `input_data_references` must be used with `output_data_references`.
	// The `input_data_references` property
	// is mutually exclusive with `input_data` property, only use one of these. The `input_data_references`
	// property is not supported for batch deployment job of Spark models and Python Functions.
	//
	// Use `output_data_references` property to specify the details pertaining to the remote source where the input
	// data for batch deployment job is available. `output_data_references` must be used with `input_data_references`.
	// The `output_data_references`
	// property is not supported for batch deployment job of Spark models and Python Functions.
	Scoring *JobScoringRequest `json:"scoring,omitempty"`

	// Details about the input/output data and other properties to be used for a batch
	// deployment job of a decision optimization problem.
	//
	// Use `input_data` property to specify the input data for batch processing as part
	// of the job's payload. The `input_data` property is mutually exclusive with `input_data_references`
	// property, only use one of these. When `input_data` is specified, the processed output
	// of batch deployment job will be available in `scoring.predictions` parameter in the deployment job response.
	//
	// Use `input_data_references` property to specify the details pertaining to the remote source
	// where the input data for batch deployment job is available. The `input_data_references` must be used
	// with `output_data_references`.
	// The `input_data_references` property is mutually exclusive with `input_data` property, only use one of these.
	//
	// Use `output_data_references` property to specify the details pertaining to the remote source where the
	// input data for batch deployment job is available. The `output_data_references` must be used with
	// `input_data_references`.
	DecisionOptimization *JobDecisionOptimizationRequest `json:"decision_optimization,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobDefinitionsCreateOptions : Instantiate DeploymentJobDefinitionsCreateOptions
func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionsCreateOptions(spaceID string, name string, deployment *SimpleRel) *DeploymentJobDefinitionsCreateOptions {
	return &DeploymentJobDefinitionsCreateOptions{
		SpaceID:    core.StringPtr(spaceID),
		Name:       core.StringPtr(name),
		Deployment: deployment,
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobDefinitionsCreateOptions) SetSpaceID(spaceID string) *DeploymentJobDefinitionsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *DeploymentJobDefinitionsCreateOptions) SetName(name string) *DeploymentJobDefinitionsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *DeploymentJobDefinitionsCreateOptions) SetDeployment(deployment *SimpleRel) *DeploymentJobDefinitionsCreateOptions {
	_options.Deployment = deployment
	return _options
}

// SetDescription : Allow user to set Description
func (_options *DeploymentJobDefinitionsCreateOptions) SetDescription(description string) *DeploymentJobDefinitionsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *DeploymentJobDefinitionsCreateOptions) SetTags(tags []string) *DeploymentJobDefinitionsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *DeploymentJobDefinitionsCreateOptions) SetCustom(custom interface{}) *DeploymentJobDefinitionsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHardwareSpec : Allow user to set HardwareSpec
func (_options *DeploymentJobDefinitionsCreateOptions) SetHardwareSpec(hardwareSpec *HardwareSpecRel) *DeploymentJobDefinitionsCreateOptions {
	_options.HardwareSpec = hardwareSpec
	return _options
}

// SetHybridPipelineHardwareSpecs : Allow user to set HybridPipelineHardwareSpecs
func (_options *DeploymentJobDefinitionsCreateOptions) SetHybridPipelineHardwareSpecs(hybridPipelineHardwareSpecs []JobEntityRequestHybridPipelineHardwareSpecsItem) *DeploymentJobDefinitionsCreateOptions {
	_options.HybridPipelineHardwareSpecs = hybridPipelineHardwareSpecs
	return _options
}

// SetScoring : Allow user to set Scoring
func (_options *DeploymentJobDefinitionsCreateOptions) SetScoring(scoring *JobScoringRequest) *DeploymentJobDefinitionsCreateOptions {
	_options.Scoring = scoring
	return _options
}

// SetDecisionOptimization : Allow user to set DecisionOptimization
func (_options *DeploymentJobDefinitionsCreateOptions) SetDecisionOptimization(decisionOptimization *JobDecisionOptimizationRequest) *DeploymentJobDefinitionsCreateOptions {
	_options.DecisionOptimization = decisionOptimization
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobDefinitionsCreateOptions) SetHeaders(param map[string]string) *DeploymentJobDefinitionsCreateOptions {
	options.Headers = param
	return options
}

// DeploymentJobDefinitionsCreateRevisionOptions : The DeploymentJobDefinitionsCreateRevision options.
type DeploymentJobDefinitionsCreateRevisionOptions struct {
	// Deployment job definition identifier.
	JobDefinitionID *string `json:"job_definition_id" validate:"required,ne="`

	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobDefinitionsCreateRevisionOptions : Instantiate DeploymentJobDefinitionsCreateRevisionOptions
func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionsCreateRevisionOptions(jobDefinitionID string, spaceID string) *DeploymentJobDefinitionsCreateRevisionOptions {
	return &DeploymentJobDefinitionsCreateRevisionOptions{
		JobDefinitionID: core.StringPtr(jobDefinitionID),
		SpaceID:         core.StringPtr(spaceID),
	}
}

// SetJobDefinitionID : Allow user to set JobDefinitionID
func (_options *DeploymentJobDefinitionsCreateRevisionOptions) SetJobDefinitionID(jobDefinitionID string) *DeploymentJobDefinitionsCreateRevisionOptions {
	_options.JobDefinitionID = core.StringPtr(jobDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobDefinitionsCreateRevisionOptions) SetSpaceID(spaceID string) *DeploymentJobDefinitionsCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *DeploymentJobDefinitionsCreateRevisionOptions) SetCommitMessage(commitMessage string) *DeploymentJobDefinitionsCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobDefinitionsCreateRevisionOptions) SetHeaders(param map[string]string) *DeploymentJobDefinitionsCreateRevisionOptions {
	options.Headers = param
	return options
}

// DeploymentJobDefinitionsDeleteOptions : The DeploymentJobDefinitionsDelete options.
type DeploymentJobDefinitionsDeleteOptions struct {
	// Deployment job definition identifier.
	JobDefinitionID *string `json:"job_definition_id" validate:"required,ne="`

	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobDefinitionsDeleteOptions : Instantiate DeploymentJobDefinitionsDeleteOptions
func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionsDeleteOptions(jobDefinitionID string, spaceID string) *DeploymentJobDefinitionsDeleteOptions {
	return &DeploymentJobDefinitionsDeleteOptions{
		JobDefinitionID: core.StringPtr(jobDefinitionID),
		SpaceID:         core.StringPtr(spaceID),
	}
}

// SetJobDefinitionID : Allow user to set JobDefinitionID
func (_options *DeploymentJobDefinitionsDeleteOptions) SetJobDefinitionID(jobDefinitionID string) *DeploymentJobDefinitionsDeleteOptions {
	_options.JobDefinitionID = core.StringPtr(jobDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobDefinitionsDeleteOptions) SetSpaceID(spaceID string) *DeploymentJobDefinitionsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobDefinitionsDeleteOptions) SetHeaders(param map[string]string) *DeploymentJobDefinitionsDeleteOptions {
	options.Headers = param
	return options
}

// DeploymentJobDefinitionsGetOptions : The DeploymentJobDefinitionsGet options.
type DeploymentJobDefinitionsGetOptions struct {
	// Deployment job definition identifier.
	JobDefinitionID *string `json:"job_definition_id" validate:"required,ne="`

	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobDefinitionsGetOptions : Instantiate DeploymentJobDefinitionsGetOptions
func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionsGetOptions(jobDefinitionID string, spaceID string) *DeploymentJobDefinitionsGetOptions {
	return &DeploymentJobDefinitionsGetOptions{
		JobDefinitionID: core.StringPtr(jobDefinitionID),
		SpaceID:         core.StringPtr(spaceID),
	}
}

// SetJobDefinitionID : Allow user to set JobDefinitionID
func (_options *DeploymentJobDefinitionsGetOptions) SetJobDefinitionID(jobDefinitionID string) *DeploymentJobDefinitionsGetOptions {
	_options.JobDefinitionID = core.StringPtr(jobDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobDefinitionsGetOptions) SetSpaceID(spaceID string) *DeploymentJobDefinitionsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *DeploymentJobDefinitionsGetOptions) SetRev(rev string) *DeploymentJobDefinitionsGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobDefinitionsGetOptions) SetHeaders(param map[string]string) *DeploymentJobDefinitionsGetOptions {
	options.Headers = param
	return options
}

// DeploymentJobDefinitionsListOptions : The DeploymentJobDefinitionsList options.
type DeploymentJobDefinitionsListOptions struct {
	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobDefinitionsListOptions : Instantiate DeploymentJobDefinitionsListOptions
func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionsListOptions(spaceID string) *DeploymentJobDefinitionsListOptions {
	return &DeploymentJobDefinitionsListOptions{
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobDefinitionsListOptions) SetSpaceID(spaceID string) *DeploymentJobDefinitionsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *DeploymentJobDefinitionsListOptions) SetStart(start string) *DeploymentJobDefinitionsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *DeploymentJobDefinitionsListOptions) SetLimit(limit int64) *DeploymentJobDefinitionsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *DeploymentJobDefinitionsListOptions) SetTagValue(tagValue string) *DeploymentJobDefinitionsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *DeploymentJobDefinitionsListOptions) SetSearch(search string) *DeploymentJobDefinitionsListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobDefinitionsListOptions) SetHeaders(param map[string]string) *DeploymentJobDefinitionsListOptions {
	options.Headers = param
	return options
}

// DeploymentJobDefinitionsListRevisionsOptions : The DeploymentJobDefinitionsListRevisions options.
type DeploymentJobDefinitionsListRevisionsOptions struct {
	// Deployment job definition identifier.
	JobDefinitionID *string `json:"job_definition_id" validate:"required,ne="`

	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobDefinitionsListRevisionsOptions : Instantiate DeploymentJobDefinitionsListRevisionsOptions
func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionsListRevisionsOptions(jobDefinitionID string, spaceID string) *DeploymentJobDefinitionsListRevisionsOptions {
	return &DeploymentJobDefinitionsListRevisionsOptions{
		JobDefinitionID: core.StringPtr(jobDefinitionID),
		SpaceID:         core.StringPtr(spaceID),
	}
}

// SetJobDefinitionID : Allow user to set JobDefinitionID
func (_options *DeploymentJobDefinitionsListRevisionsOptions) SetJobDefinitionID(jobDefinitionID string) *DeploymentJobDefinitionsListRevisionsOptions {
	_options.JobDefinitionID = core.StringPtr(jobDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobDefinitionsListRevisionsOptions) SetSpaceID(spaceID string) *DeploymentJobDefinitionsListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *DeploymentJobDefinitionsListRevisionsOptions) SetStart(start string) *DeploymentJobDefinitionsListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *DeploymentJobDefinitionsListRevisionsOptions) SetLimit(limit int64) *DeploymentJobDefinitionsListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobDefinitionsListRevisionsOptions) SetHeaders(param map[string]string) *DeploymentJobDefinitionsListRevisionsOptions {
	options.Headers = param
	return options
}

// DeploymentJobDefinitionsUpdateOptions : The DeploymentJobDefinitionsUpdate options.
type DeploymentJobDefinitionsUpdateOptions struct {
	// Deployment job definition identifier.
	JobDefinitionID *string `json:"job_definition_id" validate:"required,ne="`

	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobDefinitionsUpdateOptions : Instantiate DeploymentJobDefinitionsUpdateOptions
func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionsUpdateOptions(jobDefinitionID string, spaceID string, jsonPatch []JSONPatchOperation) *DeploymentJobDefinitionsUpdateOptions {
	return &DeploymentJobDefinitionsUpdateOptions{
		JobDefinitionID: core.StringPtr(jobDefinitionID),
		SpaceID:         core.StringPtr(spaceID),
		JSONPatch:       jsonPatch,
	}
}

// SetJobDefinitionID : Allow user to set JobDefinitionID
func (_options *DeploymentJobDefinitionsUpdateOptions) SetJobDefinitionID(jobDefinitionID string) *DeploymentJobDefinitionsUpdateOptions {
	_options.JobDefinitionID = core.StringPtr(jobDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobDefinitionsUpdateOptions) SetSpaceID(spaceID string) *DeploymentJobDefinitionsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *DeploymentJobDefinitionsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *DeploymentJobDefinitionsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobDefinitionsUpdateOptions) SetHeaders(param map[string]string) *DeploymentJobDefinitionsUpdateOptions {
	options.Headers = param
	return options
}

// DeploymentJobsCreateOptions : The DeploymentJobsCreate options.
type DeploymentJobsCreateOptions struct {
	// The space that contains the resource.
	SpaceID *string `json:"space_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// A reference to a resource.
	Deployment *SimpleRel `json:"deployment,omitempty"`

	// User defined properties.
	Custom interface{} `json:"custom,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityRequestHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`

	// Details about the input/output data and other properties to be used for a batch deployment job of a model,
	// Python Function or a Python Scripts.
	//
	// Use `input_data` property to specify the input data for batch processing as part of the job's payload.
	// The `input_data` property is mutually exclusive with `input_data_references` property, only use one of these.
	// When `input_data` is specified, the processed output of batch deployment job will be available in
	// `scoring.predictions`
	// parameter in the deployment job response. `input_data` property is not supported for batch deployment of Python
	// Scripts.
	//
	// Use `input_data_references` property to specify the details pertaining to the remote source where the input data for
	// batch deployment job is available. The `input_data_references` must be used with `output_data_references`.
	// The `input_data_references` property
	// is mutually exclusive with `input_data` property, only use one of these. The `input_data_references`
	// property is not supported for batch deployment job of Spark models and Python Functions.
	//
	// Use `output_data_references` property to specify the details pertaining to the remote source where the input
	// data for batch deployment job is available. `output_data_references` must be used with `input_data_references`.
	// The `output_data_references`
	// property is not supported for batch deployment job of Spark models and Python Functions.
	Scoring *JobScoringRequest `json:"scoring,omitempty"`

	// Details about the input/output data and other properties to be used for a batch
	// deployment job of a decision optimization problem.
	//
	// Use `input_data` property to specify the input data for batch processing as part
	// of the job's payload. The `input_data` property is mutually exclusive with `input_data_references`
	// property, only use one of these. When `input_data` is specified, the processed output
	// of batch deployment job will be available in `scoring.predictions` parameter in the deployment job response.
	//
	// Use `input_data_references` property to specify the details pertaining to the remote source
	// where the input data for batch deployment job is available. The `input_data_references` must be used
	// with `output_data_references`.
	// The `input_data_references` property is mutually exclusive with `input_data` property, only use one of these.
	//
	// Use `output_data_references` property to specify the details pertaining to the remote source where the
	// input data for batch deployment job is available. The `output_data_references` must be used with
	// `input_data_references`.
	DecisionOptimization *JobDecisionOptimizationRequest `json:"decision_optimization,omitempty"`

	// Defines number of days to retain the job meta. Job meta will be auto deleted after that. Value '-1' sets the meta to
	// be never auto deleted. accepted values are positive integer and '-1'. default value considered if the parameter not
	// passed to be '30' days.
	Retention *string `json:"retention,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobsCreateOptions : Instantiate DeploymentJobsCreateOptions
func (*WatsonMachineLearningV4) NewDeploymentJobsCreateOptions() *DeploymentJobsCreateOptions {
	return &DeploymentJobsCreateOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobsCreateOptions) SetSpaceID(spaceID string) *DeploymentJobsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *DeploymentJobsCreateOptions) SetName(name string) *DeploymentJobsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *DeploymentJobsCreateOptions) SetDescription(description string) *DeploymentJobsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *DeploymentJobsCreateOptions) SetTags(tags []string) *DeploymentJobsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetDeployment : Allow user to set Deployment
func (_options *DeploymentJobsCreateOptions) SetDeployment(deployment *SimpleRel) *DeploymentJobsCreateOptions {
	_options.Deployment = deployment
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *DeploymentJobsCreateOptions) SetCustom(custom interface{}) *DeploymentJobsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHardwareSpec : Allow user to set HardwareSpec
func (_options *DeploymentJobsCreateOptions) SetHardwareSpec(hardwareSpec *HardwareSpecRel) *DeploymentJobsCreateOptions {
	_options.HardwareSpec = hardwareSpec
	return _options
}

// SetHybridPipelineHardwareSpecs : Allow user to set HybridPipelineHardwareSpecs
func (_options *DeploymentJobsCreateOptions) SetHybridPipelineHardwareSpecs(hybridPipelineHardwareSpecs []JobEntityRequestHybridPipelineHardwareSpecsItem) *DeploymentJobsCreateOptions {
	_options.HybridPipelineHardwareSpecs = hybridPipelineHardwareSpecs
	return _options
}

// SetScoring : Allow user to set Scoring
func (_options *DeploymentJobsCreateOptions) SetScoring(scoring *JobScoringRequest) *DeploymentJobsCreateOptions {
	_options.Scoring = scoring
	return _options
}

// SetDecisionOptimization : Allow user to set DecisionOptimization
func (_options *DeploymentJobsCreateOptions) SetDecisionOptimization(decisionOptimization *JobDecisionOptimizationRequest) *DeploymentJobsCreateOptions {
	_options.DecisionOptimization = decisionOptimization
	return _options
}

// SetRetention : Allow user to set Retention
func (_options *DeploymentJobsCreateOptions) SetRetention(retention string) *DeploymentJobsCreateOptions {
	_options.Retention = core.StringPtr(retention)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobsCreateOptions) SetHeaders(param map[string]string) *DeploymentJobsCreateOptions {
	options.Headers = param
	return options
}

// DeploymentJobsDeleteOptions : The DeploymentJobsDelete options.
type DeploymentJobsDeleteOptions struct {
	// The id of the job.
	JobID *string `json:"job_id" validate:"required,ne="`

	// Cancel the deployment job that belong to this space.
	SpaceID *string `json:"space_id" validate:"required"`

	// Set to true in order to also delete the job metadata information.
	HardDelete *bool `json:"hard_delete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobsDeleteOptions : Instantiate DeploymentJobsDeleteOptions
func (*WatsonMachineLearningV4) NewDeploymentJobsDeleteOptions(jobID string, spaceID string) *DeploymentJobsDeleteOptions {
	return &DeploymentJobsDeleteOptions{
		JobID:   core.StringPtr(jobID),
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetJobID : Allow user to set JobID
func (_options *DeploymentJobsDeleteOptions) SetJobID(jobID string) *DeploymentJobsDeleteOptions {
	_options.JobID = core.StringPtr(jobID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobsDeleteOptions) SetSpaceID(spaceID string) *DeploymentJobsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetHardDelete : Allow user to set HardDelete
func (_options *DeploymentJobsDeleteOptions) SetHardDelete(hardDelete bool) *DeploymentJobsDeleteOptions {
	_options.HardDelete = core.BoolPtr(hardDelete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobsDeleteOptions) SetHeaders(param map[string]string) *DeploymentJobsDeleteOptions {
	options.Headers = param
	return options
}

// DeploymentJobsGetOptions : The DeploymentJobsGet options.
type DeploymentJobsGetOptions struct {
	// The id of the job.
	JobID *string `json:"job_id" validate:"required,ne="`

	// Retrieve the deployment job that belong to this space.
	SpaceID *string `json:"space_id" validate:"required"`

	// Retrieves only fields from 'decision_optimization' and 'scoring' section mentioned as value(s)
	// (comma separated) as output response fields. Retrieves all the fields if not mentioned.
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobsGetOptions : Instantiate DeploymentJobsGetOptions
func (*WatsonMachineLearningV4) NewDeploymentJobsGetOptions(jobID string, spaceID string) *DeploymentJobsGetOptions {
	return &DeploymentJobsGetOptions{
		JobID:   core.StringPtr(jobID),
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetJobID : Allow user to set JobID
func (_options *DeploymentJobsGetOptions) SetJobID(jobID string) *DeploymentJobsGetOptions {
	_options.JobID = core.StringPtr(jobID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobsGetOptions) SetSpaceID(spaceID string) *DeploymentJobsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetInclude : Allow user to set Include
func (_options *DeploymentJobsGetOptions) SetInclude(include string) *DeploymentJobsGetOptions {
	_options.Include = core.StringPtr(include)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobsGetOptions) SetHeaders(param map[string]string) *DeploymentJobsGetOptions {
	options.Headers = param
	return options
}

// DeploymentJobsListOptions : The DeploymentJobsList options.
type DeploymentJobsListOptions struct {
	// Retrieves the deployment jobs that belong to this space.
	SpaceID *string `json:"space_id" validate:"required"`

	// Retrieves only the jobs with these tags (comma separated).
	TagValue *string `json:"tag.value,omitempty"`

	// Filter based on on the deployment job state: queued, running, completed, failed etc.
	State *string `json:"state,omitempty"`

	// Filter based on the deployment_id.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// Retrieves only fields from 'decision_optimization' and 'scoring' section mentioned as value(s)
	// (comma separated) as output response fields. Retrieves all the fields if not mentioned.
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentJobsListOptions : Instantiate DeploymentJobsListOptions
func (*WatsonMachineLearningV4) NewDeploymentJobsListOptions(spaceID string) *DeploymentJobsListOptions {
	return &DeploymentJobsListOptions{
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentJobsListOptions) SetSpaceID(spaceID string) *DeploymentJobsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *DeploymentJobsListOptions) SetTagValue(tagValue string) *DeploymentJobsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetState : Allow user to set State
func (_options *DeploymentJobsListOptions) SetState(state string) *DeploymentJobsListOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *DeploymentJobsListOptions) SetDeploymentID(deploymentID string) *DeploymentJobsListOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetInclude : Allow user to set Include
func (_options *DeploymentJobsListOptions) SetInclude(include string) *DeploymentJobsListOptions {
	_options.Include = core.StringPtr(include)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentJobsListOptions) SetHeaders(param map[string]string) *DeploymentJobsListOptions {
	options.Headers = param
	return options
}

// DeploymentPatchRequestHelperRShiny : Specify this section if deploying an Shiny application.
type DeploymentPatchRequestHelperRShiny struct {
	// Specifies the type of users who can access the Shiny application.
	Authentication *string `json:"authentication" validate:"required"`

	// A set of parameters that specify details about the Shiny deployment.
	Parameters *DeploymentPatchRequestHelperRShinyParameters `json:"parameters,omitempty"`
}

// Constants associated with the DeploymentPatchRequestHelperRShiny.Authentication property.
// Specifies the type of users who can access the Shiny application.
const (
	DeploymentPatchRequestHelperRShiny_Authentication_AnyValidUser             = "any_valid_user"
	DeploymentPatchRequestHelperRShiny_Authentication_AnyoneWithURL            = "anyone_with_url"
	DeploymentPatchRequestHelperRShiny_Authentication_MembersOfDeploymentSpace = "members_of_deployment_space"
)

// UnmarshalDeploymentPatchRequestHelperRShiny unmarshals an instance of DeploymentPatchRequestHelperRShiny from the specified map of raw messages.
func UnmarshalDeploymentPatchRequestHelperRShiny(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentPatchRequestHelperRShiny)
	err = core.UnmarshalPrimitive(m, "authentication", &obj.Authentication)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parameters", &obj.Parameters, UnmarshalDeploymentPatchRequestHelperRShinyParameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentPatchRequestHelperRShinyParameters : A set of parameters that specify details about the Shiny deployment.
type DeploymentPatchRequestHelperRShinyParameters struct {
	// Specifies the unique `serving_name` that will be used to create a URL for the application. The creation of the
	// deployment will fail if the `serving_name` is not unique.
	ServingName *string `json:"serving_name,omitempty"`

	// Specifies the details when deploying with a code_package asset.
	CodePackage *DeploymentPatchRequestHelperRShinyParametersCodePackage `json:"code_package,omitempty"`
}

// UnmarshalDeploymentPatchRequestHelperRShinyParameters unmarshals an instance of DeploymentPatchRequestHelperRShinyParameters from the specified map of raw messages.
func UnmarshalDeploymentPatchRequestHelperRShinyParameters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentPatchRequestHelperRShinyParameters)
	err = core.UnmarshalPrimitive(m, "serving_name", &obj.ServingName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "code_package", &obj.CodePackage, UnmarshalDeploymentPatchRequestHelperRShinyParametersCodePackage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentPatchRequestHelperRShinyParametersCodePackage : Specifies the details when deploying with a code_package asset.
type DeploymentPatchRequestHelperRShinyParametersCodePackage struct {
	// The path to the application files in the code package.
	Path *string `json:"path" validate:"required"`
}

// UnmarshalDeploymentPatchRequestHelperRShinyParametersCodePackage unmarshals an instance of DeploymentPatchRequestHelperRShinyParametersCodePackage from the specified map of raw messages.
func UnmarshalDeploymentPatchRequestHelperRShinyParametersCodePackage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentPatchRequestHelperRShinyParametersCodePackage)
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentResourcesFirst : The reference to the first item in the current page.
type DeploymentResourcesFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalDeploymentResourcesFirst unmarshals an instance of DeploymentResourcesFirst from the specified map of raw messages.
func UnmarshalDeploymentResourcesFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentResourcesFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentResourcesNext : A reference to the first item of the next page, if any.
type DeploymentResourcesNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalDeploymentResourcesNext unmarshals an instance of DeploymentResourcesNext from the specified map of raw messages.
func UnmarshalDeploymentResourcesNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentResourcesNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentResourcesSystem : System details including warnings and stats. This will get populated only if 'stats' query parameter is passed with
// 'true'.
type DeploymentResourcesSystem struct {
	// Optional details provided by the service about statistics of the number of deployments created with in a space or
	// across all spaces.
	System *DeploymentSystemDetails `json:"system,omitempty"`
}

// UnmarshalDeploymentResourcesSystem unmarshals an instance of DeploymentResourcesSystem from the specified map of raw messages.
func UnmarshalDeploymentResourcesSystem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentResourcesSystem)
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalDeploymentSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentsComputePredictionsOptions : The DeploymentsComputePredictions options.
type DeploymentsComputePredictionsOptions struct {
	// The `deployment_id` can be either the `deployment_id` that identifies the deployment or a `serving_name` that allows
	// a predefined URL to be used to post a prediction.
	DeploymentID *string `json:"deployment_id" validate:"required,ne="`

	// The input data.
	InputData []InputDataArray `json:"input_data" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentsComputePredictionsOptions : Instantiate DeploymentsComputePredictionsOptions
func (*WatsonMachineLearningV4) NewDeploymentsComputePredictionsOptions(deploymentID string, inputData []InputDataArray) *DeploymentsComputePredictionsOptions {
	return &DeploymentsComputePredictionsOptions{
		DeploymentID: core.StringPtr(deploymentID),
		InputData:    inputData,
	}
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *DeploymentsComputePredictionsOptions) SetDeploymentID(deploymentID string) *DeploymentsComputePredictionsOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetInputData : Allow user to set InputData
func (_options *DeploymentsComputePredictionsOptions) SetInputData(inputData []InputDataArray) *DeploymentsComputePredictionsOptions {
	_options.InputData = inputData
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentsComputePredictionsOptions) SetHeaders(param map[string]string) *DeploymentsComputePredictionsOptions {
	options.Headers = param
	return options
}

// DeploymentsCreateOptions : The DeploymentsCreate options.
type DeploymentsCreateOptions struct {
	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// Tags that can be used when searching for resources.
	Tags []string `json:"tags,omitempty"`

	// The name of the deployment.
	Name *string `json:"name,omitempty"`

	// A description of the deployment.
	Description *string `json:"description,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// A reference to a resource.
	Asset *Rel `json:"asset,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityResultHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`

	// Indicates that this is an online deployment. An empty object has to be specified.
	// More properties will be added later on to setup the online deployment.
	// The `serving_name` can be provided in the `online.parameters`. The serving name can only have the characters
	// [a-z,0-9,_]
	// and the length should not be more than 36 characters. The `serving_name` can be used in the prediction URL in place
	// of the `deployment_id`.
	Online *DeploymentEntityRequestOnline `json:"online,omitempty"`

	// Indicates that this is a batch deployment. An empty object has to be specified.
	// More properties will be added later on to setup the batch deployment.
	Batch *DeploymentEntityRequestBatch `json:"batch,omitempty"`

	// Indicates that this is a Shiny application deployment.
	RShiny *DeploymentEntityRequestRShiny `json:"r_shiny,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentsCreateOptions : Instantiate DeploymentsCreateOptions
func (*WatsonMachineLearningV4) NewDeploymentsCreateOptions(spaceID string) *DeploymentsCreateOptions {
	return &DeploymentsCreateOptions{
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentsCreateOptions) SetSpaceID(spaceID string) *DeploymentsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *DeploymentsCreateOptions) SetTags(tags []string) *DeploymentsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetName : Allow user to set Name
func (_options *DeploymentsCreateOptions) SetName(name string) *DeploymentsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *DeploymentsCreateOptions) SetDescription(description string) *DeploymentsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *DeploymentsCreateOptions) SetCustom(custom map[string]interface{}) *DeploymentsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetAsset : Allow user to set Asset
func (_options *DeploymentsCreateOptions) SetAsset(asset *Rel) *DeploymentsCreateOptions {
	_options.Asset = asset
	return _options
}

// SetHardwareSpec : Allow user to set HardwareSpec
func (_options *DeploymentsCreateOptions) SetHardwareSpec(hardwareSpec *HardwareSpecRel) *DeploymentsCreateOptions {
	_options.HardwareSpec = hardwareSpec
	return _options
}

// SetHybridPipelineHardwareSpecs : Allow user to set HybridPipelineHardwareSpecs
func (_options *DeploymentsCreateOptions) SetHybridPipelineHardwareSpecs(hybridPipelineHardwareSpecs []JobEntityResultHybridPipelineHardwareSpecsItem) *DeploymentsCreateOptions {
	_options.HybridPipelineHardwareSpecs = hybridPipelineHardwareSpecs
	return _options
}

// SetOnline : Allow user to set Online
func (_options *DeploymentsCreateOptions) SetOnline(online *DeploymentEntityRequestOnline) *DeploymentsCreateOptions {
	_options.Online = online
	return _options
}

// SetBatch : Allow user to set Batch
func (_options *DeploymentsCreateOptions) SetBatch(batch *DeploymentEntityRequestBatch) *DeploymentsCreateOptions {
	_options.Batch = batch
	return _options
}

// SetRShiny : Allow user to set RShiny
func (_options *DeploymentsCreateOptions) SetRShiny(rShiny *DeploymentEntityRequestRShiny) *DeploymentsCreateOptions {
	_options.RShiny = rShiny
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentsCreateOptions) SetHeaders(param map[string]string) *DeploymentsCreateOptions {
	options.Headers = param
	return options
}

// DeploymentsDeleteOptions : The DeploymentsDelete options.
type DeploymentsDeleteOptions struct {
	// The deployment id.
	DeploymentID *string `json:"deployment_id" validate:"required,ne="`

	// Retrieves the deployments of assets that belong to this space.
	SpaceID *string `json:"space_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentsDeleteOptions : Instantiate DeploymentsDeleteOptions
func (*WatsonMachineLearningV4) NewDeploymentsDeleteOptions(deploymentID string, spaceID string) *DeploymentsDeleteOptions {
	return &DeploymentsDeleteOptions{
		DeploymentID: core.StringPtr(deploymentID),
		SpaceID:      core.StringPtr(spaceID),
	}
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *DeploymentsDeleteOptions) SetDeploymentID(deploymentID string) *DeploymentsDeleteOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentsDeleteOptions) SetSpaceID(spaceID string) *DeploymentsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentsDeleteOptions) SetHeaders(param map[string]string) *DeploymentsDeleteOptions {
	options.Headers = param
	return options
}

// DeploymentsGetOptions : The DeploymentsGet options.
type DeploymentsGetOptions struct {
	// The deployment id.
	DeploymentID *string `json:"deployment_id" validate:"required,ne="`

	// Retrieves the deployments of assets that belong to this space.
	SpaceID *string `json:"space_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentsGetOptions : Instantiate DeploymentsGetOptions
func (*WatsonMachineLearningV4) NewDeploymentsGetOptions(deploymentID string, spaceID string) *DeploymentsGetOptions {
	return &DeploymentsGetOptions{
		DeploymentID: core.StringPtr(deploymentID),
		SpaceID:      core.StringPtr(spaceID),
	}
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *DeploymentsGetOptions) SetDeploymentID(deploymentID string) *DeploymentsGetOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentsGetOptions) SetSpaceID(spaceID string) *DeploymentsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentsGetOptions) SetHeaders(param map[string]string) *DeploymentsGetOptions {
	options.Headers = param
	return options
}

// DeploymentsListOptions : The DeploymentsList options.
type DeploymentsListOptions struct {
	// Retrieves the deployments that belong to this space.
	SpaceID *string `json:"space_id,omitempty"`

	// Retrieves the deployment, if any, that contains this `serving_name`.
	ServingName *string `json:"serving_name,omitempty"`

	// Retrieves only the resources with the given tag value.
	TagValue *string `json:"tag.value,omitempty"`

	// Retrieves only the resources with the given asset_id, asset_id would be either model_id or function_id.
	AssetID *string `json:"asset_id,omitempty"`

	// Retrieves only the resources with the given name.
	Name *string `json:"name,omitempty"`

	// Retrieves the resources filtered with the given type. Allowed values are `model`, `function`, `py_script`, `r_shiny`
	// and `do`.
	Type *string `json:"type,omitempty"`

	// Retrieves the resources filtered by state. Allowed values are `initializing`, `updating`, `ready` and `failed`.
	State *string `json:"state,omitempty"`

	// Returns stats about deployments within a space or across spaces if it is set to true. This query parameter cannot be
	// combined with any other except for 'space_id'.
	Stats *bool `json:"stats,omitempty"`

	// Returns whether serving_name is available for use or not. This query parameter cannot be combined with any other
	// except for 'serving_name'.
	Conflict *bool `json:"conflict,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentsListOptions : Instantiate DeploymentsListOptions
func (*WatsonMachineLearningV4) NewDeploymentsListOptions() *DeploymentsListOptions {
	return &DeploymentsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentsListOptions) SetSpaceID(spaceID string) *DeploymentsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetServingName : Allow user to set ServingName
func (_options *DeploymentsListOptions) SetServingName(servingName string) *DeploymentsListOptions {
	_options.ServingName = core.StringPtr(servingName)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *DeploymentsListOptions) SetTagValue(tagValue string) *DeploymentsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetAssetID : Allow user to set AssetID
func (_options *DeploymentsListOptions) SetAssetID(assetID string) *DeploymentsListOptions {
	_options.AssetID = core.StringPtr(assetID)
	return _options
}

// SetName : Allow user to set Name
func (_options *DeploymentsListOptions) SetName(name string) *DeploymentsListOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetType : Allow user to set Type
func (_options *DeploymentsListOptions) SetType(typeVar string) *DeploymentsListOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetState : Allow user to set State
func (_options *DeploymentsListOptions) SetState(state string) *DeploymentsListOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetStats : Allow user to set Stats
func (_options *DeploymentsListOptions) SetStats(stats bool) *DeploymentsListOptions {
	_options.Stats = core.BoolPtr(stats)
	return _options
}

// SetConflict : Allow user to set Conflict
func (_options *DeploymentsListOptions) SetConflict(conflict bool) *DeploymentsListOptions {
	_options.Conflict = core.BoolPtr(conflict)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentsListOptions) SetHeaders(param map[string]string) *DeploymentsListOptions {
	options.Headers = param
	return options
}

// DeploymentsUpdateOptions : The DeploymentsUpdate options.
type DeploymentsUpdateOptions struct {
	// The deployment id.
	DeploymentID *string `json:"deployment_id" validate:"required,ne="`

	// Retrieves the deployments of assets that belong to this space.
	SpaceID *string `json:"space_id" validate:"required"`

	// The json patch.
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeploymentsUpdateOptions : Instantiate DeploymentsUpdateOptions
func (*WatsonMachineLearningV4) NewDeploymentsUpdateOptions(deploymentID string, spaceID string, jsonPatch []JSONPatchOperation) *DeploymentsUpdateOptions {
	return &DeploymentsUpdateOptions{
		DeploymentID: core.StringPtr(deploymentID),
		SpaceID:      core.StringPtr(spaceID),
		JSONPatch:    jsonPatch,
	}
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *DeploymentsUpdateOptions) SetDeploymentID(deploymentID string) *DeploymentsUpdateOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *DeploymentsUpdateOptions) SetSpaceID(spaceID string) *DeploymentsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *DeploymentsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *DeploymentsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeploymentsUpdateOptions) SetHeaders(param map[string]string) *DeploymentsUpdateOptions {
	options.Headers = param
	return options
}

// ErrorErrorsItem : ErrorErrorsItem struct
type ErrorErrorsItem struct {
	// A simple code that should convey the general sense of the error.
	Code *string `json:"code" validate:"required"`

	// The message that describes the error.
	Message *string `json:"message" validate:"required"`

	// A reference to a more detailed explanation when available.
	MoreInfo *string `json:"more_info,omitempty"`
}

// UnmarshalErrorErrorsItem unmarshals an instance of ErrorErrorsItem from the specified map of raw messages.
func UnmarshalErrorErrorsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ErrorErrorsItem)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EvaluationsSpecItem : An evaluation specification used to support evaluations for TensorFlow.
type EvaluationsSpecItem struct {
	// An identifier of this metrics set. For a DL problem this can be the output tensor id/name in order to identify on
	// which output these metrics will be computed.
	ID *string `json:"id,omitempty"`

	// The id of the `input_data.id` with the `type=target`. This points to the ground truth information that will be used
	// together with prediction information to generate metrics.
	InputTarget *string `json:"input_target,omitempty"`

	// A list of the metric names.
	MetricsNames []string `json:"metrics_names,omitempty"`
}

// UnmarshalEvaluationsSpecItem unmarshals an instance of EvaluationsSpecItem from the specified map of raw messages.
func UnmarshalEvaluationsSpecItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EvaluationsSpecItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_target", &obj.InputTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metrics_names", &obj.MetricsNames)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExperimentResourceEntity : The details of the experiment to be created.
type ExperimentResourceEntity struct {
	// The label column.
	LabelColumn *string `json:"label_column,omitempty"`

	// The optional evaluation definition.
	EvaluationDefinition *EvaluationDefinition `json:"evaluation_definition,omitempty"`

	// The optional training references used by the experiment.
	TrainingReferences []TrainingReference `json:"training_references,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// UnmarshalExperimentResourceEntity unmarshals an instance of ExperimentResourceEntity from the specified map of raw messages.
func UnmarshalExperimentResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExperimentResourceEntity)
	err = core.UnmarshalPrimitive(m, "label_column", &obj.LabelColumn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "evaluation_definition", &obj.EvaluationDefinition, UnmarshalEvaluationDefinition)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_references", &obj.TrainingReferences, UnmarshalTrainingReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExperimentResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type ExperimentResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalExperimentResourceMetadata unmarshals an instance of ExperimentResourceMetadata from the specified map of raw messages.
func UnmarshalExperimentResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExperimentResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// ExperimentsCreateOptions : The ExperimentsCreate options.
type ExperimentsCreateOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The label column.
	LabelColumn *string `json:"label_column,omitempty"`

	// The optional evaluation definition.
	EvaluationDefinition *EvaluationDefinition `json:"evaluation_definition,omitempty"`

	// The optional training references used by the experiment.
	TrainingReferences []TrainingReference `json:"training_references,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExperimentsCreateOptions : Instantiate ExperimentsCreateOptions
func (*WatsonMachineLearningV4) NewExperimentsCreateOptions(name string) *ExperimentsCreateOptions {
	return &ExperimentsCreateOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *ExperimentsCreateOptions) SetName(name string) *ExperimentsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExperimentsCreateOptions) SetProjectID(projectID string) *ExperimentsCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExperimentsCreateOptions) SetSpaceID(spaceID string) *ExperimentsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ExperimentsCreateOptions) SetDescription(description string) *ExperimentsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *ExperimentsCreateOptions) SetTags(tags []string) *ExperimentsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetLabelColumn : Allow user to set LabelColumn
func (_options *ExperimentsCreateOptions) SetLabelColumn(labelColumn string) *ExperimentsCreateOptions {
	_options.LabelColumn = core.StringPtr(labelColumn)
	return _options
}

// SetEvaluationDefinition : Allow user to set EvaluationDefinition
func (_options *ExperimentsCreateOptions) SetEvaluationDefinition(evaluationDefinition *EvaluationDefinition) *ExperimentsCreateOptions {
	_options.EvaluationDefinition = evaluationDefinition
	return _options
}

// SetTrainingReferences : Allow user to set TrainingReferences
func (_options *ExperimentsCreateOptions) SetTrainingReferences(trainingReferences []TrainingReference) *ExperimentsCreateOptions {
	_options.TrainingReferences = trainingReferences
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *ExperimentsCreateOptions) SetCustom(custom map[string]interface{}) *ExperimentsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExperimentsCreateOptions) SetHeaders(param map[string]string) *ExperimentsCreateOptions {
	options.Headers = param
	return options
}

// ExperimentsCreateRevisionOptions : The ExperimentsCreateRevision options.
type ExperimentsCreateRevisionOptions struct {
	// Experiment identifier.
	ExperimentID *string `json:"experiment_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExperimentsCreateRevisionOptions : Instantiate ExperimentsCreateRevisionOptions
func (*WatsonMachineLearningV4) NewExperimentsCreateRevisionOptions(experimentID string) *ExperimentsCreateRevisionOptions {
	return &ExperimentsCreateRevisionOptions{
		ExperimentID: core.StringPtr(experimentID),
	}
}

// SetExperimentID : Allow user to set ExperimentID
func (_options *ExperimentsCreateRevisionOptions) SetExperimentID(experimentID string) *ExperimentsCreateRevisionOptions {
	_options.ExperimentID = core.StringPtr(experimentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExperimentsCreateRevisionOptions) SetSpaceID(spaceID string) *ExperimentsCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExperimentsCreateRevisionOptions) SetProjectID(projectID string) *ExperimentsCreateRevisionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *ExperimentsCreateRevisionOptions) SetCommitMessage(commitMessage string) *ExperimentsCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExperimentsCreateRevisionOptions) SetHeaders(param map[string]string) *ExperimentsCreateRevisionOptions {
	options.Headers = param
	return options
}

// ExperimentsDeleteOptions : The ExperimentsDelete options.
type ExperimentsDeleteOptions struct {
	// Experiment identifier.
	ExperimentID *string `json:"experiment_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExperimentsDeleteOptions : Instantiate ExperimentsDeleteOptions
func (*WatsonMachineLearningV4) NewExperimentsDeleteOptions(experimentID string) *ExperimentsDeleteOptions {
	return &ExperimentsDeleteOptions{
		ExperimentID: core.StringPtr(experimentID),
	}
}

// SetExperimentID : Allow user to set ExperimentID
func (_options *ExperimentsDeleteOptions) SetExperimentID(experimentID string) *ExperimentsDeleteOptions {
	_options.ExperimentID = core.StringPtr(experimentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExperimentsDeleteOptions) SetSpaceID(spaceID string) *ExperimentsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExperimentsDeleteOptions) SetProjectID(projectID string) *ExperimentsDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExperimentsDeleteOptions) SetHeaders(param map[string]string) *ExperimentsDeleteOptions {
	options.Headers = param
	return options
}

// ExperimentsGetOptions : The ExperimentsGet options.
type ExperimentsGetOptions struct {
	// Experiment identifier.
	ExperimentID *string `json:"experiment_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExperimentsGetOptions : Instantiate ExperimentsGetOptions
func (*WatsonMachineLearningV4) NewExperimentsGetOptions(experimentID string) *ExperimentsGetOptions {
	return &ExperimentsGetOptions{
		ExperimentID: core.StringPtr(experimentID),
	}
}

// SetExperimentID : Allow user to set ExperimentID
func (_options *ExperimentsGetOptions) SetExperimentID(experimentID string) *ExperimentsGetOptions {
	_options.ExperimentID = core.StringPtr(experimentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExperimentsGetOptions) SetSpaceID(spaceID string) *ExperimentsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExperimentsGetOptions) SetProjectID(projectID string) *ExperimentsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *ExperimentsGetOptions) SetRev(rev string) *ExperimentsGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExperimentsGetOptions) SetHeaders(param map[string]string) *ExperimentsGetOptions {
	options.Headers = param
	return options
}

// ExperimentsListOptions : The ExperimentsList options.
type ExperimentsListOptions struct {
	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExperimentsListOptions : Instantiate ExperimentsListOptions
func (*WatsonMachineLearningV4) NewExperimentsListOptions() *ExperimentsListOptions {
	return &ExperimentsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExperimentsListOptions) SetSpaceID(spaceID string) *ExperimentsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExperimentsListOptions) SetProjectID(projectID string) *ExperimentsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ExperimentsListOptions) SetStart(start string) *ExperimentsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ExperimentsListOptions) SetLimit(limit int64) *ExperimentsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *ExperimentsListOptions) SetTagValue(tagValue string) *ExperimentsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *ExperimentsListOptions) SetSearch(search string) *ExperimentsListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExperimentsListOptions) SetHeaders(param map[string]string) *ExperimentsListOptions {
	options.Headers = param
	return options
}

// ExperimentsListRevisionsOptions : The ExperimentsListRevisions options.
type ExperimentsListRevisionsOptions struct {
	// Experiment identifier.
	ExperimentID *string `json:"experiment_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExperimentsListRevisionsOptions : Instantiate ExperimentsListRevisionsOptions
func (*WatsonMachineLearningV4) NewExperimentsListRevisionsOptions(experimentID string) *ExperimentsListRevisionsOptions {
	return &ExperimentsListRevisionsOptions{
		ExperimentID: core.StringPtr(experimentID),
	}
}

// SetExperimentID : Allow user to set ExperimentID
func (_options *ExperimentsListRevisionsOptions) SetExperimentID(experimentID string) *ExperimentsListRevisionsOptions {
	_options.ExperimentID = core.StringPtr(experimentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExperimentsListRevisionsOptions) SetSpaceID(spaceID string) *ExperimentsListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExperimentsListRevisionsOptions) SetProjectID(projectID string) *ExperimentsListRevisionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ExperimentsListRevisionsOptions) SetStart(start string) *ExperimentsListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ExperimentsListRevisionsOptions) SetLimit(limit int64) *ExperimentsListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExperimentsListRevisionsOptions) SetHeaders(param map[string]string) *ExperimentsListRevisionsOptions {
	options.Headers = param
	return options
}

// ExperimentsUpdateOptions : The ExperimentsUpdate options.
type ExperimentsUpdateOptions struct {
	// Experiment identifier.
	ExperimentID *string `json:"experiment_id" validate:"required,ne="`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExperimentsUpdateOptions : Instantiate ExperimentsUpdateOptions
func (*WatsonMachineLearningV4) NewExperimentsUpdateOptions(experimentID string, jsonPatch []JSONPatchOperation) *ExperimentsUpdateOptions {
	return &ExperimentsUpdateOptions{
		ExperimentID: core.StringPtr(experimentID),
		JSONPatch:    jsonPatch,
	}
}

// SetExperimentID : Allow user to set ExperimentID
func (_options *ExperimentsUpdateOptions) SetExperimentID(experimentID string) *ExperimentsUpdateOptions {
	_options.ExperimentID = core.StringPtr(experimentID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *ExperimentsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *ExperimentsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExperimentsUpdateOptions) SetSpaceID(spaceID string) *ExperimentsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExperimentsUpdateOptions) SetProjectID(projectID string) *ExperimentsUpdateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExperimentsUpdateOptions) SetHeaders(param map[string]string) *ExperimentsUpdateOptions {
	options.Headers = param
	return options
}

// FederatedLearningInfoAggregator : FederatedLearningInfoAggregator struct
type FederatedLearningInfoAggregator struct {
	Connection *FederatedLearningInfoAggregatorConnection `json:"connection,omitempty"`
}

// UnmarshalFederatedLearningInfoAggregator unmarshals an instance of FederatedLearningInfoAggregator from the specified map of raw messages.
func UnmarshalFederatedLearningInfoAggregator(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningInfoAggregator)
	err = core.UnmarshalModel(m, "connection", &obj.Connection, UnmarshalFederatedLearningInfoAggregatorConnection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningInfoAggregatorConnection : FederatedLearningInfoAggregatorConnection struct
type FederatedLearningInfoAggregatorConnection struct {
	Host *string `json:"host,omitempty"`

	Port *float64 `json:"port,omitempty"`

	Endpoint *string `json:"endpoint,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	// The headers object contains key-value pairs of headers to be supplied when connecting to the aggregator.
	HeadersVar interface{} `json:"headers,omitempty"`
}

// Constants associated with the FederatedLearningInfoAggregatorConnection.Protocol property.
const (
	FederatedLearningInfoAggregatorConnection_Protocol_Wss = "wss"
)

// UnmarshalFederatedLearningInfoAggregatorConnection unmarshals an instance of FederatedLearningInfoAggregatorConnection from the specified map of raw messages.
func UnmarshalFederatedLearningInfoAggregatorConnection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningInfoAggregatorConnection)
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "protocol", &obj.Protocol)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "headers", &obj.HeadersVar)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningInfoRemoteTrainingSystemsItem : FederatedLearningInfoRemoteTrainingSystemsItem struct
type FederatedLearningInfoRemoteTrainingSystemsItem struct {
	States *FederatedLearningInfoRemoteTrainingSystemsItemStates `json:"states,omitempty"`
}

// UnmarshalFederatedLearningInfoRemoteTrainingSystemsItem unmarshals an instance of FederatedLearningInfoRemoteTrainingSystemsItem from the specified map of raw messages.
func UnmarshalFederatedLearningInfoRemoteTrainingSystemsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningInfoRemoteTrainingSystemsItem)
	err = core.UnmarshalModel(m, "states", &obj.States, UnmarshalFederatedLearningInfoRemoteTrainingSystemsItemStates)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningInfoRemoteTrainingSystemsItemStates : FederatedLearningInfoRemoteTrainingSystemsItemStates struct
type FederatedLearningInfoRemoteTrainingSystemsItemStates struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	State *string `json:"state,omitempty"`

	RegistrationLog []FederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem `json:"registration_log,omitempty"`
}

// UnmarshalFederatedLearningInfoRemoteTrainingSystemsItemStates unmarshals an instance of FederatedLearningInfoRemoteTrainingSystemsItemStates from the specified map of raw messages.
func UnmarshalFederatedLearningInfoRemoteTrainingSystemsItemStates(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningInfoRemoteTrainingSystemsItemStates)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "registration_log", &obj.RegistrationLog, UnmarshalFederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem : FederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem struct
type FederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem struct {
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	Event *string `json:"event,omitempty"`
}

// UnmarshalFederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem unmarshals an instance of FederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem from the specified map of raw messages.
func UnmarshalFederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningInfoRemoteTrainingSystemsItemStatesRegistrationLogItem)
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event", &obj.Event)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningModel : The model type for federated_learning.
type FederatedLearningModel struct {
	Type *string `json:"type" validate:"required"`

	Spec *FederatedLearningModelSpec `json:"spec,omitempty"`

	ModelFile *string `json:"model_file,omitempty"`
}

// Constants associated with the FederatedLearningModel.Type property.
const (
	FederatedLearningModel_Type_Dt            = "dt"
	FederatedLearningModel_Type_Keras         = "keras"
	FederatedLearningModel_Type_KerasDp       = "keras_dp"
	FederatedLearningModel_Type_Sklearn       = "sklearn"
	FederatedLearningModel_Type_SklearnKmeans = "sklearn_kmeans"
	FederatedLearningModel_Type_SklearnNb     = "sklearn_nb"
	FederatedLearningModel_Type_SklearnSgd    = "sklearn_sgd"
	FederatedLearningModel_Type_Tensorflow    = "tensorflow"
	FederatedLearningModel_Type_XgbClassifier = "xgb_classifier"
	FederatedLearningModel_Type_XgbRegressor  = "xgb_regressor"
)

// NewFederatedLearningModel : Instantiate FederatedLearningModel (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewFederatedLearningModel(typeVar string) (_model *FederatedLearningModel, err error) {
	_model = &FederatedLearningModel{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalFederatedLearningModel unmarshals an instance of FederatedLearningModel from the specified map of raw messages.
func UnmarshalFederatedLearningModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningModel)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spec", &obj.Spec, UnmarshalFederatedLearningModelSpec)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_file", &obj.ModelFile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningModelSpec : FederatedLearningModelSpec struct
type FederatedLearningModelSpec struct {
	// A reference to a resource.
	Href *Rel `json:"href,omitempty"`
}

// UnmarshalFederatedLearningModelSpec unmarshals an instance of FederatedLearningModelSpec from the specified map of raw messages.
func UnmarshalFederatedLearningModelSpec(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningModelSpec)
	err = core.UnmarshalModel(m, "href", &obj.Href, UnmarshalRel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningOptimizer : The optimizer for federated learning.
type FederatedLearningOptimizer struct {
	Name *string `json:"name,omitempty"`

	// Optimizer specification.
	Spec interface{} `json:"spec,omitempty"`
}

// UnmarshalFederatedLearningOptimizer unmarshals an instance of FederatedLearningOptimizer from the specified map of raw messages.
func UnmarshalFederatedLearningOptimizer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningOptimizer)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spec", &obj.Spec)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningRemoteTraining : The remote training for federated learning.
type FederatedLearningRemoteTraining struct {
	// The quorum for federated learning.
	Quorum *float64 `json:"quorum,omitempty"`

	// The maximum timeout for federated learning.
	MaxTimeout *int64 `json:"max_timeout,omitempty"`

	RemoteTrainingSystems []FederatedLearningRemoteTrainingRemoteTrainingSystemsItem `json:"remote_training_systems" validate:"required"`
}

// NewFederatedLearningRemoteTraining : Instantiate FederatedLearningRemoteTraining (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewFederatedLearningRemoteTraining(remoteTrainingSystems []FederatedLearningRemoteTrainingRemoteTrainingSystemsItem) (_model *FederatedLearningRemoteTraining, err error) {
	_model = &FederatedLearningRemoteTraining{
		RemoteTrainingSystems: remoteTrainingSystems,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalFederatedLearningRemoteTraining unmarshals an instance of FederatedLearningRemoteTraining from the specified map of raw messages.
func UnmarshalFederatedLearningRemoteTraining(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningRemoteTraining)
	err = core.UnmarshalPrimitive(m, "quorum", &obj.Quorum)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_timeout", &obj.MaxTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "remote_training_systems", &obj.RemoteTrainingSystems, UnmarshalFederatedLearningRemoteTrainingRemoteTrainingSystemsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningRemoteTrainingRemoteTrainingSystemsItem : FederatedLearningRemoteTrainingRemoteTrainingSystemsItem struct
type FederatedLearningRemoteTrainingRemoteTrainingSystemsItem struct {
	// The remote training id of the referenced resource.
	ID *string `json:"id" validate:"required"`

	// The required of the referenced resource.
	Required *bool `json:"required,omitempty"`
}

// NewFederatedLearningRemoteTrainingRemoteTrainingSystemsItem : Instantiate FederatedLearningRemoteTrainingRemoteTrainingSystemsItem (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewFederatedLearningRemoteTrainingRemoteTrainingSystemsItem(id string) (_model *FederatedLearningRemoteTrainingRemoteTrainingSystemsItem, err error) {
	_model = &FederatedLearningRemoteTrainingRemoteTrainingSystemsItem{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalFederatedLearningRemoteTrainingRemoteTrainingSystemsItem unmarshals an instance of FederatedLearningRemoteTrainingRemoteTrainingSystemsItem from the specified map of raw messages.
func UnmarshalFederatedLearningRemoteTrainingRemoteTrainingSystemsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningRemoteTrainingRemoteTrainingSystemsItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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

// FunctionEntitySchemas : The schemas of the expected data.
type FunctionEntitySchemas struct {
	// The schema of the expected input data.
	Input []DataSchema `json:"input,omitempty"`

	// The schema of the expected output data.
	Output []DataSchema `json:"output,omitempty"`
}

// UnmarshalFunctionEntitySchemas unmarshals an instance of FunctionEntitySchemas from the specified map of raw messages.
func UnmarshalFunctionEntitySchemas(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FunctionEntitySchemas)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalDataSchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalDataSchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FunctionResourceEntity : The details of the function to be created.
type FunctionResourceEntity struct {
	// A software specification.
	SoftwareSpec *SoftwareSpecRel `json:"software_spec" validate:"required"`

	// Type of the function to be created. Only `python` type is supported as of now. If not specified, `python` is the
	// default. Functions expect a gzip file that contains a python file that make up the function. Python functions
	// specify what needs to be run when the function is deployed and what needs to be run when the scoring function is
	// called. In other words, you are able to customize what preparation WML does in the environment when you deploy the
	// function, as well as what steps WML takes to generate the output when you call the API later on. The function
	// consists of the outer function (any place that is not within the score function) and the inner score function. The
	// code that sits in the outer function runs when the function is deployed, and the environment is then frozen and
	// ready to be used whenever the online scoring or batch inline job processing API is called. The code that sits in the
	// inner score function runs when the online scoring or batch inline job processing API is called, in the environment
	// customized by the outer function. See [Deploying Python
	// function](https://dataplatform.cloud.ibm.com/docs/content/wsj/analyze-data/ml-deploy-py-function.html?context=cpdaas&audience=wdp)
	// for more details.
	//
	// This is illustrated in the example below:
	//
	// >&lt;python code used to set up the environment&gt; \
	// >\
	// >def score(payload): \
	// >&nbsp;&nbsp;&nbsp;&nbsp;df_payload = pd.DataFrame(payload["values"]) \
	// >&nbsp;&nbsp;&nbsp;&nbsp;df_payload.columns = payload["fields"] \
	// >&nbsp;&nbsp;&nbsp;&nbsp;... \
	// >&nbsp;&nbsp;&nbsp;&nbsp;output = {"result" : res} \
	// >&nbsp;&nbsp;&nbsp;&nbsp;return output \
	// >\
	// >return score.
	Type *string `json:"type,omitempty"`

	// Scoring data.
	SampleScoringInput *SyncScoringData `json:"sample_scoring_input,omitempty"`

	// The schemas of the expected data.
	Schemas *FunctionEntitySchemas `json:"schemas,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// UnmarshalFunctionResourceEntity unmarshals an instance of FunctionResourceEntity from the specified map of raw messages.
func UnmarshalFunctionResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FunctionResourceEntity)
	err = core.UnmarshalModel(m, "software_spec", &obj.SoftwareSpec, UnmarshalSoftwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sample_scoring_input", &obj.SampleScoringInput, UnmarshalSyncScoringData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schemas", &obj.Schemas, UnmarshalFunctionEntitySchemas)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FunctionResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type FunctionResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalFunctionResourceMetadata unmarshals an instance of FunctionResourceMetadata from the specified map of raw messages.
func UnmarshalFunctionResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FunctionResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// FunctionResourcesFirst : The reference to the first item in the current page.
type FunctionResourcesFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalFunctionResourcesFirst unmarshals an instance of FunctionResourcesFirst from the specified map of raw messages.
func UnmarshalFunctionResourcesFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FunctionResourcesFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FunctionResourcesNext : A reference to the first item of the next page, if any.
type FunctionResourcesNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalFunctionResourcesNext unmarshals an instance of FunctionResourcesNext from the specified map of raw messages.
func UnmarshalFunctionResourcesNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FunctionResourcesNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FunctionsCreateOptions : The FunctionsCreate options.
type FunctionsCreateOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// A software specification.
	SoftwareSpec *SoftwareSpecRel `json:"software_spec" validate:"required"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Type of the function to be created. Only `python` type is supported as of now. If not specified, `python` is the
	// default. Functions expect a gzip file that contains a python file that make up the function. Python functions
	// specify what needs to be run when the function is deployed and what needs to be run when the scoring function is
	// called. In other words, you are able to customize what preparation WML does in the environment when you deploy the
	// function, as well as what steps WML takes to generate the output when you call the API later on. The function
	// consists of the outer function (any place that is not within the score function) and the inner score function. The
	// code that sits in the outer function runs when the function is deployed, and the environment is then frozen and
	// ready to be used whenever the online scoring or batch inline job processing API is called. The code that sits in the
	// inner score function runs when the online scoring or batch inline job processing API is called, in the environment
	// customized by the outer function. See [Deploying Python
	// function](https://dataplatform.cloud.ibm.com/docs/content/wsj/analyze-data/ml-deploy-py-function.html?context=cpdaas&audience=wdp)
	// for more details.
	//
	// This is illustrated in the example below:
	//
	// >&lt;python code used to set up the environment&gt; \
	// >\
	// >def score(payload): \
	// >&nbsp;&nbsp;&nbsp;&nbsp;df_payload = pd.DataFrame(payload["values"]) \
	// >&nbsp;&nbsp;&nbsp;&nbsp;df_payload.columns = payload["fields"] \
	// >&nbsp;&nbsp;&nbsp;&nbsp;... \
	// >&nbsp;&nbsp;&nbsp;&nbsp;output = {"result" : res} \
	// >&nbsp;&nbsp;&nbsp;&nbsp;return output \
	// >\
	// >return score.
	Type *string `json:"type,omitempty"`

	// Scoring data.
	SampleScoringInput *SyncScoringData `json:"sample_scoring_input,omitempty"`

	// The schemas of the expected data.
	Schemas *FunctionEntitySchemas `json:"schemas,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsCreateOptions : Instantiate FunctionsCreateOptions
func (*WatsonMachineLearningV4) NewFunctionsCreateOptions(name string, softwareSpec *SoftwareSpecRel) *FunctionsCreateOptions {
	return &FunctionsCreateOptions{
		Name:         core.StringPtr(name),
		SoftwareSpec: softwareSpec,
	}
}

// SetName : Allow user to set Name
func (_options *FunctionsCreateOptions) SetName(name string) *FunctionsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetSoftwareSpec : Allow user to set SoftwareSpec
func (_options *FunctionsCreateOptions) SetSoftwareSpec(softwareSpec *SoftwareSpecRel) *FunctionsCreateOptions {
	_options.SoftwareSpec = softwareSpec
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsCreateOptions) SetProjectID(projectID string) *FunctionsCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsCreateOptions) SetSpaceID(spaceID string) *FunctionsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *FunctionsCreateOptions) SetDescription(description string) *FunctionsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *FunctionsCreateOptions) SetTags(tags []string) *FunctionsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetType : Allow user to set Type
func (_options *FunctionsCreateOptions) SetType(typeVar string) *FunctionsCreateOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetSampleScoringInput : Allow user to set SampleScoringInput
func (_options *FunctionsCreateOptions) SetSampleScoringInput(sampleScoringInput *SyncScoringData) *FunctionsCreateOptions {
	_options.SampleScoringInput = sampleScoringInput
	return _options
}

// SetSchemas : Allow user to set Schemas
func (_options *FunctionsCreateOptions) SetSchemas(schemas *FunctionEntitySchemas) *FunctionsCreateOptions {
	_options.Schemas = schemas
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *FunctionsCreateOptions) SetCustom(custom map[string]interface{}) *FunctionsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsCreateOptions) SetHeaders(param map[string]string) *FunctionsCreateOptions {
	options.Headers = param
	return options
}

// FunctionsCreateRevisionOptions : The FunctionsCreateRevision options.
type FunctionsCreateRevisionOptions struct {
	// Function identifier.
	FunctionID *string `json:"function_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsCreateRevisionOptions : Instantiate FunctionsCreateRevisionOptions
func (*WatsonMachineLearningV4) NewFunctionsCreateRevisionOptions(functionID string) *FunctionsCreateRevisionOptions {
	return &FunctionsCreateRevisionOptions{
		FunctionID: core.StringPtr(functionID),
	}
}

// SetFunctionID : Allow user to set FunctionID
func (_options *FunctionsCreateRevisionOptions) SetFunctionID(functionID string) *FunctionsCreateRevisionOptions {
	_options.FunctionID = core.StringPtr(functionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsCreateRevisionOptions) SetSpaceID(spaceID string) *FunctionsCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsCreateRevisionOptions) SetProjectID(projectID string) *FunctionsCreateRevisionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *FunctionsCreateRevisionOptions) SetCommitMessage(commitMessage string) *FunctionsCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsCreateRevisionOptions) SetHeaders(param map[string]string) *FunctionsCreateRevisionOptions {
	options.Headers = param
	return options
}

// FunctionsDeleteOptions : The FunctionsDelete options.
type FunctionsDeleteOptions struct {
	// Function identifier.
	FunctionID *string `json:"function_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsDeleteOptions : Instantiate FunctionsDeleteOptions
func (*WatsonMachineLearningV4) NewFunctionsDeleteOptions(functionID string) *FunctionsDeleteOptions {
	return &FunctionsDeleteOptions{
		FunctionID: core.StringPtr(functionID),
	}
}

// SetFunctionID : Allow user to set FunctionID
func (_options *FunctionsDeleteOptions) SetFunctionID(functionID string) *FunctionsDeleteOptions {
	_options.FunctionID = core.StringPtr(functionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsDeleteOptions) SetSpaceID(spaceID string) *FunctionsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsDeleteOptions) SetProjectID(projectID string) *FunctionsDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsDeleteOptions) SetHeaders(param map[string]string) *FunctionsDeleteOptions {
	options.Headers = param
	return options
}

// FunctionsDownloadCodeOptions : The FunctionsDownloadCode options.
type FunctionsDownloadCodeOptions struct {
	// Function identifier.
	FunctionID *string `json:"function_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsDownloadCodeOptions : Instantiate FunctionsDownloadCodeOptions
func (*WatsonMachineLearningV4) NewFunctionsDownloadCodeOptions(functionID string) *FunctionsDownloadCodeOptions {
	return &FunctionsDownloadCodeOptions{
		FunctionID: core.StringPtr(functionID),
	}
}

// SetFunctionID : Allow user to set FunctionID
func (_options *FunctionsDownloadCodeOptions) SetFunctionID(functionID string) *FunctionsDownloadCodeOptions {
	_options.FunctionID = core.StringPtr(functionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsDownloadCodeOptions) SetSpaceID(spaceID string) *FunctionsDownloadCodeOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsDownloadCodeOptions) SetProjectID(projectID string) *FunctionsDownloadCodeOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *FunctionsDownloadCodeOptions) SetRev(rev string) *FunctionsDownloadCodeOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsDownloadCodeOptions) SetHeaders(param map[string]string) *FunctionsDownloadCodeOptions {
	options.Headers = param
	return options
}

// FunctionsGetOptions : The FunctionsGet options.
type FunctionsGetOptions struct {
	// Function identifier.
	FunctionID *string `json:"function_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsGetOptions : Instantiate FunctionsGetOptions
func (*WatsonMachineLearningV4) NewFunctionsGetOptions(functionID string) *FunctionsGetOptions {
	return &FunctionsGetOptions{
		FunctionID: core.StringPtr(functionID),
	}
}

// SetFunctionID : Allow user to set FunctionID
func (_options *FunctionsGetOptions) SetFunctionID(functionID string) *FunctionsGetOptions {
	_options.FunctionID = core.StringPtr(functionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsGetOptions) SetSpaceID(spaceID string) *FunctionsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsGetOptions) SetProjectID(projectID string) *FunctionsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *FunctionsGetOptions) SetRev(rev string) *FunctionsGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsGetOptions) SetHeaders(param map[string]string) *FunctionsGetOptions {
	options.Headers = param
	return options
}

// FunctionsListOptions : The FunctionsList options.
type FunctionsListOptions struct {
	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsListOptions : Instantiate FunctionsListOptions
func (*WatsonMachineLearningV4) NewFunctionsListOptions() *FunctionsListOptions {
	return &FunctionsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsListOptions) SetSpaceID(spaceID string) *FunctionsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsListOptions) SetProjectID(projectID string) *FunctionsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *FunctionsListOptions) SetStart(start string) *FunctionsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *FunctionsListOptions) SetLimit(limit int64) *FunctionsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *FunctionsListOptions) SetTagValue(tagValue string) *FunctionsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *FunctionsListOptions) SetSearch(search string) *FunctionsListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsListOptions) SetHeaders(param map[string]string) *FunctionsListOptions {
	options.Headers = param
	return options
}

// FunctionsListRevisionsOptions : The FunctionsListRevisions options.
type FunctionsListRevisionsOptions struct {
	// Function identifier.
	FunctionID *string `json:"function_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsListRevisionsOptions : Instantiate FunctionsListRevisionsOptions
func (*WatsonMachineLearningV4) NewFunctionsListRevisionsOptions(functionID string) *FunctionsListRevisionsOptions {
	return &FunctionsListRevisionsOptions{
		FunctionID: core.StringPtr(functionID),
	}
}

// SetFunctionID : Allow user to set FunctionID
func (_options *FunctionsListRevisionsOptions) SetFunctionID(functionID string) *FunctionsListRevisionsOptions {
	_options.FunctionID = core.StringPtr(functionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsListRevisionsOptions) SetSpaceID(spaceID string) *FunctionsListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsListRevisionsOptions) SetProjectID(projectID string) *FunctionsListRevisionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *FunctionsListRevisionsOptions) SetStart(start string) *FunctionsListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *FunctionsListRevisionsOptions) SetLimit(limit int64) *FunctionsListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsListRevisionsOptions) SetHeaders(param map[string]string) *FunctionsListRevisionsOptions {
	options.Headers = param
	return options
}

// FunctionsUpdateOptions : The FunctionsUpdate options.
type FunctionsUpdateOptions struct {
	// Function identifier.
	FunctionID *string `json:"function_id" validate:"required,ne="`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsUpdateOptions : Instantiate FunctionsUpdateOptions
func (*WatsonMachineLearningV4) NewFunctionsUpdateOptions(functionID string, jsonPatch []JSONPatchOperation) *FunctionsUpdateOptions {
	return &FunctionsUpdateOptions{
		FunctionID: core.StringPtr(functionID),
		JSONPatch:  jsonPatch,
	}
}

// SetFunctionID : Allow user to set FunctionID
func (_options *FunctionsUpdateOptions) SetFunctionID(functionID string) *FunctionsUpdateOptions {
	_options.FunctionID = core.StringPtr(functionID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *FunctionsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *FunctionsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsUpdateOptions) SetSpaceID(spaceID string) *FunctionsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsUpdateOptions) SetProjectID(projectID string) *FunctionsUpdateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsUpdateOptions) SetHeaders(param map[string]string) *FunctionsUpdateOptions {
	options.Headers = param
	return options
}

// FunctionsUploadCodeOptions : The FunctionsUploadCode options.
type FunctionsUploadCodeOptions struct {
	// Function identifier.
	FunctionID *string `json:"function_id" validate:"required,ne="`

	// A gzip file containing code files.
	UploadCode io.ReadCloser `json:"upload-code" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewFunctionsUploadCodeOptions : Instantiate FunctionsUploadCodeOptions
func (*WatsonMachineLearningV4) NewFunctionsUploadCodeOptions(functionID string, uploadCode io.ReadCloser) *FunctionsUploadCodeOptions {
	return &FunctionsUploadCodeOptions{
		FunctionID: core.StringPtr(functionID),
		UploadCode: uploadCode,
	}
}

// SetFunctionID : Allow user to set FunctionID
func (_options *FunctionsUploadCodeOptions) SetFunctionID(functionID string) *FunctionsUploadCodeOptions {
	_options.FunctionID = core.StringPtr(functionID)
	return _options
}

// SetUploadCode : Allow user to set UploadCode
func (_options *FunctionsUploadCodeOptions) SetUploadCode(uploadCode io.ReadCloser) *FunctionsUploadCodeOptions {
	_options.UploadCode = uploadCode
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *FunctionsUploadCodeOptions) SetSpaceID(spaceID string) *FunctionsUploadCodeOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *FunctionsUploadCodeOptions) SetProjectID(projectID string) *FunctionsUploadCodeOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *FunctionsUploadCodeOptions) SetHeaders(param map[string]string) *FunctionsUploadCodeOptions {
	options.Headers = param
	return options
}

// GpuMetricsMemory : GpuMetricsMemory struct
type GpuMetricsMemory struct {
	Measure *string `json:"measure" validate:"required"`

	Value *float64 `json:"value" validate:"required"`
}

// UnmarshalGpuMetricsMemory unmarshals an instance of GpuMetricsMemory from the specified map of raw messages.
func UnmarshalGpuMetricsMemory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GpuMetricsMemory)
	err = core.UnmarshalPrimitive(m, "measure", &obj.Measure)
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

// InputDataArray : The input data.
type InputDataArray struct {
	// Discriminates the data for multi input data situation. For example in cases where multiple tensors are expected.
	ID *string `json:"id,omitempty"`

	// The names of the fields. The order of fields values must be consistent with the order of fields names.
	Fields []string `json:"fields,omitempty"`

	// Input data as a a vector for a single record or a matrix representing a mini batch of records.
	Values [][]interface{} `json:"values,omitempty"`
}

// UnmarshalInputDataArray unmarshals an instance of InputDataArray from the specified map of raw messages.
func UnmarshalInputDataArray(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InputDataArray)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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

// InstanceResourceEntityPlan : InstanceResourceEntityPlan struct
type InstanceResourceEntityPlan struct {
	// The payment plan ID.
	ID *string `json:"id" validate:"required"`

	// The payment plan name.
	Name *string `json:"name" validate:"required"`

	// 1 - for v1 plans, 2 - for the v2 plans where v2 plan means an instance is space / project aware.
	Version *int64 `json:"version" validate:"required"`
}

// Constants associated with the InstanceResourceEntityPlan.Name property.
// The payment plan name.
const (
	InstanceResourceEntityPlan_Name_Lite         = "lite"
	InstanceResourceEntityPlan_Name_Professional = "professional"
	InstanceResourceEntityPlan_Name_Standard     = "standard"
)

// UnmarshalInstanceResourceEntityPlan unmarshals an instance of InstanceResourceEntityPlan from the specified map of raw messages.
func UnmarshalInstanceResourceEntityPlan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceResourceEntityPlan)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstancesGetOptions : The InstancesGet options.
type InstancesGetOptions struct {
	// The instance identifier.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Defines authorization context that allows a space member to perform this operation.
	SpaceID *string `json:"space_id,omitempty"`

	// Defines authorization context that allows a project member to perform this operation.
	ProjectID *string `json:"project_id,omitempty"`

	// Only if `consumption_details` is set to `true` the instance `entity.consumption.details` part is available in
	// response.
	ConsumptionDetails *bool `json:"consumption_details,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstancesGetOptions : Instantiate InstancesGetOptions
func (*WatsonMachineLearningV4) NewInstancesGetOptions(instanceID string) *InstancesGetOptions {
	return &InstancesGetOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *InstancesGetOptions) SetInstanceID(instanceID string) *InstancesGetOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *InstancesGetOptions) SetSpaceID(spaceID string) *InstancesGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *InstancesGetOptions) SetProjectID(projectID string) *InstancesGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetConsumptionDetails : Allow user to set ConsumptionDetails
func (_options *InstancesGetOptions) SetConsumptionDetails(consumptionDetails bool) *InstancesGetOptions {
	_options.ConsumptionDetails = core.BoolPtr(consumptionDetails)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstancesGetOptions) SetHeaders(param map[string]string) *InstancesGetOptions {
	options.Headers = param
	return options
}

// InstancesListOptions : The InstancesList options.
type InstancesListOptions struct {
	// Return resources pertaining to any space_id listed.  Either `space_id` or `project_id` query parameter is mandatory
	// but both can be provided.
	SpaceID []string `json:"space_id,omitempty"`

	// Return resources pertaining to any project_id listed.  Either `space_id` or `project_id` query parameter is
	// mandatory but both can be provided.
	ProjectID []string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Only if `consumption_details` is set to `true` each instance `entity.consumption.details` part is available in
	// response.
	ConsumptionDetails *bool `json:"consumption_details,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstancesListOptions : Instantiate InstancesListOptions
func (*WatsonMachineLearningV4) NewInstancesListOptions() *InstancesListOptions {
	return &InstancesListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *InstancesListOptions) SetSpaceID(spaceID []string) *InstancesListOptions {
	_options.SpaceID = spaceID
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *InstancesListOptions) SetProjectID(projectID []string) *InstancesListOptions {
	_options.ProjectID = projectID
	return _options
}

// SetStart : Allow user to set Start
func (_options *InstancesListOptions) SetStart(start string) *InstancesListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *InstancesListOptions) SetLimit(limit int64) *InstancesListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetConsumptionDetails : Allow user to set ConsumptionDetails
func (_options *InstancesListOptions) SetConsumptionDetails(consumptionDetails bool) *InstancesListOptions {
	_options.ConsumptionDetails = core.BoolPtr(consumptionDetails)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstancesListOptions) SetHeaders(param map[string]string) *InstancesListOptions {
	options.Headers = param
	return options
}

// JobEntityRequestHybridPipelineHardwareSpecsItem : JobEntityRequestHybridPipelineHardwareSpecsItem struct
type JobEntityRequestHybridPipelineHardwareSpecsItem struct {
	// The id of node runtime.
	NodeRuntimeID *string `json:"node_runtime_id" validate:"required"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec" validate:"required"`
}

// NewJobEntityRequestHybridPipelineHardwareSpecsItem : Instantiate JobEntityRequestHybridPipelineHardwareSpecsItem (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewJobEntityRequestHybridPipelineHardwareSpecsItem(nodeRuntimeID string, hardwareSpec *HardwareSpecRel) (_model *JobEntityRequestHybridPipelineHardwareSpecsItem, err error) {
	_model = &JobEntityRequestHybridPipelineHardwareSpecsItem{
		NodeRuntimeID: core.StringPtr(nodeRuntimeID),
		HardwareSpec:  hardwareSpec,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJobEntityRequestHybridPipelineHardwareSpecsItem unmarshals an instance of JobEntityRequestHybridPipelineHardwareSpecsItem from the specified map of raw messages.
func UnmarshalJobEntityRequestHybridPipelineHardwareSpecsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobEntityRequestHybridPipelineHardwareSpecsItem)
	err = core.UnmarshalPrimitive(m, "node_runtime_id", &obj.NodeRuntimeID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobEntityResultHybridPipelineHardwareSpecsItem : JobEntityResultHybridPipelineHardwareSpecsItem struct
type JobEntityResultHybridPipelineHardwareSpecsItem struct {
	// The id of node runtime.
	NodeRuntimeID *string `json:"node_runtime_id" validate:"required"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec" validate:"required"`
}

// NewJobEntityResultHybridPipelineHardwareSpecsItem : Instantiate JobEntityResultHybridPipelineHardwareSpecsItem (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewJobEntityResultHybridPipelineHardwareSpecsItem(nodeRuntimeID string, hardwareSpec *HardwareSpecRel) (_model *JobEntityResultHybridPipelineHardwareSpecsItem, err error) {
	_model = &JobEntityResultHybridPipelineHardwareSpecsItem{
		NodeRuntimeID: core.StringPtr(nodeRuntimeID),
		HardwareSpec:  hardwareSpec,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJobEntityResultHybridPipelineHardwareSpecsItem unmarshals an instance of JobEntityResultHybridPipelineHardwareSpecsItem from the specified map of raw messages.
func UnmarshalJobEntityResultHybridPipelineHardwareSpecsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobEntityResultHybridPipelineHardwareSpecsItem)
	err = core.UnmarshalPrimitive(m, "node_runtime_id", &obj.NodeRuntimeID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobResourceEntity : Details about the batch deployment job.
//
// The `deployment` is a reference to the deployment associated with the deployment job or deployment job definition.
//
// The `scoring` and `decision_optimization` properties are mutually exclusive. Specify only one of these when
// submitting a batch deployment job.
//
// Use `hybrid_pipeline_hardware_specs` only in a batch deployment job of a hybrid pipeline in order to specify compute
// configuration for each pipeline node. For all other cases use `hardware_spec` to specify compute configuration.
//
// In case of output data references where the data asset is a remote database, users can specify if the batch
// deployment output must be appended to the table or if the table is to be truncated and output data updated.
// `output_data_references.location.write_mode` parameter can be used to for this purpose. The values `truncate` or
// `append` can be specified for `output_data_references.location.write_mode` parameter.
// * Specifying `truncate` as value will truncate the table and the batch output data will be inserted.
// * Specifying `append` as value will insert the batch output data to the remote database table.
// * The `write_mode` parameter is applicable only for `output_data_references` parameter.
// * The `write_mode` parameter will be applicable only for remote database related data assets. This parameter will not
// be applicable for local data asset or COS based data asset.
type JobResourceEntity struct {
	// A reference to a resource.
	Deployment *SimpleRel `json:"deployment" validate:"required"`

	// User defined properties.
	Custom interface{} `json:"custom,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityResultHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`

	// Details about the input/output data and other properties to be used for a batch deployment job of a model,
	// Python Function or a Python Scripts.
	//
	// Use `input_data` property to specify the input data for batch processing as part of the job's payload.
	// The `input_data` property is mutually exclusive with `input_data_references` property, only use one of these.
	// When `input_data` is specified, the processed output of batch deployment job will be available in
	// `scoring.predictions`
	// parameter in the deployment job response. `input_data` property is not supported for batch deployment of Python
	// Scripts.
	//
	// Use `input_data_references` property to specify the details pertaining to the remote source where the input data for
	// batch deployment job is available. The `input_data_references` must be used with `output_data_references`.
	// The `input_data_references` property
	// is mutually exclusive with `input_data` property, only use one of these. The `input_data_references`
	// property is not supported for batch deployment job of Spark models and Python Functions.
	//
	// Use `output_data_references` property to specify the details pertaining to the remote source where the input
	// data for batch deployment job is available. `output_data_references` must be used with `input_data_references`.
	// The `output_data_references`
	// property is not supported for batch deployment job of Spark models and Python Functions.
	Scoring *JobScoringRequest `json:"scoring,omitempty"`

	// Details about the input/output data and other properties to be used for a batch
	// deployment job of a decision optimization problem.
	//
	// Use `input_data` property to specify the input data for batch processing as part
	// of the job's payload. The `input_data` property is mutually exclusive with `input_data_references`
	// property, only use one of these. When `input_data` is specified, the processed output
	// of batch deployment job will be available in `scoring.predictions` parameter in the deployment job response.
	//
	// Use `input_data_references` property to specify the details pertaining to the remote source
	// where the input data for batch deployment job is available. The `input_data_references` must be used
	// with `output_data_references`.
	// The `input_data_references` property is mutually exclusive with `input_data` property, only use one of these.
	//
	// Use `output_data_references` property to specify the details pertaining to the remote source where the
	// input data for batch deployment job is available. The `output_data_references` must be used with
	// `input_data_references`.
	DecisionOptimization *JobDecisionOptimizationRequest `json:"decision_optimization,omitempty"`
}

// UnmarshalJobResourceEntity unmarshals an instance of JobResourceEntity from the specified map of raw messages.
func UnmarshalJobResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobResourceEntity)
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalSimpleRel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hybrid_pipeline_hardware_specs", &obj.HybridPipelineHardwareSpecs, UnmarshalJobEntityResultHybridPipelineHardwareSpecsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scoring", &obj.Scoring, UnmarshalJobScoringRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "decision_optimization", &obj.DecisionOptimization, UnmarshalJobDecisionOptimizationRequest)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type JobResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalJobResourceMetadata unmarshals an instance of JobResourceMetadata from the specified map of raw messages.
func UnmarshalJobResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// JobResourcesFirst : The reference to the first item in the current page.
type JobResourcesFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalJobResourcesFirst unmarshals an instance of JobResourcesFirst from the specified map of raw messages.
func UnmarshalJobResourcesFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobResourcesFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobResourcesNext : A reference to the first item of the next page, if any.
type JobResourcesNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalJobResourcesNext unmarshals an instance of JobResourcesNext from the specified map of raw messages.
func UnmarshalJobResourcesNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobResourcesNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobScoringResultEvaluationsItem : An evaluation specification used to support evaluations for TensorFlow.
type JobScoringResultEvaluationsItem struct {
	// An identifier of this metrics set. For a DL problem this can be the output tensor id/name in order to identify on
	// which output these metrics will be computed.
	ID *string `json:"id,omitempty"`

	// The id of the `input_data.id` with the `type=target`. This points to the ground truth information that will be used
	// together with prediction information to generate metrics.
	InputTarget *string `json:"input_target,omitempty"`

	// A list of the metric names.
	MetricsNames []string `json:"metrics_names,omitempty"`
}

// UnmarshalJobScoringResultEvaluationsItem unmarshals an instance of JobScoringResultEvaluationsItem from the specified map of raw messages.
func UnmarshalJobScoringResultEvaluationsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobScoringResultEvaluationsItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "input_target", &obj.InputTarget)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metrics_names", &obj.MetricsNames)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobStatusMessage : An optional message related to the job status.
type JobStatusMessage struct {
	// The level of the message, normally one of `debug`, `info` or `warning`.
	Level *string `json:"level,omitempty"`

	// The message.
	Text *string `json:"text,omitempty"`
}

// UnmarshalJobStatusMessage unmarshals an instance of JobStatusMessage from the specified map of raw messages.
func UnmarshalJobStatusMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobStatusMessage)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobsResourcesFirst : The reference to the first item in the current page.
type JobsResourcesFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalJobsResourcesFirst unmarshals an instance of JobsResourcesFirst from the specified map of raw messages.
func UnmarshalJobsResourcesFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobsResourcesFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobsResourcesNext : A reference to the first item of the next page, if any.
type JobsResourcesNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalJobsResourcesNext unmarshals an instance of JobsResourcesNext from the specified map of raw messages.
func UnmarshalJobsResourcesNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobsResourcesNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelDefinitionEntityRequestPlatform : ModelDefinitionEntityRequestPlatform struct
type ModelDefinitionEntityRequestPlatform struct {
	// The name of the platform.
	Name *string `json:"name" validate:"required"`

	// The supported versions.
	Versions []string `json:"versions" validate:"required"`
}

// NewModelDefinitionEntityRequestPlatform : Instantiate ModelDefinitionEntityRequestPlatform (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewModelDefinitionEntityRequestPlatform(name string, versions []string) (_model *ModelDefinitionEntityRequestPlatform, err error) {
	_model = &ModelDefinitionEntityRequestPlatform{
		Name:     core.StringPtr(name),
		Versions: versions,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalModelDefinitionEntityRequestPlatform unmarshals an instance of ModelDefinitionEntityRequestPlatform from the specified map of raw messages.
func UnmarshalModelDefinitionEntityRequestPlatform(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionEntityRequestPlatform)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "versions", &obj.Versions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelDefinitionResourceEntity : ModelDefinitionResourceEntity struct
type ModelDefinitionResourceEntity struct {
	// The package version.
	Version *string `json:"version" validate:"required"`

	Platform *ModelDefinitionResourceEntityPlatform `json:"platform" validate:"required"`

	// The command used to run the model.
	Command *string `json:"command,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// UnmarshalModelDefinitionResourceEntity unmarshals an instance of ModelDefinitionResourceEntity from the specified map of raw messages.
func UnmarshalModelDefinitionResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionResourceEntity)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "platform", &obj.Platform, UnmarshalModelDefinitionResourceEntityPlatform)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "command", &obj.Command)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelDefinitionResourceEntityPlatform : ModelDefinitionResourceEntityPlatform struct
type ModelDefinitionResourceEntityPlatform struct {
	// The name of the platform.
	Name *string `json:"name" validate:"required"`

	// The supported versions.
	Versions []string `json:"versions" validate:"required"`
}

// UnmarshalModelDefinitionResourceEntityPlatform unmarshals an instance of ModelDefinitionResourceEntityPlatform from the specified map of raw messages.
func UnmarshalModelDefinitionResourceEntityPlatform(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionResourceEntityPlatform)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "versions", &obj.Versions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelDefinitionResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type ModelDefinitionResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalModelDefinitionResourceMetadata unmarshals an instance of ModelDefinitionResourceMetadata from the specified map of raw messages.
func UnmarshalModelDefinitionResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// ModelDefinitionsCreateOptions : The ModelDefinitionsCreate options.
type ModelDefinitionsCreateOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// The package version.
	Version *string `json:"version" validate:"required"`

	Platform *ModelDefinitionEntityRequestPlatform `json:"platform" validate:"required"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The command used to run the model.
	Command *string `json:"command,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsCreateOptions : Instantiate ModelDefinitionsCreateOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsCreateOptions(name string, version string, platform *ModelDefinitionEntityRequestPlatform) *ModelDefinitionsCreateOptions {
	return &ModelDefinitionsCreateOptions{
		Name:     core.StringPtr(name),
		Version:  core.StringPtr(version),
		Platform: platform,
	}
}

// SetName : Allow user to set Name
func (_options *ModelDefinitionsCreateOptions) SetName(name string) *ModelDefinitionsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *ModelDefinitionsCreateOptions) SetVersion(version string) *ModelDefinitionsCreateOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetPlatform : Allow user to set Platform
func (_options *ModelDefinitionsCreateOptions) SetPlatform(platform *ModelDefinitionEntityRequestPlatform) *ModelDefinitionsCreateOptions {
	_options.Platform = platform
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsCreateOptions) SetProjectID(projectID string) *ModelDefinitionsCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsCreateOptions) SetSpaceID(spaceID string) *ModelDefinitionsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ModelDefinitionsCreateOptions) SetDescription(description string) *ModelDefinitionsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *ModelDefinitionsCreateOptions) SetTags(tags []string) *ModelDefinitionsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetCommand : Allow user to set Command
func (_options *ModelDefinitionsCreateOptions) SetCommand(command string) *ModelDefinitionsCreateOptions {
	_options.Command = core.StringPtr(command)
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *ModelDefinitionsCreateOptions) SetCustom(custom map[string]interface{}) *ModelDefinitionsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsCreateOptions) SetHeaders(param map[string]string) *ModelDefinitionsCreateOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsCreateRevisionOptions : The ModelDefinitionsCreateRevision options.
type ModelDefinitionsCreateRevisionOptions struct {
	// Model definition identifier.
	ModelDefinitionID *string `json:"model_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsCreateRevisionOptions : Instantiate ModelDefinitionsCreateRevisionOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsCreateRevisionOptions(modelDefinitionID string) *ModelDefinitionsCreateRevisionOptions {
	return &ModelDefinitionsCreateRevisionOptions{
		ModelDefinitionID: core.StringPtr(modelDefinitionID),
	}
}

// SetModelDefinitionID : Allow user to set ModelDefinitionID
func (_options *ModelDefinitionsCreateRevisionOptions) SetModelDefinitionID(modelDefinitionID string) *ModelDefinitionsCreateRevisionOptions {
	_options.ModelDefinitionID = core.StringPtr(modelDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsCreateRevisionOptions) SetSpaceID(spaceID string) *ModelDefinitionsCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsCreateRevisionOptions) SetProjectID(projectID string) *ModelDefinitionsCreateRevisionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *ModelDefinitionsCreateRevisionOptions) SetCommitMessage(commitMessage string) *ModelDefinitionsCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsCreateRevisionOptions) SetHeaders(param map[string]string) *ModelDefinitionsCreateRevisionOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsDeleteOptions : The ModelDefinitionsDelete options.
type ModelDefinitionsDeleteOptions struct {
	// Model definition identifier.
	ModelDefinitionID *string `json:"model_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsDeleteOptions : Instantiate ModelDefinitionsDeleteOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsDeleteOptions(modelDefinitionID string) *ModelDefinitionsDeleteOptions {
	return &ModelDefinitionsDeleteOptions{
		ModelDefinitionID: core.StringPtr(modelDefinitionID),
	}
}

// SetModelDefinitionID : Allow user to set ModelDefinitionID
func (_options *ModelDefinitionsDeleteOptions) SetModelDefinitionID(modelDefinitionID string) *ModelDefinitionsDeleteOptions {
	_options.ModelDefinitionID = core.StringPtr(modelDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsDeleteOptions) SetSpaceID(spaceID string) *ModelDefinitionsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsDeleteOptions) SetProjectID(projectID string) *ModelDefinitionsDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsDeleteOptions) SetHeaders(param map[string]string) *ModelDefinitionsDeleteOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsDownloadModelOptions : The ModelDefinitionsDownloadModel options.
type ModelDefinitionsDownloadModelOptions struct {
	// Model definition identifier.
	ModelDefinitionID *string `json:"model_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsDownloadModelOptions : Instantiate ModelDefinitionsDownloadModelOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsDownloadModelOptions(modelDefinitionID string) *ModelDefinitionsDownloadModelOptions {
	return &ModelDefinitionsDownloadModelOptions{
		ModelDefinitionID: core.StringPtr(modelDefinitionID),
	}
}

// SetModelDefinitionID : Allow user to set ModelDefinitionID
func (_options *ModelDefinitionsDownloadModelOptions) SetModelDefinitionID(modelDefinitionID string) *ModelDefinitionsDownloadModelOptions {
	_options.ModelDefinitionID = core.StringPtr(modelDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsDownloadModelOptions) SetSpaceID(spaceID string) *ModelDefinitionsDownloadModelOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsDownloadModelOptions) SetProjectID(projectID string) *ModelDefinitionsDownloadModelOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *ModelDefinitionsDownloadModelOptions) SetRev(rev string) *ModelDefinitionsDownloadModelOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsDownloadModelOptions) SetHeaders(param map[string]string) *ModelDefinitionsDownloadModelOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsGetOptions : The ModelDefinitionsGet options.
type ModelDefinitionsGetOptions struct {
	// Model definition identifier.
	ModelDefinitionID *string `json:"model_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsGetOptions : Instantiate ModelDefinitionsGetOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsGetOptions(modelDefinitionID string) *ModelDefinitionsGetOptions {
	return &ModelDefinitionsGetOptions{
		ModelDefinitionID: core.StringPtr(modelDefinitionID),
	}
}

// SetModelDefinitionID : Allow user to set ModelDefinitionID
func (_options *ModelDefinitionsGetOptions) SetModelDefinitionID(modelDefinitionID string) *ModelDefinitionsGetOptions {
	_options.ModelDefinitionID = core.StringPtr(modelDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsGetOptions) SetSpaceID(spaceID string) *ModelDefinitionsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsGetOptions) SetProjectID(projectID string) *ModelDefinitionsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *ModelDefinitionsGetOptions) SetRev(rev string) *ModelDefinitionsGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsGetOptions) SetHeaders(param map[string]string) *ModelDefinitionsGetOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsListOptions : The ModelDefinitionsList options.
type ModelDefinitionsListOptions struct {
	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsListOptions : Instantiate ModelDefinitionsListOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsListOptions() *ModelDefinitionsListOptions {
	return &ModelDefinitionsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsListOptions) SetSpaceID(spaceID string) *ModelDefinitionsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsListOptions) SetProjectID(projectID string) *ModelDefinitionsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ModelDefinitionsListOptions) SetStart(start string) *ModelDefinitionsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ModelDefinitionsListOptions) SetLimit(limit int64) *ModelDefinitionsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *ModelDefinitionsListOptions) SetTagValue(tagValue string) *ModelDefinitionsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *ModelDefinitionsListOptions) SetSearch(search string) *ModelDefinitionsListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsListOptions) SetHeaders(param map[string]string) *ModelDefinitionsListOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsListRevisionsOptions : The ModelDefinitionsListRevisions options.
type ModelDefinitionsListRevisionsOptions struct {
	// Model definition identifier.
	ModelDefinitionID *string `json:"model_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsListRevisionsOptions : Instantiate ModelDefinitionsListRevisionsOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsListRevisionsOptions(modelDefinitionID string) *ModelDefinitionsListRevisionsOptions {
	return &ModelDefinitionsListRevisionsOptions{
		ModelDefinitionID: core.StringPtr(modelDefinitionID),
	}
}

// SetModelDefinitionID : Allow user to set ModelDefinitionID
func (_options *ModelDefinitionsListRevisionsOptions) SetModelDefinitionID(modelDefinitionID string) *ModelDefinitionsListRevisionsOptions {
	_options.ModelDefinitionID = core.StringPtr(modelDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsListRevisionsOptions) SetSpaceID(spaceID string) *ModelDefinitionsListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsListRevisionsOptions) SetProjectID(projectID string) *ModelDefinitionsListRevisionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ModelDefinitionsListRevisionsOptions) SetStart(start string) *ModelDefinitionsListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ModelDefinitionsListRevisionsOptions) SetLimit(limit int64) *ModelDefinitionsListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsListRevisionsOptions) SetHeaders(param map[string]string) *ModelDefinitionsListRevisionsOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsUpdateOptions : The ModelDefinitionsUpdate options.
type ModelDefinitionsUpdateOptions struct {
	// Model definition identifier.
	ModelDefinitionID *string `json:"model_definition_id" validate:"required,ne="`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsUpdateOptions : Instantiate ModelDefinitionsUpdateOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsUpdateOptions(modelDefinitionID string, jsonPatch []JSONPatchOperation) *ModelDefinitionsUpdateOptions {
	return &ModelDefinitionsUpdateOptions{
		ModelDefinitionID: core.StringPtr(modelDefinitionID),
		JSONPatch:         jsonPatch,
	}
}

// SetModelDefinitionID : Allow user to set ModelDefinitionID
func (_options *ModelDefinitionsUpdateOptions) SetModelDefinitionID(modelDefinitionID string) *ModelDefinitionsUpdateOptions {
	_options.ModelDefinitionID = core.StringPtr(modelDefinitionID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *ModelDefinitionsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *ModelDefinitionsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsUpdateOptions) SetSpaceID(spaceID string) *ModelDefinitionsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsUpdateOptions) SetProjectID(projectID string) *ModelDefinitionsUpdateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsUpdateOptions) SetHeaders(param map[string]string) *ModelDefinitionsUpdateOptions {
	options.Headers = param
	return options
}

// ModelDefinitionsUploadModelOptions : The ModelDefinitionsUploadModel options.
type ModelDefinitionsUploadModelOptions struct {
	// Model definition identifier.
	ModelDefinitionID *string `json:"model_definition_id" validate:"required,ne="`

	// A gzip file containing code files.
	UploadModel io.ReadCloser `json:"upload-model" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelDefinitionsUploadModelOptions : Instantiate ModelDefinitionsUploadModelOptions
func (*WatsonMachineLearningV4) NewModelDefinitionsUploadModelOptions(modelDefinitionID string, uploadModel io.ReadCloser) *ModelDefinitionsUploadModelOptions {
	return &ModelDefinitionsUploadModelOptions{
		ModelDefinitionID: core.StringPtr(modelDefinitionID),
		UploadModel:       uploadModel,
	}
}

// SetModelDefinitionID : Allow user to set ModelDefinitionID
func (_options *ModelDefinitionsUploadModelOptions) SetModelDefinitionID(modelDefinitionID string) *ModelDefinitionsUploadModelOptions {
	_options.ModelDefinitionID = core.StringPtr(modelDefinitionID)
	return _options
}

// SetUploadModel : Allow user to set UploadModel
func (_options *ModelDefinitionsUploadModelOptions) SetUploadModel(uploadModel io.ReadCloser) *ModelDefinitionsUploadModelOptions {
	_options.UploadModel = uploadModel
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelDefinitionsUploadModelOptions) SetSpaceID(spaceID string) *ModelDefinitionsUploadModelOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelDefinitionsUploadModelOptions) SetProjectID(projectID string) *ModelDefinitionsUploadModelOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelDefinitionsUploadModelOptions) SetHeaders(param map[string]string) *ModelDefinitionsUploadModelOptions {
	options.Headers = param
	return options
}

// ModelEntitySchemas : If the schemas are provided here then they take precedent over any schemas provided in the data references.
type ModelEntitySchemas struct {
	// The schema of the expected input data.
	Input []DataSchema `json:"input,omitempty"`

	// The schema of the expected output data.
	Output []DataSchema `json:"output,omitempty"`
}

// UnmarshalModelEntitySchemas unmarshals an instance of ModelEntitySchemas from the specified map of raw messages.
func UnmarshalModelEntitySchemas(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelEntitySchemas)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalDataSchema)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalDataSchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelEntitySize : This will be used by scoring to record the size of the model.
type ModelEntitySize struct {
	// The memory size of the model.
	InMemory *float64 `json:"in_memory,omitempty"`

	// The size of the model on disk.
	Content *float64 `json:"content,omitempty"`
}

// UnmarshalModelEntitySize unmarshals an instance of ModelEntitySize from the specified map of raw messages.
func UnmarshalModelEntitySize(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelEntitySize)
	err = core.UnmarshalPrimitive(m, "in_memory", &obj.InMemory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "content", &obj.Content)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelResourceEntity : Information related to the upload of the model content.
type ModelResourceEntity struct {
	// The model type. The supported model types can be found in the documentation
	// [here](https://dataplatform.cloud.ibm.com/docs/content/wsj/wmls/wmls-deploy-python-types.html?context=analytics).
	Type *string `json:"type" validate:"required"`

	// A software specification.
	SoftwareSpec *SoftwareSpecRel `json:"software_spec" validate:"required"`

	// A reference to a resource.
	Pipeline *Rel `json:"pipeline,omitempty"`

	// The model definition.
	ModelDefinition *ModelDefinitionID `json:"model_definition,omitempty"`

	// Hyper parameters used for training this model.
	HyperParameters interface{} `json:"hyper_parameters,omitempty"`

	// User provided domain name for this model. For example: sentiment, entity, visual-recognition, finance, retail, real
	// estate etc.
	Domain *string `json:"domain,omitempty"`

	// The training data that was used to create this model.
	TrainingDataReferences []DataConnectionReference `json:"training_data_references,omitempty"`

	// The holdout/test datasets.
	TestDataReferences []DataConnectionReference `json:"test_data_references,omitempty"`

	// If the schemas are provided here then they take precedent over any schemas
	// provided in the data references.
	Schemas *ModelEntitySchemas `json:"schemas,omitempty"`

	// The name of the label column.
	LabelColumn *string `json:"label_column,omitempty"`

	// The name of the  label column seen by the estimator, which may have been transformed by the previous transformers in
	// the pipeline. This is not necessarily the same column as the `label_column` in the initial data set.
	TransformedLabelColumn *string `json:"transformed_label_column,omitempty"`

	// This will be used by scoring to record the size of the model.
	Size *ModelEntitySize `json:"size,omitempty"`

	// Metrics that can be returned by an operation.
	Metrics []Metric `json:"metrics,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// User defined objects referenced by the model. For any user defined class or function used in the model, its name, as
	// referenced in the model, must be specified as the `key` and its fully qualified class or function name must be
	// specified as the `value`. This is applicable for `Tensorflow 2.X` models serialized in `H5` format using the
	// `tf.keras` API.
	UserDefinedObjects map[string]string `json:"user_defined_objects,omitempty"`

	// The upload state.
	ContentImportState *string `json:"content_import_state,omitempty"`

	// This is the array which contains information about the software specifications that are used in the hybrid pipeline
	// (if this model includes a hybrid pipeline).
	HybridPipelineSoftwareSpecs []SoftwareSpecRel `json:"hybrid_pipeline_software_specs,omitempty"`
}

// Constants associated with the ModelResourceEntity.ContentImportState property.
// The upload state.
const (
	ModelResourceEntity_ContentImportState_Completed = "completed"
	ModelResourceEntity_ContentImportState_Failed    = "failed"
	ModelResourceEntity_ContentImportState_Running   = "running"
)

// UnmarshalModelResourceEntity unmarshals an instance of ModelResourceEntity from the specified map of raw messages.
func UnmarshalModelResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelResourceEntity)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "software_spec", &obj.SoftwareSpec, UnmarshalSoftwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pipeline", &obj.Pipeline, UnmarshalRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "model_definition", &obj.ModelDefinition, UnmarshalModelDefinitionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hyper_parameters", &obj.HyperParameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data_references", &obj.TrainingDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "test_data_references", &obj.TestDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schemas", &obj.Schemas, UnmarshalModelEntitySchemas)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label_column", &obj.LabelColumn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transformed_label_column", &obj.TransformedLabelColumn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "size", &obj.Size, UnmarshalModelEntitySize)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_defined_objects", &obj.UserDefinedObjects)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "content_import_state", &obj.ContentImportState)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hybrid_pipeline_software_specs", &obj.HybridPipelineSoftwareSpecs, UnmarshalSoftwareSpecRel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type ModelResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalModelResourceMetadata unmarshals an instance of ModelResourceMetadata from the specified map of raw messages.
func UnmarshalModelResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// ModelsCreateOptions : The ModelsCreate options.
type ModelsCreateOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// The model type. The supported model types can be found in the documentation
	// [here](https://dataplatform.cloud.ibm.com/docs/content/wsj/wmls/wmls-deploy-python-types.html?context=analytics).
	Type *string `json:"type" validate:"required"`

	// A software specification.
	SoftwareSpec *SoftwareSpecRel `json:"software_spec" validate:"required"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// A reference to a resource.
	Pipeline *Rel `json:"pipeline,omitempty"`

	// The model definition.
	ModelDefinition *ModelDefinitionID `json:"model_definition,omitempty"`

	// Hyper parameters used for training this model.
	HyperParameters interface{} `json:"hyper_parameters,omitempty"`

	// User provided domain name for this model. For example: sentiment, entity, visual-recognition, finance, retail, real
	// estate etc.
	Domain *string `json:"domain,omitempty"`

	// The training data that was used to create this model.
	TrainingDataReferences []DataConnectionReference `json:"training_data_references,omitempty"`

	// The holdout/test datasets.
	TestDataReferences []DataConnectionReference `json:"test_data_references,omitempty"`

	// If the schemas are provided here then they take precedent over any schemas
	// provided in the data references.
	Schemas *ModelEntitySchemas `json:"schemas,omitempty"`

	// The name of the label column.
	LabelColumn *string `json:"label_column,omitempty"`

	// The name of the  label column seen by the estimator, which may have been transformed by the previous transformers in
	// the pipeline. This is not necessarily the same column as the `label_column` in the initial data set.
	TransformedLabelColumn *string `json:"transformed_label_column,omitempty"`

	// This will be used by scoring to record the size of the model.
	Size *ModelEntitySize `json:"size,omitempty"`

	// Metrics that can be returned by an operation.
	Metrics []Metric `json:"metrics,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// User defined objects referenced by the model. For any user defined class or function used in the model, its name, as
	// referenced in the model, must be specified as the `key` and its fully qualified class or function name must be
	// specified as the `value`. This is applicable for `Tensorflow 2.X` models serialized in `H5` format using the
	// `tf.keras` API.
	UserDefinedObjects map[string]string `json:"user_defined_objects,omitempty"`

	// Details about the attachments that should be uploaded with this model.
	ContentLocation *ContentLocation `json:"content_location,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsCreateOptions : Instantiate ModelsCreateOptions
func (*WatsonMachineLearningV4) NewModelsCreateOptions(name string, typeVar string, softwareSpec *SoftwareSpecRel) *ModelsCreateOptions {
	return &ModelsCreateOptions{
		Name:         core.StringPtr(name),
		Type:         core.StringPtr(typeVar),
		SoftwareSpec: softwareSpec,
	}
}

// SetName : Allow user to set Name
func (_options *ModelsCreateOptions) SetName(name string) *ModelsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetType : Allow user to set Type
func (_options *ModelsCreateOptions) SetType(typeVar string) *ModelsCreateOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetSoftwareSpec : Allow user to set SoftwareSpec
func (_options *ModelsCreateOptions) SetSoftwareSpec(softwareSpec *SoftwareSpecRel) *ModelsCreateOptions {
	_options.SoftwareSpec = softwareSpec
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsCreateOptions) SetProjectID(projectID string) *ModelsCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsCreateOptions) SetSpaceID(spaceID string) *ModelsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ModelsCreateOptions) SetDescription(description string) *ModelsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *ModelsCreateOptions) SetTags(tags []string) *ModelsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetPipeline : Allow user to set Pipeline
func (_options *ModelsCreateOptions) SetPipeline(pipeline *Rel) *ModelsCreateOptions {
	_options.Pipeline = pipeline
	return _options
}

// SetModelDefinition : Allow user to set ModelDefinition
func (_options *ModelsCreateOptions) SetModelDefinition(modelDefinition *ModelDefinitionID) *ModelsCreateOptions {
	_options.ModelDefinition = modelDefinition
	return _options
}

// SetHyperParameters : Allow user to set HyperParameters
func (_options *ModelsCreateOptions) SetHyperParameters(hyperParameters interface{}) *ModelsCreateOptions {
	_options.HyperParameters = hyperParameters
	return _options
}

// SetDomain : Allow user to set Domain
func (_options *ModelsCreateOptions) SetDomain(domain string) *ModelsCreateOptions {
	_options.Domain = core.StringPtr(domain)
	return _options
}

// SetTrainingDataReferences : Allow user to set TrainingDataReferences
func (_options *ModelsCreateOptions) SetTrainingDataReferences(trainingDataReferences []DataConnectionReference) *ModelsCreateOptions {
	_options.TrainingDataReferences = trainingDataReferences
	return _options
}

// SetTestDataReferences : Allow user to set TestDataReferences
func (_options *ModelsCreateOptions) SetTestDataReferences(testDataReferences []DataConnectionReference) *ModelsCreateOptions {
	_options.TestDataReferences = testDataReferences
	return _options
}

// SetSchemas : Allow user to set Schemas
func (_options *ModelsCreateOptions) SetSchemas(schemas *ModelEntitySchemas) *ModelsCreateOptions {
	_options.Schemas = schemas
	return _options
}

// SetLabelColumn : Allow user to set LabelColumn
func (_options *ModelsCreateOptions) SetLabelColumn(labelColumn string) *ModelsCreateOptions {
	_options.LabelColumn = core.StringPtr(labelColumn)
	return _options
}

// SetTransformedLabelColumn : Allow user to set TransformedLabelColumn
func (_options *ModelsCreateOptions) SetTransformedLabelColumn(transformedLabelColumn string) *ModelsCreateOptions {
	_options.TransformedLabelColumn = core.StringPtr(transformedLabelColumn)
	return _options
}

// SetSize : Allow user to set Size
func (_options *ModelsCreateOptions) SetSize(size *ModelEntitySize) *ModelsCreateOptions {
	_options.Size = size
	return _options
}

// SetMetrics : Allow user to set Metrics
func (_options *ModelsCreateOptions) SetMetrics(metrics []Metric) *ModelsCreateOptions {
	_options.Metrics = metrics
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *ModelsCreateOptions) SetCustom(custom map[string]interface{}) *ModelsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetUserDefinedObjects : Allow user to set UserDefinedObjects
func (_options *ModelsCreateOptions) SetUserDefinedObjects(userDefinedObjects map[string]string) *ModelsCreateOptions {
	_options.UserDefinedObjects = userDefinedObjects
	return _options
}

// SetContentLocation : Allow user to set ContentLocation
func (_options *ModelsCreateOptions) SetContentLocation(contentLocation *ContentLocation) *ModelsCreateOptions {
	_options.ContentLocation = contentLocation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsCreateOptions) SetHeaders(param map[string]string) *ModelsCreateOptions {
	options.Headers = param
	return options
}

// ModelsCreateRevisionOptions : The ModelsCreateRevision options.
type ModelsCreateRevisionOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsCreateRevisionOptions : Instantiate ModelsCreateRevisionOptions
func (*WatsonMachineLearningV4) NewModelsCreateRevisionOptions(modelID string) *ModelsCreateRevisionOptions {
	return &ModelsCreateRevisionOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsCreateRevisionOptions) SetModelID(modelID string) *ModelsCreateRevisionOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsCreateRevisionOptions) SetSpaceID(spaceID string) *ModelsCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsCreateRevisionOptions) SetProjectID(projectID string) *ModelsCreateRevisionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *ModelsCreateRevisionOptions) SetCommitMessage(commitMessage string) *ModelsCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsCreateRevisionOptions) SetHeaders(param map[string]string) *ModelsCreateRevisionOptions {
	options.Headers = param
	return options
}

// ModelsDeleteContentOptions : The ModelsDeleteContent options.
type ModelsDeleteContentOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// Identifier for the attachment for resources that support separate content.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsDeleteContentOptions : Instantiate ModelsDeleteContentOptions
func (*WatsonMachineLearningV4) NewModelsDeleteContentOptions(modelID string, attachmentID string) *ModelsDeleteContentOptions {
	return &ModelsDeleteContentOptions{
		ModelID:      core.StringPtr(modelID),
		AttachmentID: core.StringPtr(attachmentID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsDeleteContentOptions) SetModelID(modelID string) *ModelsDeleteContentOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ModelsDeleteContentOptions) SetAttachmentID(attachmentID string) *ModelsDeleteContentOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsDeleteContentOptions) SetSpaceID(spaceID string) *ModelsDeleteContentOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsDeleteContentOptions) SetProjectID(projectID string) *ModelsDeleteContentOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsDeleteContentOptions) SetHeaders(param map[string]string) *ModelsDeleteContentOptions {
	options.Headers = param
	return options
}

// ModelsDeleteOptions : The ModelsDelete options.
type ModelsDeleteOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsDeleteOptions : Instantiate ModelsDeleteOptions
func (*WatsonMachineLearningV4) NewModelsDeleteOptions(modelID string) *ModelsDeleteOptions {
	return &ModelsDeleteOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsDeleteOptions) SetModelID(modelID string) *ModelsDeleteOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsDeleteOptions) SetSpaceID(spaceID string) *ModelsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsDeleteOptions) SetProjectID(projectID string) *ModelsDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsDeleteOptions) SetHeaders(param map[string]string) *ModelsDeleteOptions {
	options.Headers = param
	return options
}

// ModelsDownloadContentOptions : The ModelsDownloadContent options.
type ModelsDownloadContentOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// Identifier for the attachment for resources that support separate content.
	AttachmentID *string `json:"attachment_id" validate:"required,ne="`

	// The type of the response: application/zip, application/gzip, application/json, text/plain, or application/xml. A
	// character encoding can be specified by including a `charset` parameter. For example, 'text/plain;charset=utf-8'.
	Accept *string `json:"Accept,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsDownloadContentOptions : Instantiate ModelsDownloadContentOptions
func (*WatsonMachineLearningV4) NewModelsDownloadContentOptions(modelID string, attachmentID string) *ModelsDownloadContentOptions {
	return &ModelsDownloadContentOptions{
		ModelID:      core.StringPtr(modelID),
		AttachmentID: core.StringPtr(attachmentID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsDownloadContentOptions) SetModelID(modelID string) *ModelsDownloadContentOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetAttachmentID : Allow user to set AttachmentID
func (_options *ModelsDownloadContentOptions) SetAttachmentID(attachmentID string) *ModelsDownloadContentOptions {
	_options.AttachmentID = core.StringPtr(attachmentID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *ModelsDownloadContentOptions) SetAccept(accept string) *ModelsDownloadContentOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsDownloadContentOptions) SetSpaceID(spaceID string) *ModelsDownloadContentOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsDownloadContentOptions) SetProjectID(projectID string) *ModelsDownloadContentOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *ModelsDownloadContentOptions) SetRev(rev string) *ModelsDownloadContentOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsDownloadContentOptions) SetHeaders(param map[string]string) *ModelsDownloadContentOptions {
	options.Headers = param
	return options
}

// ModelsFilteredDownloadOptions : The ModelsFilteredDownload options.
type ModelsFilteredDownloadOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// The type of the response: application/zip, application/gzip, application/json, text/plain, or application/xml. A
	// character encoding can be specified by including a `charset` parameter. For example, 'text/plain;charset=utf-8'.
	Accept *string `json:"Accept,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Returns only resources that match this `pipeline_node_id`, this is only relevant if `content_format` is
	// `pipeline-node`.
	PipelineNodeID *string `json:"pipeline_node_id,omitempty"`

	// Returns only resources that match this `deployment_id`, this is only relevant if `content_format` is `coreml`.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// Match an attachment with this name.
	Name *string `json:"name,omitempty"`

	// This is the format of the content. Any value can be used for the format and is there to be able to easily find
	// content by format.
	ContentFormat *string `json:"content_format,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsFilteredDownloadOptions : Instantiate ModelsFilteredDownloadOptions
func (*WatsonMachineLearningV4) NewModelsFilteredDownloadOptions(modelID string) *ModelsFilteredDownloadOptions {
	return &ModelsFilteredDownloadOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsFilteredDownloadOptions) SetModelID(modelID string) *ModelsFilteredDownloadOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *ModelsFilteredDownloadOptions) SetAccept(accept string) *ModelsFilteredDownloadOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsFilteredDownloadOptions) SetSpaceID(spaceID string) *ModelsFilteredDownloadOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsFilteredDownloadOptions) SetProjectID(projectID string) *ModelsFilteredDownloadOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *ModelsFilteredDownloadOptions) SetRev(rev string) *ModelsFilteredDownloadOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetPipelineNodeID : Allow user to set PipelineNodeID
func (_options *ModelsFilteredDownloadOptions) SetPipelineNodeID(pipelineNodeID string) *ModelsFilteredDownloadOptions {
	_options.PipelineNodeID = core.StringPtr(pipelineNodeID)
	return _options
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *ModelsFilteredDownloadOptions) SetDeploymentID(deploymentID string) *ModelsFilteredDownloadOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetName : Allow user to set Name
func (_options *ModelsFilteredDownloadOptions) SetName(name string) *ModelsFilteredDownloadOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetContentFormat : Allow user to set ContentFormat
func (_options *ModelsFilteredDownloadOptions) SetContentFormat(contentFormat string) *ModelsFilteredDownloadOptions {
	_options.ContentFormat = core.StringPtr(contentFormat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsFilteredDownloadOptions) SetHeaders(param map[string]string) *ModelsFilteredDownloadOptions {
	options.Headers = param
	return options
}

// ModelsGetOptions : The ModelsGet options.
type ModelsGetOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsGetOptions : Instantiate ModelsGetOptions
func (*WatsonMachineLearningV4) NewModelsGetOptions(modelID string) *ModelsGetOptions {
	return &ModelsGetOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsGetOptions) SetModelID(modelID string) *ModelsGetOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsGetOptions) SetSpaceID(spaceID string) *ModelsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsGetOptions) SetProjectID(projectID string) *ModelsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *ModelsGetOptions) SetRev(rev string) *ModelsGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsGetOptions) SetHeaders(param map[string]string) *ModelsGetOptions {
	options.Headers = param
	return options
}

// ModelsListAttachmentsOptions : The ModelsListAttachments options.
type ModelsListAttachmentsOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// List only attachments with the given name.
	Name *string `json:"name,omitempty"`

	// This is the format of the content. Any value can be used for the format and is there to be able to easily find
	// content by format.
	ContentFormat *string `json:"content_format,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsListAttachmentsOptions : Instantiate ModelsListAttachmentsOptions
func (*WatsonMachineLearningV4) NewModelsListAttachmentsOptions(modelID string) *ModelsListAttachmentsOptions {
	return &ModelsListAttachmentsOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsListAttachmentsOptions) SetModelID(modelID string) *ModelsListAttachmentsOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsListAttachmentsOptions) SetSpaceID(spaceID string) *ModelsListAttachmentsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsListAttachmentsOptions) SetProjectID(projectID string) *ModelsListAttachmentsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *ModelsListAttachmentsOptions) SetRev(rev string) *ModelsListAttachmentsOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetName : Allow user to set Name
func (_options *ModelsListAttachmentsOptions) SetName(name string) *ModelsListAttachmentsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetContentFormat : Allow user to set ContentFormat
func (_options *ModelsListAttachmentsOptions) SetContentFormat(contentFormat string) *ModelsListAttachmentsOptions {
	_options.ContentFormat = core.StringPtr(contentFormat)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsListAttachmentsOptions) SetHeaders(param map[string]string) *ModelsListAttachmentsOptions {
	options.Headers = param
	return options
}

// ModelsListOptions : The ModelsList options.
type ModelsListOptions struct {
	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsListOptions : Instantiate ModelsListOptions
func (*WatsonMachineLearningV4) NewModelsListOptions() *ModelsListOptions {
	return &ModelsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsListOptions) SetSpaceID(spaceID string) *ModelsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsListOptions) SetProjectID(projectID string) *ModelsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ModelsListOptions) SetStart(start string) *ModelsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ModelsListOptions) SetLimit(limit int64) *ModelsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *ModelsListOptions) SetTagValue(tagValue string) *ModelsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *ModelsListOptions) SetSearch(search string) *ModelsListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsListOptions) SetHeaders(param map[string]string) *ModelsListOptions {
	options.Headers = param
	return options
}

// ModelsListRevisionsOptions : The ModelsListRevisions options.
type ModelsListRevisionsOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsListRevisionsOptions : Instantiate ModelsListRevisionsOptions
func (*WatsonMachineLearningV4) NewModelsListRevisionsOptions(modelID string) *ModelsListRevisionsOptions {
	return &ModelsListRevisionsOptions{
		ModelID: core.StringPtr(modelID),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsListRevisionsOptions) SetModelID(modelID string) *ModelsListRevisionsOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsListRevisionsOptions) SetSpaceID(spaceID string) *ModelsListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsListRevisionsOptions) SetProjectID(projectID string) *ModelsListRevisionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ModelsListRevisionsOptions) SetStart(start string) *ModelsListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ModelsListRevisionsOptions) SetLimit(limit int64) *ModelsListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsListRevisionsOptions) SetHeaders(param map[string]string) *ModelsListRevisionsOptions {
	options.Headers = param
	return options
}

// ModelsUpdateOptions : The ModelsUpdate options.
type ModelsUpdateOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsUpdateOptions : Instantiate ModelsUpdateOptions
func (*WatsonMachineLearningV4) NewModelsUpdateOptions(modelID string, jsonPatch []JSONPatchOperation) *ModelsUpdateOptions {
	return &ModelsUpdateOptions{
		ModelID:   core.StringPtr(modelID),
		JSONPatch: jsonPatch,
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsUpdateOptions) SetModelID(modelID string) *ModelsUpdateOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *ModelsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *ModelsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsUpdateOptions) SetSpaceID(spaceID string) *ModelsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsUpdateOptions) SetProjectID(projectID string) *ModelsUpdateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsUpdateOptions) SetHeaders(param map[string]string) *ModelsUpdateOptions {
	options.Headers = param
	return options
}

// ModelsUploadContentOptions : The ModelsUploadContent options.
type ModelsUploadContentOptions struct {
	// Model identifier.
	ModelID *string `json:"model_id" validate:"required,ne="`

	// This is the format of the content. Any value can be used for the format and is there to be able to easily find
	// content by format.
	ContentFormat *string `json:"content_format" validate:"required"`

	// models file.
	UploadContent map[string]interface{} `json:"upload-content,omitempty"`

	// models file.
	Body io.ReadCloser `json:"body,omitempty"`

	// The type of the input. A character encoding can be specified by including a `charset` parameter. For example,
	// 'text/plain;charset=utf-8'.
	ContentType *string `json:"Content-Type,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Returns only resources that match this `pipeline_node_id`, this is only relevant if `content_format` is
	// `pipeline-node`.
	PipelineNodeID *string `json:"pipeline_node_id,omitempty"`

	// Returns only resources that match this `deployment_id`, this is only relevant if `content_format` is `coreml`.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// Provide the name of the attachment.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewModelsUploadContentOptions : Instantiate ModelsUploadContentOptions
func (*WatsonMachineLearningV4) NewModelsUploadContentOptions(modelID string, contentFormat string) *ModelsUploadContentOptions {
	return &ModelsUploadContentOptions{
		ModelID:       core.StringPtr(modelID),
		ContentFormat: core.StringPtr(contentFormat),
	}
}

// SetModelID : Allow user to set ModelID
func (_options *ModelsUploadContentOptions) SetModelID(modelID string) *ModelsUploadContentOptions {
	_options.ModelID = core.StringPtr(modelID)
	return _options
}

// SetContentFormat : Allow user to set ContentFormat
func (_options *ModelsUploadContentOptions) SetContentFormat(contentFormat string) *ModelsUploadContentOptions {
	_options.ContentFormat = core.StringPtr(contentFormat)
	return _options
}

// SetUploadContent : Allow user to set UploadContent
func (_options *ModelsUploadContentOptions) SetUploadContent(uploadContent map[string]interface{}) *ModelsUploadContentOptions {
	_options.UploadContent = uploadContent
	return _options
}

// SetBody : Allow user to set Body
func (_options *ModelsUploadContentOptions) SetBody(body io.ReadCloser) *ModelsUploadContentOptions {
	_options.Body = body
	return _options
}

// SetContentType : Allow user to set ContentType
func (_options *ModelsUploadContentOptions) SetContentType(contentType string) *ModelsUploadContentOptions {
	_options.ContentType = core.StringPtr(contentType)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ModelsUploadContentOptions) SetSpaceID(spaceID string) *ModelsUploadContentOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ModelsUploadContentOptions) SetProjectID(projectID string) *ModelsUploadContentOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetPipelineNodeID : Allow user to set PipelineNodeID
func (_options *ModelsUploadContentOptions) SetPipelineNodeID(pipelineNodeID string) *ModelsUploadContentOptions {
	_options.PipelineNodeID = core.StringPtr(pipelineNodeID)
	return _options
}

// SetDeploymentID : Allow user to set DeploymentID
func (_options *ModelsUploadContentOptions) SetDeploymentID(deploymentID string) *ModelsUploadContentOptions {
	_options.DeploymentID = core.StringPtr(deploymentID)
	return _options
}

// SetName : Allow user to set Name
func (_options *ModelsUploadContentOptions) SetName(name string) *ModelsUploadContentOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ModelsUploadContentOptions) SetHeaders(param map[string]string) *ModelsUploadContentOptions {
	options.Headers = param
	return options
}

// PaginationFirst : The reference to the first item in the current page.
type PaginationFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalPaginationFirst unmarshals an instance of PaginationFirst from the specified map of raw messages.
func UnmarshalPaginationFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginationNext : A reference to the first item of the next page, if any.
type PaginationNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalPaginationNext unmarshals an instance of PaginationNext from the specified map of raw messages.
func UnmarshalPaginationNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelineRelDataBindingsItem : PipelineRelDataBindingsItem struct
type PipelineRelDataBindingsItem struct {
	// The input_data_reference name that is the input for the node identified by node_id.
	DataReferenceName *string `json:"data_reference_name" validate:"required"`

	// The id of the pipeline node that will receive the associated input_data_reference input.
	NodeID *string `json:"node_id" validate:"required"`
}

// NewPipelineRelDataBindingsItem : Instantiate PipelineRelDataBindingsItem (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewPipelineRelDataBindingsItem(dataReferenceName string, nodeID string) (_model *PipelineRelDataBindingsItem, err error) {
	_model = &PipelineRelDataBindingsItem{
		DataReferenceName: core.StringPtr(dataReferenceName),
		NodeID:            core.StringPtr(nodeID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPipelineRelDataBindingsItem unmarshals an instance of PipelineRelDataBindingsItem from the specified map of raw messages.
func UnmarshalPipelineRelDataBindingsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineRelDataBindingsItem)
	err = core.UnmarshalPrimitive(m, "data_reference_name", &obj.DataReferenceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "node_id", &obj.NodeID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelineRelNodesParametersItem : PipelineRelNodesParametersItem struct
type PipelineRelNodesParametersItem struct {
	// The id of the pipeline node that will receive the associated input_data_reference input.
	NodeID *string `json:"node_id" validate:"required"`

	// The parameters.
	Parameters interface{} `json:"parameters" validate:"required"`
}

// NewPipelineRelNodesParametersItem : Instantiate PipelineRelNodesParametersItem (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewPipelineRelNodesParametersItem(nodeID string, parameters interface{}) (_model *PipelineRelNodesParametersItem, err error) {
	_model = &PipelineRelNodesParametersItem{
		NodeID:     core.StringPtr(nodeID),
		Parameters: parameters,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPipelineRelNodesParametersItem unmarshals an instance of PipelineRelNodesParametersItem from the specified map of raw messages.
func UnmarshalPipelineRelNodesParametersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineRelNodesParametersItem)
	err = core.UnmarshalPrimitive(m, "node_id", &obj.NodeID)
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

// PipelineResourceEntity : The details of the pipeline to be created.
type PipelineResourceEntity struct {
	// The pipeline document, see
	// [pipeline-flow-v2-schema](https://raw.githubusercontent.com/elyra-ai/pipeline-schemas/master/common-pipeline/pipeline-flow/pipeline-flow-v2-schema.json)
	// for the schema definition.
	Document interface{} `json:"document" validate:"required"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// UnmarshalPipelineResourceEntity unmarshals an instance of PipelineResourceEntity from the specified map of raw messages.
func UnmarshalPipelineResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineResourceEntity)
	err = core.UnmarshalPrimitive(m, "document", &obj.Document)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelineResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type PipelineResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalPipelineResourceMetadata unmarshals an instance of PipelineResourceMetadata from the specified map of raw messages.
func UnmarshalPipelineResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// PipelineResourcesFirst : The reference to the first item in the current page.
type PipelineResourcesFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalPipelineResourcesFirst unmarshals an instance of PipelineResourcesFirst from the specified map of raw messages.
func UnmarshalPipelineResourcesFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineResourcesFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelineResourcesNext : A reference to the first item of the next page, if any.
type PipelineResourcesNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalPipelineResourcesNext unmarshals an instance of PipelineResourcesNext from the specified map of raw messages.
func UnmarshalPipelineResourcesNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineResourcesNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelinesCreateOptions : The PipelinesCreate options.
type PipelinesCreateOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// The pipeline document, see
	// [pipeline-flow-v2-schema](https://raw.githubusercontent.com/elyra-ai/pipeline-schemas/master/common-pipeline/pipeline-flow/pipeline-flow-v2-schema.json)
	// for the schema definition.
	Document interface{} `json:"document" validate:"required"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPipelinesCreateOptions : Instantiate PipelinesCreateOptions
func (*WatsonMachineLearningV4) NewPipelinesCreateOptions(name string, document interface{}) *PipelinesCreateOptions {
	return &PipelinesCreateOptions{
		Name:     core.StringPtr(name),
		Document: document,
	}
}

// SetName : Allow user to set Name
func (_options *PipelinesCreateOptions) SetName(name string) *PipelinesCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDocument : Allow user to set Document
func (_options *PipelinesCreateOptions) SetDocument(document interface{}) *PipelinesCreateOptions {
	_options.Document = document
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *PipelinesCreateOptions) SetProjectID(projectID string) *PipelinesCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *PipelinesCreateOptions) SetSpaceID(spaceID string) *PipelinesCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *PipelinesCreateOptions) SetDescription(description string) *PipelinesCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *PipelinesCreateOptions) SetTags(tags []string) *PipelinesCreateOptions {
	_options.Tags = tags
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *PipelinesCreateOptions) SetCustom(custom map[string]interface{}) *PipelinesCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PipelinesCreateOptions) SetHeaders(param map[string]string) *PipelinesCreateOptions {
	options.Headers = param
	return options
}

// PipelinesCreateRevisionOptions : The PipelinesCreateRevision options.
type PipelinesCreateRevisionOptions struct {
	// Pipeline identifier.
	PipelineID *string `json:"pipeline_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPipelinesCreateRevisionOptions : Instantiate PipelinesCreateRevisionOptions
func (*WatsonMachineLearningV4) NewPipelinesCreateRevisionOptions(pipelineID string) *PipelinesCreateRevisionOptions {
	return &PipelinesCreateRevisionOptions{
		PipelineID: core.StringPtr(pipelineID),
	}
}

// SetPipelineID : Allow user to set PipelineID
func (_options *PipelinesCreateRevisionOptions) SetPipelineID(pipelineID string) *PipelinesCreateRevisionOptions {
	_options.PipelineID = core.StringPtr(pipelineID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *PipelinesCreateRevisionOptions) SetSpaceID(spaceID string) *PipelinesCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *PipelinesCreateRevisionOptions) SetProjectID(projectID string) *PipelinesCreateRevisionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *PipelinesCreateRevisionOptions) SetCommitMessage(commitMessage string) *PipelinesCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PipelinesCreateRevisionOptions) SetHeaders(param map[string]string) *PipelinesCreateRevisionOptions {
	options.Headers = param
	return options
}

// PipelinesDeleteOptions : The PipelinesDelete options.
type PipelinesDeleteOptions struct {
	// Pipeline identifier.
	PipelineID *string `json:"pipeline_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPipelinesDeleteOptions : Instantiate PipelinesDeleteOptions
func (*WatsonMachineLearningV4) NewPipelinesDeleteOptions(pipelineID string) *PipelinesDeleteOptions {
	return &PipelinesDeleteOptions{
		PipelineID: core.StringPtr(pipelineID),
	}
}

// SetPipelineID : Allow user to set PipelineID
func (_options *PipelinesDeleteOptions) SetPipelineID(pipelineID string) *PipelinesDeleteOptions {
	_options.PipelineID = core.StringPtr(pipelineID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *PipelinesDeleteOptions) SetSpaceID(spaceID string) *PipelinesDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *PipelinesDeleteOptions) SetProjectID(projectID string) *PipelinesDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PipelinesDeleteOptions) SetHeaders(param map[string]string) *PipelinesDeleteOptions {
	options.Headers = param
	return options
}

// PipelinesGetOptions : The PipelinesGet options.
type PipelinesGetOptions struct {
	// Pipeline identifier.
	PipelineID *string `json:"pipeline_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPipelinesGetOptions : Instantiate PipelinesGetOptions
func (*WatsonMachineLearningV4) NewPipelinesGetOptions(pipelineID string) *PipelinesGetOptions {
	return &PipelinesGetOptions{
		PipelineID: core.StringPtr(pipelineID),
	}
}

// SetPipelineID : Allow user to set PipelineID
func (_options *PipelinesGetOptions) SetPipelineID(pipelineID string) *PipelinesGetOptions {
	_options.PipelineID = core.StringPtr(pipelineID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *PipelinesGetOptions) SetSpaceID(spaceID string) *PipelinesGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *PipelinesGetOptions) SetProjectID(projectID string) *PipelinesGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *PipelinesGetOptions) SetRev(rev string) *PipelinesGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PipelinesGetOptions) SetHeaders(param map[string]string) *PipelinesGetOptions {
	options.Headers = param
	return options
}

// PipelinesListOptions : The PipelinesList options.
type PipelinesListOptions struct {
	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPipelinesListOptions : Instantiate PipelinesListOptions
func (*WatsonMachineLearningV4) NewPipelinesListOptions() *PipelinesListOptions {
	return &PipelinesListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *PipelinesListOptions) SetSpaceID(spaceID string) *PipelinesListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *PipelinesListOptions) SetProjectID(projectID string) *PipelinesListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *PipelinesListOptions) SetStart(start string) *PipelinesListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *PipelinesListOptions) SetLimit(limit int64) *PipelinesListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *PipelinesListOptions) SetTagValue(tagValue string) *PipelinesListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *PipelinesListOptions) SetSearch(search string) *PipelinesListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PipelinesListOptions) SetHeaders(param map[string]string) *PipelinesListOptions {
	options.Headers = param
	return options
}

// PipelinesListRevisionsOptions : The PipelinesListRevisions options.
type PipelinesListRevisionsOptions struct {
	// Pipeline identifier.
	PipelineID *string `json:"pipeline_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPipelinesListRevisionsOptions : Instantiate PipelinesListRevisionsOptions
func (*WatsonMachineLearningV4) NewPipelinesListRevisionsOptions(pipelineID string) *PipelinesListRevisionsOptions {
	return &PipelinesListRevisionsOptions{
		PipelineID: core.StringPtr(pipelineID),
	}
}

// SetPipelineID : Allow user to set PipelineID
func (_options *PipelinesListRevisionsOptions) SetPipelineID(pipelineID string) *PipelinesListRevisionsOptions {
	_options.PipelineID = core.StringPtr(pipelineID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *PipelinesListRevisionsOptions) SetSpaceID(spaceID string) *PipelinesListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *PipelinesListRevisionsOptions) SetProjectID(projectID string) *PipelinesListRevisionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *PipelinesListRevisionsOptions) SetStart(start string) *PipelinesListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *PipelinesListRevisionsOptions) SetLimit(limit int64) *PipelinesListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PipelinesListRevisionsOptions) SetHeaders(param map[string]string) *PipelinesListRevisionsOptions {
	options.Headers = param
	return options
}

// PipelinesUpdateOptions : The PipelinesUpdate options.
type PipelinesUpdateOptions struct {
	// Pipeline identifier.
	PipelineID *string `json:"pipeline_id" validate:"required,ne="`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPipelinesUpdateOptions : Instantiate PipelinesUpdateOptions
func (*WatsonMachineLearningV4) NewPipelinesUpdateOptions(pipelineID string, jsonPatch []JSONPatchOperation) *PipelinesUpdateOptions {
	return &PipelinesUpdateOptions{
		PipelineID: core.StringPtr(pipelineID),
		JSONPatch:  jsonPatch,
	}
}

// SetPipelineID : Allow user to set PipelineID
func (_options *PipelinesUpdateOptions) SetPipelineID(pipelineID string) *PipelinesUpdateOptions {
	_options.PipelineID = core.StringPtr(pipelineID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *PipelinesUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *PipelinesUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *PipelinesUpdateOptions) SetSpaceID(spaceID string) *PipelinesUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *PipelinesUpdateOptions) SetProjectID(projectID string) *PipelinesUpdateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PipelinesUpdateOptions) SetHeaders(param map[string]string) *PipelinesUpdateOptions {
	options.Headers = param
	return options
}

// RemoteTrainingSystemResourceEntity : The definition of a remote system used by Federated Learning.
type RemoteTrainingSystemResourceEntity struct {
	// The list of allowed identities that are allowed to access the remote system.
	AllowedIdentities []AllowedIdentity `json:"allowed_identities" validate:"required"`

	// A remote organization.
	Organization *Organization `json:"organization,omitempty"`

	// The details of the remote administrator for the organization and identities.
	RemoteAdmin *RemoteAdmin `json:"remote_admin,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// UnmarshalRemoteTrainingSystemResourceEntity unmarshals an instance of RemoteTrainingSystemResourceEntity from the specified map of raw messages.
func UnmarshalRemoteTrainingSystemResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoteTrainingSystemResourceEntity)
	err = core.UnmarshalModel(m, "allowed_identities", &obj.AllowedIdentities, UnmarshalAllowedIdentity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "organization", &obj.Organization, UnmarshalOrganization)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "remote_admin", &obj.RemoteAdmin, UnmarshalRemoteAdmin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoteTrainingSystemResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type RemoteTrainingSystemResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalRemoteTrainingSystemResourceMetadata unmarshals an instance of RemoteTrainingSystemResourceMetadata from the specified map of raw messages.
func UnmarshalRemoteTrainingSystemResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoteTrainingSystemResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// RemoteTrainingSystemsCreateOptions : The RemoteTrainingSystemsCreate options.
type RemoteTrainingSystemsCreateOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// The list of allowed identities that are allowed to access the remote system.
	AllowedIdentities []AllowedIdentity `json:"allowed_identities" validate:"required"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// A remote organization.
	Organization *Organization `json:"organization,omitempty"`

	// The details of the remote administrator for the organization and identities.
	RemoteAdmin *RemoteAdmin `json:"remote_admin,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoteTrainingSystemsCreateOptions : Instantiate RemoteTrainingSystemsCreateOptions
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemsCreateOptions(name string, allowedIdentities []AllowedIdentity) *RemoteTrainingSystemsCreateOptions {
	return &RemoteTrainingSystemsCreateOptions{
		Name:              core.StringPtr(name),
		AllowedIdentities: allowedIdentities,
	}
}

// SetName : Allow user to set Name
func (_options *RemoteTrainingSystemsCreateOptions) SetName(name string) *RemoteTrainingSystemsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetAllowedIdentities : Allow user to set AllowedIdentities
func (_options *RemoteTrainingSystemsCreateOptions) SetAllowedIdentities(allowedIdentities []AllowedIdentity) *RemoteTrainingSystemsCreateOptions {
	_options.AllowedIdentities = allowedIdentities
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *RemoteTrainingSystemsCreateOptions) SetProjectID(projectID string) *RemoteTrainingSystemsCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *RemoteTrainingSystemsCreateOptions) SetSpaceID(spaceID string) *RemoteTrainingSystemsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *RemoteTrainingSystemsCreateOptions) SetDescription(description string) *RemoteTrainingSystemsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *RemoteTrainingSystemsCreateOptions) SetTags(tags []string) *RemoteTrainingSystemsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetOrganization : Allow user to set Organization
func (_options *RemoteTrainingSystemsCreateOptions) SetOrganization(organization *Organization) *RemoteTrainingSystemsCreateOptions {
	_options.Organization = organization
	return _options
}

// SetRemoteAdmin : Allow user to set RemoteAdmin
func (_options *RemoteTrainingSystemsCreateOptions) SetRemoteAdmin(remoteAdmin *RemoteAdmin) *RemoteTrainingSystemsCreateOptions {
	_options.RemoteAdmin = remoteAdmin
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *RemoteTrainingSystemsCreateOptions) SetCustom(custom map[string]interface{}) *RemoteTrainingSystemsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoteTrainingSystemsCreateOptions) SetHeaders(param map[string]string) *RemoteTrainingSystemsCreateOptions {
	options.Headers = param
	return options
}

// RemoteTrainingSystemsCreateRevisionOptions : The RemoteTrainingSystemsCreateRevision options.
type RemoteTrainingSystemsCreateRevisionOptions struct {
	// Remote training system identifier.
	RemoteTrainingSystemID *string `json:"remote_training_system_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoteTrainingSystemsCreateRevisionOptions : Instantiate RemoteTrainingSystemsCreateRevisionOptions
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemsCreateRevisionOptions(remoteTrainingSystemID string) *RemoteTrainingSystemsCreateRevisionOptions {
	return &RemoteTrainingSystemsCreateRevisionOptions{
		RemoteTrainingSystemID: core.StringPtr(remoteTrainingSystemID),
	}
}

// SetRemoteTrainingSystemID : Allow user to set RemoteTrainingSystemID
func (_options *RemoteTrainingSystemsCreateRevisionOptions) SetRemoteTrainingSystemID(remoteTrainingSystemID string) *RemoteTrainingSystemsCreateRevisionOptions {
	_options.RemoteTrainingSystemID = core.StringPtr(remoteTrainingSystemID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *RemoteTrainingSystemsCreateRevisionOptions) SetSpaceID(spaceID string) *RemoteTrainingSystemsCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *RemoteTrainingSystemsCreateRevisionOptions) SetProjectID(projectID string) *RemoteTrainingSystemsCreateRevisionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *RemoteTrainingSystemsCreateRevisionOptions) SetCommitMessage(commitMessage string) *RemoteTrainingSystemsCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoteTrainingSystemsCreateRevisionOptions) SetHeaders(param map[string]string) *RemoteTrainingSystemsCreateRevisionOptions {
	options.Headers = param
	return options
}

// RemoteTrainingSystemsDeleteOptions : The RemoteTrainingSystemsDelete options.
type RemoteTrainingSystemsDeleteOptions struct {
	// Remote training system identifier.
	RemoteTrainingSystemID *string `json:"remote_training_system_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoteTrainingSystemsDeleteOptions : Instantiate RemoteTrainingSystemsDeleteOptions
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemsDeleteOptions(remoteTrainingSystemID string) *RemoteTrainingSystemsDeleteOptions {
	return &RemoteTrainingSystemsDeleteOptions{
		RemoteTrainingSystemID: core.StringPtr(remoteTrainingSystemID),
	}
}

// SetRemoteTrainingSystemID : Allow user to set RemoteTrainingSystemID
func (_options *RemoteTrainingSystemsDeleteOptions) SetRemoteTrainingSystemID(remoteTrainingSystemID string) *RemoteTrainingSystemsDeleteOptions {
	_options.RemoteTrainingSystemID = core.StringPtr(remoteTrainingSystemID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *RemoteTrainingSystemsDeleteOptions) SetSpaceID(spaceID string) *RemoteTrainingSystemsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *RemoteTrainingSystemsDeleteOptions) SetProjectID(projectID string) *RemoteTrainingSystemsDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoteTrainingSystemsDeleteOptions) SetHeaders(param map[string]string) *RemoteTrainingSystemsDeleteOptions {
	options.Headers = param
	return options
}

// RemoteTrainingSystemsGetOptions : The RemoteTrainingSystemsGet options.
type RemoteTrainingSystemsGetOptions struct {
	// Remote training system identifier.
	RemoteTrainingSystemID *string `json:"remote_training_system_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoteTrainingSystemsGetOptions : Instantiate RemoteTrainingSystemsGetOptions
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemsGetOptions(remoteTrainingSystemID string) *RemoteTrainingSystemsGetOptions {
	return &RemoteTrainingSystemsGetOptions{
		RemoteTrainingSystemID: core.StringPtr(remoteTrainingSystemID),
	}
}

// SetRemoteTrainingSystemID : Allow user to set RemoteTrainingSystemID
func (_options *RemoteTrainingSystemsGetOptions) SetRemoteTrainingSystemID(remoteTrainingSystemID string) *RemoteTrainingSystemsGetOptions {
	_options.RemoteTrainingSystemID = core.StringPtr(remoteTrainingSystemID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *RemoteTrainingSystemsGetOptions) SetSpaceID(spaceID string) *RemoteTrainingSystemsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *RemoteTrainingSystemsGetOptions) SetProjectID(projectID string) *RemoteTrainingSystemsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *RemoteTrainingSystemsGetOptions) SetRev(rev string) *RemoteTrainingSystemsGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoteTrainingSystemsGetOptions) SetHeaders(param map[string]string) *RemoteTrainingSystemsGetOptions {
	options.Headers = param
	return options
}

// RemoteTrainingSystemsListOptions : The RemoteTrainingSystemsList options.
type RemoteTrainingSystemsListOptions struct {
	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoteTrainingSystemsListOptions : Instantiate RemoteTrainingSystemsListOptions
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemsListOptions() *RemoteTrainingSystemsListOptions {
	return &RemoteTrainingSystemsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *RemoteTrainingSystemsListOptions) SetSpaceID(spaceID string) *RemoteTrainingSystemsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *RemoteTrainingSystemsListOptions) SetProjectID(projectID string) *RemoteTrainingSystemsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *RemoteTrainingSystemsListOptions) SetStart(start string) *RemoteTrainingSystemsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *RemoteTrainingSystemsListOptions) SetLimit(limit int64) *RemoteTrainingSystemsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *RemoteTrainingSystemsListOptions) SetTagValue(tagValue string) *RemoteTrainingSystemsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *RemoteTrainingSystemsListOptions) SetSearch(search string) *RemoteTrainingSystemsListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoteTrainingSystemsListOptions) SetHeaders(param map[string]string) *RemoteTrainingSystemsListOptions {
	options.Headers = param
	return options
}

// RemoteTrainingSystemsListRevisionsOptions : The RemoteTrainingSystemsListRevisions options.
type RemoteTrainingSystemsListRevisionsOptions struct {
	// Remote training system identifier.
	RemoteTrainingSystemID *string `json:"remote_training_system_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoteTrainingSystemsListRevisionsOptions : Instantiate RemoteTrainingSystemsListRevisionsOptions
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemsListRevisionsOptions(remoteTrainingSystemID string) *RemoteTrainingSystemsListRevisionsOptions {
	return &RemoteTrainingSystemsListRevisionsOptions{
		RemoteTrainingSystemID: core.StringPtr(remoteTrainingSystemID),
	}
}

// SetRemoteTrainingSystemID : Allow user to set RemoteTrainingSystemID
func (_options *RemoteTrainingSystemsListRevisionsOptions) SetRemoteTrainingSystemID(remoteTrainingSystemID string) *RemoteTrainingSystemsListRevisionsOptions {
	_options.RemoteTrainingSystemID = core.StringPtr(remoteTrainingSystemID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *RemoteTrainingSystemsListRevisionsOptions) SetSpaceID(spaceID string) *RemoteTrainingSystemsListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *RemoteTrainingSystemsListRevisionsOptions) SetProjectID(projectID string) *RemoteTrainingSystemsListRevisionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *RemoteTrainingSystemsListRevisionsOptions) SetStart(start string) *RemoteTrainingSystemsListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *RemoteTrainingSystemsListRevisionsOptions) SetLimit(limit int64) *RemoteTrainingSystemsListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoteTrainingSystemsListRevisionsOptions) SetHeaders(param map[string]string) *RemoteTrainingSystemsListRevisionsOptions {
	options.Headers = param
	return options
}

// RemoteTrainingSystemsUpdateOptions : The RemoteTrainingSystemsUpdate options.
type RemoteTrainingSystemsUpdateOptions struct {
	// Remote training system identifier.
	RemoteTrainingSystemID *string `json:"remote_training_system_id" validate:"required,ne="`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoteTrainingSystemsUpdateOptions : Instantiate RemoteTrainingSystemsUpdateOptions
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemsUpdateOptions(remoteTrainingSystemID string, jsonPatch []JSONPatchOperation) *RemoteTrainingSystemsUpdateOptions {
	return &RemoteTrainingSystemsUpdateOptions{
		RemoteTrainingSystemID: core.StringPtr(remoteTrainingSystemID),
		JSONPatch:              jsonPatch,
	}
}

// SetRemoteTrainingSystemID : Allow user to set RemoteTrainingSystemID
func (_options *RemoteTrainingSystemsUpdateOptions) SetRemoteTrainingSystemID(remoteTrainingSystemID string) *RemoteTrainingSystemsUpdateOptions {
	_options.RemoteTrainingSystemID = core.StringPtr(remoteTrainingSystemID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *RemoteTrainingSystemsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *RemoteTrainingSystemsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *RemoteTrainingSystemsUpdateOptions) SetSpaceID(spaceID string) *RemoteTrainingSystemsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *RemoteTrainingSystemsUpdateOptions) SetProjectID(projectID string) *RemoteTrainingSystemsUpdateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoteTrainingSystemsUpdateOptions) SetHeaders(param map[string]string) *RemoteTrainingSystemsUpdateOptions {
	options.Headers = param
	return options
}

// ResourceMetaBaseCommitInfo : Information related to the revision.
type ResourceMetaBaseCommitInfo struct {
	// The time when the revision was committed.
	CommittedAt *strfmt.DateTime `json:"committed_at" validate:"required"`

	// The message that was provided when the revision was created.
	CommitMessage *string `json:"commit_message,omitempty"`
}

// UnmarshalResourceMetaBaseCommitInfo unmarshals an instance of ResourceMetaBaseCommitInfo from the specified map of raw messages.
func UnmarshalResourceMetaBaseCommitInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceMetaBaseCommitInfo)
	err = core.UnmarshalPrimitive(m, "committed_at", &obj.CommittedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "commit_message", &obj.CommitMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDefinitionResourceEntity : The `training_data_references` contain the training datasets and the
// `results_reference` the connection where results will be stored.
type TrainingDefinitionResourceEntity struct {
	// A reference to a resource.
	Experiment *Rel `json:"experiment,omitempty"`

	// A pipeline.
	// The `hardware_spec` is a reference to the hardware specification.
	// The `hybrid_pipeline_hardware_specs` are used only when training a hybrid pipeline in order to
	// specify compute requirement for each pipeline node.
	Pipeline *PipelineRel `json:"pipeline,omitempty"`

	// A model.
	// The `software_spec` is a reference to a software specification.
	// The `hardware_spec` is a reference to a hardware specification.
	ModelDefinition *ModelDefinitionRel `json:"model_definition,omitempty"`

	// Federated Learning is a Technical Preview.
	FederatedLearning *FederatedLearning `json:"federated_learning,omitempty"`

	// Training datasets.
	TrainingDataReferences []DataConnectionReference `json:"training_data_references,omitempty"`

	// The training results.
	ResultsReference *ObjectLocation `json:"results_reference" validate:"required"`

	// The holdout/test datasets.
	TestDataReferences []DataConnectionReference `json:"test_data_references,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// UnmarshalTrainingDefinitionResourceEntity unmarshals an instance of TrainingDefinitionResourceEntity from the specified map of raw messages.
func UnmarshalTrainingDefinitionResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDefinitionResourceEntity)
	err = core.UnmarshalModel(m, "experiment", &obj.Experiment, UnmarshalRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pipeline", &obj.Pipeline, UnmarshalPipelineRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "model_definition", &obj.ModelDefinition, UnmarshalModelDefinitionRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "federated_learning", &obj.FederatedLearning, UnmarshalFederatedLearning)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data_references", &obj.TrainingDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results_reference", &obj.ResultsReference, UnmarshalObjectLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "test_data_references", &obj.TestDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDefinitionResourceMetadata : Common metadata for a resource where `project_id` or `space_id` must be present.
type TrainingDefinitionResourceMetadata struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalTrainingDefinitionResourceMetadata unmarshals an instance of TrainingDefinitionResourceMetadata from the specified map of raw messages.
func UnmarshalTrainingDefinitionResourceMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDefinitionResourceMetadata)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// TrainingDefinitionResourcesFirst : The reference to the first item in the current page.
type TrainingDefinitionResourcesFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalTrainingDefinitionResourcesFirst unmarshals an instance of TrainingDefinitionResourcesFirst from the specified map of raw messages.
func UnmarshalTrainingDefinitionResourcesFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDefinitionResourcesFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDefinitionResourcesNext : A reference to the first item of the next page, if any.
type TrainingDefinitionResourcesNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalTrainingDefinitionResourcesNext unmarshals an instance of TrainingDefinitionResourcesNext from the specified map of raw messages.
func UnmarshalTrainingDefinitionResourcesNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDefinitionResourcesNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDefinitionsCreateOptions : The TrainingDefinitionsCreate options.
type TrainingDefinitionsCreateOptions struct {
	// The name of the resource.
	Name *string `json:"name" validate:"required"`

	// The training results.
	ResultsReference *ObjectLocation `json:"results_reference" validate:"required"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// A reference to a resource.
	Experiment *Rel `json:"experiment,omitempty"`

	// A pipeline.
	// The `hardware_spec` is a reference to the hardware specification.
	// The `hybrid_pipeline_hardware_specs` are used only when training a hybrid pipeline in order to
	// specify compute requirement for each pipeline node.
	Pipeline *PipelineRel `json:"pipeline,omitempty"`

	// A model.
	// The `software_spec` is a reference to a software specification.
	// The `hardware_spec` is a reference to a hardware specification.
	ModelDefinition *ModelDefinitionRel `json:"model_definition,omitempty"`

	// Federated Learning is a Technical Preview.
	FederatedLearning *FederatedLearning `json:"federated_learning,omitempty"`

	// Training datasets.
	TrainingDataReferences []DataConnectionReference `json:"training_data_references,omitempty"`

	// The holdout/test datasets.
	TestDataReferences []DataConnectionReference `json:"test_data_references,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingDefinitionsCreateOptions : Instantiate TrainingDefinitionsCreateOptions
func (*WatsonMachineLearningV4) NewTrainingDefinitionsCreateOptions(name string, resultsReference *ObjectLocation) *TrainingDefinitionsCreateOptions {
	return &TrainingDefinitionsCreateOptions{
		Name:             core.StringPtr(name),
		ResultsReference: resultsReference,
	}
}

// SetName : Allow user to set Name
func (_options *TrainingDefinitionsCreateOptions) SetName(name string) *TrainingDefinitionsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetResultsReference : Allow user to set ResultsReference
func (_options *TrainingDefinitionsCreateOptions) SetResultsReference(resultsReference *ObjectLocation) *TrainingDefinitionsCreateOptions {
	_options.ResultsReference = resultsReference
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingDefinitionsCreateOptions) SetProjectID(projectID string) *TrainingDefinitionsCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingDefinitionsCreateOptions) SetSpaceID(spaceID string) *TrainingDefinitionsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *TrainingDefinitionsCreateOptions) SetDescription(description string) *TrainingDefinitionsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *TrainingDefinitionsCreateOptions) SetTags(tags []string) *TrainingDefinitionsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetExperiment : Allow user to set Experiment
func (_options *TrainingDefinitionsCreateOptions) SetExperiment(experiment *Rel) *TrainingDefinitionsCreateOptions {
	_options.Experiment = experiment
	return _options
}

// SetPipeline : Allow user to set Pipeline
func (_options *TrainingDefinitionsCreateOptions) SetPipeline(pipeline *PipelineRel) *TrainingDefinitionsCreateOptions {
	_options.Pipeline = pipeline
	return _options
}

// SetModelDefinition : Allow user to set ModelDefinition
func (_options *TrainingDefinitionsCreateOptions) SetModelDefinition(modelDefinition *ModelDefinitionRel) *TrainingDefinitionsCreateOptions {
	_options.ModelDefinition = modelDefinition
	return _options
}

// SetFederatedLearning : Allow user to set FederatedLearning
func (_options *TrainingDefinitionsCreateOptions) SetFederatedLearning(federatedLearning *FederatedLearning) *TrainingDefinitionsCreateOptions {
	_options.FederatedLearning = federatedLearning
	return _options
}

// SetTrainingDataReferences : Allow user to set TrainingDataReferences
func (_options *TrainingDefinitionsCreateOptions) SetTrainingDataReferences(trainingDataReferences []DataConnectionReference) *TrainingDefinitionsCreateOptions {
	_options.TrainingDataReferences = trainingDataReferences
	return _options
}

// SetTestDataReferences : Allow user to set TestDataReferences
func (_options *TrainingDefinitionsCreateOptions) SetTestDataReferences(testDataReferences []DataConnectionReference) *TrainingDefinitionsCreateOptions {
	_options.TestDataReferences = testDataReferences
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *TrainingDefinitionsCreateOptions) SetCustom(custom map[string]interface{}) *TrainingDefinitionsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingDefinitionsCreateOptions) SetHeaders(param map[string]string) *TrainingDefinitionsCreateOptions {
	options.Headers = param
	return options
}

// TrainingDefinitionsCreateRevisionOptions : The TrainingDefinitionsCreateRevision options.
type TrainingDefinitionsCreateRevisionOptions struct {
	// Training definition identifier.
	TrainingDefinitionID *string `json:"training_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// An optional commit message for the revision.
	CommitMessage *string `json:"commit_message,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingDefinitionsCreateRevisionOptions : Instantiate TrainingDefinitionsCreateRevisionOptions
func (*WatsonMachineLearningV4) NewTrainingDefinitionsCreateRevisionOptions(trainingDefinitionID string) *TrainingDefinitionsCreateRevisionOptions {
	return &TrainingDefinitionsCreateRevisionOptions{
		TrainingDefinitionID: core.StringPtr(trainingDefinitionID),
	}
}

// SetTrainingDefinitionID : Allow user to set TrainingDefinitionID
func (_options *TrainingDefinitionsCreateRevisionOptions) SetTrainingDefinitionID(trainingDefinitionID string) *TrainingDefinitionsCreateRevisionOptions {
	_options.TrainingDefinitionID = core.StringPtr(trainingDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingDefinitionsCreateRevisionOptions) SetSpaceID(spaceID string) *TrainingDefinitionsCreateRevisionOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingDefinitionsCreateRevisionOptions) SetProjectID(projectID string) *TrainingDefinitionsCreateRevisionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCommitMessage : Allow user to set CommitMessage
func (_options *TrainingDefinitionsCreateRevisionOptions) SetCommitMessage(commitMessage string) *TrainingDefinitionsCreateRevisionOptions {
	_options.CommitMessage = core.StringPtr(commitMessage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingDefinitionsCreateRevisionOptions) SetHeaders(param map[string]string) *TrainingDefinitionsCreateRevisionOptions {
	options.Headers = param
	return options
}

// TrainingDefinitionsDeleteOptions : The TrainingDefinitionsDelete options.
type TrainingDefinitionsDeleteOptions struct {
	// Training definition identifier.
	TrainingDefinitionID *string `json:"training_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingDefinitionsDeleteOptions : Instantiate TrainingDefinitionsDeleteOptions
func (*WatsonMachineLearningV4) NewTrainingDefinitionsDeleteOptions(trainingDefinitionID string) *TrainingDefinitionsDeleteOptions {
	return &TrainingDefinitionsDeleteOptions{
		TrainingDefinitionID: core.StringPtr(trainingDefinitionID),
	}
}

// SetTrainingDefinitionID : Allow user to set TrainingDefinitionID
func (_options *TrainingDefinitionsDeleteOptions) SetTrainingDefinitionID(trainingDefinitionID string) *TrainingDefinitionsDeleteOptions {
	_options.TrainingDefinitionID = core.StringPtr(trainingDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingDefinitionsDeleteOptions) SetSpaceID(spaceID string) *TrainingDefinitionsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingDefinitionsDeleteOptions) SetProjectID(projectID string) *TrainingDefinitionsDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingDefinitionsDeleteOptions) SetHeaders(param map[string]string) *TrainingDefinitionsDeleteOptions {
	options.Headers = param
	return options
}

// TrainingDefinitionsGetOptions : The TrainingDefinitionsGet options.
type TrainingDefinitionsGetOptions struct {
	// Training definition identifier.
	TrainingDefinitionID *string `json:"training_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// The revision number of the resource.
	Rev *string `json:"rev,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingDefinitionsGetOptions : Instantiate TrainingDefinitionsGetOptions
func (*WatsonMachineLearningV4) NewTrainingDefinitionsGetOptions(trainingDefinitionID string) *TrainingDefinitionsGetOptions {
	return &TrainingDefinitionsGetOptions{
		TrainingDefinitionID: core.StringPtr(trainingDefinitionID),
	}
}

// SetTrainingDefinitionID : Allow user to set TrainingDefinitionID
func (_options *TrainingDefinitionsGetOptions) SetTrainingDefinitionID(trainingDefinitionID string) *TrainingDefinitionsGetOptions {
	_options.TrainingDefinitionID = core.StringPtr(trainingDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingDefinitionsGetOptions) SetSpaceID(spaceID string) *TrainingDefinitionsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingDefinitionsGetOptions) SetProjectID(projectID string) *TrainingDefinitionsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetRev : Allow user to set Rev
func (_options *TrainingDefinitionsGetOptions) SetRev(rev string) *TrainingDefinitionsGetOptions {
	_options.Rev = core.StringPtr(rev)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingDefinitionsGetOptions) SetHeaders(param map[string]string) *TrainingDefinitionsGetOptions {
	options.Headers = param
	return options
}

// TrainingDefinitionsListOptions : The TrainingDefinitionsList options.
type TrainingDefinitionsListOptions struct {
	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Return only the resources with the given tag values, separated by `or` or `and` to support multiple tags.
	TagValue *string `json:"tag.value,omitempty"`

	// Returns only resources that match this search string. The path to the field must be the complete path to the field,
	// and this field must be one of the indexed fields for this resource type. Note that the search string must be URL
	// encoded.
	Search *string `json:"search,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingDefinitionsListOptions : Instantiate TrainingDefinitionsListOptions
func (*WatsonMachineLearningV4) NewTrainingDefinitionsListOptions() *TrainingDefinitionsListOptions {
	return &TrainingDefinitionsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingDefinitionsListOptions) SetSpaceID(spaceID string) *TrainingDefinitionsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingDefinitionsListOptions) SetProjectID(projectID string) *TrainingDefinitionsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *TrainingDefinitionsListOptions) SetStart(start string) *TrainingDefinitionsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *TrainingDefinitionsListOptions) SetLimit(limit int64) *TrainingDefinitionsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *TrainingDefinitionsListOptions) SetTagValue(tagValue string) *TrainingDefinitionsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *TrainingDefinitionsListOptions) SetSearch(search string) *TrainingDefinitionsListOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingDefinitionsListOptions) SetHeaders(param map[string]string) *TrainingDefinitionsListOptions {
	options.Headers = param
	return options
}

// TrainingDefinitionsListRevisionsOptions : The TrainingDefinitionsListRevisions options.
type TrainingDefinitionsListRevisionsOptions struct {
	// Training definition identifier.
	TrainingDefinitionID *string `json:"training_definition_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. By default limit is 100. Max limit allowed is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingDefinitionsListRevisionsOptions : Instantiate TrainingDefinitionsListRevisionsOptions
func (*WatsonMachineLearningV4) NewTrainingDefinitionsListRevisionsOptions(trainingDefinitionID string) *TrainingDefinitionsListRevisionsOptions {
	return &TrainingDefinitionsListRevisionsOptions{
		TrainingDefinitionID: core.StringPtr(trainingDefinitionID),
	}
}

// SetTrainingDefinitionID : Allow user to set TrainingDefinitionID
func (_options *TrainingDefinitionsListRevisionsOptions) SetTrainingDefinitionID(trainingDefinitionID string) *TrainingDefinitionsListRevisionsOptions {
	_options.TrainingDefinitionID = core.StringPtr(trainingDefinitionID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingDefinitionsListRevisionsOptions) SetSpaceID(spaceID string) *TrainingDefinitionsListRevisionsOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingDefinitionsListRevisionsOptions) SetProjectID(projectID string) *TrainingDefinitionsListRevisionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *TrainingDefinitionsListRevisionsOptions) SetStart(start string) *TrainingDefinitionsListRevisionsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *TrainingDefinitionsListRevisionsOptions) SetLimit(limit int64) *TrainingDefinitionsListRevisionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingDefinitionsListRevisionsOptions) SetHeaders(param map[string]string) *TrainingDefinitionsListRevisionsOptions {
	options.Headers = param
	return options
}

// TrainingDefinitionsUpdateOptions : The TrainingDefinitionsUpdate options.
type TrainingDefinitionsUpdateOptions struct {
	// Training definition identifier.
	TrainingDefinitionID *string `json:"training_definition_id" validate:"required,ne="`

	// Input For Patch. This is the patch body which corresponds to the JavaScript Object Notation (JSON) Patch standard
	// (RFC 6902).
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingDefinitionsUpdateOptions : Instantiate TrainingDefinitionsUpdateOptions
func (*WatsonMachineLearningV4) NewTrainingDefinitionsUpdateOptions(trainingDefinitionID string, jsonPatch []JSONPatchOperation) *TrainingDefinitionsUpdateOptions {
	return &TrainingDefinitionsUpdateOptions{
		TrainingDefinitionID: core.StringPtr(trainingDefinitionID),
		JSONPatch:            jsonPatch,
	}
}

// SetTrainingDefinitionID : Allow user to set TrainingDefinitionID
func (_options *TrainingDefinitionsUpdateOptions) SetTrainingDefinitionID(trainingDefinitionID string) *TrainingDefinitionsUpdateOptions {
	_options.TrainingDefinitionID = core.StringPtr(trainingDefinitionID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *TrainingDefinitionsUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *TrainingDefinitionsUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingDefinitionsUpdateOptions) SetSpaceID(spaceID string) *TrainingDefinitionsUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingDefinitionsUpdateOptions) SetProjectID(projectID string) *TrainingDefinitionsUpdateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingDefinitionsUpdateOptions) SetHeaders(param map[string]string) *TrainingDefinitionsUpdateOptions {
	options.Headers = param
	return options
}

// TrainingReferenceHyperParametersOptimization : The hyper parameters used in the experiment.
type TrainingReferenceHyperParametersOptimization struct {
	// Optimization algorithm.
	Method *TrainingReferenceHyperParametersOptimizationMethod `json:"method,omitempty"`

	// Hyper parameters that can be a range or an array of number or strings.
	HyperParameters []HyperParameter `json:"hyper_parameters,omitempty"`
}

// UnmarshalTrainingReferenceHyperParametersOptimization unmarshals an instance of TrainingReferenceHyperParametersOptimization from the specified map of raw messages.
func UnmarshalTrainingReferenceHyperParametersOptimization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingReferenceHyperParametersOptimization)
	err = core.UnmarshalModel(m, "method", &obj.Method, UnmarshalTrainingReferenceHyperParametersOptimizationMethod)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hyper_parameters", &obj.HyperParameters, UnmarshalHyperParameter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingReferenceHyperParametersOptimizationMethod : Optimization algorithm.
type TrainingReferenceHyperParametersOptimizationMethod struct {
	Name *string `json:"name,omitempty"`

	// Optimizer configuration parameters.
	Parameters interface{} `json:"parameters,omitempty"`
}

// Constants associated with the TrainingReferenceHyperParametersOptimizationMethod.Name property.
const (
	TrainingReferenceHyperParametersOptimizationMethod_Name_Bandit    = "bandit"
	TrainingReferenceHyperParametersOptimizationMethod_Name_Hyperband = "hyperband"
	TrainingReferenceHyperParametersOptimizationMethod_Name_Random    = "random"
	TrainingReferenceHyperParametersOptimizationMethod_Name_Rbfopt    = "rbfopt"
)

// UnmarshalTrainingReferenceHyperParametersOptimizationMethod unmarshals an instance of TrainingReferenceHyperParametersOptimizationMethod from the specified map of raw messages.
func UnmarshalTrainingReferenceHyperParametersOptimizationMethod(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingReferenceHyperParametersOptimizationMethod)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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

// TrainingResourcesFirst : The reference to the first item in the current page.
type TrainingResourcesFirst struct {
	// The uri of the first resource returned.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalTrainingResourcesFirst unmarshals an instance of TrainingResourcesFirst from the specified map of raw messages.
func UnmarshalTrainingResourcesFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingResourcesFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingResourcesNext : A reference to the first item of the next page, if any.
type TrainingResourcesNext struct {
	// The uri of the next set of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalTrainingResourcesNext unmarshals an instance of TrainingResourcesNext from the specified map of raw messages.
func UnmarshalTrainingResourcesNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingResourcesNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingStatusHpo : Hyperparameter optimization.
type TrainingStatusHpo struct {
	NodeID *string `json:"node_id,omitempty"`

	HyperParameters interface{} `json:"hyper_parameters" validate:"required"`
}

// UnmarshalTrainingStatusHpo unmarshals an instance of TrainingStatusHpo from the specified map of raw messages.
func UnmarshalTrainingStatusHpo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingStatusHpo)
	err = core.UnmarshalPrimitive(m, "node_id", &obj.NodeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hyper_parameters", &obj.HyperParameters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingStatusMessage : Message.
type TrainingStatusMessage struct {
	Level *string `json:"level,omitempty"`

	Text *string `json:"text,omitempty"`
}

// UnmarshalTrainingStatusMessage unmarshals an instance of TrainingStatusMessage from the specified map of raw messages.
func UnmarshalTrainingStatusMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingStatusMessage)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingsCreateOptions : The TrainingsCreate options.
type TrainingsCreateOptions struct {
	// The training results.
	ResultsReference *ObjectLocation `json:"results_reference" validate:"required"`

	// A reference to a resource.
	Experiment *Rel `json:"experiment,omitempty"`

	// A pipeline.
	// The `hardware_spec` is a reference to the hardware specification.
	// The `hybrid_pipeline_hardware_specs` are used only when training a hybrid pipeline in order to
	// specify compute requirement for each pipeline node.
	Pipeline *PipelineRel `json:"pipeline,omitempty"`

	// A model.
	// The `software_spec` is a reference to a software specification.
	// The `hardware_spec` is a reference to a hardware specification.
	ModelDefinition *ModelDefinitionRel `json:"model_definition,omitempty"`

	// Federated Learning is a Technical Preview.
	FederatedLearning *FederatedLearning `json:"federated_learning,omitempty"`

	// Training datasets.
	TrainingDataReferences []DataConnectionReference `json:"training_data_references,omitempty"`

	// The holdout/test datasets.
	TestDataReferences []DataConnectionReference `json:"test_data_references,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The name of the training.
	Name *string `json:"name,omitempty"`

	// A description of the training.
	Description *string `json:"description,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingsCreateOptions : Instantiate TrainingsCreateOptions
func (*WatsonMachineLearningV4) NewTrainingsCreateOptions(resultsReference *ObjectLocation) *TrainingsCreateOptions {
	return &TrainingsCreateOptions{
		ResultsReference: resultsReference,
	}
}

// SetResultsReference : Allow user to set ResultsReference
func (_options *TrainingsCreateOptions) SetResultsReference(resultsReference *ObjectLocation) *TrainingsCreateOptions {
	_options.ResultsReference = resultsReference
	return _options
}

// SetExperiment : Allow user to set Experiment
func (_options *TrainingsCreateOptions) SetExperiment(experiment *Rel) *TrainingsCreateOptions {
	_options.Experiment = experiment
	return _options
}

// SetPipeline : Allow user to set Pipeline
func (_options *TrainingsCreateOptions) SetPipeline(pipeline *PipelineRel) *TrainingsCreateOptions {
	_options.Pipeline = pipeline
	return _options
}

// SetModelDefinition : Allow user to set ModelDefinition
func (_options *TrainingsCreateOptions) SetModelDefinition(modelDefinition *ModelDefinitionRel) *TrainingsCreateOptions {
	_options.ModelDefinition = modelDefinition
	return _options
}

// SetFederatedLearning : Allow user to set FederatedLearning
func (_options *TrainingsCreateOptions) SetFederatedLearning(federatedLearning *FederatedLearning) *TrainingsCreateOptions {
	_options.FederatedLearning = federatedLearning
	return _options
}

// SetTrainingDataReferences : Allow user to set TrainingDataReferences
func (_options *TrainingsCreateOptions) SetTrainingDataReferences(trainingDataReferences []DataConnectionReference) *TrainingsCreateOptions {
	_options.TrainingDataReferences = trainingDataReferences
	return _options
}

// SetTestDataReferences : Allow user to set TestDataReferences
func (_options *TrainingsCreateOptions) SetTestDataReferences(testDataReferences []DataConnectionReference) *TrainingsCreateOptions {
	_options.TestDataReferences = testDataReferences
	return _options
}

// SetCustom : Allow user to set Custom
func (_options *TrainingsCreateOptions) SetCustom(custom map[string]interface{}) *TrainingsCreateOptions {
	_options.Custom = custom
	return _options
}

// SetTags : Allow user to set Tags
func (_options *TrainingsCreateOptions) SetTags(tags []string) *TrainingsCreateOptions {
	_options.Tags = tags
	return _options
}

// SetName : Allow user to set Name
func (_options *TrainingsCreateOptions) SetName(name string) *TrainingsCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *TrainingsCreateOptions) SetDescription(description string) *TrainingsCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingsCreateOptions) SetSpaceID(spaceID string) *TrainingsCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingsCreateOptions) SetProjectID(projectID string) *TrainingsCreateOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingsCreateOptions) SetHeaders(param map[string]string) *TrainingsCreateOptions {
	options.Headers = param
	return options
}

// TrainingsDeleteOptions : The TrainingsDelete options.
type TrainingsDeleteOptions struct {
	// The training identifier.
	TrainingID *string `json:"training_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Set to true in order to also delete the job metadata information.
	HardDelete *bool `json:"hard_delete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingsDeleteOptions : Instantiate TrainingsDeleteOptions
func (*WatsonMachineLearningV4) NewTrainingsDeleteOptions(trainingID string) *TrainingsDeleteOptions {
	return &TrainingsDeleteOptions{
		TrainingID: core.StringPtr(trainingID),
	}
}

// SetTrainingID : Allow user to set TrainingID
func (_options *TrainingsDeleteOptions) SetTrainingID(trainingID string) *TrainingsDeleteOptions {
	_options.TrainingID = core.StringPtr(trainingID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingsDeleteOptions) SetSpaceID(spaceID string) *TrainingsDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingsDeleteOptions) SetProjectID(projectID string) *TrainingsDeleteOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHardDelete : Allow user to set HardDelete
func (_options *TrainingsDeleteOptions) SetHardDelete(hardDelete bool) *TrainingsDeleteOptions {
	_options.HardDelete = core.BoolPtr(hardDelete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingsDeleteOptions) SetHeaders(param map[string]string) *TrainingsDeleteOptions {
	options.Headers = param
	return options
}

// TrainingsGetOptions : The TrainingsGet options.
type TrainingsGetOptions struct {
	// The training identifier.
	TrainingID *string `json:"training_id" validate:"required,ne="`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingsGetOptions : Instantiate TrainingsGetOptions
func (*WatsonMachineLearningV4) NewTrainingsGetOptions(trainingID string) *TrainingsGetOptions {
	return &TrainingsGetOptions{
		TrainingID: core.StringPtr(trainingID),
	}
}

// SetTrainingID : Allow user to set TrainingID
func (_options *TrainingsGetOptions) SetTrainingID(trainingID string) *TrainingsGetOptions {
	_options.TrainingID = core.StringPtr(trainingID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingsGetOptions) SetSpaceID(spaceID string) *TrainingsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingsGetOptions) SetProjectID(projectID string) *TrainingsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingsGetOptions) SetHeaders(param map[string]string) *TrainingsGetOptions {
	options.Headers = param
	return options
}

// TrainingsListOptions : The TrainingsList options.
type TrainingsListOptions struct {
	// Token required for token-based pagination. This token cannot be determined by end user. It is generated by the
	// service and it is set in the href available in the `next` field.
	Start *string `json:"start,omitempty"`

	// How many resources should be returned. Default value is 100. Max value is 200.
	Limit *int64 `json:"limit,omitempty"`

	// Compute the total count. May have performance impact.
	TotalCount *bool `json:"total_count,omitempty"`

	// Return only the resources with the given tag value.
	TagValue *string `json:"tag.value,omitempty"`

	// Filter based on 'pipeline' or 'experiment' trainings. Example: `type=pipeline`. If not provided, returns all the
	// trainings.
	Type *string `json:"type,omitempty"`

	// Filter based on on the training job state: queued, running, completed, failed etc.
	State *string `json:"state,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` query parameter has to be given.
	ProjectID *string `json:"project_id,omitempty"`

	// Return the training jobs that are sub-jobs of this parent_id job. The parent_id can be the experiment job ID.
	ParentID *string `json:"parent_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewTrainingsListOptions : Instantiate TrainingsListOptions
func (*WatsonMachineLearningV4) NewTrainingsListOptions() *TrainingsListOptions {
	return &TrainingsListOptions{}
}

// SetStart : Allow user to set Start
func (_options *TrainingsListOptions) SetStart(start string) *TrainingsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *TrainingsListOptions) SetLimit(limit int64) *TrainingsListOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetTotalCount : Allow user to set TotalCount
func (_options *TrainingsListOptions) SetTotalCount(totalCount bool) *TrainingsListOptions {
	_options.TotalCount = core.BoolPtr(totalCount)
	return _options
}

// SetTagValue : Allow user to set TagValue
func (_options *TrainingsListOptions) SetTagValue(tagValue string) *TrainingsListOptions {
	_options.TagValue = core.StringPtr(tagValue)
	return _options
}

// SetType : Allow user to set Type
func (_options *TrainingsListOptions) SetType(typeVar string) *TrainingsListOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetState : Allow user to set State
func (_options *TrainingsListOptions) SetState(state string) *TrainingsListOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *TrainingsListOptions) SetSpaceID(spaceID string) *TrainingsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *TrainingsListOptions) SetProjectID(projectID string) *TrainingsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetParentID : Allow user to set ParentID
func (_options *TrainingsListOptions) SetParentID(parentID string) *TrainingsListOptions {
	_options.ParentID = core.StringPtr(parentID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TrainingsListOptions) SetHeaders(param map[string]string) *TrainingsListOptions {
	options.Headers = param
	return options
}

// AllContentMetadata : The metadata related to the attachments.
type AllContentMetadata struct {
	// The number of attachments associated with the resource.
	TotalCount *int64 `json:"total_count" validate:"required"`

	Attachments []ContentMetadata `json:"attachments" validate:"required"`
}

// UnmarshalAllContentMetadata unmarshals an instance of AllContentMetadata from the specified map of raw messages.
func UnmarshalAllContentMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AllContentMetadata)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalContentMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AllowedIdentity : The allowed identity.
type AllowedIdentity struct {
	// The id of the identity.
	ID *string `json:"id" validate:"required"`

	// The type of the identity.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the AllowedIdentity.Type property.
// The type of the identity.
const (
	AllowedIdentity_Type_Service = "service"
	AllowedIdentity_Type_User    = "user"
)

// NewAllowedIdentity : Instantiate AllowedIdentity (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewAllowedIdentity(id string, typeVar string) (_model *AllowedIdentity, err error) {
	_model = &AllowedIdentity{
		ID:   core.StringPtr(id),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAllowedIdentity unmarshals an instance of AllowedIdentity from the specified map of raw messages.
func UnmarshalAllowedIdentity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AllowedIdentity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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

// BinaryClassification : BinaryClassification struct
type BinaryClassification struct {
	ConfusionMatrices []ConfusionMatrix `json:"confusion_matrices,omitempty"`

	RocCurves []RocCurve `json:"roc_curves,omitempty"`
}

// UnmarshalBinaryClassification unmarshals an instance of BinaryClassification from the specified map of raw messages.
func UnmarshalBinaryClassification(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BinaryClassification)
	err = core.UnmarshalModel(m, "confusion_matrices", &obj.ConfusionMatrices, UnmarshalConfusionMatrix)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "roc_curves", &obj.RocCurves, UnmarshalRocCurve)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BluemixAccount : BluemixAccount struct
type BluemixAccount struct {
	// The account ID.
	ID *string `json:"id" validate:"required"`
}

// UnmarshalBluemixAccount unmarshals an instance of BluemixAccount from the specified map of raw messages.
func UnmarshalBluemixAccount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BluemixAccount)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CommonPatchRequestHelper : CommonPatchRequestHelper struct
type CommonPatchRequestHelper struct {
	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`
}

// UnmarshalCommonPatchRequestHelper unmarshals an instance of CommonPatchRequestHelper from the specified map of raw messages.
func UnmarshalCommonPatchRequestHelper(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CommonPatchRequestHelper)
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonMachineLearningV4) NewCommonPatchRequestHelperPatch(commonPatchRequestHelper *CommonPatchRequestHelper) (_patch []JSONPatchOperation) {
	if commonPatchRequestHelper.Tags != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/tags"),
			Value: commonPatchRequestHelper.Tags,
		})
	}
	if commonPatchRequestHelper.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: commonPatchRequestHelper.Name,
		})
	}
	if commonPatchRequestHelper.Description != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/description"),
			Value: commonPatchRequestHelper.Description,
		})
	}
	if commonPatchRequestHelper.Custom != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/custom"),
			Value: commonPatchRequestHelper.Custom,
		})
	}
	return
}

// ComputeUsageMetrics : Compute usage metrics.
type ComputeUsageMetrics struct {
	Timestamp *strfmt.DateTime `json:"timestamp" validate:"required"`

	GpuCount *float64 `json:"gpu_count,omitempty"`

	Gpus []GpuMetrics `json:"gpus,omitempty"`
}

// UnmarshalComputeUsageMetrics unmarshals an instance of ComputeUsageMetrics from the specified map of raw messages.
func UnmarshalComputeUsageMetrics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComputeUsageMetrics)
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "gpu_count", &obj.GpuCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "gpus", &obj.Gpus, UnmarshalGpuMetrics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfusionMatrix : The confusion matrix for the selected class.
type ConfusionMatrix struct {
	TrueClass *string `json:"true_class" validate:"required"`

	Tp *int64 `json:"tp" validate:"required"`

	Tn *int64 `json:"tn" validate:"required"`

	Fp *int64 `json:"fp" validate:"required"`

	Fn *int64 `json:"fn" validate:"required"`
}

// NewConfusionMatrix : Instantiate ConfusionMatrix (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewConfusionMatrix(trueClass string, tp int64, tn int64, fp int64, fn int64) (_model *ConfusionMatrix, err error) {
	_model = &ConfusionMatrix{
		TrueClass: core.StringPtr(trueClass),
		Tp:        core.Int64Ptr(tp),
		Tn:        core.Int64Ptr(tn),
		Fp:        core.Int64Ptr(fp),
		Fn:        core.Int64Ptr(fn),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalConfusionMatrix unmarshals an instance of ConfusionMatrix from the specified map of raw messages.
func UnmarshalConfusionMatrix(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfusionMatrix)
	err = core.UnmarshalPrimitive(m, "true_class", &obj.TrueClass)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tp", &obj.Tp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tn", &obj.Tn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fp", &obj.Fp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fn", &obj.Fn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Consumption : The consumption part is available only when `plan.version` is `2`. All the values are calculated at the account level
// the instance belongs to, not the instance itself.
type Consumption struct {
	CapacityUnitHours *ConsumptionCapacityUnitHours `json:"capacity_unit_hours,omitempty"`

	GpuCount *ConsumptionGpuCount `json:"gpu_count,omitempty"`

	DoJobCount *ConsumptionDoJobCount `json:"do_job_count,omitempty"`

	// Limit for deployment jobs.
	DeploymentJobCount *ConsumptionDeploymentJobCount `json:"deployment_job_count,omitempty"`

	// Grouped compute usage details presented at the account level the instance belongs to.
	Details []ConsumptionDetails `json:"details,omitempty"`
}

// UnmarshalConsumption unmarshals an instance of Consumption from the specified map of raw messages.
func UnmarshalConsumption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Consumption)
	err = core.UnmarshalModel(m, "capacity_unit_hours", &obj.CapacityUnitHours, UnmarshalConsumptionCapacityUnitHours)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "gpu_count", &obj.GpuCount, UnmarshalConsumptionGpuCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "do_job_count", &obj.DoJobCount, UnmarshalConsumptionDoJobCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deployment_job_count", &obj.DeploymentJobCount, UnmarshalConsumptionDeploymentJobCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "details", &obj.Details, UnmarshalConsumptionDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConsumptionDetails : Compute usage details in a given context and framework.
type ConsumptionDetails struct {
	// Context in which compute resources are consumed.
	Context *string `json:"context" validate:"required"`

	// Machine learning framework or tool.
	Framework *string `json:"framework,omitempty"`

	// The current total computation time (in capacity unit hours) in a given context and framework. It is a sum of both,
	// reserved and already send to BSS, units.
	CapacityUnitHours *float64 `json:"capacity_unit_hours" validate:"required"`
}

// UnmarshalConsumptionDetails unmarshals an instance of ConsumptionDetails from the specified map of raw messages.
func UnmarshalConsumptionDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConsumptionDetails)
	err = core.UnmarshalPrimitive(m, "context", &obj.Context)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "framework", &obj.Framework)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "capacity_unit_hours", &obj.CapacityUnitHours)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContentInfo : The content information to be uploaded.
type ContentInfo struct {
	// The content format of the attachment. This can be one of `native`, `coreML`, `pipeline-node`.
	ContentFormat *string `json:"content_format" validate:"required"`

	// The location of the content to be uploaded.
	Location *string `json:"location" validate:"required"`

	// The file name that will be used when downloading the content from the UI.
	FileName *string `json:"file_name" validate:"required"`

	// The `pipeline_node_id` that corresponds to this content. This is required only if the `content_format` is
	// `pipeline-node` otherwise it is rejected or ignored.
	PipelineNodeID *string `json:"pipeline_node_id,omitempty"`

	// The `deployment_id` that corresponds to this content. This is required only if the `content_format` is `coreml`
	// otherwise it is rejected or ignored.
	DeploymentID *string `json:"deployment_id,omitempty"`
}

// NewContentInfo : Instantiate ContentInfo (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewContentInfo(contentFormat string, location string, fileName string) (_model *ContentInfo, err error) {
	_model = &ContentInfo{
		ContentFormat: core.StringPtr(contentFormat),
		Location:      core.StringPtr(location),
		FileName:      core.StringPtr(fileName),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalContentInfo unmarshals an instance of ContentInfo from the specified map of raw messages.
func UnmarshalContentInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContentInfo)
	err = core.UnmarshalPrimitive(m, "content_format", &obj.ContentFormat)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_name", &obj.FileName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_node_id", &obj.PipelineNodeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContentLocation : Details about the attachments that should be uploaded with this model.
type ContentLocation struct {
	// The content information to be uploaded.
	Contents []ContentInfo `json:"contents" validate:"required"`

	// The data source type like `connection_asset` or `data_asset`.
	Type *string `json:"type" validate:"required"`

	// Connection properties.
	Connection map[string]string `json:"connection,omitempty"`

	// Location properties.
	Location map[string]string `json:"location,omitempty"`
}

// Constants associated with the ContentLocation.Type property.
// The data source type like `connection_asset` or `data_asset`.
const (
	ContentLocation_Type_ConnectionAsset = "connection_asset"
	ContentLocation_Type_DataAsset       = "data_asset"
	ContentLocation_Type_Fs              = "fs"
	ContentLocation_Type_URL             = "url"
)

// NewContentLocation : Instantiate ContentLocation (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewContentLocation(contents []ContentInfo, typeVar string) (_model *ContentLocation, err error) {
	_model = &ContentLocation{
		Contents: contents,
		Type:     core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalContentLocation unmarshals an instance of ContentLocation from the specified map of raw messages.
func UnmarshalContentLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContentLocation)
	err = core.UnmarshalModel(m, "contents", &obj.Contents, UnmarshalContentInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection", &obj.Connection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContentMetadata : The metadata related to the attachment.
type ContentMetadata struct {
	// The content id for the attachment.
	AttachmentID *string `json:"attachment_id" validate:"required"`

	// The type of the content.
	ContentFormat *string `json:"content_format" validate:"required"`

	// The name of the attachment.
	Name *string `json:"name,omitempty"`

	// The `pipeline_node_id` that corresponds to this content.
	PipelineNodeID *string `json:"pipeline_node_id,omitempty"`

	// The `deployment_id` that corresponds to this content.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// This will be set to `true` if the content has been persisted. If this content was part of the import process and the
	// upload of the content failed then the message can be found in the
	// `warnings` field of the `system_details` that are returned with the model details.
	Persisted *bool `json:"persisted,omitempty"`
}

// UnmarshalContentMetadata unmarshals an instance of ContentMetadata from the specified map of raw messages.
func UnmarshalContentMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContentMetadata)
	err = core.UnmarshalPrimitive(m, "attachment_id", &obj.AttachmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "content_format", &obj.ContentFormat)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_node_id", &obj.PipelineNodeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "persisted", &obj.Persisted)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataConnectionReference : A reference to data with an optional data schema.
type DataConnectionReference struct {
	// Optional item identification inside a collection.
	ID *string `json:"id,omitempty"`

	// The data source type like `connection_asset` or `data_asset`.
	Type *string `json:"type" validate:"required"`

	// Contains a set of fields specific to each connection. See here for [details about specifying
	// connections](#datareferences).
	Connection interface{} `json:"connection,omitempty"`

	// Contains a set of fields that describe the location of the data with respect to the `connection`.
	Location map[string]string `json:"location" validate:"required"`

	// The schema of the expected data, see
	// [datarecord-metadata-v2-schema](https://raw.githubusercontent.com/elyra-ai/pipeline-schemas/master/common-pipeline/datarecord-metadata/datarecord-metadata-v2-schema.json)
	// for the schema definition.
	Schema *DataSchema `json:"schema,omitempty"`
}

// Constants associated with the DataConnectionReference.Type property.
// The data source type like `connection_asset` or `data_asset`.
const (
	DataConnectionReference_Type_ConnectionAsset = "connection_asset"
	DataConnectionReference_Type_DataAsset       = "data_asset"
	DataConnectionReference_Type_Fs              = "fs"
	DataConnectionReference_Type_URL             = "url"
)

// NewDataConnectionReference : Instantiate DataConnectionReference (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewDataConnectionReference(typeVar string, location map[string]string) (_model *DataConnectionReference, err error) {
	_model = &DataConnectionReference{
		Type:     core.StringPtr(typeVar),
		Location: location,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDataConnectionReference unmarshals an instance of DataConnectionReference from the specified map of raw messages.
func UnmarshalDataConnectionReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataConnectionReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection", &obj.Connection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schema", &obj.Schema, UnmarshalDataSchema)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataSchema : The schema of the expected data, see
// [datarecord-metadata-v2-schema](https://raw.githubusercontent.com/elyra-ai/pipeline-schemas/master/common-pipeline/datarecord-metadata/datarecord-metadata-v2-schema.json)
// for the schema definition.
type DataSchema struct {
	// An id to identify a schema.
	ID *string `json:"id" validate:"required"`

	// A name for the schema.
	Name *string `json:"name,omitempty"`

	// The fields that describe the data schema.
	Fields []interface{} `json:"fields" validate:"required"`
}

// NewDataSchema : Instantiate DataSchema (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewDataSchema(id string, fields []interface{}) (_model *DataSchema, err error) {
	_model = &DataSchema{
		ID:     core.StringPtr(id),
		Fields: fields,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalDataSchema unmarshals an instance of DataSchema from the specified map of raw messages.
func UnmarshalDataSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataSchema)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentEntity : The definition of the deployment.
type DeploymentEntity struct {
	// Tags that can be used when searching for resources.
	Tags []string `json:"tags,omitempty"`

	// The space that contains the resource.
	SpaceID *string `json:"space_id" validate:"required"`

	// The name of the deployment.
	Name *string `json:"name,omitempty"`

	// A description of the deployment.
	Description *string `json:"description,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// A reference to a resource.
	Asset *Rel `json:"asset,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityResultHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`

	// Indicates that this is an online deployment. An empty object has to be specified.
	// More properties will be added later on to setup the online deployment.
	// The `serving_name` can be provided in the `online.parameters`. The serving name can only have the characters
	// [a-z,0-9,_]
	// and the length should not be more than 36 characters. The `serving_name` can be used in the prediction URL in place
	// of the `deployment_id`.
	Online *DeploymentEntityRequestOnline `json:"online,omitempty"`

	// Indicates that this is a batch deployment. An empty object has to be specified.
	// More properties will be added later on to setup the batch deployment.
	Batch *DeploymentEntityRequestBatch `json:"batch,omitempty"`

	// Indicates that this is a Shiny application deployment.
	RShiny *DeploymentEntityRequestRShiny `json:"r_shiny,omitempty"`

	// Specifies the type of the asset on which deployment is created.
	DeployedAssetType *string `json:"deployed_asset_type,omitempty"`

	// Specifies the current status, additional information about the deployment
	// and any failure messages in case of deployment failures.
	Status *DeploymentEntityStatus `json:"status,omitempty"`
}

// Constants associated with the DeploymentEntity.DeployedAssetType property.
// Specifies the type of the asset on which deployment is created.
const (
	DeploymentEntity_DeployedAssetType_Do       = "do"
	DeploymentEntity_DeployedAssetType_Function = "function"
	DeploymentEntity_DeployedAssetType_Model    = "model"
	DeploymentEntity_DeployedAssetType_PyScript = "py_script"
	DeploymentEntity_DeployedAssetType_RShiny   = "r_shiny"
)

// UnmarshalDeploymentEntity unmarshals an instance of DeploymentEntity from the specified map of raw messages.
func UnmarshalDeploymentEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentEntity)
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hybrid_pipeline_hardware_specs", &obj.HybridPipelineHardwareSpecs, UnmarshalJobEntityResultHybridPipelineHardwareSpecsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "online", &obj.Online, UnmarshalDeploymentEntityRequestOnline)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "batch", &obj.Batch, UnmarshalDeploymentEntityRequestBatch)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "r_shiny", &obj.RShiny, UnmarshalDeploymentEntityRequestRShiny)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deployed_asset_type", &obj.DeployedAssetType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalDeploymentEntityStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentJobDefinitionPatchHelper : DeploymentJobDefinitionPatchHelper struct
type DeploymentJobDefinitionPatchHelper struct {
	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// A reference to a resource.
	Deployment *SimpleRel `json:"deployment,omitempty"`
}

// UnmarshalDeploymentJobDefinitionPatchHelper unmarshals an instance of DeploymentJobDefinitionPatchHelper from the specified map of raw messages.
func UnmarshalDeploymentJobDefinitionPatchHelper(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentJobDefinitionPatchHelper)
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalSimpleRel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonMachineLearningV4) NewDeploymentJobDefinitionPatchHelperPatch(deploymentJobDefinitionPatchHelper *DeploymentJobDefinitionPatchHelper) (_patch []JSONPatchOperation) {
	if deploymentJobDefinitionPatchHelper.Tags != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/tags"),
			Value: deploymentJobDefinitionPatchHelper.Tags,
		})
	}
	if deploymentJobDefinitionPatchHelper.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: deploymentJobDefinitionPatchHelper.Name,
		})
	}
	if deploymentJobDefinitionPatchHelper.Description != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/description"),
			Value: deploymentJobDefinitionPatchHelper.Description,
		})
	}
	if deploymentJobDefinitionPatchHelper.Custom != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/custom"),
			Value: deploymentJobDefinitionPatchHelper.Custom,
		})
	}
	if deploymentJobDefinitionPatchHelper.Deployment != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/deployment"),
			Value: deploymentJobDefinitionPatchHelper.Deployment,
		})
	}
	return
}

// DeploymentPatchRequestHelper : DeploymentPatchRequestHelper struct
type DeploymentPatchRequestHelper struct {
	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// A reference to a resource.
	Asset *Rel `json:"asset,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityResultHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`

	// Specify this section if deploying an Shiny application.
	RShiny *DeploymentPatchRequestHelperRShiny `json:"r_shiny,omitempty"`
}

// UnmarshalDeploymentPatchRequestHelper unmarshals an instance of DeploymentPatchRequestHelper from the specified map of raw messages.
func UnmarshalDeploymentPatchRequestHelper(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentPatchRequestHelper)
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "asset", &obj.Asset, UnmarshalRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hybrid_pipeline_hardware_specs", &obj.HybridPipelineHardwareSpecs, UnmarshalJobEntityResultHybridPipelineHardwareSpecsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "r_shiny", &obj.RShiny, UnmarshalDeploymentPatchRequestHelperRShiny)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonMachineLearningV4) NewDeploymentPatchRequestHelperPatch(deploymentPatchRequestHelper *DeploymentPatchRequestHelper) (_patch []JSONPatchOperation) {
	if deploymentPatchRequestHelper.Tags != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/tags"),
			Value: deploymentPatchRequestHelper.Tags,
		})
	}
	if deploymentPatchRequestHelper.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: deploymentPatchRequestHelper.Name,
		})
	}
	if deploymentPatchRequestHelper.Description != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/description"),
			Value: deploymentPatchRequestHelper.Description,
		})
	}
	if deploymentPatchRequestHelper.Custom != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/custom"),
			Value: deploymentPatchRequestHelper.Custom,
		})
	}
	if deploymentPatchRequestHelper.Asset != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/asset"),
			Value: deploymentPatchRequestHelper.Asset,
		})
	}
	if deploymentPatchRequestHelper.HardwareSpec != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/hardware_spec"),
			Value: deploymentPatchRequestHelper.HardwareSpec,
		})
	}
	if deploymentPatchRequestHelper.HybridPipelineHardwareSpecs != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/hybrid_pipeline_hardware_specs"),
			Value: deploymentPatchRequestHelper.HybridPipelineHardwareSpecs,
		})
	}
	if deploymentPatchRequestHelper.RShiny != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/r_shiny"),
			Value: deploymentPatchRequestHelper.RShiny,
		})
	}
	return
}

// DeploymentResource : DeploymentResource struct
type DeploymentResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *ResourceMeta `json:"metadata,omitempty"`

	// The definition of the deployment.
	Entity *DeploymentEntity `json:"entity,omitempty"`
}

// UnmarshalDeploymentResource unmarshals an instance of DeploymentResource from the specified map of raw messages.
func UnmarshalDeploymentResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalResourceMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalDeploymentEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentResources : The deployment resources.
type DeploymentResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *DeploymentResourcesFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *DeploymentResourcesNext `json:"next,omitempty"`

	// A list of deployment resources. This will be empty if 'stats' query parameter is passed with 'true'.
	Resources []DeploymentResource `json:"resources,omitempty"`

	// System details including warnings and stats. This will get populated only if 'stats' query parameter is passed with
	// 'true'.
	System *DeploymentResourcesSystem `json:"system,omitempty"`
}

// UnmarshalDeploymentResources unmarshals an instance of DeploymentResources from the specified map of raw messages.
func UnmarshalDeploymentResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalDeploymentResourcesFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalDeploymentResourcesNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalDeploymentResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalDeploymentResourcesSystem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeploymentSystemDetails : Optional details provided by the service about statistics of the number of deployments created with in a space or
// across all spaces.
type DeploymentSystemDetails struct {
	// Any warnings coming from the system.
	Warnings []Warning `json:"warnings,omitempty"`

	// The stats about deployments for a space or across all spaces.
	Stats []Stats `json:"stats,omitempty"`
}

// UnmarshalDeploymentSystemDetails unmarshals an instance of DeploymentSystemDetails from the specified map of raw messages.
func UnmarshalDeploymentSystemDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeploymentSystemDetails)
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalWarning)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stats", &obj.Stats, UnmarshalStats)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Error : The data returned when an error is encountered.
type Error struct {
	// An identifier that can be used to trace the request.
	Trace *string `json:"trace" validate:"required"`

	// The list of errors.
	Errors []ErrorErrorsItem `json:"errors" validate:"required"`
}

// UnmarshalError unmarshals an instance of Error from the specified map of raw messages.
func UnmarshalError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Error)
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalErrorErrorsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EvaluationDefinition : The optional evaluation definition.
type EvaluationDefinition struct {
	// The evaluation method.
	Method *string `json:"method,omitempty"`

	// The evaluation metrics.
	Metrics []EvaluationMetric `json:"metrics" validate:"required"`
}

// Constants associated with the EvaluationDefinition.Method property.
// The evaluation method.
const (
	EvaluationDefinition_Method_Binary     = "binary"
	EvaluationDefinition_Method_Multiclass = "multiclass"
	EvaluationDefinition_Method_Regression = "regression"
)

// NewEvaluationDefinition : Instantiate EvaluationDefinition (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewEvaluationDefinition(metrics []EvaluationMetric) (_model *EvaluationDefinition, err error) {
	_model = &EvaluationDefinition{
		Metrics: metrics,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalEvaluationDefinition unmarshals an instance of EvaluationDefinition from the specified map of raw messages.
func UnmarshalEvaluationDefinition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EvaluationDefinition)
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalEvaluationMetric)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EvaluationMetric : An evaluation metric.
type EvaluationMetric struct {
	Name *string `json:"name" validate:"required"`

	Maximize *bool `json:"maximize,omitempty"`
}

// NewEvaluationMetric : Instantiate EvaluationMetric (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewEvaluationMetric(name string) (_model *EvaluationMetric, err error) {
	_model = &EvaluationMetric{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalEvaluationMetric unmarshals an instance of EvaluationMetric from the specified map of raw messages.
func UnmarshalEvaluationMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EvaluationMetric)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximize", &obj.Maximize)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExperimentResource : The information for an experiment.
type ExperimentResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *ExperimentResourceMetadata `json:"metadata" validate:"required"`

	// The details of the experiment to be created.
	Entity *ExperimentResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalExperimentResource unmarshals an instance of ExperimentResource from the specified map of raw messages.
func UnmarshalExperimentResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExperimentResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalExperimentResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalExperimentResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExperimentResources : System details.
type ExperimentResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	// A list of experiments.
	Resources []ExperimentResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalExperimentResources unmarshals an instance of ExperimentResources from the specified map of raw messages.
func UnmarshalExperimentResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExperimentResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalExperimentResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ExperimentResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// FeatureImportance : FeatureImportance struct
type FeatureImportance struct {
	Stage *string `json:"stage" validate:"required"`

	ComputationType *string `json:"computation_type" validate:"required"`

	// The feature names and importance values as numbers.
	Features interface{} `json:"features" validate:"required"`
}

// NewFeatureImportance : Instantiate FeatureImportance (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewFeatureImportance(stage string, computationType string, features interface{}) (_model *FeatureImportance, err error) {
	_model = &FeatureImportance{
		Stage:           core.StringPtr(stage),
		ComputationType: core.StringPtr(computationType),
		Features:        features,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalFeatureImportance unmarshals an instance of FeatureImportance from the specified map of raw messages.
func UnmarshalFeatureImportance(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FeatureImportance)
	err = core.UnmarshalPrimitive(m, "stage", &obj.Stage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "computation_type", &obj.ComputationType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearning : Federated Learning is a Technical Preview.
type FederatedLearning struct {
	// The model type for federated_learning.
	Model *FederatedLearningModel `json:"model,omitempty"`

	// The fusion type for federated learning.
	FusionType *string `json:"fusion_type" validate:"required"`

	// The remote training for federated learning.
	RemoteTraining *FederatedLearningRemoteTraining `json:"remote_training" validate:"required"`

	// The number of training iterations to complete between the aggregator and the remote systems unless termination
	// accuracy is achieved first.
	Rounds *int64 `json:"rounds,omitempty"`

	// A boolean expression that describes the conditions, in terms of model metrics, under which training should complete.
	TerminationPredicate *string `json:"termination_predicate,omitempty"`

	// The total number of passes over the local training dataset to train a model.
	Epochs *int64 `json:"epochs,omitempty"`

	// The optimizer for federated learning.
	Optimizer *FederatedLearningOptimizer `json:"optimizer,omitempty"`

	// The loss function to use in the boosting process.
	Loss *string `json:"loss,omitempty"`

	// The metrics for federated learning.
	Metrics *string `json:"metrics,omitempty"`

	// The maximum depth of each tree.
	MaxDepth *int64 `json:"max_depth,omitempty"`

	// Controls how rapidly to change the model in response to the estimated error each time the model weights are updated.
	LearningRate *float64 `json:"learning_rate,omitempty"`

	// The coefficient for the L2 regularizer.
	L2Regularization *float64 `json:"l2_regularization,omitempty"`

	// The maximum number of bins to use for non-missing values across all features.
	MaxBins *int64 `json:"max_bins,omitempty"`

	// The maximum number of leaves for each tree.
	MaxLeafNodes *int64 `json:"max_leaf_nodes,omitempty"`

	// The minimum number of samples per leaf.
	MinSamplesLeaf *int64 `json:"min_samples_leaf,omitempty"`

	// Seed used to set the pseudo-random number generator.
	RandomState *int64 `json:"random_state,omitempty"`

	// A general programming term to produce logging output.
	Verbose *bool `json:"verbose,omitempty"`

	// Number of target classes for the classification model.
	NumClasses *int64 `json:"num_classes,omitempty"`

	// Specifies the failure tolerance, for example the maximum number of parties that would possibly be failing during
	// training.
	ByzantineThreshold *int64 `json:"byzantine_threshold,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// The version for federated learning.
	Version *string `json:"version,omitempty"`

	// The log level (`critical`, `error`, `warning`, `info`, `debug`) for federated learning.
	LogLevel *string `json:"log_level,omitempty"`
}

// Constants associated with the FederatedLearning.FusionType property.
// The fusion type for federated learning.
const (
	FederatedLearning_FusionType_Avg              = "avg"
	FederatedLearning_FusionType_CoordinateMedian = "coordinate_median"
	FederatedLearning_FusionType_Dt               = "dt"
	FederatedLearning_FusionType_Gradient         = "gradient"
	FederatedLearning_FusionType_IterAvg          = "iter_avg"
	FederatedLearning_FusionType_IterAvgDp        = "iter_avg_dp"
	FederatedLearning_FusionType_Krum             = "krum"
	FederatedLearning_FusionType_NaiveBayes       = "naive_bayes"
	FederatedLearning_FusionType_Pfnm             = "pfnm"
	FederatedLearning_FusionType_Spahm            = "spahm"
	FederatedLearning_FusionType_XgbClassifier    = "xgb_classifier"
	FederatedLearning_FusionType_XgbRegressor     = "xgb_regressor"
)

// Constants associated with the FederatedLearning.LogLevel property.
// The log level (`critical`, `error`, `warning`, `info`, `debug`) for federated learning.
const (
	FederatedLearning_LogLevel_Critical = "critical"
	FederatedLearning_LogLevel_Debug    = "debug"
	FederatedLearning_LogLevel_Error    = "error"
	FederatedLearning_LogLevel_Info     = "info"
	FederatedLearning_LogLevel_Warning  = "warning"
)

// NewFederatedLearning : Instantiate FederatedLearning (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewFederatedLearning(fusionType string, remoteTraining *FederatedLearningRemoteTraining) (_model *FederatedLearning, err error) {
	_model = &FederatedLearning{
		FusionType:     core.StringPtr(fusionType),
		RemoteTraining: remoteTraining,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalFederatedLearning unmarshals an instance of FederatedLearning from the specified map of raw messages.
func UnmarshalFederatedLearning(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearning)
	err = core.UnmarshalModel(m, "model", &obj.Model, UnmarshalFederatedLearningModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fusion_type", &obj.FusionType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "remote_training", &obj.RemoteTraining, UnmarshalFederatedLearningRemoteTraining)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rounds", &obj.Rounds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "termination_predicate", &obj.TerminationPredicate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "epochs", &obj.Epochs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "optimizer", &obj.Optimizer, UnmarshalFederatedLearningOptimizer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "loss", &obj.Loss)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metrics", &obj.Metrics)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_depth", &obj.MaxDepth)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "learning_rate", &obj.LearningRate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "l2_regularization", &obj.L2Regularization)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_bins", &obj.MaxBins)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_leaf_nodes", &obj.MaxLeafNodes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min_samples_leaf", &obj.MinSamplesLeaf)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "random_state", &obj.RandomState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "verbose", &obj.Verbose)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "num_classes", &obj.NumClasses)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "byzantine_threshold", &obj.ByzantineThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_level", &obj.LogLevel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FederatedLearningInfo : Federated learning info.
type FederatedLearningInfo struct {
	Aggregator *FederatedLearningInfoAggregator `json:"aggregator,omitempty"`

	// Remote training systems.
	RemoteTrainingSystems []FederatedLearningInfoRemoteTrainingSystemsItem `json:"remote_training_systems,omitempty"`
}

// UnmarshalFederatedLearningInfo unmarshals an instance of FederatedLearningInfo from the specified map of raw messages.
func UnmarshalFederatedLearningInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FederatedLearningInfo)
	err = core.UnmarshalModel(m, "aggregator", &obj.Aggregator, UnmarshalFederatedLearningInfoAggregator)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "remote_training_systems", &obj.RemoteTrainingSystems, UnmarshalFederatedLearningInfoRemoteTrainingSystemsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FunctionResource : The information for a function.
type FunctionResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *FunctionResourceMetadata `json:"metadata" validate:"required"`

	// The details of the function to be created.
	Entity *FunctionResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalFunctionResource unmarshals an instance of FunctionResource from the specified map of raw messages.
func UnmarshalFunctionResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FunctionResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalFunctionResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalFunctionResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FunctionResources : System details.
type FunctionResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *FunctionResourcesFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *FunctionResourcesNext `json:"next,omitempty"`

	// A list of functions.
	Resources []FunctionResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalFunctionResources unmarshals an instance of FunctionResources from the specified map of raw messages.
func UnmarshalFunctionResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FunctionResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFunctionResourcesFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalFunctionResourcesNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalFunctionResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *FunctionResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// GpuMetrics : GPU metrics.
type GpuMetrics struct {
	Type *string `json:"type" validate:"required"`

	Memory *GpuMetricsMemory `json:"memory,omitempty"`
}

// UnmarshalGpuMetrics unmarshals an instance of GpuMetrics from the specified map of raw messages.
func UnmarshalGpuMetrics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GpuMetrics)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "memory", &obj.Memory, UnmarshalGpuMetricsMemory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HardwareSpecRel : A hardware specification.
type HardwareSpecRel struct {
	// The id of the hardware specification. One, and only one, of `id` or `name` must be set.
	ID *string `json:"id,omitempty"`

	// The revision of the hardware specification.
	Rev *string `json:"rev,omitempty"`

	// The name of the hardware specification. One, and only one, of `id` or `name` must be set.
	Name *string `json:"name,omitempty"`

	// The number of nodes applied to a computation.
	NumNodes *int64 `json:"num_nodes,omitempty"`
}

// UnmarshalHardwareSpecRel unmarshals an instance of HardwareSpecRel from the specified map of raw messages.
func UnmarshalHardwareSpecRel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HardwareSpecRel)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "num_nodes", &obj.NumNodes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HyperParameter : A set of hyper parameters.
type HyperParameter struct {
	// The name of the hyper parameters.
	Name *string `json:"name" validate:"required"`

	// An object containing floating point properties 'min_value', 'max_value', 'step', and 'power', or a list of either
	// strings or floating point values.
	Items interface{} `json:"items" validate:"required"`
}

// NewHyperParameter : Instantiate HyperParameter (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewHyperParameter(name string, items interface{}) (_model *HyperParameter, err error) {
	_model = &HyperParameter{
		Name:  core.StringPtr(name),
		Items: items,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalHyperParameter unmarshals an instance of HyperParameter from the specified map of raw messages.
func UnmarshalHyperParameter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HyperParameter)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "items", &obj.Items)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceResource : InstanceResource struct
type InstanceResource struct {
	// Common metadata for a resource.
	Metadata *ResourceMetaBase `json:"metadata,omitempty"`

	Entity *InstanceResourceEntity `json:"entity,omitempty"`
}

// UnmarshalInstanceResource unmarshals an instance of InstanceResource from the specified map of raw messages.
func UnmarshalInstanceResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalResourceMetaBase)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalInstanceResourceEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceResourceEntity : InstanceResourceEntity struct
type InstanceResourceEntity struct {
	// Status of the service instance (active|inactive).
	Status *string `json:"status" validate:"required"`

	Plan *InstanceResourceEntityPlan `json:"plan" validate:"required"`

	Crn *string `json:"crn,omitempty"`

	Account *BluemixAccount `json:"account,omitempty"`

	// The consumption part is available only when `plan.version` is `2`. All the values are calculated at the account
	// level the instance belongs to, not the instance itself.
	Consumption *Consumption `json:"consumption,omitempty"`

	// Cloud Service Endpoints (CSE) the instance is enabled for. Possible values are `public`, `private` and
	// `public-and-private`.
	ServiceEndpoints *string `json:"service_endpoints" validate:"required"`
}

// Constants associated with the InstanceResourceEntity.Status property.
// Status of the service instance (active|inactive).
const (
	InstanceResourceEntity_Status_Active   = "active"
	InstanceResourceEntity_Status_Inactive = "inactive"
)

// UnmarshalInstanceResourceEntity unmarshals an instance of InstanceResourceEntity from the specified map of raw messages.
func UnmarshalInstanceResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceResourceEntity)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "plan", &obj.Plan, UnmarshalInstanceResourceEntityPlan)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account", &obj.Account, UnmarshalBluemixAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "consumption", &obj.Consumption, UnmarshalConsumption)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_endpoints", &obj.ServiceEndpoints)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstanceResources : Information for paging when querying resources.
type InstanceResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	Resources []InstanceResource `json:"resources,omitempty"`
}

// UnmarshalInstanceResources unmarshals an instance of InstanceResources from the specified map of raw messages.
func UnmarshalInstanceResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalInstanceResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *InstanceResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// IntermediateModel : The details of the intermediate model.
type IntermediateModel struct {
	// The name of the pipeline.
	Name *string `json:"name" validate:"required"`

	// The process that generated this pipeline.
	Process *string `json:"process" validate:"required"`

	// The location of the intermediate model.
	Location *ModelLocation `json:"location,omitempty"`

	NotebookLocation *string `json:"notebook_location,omitempty"`

	SdkNotebookLocation *string `json:"sdk_notebook_location,omitempty"`

	// The nodes of this pipeline.
	PipelineNodes []string `json:"pipeline_nodes,omitempty"`

	// The steps that lead to the creation of this pipeline.
	CompositionSteps []string `json:"composition_steps,omitempty"`

	Duration *int64 `json:"duration,omitempty"`

	ModelAsset *string `json:"model_asset,omitempty"`
}

// NewIntermediateModel : Instantiate IntermediateModel (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewIntermediateModel(name string, process string) (_model *IntermediateModel, err error) {
	_model = &IntermediateModel{
		Name:    core.StringPtr(name),
		Process: core.StringPtr(process),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalIntermediateModel unmarshals an instance of IntermediateModel from the specified map of raw messages.
func UnmarshalIntermediateModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntermediateModel)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "process", &obj.Process)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "location", &obj.Location, UnmarshalModelLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notebook_location", &obj.NotebookLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sdk_notebook_location", &obj.SdkNotebookLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_nodes", &obj.PipelineNodes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "composition_steps", &obj.CompositionSteps)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "duration", &obj.Duration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_asset", &obj.ModelAsset)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobDecisionOptimizationRequest : Details about the input/output data and other properties to be used for a batch deployment job of a decision
// optimization problem.
//
// Use `input_data` property to specify the input data for batch processing as part of the job's payload. The
// `input_data` property is mutually exclusive with `input_data_references` property, only use one of these. When
// `input_data` is specified, the processed output of batch deployment job will be available in `scoring.predictions`
// parameter in the deployment job response.
//
// Use `input_data_references` property to specify the details pertaining to the remote source where the input data for
// batch deployment job is available. The `input_data_references` must be used with `output_data_references`. The
// `input_data_references` property is mutually exclusive with `input_data` property, only use one of these.
//
// Use `output_data_references` property to specify the details pertaining to the remote source where the input data for
// batch deployment job is available. The `output_data_references` must be used with `input_data_references`.
type JobDecisionOptimizationRequest struct {
	// To control solve behavior, you can specify solve parameters in your request as key-value pairs.
	SolveParameters interface{} `json:"solve_parameters,omitempty"`

	// A list of payloads.
	InputData []ScoringPayloadOptim `json:"input_data,omitempty"`

	// A list of input data references.
	InputDataReferences []ObjectLocationOptim `json:"input_data_references,omitempty"`

	// A list of output payloads.
	OutputData []ScoringPayloadOptim `json:"output_data,omitempty"`

	// A list of output data references.
	OutputDataReferences []ObjectLocationOptim `json:"output_data_references,omitempty"`
}

// UnmarshalJobDecisionOptimizationRequest unmarshals an instance of JobDecisionOptimizationRequest from the specified map of raw messages.
func UnmarshalJobDecisionOptimizationRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobDecisionOptimizationRequest)
	err = core.UnmarshalPrimitive(m, "solve_parameters", &obj.SolveParameters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data", &obj.InputData, UnmarshalScoringPayloadOptim)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data_references", &obj.InputDataReferences, UnmarshalObjectLocationOptim)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output_data", &obj.OutputData, UnmarshalScoringPayloadOptim)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output_data_references", &obj.OutputDataReferences, UnmarshalObjectLocationOptim)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobDecisionOptimizationResult : The solve state for a Decision Optimzation job.
type JobDecisionOptimizationResult struct {
	// To control solve behavior, you can specify solve parameters in your request as key-value pairs.
	SolveParameters interface{} `json:"solve_parameters,omitempty"`

	// A list of payloads.
	InputData []ScoringPayloadOptim `json:"input_data,omitempty"`

	// A list of input data references.
	InputDataReferences []ObjectLocationOptim `json:"input_data_references,omitempty"`

	// A list of output payloads.
	OutputData []ScoringPayloadOptim `json:"output_data,omitempty"`

	// A list of output data references.
	OutputDataReferences []ObjectLocationOptim `json:"output_data_references,omitempty"`

	// The status of the job.
	Status *JobStatus `json:"status,omitempty"`

	// The solve state for a Decision Optimzation job.
	SolveState *SolveState `json:"solve_state,omitempty"`
}

// UnmarshalJobDecisionOptimizationResult unmarshals an instance of JobDecisionOptimizationResult from the specified map of raw messages.
func UnmarshalJobDecisionOptimizationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobDecisionOptimizationResult)
	err = core.UnmarshalPrimitive(m, "solve_parameters", &obj.SolveParameters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data", &obj.InputData, UnmarshalScoringPayloadOptim)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data_references", &obj.InputDataReferences, UnmarshalObjectLocationOptim)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output_data", &obj.OutputData, UnmarshalScoringPayloadOptim)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output_data_references", &obj.OutputDataReferences, UnmarshalObjectLocationOptim)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalJobStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "solve_state", &obj.SolveState, UnmarshalSolveState)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobResource : The information for a deployment job definition.
type JobResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *JobResourceMetadata `json:"metadata" validate:"required"`

	// Details about the batch deployment job.
	//
	// The `deployment` is a reference to the deployment associated with the deployment job or deployment job definition.
	//
	// The `scoring` and `decision_optimization` properties are mutually exclusive.
	// Specify only one of these when submitting a batch deployment job.
	//
	// Use `hybrid_pipeline_hardware_specs` only in a batch deployment job of a hybrid pipeline
	// in order to specify compute configuration for each pipeline node. For all other cases use `hardware_spec`
	// to specify compute configuration.
	//
	// In case of output data references where the data asset is a remote database, users can specify if the batch
	// deployment output must be appended to the table or if the table is to be truncated and output data updated.
	// `output_data_references.location.write_mode` parameter can be used to for this purpose.
	// The values `truncate` or `append` can be specified for `output_data_references.location.write_mode`
	// parameter.
	// * Specifying `truncate` as value will truncate the table and the batch output data will be inserted.
	// * Specifying `append` as value will insert the batch output data to the remote database table.
	// * The `write_mode` parameter is applicable only for `output_data_references` parameter.
	// * The `write_mode` parameter will be applicable only for remote database related data assets.
	// This parameter will not be applicable for local data asset or COS based data asset.
	Entity *JobResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalJobResource unmarshals an instance of JobResource from the specified map of raw messages.
func UnmarshalJobResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalJobResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalJobResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobResources : System details.
type JobResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *JobResourcesFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *JobResourcesNext `json:"next,omitempty"`

	// A list of deployment job definitions.
	Resources []JobResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalJobResources unmarshals an instance of JobResources from the specified map of raw messages.
func UnmarshalJobResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalJobResourcesFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalJobResourcesNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalJobResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *JobResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// JobScoringRequest : Details about the input/output data and other properties to be used for a batch deployment job of a model, Python
// Function or a Python Scripts.
//
// Use `input_data` property to specify the input data for batch processing as part of the job's payload. The
// `input_data` property is mutually exclusive with `input_data_references` property, only use one of these. When
// `input_data` is specified, the processed output of batch deployment job will be available in `scoring.predictions`
// parameter in the deployment job response. `input_data` property is not supported for batch deployment of Python
// Scripts.
//
// Use `input_data_references` property to specify the details pertaining to the remote source where the input data for
// batch deployment job is available. The `input_data_references` must be used with `output_data_references`. The
// `input_data_references` property is mutually exclusive with `input_data` property, only use one of these. The
// `input_data_references` property is not supported for batch deployment job of Spark models and Python Functions.
//
// Use `output_data_references` property to specify the details pertaining to the remote source where the input data for
// batch deployment job is available. `output_data_references` must be used with `input_data_references`. The
// `output_data_references` property is not supported for batch deployment job of Spark models and Python Functions.
type JobScoringRequest struct {
	// A list of payloads.
	InputData []ScoringPayload `json:"input_data,omitempty"`

	// A list of input data references.
	InputDataReferences []DataConnectionReference `json:"input_data_references,omitempty"`

	// A reference to data with an optional data schema.
	OutputDataReference *DataConnectionReference `json:"output_data_reference,omitempty"`

	// A list of evaluation specifications.
	Evaluations []EvaluationsSpecItem `json:"evaluations,omitempty"`

	// This property is used to specify environment variables and their values required to be consumed by the batch
	// deployment job. The environment variables and values must be specified as key-value pairs.
	//
	// This property is currently supported only for Python Scripts in batch deployment jobs.
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`
}

// UnmarshalJobScoringRequest unmarshals an instance of JobScoringRequest from the specified map of raw messages.
func UnmarshalJobScoringRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobScoringRequest)
	err = core.UnmarshalModel(m, "input_data", &obj.InputData, UnmarshalScoringPayload)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data_references", &obj.InputDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output_data_reference", &obj.OutputDataReference, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "evaluations", &obj.Evaluations, UnmarshalEvaluationsSpecItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_variables", &obj.EnvironmentVariables)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobScoringResult : The status of the job.
type JobScoringResult struct {
	// A list of payloads.
	InputData []ScoringPayload `json:"input_data,omitempty"`

	// A list of input data references.
	InputDataReferences []DataConnectionReference `json:"input_data_references,omitempty"`

	// A reference to data with an optional data schema.
	OutputDataReference *DataConnectionReference `json:"output_data_reference,omitempty"`

	// A list of evaluation specifications.
	Evaluations []JobScoringResultEvaluationsItem `json:"evaluations,omitempty"`

	// This property is used to specify environment variables and their values required to be consumed by the batch
	// deployment job. The environment variables and values must be specified as key-value pairs.
	//
	// This property is currently supported only for Python Scripts in batch deployment jobs.
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`

	// The predictions.
	Predictions []ScoringPayload `json:"predictions,omitempty"`

	// The status of the job.
	Status *JobStatus `json:"status,omitempty"`
}

// UnmarshalJobScoringResult unmarshals an instance of JobScoringResult from the specified map of raw messages.
func UnmarshalJobScoringResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobScoringResult)
	err = core.UnmarshalModel(m, "input_data", &obj.InputData, UnmarshalScoringPayload)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input_data_references", &obj.InputDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output_data_reference", &obj.OutputDataReference, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "evaluations", &obj.Evaluations, UnmarshalJobScoringResultEvaluationsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_variables", &obj.EnvironmentVariables)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "predictions", &obj.Predictions, UnmarshalScoringPayload)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalJobStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobStatus : The status of the job.
type JobStatus struct {
	// The state of the job.
	State *string `json:"state,omitempty"`

	// The date when the job started to run.
	RunningAt *strfmt.DateTime `json:"running_at,omitempty"`

	// The date that the job finished.
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	// An optional message related to the job status.
	Message *JobStatusMessage `json:"message,omitempty"`

	// The data returned when an error is encountered.
	Failure *Error `json:"failure,omitempty"`
}

// Constants associated with the JobStatus.State property.
// The state of the job.
const (
	JobStatus_State_Canceled  = "canceled"
	JobStatus_State_Completed = "completed"
	JobStatus_State_Failed    = "failed"
	JobStatus_State_Queued    = "queued"
	JobStatus_State_Running   = "running"
)

// UnmarshalJobStatus unmarshals an instance of JobStatus from the specified map of raw messages.
func UnmarshalJobStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobStatus)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "running_at", &obj.RunningAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_at", &obj.CompletedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "message", &obj.Message, UnmarshalJobStatusMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobStatusEntity : Information about the platform job assets related to this execution.
type JobStatusEntity struct {
	// A reference to a resource.
	Deployment *SimpleRel `json:"deployment" validate:"required"`

	// User defined properties.
	Custom interface{} `json:"custom,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityResultHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`

	// The status of the job.
	Scoring *JobScoringResult `json:"scoring,omitempty"`

	// The solve state for a Decision Optimzation job.
	DecisionOptimization *JobDecisionOptimizationResult `json:"decision_optimization,omitempty"`

	// Information about the platform job assets related to this execution.
	// Depending on the `version` date passed, the `platform_jobs` section in the response may or may not be populated.
	// Use the GET call to retrieve the deployment job, this GET call will eventually populate the `platform_jobs` section.
	// Refer to the `version date` description for more details.
	PlatformJob *PlatformJob `json:"platform_job,omitempty"`
}

// UnmarshalJobStatusEntity unmarshals an instance of JobStatusEntity from the specified map of raw messages.
func UnmarshalJobStatusEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobStatusEntity)
	err = core.UnmarshalModel(m, "deployment", &obj.Deployment, UnmarshalSimpleRel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hybrid_pipeline_hardware_specs", &obj.HybridPipelineHardwareSpecs, UnmarshalJobEntityResultHybridPipelineHardwareSpecsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scoring", &obj.Scoring, UnmarshalJobScoringResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "decision_optimization", &obj.DecisionOptimization, UnmarshalJobDecisionOptimizationResult)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "platform_job", &obj.PlatformJob, UnmarshalPlatformJob)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobsResource : The information related to the job.
type JobsResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *ResourceMeta `json:"metadata,omitempty"`

	// Information about the platform job assets related to this execution.
	Entity *JobStatusEntity `json:"entity,omitempty"`
}

// UnmarshalJobsResource unmarshals an instance of JobsResource from the specified map of raw messages.
func UnmarshalJobsResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobsResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalResourceMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalJobStatusEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JobsResources : The information related to the jobs.
type JobsResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *JobsResourcesFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *JobsResourcesNext `json:"next,omitempty"`

	// A list of jobs.
	Resources []JobsResource `json:"resources,omitempty"`
}

// UnmarshalJobsResources unmarshals an instance of JobsResources from the specified map of raw messages.
func UnmarshalJobsResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JobsResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalJobsResourcesFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalJobsResourcesNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalJobsResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on an object, as defined by
// [RFC 6902](https://tools.ietf.org/html/rfc6902).
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add     = "add"
	JSONPatchOperation_Op_Remove  = "remove"
	JSONPatchOperation_Op_Replace = "replace"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
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
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Message : Optional messages related to the deployment.
type Message struct {
	// The level of the message, normally one of `debug`, `info` or `warning`.
	Level *string `json:"level,omitempty"`

	// The message.
	Text *string `json:"text,omitempty"`
}

// UnmarshalMessage unmarshals an instance of Message from the specified map of raw messages.
func UnmarshalMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Message)
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Metric : A metric.
type Metric struct {
	// A timestamp for the metrics.
	Timestamp *strfmt.DateTime `json:"timestamp" validate:"required"`

	// The iteration number.
	Iteration *int64 `json:"iteration,omitempty"`

	// The metrics.
	MlMetrics map[string]float64 `json:"ml_metrics,omitempty"`

	// The metrics from the time series.
	TsMetrics interface{} `json:"ts_metrics,omitempty"`

	// The metrics from federated training.
	MlFederatedMetrics map[string]MlFederatedMetric `json:"ml_federated_metrics,omitempty"`

	// Provides extra information for this training stage in the context of auto-ml.
	Context *MetricsContext `json:"context,omitempty"`
}

// NewMetric : Instantiate Metric (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewMetric(timestamp *strfmt.DateTime) (_model *Metric, err error) {
	_model = &Metric{
		Timestamp: timestamp,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMetric unmarshals an instance of Metric from the specified map of raw messages.
func UnmarshalMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Metric)
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iteration", &obj.Iteration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ml_metrics", &obj.MlMetrics)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ts_metrics", &obj.TsMetrics)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ml_federated_metrics", &obj.MlFederatedMetrics, UnmarshalMlFederatedMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalMetricsContext)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetricsContext : Provides extra information for this training stage in the context of auto-ml.
type MetricsContext struct {
	// The deployment that created the metrics.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// The details of the intermediate model.
	IntermediateModel *IntermediateModel `json:"intermediate_model,omitempty"`

	Phase *string `json:"phase,omitempty"`

	// Details about the step.
	Step *StepInfo `json:"step,omitempty"`

	Classes []string `json:"classes,omitempty"`

	BinaryClassfication *BinaryClassification `json:"binary_classfication,omitempty"`

	MultiClassClassification *MultiClassClassifications `json:"multi_class_classification,omitempty"`

	FeaturesImportance []FeatureImportance `json:"features_importance,omitempty"`

	Schema *string `json:"schema,omitempty"`

	Estimators []string `json:"estimators,omitempty"`
}

// UnmarshalMetricsContext unmarshals an instance of MetricsContext from the specified map of raw messages.
func UnmarshalMetricsContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricsContext)
	err = core.UnmarshalPrimitive(m, "deployment_id", &obj.DeploymentID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "intermediate_model", &obj.IntermediateModel, UnmarshalIntermediateModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "phase", &obj.Phase)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "step", &obj.Step, UnmarshalStepInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "classes", &obj.Classes)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "binary_classfication", &obj.BinaryClassfication, UnmarshalBinaryClassification)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "multi_class_classification", &obj.MultiClassClassification, UnmarshalMultiClassClassifications)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "features_importance", &obj.FeaturesImportance, UnmarshalFeatureImportance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schema", &obj.Schema)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "estimators", &obj.Estimators)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MlFederatedMetric : The metrics from federated training.
type MlFederatedMetric struct {
	// Remote training systems.
	RemoteTrainingSystems []RemoteTrainingSystemMetric `json:"remote_training_systems,omitempty"`

	Global *float64 `json:"global,omitempty"`
}

// UnmarshalMlFederatedMetric unmarshals an instance of MlFederatedMetric from the specified map of raw messages.
func UnmarshalMlFederatedMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MlFederatedMetric)
	err = core.UnmarshalModel(m, "remote_training_systems", &obj.RemoteTrainingSystems, UnmarshalRemoteTrainingSystemMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "global", &obj.Global)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelDefinitionID : The model definition.
type ModelDefinitionID struct {
	// The id of the model definition.
	ID *string `json:"id,omitempty"`
}

// UnmarshalModelDefinitionID unmarshals an instance of ModelDefinitionID from the specified map of raw messages.
func UnmarshalModelDefinitionID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionID)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelDefinitionRel : A model. The `software_spec` is a reference to a software specification. The `hardware_spec` is a reference to a
// hardware specification.
type ModelDefinitionRel struct {
	// The id of the referenced resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the referenced resource.
	Rev *string `json:"rev,omitempty"`

	// The underlying model type produced by the pipeline or by the model_definition.
	ModelType *string `json:"model_type,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// A software specification.
	SoftwareSpec *SoftwareSpecRel `json:"software_spec,omitempty"`

	// If present, it overrides the command specified to the library resource itself.
	Command *string `json:"command,omitempty"`

	// Optional key-value pairs parameters.
	Parameters interface{} `json:"parameters,omitempty"`
}

// NewModelDefinitionRel : Instantiate ModelDefinitionRel (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewModelDefinitionRel(id string) (_model *ModelDefinitionRel, err error) {
	_model = &ModelDefinitionRel{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalModelDefinitionRel unmarshals an instance of ModelDefinitionRel from the specified map of raw messages.
func UnmarshalModelDefinitionRel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionRel)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_type", &obj.ModelType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "software_spec", &obj.SoftwareSpec, UnmarshalSoftwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "command", &obj.Command)
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

// ModelDefinitionResource : The information for a model definition.
type ModelDefinitionResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *ModelDefinitionResourceMetadata `json:"metadata" validate:"required"`

	Entity *ModelDefinitionResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalModelDefinitionResource unmarshals an instance of ModelDefinitionResource from the specified map of raw messages.
func UnmarshalModelDefinitionResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalModelDefinitionResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalModelDefinitionResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelDefinitionResources : System details.
type ModelDefinitionResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	// A list of model definitions.
	Resources []ModelDefinitionResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalModelDefinitionResources unmarshals an instance of ModelDefinitionResources from the specified map of raw messages.
func UnmarshalModelDefinitionResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelDefinitionResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalModelDefinitionResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ModelDefinitionResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// ModelLocation : The location of the intermediate model.
type ModelLocation struct {
	// The generated pipeline at this stage.
	Pipeline *string `json:"pipeline,omitempty"`

	// The generated pipeline model.
	PipelineModel *string `json:"pipeline_model,omitempty"`

	// The generated model at this stage.
	Model *string `json:"model,omitempty"`
}

// UnmarshalModelLocation unmarshals an instance of ModelLocation from the specified map of raw messages.
func UnmarshalModelLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelLocation)
	err = core.UnmarshalPrimitive(m, "pipeline", &obj.Pipeline)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_model", &obj.PipelineModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model", &obj.Model)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelResource : The information for a model.
type ModelResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *ModelResourceMetadata `json:"metadata" validate:"required"`

	// Information related to the upload of the model content.
	Entity *ModelResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalModelResource unmarshals an instance of ModelResource from the specified map of raw messages.
func UnmarshalModelResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalModelResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalModelResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ModelResources : System details.
type ModelResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	// A list of models.
	Resources []ModelResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalModelResources unmarshals an instance of ModelResources from the specified map of raw messages.
func UnmarshalModelResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ModelResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalModelResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ModelResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// MultiClassClassification : MultiClassClassification struct
type MultiClassClassification struct {
	// The class name.
	Class *string `json:"class" validate:"required"`

	// The location of the confusion_matrix details in associated storage.
	ConfusionMatrixLocation *string `json:"confusion_matrix_location,omitempty"`

	// The confusion matrix for the selected class.
	ConfusionMatrix *ConfusionMatrix `json:"confusion_matrix" validate:"required"`

	// The location of the roc (receiver operating characteristic) curve details in the associated storage.
	RocCurveLocation *string `json:"roc_curve_location,omitempty"`

	// The roc (receiver operating characteristic) curve for the selected class.
	RocCurve *RocCurve `json:"roc_curve" validate:"required"`
}

// NewMultiClassClassification : Instantiate MultiClassClassification (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewMultiClassClassification(class string, confusionMatrix *ConfusionMatrix, rocCurve *RocCurve) (_model *MultiClassClassification, err error) {
	_model = &MultiClassClassification{
		Class:           core.StringPtr(class),
		ConfusionMatrix: confusionMatrix,
		RocCurve:        rocCurve,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMultiClassClassification unmarshals an instance of MultiClassClassification from the specified map of raw messages.
func UnmarshalMultiClassClassification(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MultiClassClassification)
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confusion_matrix_location", &obj.ConfusionMatrixLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "confusion_matrix", &obj.ConfusionMatrix, UnmarshalConfusionMatrix)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "roc_curve_location", &obj.RocCurveLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "roc_curve", &obj.RocCurve, UnmarshalRocCurve)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MultiClassClassifications : MultiClassClassifications struct
type MultiClassClassifications struct {
	// The classifications details for each class.
	OneVsAll []MultiClassClassification `json:"one_vs_all,omitempty"`

	// The location of the classifications details in associated storage.
	OneVsAllLocation *string `json:"one_vs_all_location,omitempty"`
}

// UnmarshalMultiClassClassifications unmarshals an instance of MultiClassClassifications from the specified map of raw messages.
func UnmarshalMultiClassClassifications(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MultiClassClassifications)
	err = core.UnmarshalModel(m, "one_vs_all", &obj.OneVsAll, UnmarshalMultiClassClassification)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "one_vs_all_location", &obj.OneVsAllLocation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectLocation : A reference to data.
type ObjectLocation struct {
	// Item identification inside a collection.
	ID *string `json:"id,omitempty"`

	// The data source type like `connection_asset` or `data_asset`.
	Type *string `json:"type" validate:"required"`

	// Contains a set of fields specific to each connection. See here for [details about specifying
	// connections](#datareferences).
	Connection interface{} `json:"connection,omitempty"`

	// Contains a set of fields specific to each connection.
	Location map[string]string `json:"location" validate:"required"`
}

// Constants associated with the ObjectLocation.Type property.
// The data source type like `connection_asset` or `data_asset`.
const (
	ObjectLocation_Type_ConnectionAsset = "connection_asset"
	ObjectLocation_Type_DataAsset       = "data_asset"
	ObjectLocation_Type_Fs              = "fs"
	ObjectLocation_Type_URL             = "url"
)

// NewObjectLocation : Instantiate ObjectLocation (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewObjectLocation(typeVar string, location map[string]string) (_model *ObjectLocation, err error) {
	_model = &ObjectLocation{
		Type:     core.StringPtr(typeVar),
		Location: location,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalObjectLocation unmarshals an instance of ObjectLocation from the specified map of raw messages.
func UnmarshalObjectLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectLocation)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection", &obj.Connection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ObjectLocationOptim : A reference to data.
type ObjectLocationOptim struct {
	// Item identification inside a collection.
	ID *string `json:"id" validate:"required"`

	// The data source type like `connection_asset` or `data_asset`.
	Type *string `json:"type" validate:"required"`

	// Contains a set of fields specific to each connection. See here for [details about specifying
	// connections](#datareferences).
	Connection interface{} `json:"connection,omitempty"`

	// Contains a set of fields that describe the location of the data with respect to the `connection`.
	Location map[string]string `json:"location" validate:"required"`
}

// Constants associated with the ObjectLocationOptim.Type property.
// The data source type like `connection_asset` or `data_asset`.
const (
	ObjectLocationOptim_Type_ConnectionAsset = "connection_asset"
	ObjectLocationOptim_Type_DataAsset       = "data_asset"
	ObjectLocationOptim_Type_Fs              = "fs"
	ObjectLocationOptim_Type_URL             = "url"
)

// NewObjectLocationOptim : Instantiate ObjectLocationOptim (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewObjectLocationOptim(id string, typeVar string, location map[string]string) (_model *ObjectLocationOptim, err error) {
	_model = &ObjectLocationOptim{
		ID:       core.StringPtr(id),
		Type:     core.StringPtr(typeVar),
		Location: location,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalObjectLocationOptim unmarshals an instance of ObjectLocationOptim from the specified map of raw messages.
func UnmarshalObjectLocationOptim(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ObjectLocationOptim)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "connection", &obj.Connection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Organization : A remote organization.
type Organization struct {
	// The name of the organization.
	Name *string `json:"name" validate:"required"`

	// The region for the organization.
	Region *string `json:"region,omitempty"`
}

// NewOrganization : Instantiate Organization (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewOrganization(name string) (_model *Organization, err error) {
	_model = &Organization{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalOrganization unmarshals an instance of Organization from the specified map of raw messages.
func UnmarshalOrganization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Organization)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelineRel : A pipeline. The `hardware_spec` is a reference to the hardware specification. The `hybrid_pipeline_hardware_specs`
// are used only when training a hybrid pipeline in order to specify compute requirement for each pipeline node.
type PipelineRel struct {
	// The id of the referenced resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the referenced resource.
	Rev *string `json:"rev,omitempty"`

	// The underlying model type produced by the pipeline or by the model_definition.
	ModelType *string `json:"model_type,omitempty"`

	// The data bindings.
	DataBindings []PipelineRelDataBindingsItem `json:"data_bindings,omitempty"`

	// The node parameters.
	NodesParameters []PipelineRelNodesParametersItem `json:"nodes_parameters,omitempty"`

	// A hardware specification.
	HardwareSpec *HardwareSpecRel `json:"hardware_spec,omitempty"`

	// Hybrid pipeline hardware specification.
	HybridPipelineHardwareSpecs []JobEntityResultHybridPipelineHardwareSpecsItem `json:"hybrid_pipeline_hardware_specs,omitempty"`
}

// NewPipelineRel : Instantiate PipelineRel (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewPipelineRel(id string) (_model *PipelineRel, err error) {
	_model = &PipelineRel{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPipelineRel unmarshals an instance of PipelineRel from the specified map of raw messages.
func UnmarshalPipelineRel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineRel)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "model_type", &obj.ModelType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data_bindings", &obj.DataBindings, UnmarshalPipelineRelDataBindingsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "nodes_parameters", &obj.NodesParameters, UnmarshalPipelineRelNodesParametersItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hardware_spec", &obj.HardwareSpec, UnmarshalHardwareSpecRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hybrid_pipeline_hardware_specs", &obj.HybridPipelineHardwareSpecs, UnmarshalJobEntityResultHybridPipelineHardwareSpecsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelineResource : The information for a pipeline.
type PipelineResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *PipelineResourceMetadata `json:"metadata" validate:"required"`

	// The details of the pipeline to be created.
	Entity *PipelineResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalPipelineResource unmarshals an instance of PipelineResource from the specified map of raw messages.
func UnmarshalPipelineResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalPipelineResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalPipelineResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PipelineResources : System details.
type PipelineResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *PipelineResourcesFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *PipelineResourcesNext `json:"next,omitempty"`

	// A list of pipelines.
	Resources []PipelineResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalPipelineResources unmarshals an instance of PipelineResources from the specified map of raw messages.
func UnmarshalPipelineResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PipelineResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPipelineResourcesFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPipelineResourcesNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalPipelineResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *PipelineResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// PlatformJob : Information about the platform job assets related to this execution. Depending on the `version` date passed, the
// `platform_jobs` section in the response may or may not be populated. Use the GET call to retrieve the deployment job,
// this GET call will eventually populate the `platform_jobs` section. Refer to the `version date` description for more
// details.
type PlatformJob struct {
	// The id of the platform job.
	JobID *string `json:"job_id" validate:"required"`

	// The run id of the platform job.
	RunID *string `json:"run_id" validate:"required"`
}

// UnmarshalPlatformJob unmarshals an instance of PlatformJob from the specified map of raw messages.
func UnmarshalPlatformJob(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PlatformJob)
	err = core.UnmarshalPrimitive(m, "job_id", &obj.JobID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "run_id", &obj.RunID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rel : A reference to a resource.
type Rel struct {
	// The id of the referenced resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the referenced resource.
	Rev *string `json:"rev,omitempty"`
}

// NewRel : Instantiate Rel (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewRel(id string) (_model *Rel, err error) {
	_model = &Rel{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRel unmarshals an instance of Rel from the specified map of raw messages.
func UnmarshalRel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rel)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoteAdmin : The details of the remote administrator for the organization and identities.
type RemoteAdmin struct {
	// The name of the remote administrator.
	Name *string `json:"name" validate:"required"`

	// The email of the remote administrator.
	Email *string `json:"email,omitempty"`
}

// NewRemoteAdmin : Instantiate RemoteAdmin (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewRemoteAdmin(name string) (_model *RemoteAdmin, err error) {
	_model = &RemoteAdmin{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRemoteAdmin unmarshals an instance of RemoteAdmin from the specified map of raw messages.
func UnmarshalRemoteAdmin(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoteAdmin)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoteTrainingSystemMetric : The remote training system metric.
type RemoteTrainingSystemMetric struct {
	ID *string `json:"id" validate:"required"`

	Local *float64 `json:"local,omitempty"`

	Fused *float64 `json:"fused,omitempty"`
}

// NewRemoteTrainingSystemMetric : Instantiate RemoteTrainingSystemMetric (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewRemoteTrainingSystemMetric(id string) (_model *RemoteTrainingSystemMetric, err error) {
	_model = &RemoteTrainingSystemMetric{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRemoteTrainingSystemMetric unmarshals an instance of RemoteTrainingSystemMetric from the specified map of raw messages.
func UnmarshalRemoteTrainingSystemMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoteTrainingSystemMetric)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "local", &obj.Local)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fused", &obj.Fused)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoteTrainingSystemPatchHelper : RemoteTrainingSystemPatchHelper struct
type RemoteTrainingSystemPatchHelper struct {
	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// A remote organization.
	Organization *Organization `json:"organization,omitempty"`

	// The list of allowed identities that are allowed to access the remote system.
	AllowedIdentities []AllowedIdentity `json:"allowed_identities,omitempty"`

	// The details of the remote administrator for the organization and identities.
	RemoteAdmin *RemoteAdmin `json:"remote_admin,omitempty"`
}

// UnmarshalRemoteTrainingSystemPatchHelper unmarshals an instance of RemoteTrainingSystemPatchHelper from the specified map of raw messages.
func UnmarshalRemoteTrainingSystemPatchHelper(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoteTrainingSystemPatchHelper)
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "organization", &obj.Organization, UnmarshalOrganization)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "allowed_identities", &obj.AllowedIdentities, UnmarshalAllowedIdentity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "remote_admin", &obj.RemoteAdmin, UnmarshalRemoteAdmin)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonMachineLearningV4) NewRemoteTrainingSystemPatchHelperPatch(remoteTrainingSystemPatchHelper *RemoteTrainingSystemPatchHelper) (_patch []JSONPatchOperation) {
	if remoteTrainingSystemPatchHelper.Tags != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/tags"),
			Value: remoteTrainingSystemPatchHelper.Tags,
		})
	}
	if remoteTrainingSystemPatchHelper.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: remoteTrainingSystemPatchHelper.Name,
		})
	}
	if remoteTrainingSystemPatchHelper.Description != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/description"),
			Value: remoteTrainingSystemPatchHelper.Description,
		})
	}
	if remoteTrainingSystemPatchHelper.Custom != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/custom"),
			Value: remoteTrainingSystemPatchHelper.Custom,
		})
	}
	if remoteTrainingSystemPatchHelper.Organization != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/organization"),
			Value: remoteTrainingSystemPatchHelper.Organization,
		})
	}
	if remoteTrainingSystemPatchHelper.AllowedIdentities != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/allowed_identities"),
			Value: remoteTrainingSystemPatchHelper.AllowedIdentities,
		})
	}
	if remoteTrainingSystemPatchHelper.RemoteAdmin != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/remote_admin"),
			Value: remoteTrainingSystemPatchHelper.RemoteAdmin,
		})
	}
	return
}

// RemoteTrainingSystemResource : The information for a remote training system.
type RemoteTrainingSystemResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *RemoteTrainingSystemResourceMetadata `json:"metadata" validate:"required"`

	// The definition of a remote system used by Federated Learning.
	Entity *RemoteTrainingSystemResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalRemoteTrainingSystemResource unmarshals an instance of RemoteTrainingSystemResource from the specified map of raw messages.
func UnmarshalRemoteTrainingSystemResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoteTrainingSystemResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalRemoteTrainingSystemResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalRemoteTrainingSystemResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoteTrainingSystemResources : System details.
type RemoteTrainingSystemResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	// A list of remote training systems.
	Resources []RemoteTrainingSystemResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalRemoteTrainingSystemResources unmarshals an instance of RemoteTrainingSystemResources from the specified map of raw messages.
func UnmarshalRemoteTrainingSystemResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RemoteTrainingSystemResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalRemoteTrainingSystemResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *RemoteTrainingSystemResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// ResourceMeta : Common metadata for a resource where `project_id` or `space_id` must be present.
type ResourceMeta struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`

	// The space that contains the resource. Either `space_id` or `project_id` has to be given.
	SpaceID *string `json:"space_id,omitempty"`

	// The project that contains the resource. Either `space_id` or `project_id` has to be given.
	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalResourceMeta unmarshals an instance of ResourceMeta from the specified map of raw messages.
func UnmarshalResourceMeta(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceMeta)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
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

// ResourceMetaBase : Common metadata for a resource.
type ResourceMetaBase struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The revision of the resource.
	Rev *string `json:"rev,omitempty"`

	// The user id which created this resource.
	Owner *string `json:"owner,omitempty"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was last modified.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// The id of the parent resource where applicable.
	ParentID *string `json:"parent_id,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// Information related to the revision.
	CommitInfo *ResourceMetaBaseCommitInfo `json:"commit_info,omitempty"`
}

// UnmarshalResourceMetaBase unmarshals an instance of ResourceMetaBase from the specified map of raw messages.
func UnmarshalResourceMetaBase(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceMetaBase)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parent_id", &obj.ParentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalModel(m, "commit_info", &obj.CommitInfo, UnmarshalResourceMetaBaseCommitInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RocCurve : The roc (receiver operating characteristic) curve for the selected class.
type RocCurve struct {
	TrueClass *string `json:"true_class" validate:"required"`

	// The true posistive rates.
	Tpr []float64 `json:"tpr" validate:"required"`

	// The false posistive rates.
	Fpr []float64 `json:"fpr" validate:"required"`

	// The thresholds.
	Thresholds []float64 `json:"thresholds,omitempty"`
}

// NewRocCurve : Instantiate RocCurve (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewRocCurve(trueClass string, tpr []float64, fpr []float64) (_model *RocCurve, err error) {
	_model = &RocCurve{
		TrueClass: core.StringPtr(trueClass),
		Tpr:       tpr,
		Fpr:       fpr,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRocCurve unmarshals an instance of RocCurve from the specified map of raw messages.
func UnmarshalRocCurve(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RocCurve)
	err = core.UnmarshalPrimitive(m, "true_class", &obj.TrueClass)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tpr", &obj.Tpr)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fpr", &obj.Fpr)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "thresholds", &obj.Thresholds)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringPayload : The payload for scoring.
type ScoringPayload struct {
	// Discriminates the data for multi input data situation. For example in cases where multiple tensors are expected.
	ID *string `json:"id,omitempty"`

	// If specified, the `values` represents the ground truth data (the label information) for the input data provided.
	// This information will be used for computing machine learning metrics.
	Type *string `json:"type,omitempty"`

	// The names of the fields. The order of fields values must be consistent with the order of fields names.
	Fields []string `json:"fields,omitempty"`

	// Input data as a vector for a single record or a matrix representing a mini batch of records.
	Values [][]interface{} `json:"values,omitempty"`

	// Used when performing evaluation. This contains the groud truths for the input data.
	Targets [][]interface{} `json:"targets,omitempty"`
}

// Constants associated with the ScoringPayload.Type property.
// If specified, the `values` represents the ground truth data (the label information) for the input data provided. This
// information will be used for computing machine learning metrics.
const (
	ScoringPayload_Type_Target = "target"
)

// UnmarshalScoringPayload unmarshals an instance of ScoringPayload from the specified map of raw messages.
func UnmarshalScoringPayload(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringPayload)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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
	err = core.UnmarshalPrimitive(m, "targets", &obj.Targets)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScoringPayloadOptim : The payload for scoring.
type ScoringPayloadOptim struct {
	// Discriminates the data for multi input data situation. For example in cases where multiple tensors are expected.
	ID *string `json:"id" validate:"required"`

	// The names of the fields. The order of fields values must be consistent with the order of fields names. Mutually
	// exclusive with `content` field.
	Fields []string `json:"fields,omitempty"`

	// Input data as a vector for a single record or a matrix representing a mini batch of records. Mutually exclusive with
	// `content` field.
	Values [][]interface{} `json:"values,omitempty"`

	// Input data as a base64 encoded string. Mutually exclusive with `fields` or `values` field.
	Content *string `json:"content,omitempty"`
}

// NewScoringPayloadOptim : Instantiate ScoringPayloadOptim (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewScoringPayloadOptim(id string) (_model *ScoringPayloadOptim, err error) {
	_model = &ScoringPayloadOptim{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalScoringPayloadOptim unmarshals an instance of ScoringPayloadOptim from the specified map of raw messages.
func UnmarshalScoringPayloadOptim(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScoringPayloadOptim)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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
	err = core.UnmarshalPrimitive(m, "content", &obj.Content)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SimpleRel : A reference to a resource.
type SimpleRel struct {
	// The id of the referenced resource.
	ID *string `json:"id" validate:"required"`
}

// NewSimpleRel : Instantiate SimpleRel (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewSimpleRel(id string) (_model *SimpleRel, err error) {
	_model = &SimpleRel{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSimpleRel unmarshals an instance of SimpleRel from the specified map of raw messages.
func UnmarshalSimpleRel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SimpleRel)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SoftwareSpecRel : A software specification.
type SoftwareSpecRel struct {
	// The id of the software specification. One, and only one, of `id` or `name` must be set.
	ID *string `json:"id,omitempty"`

	// The revision of the software specification.
	Rev *string `json:"rev,omitempty"`

	// The name of the software specification. One, and only one, of `id` or `name` must be set.
	Name *string `json:"name,omitempty"`
}

// UnmarshalSoftwareSpecRel unmarshals an instance of SoftwareSpecRel from the specified map of raw messages.
func UnmarshalSoftwareSpecRel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SoftwareSpecRel)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rev", &obj.Rev)
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

// SolveState : The solve state for a Decision Optimzation job.
type SolveState struct {
	// Details related to the job.
	Details interface{} `json:"details,omitempty"`

	// The solve status.
	SolveStatus *string `json:"solve_status,omitempty"`

	// The interrupted status.
	InterruptionStatus *string `json:"interruption_status,omitempty"`

	// The latest engine activity.
	LatestEngineActivity []string `json:"latest_engine_activity,omitempty"`
}

// Constants associated with the SolveState.SolveStatus property.
// The solve status.
const (
	SolveState_SolveStatus_FeasibleSolution              = "feasible_solution"
	SolveState_SolveStatus_InfeasibleOrUnboundedSolution = "infeasible_or_unbounded_solution"
	SolveState_SolveStatus_InfeasibleSolution            = "infeasible_solution"
	SolveState_SolveStatus_OptimalSolution               = "optimal_solution"
	SolveState_SolveStatus_UnboundedSolution             = "unbounded_solution"
	SolveState_SolveStatus_Unknown                       = "unknown"
)

// Constants associated with the SolveState.InterruptionStatus property.
// The interrupted status.
const (
	SolveState_InterruptionStatus_EngineLimit = "engine_limit"
	SolveState_InterruptionStatus_Kill        = "kill"
	SolveState_InterruptionStatus_OutOfMemory = "out_of_memory"
	SolveState_InterruptionStatus_Stop        = "stop"
	SolveState_InterruptionStatus_Timeout     = "timeout"
)

// UnmarshalSolveState unmarshals an instance of SolveState from the specified map of raw messages.
func UnmarshalSolveState(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SolveState)
	err = core.UnmarshalPrimitive(m, "details", &obj.Details)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "solve_status", &obj.SolveStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interruption_status", &obj.InterruptionStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "latest_engine_activity", &obj.LatestEngineActivity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Stats : The stats about deployments for a space.
type Stats struct {
	// An `id` associated with the space.
	SpaceID *string `json:"space_id,omitempty"`

	// The total number of deployments created in a space including online, rshiny and batch.
	TotalCount *float64 `json:"total_count,omitempty"`

	// The number of online deployments created in a space.
	OnlineCount *float64 `json:"online_count,omitempty"`

	// The number of batch deployments created in a space.
	BatchCount *float64 `json:"batch_count,omitempty"`

	// The number of rshiny deployments created in a space.x.
	RshinyCount *float64 `json:"rshiny_count,omitempty"`
}

// UnmarshalStats unmarshals an instance of Stats from the specified map of raw messages.
func UnmarshalStats(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Stats)
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "online_count", &obj.OnlineCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "batch_count", &obj.BatchCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rshiny_count", &obj.RshinyCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StepInfo : Details about the step.
type StepInfo struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name" validate:"required"`

	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`

	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	HyperParameters interface{} `json:"hyper_parameters,omitempty"`

	DataAllocation *int64 `json:"data_allocation,omitempty"`

	Estimator *string `json:"estimator,omitempty"`

	Transformer *string `json:"transformer,omitempty"`

	Score *float64 `json:"score,omitempty"`
}

// NewStepInfo : Instantiate StepInfo (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewStepInfo(name string) (_model *StepInfo, err error) {
	_model = &StepInfo{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalStepInfo unmarshals an instance of StepInfo from the specified map of raw messages.
func UnmarshalStepInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StepInfo)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_at", &obj.StartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_at", &obj.CompletedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hyper_parameters", &obj.HyperParameters)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_allocation", &obj.DataAllocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "estimator", &obj.Estimator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "transformer", &obj.Transformer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyncScoringData : Scoring data.
type SyncScoringData struct {
	// The input data.
	InputData []InputDataArray `json:"input_data" validate:"required"`
}

// NewSyncScoringData : Instantiate SyncScoringData (Generic Model Constructor)
func (*WatsonMachineLearningV4) NewSyncScoringData(inputData []InputDataArray) (_model *SyncScoringData, err error) {
	_model = &SyncScoringData{
		InputData: inputData,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSyncScoringData unmarshals an instance of SyncScoringData from the specified map of raw messages.
func UnmarshalSyncScoringData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyncScoringData)
	err = core.UnmarshalModel(m, "input_data", &obj.InputData, UnmarshalInputDataArray)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyncScoringDataResults : Scoring results.
type SyncScoringDataResults struct {
	// The predictions.
	Predictions []ScoringPayload `json:"predictions,omitempty"`
}

// UnmarshalSyncScoringDataResults unmarshals an instance of SyncScoringDataResults from the specified map of raw messages.
func UnmarshalSyncScoringDataResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyncScoringDataResults)
	err = core.UnmarshalModel(m, "predictions", &obj.Predictions, UnmarshalScoringPayload)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SystemDetails : Optional details coming from the service and related to the API call or the associated resource.
type SystemDetails struct {
	// Any warnings coming from the system.
	Warnings []Warning `json:"warnings,omitempty"`
}

// UnmarshalSystemDetails unmarshals an instance of SystemDetails from the specified map of raw messages.
func UnmarshalSystemDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SystemDetails)
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalWarning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDefinitionPatchHelper : TrainingDefinitionPatchHelper struct
type TrainingDefinitionPatchHelper struct {
	// A list of tags for this resource.
	Tags []string `json:"tags,omitempty"`

	// The name of the resource.
	Name *string `json:"name,omitempty"`

	// A description of the resource.
	Description *string `json:"description,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Federated Learning is a Technical Preview.
	FederatedLearning *FederatedLearning `json:"federated_learning,omitempty"`
}

// UnmarshalTrainingDefinitionPatchHelper unmarshals an instance of TrainingDefinitionPatchHelper from the specified map of raw messages.
func UnmarshalTrainingDefinitionPatchHelper(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDefinitionPatchHelper)
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		//return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "federated_learning", &obj.FederatedLearning, UnmarshalFederatedLearning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*WatsonMachineLearningV4) NewTrainingDefinitionPatchHelperPatch(trainingDefinitionPatchHelper *TrainingDefinitionPatchHelper) (_patch []JSONPatchOperation) {
	if trainingDefinitionPatchHelper.Tags != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/tags"),
			Value: trainingDefinitionPatchHelper.Tags,
		})
	}
	if trainingDefinitionPatchHelper.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: trainingDefinitionPatchHelper.Name,
		})
	}
	if trainingDefinitionPatchHelper.Description != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/description"),
			Value: trainingDefinitionPatchHelper.Description,
		})
	}
	if trainingDefinitionPatchHelper.Custom != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/custom"),
			Value: trainingDefinitionPatchHelper.Custom,
		})
	}
	if trainingDefinitionPatchHelper.FederatedLearning != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/federated_learning"),
			Value: trainingDefinitionPatchHelper.FederatedLearning,
		})
	}
	return
}

// TrainingDefinitionResource : The information for a training definition.
type TrainingDefinitionResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *TrainingDefinitionResourceMetadata `json:"metadata" validate:"required"`

	// The `training_data_references` contain the training datasets and the
	// `results_reference` the connection where results will be stored.
	Entity *TrainingDefinitionResourceEntity `json:"entity" validate:"required"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalTrainingDefinitionResource unmarshals an instance of TrainingDefinitionResource from the specified map of raw messages.
func UnmarshalTrainingDefinitionResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDefinitionResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalTrainingDefinitionResourceMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalTrainingDefinitionResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingDefinitionResources : System details.
type TrainingDefinitionResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *TrainingDefinitionResourcesFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *TrainingDefinitionResourcesNext `json:"next,omitempty"`

	// A list of training definitions.
	Resources []TrainingDefinitionResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalTrainingDefinitionResources unmarshals an instance of TrainingDefinitionResources from the specified map of raw messages.
func UnmarshalTrainingDefinitionResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingDefinitionResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalTrainingDefinitionResourcesFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalTrainingDefinitionResourcesNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalTrainingDefinitionResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *TrainingDefinitionResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// TrainingReference : The `pipeline` is a reference to the pipeline. The `model_definition` is the library reference that contains the
// training library.
type TrainingReference struct {
	// A pipeline.
	// The `hardware_spec` is a reference to the hardware specification.
	// The `hybrid_pipeline_hardware_specs` are used only when training a hybrid pipeline in order to
	// specify compute requirement for each pipeline node.
	Pipeline *PipelineRel `json:"pipeline,omitempty"`

	// The model definition.
	ModelDefinition *ModelDefinitionID `json:"model_definition,omitempty"`

	// The hyper parameters used in the experiment.
	HyperParametersOptimization *TrainingReferenceHyperParametersOptimization `json:"hyper_parameters_optimization,omitempty"`
}

// UnmarshalTrainingReference unmarshals an instance of TrainingReference from the specified map of raw messages.
func UnmarshalTrainingReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingReference)
	err = core.UnmarshalModel(m, "pipeline", &obj.Pipeline, UnmarshalPipelineRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "model_definition", &obj.ModelDefinition, UnmarshalModelDefinitionID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hyper_parameters_optimization", &obj.HyperParametersOptimization, UnmarshalTrainingReferenceHyperParametersOptimization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingResource : Training resource.
type TrainingResource struct {
	// Common metadata for a resource where `project_id` or `space_id` must be present.
	Metadata *ResourceMeta `json:"metadata,omitempty"`

	// The `training_data_references` contain the training datasets and the
	// `results_reference` the connection where results will be stored.
	Entity *TrainingResourceEntity `json:"entity,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalTrainingResource unmarshals an instance of TrainingResource from the specified map of raw messages.
func UnmarshalTrainingResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalResourceMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalTrainingResourceEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingResourceEntity : The `training_data_references` contain the training datasets and the
// `results_reference` the connection where results will be stored.
type TrainingResourceEntity struct {
	// A reference to a resource.
	Experiment *Rel `json:"experiment,omitempty"`

	// A pipeline.
	// The `hardware_spec` is a reference to the hardware specification.
	// The `hybrid_pipeline_hardware_specs` are used only when training a hybrid pipeline in order to
	// specify compute requirement for each pipeline node.
	Pipeline *PipelineRel `json:"pipeline,omitempty"`

	// A model.
	// The `software_spec` is a reference to a software specification.
	// The `hardware_spec` is a reference to a hardware specification.
	ModelDefinition *ModelDefinitionRel `json:"model_definition,omitempty"`

	// Federated Learning is a Technical Preview.
	FederatedLearning *FederatedLearning `json:"federated_learning,omitempty"`

	// Training datasets.
	TrainingDataReferences []DataConnectionReference `json:"training_data_references,omitempty"`

	// The training results.
	ResultsReference *ObjectLocation `json:"results_reference" validate:"required"`

	// The holdout/test datasets.
	TestDataReferences []DataConnectionReference `json:"test_data_references,omitempty"`

	// User defined properties specified as key-value pairs.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Status of the model.
	Status *TrainingStatus `json:"status,omitempty"`
}

// UnmarshalTrainingResourceEntity unmarshals an instance of TrainingResourceEntity from the specified map of raw messages.
func UnmarshalTrainingResourceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingResourceEntity)
	err = core.UnmarshalModel(m, "experiment", &obj.Experiment, UnmarshalRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pipeline", &obj.Pipeline, UnmarshalPipelineRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "model_definition", &obj.ModelDefinition, UnmarshalModelDefinitionRel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "federated_learning", &obj.FederatedLearning, UnmarshalFederatedLearning)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "training_data_references", &obj.TrainingDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "results_reference", &obj.ResultsReference, UnmarshalObjectLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "test_data_references", &obj.TestDataReferences, UnmarshalDataConnectionReference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "custom", &obj.Custom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalTrainingStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TrainingResources : Information for paging when querying resources.
type TrainingResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit" validate:"required"`

	// The reference to the first item in the current page.
	First *TrainingResourcesFirst `json:"first" validate:"required"`

	// A reference to the first item of the next page, if any.
	Next *TrainingResourcesNext `json:"next,omitempty"`

	Resources []TrainingResource `json:"resources,omitempty"`

	// Optional details coming from the service and related to the API call or the associated resource.
	System *SystemDetails `json:"system,omitempty"`
}

// UnmarshalTrainingResources unmarshals an instance of TrainingResources from the specified map of raw messages.
func UnmarshalTrainingResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingResources)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalTrainingResourcesFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalTrainingResourcesNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalTrainingResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system", &obj.System, UnmarshalSystemDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *TrainingResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// TrainingStatus : Status of the model.
type TrainingStatus struct {
	// Date and Time in which current training state has started.
	RunningAt *strfmt.DateTime `json:"running_at,omitempty"`

	// Date and Time in which training had completed.
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	// Current training iteration.
	Iteration *float64 `json:"iteration,omitempty"`

	// Total number of iterations training must complete.
	TotalIterations *float64 `json:"total_iterations,omitempty"`

	// Current state of training.
	State *string `json:"state" validate:"required"`

	// Compute usage metrics.
	ComputeUsageMetrics *ComputeUsageMetrics `json:"compute_usage_metrics,omitempty"`

	// Hyperparameter optimization.
	Hpo *TrainingStatusHpo `json:"hpo,omitempty"`

	// Federated learning info.
	FederatedLearningInfo *FederatedLearningInfo `json:"federated_learning_info,omitempty"`

	// Message.
	Message *TrainingStatusMessage `json:"message,omitempty"`

	// Metrics that can be returned by an operation.
	Metrics []Metric `json:"metrics,omitempty"`

	// The data returned when an error is encountered.
	Failure *Error `json:"failure,omitempty"`
}

// Constants associated with the TrainingStatus.State property.
// Current state of training.
const (
	TrainingStatus_State_Canceled  = "canceled"
	TrainingStatus_State_Completed = "completed"
	TrainingStatus_State_Failed    = "failed"
	TrainingStatus_State_Pending   = "pending"
	TrainingStatus_State_Queued    = "queued"
	TrainingStatus_State_Running   = "running"
	TrainingStatus_State_Storing   = "storing"
)

// UnmarshalTrainingStatus unmarshals an instance of TrainingStatus from the specified map of raw messages.
func UnmarshalTrainingStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TrainingStatus)
	err = core.UnmarshalPrimitive(m, "running_at", &obj.RunningAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "completed_at", &obj.CompletedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iteration", &obj.Iteration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_iterations", &obj.TotalIterations)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "compute_usage_metrics", &obj.ComputeUsageMetrics, UnmarshalComputeUsageMetrics)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hpo", &obj.Hpo, UnmarshalTrainingStatusHpo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "federated_learning_info", &obj.FederatedLearningInfo, UnmarshalFederatedLearningInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "message", &obj.Message, UnmarshalTrainingStatusMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metrics", &obj.Metrics, UnmarshalMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failure", &obj.Failure, UnmarshalError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Warning : A warning message.
type Warning struct {
	// The message.
	Message *string `json:"message" validate:"required"`

	// An `id` associated with the message.
	ID *string `json:"id,omitempty"`
}

// UnmarshalWarning unmarshals an instance of Warning from the specified map of raw messages.
func UnmarshalWarning(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Warning)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
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

// DeploymentJobDefinitionsListPager can be used to simplify the use of the "DeploymentJobDefinitionsList" method.
type DeploymentJobDefinitionsListPager struct {
	hasNext     bool
	options     *DeploymentJobDefinitionsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewDeploymentJobDefinitionsListPager returns a new DeploymentJobDefinitionsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewDeploymentJobDefinitionsListPager(options *DeploymentJobDefinitionsListOptions) (pager *DeploymentJobDefinitionsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy DeploymentJobDefinitionsListOptions = *options
	pager = &DeploymentJobDefinitionsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *DeploymentJobDefinitionsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *DeploymentJobDefinitionsListPager) GetNextWithContext(ctx context.Context) (page []JobResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.DeploymentJobDefinitionsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *DeploymentJobDefinitionsListPager) GetAllWithContext(ctx context.Context) (allItems []JobResource, err error) {
	for pager.HasNext() {
		var nextPage []JobResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *DeploymentJobDefinitionsListPager) GetNext() (page []JobResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DeploymentJobDefinitionsListPager) GetAll() (allItems []JobResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// DeploymentJobDefinitionsListRevisionsPager can be used to simplify the use of the "DeploymentJobDefinitionsListRevisions" method.
type DeploymentJobDefinitionsListRevisionsPager struct {
	hasNext     bool
	options     *DeploymentJobDefinitionsListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewDeploymentJobDefinitionsListRevisionsPager returns a new DeploymentJobDefinitionsListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewDeploymentJobDefinitionsListRevisionsPager(options *DeploymentJobDefinitionsListRevisionsOptions) (pager *DeploymentJobDefinitionsListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy DeploymentJobDefinitionsListRevisionsOptions = *options
	pager = &DeploymentJobDefinitionsListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *DeploymentJobDefinitionsListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *DeploymentJobDefinitionsListRevisionsPager) GetNextWithContext(ctx context.Context) (page []JobResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.DeploymentJobDefinitionsListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *DeploymentJobDefinitionsListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []JobResource, err error) {
	for pager.HasNext() {
		var nextPage []JobResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *DeploymentJobDefinitionsListRevisionsPager) GetNext() (page []JobResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *DeploymentJobDefinitionsListRevisionsPager) GetAll() (allItems []JobResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ExperimentsListPager can be used to simplify the use of the "ExperimentsList" method.
type ExperimentsListPager struct {
	hasNext     bool
	options     *ExperimentsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewExperimentsListPager returns a new ExperimentsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewExperimentsListPager(options *ExperimentsListOptions) (pager *ExperimentsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ExperimentsListOptions = *options
	pager = &ExperimentsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ExperimentsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ExperimentsListPager) GetNextWithContext(ctx context.Context) (page []ExperimentResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ExperimentsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ExperimentsListPager) GetAllWithContext(ctx context.Context) (allItems []ExperimentResource, err error) {
	for pager.HasNext() {
		var nextPage []ExperimentResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ExperimentsListPager) GetNext() (page []ExperimentResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ExperimentsListPager) GetAll() (allItems []ExperimentResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ExperimentsListRevisionsPager can be used to simplify the use of the "ExperimentsListRevisions" method.
type ExperimentsListRevisionsPager struct {
	hasNext     bool
	options     *ExperimentsListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewExperimentsListRevisionsPager returns a new ExperimentsListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewExperimentsListRevisionsPager(options *ExperimentsListRevisionsOptions) (pager *ExperimentsListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ExperimentsListRevisionsOptions = *options
	pager = &ExperimentsListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ExperimentsListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ExperimentsListRevisionsPager) GetNextWithContext(ctx context.Context) (page []ExperimentResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ExperimentsListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ExperimentsListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []ExperimentResource, err error) {
	for pager.HasNext() {
		var nextPage []ExperimentResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ExperimentsListRevisionsPager) GetNext() (page []ExperimentResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ExperimentsListRevisionsPager) GetAll() (allItems []ExperimentResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// FunctionsListPager can be used to simplify the use of the "FunctionsList" method.
type FunctionsListPager struct {
	hasNext     bool
	options     *FunctionsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewFunctionsListPager returns a new FunctionsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewFunctionsListPager(options *FunctionsListOptions) (pager *FunctionsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy FunctionsListOptions = *options
	pager = &FunctionsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *FunctionsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *FunctionsListPager) GetNextWithContext(ctx context.Context) (page []FunctionResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.FunctionsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *FunctionsListPager) GetAllWithContext(ctx context.Context) (allItems []FunctionResource, err error) {
	for pager.HasNext() {
		var nextPage []FunctionResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *FunctionsListPager) GetNext() (page []FunctionResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *FunctionsListPager) GetAll() (allItems []FunctionResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// FunctionsListRevisionsPager can be used to simplify the use of the "FunctionsListRevisions" method.
type FunctionsListRevisionsPager struct {
	hasNext     bool
	options     *FunctionsListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewFunctionsListRevisionsPager returns a new FunctionsListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewFunctionsListRevisionsPager(options *FunctionsListRevisionsOptions) (pager *FunctionsListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy FunctionsListRevisionsOptions = *options
	pager = &FunctionsListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *FunctionsListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *FunctionsListRevisionsPager) GetNextWithContext(ctx context.Context) (page []FunctionResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.FunctionsListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *FunctionsListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []FunctionResource, err error) {
	for pager.HasNext() {
		var nextPage []FunctionResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *FunctionsListRevisionsPager) GetNext() (page []FunctionResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *FunctionsListRevisionsPager) GetAll() (allItems []FunctionResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// InstancesListPager can be used to simplify the use of the "InstancesList" method.
type InstancesListPager struct {
	hasNext     bool
	options     *InstancesListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewInstancesListPager returns a new InstancesListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewInstancesListPager(options *InstancesListOptions) (pager *InstancesListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy InstancesListOptions = *options
	pager = &InstancesListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *InstancesListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *InstancesListPager) GetNextWithContext(ctx context.Context) (page []InstanceResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.InstancesListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *InstancesListPager) GetAllWithContext(ctx context.Context) (allItems []InstanceResource, err error) {
	for pager.HasNext() {
		var nextPage []InstanceResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *InstancesListPager) GetNext() (page []InstanceResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *InstancesListPager) GetAll() (allItems []InstanceResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ModelsListPager can be used to simplify the use of the "ModelsList" method.
type ModelsListPager struct {
	hasNext     bool
	options     *ModelsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewModelsListPager returns a new ModelsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewModelsListPager(options *ModelsListOptions) (pager *ModelsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ModelsListOptions = *options
	pager = &ModelsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ModelsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ModelsListPager) GetNextWithContext(ctx context.Context) (page []ModelResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ModelsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ModelsListPager) GetAllWithContext(ctx context.Context) (allItems []ModelResource, err error) {
	for pager.HasNext() {
		var nextPage []ModelResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ModelsListPager) GetNext() (page []ModelResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ModelsListPager) GetAll() (allItems []ModelResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ModelsListRevisionsPager can be used to simplify the use of the "ModelsListRevisions" method.
type ModelsListRevisionsPager struct {
	hasNext     bool
	options     *ModelsListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewModelsListRevisionsPager returns a new ModelsListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewModelsListRevisionsPager(options *ModelsListRevisionsOptions) (pager *ModelsListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ModelsListRevisionsOptions = *options
	pager = &ModelsListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ModelsListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ModelsListRevisionsPager) GetNextWithContext(ctx context.Context) (page []ModelResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ModelsListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ModelsListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []ModelResource, err error) {
	for pager.HasNext() {
		var nextPage []ModelResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ModelsListRevisionsPager) GetNext() (page []ModelResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ModelsListRevisionsPager) GetAll() (allItems []ModelResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ModelDefinitionsListPager can be used to simplify the use of the "ModelDefinitionsList" method.
type ModelDefinitionsListPager struct {
	hasNext     bool
	options     *ModelDefinitionsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewModelDefinitionsListPager returns a new ModelDefinitionsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewModelDefinitionsListPager(options *ModelDefinitionsListOptions) (pager *ModelDefinitionsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ModelDefinitionsListOptions = *options
	pager = &ModelDefinitionsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ModelDefinitionsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ModelDefinitionsListPager) GetNextWithContext(ctx context.Context) (page []ModelDefinitionResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ModelDefinitionsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ModelDefinitionsListPager) GetAllWithContext(ctx context.Context) (allItems []ModelDefinitionResource, err error) {
	for pager.HasNext() {
		var nextPage []ModelDefinitionResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ModelDefinitionsListPager) GetNext() (page []ModelDefinitionResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ModelDefinitionsListPager) GetAll() (allItems []ModelDefinitionResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ModelDefinitionsListRevisionsPager can be used to simplify the use of the "ModelDefinitionsListRevisions" method.
type ModelDefinitionsListRevisionsPager struct {
	hasNext     bool
	options     *ModelDefinitionsListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewModelDefinitionsListRevisionsPager returns a new ModelDefinitionsListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewModelDefinitionsListRevisionsPager(options *ModelDefinitionsListRevisionsOptions) (pager *ModelDefinitionsListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ModelDefinitionsListRevisionsOptions = *options
	pager = &ModelDefinitionsListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ModelDefinitionsListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ModelDefinitionsListRevisionsPager) GetNextWithContext(ctx context.Context) (page []ModelDefinitionResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ModelDefinitionsListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ModelDefinitionsListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []ModelDefinitionResource, err error) {
	for pager.HasNext() {
		var nextPage []ModelDefinitionResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ModelDefinitionsListRevisionsPager) GetNext() (page []ModelDefinitionResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ModelDefinitionsListRevisionsPager) GetAll() (allItems []ModelDefinitionResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// PipelinesListPager can be used to simplify the use of the "PipelinesList" method.
type PipelinesListPager struct {
	hasNext     bool
	options     *PipelinesListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewPipelinesListPager returns a new PipelinesListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewPipelinesListPager(options *PipelinesListOptions) (pager *PipelinesListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy PipelinesListOptions = *options
	pager = &PipelinesListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *PipelinesListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *PipelinesListPager) GetNextWithContext(ctx context.Context) (page []PipelineResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.PipelinesListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *PipelinesListPager) GetAllWithContext(ctx context.Context) (allItems []PipelineResource, err error) {
	for pager.HasNext() {
		var nextPage []PipelineResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *PipelinesListPager) GetNext() (page []PipelineResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *PipelinesListPager) GetAll() (allItems []PipelineResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// PipelinesListRevisionsPager can be used to simplify the use of the "PipelinesListRevisions" method.
type PipelinesListRevisionsPager struct {
	hasNext     bool
	options     *PipelinesListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewPipelinesListRevisionsPager returns a new PipelinesListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewPipelinesListRevisionsPager(options *PipelinesListRevisionsOptions) (pager *PipelinesListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy PipelinesListRevisionsOptions = *options
	pager = &PipelinesListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *PipelinesListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *PipelinesListRevisionsPager) GetNextWithContext(ctx context.Context) (page []PipelineResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.PipelinesListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *PipelinesListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []PipelineResource, err error) {
	for pager.HasNext() {
		var nextPage []PipelineResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *PipelinesListRevisionsPager) GetNext() (page []PipelineResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *PipelinesListRevisionsPager) GetAll() (allItems []PipelineResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// RemoteTrainingSystemsListPager can be used to simplify the use of the "RemoteTrainingSystemsList" method.
type RemoteTrainingSystemsListPager struct {
	hasNext     bool
	options     *RemoteTrainingSystemsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewRemoteTrainingSystemsListPager returns a new RemoteTrainingSystemsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewRemoteTrainingSystemsListPager(options *RemoteTrainingSystemsListOptions) (pager *RemoteTrainingSystemsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy RemoteTrainingSystemsListOptions = *options
	pager = &RemoteTrainingSystemsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *RemoteTrainingSystemsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *RemoteTrainingSystemsListPager) GetNextWithContext(ctx context.Context) (page []RemoteTrainingSystemResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.RemoteTrainingSystemsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *RemoteTrainingSystemsListPager) GetAllWithContext(ctx context.Context) (allItems []RemoteTrainingSystemResource, err error) {
	for pager.HasNext() {
		var nextPage []RemoteTrainingSystemResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *RemoteTrainingSystemsListPager) GetNext() (page []RemoteTrainingSystemResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *RemoteTrainingSystemsListPager) GetAll() (allItems []RemoteTrainingSystemResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// RemoteTrainingSystemsListRevisionsPager can be used to simplify the use of the "RemoteTrainingSystemsListRevisions" method.
type RemoteTrainingSystemsListRevisionsPager struct {
	hasNext     bool
	options     *RemoteTrainingSystemsListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewRemoteTrainingSystemsListRevisionsPager returns a new RemoteTrainingSystemsListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewRemoteTrainingSystemsListRevisionsPager(options *RemoteTrainingSystemsListRevisionsOptions) (pager *RemoteTrainingSystemsListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy RemoteTrainingSystemsListRevisionsOptions = *options
	pager = &RemoteTrainingSystemsListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *RemoteTrainingSystemsListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *RemoteTrainingSystemsListRevisionsPager) GetNextWithContext(ctx context.Context) (page []RemoteTrainingSystemResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.RemoteTrainingSystemsListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *RemoteTrainingSystemsListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []RemoteTrainingSystemResource, err error) {
	for pager.HasNext() {
		var nextPage []RemoteTrainingSystemResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *RemoteTrainingSystemsListRevisionsPager) GetNext() (page []RemoteTrainingSystemResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *RemoteTrainingSystemsListRevisionsPager) GetAll() (allItems []RemoteTrainingSystemResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// TrainingsListPager can be used to simplify the use of the "TrainingsList" method.
type TrainingsListPager struct {
	hasNext     bool
	options     *TrainingsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewTrainingsListPager returns a new TrainingsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewTrainingsListPager(options *TrainingsListOptions) (pager *TrainingsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy TrainingsListOptions = *options
	pager = &TrainingsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *TrainingsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *TrainingsListPager) GetNextWithContext(ctx context.Context) (page []TrainingResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.TrainingsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *TrainingsListPager) GetAllWithContext(ctx context.Context) (allItems []TrainingResource, err error) {
	for pager.HasNext() {
		var nextPage []TrainingResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *TrainingsListPager) GetNext() (page []TrainingResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *TrainingsListPager) GetAll() (allItems []TrainingResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// TrainingDefinitionsListPager can be used to simplify the use of the "TrainingDefinitionsList" method.
type TrainingDefinitionsListPager struct {
	hasNext     bool
	options     *TrainingDefinitionsListOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewTrainingDefinitionsListPager returns a new TrainingDefinitionsListPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewTrainingDefinitionsListPager(options *TrainingDefinitionsListOptions) (pager *TrainingDefinitionsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy TrainingDefinitionsListOptions = *options
	pager = &TrainingDefinitionsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *TrainingDefinitionsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *TrainingDefinitionsListPager) GetNextWithContext(ctx context.Context) (page []TrainingDefinitionResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.TrainingDefinitionsListWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *TrainingDefinitionsListPager) GetAllWithContext(ctx context.Context) (allItems []TrainingDefinitionResource, err error) {
	for pager.HasNext() {
		var nextPage []TrainingDefinitionResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *TrainingDefinitionsListPager) GetNext() (page []TrainingDefinitionResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *TrainingDefinitionsListPager) GetAll() (allItems []TrainingDefinitionResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// TrainingDefinitionsListRevisionsPager can be used to simplify the use of the "TrainingDefinitionsListRevisions" method.
type TrainingDefinitionsListRevisionsPager struct {
	hasNext     bool
	options     *TrainingDefinitionsListRevisionsOptions
	client      *WatsonMachineLearningV4
	pageContext struct {
		next *string
	}
}

// NewTrainingDefinitionsListRevisionsPager returns a new TrainingDefinitionsListRevisionsPager instance.
func (watsonMachineLearning *WatsonMachineLearningV4) NewTrainingDefinitionsListRevisionsPager(options *TrainingDefinitionsListRevisionsOptions) (pager *TrainingDefinitionsListRevisionsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy TrainingDefinitionsListRevisionsOptions = *options
	pager = &TrainingDefinitionsListRevisionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  watsonMachineLearning,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *TrainingDefinitionsListRevisionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *TrainingDefinitionsListRevisionsPager) GetNextWithContext(ctx context.Context) (page []TrainingDefinitionResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.TrainingDefinitionsListRevisionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		var start *string
		start, err = core.GetQueryParam(result.Next.Href, "start")
		if err != nil {
			err = fmt.Errorf("error retrieving 'start' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Resources

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *TrainingDefinitionsListRevisionsPager) GetAllWithContext(ctx context.Context) (allItems []TrainingDefinitionResource, err error) {
	for pager.HasNext() {
		var nextPage []TrainingDefinitionResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *TrainingDefinitionsListRevisionsPager) GetNext() (page []TrainingDefinitionResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *TrainingDefinitionsListRevisionsPager) GetAll() (allItems []TrainingDefinitionResource, err error) {
	return pager.GetAllWithContext(context.Background())
}
