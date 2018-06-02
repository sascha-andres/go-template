package engine

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	yaml "gopkg.in/yaml.v1"
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
				logger.Info("found .go-template.yml")
				f, err := os.Open(path.Join(fullPath, ".go-template.yml"))
				if err != nil {
					return nil, err
				}
				defer func() {
					if err := f.Close(); err != nil {
						logger.Error(err.Error())
					}
				}()
				content, err := ioutil.ReadAll(f)
				if err != nil {
					return nil, err
				}
				var templateFile TemplateFile
				err = yaml.Unmarshal(content, &templateFile)
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
