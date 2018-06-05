package engine

import (
	"fmt"
	"path"
)

// loadTemplateFile reads template configuration and does a validation run
func (e *Engine) loadTemplateFile(templateName string, arguments map[string]string) *TemplateFile {
	if e.err != nil {
		return nil
	}
	templateFile, err := e.readTemplateFile(path.Join(e.storageDirectory, templateName, ".go-template.yml"))
	if err != nil {
		e.err = err
		return nil
	}
	for _, arg := range templateFile.Arguments {
		if _, ok := arguments[arg]; !ok {
			e.err = fmt.Errorf("argument not provided: [%s]", arg)
			return nil
		}
	}
	return templateFile
}
