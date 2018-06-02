package engine

import (
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

// GetValue asks interactively for a value
func getValue(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	return result
}
