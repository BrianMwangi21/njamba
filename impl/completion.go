package impl

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RunCompletion(cmd *cobra.Command, args []string) {
	userPromptContent := promptContent{
		"Please enter your prompt",
		"What would you like to know ?",
	}
	userPrompt := promptGetInput(userPromptContent)

	fmt.Println("This is the userPrompt", userPrompt)
}
