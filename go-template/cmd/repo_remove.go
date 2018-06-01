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

// repoRemoveCmd represents the remove command
var repoRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a repository",
	Long: `add a new repository to local storage

This removes a clone of the repository from the storage directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	repoCmd.AddCommand(repoRemoveCmd)

	repoRemoveCmd.Flags().StringP("name", "n", "", "Provide name for project to remove")
	repoRemoveCmd.MarkFlagRequired("name")
}
