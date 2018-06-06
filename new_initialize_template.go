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
