package engine

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

// ListRepositories lists all repositories downloaded
func (e *Engine) ListRepositories() ([]*Repository, error) {
	logger := e.logger.WithField("method", "ListRepositories")
	result := make([]*Repository, 0)
	entries, err := ioutil.ReadDir(e.storageDirectory)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			fullPath := path.Join(e.storageDirectory, entry.Name())
			logger.Infof("looking in [%s]", fullPath)
			if _, err := os.Stat(path.Join(fullPath, ".go-template.yml")); err == nil {
				logger.Debug("found .go-template.yml")
				templateFile, err := e.readTemplateFile(path.Join(fullPath, ".go-template.yml"))
				if err != nil {
					return nil, err
				}
				result = append(result, &templateFile.Repository)
			} else {
				logger.Warnf("no template in [%s]", fullPath)
			}
		}
	}
	return result, nil
}
