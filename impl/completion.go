package impl

import (
	"context"

	"github.com/BrianMwangi21/njamba/config"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/cobra"
)

func RunCompletion(cmd *cobra.Command, args []string) {
	userPromptContent := promptContent{
		"Please enter your prompt",
		"What would you like to know ?",
	}
	userPrompt := promptGetInput(userPromptContent)

	c := gogpt.NewClient(configs.EnvOpenAI())
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Davinci,
		MaxTokens: 10,
		Prompt:    userPrompt,
	}
	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		printSlowly(string(err.Error()))
	}

	printSlowly(resp.Choices[0].Text)
}
