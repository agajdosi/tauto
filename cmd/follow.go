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
	"log"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/database"
	"github.com/agajdosi/twitter-storm-toolkit/pkg/twitter"

	"github.com/spf13/cobra"
)

var who []string

// followCmd represents the follow command
var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Will follow a selected user on twitter.",
	Long:  `Will follow a selected user on twitter.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.EnsureExists()
		if err != nil {
			log.Fatal(err)
		}

		if username != "" {
			followBySingle()
		} else {
			followByAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(followCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// followCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// followCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	followCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the account who will follow. When left empty it will use all usernames available in the database.")

	followCmd.Flags().StringSliceVarP(&who, "who", "w", nil, "Which account's username(s) to follow. Can be a string or list of strings.")
	followCmd.MarkFlagRequired("who")
}

func followByAll() error {
	users, err := database.GetAllBots()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		u := twitter.NewUser(user.Username, user.Password)

		for _, toFollow := range who {
			err = u.Follow(toFollow)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}

func followBySingle() error {
	password, err := database.GetBot(username)
	if err != nil {
		log.Fatal(err)
	}

	user := twitter.NewUser(username, password)
	for _, toFollow := range who {
		err = user.Follow(toFollow)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
