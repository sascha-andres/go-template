package engine

import (
	"io/ioutil"
	"os"
	"path"
)

// setupDirectories creates all necessary directories
func (e *Engine) setupDirectories(name, templateName string) (string, string, string, func()) {
	if e.err != nil {
		return "", "", "", nil
	}
	logger := e.logger.WithField("method", "setupDirectories")
	currentDirectory, err := os.Getwd()
	if err != nil {
		e.err = nil
		return "", "", "", nil
	}
	projectDirectory := path.Join(currentDirectory, name)
	logger.Debugf("working for new project in %s", projectDirectory)
	if _, err := os.Stat(projectDirectory); err != nil && !os.IsNotExist(err) {
		e.err = nil
		return "", "", "", nil
	}
	templateDirectory := path.Join(e.storageDirectory, templateName)
	logger.Debugf("working with templateName in [%s]", templateDirectory)
	if _, err := os.Stat(templateDirectory); err != nil && !os.IsExist(err) {
		e.err = nil
		return "", "", "", nil
	}
	temporaryDirectory, err := ioutil.TempDir("", "go-templateName")
	if err != nil {
		e.err = nil
		return "", "", "", nil
	}
	return projectDirectory, templateDirectory, path.Join(temporaryDirectory, "work"), func() {
		logger.Debugf("removing [%s]", temporaryDirectory)
		if err := os.RemoveAll(temporaryDirectory); err != nil {
			logger.Errorf("unable to clean up: [%s]", err.Error())
		}
	}
}
