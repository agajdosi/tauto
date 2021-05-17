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

	"github.com/agajdosi/tauto/pkg/generate"
	"github.com/spf13/cobra"
)

var templateArg string

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Use this command to test your comment template before adding it into the config file.",
	Long:  `Use this command to test your comment template before adding it into the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < 20; i++ {
			tweet := generate.FromTemplate(templateArg)
			fmt.Println(tweet)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	templateCmd.Flags().StringVarP(&templateArg, "template", "t", "", "Template to test. Will generate 30 variants of the template.")
	templateCmd.MarkFlagRequired("template")
}
