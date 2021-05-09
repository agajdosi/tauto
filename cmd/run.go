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
		handleBots()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the bot which will follow. When left empty it will use all bots available in the database.")
}

func handleBots() {
	//allies, _ := database.GetOthers("", "ally")
	//neutrals, _ := database.GetOthers("", "neutral")
	enemies, _ := database.GetOthers("", "enemy")

	bots, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, bot := range bots {
		b, cancel := twitter.NewUser(bot.ID, bot.Username, bot.Password, 6000)
		//handleNeutrals(b, neutrals)
		//handleAllies(b, allies)
		handleEnemies(b, enemies)
		cancel()
	}
}

func handleAllies(b twitter.Bot, allies []database.Other) {
	for _, ally := range allies {
		tweets := twitter.GetTweets(ally.Username)
		for tweet := range tweets {
			b.MaybeLike(tweet.PermanentURL, 1)
			b.MaybeRetweet(tweet.PermanentURL, 1)
		}
	}
}

func handleNeutrals(b twitter.Bot, neutrals []database.Other) {
	for _, neutral := range neutrals {
		tweets := twitter.GetTweets(neutral.Username)
		for tweet := range tweets {
			b.MaybeLike(tweet.PermanentURL, 0.1)
			b.MaybeRetweet(tweet.PermanentURL, 0.2)
		}
	}
}

func handleEnemies(b twitter.Bot, enemies []database.Other) {
	for _, enemy := range enemies {
		tweets := twitter.GetTweets(enemy.Username)
		for tweet := range tweets {
			if tweet.IsRetweet && !tweet.IsQuoted {
				continue
			}

			if tweet.IsReply {
				continue
			}

			b.TrollComment(tweet.PermanentURL, b.Username)
		}
	}
}
