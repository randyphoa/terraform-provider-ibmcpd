package spacev2

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	common "terraform-provider-ibmcpd/internal/go-sdk/common"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

// SpaceV2 : ## API overview
//
// Step by step instructions on how to use spaces API (v2) can be found
// [here](https://github.ibm.com/AILifecycle/cpd-spaces-api/blob/master/README.md).
//
// API Version: 2.0.0
type SpaceV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.dataplatform.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "space"

// SpaceV2Options : Service options
type SpaceV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewSpaceV2UsingExternalConfig : constructs an instance of SpaceV2 with passed in options and external configuration.
func NewSpaceV2UsingExternalConfig(options *SpaceV2Options) (space *SpaceV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	space, err = NewSpaceV2(options)
	if err != nil {
		return
	}

	err = space.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = space.Service.SetServiceURL(options.URL)
	}
	return
}

// NewSpaceV2 : constructs an instance of SpaceV2 with passed in options.
func NewSpaceV2(options *SpaceV2Options) (service *SpaceV2, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
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

	service = &SpaceV2{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "space" suitable for processing requests.
func (space *SpaceV2) Clone() *SpaceV2 {
	if core.IsNil(space) {
		return nil
	}
	clone := *space
	clone.Service = space.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (space *SpaceV2) SetServiceURL(url string) error {
	return space.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (space *SpaceV2) GetServiceURL() string {
	return space.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (space *SpaceV2) SetDefaultHeaders(headers http.Header) {
	space.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (space *SpaceV2) SetEnableGzipCompression(enableGzip bool) {
	space.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (space *SpaceV2) GetEnableGzipCompression() bool {
	return space.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (space *SpaceV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	space.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (space *SpaceV2) DisableRetries() {
	space.Service.DisableRetries()
}

// SpacesCreate : Create a new space
// Creates a new space to scope other assets. Authorized user must have the following roles (see
// https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-iams)
// - Platform management role: Administrator
// - Service access role: Manager
//
// On Public Cloud user is required to provide Cloud Object Storage instance details in the 'storage' property. On
// private CPD installations the default storage is used instead.
func (space *SpaceV2) SpacesCreate(spacesCreateOptions *SpacesCreateOptions) (result *SpaceResource, response *core.DetailedResponse, err error) {
	return space.SpacesCreateWithContext(context.Background(), spacesCreateOptions)
}

// SpacesCreateWithContext is an alternate form of the SpacesCreate method which supports a Context parameter
func (space *SpaceV2) SpacesCreateWithContext(ctx context.Context, spacesCreateOptions *SpacesCreateOptions) (result *SpaceResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(spacesCreateOptions, "spacesCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(spacesCreateOptions, "spacesCreateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range spacesCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "SpacesCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if spacesCreateOptions.Name != nil {
		body["name"] = spacesCreateOptions.Name
	}
	if spacesCreateOptions.Description != nil {
		body["description"] = spacesCreateOptions.Description
	}
	if spacesCreateOptions.Storage != nil {
		body["storage"] = spacesCreateOptions.Storage
	}
	if spacesCreateOptions.Compute != nil {
		body["compute"] = spacesCreateOptions.Compute
	}
	if spacesCreateOptions.Tags != nil {
		body["tags"] = spacesCreateOptions.Tags
	}
	if spacesCreateOptions.Generator != nil {
		body["generator"] = spacesCreateOptions.Generator
	}
	if spacesCreateOptions.Stage != nil {
		body["stage"] = spacesCreateOptions.Stage
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
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpaceResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SpacesList : Retrieve the spaces
// Retrieves the space list.
func (space *SpaceV2) SpacesList(spacesListOptions *SpacesListOptions) (result *SpaceResources, response *core.DetailedResponse, err error) {
	return space.SpacesListWithContext(context.Background(), spacesListOptions)
}

// SpacesListWithContext is an alternate form of the SpacesList method which supports a Context parameter
func (space *SpaceV2) SpacesListWithContext(ctx context.Context, spacesListOptions *SpacesListOptions) (result *SpaceResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(spacesListOptions, "spacesListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range spacesListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "SpacesList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if spacesListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*spacesListOptions.Start))
	}
	if spacesListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*spacesListOptions.Limit))
	}
	if spacesListOptions.TotalCount != nil {
		builder.AddQuery("total_count", fmt.Sprint(*spacesListOptions.TotalCount))
	}
	if spacesListOptions.ID != nil {
		builder.AddQuery("id", fmt.Sprint(*spacesListOptions.ID))
	}
	if spacesListOptions.Tags != nil {
		builder.AddQuery("tags", fmt.Sprint(*spacesListOptions.Tags))
	}
	if spacesListOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*spacesListOptions.Include))
	}
	if spacesListOptions.Member != nil {
		builder.AddQuery("member", fmt.Sprint(*spacesListOptions.Member))
	}
	if spacesListOptions.Roles != nil {
		builder.AddQuery("roles", fmt.Sprint(*spacesListOptions.Roles))
	}
	if spacesListOptions.BssAccountID != nil {
		builder.AddQuery("bss_account_id", fmt.Sprint(*spacesListOptions.BssAccountID))
	}
	if spacesListOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*spacesListOptions.Name))
	}
	if spacesListOptions.ComputeCrn != nil {
		builder.AddQuery("compute.crn", fmt.Sprint(*spacesListOptions.ComputeCrn))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpaceResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SpacesGet : Retrieve the space
// Retrieves the space with the specified identifier.
func (space *SpaceV2) SpacesGet(spacesGetOptions *SpacesGetOptions) (result *SpaceResource, response *core.DetailedResponse, err error) {
	return space.SpacesGetWithContext(context.Background(), spacesGetOptions)
}

// SpacesGetWithContext is an alternate form of the SpacesGet method which supports a Context parameter
func (space *SpaceV2) SpacesGetWithContext(ctx context.Context, spacesGetOptions *SpacesGetOptions) (result *SpaceResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(spacesGetOptions, "spacesGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(spacesGetOptions, "spacesGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id": *spacesGetOptions.SpaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range spacesGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "SpacesGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if spacesGetOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*spacesGetOptions.Include))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpaceResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SpacesDelete : Delete the space
// Deletes the space with the specified identifier.
func (space *SpaceV2) SpacesDelete(spacesDeleteOptions *SpacesDeleteOptions) (response *core.DetailedResponse, err error) {
	return space.SpacesDeleteWithContext(context.Background(), spacesDeleteOptions)
}

// SpacesDeleteWithContext is an alternate form of the SpacesDelete method which supports a Context parameter
func (space *SpaceV2) SpacesDeleteWithContext(ctx context.Context, spacesDeleteOptions *SpacesDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(spacesDeleteOptions, "spacesDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(spacesDeleteOptions, "spacesDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id": *spacesDeleteOptions.SpaceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range spacesDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "SpacesDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = space.Service.Request(request, nil)

	return
}

// SpacesUpdate : Update the space
// Partially update this space. Allowed paths are:
//   - /name
//   - /description
//   - /compute
//   - /stage/name.
func (space *SpaceV2) SpacesUpdate(spacesUpdateOptions *SpacesUpdateOptions) (result *SpaceResource, response *core.DetailedResponse, err error) {
	return space.SpacesUpdateWithContext(context.Background(), spacesUpdateOptions)
}

// SpacesUpdateWithContext is an alternate form of the SpacesUpdate method which supports a Context parameter
func (space *SpaceV2) SpacesUpdateWithContext(ctx context.Context, spacesUpdateOptions *SpacesUpdateOptions) (result *SpaceResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(spacesUpdateOptions, "spacesUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(spacesUpdateOptions, "spacesUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id": *spacesUpdateOptions.SpaceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range spacesUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "SpacesUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(spacesUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpaceResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MembersList : Retrieve the space members
// Retrieves the member list for the specified space.
func (space *SpaceV2) MembersList(membersListOptions *MembersListOptions) (result *MemberResources, response *core.DetailedResponse, err error) {
	return space.MembersListWithContext(context.Background(), membersListOptions)
}

// MembersListWithContext is an alternate form of the MembersList method which supports a Context parameter
func (space *SpaceV2) MembersListWithContext(ctx context.Context, membersListOptions *MembersListOptions) (result *MemberResources, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(membersListOptions, "membersListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(membersListOptions, "membersListOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id": *membersListOptions.SpaceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}/members`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range membersListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "MembersList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if membersListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*membersListOptions.Start))
	}
	if membersListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*membersListOptions.Limit))
	}
	if membersListOptions.TotalCount != nil {
		builder.AddQuery("total_count", fmt.Sprint(*membersListOptions.TotalCount))
	}
	if membersListOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*membersListOptions.Type))
	}
	if membersListOptions.Role != nil {
		builder.AddQuery("role", fmt.Sprint(*membersListOptions.Role))
	}
	if membersListOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*membersListOptions.State))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMemberResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MembersCreate : Add a member to space
// Adds a member to the specified space.
func (space *SpaceV2) MembersCreate(membersCreateOptions *MembersCreateOptions) (result *MemberArray, response *core.DetailedResponse, err error) {
	return space.MembersCreateWithContext(context.Background(), membersCreateOptions)
}

