// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sync

import (
	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/cmd/shared"
	"github.com/srinandan/apigeecli/cmd/sync/getsync"
	"github.com/srinandan/apigeecli/cmd/sync/resetsync"
	"github.com/srinandan/apigeecli/cmd/sync/setsync"
)

//Cmd to manage identities
var Cmd = &cobra.Command{
	Use:   "sync",
	Short: "Manage identities with grant access to control plane resources",
	Long:  "Manage identities with grant access to control plane resources",
}

func init() {

	Cmd.PersistentFlags().StringVarP(&shared.RootArgs.Org, "org", "o",
		"", "Apigee organization name")

	_ = Cmd.MarkPersistentFlagRequired("org")
	Cmd.AddCommand(setsync.Cmd)
	Cmd.AddCommand(getsync.Cmd)
	Cmd.AddCommand(resetsync.Cmd)
}
