package provider

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"time"

	"terraform-provider-ibmcpd/internal/go-sdk/client"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonopenscalev2"
	"terraform-provider-ibmcpd/internal/utils"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &recordResource{}
	_ resource.ResourceWithConfigure = &recordResource{}
)

type recordResource struct {
	client *client.Client
}

type recordResourceModel struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
	FilePath       types.String `tfsdk:"file_path"`
	Type           types.String `tfsdk:"type"`
}

func NewRecordResource() resource.Resource {
	return &recordResource{}
}

func (r *recordResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.Client)
}

func (r *recordResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_record"
}

func (r *recordResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"subscription_id": schema.StringAttribute{
				Required: true,
			},
			"file_path": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *recordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan recordResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WOS Client", err.Error())
		return
	}

	file, err := os.Open(plan.FilePath.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to open record file", err.Error())
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read record file", err.Error())
		return
	}
	var jsonFeedback map[string][]interface{}
	json.Unmarshal(content, &jsonFeedback)
	fields := make([]string, len(jsonFeedback["fields"]))
	for i, v := range jsonFeedback["fields"] {
		fields[i] = v.(string)
	}
	objValues, err := json.Marshal(jsonFeedback["values"])
	if err != nil {
		panic(err)
	}
	var values [][]interface{}
	json.Unmarshal(objValues, &values)

	var dataSetID *string

	for i := 1; i < 20; i++ {
		time.Sleep(5 * time.Second)
		resultDataSets, response, err := wosClient.DataSetsList(&watsonopenscalev2.DataSetsListOptions{
			TargetTargetID:   core.StringPtr(plan.SubscriptionID.ValueString()),
			Type:             core.StringPtr(plan.Type.ValueString()),
			TargetTargetType: core.StringPtr("subscription"),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Listing Datasets", "Could not list datasets, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to List Datasets", err.Error())
			return
		}
		if len(resultDataSets.DataSets) > 0 {
			dataSetID = resultDataSets.DataSets[0].Metadata.ID
			if *resultDataSets.DataSets[0].Entity.Status.State == "active" {
				break
			}
		}
	}

	_, response, err := wosClient.RecordsAdd(&watsonopenscalev2.RecordsAddOptions{
		DataSetID: dataSetID,
		DatasetRecordsPayloadItem: []watsonopenscalev2.DatasetRecordsPayloadItemIntf{
			&watsonopenscalev2.DatasetRecordsPayloadItemJsList{
				Fields: fields,
				Values: values,
			},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Adding Records", "Could not add records, unexpected error: "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to List Datasets", err.Error())
		return
	}

	if _, err := os.Stat(plan.FilePath.ValueString()); err == nil || os.IsExist(err) {
		os.Remove(plan.FilePath.ValueString())
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *recordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *recordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *recordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
