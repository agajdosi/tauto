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
	"log"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/database"
	"github.com/spf13/cobra"
)

var enemyCmd = &cobra.Command{
	Use:   "enemy",
	Short: "Manipulates available enemy accounts.",
	Long:  `Manipulates available enemy accounts. Enemies are Twitter accounts which you want want to fight. Their posts gonna be retweeted or commented negatively.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var enemyAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new enemy account into the database.",
	Long:  `Adds a new enemy account into the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()

		_, err := database.AddOther(username, "enemy")
		if err != nil {
			log.Fatal(err)
		}
	},
}

var enemyRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes enemy account from the database.",
	Long:  `Removes enemy account from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.DeleteOther(username, "enemy")
		if err != nil {
			log.Fatal(err)
		}
	},
}

var enemyListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all enemy accounts in the database.",
	Long:  `Lists all enemy accounts in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		database.ListOthers("enemy")
	},
}

func init() {
	enemyCmd.AddCommand(enemyAddCmd)
	enemyCmd.AddCommand(enemyListCmd)
	enemyCmd.AddCommand(enemyRemoveCmd)
	rootCmd.AddCommand(enemyCmd)

	enemyAddCmd.Flags().StringVarP(&username, "username", "u", "", "Username of enemy account.")
	enemyAddCmd.MarkFlagRequired("username")

	enemyRemoveCmd.Flags().StringVarP(&username, "username", "u", "", "Username of enemy account.")
	enemyRemoveCmd.MarkFlagRequired("username")
}
