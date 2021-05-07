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
	"github.com/agajdosi/twitter-storm-toolkit/pkg/twitter"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the TST: love allies, checkout neutrals, hate enemies.",
	Long:  `Run the TST: love allies, checkout neutrals, hate enemies.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		supportAllies()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the bot which will follow. When left empty it will use all bots available in the database.")
}

func supportAllies() {
	allies, _ := database.GetOthers("", "ally")
	bots, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, bot := range bots {
		b, cancel := twitter.NewUser(bot.ID, bot.Username, bot.Password, 6000)
		for _, ally := range allies {
			tweets := twitter.GetTweets(ally.Username)
			for _, tweet := range tweets {
				b.MaybeLike(tweet, 1)
				b.MaybeRetweet(tweet, 0.4)
			}
		}
		cancel()
	}
}
