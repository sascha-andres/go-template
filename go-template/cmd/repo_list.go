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
	"text/tabwriter"

	"github.com/sascha-andres/go-template"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// repoListCmd represents the list command
var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all local repositories",
	Long:  `List information about local repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		e, err := engine.New(viper.GetString("storage"), viper.GetString("log-level"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		list, err := e.ListRepositories()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)
		fmt.Fprintln(w, "name\tauthor\turl\t")
		for _, entry := range list {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t", entry.Name, entry.Author, entry.URL))
		}
		w.Flush()
	},
}

func init() {
	repoCmd.AddCommand(repoListCmd)
}
