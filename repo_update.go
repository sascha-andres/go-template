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
	"fmt"
	"path"

	"io/ioutil"

	"livingit.de/code/go-template/wrapper"
)

// UpdateRepository issues a git pull
func (e *Engine) UpdateRepository(name string) error {
	if ok, _ := e.exists(name); ok {
		_, err := wrapper.Git("-C", path.Join(e.storageDirectory, name), "pull")
		return err
	}
	return fmt.Errorf("no such repository: %s", name)
}

// UpdateRepositories issues a git pull on all template repositories
func (e *Engine) UpdateRepositories() error {
	entries, err := ioutil.ReadDir(e.storageDirectory)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			if ok, _ := e.exists(entry.Name()); ok {
				if err := e.UpdateRepository(entry.Name()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
