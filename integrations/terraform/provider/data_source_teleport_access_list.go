// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"

	convert "github.com/gravitational/teleport/api/types/accesslist/convert/v1"
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	schemav1 "github.com/gravitational/teleport/integrations/terraform/tfschema/accesslist/v1"
)

// dataSourceTeleportAccessListType is the data source metadata type
type dataSourceTeleportAccessListType struct{}

// dataSourceTeleportAccessList is the resource
type dataSourceTeleportAccessList struct {
	p Provider
}

// GetSchema returns the data source schema
func (r dataSourceTeleportAccessListType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return schemav1.GenSchemaAccessList(ctx)
}

// NewDataSource creates the empty data source
func (r dataSourceTeleportAccessListType) NewDataSource(_ context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceTeleportAccessList{
		p: *(p.(*Provider)),
	}, nil
}

// Read reads teleport AccessList
func (r dataSourceTeleportAccessList) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var id types.String
	diags := req.Config.GetAttribute(ctx, path.Root("header").AtName("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accessListI, err := r.p.Client.AccessListClient().GetAccessList(ctx, id.Value)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading AccessList", trace.Wrap(err), "access_list"))
		return
	}

    var state types.Object
	accessList := convert.ToProto(accessListI)
	
	diags = schemav1.CopyAccessListToTerraform(ctx, accessList, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}