// MembersCreateWithContext is an alternate form of the MembersCreate method which supports a Context parameter
func (space *SpaceV2) MembersCreateWithContext(ctx context.Context, membersCreateOptions *MembersCreateOptions) (result *MemberArray, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(membersCreateOptions, "membersCreateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(membersCreateOptions, "membersCreateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id": *membersCreateOptions.SpaceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}/members`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range membersCreateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "MembersCreate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if membersCreateOptions.Members != nil {
		body["members"] = membersCreateOptions.Members
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
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMemberArray)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MembersGet : Retrieve the space member information
// Retrieves the member information for the member and space with the specified identifiers.
func (space *SpaceV2) MembersGet(membersGetOptions *MembersGetOptions) (result *MemberResource, response *core.DetailedResponse, err error) {
	return space.MembersGetWithContext(context.Background(), membersGetOptions)
}

// MembersGetWithContext is an alternate form of the MembersGet method which supports a Context parameter
func (space *SpaceV2) MembersGetWithContext(ctx context.Context, membersGetOptions *MembersGetOptions) (result *MemberResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(membersGetOptions, "membersGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(membersGetOptions, "membersGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id":  *membersGetOptions.SpaceID,
		"member_id": *membersGetOptions.MemberID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}/members/{member_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range membersGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "MembersGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMemberResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// MembersDelete : Remove the space member
// Removes the specified member from the space.
func (space *SpaceV2) MembersDelete(membersDeleteOptions *MembersDeleteOptions) (response *core.DetailedResponse, err error) {
	return space.MembersDeleteWithContext(context.Background(), membersDeleteOptions)
}

// MembersDeleteWithContext is an alternate form of the MembersDelete method which supports a Context parameter
func (space *SpaceV2) MembersDeleteWithContext(ctx context.Context, membersDeleteOptions *MembersDeleteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(membersDeleteOptions, "membersDeleteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(membersDeleteOptions, "membersDeleteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id":  *membersDeleteOptions.SpaceID,
		"member_id": *membersDeleteOptions.MemberID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}/members/{member_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range membersDeleteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "MembersDelete")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = space.Service.Request(request, nil)

	return
}

// MembersUpdate : Update the space member
// Partially update the member selected with the specified identifier. Allowed paths are:
//   - /role
//   - /state.
func (space *SpaceV2) MembersUpdate(membersUpdateOptions *MembersUpdateOptions) (result *MemberResource, response *core.DetailedResponse, err error) {
	return space.MembersUpdateWithContext(context.Background(), membersUpdateOptions)
}

// MembersUpdateWithContext is an alternate form of the MembersUpdate method which supports a Context parameter
func (space *SpaceV2) MembersUpdateWithContext(ctx context.Context, membersUpdateOptions *MembersUpdateOptions) (result *MemberResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(membersUpdateOptions, "membersUpdateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(membersUpdateOptions, "membersUpdateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"space_id":  *membersUpdateOptions.SpaceID,
		"member_id": *membersUpdateOptions.MemberID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/spaces/{space_id}/members/{member_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range membersUpdateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "MembersUpdate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(membersUpdateOptions.JSONPatch)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMemberResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExportsList : Retrieve the asset exports
// Retrieves the asset export list for the specified space, project, or catalog.
func (space *SpaceV2) ExportsList(exportsListOptions *ExportsListOptions) (result *ExportResources, response *core.DetailedResponse, err error) {
	return space.ExportsListWithContext(context.Background(), exportsListOptions)
}

// ExportsListWithContext is an alternate form of the ExportsList method which supports a Context parameter
func (space *SpaceV2) ExportsListWithContext(ctx context.Context, exportsListOptions *ExportsListOptions) (result *ExportResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(exportsListOptions, "exportsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_exports`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range exportsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ExportsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if exportsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*exportsListOptions.SpaceID))
	}
	if exportsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*exportsListOptions.ProjectID))
	}
	if exportsListOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*exportsListOptions.CatalogID))
	}
	if exportsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*exportsListOptions.Start))
	}
	if exportsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*exportsListOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExportResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExportsStart : Start the asset export process
// Starts the asset export process for the specified space, project, or catalog. On CPD 3.0.1 assets export is supported
// only in the context of a space.
func (space *SpaceV2) ExportsStart(exportsStartOptions *ExportsStartOptions) (result *ExportResource, response *core.DetailedResponse, err error) {
	return space.ExportsStartWithContext(context.Background(), exportsStartOptions)
}

// ExportsStartWithContext is an alternate form of the ExportsStart method which supports a Context parameter
func (space *SpaceV2) ExportsStartWithContext(ctx context.Context, exportsStartOptions *ExportsStartOptions) (result *ExportResource, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(exportsStartOptions, "exportsStartOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_exports`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range exportsStartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ExportsStart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if exportsStartOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*exportsStartOptions.SpaceID))
	}
	if exportsStartOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*exportsStartOptions.ProjectID))
	}
	if exportsStartOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*exportsStartOptions.CatalogID))
	}

	body := make(map[string]interface{})
	if exportsStartOptions.Name != nil {
		body["name"] = exportsStartOptions.Name
	}
	if exportsStartOptions.Description != nil {
		body["description"] = exportsStartOptions.Description
	}
	if exportsStartOptions.Assets != nil {
		body["assets"] = exportsStartOptions.Assets
	}
	if exportsStartOptions.EncryptionKey != nil {
		body["encryption_key"] = exportsStartOptions.EncryptionKey
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
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExportResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExportsCancel : Cancel the asset export process
// Cancels the asset export process with the specified identifier.
func (space *SpaceV2) ExportsCancel(exportsCancelOptions *ExportsCancelOptions) (response *core.DetailedResponse, err error) {
	return space.ExportsCancelWithContext(context.Background(), exportsCancelOptions)
}

// ExportsCancelWithContext is an alternate form of the ExportsCancel method which supports a Context parameter
func (space *SpaceV2) ExportsCancelWithContext(ctx context.Context, exportsCancelOptions *ExportsCancelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(exportsCancelOptions, "exportsCancelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(exportsCancelOptions, "exportsCancelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"export_id": *exportsCancelOptions.ExportID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_exports/{export_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range exportsCancelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ExportsCancel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if exportsCancelOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*exportsCancelOptions.SpaceID))
	}
	if exportsCancelOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*exportsCancelOptions.ProjectID))
	}
	if exportsCancelOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*exportsCancelOptions.CatalogID))
	}
	if exportsCancelOptions.HardDelete != nil {
		builder.AddQuery("hard_delete", fmt.Sprint(*exportsCancelOptions.HardDelete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = space.Service.Request(request, nil)

	return
}

// ExportsGet : Retrieve the asset export
// Retrieves the asset export with the specified identifier.
func (space *SpaceV2) ExportsGet(exportsGetOptions *ExportsGetOptions) (result *ExportResource, response *core.DetailedResponse, err error) {
	return space.ExportsGetWithContext(context.Background(), exportsGetOptions)
}

// ExportsGetWithContext is an alternate form of the ExportsGet method which supports a Context parameter
func (space *SpaceV2) ExportsGetWithContext(ctx context.Context, exportsGetOptions *ExportsGetOptions) (result *ExportResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(exportsGetOptions, "exportsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(exportsGetOptions, "exportsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"export_id": *exportsGetOptions.ExportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_exports/{export_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range exportsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ExportsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if exportsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*exportsGetOptions.SpaceID))
	}
	if exportsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*exportsGetOptions.ProjectID))
	}
	if exportsGetOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*exportsGetOptions.CatalogID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExportResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExportsDownload : Download the asset export content
// Downloads the content for the asset export process with the specified identifier.
func (space *SpaceV2) ExportsDownload(exportsDownloadOptions *ExportsDownloadOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return space.ExportsDownloadWithContext(context.Background(), exportsDownloadOptions)
}

// ExportsDownloadWithContext is an alternate form of the ExportsDownload method which supports a Context parameter
func (space *SpaceV2) ExportsDownloadWithContext(ctx context.Context, exportsDownloadOptions *ExportsDownloadOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(exportsDownloadOptions, "exportsDownloadOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(exportsDownloadOptions, "exportsDownloadOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"export_id": *exportsDownloadOptions.ExportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_exports/{export_id}/content`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range exportsDownloadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ExportsDownload")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/zip")

	if exportsDownloadOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*exportsDownloadOptions.SpaceID))
	}
	if exportsDownloadOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*exportsDownloadOptions.ProjectID))
	}
	if exportsDownloadOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*exportsDownloadOptions.CatalogID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = space.Service.Request(request, &result)

	return
}

// ImportsList : Retrieve the asset imports
// Retrieves the asset import list for the specified space, project, or catalog.
func (space *SpaceV2) ImportsList(importsListOptions *ImportsListOptions) (result *ImportResources, response *core.DetailedResponse, err error) {
	return space.ImportsListWithContext(context.Background(), importsListOptions)
}

// ImportsListWithContext is an alternate form of the ImportsList method which supports a Context parameter
func (space *SpaceV2) ImportsListWithContext(ctx context.Context, importsListOptions *ImportsListOptions) (result *ImportResources, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(importsListOptions, "importsListOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_imports`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range importsListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ImportsList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if importsListOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*importsListOptions.SpaceID))
	}
	if importsListOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*importsListOptions.ProjectID))
	}
	if importsListOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*importsListOptions.CatalogID))
	}
	if importsListOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*importsListOptions.Start))
	}
	if importsListOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*importsListOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImportResources)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ImportsStart : Start the asset import process
