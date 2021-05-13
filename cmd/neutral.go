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

	"github.com/agajdosi/tauto/pkg/database"
	"github.com/spf13/cobra"
)

var neutralCmd = &cobra.Command{
	Use:   "neutral",
	Short: "Manipulates available neutral accounts.",
	Long:  `Manipulates available neutral accounts. Neutrals are Twitter accounts which we take as neutral sources of information - like news sources. Their posts gonna be sometimes retweeted, liked or commented. Its purpose is to fake that our bots are normal persons.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var neutralAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new neutral account into the database.",
	Long:  `Adds a new neutral account into the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()

		_, err := database.AddOther(username, "neutral")
		if err != nil {
			log.Fatal(err)
		}
	},
}

var neutralRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes neutral account from the database.",
	Long:  `Removes neutral account from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.DeleteOther(username, "neutral")
		if err != nil {
			log.Fatal(err)
		}
	},
}

var neutralListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all neutrals in the database.",
	Long:  `Lists all neutrals in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		database.ListOthers("neutral")
	},
}

func init() {
	neutralCmd.AddCommand(neutralAddCmd)
	neutralCmd.AddCommand(neutralListCmd)
	neutralCmd.AddCommand(neutralRemoveCmd)
	rootCmd.AddCommand(neutralCmd)

	neutralAddCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the neutral account.")
	neutralAddCmd.MarkFlagRequired("username")

	neutralRemoveCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the neutral account.")
	neutralRemoveCmd.MarkFlagRequired("username")
}
