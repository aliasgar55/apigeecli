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

package crtsf

import (
	"net/url"
	"path"

	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/cmd/shared"
)

//Cmd to create shared flow
var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a sharedflow in an Apigee Org",
	Long:  "Creates a sharedflow in an Apigee Org",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		if proxy != "" {
			return shared.ImportBundle("sharedflows", name, proxy)
		}
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "sharedflows")
		proxyName := "{\"name\":\"" + name + "\"}"
		_, err = shared.HttpClient(true, u.String(), proxyName)
		return
	},
}

var name, proxy string

func init() {

	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "Sharedflow name")
	Cmd.Flags().StringVarP(&proxy, "proxy", "p",
		"", "Sharedflow bundle path")

	_ = Cmd.MarkFlagRequired("name")
}
