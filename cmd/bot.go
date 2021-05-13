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
	Short: "Manipulates bot accounts.",
	Long:  `Manipulates bot accounts. Bots are our Twitter accounts which we automatically use to spread negative or positive messages.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var botAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new bot account into the database.",
	Long:  `Adds a new bot account into the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		_, err := database.AddBot(username, password)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var botRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes bot account from the database.",
	Long:  `Removes bot account from the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		err := database.DeleteBot(username)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var botListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all bot accounts in the database.",
	Long:  `Lists all bot accounts in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		err := database.ListBots()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	botCmd.AddCommand(botAddCmd)
	botCmd.AddCommand(botListCmd)
	botCmd.AddCommand(botRemoveCmd)
	rootCmd.AddCommand(botCmd)

	botAddCmd.Flags().StringVarP(&username, "username", "u", "", "Username to be used to log in the bot account.")
	botAddCmd.MarkFlagRequired("username")
	botAddCmd.Flags().StringVarP(&password, "password", "p", "", "Password to be used to log in the bot account.")
	botAddCmd.MarkFlagRequired("password")

	botRemoveCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the bot account which is going to be removed.")
	botRemoveCmd.MarkFlagRequired("username")
}
