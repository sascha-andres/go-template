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