// Starts the asset import process for the specified space, project, or catalog. On CPD 3.0.1 assets import is supported
// only in the context of a space.
func (space *SpaceV2) ImportsStart(importsStartOptions *ImportsStartOptions) (result *ImportResource, response *core.DetailedResponse, err error) {
	return space.ImportsStartWithContext(context.Background(), importsStartOptions)
}

// ImportsStartWithContext is an alternate form of the ImportsStart method which supports a Context parameter
func (space *SpaceV2) ImportsStartWithContext(ctx context.Context, importsStartOptions *ImportsStartOptions) (result *ImportResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importsStartOptions, "importsStartOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importsStartOptions, "importsStartOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_imports`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range importsStartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ImportsStart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if importsStartOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*importsStartOptions.SpaceID))
	}
	if importsStartOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*importsStartOptions.ProjectID))
	}
	if importsStartOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*importsStartOptions.CatalogID))
	}

	builder.AddFormData("file", "filename",
		core.StringNilMapper(importsStartOptions.FileContentType), importsStartOptions.File)
	if importsStartOptions.EncryptionKey != nil {
		builder.AddFormData("encryption_key", "", "", fmt.Sprint(*importsStartOptions.EncryptionKey))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImportResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ImportsCancel : Cancel the asset import process
// Cancels the asset import process with the specified identifier.
func (space *SpaceV2) ImportsCancel(importsCancelOptions *ImportsCancelOptions) (response *core.DetailedResponse, err error) {
	return space.ImportsCancelWithContext(context.Background(), importsCancelOptions)
}

// ImportsCancelWithContext is an alternate form of the ImportsCancel method which supports a Context parameter
func (space *SpaceV2) ImportsCancelWithContext(ctx context.Context, importsCancelOptions *ImportsCancelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importsCancelOptions, "importsCancelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importsCancelOptions, "importsCancelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"import_id": *importsCancelOptions.ImportID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_imports/{import_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range importsCancelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ImportsCancel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if importsCancelOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*importsCancelOptions.SpaceID))
	}
	if importsCancelOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*importsCancelOptions.ProjectID))
	}
	if importsCancelOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*importsCancelOptions.CatalogID))
	}
	if importsCancelOptions.HardDelete != nil {
		builder.AddQuery("hard_delete", fmt.Sprint(*importsCancelOptions.HardDelete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = space.Service.Request(request, nil)

	return
}

// ImportsGet : Retrieve the asset import
// Retrieves the asset import with the specified identifier.
func (space *SpaceV2) ImportsGet(importsGetOptions *ImportsGetOptions) (result *ImportResource, response *core.DetailedResponse, err error) {
	return space.ImportsGetWithContext(context.Background(), importsGetOptions)
}

// ImportsGetWithContext is an alternate form of the ImportsGet method which supports a Context parameter
func (space *SpaceV2) ImportsGetWithContext(ctx context.Context, importsGetOptions *ImportsGetOptions) (result *ImportResource, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importsGetOptions, "importsGetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importsGetOptions, "importsGetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"import_id": *importsGetOptions.ImportID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = space.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(space.Service.Options.URL, `/v2/asset_imports/{import_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range importsGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("space", "V2", "ImportsGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if importsGetOptions.SpaceID != nil {
		builder.AddQuery("space_id", fmt.Sprint(*importsGetOptions.SpaceID))
	}
	if importsGetOptions.ProjectID != nil {
		builder.AddQuery("project_id", fmt.Sprint(*importsGetOptions.ProjectID))
	}
	if importsGetOptions.CatalogID != nil {
		builder.AddQuery("catalog_id", fmt.Sprint(*importsGetOptions.CatalogID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = space.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImportResource)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ErrorErrorsItem : ErrorErrorsItem struct
type ErrorErrorsItem struct {
	// A simple code that should convey the general sense of the error.
	Code *string `json:"code" validate:"required"`

	// The message that describes the error.
	Message *string `json:"message" validate:"required"`

	Parameters []string `json:"parameters,omitempty"`

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
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
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

// ExportsCancelOptions : The ExportsCancel options.
type ExportsCancelOptions struct {
	// The export identification.
	ExportID *string `json:"export_id" validate:"required,ne="`

	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Default is false.
	HardDelete *bool `json:"hard_delete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExportsCancelOptions : Instantiate ExportsCancelOptions
func (*SpaceV2) NewExportsCancelOptions(exportID string) *ExportsCancelOptions {
	return &ExportsCancelOptions{
		ExportID: core.StringPtr(exportID),
	}
}

// SetExportID : Allow user to set ExportID
func (_options *ExportsCancelOptions) SetExportID(exportID string) *ExportsCancelOptions {
	_options.ExportID = core.StringPtr(exportID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExportsCancelOptions) SetSpaceID(spaceID string) *ExportsCancelOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExportsCancelOptions) SetProjectID(projectID string) *ExportsCancelOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ExportsCancelOptions) SetCatalogID(catalogID string) *ExportsCancelOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetHardDelete : Allow user to set HardDelete
func (_options *ExportsCancelOptions) SetHardDelete(hardDelete bool) *ExportsCancelOptions {
	_options.HardDelete = core.BoolPtr(hardDelete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExportsCancelOptions) SetHeaders(param map[string]string) *ExportsCancelOptions {
	options.Headers = param
	return options
}

// ExportsDownloadOptions : The ExportsDownload options.
type ExportsDownloadOptions struct {
	// The export identification.
	ExportID *string `json:"export_id" validate:"required,ne="`

	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExportsDownloadOptions : Instantiate ExportsDownloadOptions
func (*SpaceV2) NewExportsDownloadOptions(exportID string) *ExportsDownloadOptions {
	return &ExportsDownloadOptions{
		ExportID: core.StringPtr(exportID),
	}
}

// SetExportID : Allow user to set ExportID
func (_options *ExportsDownloadOptions) SetExportID(exportID string) *ExportsDownloadOptions {
	_options.ExportID = core.StringPtr(exportID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExportsDownloadOptions) SetSpaceID(spaceID string) *ExportsDownloadOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExportsDownloadOptions) SetProjectID(projectID string) *ExportsDownloadOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ExportsDownloadOptions) SetCatalogID(catalogID string) *ExportsDownloadOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExportsDownloadOptions) SetHeaders(param map[string]string) *ExportsDownloadOptions {
	options.Headers = param
	return options
}

// ExportsGetOptions : The ExportsGet options.
type ExportsGetOptions struct {
	// The export identification.
	ExportID *string `json:"export_id" validate:"required,ne="`

	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExportsGetOptions : Instantiate ExportsGetOptions
func (*SpaceV2) NewExportsGetOptions(exportID string) *ExportsGetOptions {
	return &ExportsGetOptions{
		ExportID: core.StringPtr(exportID),
	}
}

// SetExportID : Allow user to set ExportID
func (_options *ExportsGetOptions) SetExportID(exportID string) *ExportsGetOptions {
	_options.ExportID = core.StringPtr(exportID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExportsGetOptions) SetSpaceID(spaceID string) *ExportsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExportsGetOptions) SetProjectID(projectID string) *ExportsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ExportsGetOptions) SetCatalogID(catalogID string) *ExportsGetOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExportsGetOptions) SetHeaders(param map[string]string) *ExportsGetOptions {
	options.Headers = param
	return options
}

// ExportsListOptions : The ExportsList options.
type ExportsListOptions struct {
	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Token representing first resource.
	Start *string `json:"start,omitempty"`

	// The number of resources returned. Default value is 100. Max value is 200.
	Limit *float64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExportsListOptions : Instantiate ExportsListOptions
func (*SpaceV2) NewExportsListOptions() *ExportsListOptions {
	return &ExportsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExportsListOptions) SetSpaceID(spaceID string) *ExportsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExportsListOptions) SetProjectID(projectID string) *ExportsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ExportsListOptions) SetCatalogID(catalogID string) *ExportsListOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ExportsListOptions) SetStart(start string) *ExportsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ExportsListOptions) SetLimit(limit float64) *ExportsListOptions {
	_options.Limit = core.Float64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExportsListOptions) SetHeaders(param map[string]string) *ExportsListOptions {
	options.Headers = param
	return options
}

// ExportsStartOptions : The ExportsStart options.
type ExportsStartOptions struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Assets *ExportAssets `json:"assets,omitempty"`

	// Encryption key used to encrypt the sensitive data during export.
	EncryptionKey *string `json:"encryption_key,omitempty"`

	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExportsStartOptions : Instantiate ExportsStartOptions
func (*SpaceV2) NewExportsStartOptions() *ExportsStartOptions {
	return &ExportsStartOptions{}
}

// SetName : Allow user to set Name
func (_options *ExportsStartOptions) SetName(name string) *ExportsStartOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ExportsStartOptions) SetDescription(description string) *ExportsStartOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAssets : Allow user to set Assets
func (_options *ExportsStartOptions) SetAssets(assets *ExportAssets) *ExportsStartOptions {
	_options.Assets = assets
	return _options
}

// SetEncryptionKey : Allow user to set EncryptionKey
func (_options *ExportsStartOptions) SetEncryptionKey(encryptionKey string) *ExportsStartOptions {
	_options.EncryptionKey = core.StringPtr(encryptionKey)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ExportsStartOptions) SetSpaceID(spaceID string) *ExportsStartOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ExportsStartOptions) SetProjectID(projectID string) *ExportsStartOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ExportsStartOptions) SetCatalogID(catalogID string) *ExportsStartOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExportsStartOptions) SetHeaders(param map[string]string) *ExportsStartOptions {
	options.Headers = param
	return options
}

// ImportsCancelOptions : The ImportsCancel options.
type ImportsCancelOptions struct {
	// The import identification.
	ImportID *string `json:"import_id" validate:"required,ne="`

	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Default is false.
	HardDelete *bool `json:"hard_delete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportsCancelOptions : Instantiate ImportsCancelOptions
func (*SpaceV2) NewImportsCancelOptions(importID string) *ImportsCancelOptions {
	return &ImportsCancelOptions{
		ImportID: core.StringPtr(importID),
	}
}

// SetImportID : Allow user to set ImportID
func (_options *ImportsCancelOptions) SetImportID(importID string) *ImportsCancelOptions {
	_options.ImportID = core.StringPtr(importID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ImportsCancelOptions) SetSpaceID(spaceID string) *ImportsCancelOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ImportsCancelOptions) SetProjectID(projectID string) *ImportsCancelOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ImportsCancelOptions) SetCatalogID(catalogID string) *ImportsCancelOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetHardDelete : Allow user to set HardDelete
func (_options *ImportsCancelOptions) SetHardDelete(hardDelete bool) *ImportsCancelOptions {
	_options.HardDelete = core.BoolPtr(hardDelete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ImportsCancelOptions) SetHeaders(param map[string]string) *ImportsCancelOptions {
	options.Headers = param
	return options
}

// ImportsGetOptions : The ImportsGet options.
type ImportsGetOptions struct {
	// The import identification.
	ImportID *string `json:"import_id" validate:"required,ne="`

	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportsGetOptions : Instantiate ImportsGetOptions
func (*SpaceV2) NewImportsGetOptions(importID string) *ImportsGetOptions {
	return &ImportsGetOptions{
		ImportID: core.StringPtr(importID),
	}
}

// SetImportID : Allow user to set ImportID
func (_options *ImportsGetOptions) SetImportID(importID string) *ImportsGetOptions {
	_options.ImportID = core.StringPtr(importID)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ImportsGetOptions) SetSpaceID(spaceID string) *ImportsGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ImportsGetOptions) SetProjectID(projectID string) *ImportsGetOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ImportsGetOptions) SetCatalogID(catalogID string) *ImportsGetOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ImportsGetOptions) SetHeaders(param map[string]string) *ImportsGetOptions {
	options.Headers = param
	return options
}

// ImportsListOptions : The ImportsList options.
type ImportsListOptions struct {
	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Token representing first resource.
	Start *string `json:"start,omitempty"`

	// The number of resources returned. Default value is 100. Max value is 200.
	Limit *float64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportsListOptions : Instantiate ImportsListOptions
func (*SpaceV2) NewImportsListOptions() *ImportsListOptions {
	return &ImportsListOptions{}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ImportsListOptions) SetSpaceID(spaceID string) *ImportsListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ImportsListOptions) SetProjectID(projectID string) *ImportsListOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ImportsListOptions) SetCatalogID(catalogID string) *ImportsListOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ImportsListOptions) SetStart(start string) *ImportsListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ImportsListOptions) SetLimit(limit float64) *ImportsListOptions {
	_options.Limit = core.Float64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ImportsListOptions) SetHeaders(param map[string]string) *ImportsListOptions {
	options.Headers = param
	return options
}

// ImportsStartOptions : The ImportsStart options.
type ImportsStartOptions struct {
	// Archive with assets to be imported.
	File io.ReadCloser `json:"file" validate:"required"`

	// The content type of file.
	FileContentType *string `json:"file_content_type,omitempty"`

	// Encryption key used to decrypt the sensitive data during import.
	EncryptionKey *string `json:"encryption_key,omitempty"`

	// Return resources pertaining to this space. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	SpaceID *string `json:"space_id,omitempty"`

	// Return resources pertaining to this project. Either 'space_id', 'project_id', 'catalog_id' query parameter has to be
	// given and is mandatory.
	ProjectID *string `json:"project_id,omitempty"`

	// This parameter is only supported on CPD 3.5. Return resources pertaining to this catalog. Either 'space_id',
	// 'project_id', 'catalog_id' query parameter has to be given and is mandatory.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportsStartOptions : Instantiate ImportsStartOptions
func (*SpaceV2) NewImportsStartOptions(file io.ReadCloser) *ImportsStartOptions {
	return &ImportsStartOptions{
		File: file,
	}
}

// SetFile : Allow user to set File
func (_options *ImportsStartOptions) SetFile(file io.ReadCloser) *ImportsStartOptions {
	_options.File = file
	return _options
}

// SetFileContentType : Allow user to set FileContentType
func (_options *ImportsStartOptions) SetFileContentType(fileContentType string) *ImportsStartOptions {
	_options.FileContentType = core.StringPtr(fileContentType)
	return _options
}

// SetEncryptionKey : Allow user to set EncryptionKey
func (_options *ImportsStartOptions) SetEncryptionKey(encryptionKey string) *ImportsStartOptions {
	_options.EncryptionKey = core.StringPtr(encryptionKey)
	return _options
}

// SetSpaceID : Allow user to set SpaceID
func (_options *ImportsStartOptions) SetSpaceID(spaceID string) *ImportsStartOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetProjectID : Allow user to set ProjectID
func (_options *ImportsStartOptions) SetProjectID(projectID string) *ImportsStartOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetCatalogID : Allow user to set CatalogID
func (_options *ImportsStartOptions) SetCatalogID(catalogID string) *ImportsStartOptions {
	_options.CatalogID = core.StringPtr(catalogID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ImportsStartOptions) SetHeaders(param map[string]string) *ImportsStartOptions {
	options.Headers = param
	return options
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
func (*SpaceV2) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
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

// MembersCreateOptions : The MembersCreate options.
type MembersCreateOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	Members []MemberResource `json:"members" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMembersCreateOptions : Instantiate MembersCreateOptions
func (*SpaceV2) NewMembersCreateOptions(spaceID string, members []MemberResource) *MembersCreateOptions {
	return &MembersCreateOptions{
		SpaceID: core.StringPtr(spaceID),
		Members: members,
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *MembersCreateOptions) SetSpaceID(spaceID string) *MembersCreateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetMembers : Allow user to set Members
func (_options *MembersCreateOptions) SetMembers(members []MemberResource) *MembersCreateOptions {
	_options.Members = members
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MembersCreateOptions) SetHeaders(param map[string]string) *MembersCreateOptions {
	options.Headers = param
	return options
}

// MembersDeleteOptions : The MembersDelete options.
type MembersDeleteOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	// The member identification.
	MemberID *string `json:"member_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMembersDeleteOptions : Instantiate MembersDeleteOptions
func (*SpaceV2) NewMembersDeleteOptions(spaceID string, memberID string) *MembersDeleteOptions {
	return &MembersDeleteOptions{
		SpaceID:  core.StringPtr(spaceID),
		MemberID: core.StringPtr(memberID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *MembersDeleteOptions) SetSpaceID(spaceID string) *MembersDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetMemberID : Allow user to set MemberID
func (_options *MembersDeleteOptions) SetMemberID(memberID string) *MembersDeleteOptions {
	_options.MemberID = core.StringPtr(memberID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MembersDeleteOptions) SetHeaders(param map[string]string) *MembersDeleteOptions {
	options.Headers = param
	return options
}

// MembersGetOptions : The MembersGet options.
type MembersGetOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	// The member identification.
	MemberID *string `json:"member_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMembersGetOptions : Instantiate MembersGetOptions
func (*SpaceV2) NewMembersGetOptions(spaceID string, memberID string) *MembersGetOptions {
	return &MembersGetOptions{
		SpaceID:  core.StringPtr(spaceID),
		MemberID: core.StringPtr(memberID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *MembersGetOptions) SetSpaceID(spaceID string) *MembersGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetMemberID : Allow user to set MemberID
func (_options *MembersGetOptions) SetMemberID(memberID string) *MembersGetOptions {
	_options.MemberID = core.StringPtr(memberID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MembersGetOptions) SetHeaders(param map[string]string) *MembersGetOptions {
	options.Headers = param
	return options
}

// MembersListOptions : The MembersList options.
type MembersListOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	// Token representing first resource.
	Start *string `json:"start,omitempty"`

	// The number of resources returned. Default value is 100. Max value is 200.
	Limit *float64 `json:"limit,omitempty"`

	// Include details about total number of resources. This flag is not supported on CPD 3.0.1.
	TotalCount *bool `json:"total_count,omitempty"`

	// Find the member by 'type'.
	Type *string `json:"type,omitempty"`

	// Find the member by 'role'.
	Role *string `json:"role,omitempty"`

	// Find the member by 'state'.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMembersListOptions : Instantiate MembersListOptions
func (*SpaceV2) NewMembersListOptions(spaceID string) *MembersListOptions {
	return &MembersListOptions{
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *MembersListOptions) SetSpaceID(spaceID string) *MembersListOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *MembersListOptions) SetStart(start string) *MembersListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *MembersListOptions) SetLimit(limit float64) *MembersListOptions {
	_options.Limit = core.Float64Ptr(limit)
	return _options
}

// SetTotalCount : Allow user to set TotalCount
func (_options *MembersListOptions) SetTotalCount(totalCount bool) *MembersListOptions {
	_options.TotalCount = core.BoolPtr(totalCount)
	return _options
}

// SetType : Allow user to set Type
func (_options *MembersListOptions) SetType(typeVar string) *MembersListOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetRole : Allow user to set Role
func (_options *MembersListOptions) SetRole(role string) *MembersListOptions {
	_options.Role = core.StringPtr(role)
	return _options
}

// SetState : Allow user to set State
func (_options *MembersListOptions) SetState(state string) *MembersListOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MembersListOptions) SetHeaders(param map[string]string) *MembersListOptions {
	options.Headers = param
	return options
}

// MembersUpdateOptions : The MembersUpdate options.
type MembersUpdateOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	// The member identification.
	MemberID *string `json:"member_id" validate:"required,ne="`

	// The json patch.
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMembersUpdateOptions : Instantiate MembersUpdateOptions
func (*SpaceV2) NewMembersUpdateOptions(spaceID string, memberID string, jsonPatch []JSONPatchOperation) *MembersUpdateOptions {
	return &MembersUpdateOptions{
		SpaceID:   core.StringPtr(spaceID),
		MemberID:  core.StringPtr(memberID),
		JSONPatch: jsonPatch,
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *MembersUpdateOptions) SetSpaceID(spaceID string) *MembersUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetMemberID : Allow user to set MemberID
func (_options *MembersUpdateOptions) SetMemberID(memberID string) *MembersUpdateOptions {
	_options.MemberID = core.StringPtr(memberID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *MembersUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *MembersUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MembersUpdateOptions) SetHeaders(param map[string]string) *MembersUpdateOptions {
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

// SpacesCreateOptions : The SpacesCreate options.
type SpacesCreateOptions struct {
	// Name of space.
	Name *string `json:"name" validate:"required"`

	// Description of space.
	Description *string `json:"description,omitempty"`

	// Cloud Object Storage instance is required for spaces created on Public Cloud. On private CPD installations default
	// storage is used instead. This flag is not supported on CPD.
	Storage *StorageRequest `json:"storage,omitempty"`

	// This flag is not supported on CPD.
	Compute []ComputeRequest `json:"compute,omitempty"`

	// User-defined tags associated with a space.
	Tags []string `json:"tags,omitempty"`

	// A consistent label used to identify a client that created a space. A generator must be comprised of the following
	// characters - alphanumeric, dash, underscore, space.
	Generator *string `json:"generator,omitempty"`

	// Space production and stage name.
	Stage *StageRequest `json:"stage,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSpacesCreateOptions : Instantiate SpacesCreateOptions
func (*SpaceV2) NewSpacesCreateOptions(name string) *SpacesCreateOptions {
	return &SpacesCreateOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *SpacesCreateOptions) SetName(name string) *SpacesCreateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *SpacesCreateOptions) SetDescription(description string) *SpacesCreateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetStorage : Allow user to set Storage
func (_options *SpacesCreateOptions) SetStorage(storage *StorageRequest) *SpacesCreateOptions {
	_options.Storage = storage
	return _options
}

// SetCompute : Allow user to set Compute
func (_options *SpacesCreateOptions) SetCompute(compute []ComputeRequest) *SpacesCreateOptions {
	_options.Compute = compute
	return _options
}

// SetTags : Allow user to set Tags
func (_options *SpacesCreateOptions) SetTags(tags []string) *SpacesCreateOptions {
	_options.Tags = tags
	return _options
}

// SetGenerator : Allow user to set Generator
func (_options *SpacesCreateOptions) SetGenerator(generator string) *SpacesCreateOptions {
	_options.Generator = core.StringPtr(generator)
	return _options
}

// SetStage : Allow user to set Stage
func (_options *SpacesCreateOptions) SetStage(stage *StageRequest) *SpacesCreateOptions {
	_options.Stage = stage
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SpacesCreateOptions) SetHeaders(param map[string]string) *SpacesCreateOptions {
	options.Headers = param
	return options
}

// SpacesDeleteOptions : The SpacesDelete options.
type SpacesDeleteOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSpacesDeleteOptions : Instantiate SpacesDeleteOptions
func (*SpaceV2) NewSpacesDeleteOptions(spaceID string) *SpacesDeleteOptions {
	return &SpacesDeleteOptions{
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *SpacesDeleteOptions) SetSpaceID(spaceID string) *SpacesDeleteOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SpacesDeleteOptions) SetHeaders(param map[string]string) *SpacesDeleteOptions {
	options.Headers = param
	return options
}

// SpacesGetOptions : The SpacesGet options.
type SpacesGetOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	// A list of comma-separated space sections to include in the query results. Example: '?include=members'.
	//
	// Available fields:
	//  * members (returns up to 100 members)
	//  * nothing (does not return space entity and metadata).
	Include *string `json:"include,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSpacesGetOptions : Instantiate SpacesGetOptions
func (*SpaceV2) NewSpacesGetOptions(spaceID string) *SpacesGetOptions {
	return &SpacesGetOptions{
		SpaceID: core.StringPtr(spaceID),
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *SpacesGetOptions) SetSpaceID(spaceID string) *SpacesGetOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetInclude : Allow user to set Include
func (_options *SpacesGetOptions) SetInclude(include string) *SpacesGetOptions {
	_options.Include = core.StringPtr(include)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SpacesGetOptions) SetHeaders(param map[string]string) *SpacesGetOptions {
	options.Headers = param
	return options
}

// SpacesListOptions : The SpacesList options.
type SpacesListOptions struct {
	// Token representing first resource.
	Start *string `json:"start,omitempty"`

	// The number of resources returned. Default value is 100. Max value is 200.
	Limit *float64 `json:"limit,omitempty"`

	// Include details about total number of resources. This flag is not supported on CPD 3.0.1.
	TotalCount *bool `json:"total_count,omitempty"`

	// Comma-separated list of ids to be returned. This flag is not supported on CPD 3.0.1.
	ID *string `json:"id,omitempty"`

	// A list of comma-separated, user-defined tags to use to filter the query results.
	Tags *string `json:"tags,omitempty"`

	// A list of comma-separated space sections to include in the query results. Example: '?include=members'.
	//
	// Available fields:
	//  * members (returns up to 100 members)
	//  * nothing (does not return space entity and metadata).
	Include *string `json:"include,omitempty"`

	// Filters the result list to only include spaces where the user with a matching user id is a member.
	Member *string `json:"member,omitempty"`

	// Must be used in conjunction with the member query parameter. Filters the result set to include only spaces where the
	// specified member has one of the roles specified.
	//
	// Values:
	//
	//   * admin
	//   * editor
	//   * viewer.
	Roles *string `json:"roles,omitempty"`

	// Filtering by bss_account_id is allowed only for accredited services.
	BssAccountID *string `json:"bss_account_id,omitempty"`

	// Filters the result list to only include space with specified name.
	Name *string `json:"name,omitempty"`

	// Filters the result list to only include spaces with specified compute.crn.
	ComputeCrn *string `json:"compute.crn,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSpacesListOptions : Instantiate SpacesListOptions
func (*SpaceV2) NewSpacesListOptions() *SpacesListOptions {
	return &SpacesListOptions{}
}

// SetStart : Allow user to set Start
func (_options *SpacesListOptions) SetStart(start string) *SpacesListOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *SpacesListOptions) SetLimit(limit float64) *SpacesListOptions {
	_options.Limit = core.Float64Ptr(limit)
	return _options
}

// SetTotalCount : Allow user to set TotalCount
func (_options *SpacesListOptions) SetTotalCount(totalCount bool) *SpacesListOptions {
	_options.TotalCount = core.BoolPtr(totalCount)
	return _options
}

// SetID : Allow user to set ID
func (_options *SpacesListOptions) SetID(id string) *SpacesListOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetTags : Allow user to set Tags
func (_options *SpacesListOptions) SetTags(tags string) *SpacesListOptions {
	_options.Tags = core.StringPtr(tags)
	return _options
}

// SetInclude : Allow user to set Include
func (_options *SpacesListOptions) SetInclude(include string) *SpacesListOptions {
	_options.Include = core.StringPtr(include)
	return _options
}

// SetMember : Allow user to set Member
func (_options *SpacesListOptions) SetMember(member string) *SpacesListOptions {
	_options.Member = core.StringPtr(member)
	return _options
}

// SetRoles : Allow user to set Roles
func (_options *SpacesListOptions) SetRoles(roles string) *SpacesListOptions {
	_options.Roles = core.StringPtr(roles)
	return _options
}

// SetBssAccountID : Allow user to set BssAccountID
func (_options *SpacesListOptions) SetBssAccountID(bssAccountID string) *SpacesListOptions {
	_options.BssAccountID = core.StringPtr(bssAccountID)
	return _options
}

// SetName : Allow user to set Name
func (_options *SpacesListOptions) SetName(name string) *SpacesListOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetComputeCrn : Allow user to set ComputeCrn
func (_options *SpacesListOptions) SetComputeCrn(computeCrn string) *SpacesListOptions {
	_options.ComputeCrn = core.StringPtr(computeCrn)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SpacesListOptions) SetHeaders(param map[string]string) *SpacesListOptions {
	options.Headers = param
	return options
}

// SpacesUpdateOptions : The SpacesUpdate options.
type SpacesUpdateOptions struct {
	// The space identification.
	SpaceID *string `json:"space_id" validate:"required,ne="`

	// Input payload for the space.
	JSONPatch []JSONPatchOperation `json:"json-patch" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSpacesUpdateOptions : Instantiate SpacesUpdateOptions
func (*SpaceV2) NewSpacesUpdateOptions(spaceID string, jsonPatch []JSONPatchOperation) *SpacesUpdateOptions {
	return &SpacesUpdateOptions{
		SpaceID:   core.StringPtr(spaceID),
		JSONPatch: jsonPatch,
	}
}

// SetSpaceID : Allow user to set SpaceID
func (_options *SpacesUpdateOptions) SetSpaceID(spaceID string) *SpacesUpdateOptions {
	_options.SpaceID = core.StringPtr(spaceID)
	return _options
}

// SetJSONPatch : Allow user to set JSONPatch
func (_options *SpacesUpdateOptions) SetJSONPatch(jsonPatch []JSONPatchOperation) *SpacesUpdateOptions {
	_options.JSONPatch = jsonPatch
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SpacesUpdateOptions) SetHeaders(param map[string]string) *SpacesUpdateOptions {
	options.Headers = param
	return options
}

// ComputeEntity : ComputeEntity struct
type ComputeEntity struct {
	// Display name of the compute instance associated with the space.
	Name *string `json:"name" validate:"required"`

	// The Cloud Resource Name (CRN) for the compute service associated with the space.
	Crn *string `json:"crn" validate:"required"`

	// Unique GUID of the associated space compute.
	Guid *string `json:"guid" validate:"required"`

	// The type of compute associated with the space.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the ComputeEntity.Type property.
// The type of compute associated with the space.
const (
	ComputeEntity_Type_MachineLearning = "machine_learning"
	ComputeEntity_Type_Openscale       = "openscale"
)

// UnmarshalComputeEntity unmarshals an instance of ComputeEntity from the specified map of raw messages.
func UnmarshalComputeEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComputeEntity)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "guid", &obj.Guid)
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

// ComputeRequest : ComputeRequest struct
type ComputeRequest struct {
	// Display name of the compute instance associated with the space.
	Name *string `json:"name" validate:"required"`

	// The Cloud Resource Name (CRN) for the compute service associated with the space.
	Crn *string `json:"crn" validate:"required"`
}

// NewComputeRequest : Instantiate ComputeRequest (Generic Model Constructor)
func (*SpaceV2) NewComputeRequest(name string, crn string) (_model *ComputeRequest, err error) {
	_model = &ComputeRequest{
		Name: core.StringPtr(name),
		Crn:  core.StringPtr(crn),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalComputeRequest unmarshals an instance of ComputeRequest from the specified map of raw messages.
func UnmarshalComputeRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ComputeRequest)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
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

// ExportAssets : ExportAssets struct
type ExportAssets struct {
	AllAssets *bool `json:"all_assets,omitempty"`

	AssetTypes []string `json:"asset_types,omitempty"`

	AssetIds []string `json:"asset_ids,omitempty"`
}

// UnmarshalExportAssets unmarshals an instance of ExportAssets from the specified map of raw messages.
func UnmarshalExportAssets(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportAssets)
	err = core.UnmarshalPrimitive(m, "all_assets", &obj.AllAssets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_types", &obj.AssetTypes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "asset_ids", &obj.AssetIds)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportEntity : ExportEntity struct
type ExportEntity struct {
	Assets *ExportAssets `json:"assets,omitempty"`

	Status *ImportExportStatus `json:"status,omitempty"`
}

// UnmarshalExportEntity unmarshals an instance of ExportEntity from the specified map of raw messages.
func UnmarshalExportEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportEntity)
	err = core.UnmarshalModel(m, "assets", &obj.Assets, UnmarshalExportAssets)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalImportExportStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportResource : ExportResource struct
type ExportResource struct {
	// Common metadata for a import/export.
	Metadata *ImportExportMeta `json:"metadata,omitempty"`

	Entity *ExportEntity `json:"entity,omitempty"`
}

// UnmarshalExportResource unmarshals an instance of ExportResource from the specified map of raw messages.
func UnmarshalExportResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalImportExportMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalExportEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportResources : Information for paging when queerying resources.
type ExportResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first,omitempty"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	Resources []ExportResource `json:"resources,omitempty"`
}

// UnmarshalExportResources unmarshals an instance of ExportResources from the specified map of raw messages.
func UnmarshalExportResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportResources)
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
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalExportResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ExportResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// ImportEntity : ImportEntity struct
type ImportEntity struct {
	Status *ImportExportStatus `json:"status,omitempty"`
}

// UnmarshalImportEntity unmarshals an instance of ImportEntity from the specified map of raw messages.
func UnmarshalImportEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportEntity)
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalImportExportStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportExportMeta : Common metadata for a import/export.
type ImportExportMeta struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The url of the resource.
	URL *string `json:"url" validate:"required"`

	// The user id which created this resource.
	CreatorID *string `json:"creator_id" validate:"required"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	SpaceID *string `json:"space_id,omitempty"`

	ProjectID *string `json:"project_id,omitempty"`

	CatalogID *string `json:"catalog_id,omitempty"`
}

// UnmarshalImportExportMeta unmarshals an instance of ImportExportMeta from the specified map of raw messages.
func UnmarshalImportExportMeta(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportExportMeta)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creator_id", &obj.CreatorID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
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
	err = core.UnmarshalPrimitive(m, "space_id", &obj.SpaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportExportStatus : ImportExportStatus struct
type ImportExportStatus struct {
	// The state of the process.
	State *string `json:"state,omitempty"`

	// A value between 0 and 100 indicating the % complete. This may be missing in cases where it is not possible to obtain
	// it.
	Progress *float64 `json:"progress,omitempty"`

	// The data returned when an error is encountered.
	Failure *Error `json:"failure,omitempty"`
}

// Constants associated with the ImportExportStatus.State property.
// The state of the process.
const (
	ImportExportStatus_State_Canceled  = "canceled"
	ImportExportStatus_State_Completed = "completed"
	ImportExportStatus_State_Deleting  = "deleting"
	ImportExportStatus_State_Failed    = "failed"
	ImportExportStatus_State_Pending   = "pending"
	ImportExportStatus_State_Running   = "running"
)

// UnmarshalImportExportStatus unmarshals an instance of ImportExportStatus from the specified map of raw messages.
func UnmarshalImportExportStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportExportStatus)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "progress", &obj.Progress)
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

// ImportResource : ImportResource struct
type ImportResource struct {
	// Common metadata for a import/export.
	Metadata *ImportExportMeta `json:"metadata,omitempty"`

	Entity *ImportEntity `json:"entity,omitempty"`
}

// UnmarshalImportResource unmarshals an instance of ImportResource from the specified map of raw messages.
func UnmarshalImportResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalImportExportMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalImportEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportResources : Information for paging when queerying resources.
type ImportResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first,omitempty"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	Resources []ImportResource `json:"resources,omitempty"`
}

// UnmarshalImportResources unmarshals an instance of ImportResources from the specified map of raw messages.
func UnmarshalImportResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportResources)
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
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalImportResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ImportResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// MemberArray : MemberArray struct
type MemberArray struct {
	Resources []MemberResource `json:"resources" validate:"required"`
}

// UnmarshalMemberArray unmarshals an instance of MemberArray from the specified map of raw messages.
func UnmarshalMemberArray(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MemberArray)
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalMemberResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MemberPatchRequest : MemberPatchRequest struct
type MemberPatchRequest struct {
	// The role of the space member.
	Role *string `json:"role" validate:"required"`

	// The state of the space member. This field is only supported for members of type 'user'.
	State *string `json:"state,omitempty"`
}

// Constants associated with the MemberPatchRequest.Role property.
// The role of the space member.
const (
	MemberPatchRequest_Role_Admin  = "admin"
	MemberPatchRequest_Role_Editor = "editor"
	MemberPatchRequest_Role_Viewer = "viewer"
)

// Constants associated with the MemberPatchRequest.State property.
// The state of the space member. This field is only supported for members of type 'user'.
const (
	MemberPatchRequest_State_Active  = "active"
	MemberPatchRequest_State_Pending = "pending"
)

// UnmarshalMemberPatchRequest unmarshals an instance of MemberPatchRequest from the specified map of raw messages.
func UnmarshalMemberPatchRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MemberPatchRequest)
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
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

func (*SpaceV2) NewMemberPatchRequestPatch(memberPatchRequest *MemberPatchRequest) (_patch []JSONPatchOperation) {
	if memberPatchRequest.Role != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/role"),
			Value: memberPatchRequest.Role,
		})
	}
	if memberPatchRequest.State != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/state"),
			Value: memberPatchRequest.State,
		})
	}
	return
}

