// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package tenant

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/cobra"
)

var tenantNameArgsError = pulsar.Output{
	Desc: "the tenant name is not specified or the tenant name is specified more than one",
	Out:  "[✖]  only one argument is allowed to be used as a name",
}

var tenantNotExistError = pulsar.Output{
	Desc: "the specified tenant does not exist in the broker",
	Out:  "[✖]  code: 404 reason: The tenant does not exist",
}

var tenantAlreadyExistError = pulsar.Output{
	Desc: "the specified tenant has been created",
	Out:  "[✖]  code: 409 reason: Tenant already exists",
}

func Command(flagGrouping *cmdutils.FlagGrouping) *cobra.Command {
	resourceCmd := cmdutils.NewResourceCmd(
		"tenants",
		"Operations about tenant(s)",
		"",
		"tenant")

	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, createTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, deleteTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, UpdateTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, listTenantCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, getTenantCmd)

	return resourceCmd
}