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
	"os"
	"path"
)

// handleExplicitTemplate applies templating to a file
func (e *Engine) handleExplicitTemplate(workingDirectory, explicitTemplateFile, name string, arguments map[string]string) {
	if e.err != nil {
		return
	}
	fileToProcess := path.Join(workingDirectory, explicitTemplateFile)
	f, err := os.Open(fileToProcess)
	if err != nil {
		e.err = err
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			e.logger.WithField("method", "handleExplicitTemplate").Errorf("error closing file: [%s]", err.Error())
		}
	}()
	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		e.err = err
		return
	}
	replacedContent, err := applyVariables(string(fileContent), name, arguments)
	if err != nil {
		e.err = err
		return
	}
	stat, err := os.Stat(fileToProcess)
	if err != nil {
		e.err = err
		return
	}
	if err := ioutil.WriteFile(fileToProcess, []byte(replacedContent), stat.Mode()); err != nil {
		e.err = err
		return
	}
}

func (e *Engine) handleExplicitTemplates(workingDirectory, name string, arguments map[string]string) {
	if e.err != nil {
		return
	}
	for _, explicitTemplateFile := range e.templateFile.Transformation.Templates {
		e.handleExplicitTemplate(workingDirectory, explicitTemplateFile, name, arguments)
	}
}
