package engine

import (
	"os"
	"path"
)

// handleRename renames files or directories
func (e *Engine) handleRename(workingDirectory, name string, arguments map[string]string, rename FromToInformation) {
	if e.err != nil {
		return
	}
	logger := e.logger.WithField("method", "handleRename")
	localTo, err := applyVariables(rename.To, name, arguments)
	if err != nil {
		e.err = err
		return
	}
	e.err = os.Rename(path.Join(workingDirectory, rename.From), path.Join(workingDirectory, localTo))
	if err != nil {
		logger.Warnf("could not rename [%s]: %s", rename.From, err.Error())
	} else {
		logger.Debugf("renamed [%s]", rename.From)
	}
}
