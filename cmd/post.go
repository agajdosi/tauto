/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Posts a new tweet on Twitter.",
	Long:  `Posts a new tweet on Twitter.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.EnsureExists()
		if err != nil {
			log.Fatal(err)
		}

		if username != "" {
			postToSingle()
		} else {
			postToAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	postCmd.Flags().StringVar(&username, "username", "", "Username to be used to log into Twitter. If username is not provided, the command will tweet from all available accounts.")

	postCmd.Flags().StringVar(&tweet, "tweet", "", "Tweet which will be tweeted.")
	postCmd.MarkFlagRequired("tweet")
}

func postToAll() error {
	fmt.Println("posting to all is not yet supported")
	return nil
}

func postToSingle() error {
	password, err := database.GetBot(username)
	if err != nil {
		log.Fatal(err)
	}

	user := twitter.NewUser(username, password)
	err = user.Post(tweet)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
