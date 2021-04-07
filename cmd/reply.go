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
	"fmt"
	"log"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/database"
	"github.com/agajdosi/twitter-storm-toolkit/pkg/twitter"
	"github.com/spf13/cobra"
)

var where []string

var replyCmd = &cobra.Command{
	Use:   "reply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		reply()
	},
}

func init() {
	rootCmd.AddCommand(replyCmd)

	replyCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the account who will reply. When left empty it will use all usernames available in the database.")

	replyCmd.Flags().StringSliceVarP(&where, "where", "w", nil, "Where to reply - a url of the tweet(s).")
	replyCmd.MarkFlagRequired("who")

	replyCmd.Flags().StringVarP(&tweet, "tweet", "t", "", "Tweet which will be replied.")
	replyCmd.MarkFlagRequired("tweet")
}

func reply() error {
	users, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		u, cancel := twitter.NewUser(user.ID, user.Username, user.Password)

		for _, w := range where {
			err = u.Reply(tweet, w)
			if err != nil {
				fmt.Println(err)
			}
		}
		cancel()
	}

	return nil
}
