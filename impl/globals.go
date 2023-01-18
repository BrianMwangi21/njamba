package impl

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

var (
	GPT_MODELS = []string{
		"text-davinci-003",
		"text-davinci-002",
		"text-curie-001",
		"text-babbage-001",
		"text-ada-001",
		"text-davinci-001",
		"davinci-instruct-beta",
		"davinci",
		"curie-instruct-beta",
		"curie",
		"ada",
		"babbage",
	}
)

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func promptGetSelect(pc promptContent, items []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func printSlowly(s string) {
	for _, r := range s {
		fmt.Printf("%c", r)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}
