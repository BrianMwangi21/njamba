package impl

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMsg string
	label    string
}

type GPT_MODELS struct {
	names        []string
	descriptions []string
}

var (
	MODELS = GPT_MODELS{
		[]string{"text-davinci-003", "text-curie-001", "text-babbage-001", "text-ada-001"},
		[]string{"Most capable GPT-3 model, often with higher quality output.", "Very capable, but faster and lower cost than Davinci.", "Capable of straightforward tasks, very fast, and lower cost.", "Capable of very simple tasks, usually the fastest model in the GPT-3 series, and lowest cost."},
	}
	MULTIPLE_RESPONSES_CHOICES = []string{
		"Yes! Show me all - I am rather curious",
		"Nope. I'm good with just the first one - I'm feeling lucky",
	}
)

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

func promptGetSelect(pc promptContent, items []string) (int, string) {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label: pc.label,
			Items: items,
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

	return index, result
}

func printSlowly(s string) {
	for _, r := range s {
		fmt.Printf("%c", r)
		time.Sleep(100 * time.Millisecond)
	}
}

func prettyModelDesc() []string {
	var result []string
	for i, name := range MODELS.names {
		result = append(result, fmt.Sprintf("%s - %s", name, MODELS.descriptions[i]))
	}
	return result
}
