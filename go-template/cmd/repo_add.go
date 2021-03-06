// Copyright © 2018 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"livingit.de/code/go-template"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// repoAddCmd represents the add command
var repoAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new repository",
	Long: `add a new repository to local storage

This does a clone of the repository to the storage directory`,
	Run: func(cmd *cobra.Command, args []string) {
		e, err := engine.New(viper.GetString("storage"), viper.GetString("log-level"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		repository, err := e.AddRepository(cmd.Flag("url").Value.String(), cmd.Flag("branch").Value.String())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Name:        %s", repository.Name))
		fmt.Println(fmt.Sprintf("Author:      %s", repository.Author))
		fmt.Println(fmt.Sprintf("URL:         %s", repository.URL))
		fmt.Println(fmt.Sprintf("Description:\n%s", repository.Description))
	},
}

func init() {
	repoCmd.AddCommand(repoAddCmd)

	repoAddCmd.Flags().StringP("url", "u", "", "provide url where to look for repository")
	repoAddCmd.MarkFlagRequired("url")

	repoAddCmd.Flags().StringP("branch", "b", "", "provide a branch or tag")
}
