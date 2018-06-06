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
