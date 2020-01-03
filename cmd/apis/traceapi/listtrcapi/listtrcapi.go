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

package listtrcapi

import (
	"net/url"
	"path"

	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/cmd/shared"
)

//Cmd to manage tracing of apis
var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List all debug sessions for an API proxy revision",
	Long:  "List all debug sessions for an API proxy revision deployed in an environment",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "environments", shared.RootArgs.Env,
			"apis", name, "revisions", revision, "debugsessions")
		_, err = shared.HttpClient(true, u.String())
		return

	},
}

var name, revision string

func init() {

	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "API proxy name")
	Cmd.Flags().StringVarP(&revision, "rev", "v",
		"", "API Proxy revision")

	_ = Cmd.MarkFlagRequired("name")
	_ = Cmd.MarkFlagRequired("rev")

}
