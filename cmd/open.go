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

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens single profile and leaves it for manual tweaks.",
	Long:  `Opens single profile and leaves it for manual tweaks.`,
	Run: func(cmd *cobra.Command, args []string) {
		users, err := database.GetBots(username, false)
		if err != nil {
			log.Fatal(err)
		}

		for _, user := range users {
			u, cancel := twitter.NewUser(user.ID, user.Username, user.Password, 9999)
			err = u.Open()
			if err != nil {
				fmt.Println(err)
			}
			cancel()
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringVarP(&username, "username", "u", "", "Username of the account which will be opened.")
}
