package engine

import (
	"os"
	"path"
)

// exists checks if a repository exists
func (e *Engine) exists(name string) (bool, error) {
	logger := e.logger.WithField("method", "exists")
	fullPath := path.Join(e.storageDirectory, name)
	logger.Debugf("checking [%s] for repository", fullPath)
	if stat, err := os.Stat(fullPath); err == nil && stat.IsDir() {
		if _, err := os.Stat(path.Join(fullPath, ".go-template.yml")); err == nil {
			return true, nil
		}
	} else {
		if err != nil {
			if !os.IsNotExist(err) {
				logger.Infof("error while checking for directory: %s", err.Error())
				return false, err
			}
		}
	}
	return false, nil
}
