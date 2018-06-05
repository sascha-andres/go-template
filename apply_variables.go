package engine

import (
	"bytes"
	"text/template"
)

// applyVariables takes a template and returns the result
func applyVariables(templateContent, name string, arguments map[string]string) (string, error) {
	textTemplate, err := template.New("replacement").Parse(templateContent)
	if err != nil {
		return "", err
	}
	var result []byte
	buff := bytes.NewBuffer(result)
	err = textTemplate.Execute(buff, struct {
		Name      string
		Arguments map[string]string
	}{
		Name:      name,
		Arguments: arguments,
	})
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}
