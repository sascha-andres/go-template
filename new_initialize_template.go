package engine

import (
	"os"
	"path"
)

// initializeTemplate copies template into place and removes the git reference
func (e *Engine) initializeTemplate(templateDirectory, workingDirectory string) {
	if e.err != nil {
		return
	}
	logger := e.logger.WithField("method", "initializeTemplate")
	logger.Debugf("creating working directory [%s]", workingDirectory)
	err := copyDir(templateDirectory, workingDirectory)
	if err != nil {
		e.err = err
		return
	}
	logger.Debugf("remove [%s] folder", path.Join(workingDirectory, ".git"))
	e.err = os.RemoveAll(path.Join(workingDirectory, ".git"))
}
