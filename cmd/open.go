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

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens single profile and leaves it for manual tweaks.",
	Long:  `Opens single profile and leaves it for manual tweaks.`,
	Run: func(cmd *cobra.Command, args []string) {
		users, err := database.GetBots(username)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			u := twitter.NewUser(user.Username, user.Password)
			err = u.Open()
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	openCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the account which will be opened.")
}