// MemberResource : MemberResource struct
type MemberResource struct {
	// The IAM ID of the space member. This field is required for space members with the 'active' state.
	ID *string `json:"id,omitempty"`

	// The role of the space member.
	Role *string `json:"role" validate:"required"`

	// The state of the space member. This field is only supported for members of type 'user'.
	State *string `json:"state,omitempty"`

	// The space member type.
	Type *string `json:"type,omitempty"`
}

// Constants associated with the MemberResource.Role property.
// The role of the space member.
const (
	MemberResource_Role_Admin  = "admin"
	MemberResource_Role_Editor = "editor"
	MemberResource_Role_Viewer = "viewer"
)

// Constants associated with the MemberResource.State property.
// The state of the space member. This field is only supported for members of type 'user'.
const (
	MemberResource_State_Active  = "active"
	MemberResource_State_Pending = "pending"
)

// Constants associated with the MemberResource.Type property.
// The space member type.
const (
	MemberResource_Type_Group   = "group"
	MemberResource_Type_Service = "service"
	MemberResource_Type_User    = "user"
)

// NewMemberResource : Instantiate MemberResource (Generic Model Constructor)
func (*SpaceV2) NewMemberResource(role string) (_model *MemberResource, err error) {
	_model = &MemberResource{
		Role: core.StringPtr(role),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMemberResource unmarshals an instance of MemberResource from the specified map of raw messages.
func UnmarshalMemberResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MemberResource)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role", &obj.Role)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
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

// MemberResources : Information for paging when queerying resources.
type MemberResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first,omitempty"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	Resources []MemberResource `json:"resources" validate:"required"`
}

