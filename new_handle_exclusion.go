package engine

import (
	"os"
	"path"
)

// handleExclusion removes files that should not be part of generated project
func (e *Engine) handleExclusion(workingDirectory, excluded string) {
	if e.err != nil {
		return
	}
	logger := e.logger.WithField("method", "handleExclusion")
	e.err = os.RemoveAll(path.Join(workingDirectory, excluded))
	if e.err != nil {
		logger.Warnf("could not remove [%s]", excluded)
	} else {
		logger.Debugf("removed [%s]", excluded)
	}
}
