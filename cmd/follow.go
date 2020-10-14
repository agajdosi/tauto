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

	"github.com/spf13/cobra"
)

var who string

// followCmd represents the follow command
var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Will follow a selected user on twitter.",
	Long:  `Will follow a selected user on twitter.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("follow called")
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

	followCmd.Flags().StringVarP(&who, "who", "w", "", "Which account to follow. Please enter a username of account to follow.")
	followCmd.MarkFlagRequired("who")
}
