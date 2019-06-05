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

// handleExclusions removes files from the template that should not be part of the project
func (e *Engine) handleExclusions(workingDirectory string) {
	if e.err != nil {
		return
	}
	handle := func(workingDirectory, excluded string) error {
		logger := e.logger.WithField("method", "handleExclusion")
		e.err = os.RemoveAll(path.Join(workingDirectory, excluded))
		if e.err != nil {
			logger.Warnf("could not remove [%s]", excluded)
		} else {
			logger.Debugf("removed [%s]", excluded)
		}
		return e.err
	}
	for _, excluded := range e.templateFile.Transformation.ExcludedFiles {
		if err := handle(workingDirectory, excluded); err != nil {
			return
		}
	}
}
