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

// handleExclusion removes files that should not be part of generated project
func (e *Engine) handleExclusion(workingDirectory, excluded string) {
	if e.err != nil {
		return
	}
	logger := e.logger.WithField("method", "handleExclusion")
	e.err = os.RemoveAll(path.Join(workingDirectory, excluded))
	if e.err != nil {
		logger.Warnf("could not remove [%s]", excluded)
	} else {
		logger.Debugf("removed [%s]", excluded)
	}
}

func (e *Engine) handleExclusions(workingDirectory string) {
	if e.err != nil {
		return
	}
	for _, excluded := range e.templateFile.Transformation.ExcludedFiles {
		e.handleExclusion(workingDirectory, excluded)
	}
}
