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

var username string
var password string
var tweet string

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Posts a new tweet on Twitter.",
	Long:  `Posts a new tweet on Twitter.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		post()
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringVarP(&username, "username", "u", "", "Username to be used to log into Twitter. If username is not provided, the command will tweet from all available accounts.")

	postCmd.Flags().StringVarP(&tweet, "tweet", "t", "", "Tweet which will be tweeted.")
	postCmd.MarkFlagRequired("tweet")
}

func post() error {
	users, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		u := twitter.NewUser(user.ID, user.Username, user.Password)
		err = u.Post(tweet)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
