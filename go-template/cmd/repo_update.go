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

	"github.com/spf13/cobra"
)

// repoUpdateCmd represents the update command
var repoUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update repositories",
	Long: `Update repositories in your local storage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	repoCmd.AddCommand(repoUpdateCmd)

	repoUpdateCmd.Flags().StringP("limit-to", "l", "", "Provide name for repository to update, omit for all")
	repoUpdateCmd.MarkFlagRequired("limit-to")
}
