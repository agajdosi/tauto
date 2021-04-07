/*
Copyright © 2020 Andreas Gajdosik <andreas.gajdosik@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// TBD
var allyCmd = &cobra.Command{
	Use:   "ally",
	Short: "TBD. Manipulates allied accounts.",
	Long:  `TBD. Manipulates allied accounts. Allies are Twitter accounts which you want want to support. Their posts gonna be retweeted, liked or commented positively.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// TBD
var allyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "TBD. Adds a new allied account into the database.",
	Long:  `TBD. Adds a new allied account the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// TBD
var allyRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "TBD. Removes allied account from the database.",
	Long:  `TBD. Removes allied account from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// TBD
var allyListCmd = &cobra.Command{
	Use:   "list",
	Short: "TBD. Lists all allied accounts in the database.",
	Long:  `TBD. Lists all allied accounts in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	allyCmd.AddCommand(allyAddCmd)
	allyCmd.AddCommand(allyListCmd)
	allyCmd.AddCommand(allyRemoveCmd)
	rootCmd.AddCommand(allyCmd)

	allyAddCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the allied account.")
	allyAddCmd.MarkFlagRequired("username")

	allyRemoveCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the allied account.")
	allyRemoveCmd.MarkFlagRequired("username")
}
