/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"github.com/BrianMwangi21/njamba/impl"
	"github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "This command allows a user to send a prompt to the OpenAI API",
	Long: `This command allows a user to send a prompt to the OpenAI API.
The user will be guided through the process of creating a prompt.
After the necessary steps, the prompt will be sent to the OpenAI API and 
a response returned to the user.

This command uses the completion models.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		codeFlag, _ := cmd.Flags().GetBool("code")

		impl.RunCompletion(cmd, args, codeFlag)
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	promptCmd.Flags().BoolP("code", "c", false, "Use this flag to send a request to the code completion models instead of the default text completion models.")
}
