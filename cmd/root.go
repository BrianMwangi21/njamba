/*
Copyright © 2023 Mwangi Kabiru mwangikabiru21@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "njamba",
	Short: "njamba is your OpenAI assistant on the terminal",
	Long: `
njamba is your OpenAI assistant on the terminal. It is a command line tool that
uses the OpenAI API for natural language processing. It can be used to generate
text, search for information, and answer questions. It is a work in progress and
is currently in alpha stage. It is written in Go and uses Cobra for command line interface. 
It is also licensed under the GNU General Public License v3.0.

Feel free to contribute to the project on Github: https://github.com/BrianMwangi21/njamba
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.njamba.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
