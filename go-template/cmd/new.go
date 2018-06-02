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

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create project from template",
	Long:  `Create a project from a template`,
	Run: func(cmd *cobra.Command, args []string) {
		arguments, err := splitArguments(cmd.Flag("arguments").Value.String())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		e, err := engine.New(viper.GetString("storage"), viper.GetString("log-level"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = e.New(cmd.Flag("name").Value.String(), cmd.Flag("template").Value.String(), arguments)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("name", "n", "", "provide name for new project")
	newCmd.MarkFlagRequired("name")

	newCmd.Flags().StringP("template", "t", "", "provide name for template to use")
	newCmd.MarkFlagRequired("template")

	newCmd.Flags().String("arguments", "", "add arguments like this: --arguments Namespace=github.com/sascha.andres,Author=Someone")
}
