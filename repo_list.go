// Copyright © 2018 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