// UnmarshalMemberResources unmarshals an instance of MemberResources from the specified map of raw messages.
func UnmarshalMemberResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MemberResources)
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
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalMemberResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *MemberResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// ResourceMeta : Common metadata for a resource.
type ResourceMeta struct {
	// The id of the resource.
	ID *string `json:"id" validate:"required"`

	// The url of the resource.
	URL *string `json:"url" validate:"required"`

	// The user id which created this resource.
	CreatorID *string `json:"creator_id" validate:"required"`

	// The time when the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The time when the resource was updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`
}

// UnmarshalResourceMeta unmarshals an instance of ResourceMeta from the specified map of raw messages.
func UnmarshalResourceMeta(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceMeta)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creator_id", &obj.CreatorID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
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

// Scope : The scope definition.
type Scope struct {
	// The BSS Account ID associated with the space.
	BssAccountID *string `json:"bss_account_id" validate:"required"`
}

// UnmarshalScope unmarshals an instance of Scope from the specified map of raw messages.
func UnmarshalScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scope)
	err = core.UnmarshalPrimitive(m, "bss_account_id", &obj.BssAccountID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpaceEntity : SpaceEntity struct
type SpaceEntity struct {
	Name *string `json:"name" validate:"required"`

	Description *string `json:"description,omitempty"`

	// The scope definition.
	Scope *Scope `json:"scope" validate:"required"`

	Storage *Storage `json:"storage,omitempty"`

	Compute []ComputeEntity `json:"compute,omitempty"`

	Members []MemberResource `json:"members,omitempty"`

	// User-defined tags associated with a space.
	Tags []string `json:"tags,omitempty"`

	// The status of the space.
	Status *SpaceStatus `json:"status" validate:"required"`

	// A consistent label used to identify a client that created a space.
	Generator *string `json:"generator,omitempty"`

	// Space production and stage name.
	Stage *StageEntity `json:"stage" validate:"required"`
}

// UnmarshalSpaceEntity unmarshals an instance of SpaceEntity from the specified map of raw messages.
func UnmarshalSpaceEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpaceEntity)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scope", &obj.Scope, UnmarshalScope)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "storage", &obj.Storage, UnmarshalStorage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "compute", &obj.Compute, UnmarshalComputeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalMemberResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalSpaceStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "generator", &obj.Generator)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stage", &obj.Stage, UnmarshalStageEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpacePatchRequest : SpacePatchRequest struct
type SpacePatchRequest struct {
	// Updated name.
	Name *string `json:"name,omitempty"`

	// Updated description.
	Description *string `json:"description,omitempty"`

	// Updated compute definition.
	Compute []ComputeEntity `json:"compute,omitempty"`

	// Updated stage name.
	Stagename *string `json:"stage/name,omitempty"`
}

