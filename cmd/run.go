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
	"context"
	"log"
	"math/rand"

	"github.com/agajdosi/tauto/pkg/database"
	"github.com/agajdosi/tauto/pkg/twitter"
	twitterscraper "github.com/n0madic/twitter-scraper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Tauto: love allies, checkout neutrals, troll enemies. This is the ultimate command.",
	Long:  `Run the Tauto: love allies, checkout neutrals, troll enemies. This is the ultimate command.`,
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
	allies, _ := database.GetOthers("", "ally")
	neutrals, _ := database.GetOthers("", "neutral")
	enemies, _ := database.GetOthers("", "enemy")

	bots, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, bot := range bots {
		b, cancel := twitter.NewUser(bot.ID, bot.Username, bot.Password, 999999)
		accessible, err := b.IsProfileAccessible()
		if err != nil {
			log.Fatal(err)
		}
		if accessible == false {
			continue
		}

		slander(b)
		handleNeutrals(b, neutrals)
		handleAllies(b, allies)
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

			b.TrollReply(tweet.PermanentURL)
		}
	}
}

func slander(b twitter.Bot) {
	//CHAOTIC UNMARSHAL - SHOULD BE REDESIGNED
	targets := viper.Get("slanderTargets")
	for _, target := range targets.([]interface{}) {
		for name, templates := range target.(map[interface{}]interface{}) {
			var ts []string
			for _, template := range templates.([]interface{}) {
				ts = append(ts, template.(string))
			}

			//Search Twitter for mentions
			scraper := twitterscraper.New()
			scraper.SetSearchMode(twitterscraper.SearchLatest)
			tweets := scraper.SearchTweets(context.Background(), name.(string), 40)
			for tweet := range tweets {
				if rand.Float32() > 0.25 {
					continue
				}
				template := ts[rand.Intn(len(ts))]
				b.ReplyFromTemplate(tweet.PermanentURL, template)
			}
		}
	}
}

//You’re unable to view this Tweet because this account owner limits who can view their Tweets.
