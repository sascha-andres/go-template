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
)

// loadTemplateFile reads template configuration and does a validation run
func (e *Engine) loadTemplateFile(templateName string, arguments map[string]string) {

	err := e.readTemplateFile(path.Join(e.storageDirectory, templateName, ".go-template.yml"))
	if err != nil {
		e.err = err
		return
	}
	for _, arg := range e.templateFile.Arguments {
		fmt.Println("Testing " + arg)
		if _, ok := arguments[arg]; !ok {
			e.err = fmt.Errorf("argument not provided: [%s]", arg)
			break
		}
	}
}