// UnmarshalSpacePatchRequest unmarshals an instance of SpacePatchRequest from the specified map of raw messages.
func UnmarshalSpacePatchRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpacePatchRequest)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "compute", &obj.Compute, UnmarshalComputeEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stage/name", &obj.Stagename)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*SpaceV2) NewSpacePatchRequestPatch(spacePatchRequest *SpacePatchRequest) (_patch []JSONPatchOperation) {
	if spacePatchRequest.Name != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/name"),
			Value: spacePatchRequest.Name,
		})
	}
	if spacePatchRequest.Description != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/description"),
			Value: spacePatchRequest.Description,
		})
	}
	if spacePatchRequest.Compute != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/compute"),
			Value: spacePatchRequest.Compute,
		})
	}
	if spacePatchRequest.Stagename != nil {
		_patch = append(_patch, JSONPatchOperation{
			Op:    core.StringPtr(JSONPatchOperation_Op_Add),
			Path:  core.StringPtr("/stage/name"),
			Value: spacePatchRequest.Stagename,
		})
	}
	return
}

// SpaceResource : SpaceResource struct
type SpaceResource struct {
	// Common metadata for a resource.
	Metadata *ResourceMeta `json:"metadata" validate:"required"`

	Entity *SpaceEntity `json:"entity" validate:"required"`
}

