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

	"github.com/agajdosi/tauto/pkg/database"
	"github.com/agajdosi/tauto/pkg/twitter"

	"github.com/spf13/cobra"
)

var who []string

var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Will follow a selected user on twitter.",
	Long:  `Will follow a selected user on twitter.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.EnsureExists()
		follow()
	},
}

func init() {
	rootCmd.AddCommand(followCmd)

	followCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the account who will follow. When left empty it will use all usernames available in the database.")

	followCmd.Flags().StringSliceVarP(&who, "who", "w", nil, "Which account's username(s) to follow. Can be a string or list of strings.")
	followCmd.MarkFlagRequired("who")
}

func follow() error {
	users, err := database.GetBots(username, true)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		u, cancel := twitter.NewUser(user.ID, user.Username, user.Password, 300)
		for _, toFollow := range who {
			err = u.EnsureFollowed(toFollow)
			if err != nil {
				fmt.Println(err)
			}
		}
		cancel()
	}

	return nil
}
