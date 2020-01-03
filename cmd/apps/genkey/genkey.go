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

package crtkey

import (
	"net/url"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/srinandan/apigeecli/cmd/shared"
)

//Cmd to generate key
var Cmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate a new developer KeyPair",
	Long:  "Generate a new developer KeyPair",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		u, _ := url.Parse(shared.BaseURL)

		key := []string{}

		key = append(key, "\"name\":\""+name+"\"")
		key = append(key, "\"apiProducts\":[\""+getArrayStr(apiProducts)+"\"]")

		if callback != "" {
			key = append(key, "\"callbackUrl\":\""+callback+"\"")
		}

		if expires != "" {
			key = append(key, "\"keyExpiresIn\":\""+expires+"\"")
		}

		if len(scopes) > 0 {
			key = append(key, "\"scopes\":[\""+getArrayStr(scopes)+"\"]")
		}

		payload := "{" + strings.Join(key, ",") + "}"
		u.Path = path.Join(u.Path, shared.RootArgs.Org, "developers", devID, "apps", name)
		_, err = shared.HttpClient(true, u.String(), payload)
		return
	},
}

var name, devID, expires, callback string
var apiProducts, scopes []string

func init() {

	Cmd.Flags().StringVarP(&name, "name", "n",
		"", "Name of the developer app")
	Cmd.Flags().StringVarP(&devID, "devid", "d",
		"", "Developer Id")
	Cmd.Flags().StringVarP(&expires, "expires", "x",
		"", "A setting, in milliseconds, for the lifetime of the consumer key")
	Cmd.Flags().StringVarP(&callback, "callback", "c",
		"", "The callbackUrl is used by OAuth")
	Cmd.Flags().StringArrayVarP(&apiProducts, "prods", "p",
		[]string{}, "A list of api products")
	Cmd.Flags().StringArrayVarP(&scopes, "scopes", "s",
		[]string{}, "OAuth scopes")

	_ = Cmd.MarkFlagRequired("name")
	_ = Cmd.MarkFlagRequired("devid")
	_ = Cmd.MarkFlagRequired("prods")
}

func getArrayStr(str []string) string {
	tmp := strings.Join(str, ",")
	tmp = strings.ReplaceAll(tmp, ",", "\",\"")
	return tmp
}
