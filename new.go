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
	templateFile := e.loadTemplateFile(templateName, arguments)
	if e.err == nil && templateFile.InitializeGit {
		wrapper.Git("-C", workingDirectory, "init")
		commit(workingDirectory, "\"feat: initial commit\"")
		wrapper.Git("-C", workingDirectory, "checkout", "-b", "develop")
	}
	if e.err == nil {
		for _, excluded := range templateFile.Transformation.ExcludedFiles {
			e.handleExclusion(workingDirectory, excluded)
		}
		if templateFile.InitializeGit {
			commit(workingDirectory, "\"feat: removed excluded files from template\"")
		}
		for _, rename := range templateFile.Transformation.Renames {
			e.handleRename(workingDirectory, name, arguments, rename)
		}
		if templateFile.InitializeGit {
			commit(workingDirectory, "\"feat: rename transformations\"")
		}
		for _, replacements := range templateFile.Transformation.Replacements {
			e.handleReplacements(templateFile, replacements, workingDirectory, name, arguments)
		}
		if templateFile.InitializeGit {
			commit(workingDirectory, "\"feat: replacements in files\"")
		}
		for _, explicitTemplateFile := range templateFile.Transformation.Templates {
			e.handleExplicitTemplate(workingDirectory, explicitTemplateFile, name, arguments)
		}
		if templateFile.InitializeGit {
			commit(workingDirectory, "\"feat: handle explicit templates\"")
		}
		e.err = copyDir(workingDirectory, projectDirectory)
	}
	return e.err
}
