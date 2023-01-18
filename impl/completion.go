package impl

import (
	"context"
	"os"
	"strconv"

	"github.com/BrianMwangi21/njamba/config"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/cobra"
)

func RunCompletion(cmd *cobra.Command, args []string) {
	printSlowly("Running prompt completion...\n")

	// Get user prompt
	userPromptContent := promptContent{
		"Please enter your prompt",
		"What would you like to know ?",
	}
	userPrompt := promptGetInput(userPromptContent)

	// Get model
	modelPromptContent := promptContent{
		"Please select a model to use",
		"Which model would you like to use for this completion ?",
	}
	modelPrompt := promptGetSelect(modelPromptContent, GPT_MODELS)

	// Get temperature
	tempPromptContent := promptContent{
		"Please enter your preferred temperature",
		"What would you like the temperature of the request to be ? [Between 0 - 1, Recommended : 0.7] : ",
	}
	tempPrompt, err := strconv.ParseFloat(promptGetInput(tempPromptContent), 32)
	if err != nil {
		printSlowly("Temperature must be a number between 0 and 1. Please try again\n")
		os.Exit(1)
	}

	// Get max tokens
	tokensPromptContent := promptContent{
		"Please enter your preferred max tokens",
		"What would you like the max tokens for the response to be ? [Between 0 - 2024, Recommended : > 256] : ",
	}
	tokensPrompt, err := strconv.Atoi(promptGetInput(tokensPromptContent))
	if err != nil || tokensPrompt < 0 || tokensPrompt > 2024 {
		printSlowly("Max tokens must be a number between 0 and 2024. Please try again\n")
		os.Exit(1)
	}

	c := gogpt.NewClient(configs.EnvOpenAI())
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Prompt:      userPrompt,
		Model:       modelPrompt,
		Temperature: float32(tempPrompt),
		MaxTokens:   tokensPrompt,
	}

	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		printSlowly(string(err.Error()))
		os.Exit(1)
	}

	printSlowly(resp.Choices[0].Text)
}