// UnmarshalSpaceResource unmarshals an instance of SpaceResource from the specified map of raw messages.
func UnmarshalSpaceResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpaceResource)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalResourceMeta)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalSpaceEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpaceResources : Information for paging when queerying resources.
type SpaceResources struct {
	// Computed explicitly only when 'total_count=true' query parameter is present. This is in order to avoid performance
	// penalties.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of items to return in each page.
	Limit *int64 `json:"limit,omitempty"`

	// The reference to the first item in the current page.
	First *PaginationFirst `json:"first,omitempty"`

	// A reference to the first item of the next page, if any.
	Next *PaginationNext `json:"next,omitempty"`

	Resources []SpaceResource `json:"resources" validate:"required"`
}

// UnmarshalSpaceResources unmarshals an instance of SpaceResources from the specified map of raw messages.
func UnmarshalSpaceResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpaceResources)
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
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalSpaceResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *SpaceResources) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	start, err := core.GetQueryParam(resp.Next.Href, "start")
	if err != nil || start == nil {
		return nil, err
	}
	return start, nil
}

// SpaceStatus : The status of the space.
type SpaceStatus struct {
	State *string `json:"state" validate:"required"`

	// The data returned when an error is encountered.
	Failure *Error `json:"failure,omitempty"`
}

// Constants associated with the SpaceStatus.State property.
const (
	SpaceStatus_State_Active    = "active"
	SpaceStatus_State_Deleted   = "deleted"
	SpaceStatus_State_Deleting  = "deleting"
	SpaceStatus_State_Error     = "error"
	SpaceStatus_State_Preparing = "preparing"
)

// UnmarshalSpaceStatus unmarshals an instance of SpaceStatus from the specified map of raw messages.
func UnmarshalSpaceStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpaceStatus)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
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

// StageEntity : Space production and stage name.
type StageEntity struct {
	// A boolean flag which specifies if a space is production or not. It cannot be modified.
	Production *bool `json:"production" validate:"required"`

	// Stage name. E.g. development, test, pre-production, production.
	Name *string `json:"name,omitempty"`
}

// UnmarshalStageEntity unmarshals an instance of StageEntity from the specified map of raw messages.
func UnmarshalStageEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StageEntity)
	err = core.UnmarshalPrimitive(m, "production", &obj.Production)
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

// StageRequest : Space production and stage name.
type StageRequest struct {
	// A boolean flag which specifies if a space is production or not. It cannot be modified.
	Production *bool `json:"production,omitempty"`

	// Stage name. E.g. development, test, pre-production, production.
	Name *string `json:"name,omitempty"`
}

// UnmarshalStageRequest unmarshals an instance of StageRequest from the specified map of raw messages.
func UnmarshalStageRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StageRequest)
	err = core.UnmarshalPrimitive(m, "production", &obj.Production)
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

