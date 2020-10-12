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
	"os"

	"github.com/agajdosi/twitter-storm-toolkit/pkg/twitter"
	"github.com/spf13/cobra"
)

var username string
var password string
var tweet string

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := twitter.NewUser(username, password)
		err := user.Post(tweet)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
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
	postCmd.Flags().StringVar(&username, "username", "", "Username to be used to log into Twitter.")
	postCmd.MarkFlagRequired("username")

	postCmd.Flags().StringVar(&password, "password", "", "Password to be used to log into Twitter.")
	postCmd.MarkFlagRequired("password")

	postCmd.Flags().StringVar(&tweet, "tweet", "", "Tweet which will be tweeted.")
	postCmd.MarkFlagRequired("tweet")
}
