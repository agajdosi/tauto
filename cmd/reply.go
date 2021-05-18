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
	"github.com/agajdosi/tauto/pkg/twitter"
	"github.com/spf13/cobra"
)

var urls []string

var replyCmd = &cobra.Command{
	Use:   "reply",
	Short: "Replies to a specified tweet with the text.",
	Long:  `Replies to a specified tweet with the text.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		reply()
	},
}

func init() {
	rootCmd.AddCommand(replyCmd)

	replyCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the account who will reply. When left empty it will use all usernames available in the database.")

	replyCmd.Flags().StringSliceVarP(&urls, "url", "U", nil, "Where to reply - a url of the tweet(s).")
	replyCmd.MarkFlagRequired("url")

	replyCmd.Flags().StringVarP(&tweet, "tweet", "t", "", "Tweet which will be replied.")
	replyCmd.MarkFlagRequired("tweet")
}

func reply() error {
	bots, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, bot := range bots {
		b, cancel, err := twitter.NewUser(bot.ID, bot.Username, bot.Password, 300)
		if err != nil {
			return err
		}

		for _, url := range urls {
			b.Reply(url, tweet)
		}
		cancel()
	}

	return nil
}
