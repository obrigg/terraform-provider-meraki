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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksApplianceSecurityIntrusionDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceSecurityIntrusionDataSource{}
)

func NewNetworksApplianceSecurityIntrusionDataSource() datasource.DataSource {
	return &NetworksApplianceSecurityIntrusionDataSource{}
}

type NetworksApplianceSecurityIntrusionDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceSecurityIntrusionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceSecurityIntrusionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_security_intrusion"
}

func (d *NetworksApplianceSecurityIntrusionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ids_rulesets": schema.StringAttribute{
						Computed: true,
					},
					"mode": schema.StringAttribute{
						Computed: true,
					},
					"protected_networks": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"excluded_cidr": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"included_cidr": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"use_default": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceSecurityIntrusionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceSecurityIntrusion NetworksApplianceSecurityIntrusion
	diags := req.Config.Get(ctx, &networksApplianceSecurityIntrusion)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceSecurityIntrusion")
		vvNetworkID := networksApplianceSecurityIntrusion.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceSecurityIntrusion(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}

		networksApplianceSecurityIntrusion = ResponseApplianceGetNetworkApplianceSecurityIntrusionItemToBody(networksApplianceSecurityIntrusion, response1)
		diags = resp.State.Set(ctx, &networksApplianceSecurityIntrusion)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceSecurityIntrusion struct {
	NetworkID types.String                                           `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceSecurityIntrusion `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceSecurityIntrusion struct {
	IDsRulesets       types.String                                                            `tfsdk:"ids_rulesets"`
	Mode              types.String                                                            `tfsdk:"mode"`
	ProtectedNetworks *ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks `tfsdk:"protected_networks"`
}

type ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks struct {
	ExcludedCidr types.List `tfsdk:"excluded_cidr"`
	IncludedCidr types.List `tfsdk:"included_cidr"`
	UseDefault   types.Bool `tfsdk:"use_default"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceSecurityIntrusionItemToBody(state NetworksApplianceSecurityIntrusion, response *merakigosdk.ResponseApplianceGetNetworkApplianceSecurityIntrusion) NetworksApplianceSecurityIntrusion {
	itemState := ResponseApplianceGetNetworkApplianceSecurityIntrusion{
		IDsRulesets: types.StringValue(response.IDsRulesets),
		Mode:        types.StringValue(response.Mode),
		ProtectedNetworks: func() *ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks {
			if response.ProtectedNetworks != nil {
				return &ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks{
					ExcludedCidr: StringSliceToList(response.ProtectedNetworks.ExcludedCidr),
					IncludedCidr: StringSliceToList(response.ProtectedNetworks.IncludedCidr),
					UseDefault: func() types.Bool {
						if response.ProtectedNetworks.UseDefault != nil {
							return types.BoolValue(*response.ProtectedNetworks.UseDefault)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceSecurityIntrusionProtectedNetworks{}
		}(),
	}
	state.Item = &itemState
	return state
}
