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
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsBonjourForwardingResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsBonjourForwardingResource{}
)

func NewNetworksWirelessSSIDsBonjourForwardingResource() resource.Resource {
	return &NetworksWirelessSSIDsBonjourForwardingResource{}
}

type NetworksWirelessSSIDsBonjourForwardingResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsBonjourForwardingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_bonjour_forwarding"
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, Bonjour forwarding is enabled on this SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `List of bonjour forwarding rules.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"description": schema.StringAttribute{
							MarkdownDescription: `A description for your Bonjour forwarding rule. Optional.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"services": schema.SetAttribute{
							MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"vlan_id": schema.StringAttribute{
							MarkdownDescription: `The ID of the service VLAN. Required.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsBonjourForwardingRs

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
	vvNumber := data.Number.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsBonjourForwarding only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsBonjourForwarding only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsBonjourForwardingRs

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
	vvNumber := data.Number.ValueString()
	// number
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsBonjourForwardingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsBonjourForwardingRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvNumber := data.Number.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsBonjourForwardingRs struct {
	NetworkID types.String                                                      `tfsdk:"network_id"`
	Number    types.String                                                      `tfsdk:"number"`
	Enabled   types.Bool                                                        `tfsdk:"enabled"`
	Rules     *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs struct {
	Description types.String `tfsdk:"description"`
	Services    types.Set    `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

// FromBody
func (r *NetworksWirelessSSIDsBonjourForwardingRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			description := rItem1.Description.ValueString()
			var services []string = nil

			rItem1.Services.ElementsAs(ctx, &services, false)
			vLANID := rItem1.VLANID.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules = append(requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules{
				Description: description,
				Services:    services,
				VLANID:      vLANID,
			})
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding{
		Enabled: enabled,
		Rules: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules {
			if len(requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBodyRs(state NetworksWirelessSSIDsBonjourForwardingRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding, is_read bool) NetworksWirelessSSIDsBonjourForwardingRs {
	itemState := NetworksWirelessSSIDsBonjourForwardingRs{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs{
						Description: types.StringValue(rules.Description),
						Services:    StringSliceToSet(rules.Services),
						VLANID:      types.StringValue(rules.VLANID),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsBonjourForwardingRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsBonjourForwardingRs)
}
