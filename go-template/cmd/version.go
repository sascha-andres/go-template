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

var (
	Version = "develop"
	Commit  = "develop"
	Date    = "today"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print out version info",
	Long: `Prints out some version information such as version number,
commit and date of compile`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("version: %s", Version))
		fmt.Println(fmt.Sprintf("commit:  %s", Commit))
		fmt.Println(fmt.Sprintf("date:    %s", Date))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
