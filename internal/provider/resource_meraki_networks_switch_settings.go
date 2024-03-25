// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0
package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchSettingsResource{}
)

func NewNetworksSwitchSettingsResource() resource.Resource {
	return &NetworksSwitchSettingsResource{}
}

type NetworksSwitchSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_settings"
}

func (r *NetworksSwitchSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"power_exceptions": schema.SetNestedAttribute{
				MarkdownDescription: `Exceptions on a per switch basis to "useCombinedPower"`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"power_type": schema.StringAttribute{
							MarkdownDescription: `Per switch exception (combined, redundant, useNetworkSetting)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the switch`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"use_combined_power": schema.BoolAttribute{
				MarkdownDescription: `The use Combined Power as the default behavior of secondary power supplies on supported devices.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: `Management VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksSwitchSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchSettingsRs

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	//Item
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchSettings(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchSettings(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchSettings(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchSettings",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchSettingsRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchSettings(vvNetworkID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchSettings",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchSettingsRs struct {
	NetworkID        types.String                                               `tfsdk:"network_id"`
	PowerExceptions  *[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptionsRs `tfsdk:"power_exceptions"`
	UseCombinedPower types.Bool                                                 `tfsdk:"use_combined_power"`
	VLAN             types.Int64                                                `tfsdk:"vlan"`
}

type ResponseSwitchGetNetworkSwitchSettingsPowerExceptionsRs struct {
	PowerType types.String `tfsdk:"power_type"`
	Serial    types.String `tfsdk:"serial"`
}

// FromBody
func (r *NetworksSwitchSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchSettings {
	var requestSwitchUpdateNetworkSwitchSettingsPowerExceptions []merakigosdk.RequestSwitchUpdateNetworkSwitchSettingsPowerExceptions
	if r.PowerExceptions != nil {
		for _, rItem1 := range *r.PowerExceptions {
			powerType := rItem1.PowerType.ValueString()
			serial := rItem1.Serial.ValueString()
			requestSwitchUpdateNetworkSwitchSettingsPowerExceptions = append(requestSwitchUpdateNetworkSwitchSettingsPowerExceptions, merakigosdk.RequestSwitchUpdateNetworkSwitchSettingsPowerExceptions{
				PowerType: powerType,
				Serial:    serial,
			})
		}
	}
	useCombinedPower := new(bool)
	if !r.UseCombinedPower.IsUnknown() && !r.UseCombinedPower.IsNull() {
		*useCombinedPower = r.UseCombinedPower.ValueBool()
	} else {
		useCombinedPower = nil
	}
	vLAN := new(int64)
	if !r.VLAN.IsUnknown() && !r.VLAN.IsNull() {
		*vLAN = r.VLAN.ValueInt64()
	} else {
		vLAN = nil
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchSettings{
		PowerExceptions: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchSettingsPowerExceptions {
			if len(requestSwitchUpdateNetworkSwitchSettingsPowerExceptions) > 0 {
				return &requestSwitchUpdateNetworkSwitchSettingsPowerExceptions
			}
			return nil
		}(),
		UseCombinedPower: useCombinedPower,
		VLAN:             int64ToIntPointer(vLAN),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchSettingsItemToBodyRs(state NetworksSwitchSettingsRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchSettings, is_read bool) NetworksSwitchSettingsRs {
	itemState := NetworksSwitchSettingsRs{
		PowerExceptions: func() *[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptionsRs {
			if response.PowerExceptions != nil {
				result := make([]ResponseSwitchGetNetworkSwitchSettingsPowerExceptionsRs, len(*response.PowerExceptions))
				for i, powerExceptions := range *response.PowerExceptions {
					result[i] = ResponseSwitchGetNetworkSwitchSettingsPowerExceptionsRs{
						PowerType: types.StringValue(powerExceptions.PowerType),
						Serial:    types.StringValue(powerExceptions.Serial),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchSettingsPowerExceptionsRs{}
		}(),
		UseCombinedPower: func() types.Bool {
			if response.UseCombinedPower != nil {
				return types.BoolValue(*response.UseCombinedPower)
			}
			return types.Bool{}
		}(),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchSettingsRs)
}
