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
	"strings"

	"os"

	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// ensureStorage checks storage and asks for creation if required
func (e *Engine) ensureStorage() error {
	logger := e.logger.WithField("method", "ensureStorage")
	logger.Debugf("checking [%s]", e.storageDirectory)
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	localStorageDirectory := strings.Replace(e.storageDirectory, "${HOME}", home, 1)
	if stat, err := os.Stat(localStorageDirectory); err != nil || !stat.IsDir() {
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		if stat != nil && !stat.IsDir() {
			return errors.New("storage directory is a file")
		}
		consent := getValue(fmt.Sprintf("Create [%s]? Type yes to do so", localStorageDirectory))
		if strings.ToLower(consent) == "yes" {
			if err := os.Mkdir(localStorageDirectory, 0750); err != nil {
				return err
			}
		} else {
			return errors.New("no consent to create empty directory")
		}
	}
	logger.Debug("directory exists")
	e.storageDirectory = localStorageDirectory
	return nil
}
