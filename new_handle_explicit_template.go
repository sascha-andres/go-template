package engine

import (
	"io/ioutil"
	"os"
	"path"
)

// handleExplicitTemplate applies templating to a file
func (e *Engine) handleExplicitTemplate(workingDirectory, explicitTemplateFile, name string, arguments map[string]string) {
	if e.err != nil {
		return
	}
	fileToProcess := path.Join(workingDirectory, explicitTemplateFile)
	f, err := os.Open(fileToProcess)
	if err != nil {
		e.err = err
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			e.logger.WithField("method", "handleExplicitTemplate").Errorf("error closing file: [%s]", err.Error())
		}
	}()
	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		e.err = err
		return
	}
	replacedContent, err := applyVariables(string(fileContent), name, arguments)
	if err != nil {
		e.err = err
		return
	}
	stat, err := os.Stat(fileToProcess)
	if err != nil {
		e.err = err
		return
	}
	if err := ioutil.WriteFile(fileToProcess, []byte(replacedContent), stat.Mode()); err != nil {
		e.err = err
		return
	}
}
