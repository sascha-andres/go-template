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

	"github.com/sascha-andres/go-template"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// repoInfoCmd represents the info command
var repoInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "print information about a repository",
	Long: `Prints information about a local repository:

- Name
- Author
- Description
- Homepage
- Arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		e, err := engine.New(viper.GetString("storage"), viper.GetString("log-level"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		repository, arguments, err := e.InfoRepository(cmd.Flag("name").Value.String())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Name:        %s", repository.Name))
		fmt.Println(fmt.Sprintf("Author:      %s", repository.Author))
		fmt.Println(fmt.Sprintf("URL:         %s", repository.URL))
		fmt.Println(fmt.Sprintf("Description:\n%s", repository.Description))
		fmt.Println()
		fmt.Println("Arguments:")
		for _, arg := range arguments {
			fmt.Println(fmt.Sprintf("             - %s", arg))
		}
	},
}

func init() {
	repoCmd.AddCommand(repoInfoCmd)

	repoInfoCmd.Flags().StringP("name", "n", "", "provide name of repository")
	repoInfoCmd.MarkFlagRequired("name")
}
