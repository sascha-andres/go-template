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

import "path"

// InfoRepository returns information about a repository
func (e *Engine) InfoRepository(name string) (*Repository, []string, error) {
	if ok, err := e.exists(path.Join(name)); err != nil || !ok {
		return nil, nil, err
	}
	err := e.readTemplateFile(path.Join(e.storageDirectory, name, ".go-template.yml"))
	if err != nil {
		return nil, nil, err
	}
	return &e.templateFile.Repository, e.templateFile.Arguments, nil
}
