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
	"livingit.de/code/go-template/wrapper"
)

func (e *Engine) New(name, templateName string, arguments map[string]string) error {
	projectDirectory, templateDirectory, workingDirectory, deferFunc := e.setupDirectories(name, templateName)
	if e.err == nil {
		defer func() {
			deferFunc()
		}()
	}
	e.initializeTemplate(templateDirectory, workingDirectory)
	e.loadTemplateFile(templateName, arguments)
	if e.err != nil {
		return e.err
	}
	if e.templateFile.InitializeGit {
		wrapper.Git("-C", workingDirectory, "init")
		e.templateFile.commit(workingDirectory, "\"feat: initial commit\"")
		wrapper.Git("-C", workingDirectory, "checkout", "-b", "develop")
	}
	e.handleExclusions(workingDirectory)
	e.templateFile.commit(workingDirectory, "\"feat: removed excluded files from template\"")
	e.handleRenames(workingDirectory, name, arguments)
	e.templateFile.commit(workingDirectory, "\"feat: rename transformations\"")
	e.handleAllReplacements(workingDirectory, name, arguments)
	e.templateFile.commit(workingDirectory, "\"feat: replacements in files\"")
	e.handleExplicitTemplates(workingDirectory, name, arguments)
	e.templateFile.commit(workingDirectory, "\"feat: handle explicit templates\"")
	if e.err == nil {
		e.err = copyDir(workingDirectory, projectDirectory)
	}
	return e.err
}
