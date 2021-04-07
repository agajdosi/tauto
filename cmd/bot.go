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

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "manipulates bots",
	Long:  `Manipulates bots.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new bot into the database.",
	Long:  `Adds a new bot into the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.EnsureExists()
		if err != nil {
			log.Fatal(err)
		}

		_, err = database.AddBot(username, password, "twitter")
		if err != nil {
			log.Fatal(err)
		}
	},
}

// TBD
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "TBD. Removes bot from the database.",
	Long:  `TBD. Removes bot from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// TBD
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "TBD. Lists all bots in the database.",
	Long:  `TBD. Lists all bots in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	botCmd.AddCommand(addCmd)
	botCmd.AddCommand(listCmd)
	botCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(botCmd)

	addCmd.Flags().StringVarP(&username, "username", "u", "", "Username to be used to log in the bot.")
	addCmd.MarkFlagRequired("username")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "Password to be used to log in the bot.")
	addCmd.MarkFlagRequired("password")

	removeCmd.Flags().StringVarP(&username, "username", "u", "", "Username of bot which is going to be removed.")
	removeCmd.MarkFlagRequired("username")
}
