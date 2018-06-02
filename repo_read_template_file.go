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

	yaml "gopkg.in/yaml.v2"
)

// readTemplateFile reads a specified template file and returns it
func (e *Engine) readTemplateFile(path string) (*TemplateFile, error) {
	logger := e.logger.WithField("method", "readTemplateFile")
	logger.Debugf("reading [%s]", path)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var templateFile TemplateFile
	err = yaml.Unmarshal(content, &templateFile)
	if err != nil {
		return nil, err
	}
	return &templateFile, err
}
