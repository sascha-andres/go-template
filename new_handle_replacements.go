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
	"os"
	"path/filepath"
	"strings"
)

// handleReplacements changes the contents of files in the template by replacements
func (e *Engine) handleReplacements(replacements FromToInformation, workingDirectory, name string, arguments map[string]string) {
	if e.err != nil {
		return
	}
	localTo, err := applyVariables(replacements.To, name, arguments)
	if err != nil {
		e.err = err
		return
	}
	err = filepath.Walk(workingDirectory, func(path string, info os.FileInfo, err error) error {
		if nil != e.templateFile.Transformation.Templates && stringInSlice(strings.Replace(path, workingDirectory+"/", "", 1), e.templateFile.Transformation.Templates) {
			e.logger.WithField("method", "handleReplacements").Debugf("[%s] is a template", path)
			return nil
		}
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", workingDirectory, err)
			return err
		}
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			return replaceInFile(replacements.From, localTo, path)
		}
		return nil
	})
}

func (e *Engine) handleAllReplacements(workingDirectory, name string, arguments map[string]string) {
	if e.err != nil {
		return
	}
	for _, replacements := range e.templateFile.Transformation.Replacements {
		e.handleReplacements(replacements, workingDirectory, name, arguments)
	}
}