// Storage : Storage struct
type Storage struct {
	// The associated storage type.
	Type *string `json:"type" validate:"required"`

	Properties *StorageProperties `json:"properties" validate:"required"`
}

// Constants associated with the Storage.Type property.
// The associated storage type.
const (
	Storage_Type_BmcosObjectStorage = "bmcos_object_storage"
)

// UnmarshalStorage unmarshals an instance of Storage from the specified map of raw messages.
func UnmarshalStorage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Storage)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "properties", &obj.Properties, UnmarshalStorageProperties)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageCredentials : StorageCredentials struct
type StorageCredentials struct {
	// An api key for specific user role.
	ApiKey *string `json:"api_key" validate:"required"`

	// Service iam-id for specific user role.
	ServiceID *string `json:"service_id" validate:"required"`

	// HMAC access key id for associated cloud object storage and specific user role.
	AccessKeyID *string `json:"access_key_id,omitempty"`

	// HMAC secret access key for associated cloud object storage and specific user role.
	SecretAccessKey *string `json:"secret_access_key,omitempty"`

	// Resource key CRN. Must be provided with other HMAC credentials for non-admin entries.
	ResourceKeyCrn *string `json:"resource_key_crn,omitempty"`
}

// UnmarshalStorageCredentials unmarshals an instance of StorageCredentials from the specified map of raw messages.
func UnmarshalStorageCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageCredentials)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_id", &obj.ServiceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "access_key_id", &obj.AccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret_access_key", &obj.SecretAccessKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_key_crn", &obj.ResourceKeyCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageCredentialsAdmin : StorageCredentialsAdmin struct
type StorageCredentialsAdmin struct {
	// An api key for specific user role.
	ApiKey *string `json:"api_key" validate:"required"`

	// Service iam-id for specific user role.
	ServiceID *string `json:"service_id" validate:"required"`

	// HMAC access key id for associated cloud object storage and specific user role.
	AccessKeyID *string `json:"access_key_id,omitempty"`

	// HMAC secret access key for associated cloud object storage and specific user role.
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
}

// UnmarshalStorageCredentialsAdmin unmarshals an instance of StorageCredentialsAdmin from the specified map of raw messages.
func UnmarshalStorageCredentialsAdmin(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageCredentialsAdmin)
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_id", &obj.ServiceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "access_key_id", &obj.AccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret_access_key", &obj.SecretAccessKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageCredentialsFull : Credentials required to access the bucket.
type StorageCredentialsFull struct {
	Admin *StorageCredentialsAdmin `json:"admin,omitempty"`

	Editor *StorageCredentials `json:"editor,omitempty"`

	Viewer *StorageCredentials `json:"viewer" validate:"required"`
}

// UnmarshalStorageCredentialsFull unmarshals an instance of StorageCredentialsFull from the specified map of raw messages.
func UnmarshalStorageCredentialsFull(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageCredentialsFull)
	err = core.UnmarshalModel(m, "admin", &obj.Admin, UnmarshalStorageCredentialsAdmin)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "editor", &obj.Editor, UnmarshalStorageCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "viewer", &obj.Viewer, UnmarshalStorageCredentials)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageProperties : StorageProperties struct
type StorageProperties struct {
	// A cloud resource name of the Cloud Object Storage instance.
	ResourceCrn *string `json:"resource_crn" validate:"required"`

	// Guid of the Cloud Object Storage instance.
	Guid *string `json:"guid" validate:"required"`

	// The bucket name to connect to.
	BucketName *string `json:"bucket_name" validate:"required"`

	// The bucket region to connect to.
	BucketRegion *string `json:"bucket_region" validate:"required"`

	// The endpoint URL to use for connecting clients to the object store.
	EndpointURL *string `json:"endpoint_url" validate:"required"`

	// Used to indicate if the bucket created for the project is encrypted using key protect.
	KeyProtect *bool `json:"key_protect,omitempty"`

	// Credentials required to access the bucket.
	Credentials *StorageCredentialsFull `json:"credentials" validate:"required"`
}

// UnmarshalStorageProperties unmarshals an instance of StorageProperties from the specified map of raw messages.
func UnmarshalStorageProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageProperties)
	err = core.UnmarshalPrimitive(m, "resource_crn", &obj.ResourceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "guid", &obj.Guid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_name", &obj.BucketName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_region", &obj.BucketRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint_url", &obj.EndpointURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key_protect", &obj.KeyProtect)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalStorageCredentialsFull)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StorageRequest : Cloud Object Storage instance is required for spaces created on Public Cloud. On private CPD installations default
// storage is used instead. This flag is not supported on CPD.
type StorageRequest struct {
	// A cloud resource name of the Cloud Object Storage instance.
	ResourceCrn *string `json:"resource_crn" validate:"required"`

	// Set to true of the COS instance is delegated by the account admin.
	Delegated *bool `json:"delegated,omitempty"`
}

// NewStorageRequest : Instantiate StorageRequest (Generic Model Constructor)
func (*SpaceV2) NewStorageRequest(resourceCrn string) (_model *StorageRequest, err error) {
	_model = &StorageRequest{
		ResourceCrn: core.StringPtr(resourceCrn),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalStorageRequest unmarshals an instance of StorageRequest from the specified map of raw messages.
func UnmarshalStorageRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StorageRequest)
	err = core.UnmarshalPrimitive(m, "resource_crn", &obj.ResourceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "delegated", &obj.Delegated)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpacesListPager can be used to simplify the use of the "SpacesList" method.
type SpacesListPager struct {
	hasNext     bool
	options     *SpacesListOptions
	client      *SpaceV2
	pageContext struct {
		next *string
	}
}

// NewSpacesListPager returns a new SpacesListPager instance.
func (space *SpaceV2) NewSpacesListPager(options *SpacesListOptions) (pager *SpacesListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy SpacesListOptions = *options
	pager = &SpacesListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  space,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *SpacesListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *SpacesListPager) GetNextWithContext(ctx context.Context) (page []SpaceResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.SpacesListWithContext(ctx, pager.options)
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
func (pager *SpacesListPager) GetAllWithContext(ctx context.Context) (allItems []SpaceResource, err error) {
	for pager.HasNext() {
		var nextPage []SpaceResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *SpacesListPager) GetNext() (page []SpaceResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *SpacesListPager) GetAll() (allItems []SpaceResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// MembersListPager can be used to simplify the use of the "MembersList" method.
type MembersListPager struct {
	hasNext     bool
	options     *MembersListOptions
	client      *SpaceV2
	pageContext struct {
		next *string
	}
}

// NewMembersListPager returns a new MembersListPager instance.
func (space *SpaceV2) NewMembersListPager(options *MembersListOptions) (pager *MembersListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy MembersListOptions = *options
	pager = &MembersListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  space,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *MembersListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *MembersListPager) GetNextWithContext(ctx context.Context) (page []MemberResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.MembersListWithContext(ctx, pager.options)
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
func (pager *MembersListPager) GetAllWithContext(ctx context.Context) (allItems []MemberResource, err error) {
	for pager.HasNext() {
		var nextPage []MemberResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *MembersListPager) GetNext() (page []MemberResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *MembersListPager) GetAll() (allItems []MemberResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ExportsListPager can be used to simplify the use of the "ExportsList" method.
type ExportsListPager struct {
	hasNext     bool
	options     *ExportsListOptions
	client      *SpaceV2
	pageContext struct {
		next *string
	}
}

// NewExportsListPager returns a new ExportsListPager instance.
func (space *SpaceV2) NewExportsListPager(options *ExportsListOptions) (pager *ExportsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ExportsListOptions = *options
	pager = &ExportsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  space,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ExportsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ExportsListPager) GetNextWithContext(ctx context.Context) (page []ExportResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ExportsListWithContext(ctx, pager.options)
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
func (pager *ExportsListPager) GetAllWithContext(ctx context.Context) (allItems []ExportResource, err error) {
	for pager.HasNext() {
		var nextPage []ExportResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ExportsListPager) GetNext() (page []ExportResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ExportsListPager) GetAll() (allItems []ExportResource, err error) {
	return pager.GetAllWithContext(context.Background())
}

// ImportsListPager can be used to simplify the use of the "ImportsList" method.
type ImportsListPager struct {
	hasNext     bool
	options     *ImportsListOptions
	client      *SpaceV2
	pageContext struct {
		next *string
	}
}

// NewImportsListPager returns a new ImportsListPager instance.
func (space *SpaceV2) NewImportsListPager(options *ImportsListOptions) (pager *ImportsListPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ImportsListOptions = *options
	pager = &ImportsListPager{
		hasNext: true,
		options: &optionsCopy,
		client:  space,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ImportsListPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ImportsListPager) GetNextWithContext(ctx context.Context) (page []ImportResource, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ImportsListWithContext(ctx, pager.options)
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
func (pager *ImportsListPager) GetAllWithContext(ctx context.Context) (allItems []ImportResource, err error) {
	for pager.HasNext() {
		var nextPage []ImportResource
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ImportsListPager) GetNext() (page []ImportResource, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ImportsListPager) GetAll() (allItems []ImportResource, err error) {
	return pager.GetAllWithContext(context.Background())
}
