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

	"github.com/pkg/errors"
	"livingit.de/code/go-template/wrapper"
)

// AddRepository clones a repository and moves it into storage if it is a template repository
func (e *Engine) AddRepository(url, branch string) (*Repository, error) {
	logger := e.logger.WithField("method", "AddRepository")
	logger.Debugf("add from [%s]", url)
	temporaryDirectory, err := ioutil.TempDir("", "go-template")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := os.RemoveAll(temporaryDirectory); err != nil {
			logger.Errorf("error removing temporary directory: %s", err.Error())
		}
	}()
	if "" != branch {
		wrapper.Git("clone", "-b", branch, url, temporaryDirectory)
	} else {
		wrapper.Git("clone", url, temporaryDirectory)
	}
	if _, err := os.Stat(path.Join(temporaryDirectory, ".go-template.yml")); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("not a go-template repository")
		}
		return nil, err
	}
	e.readTemplateFile(path.Join(temporaryDirectory, ".go-template.yml"))
	if e.err != nil {
		return nil, err
	}
	if ok, err := e.exists(e.templateFile.Repository.Name); ok || err != nil {
		if ok {
			// TODO [resolve conflicting names]
			return nil, errors.New("template already known")
		}
		return nil, err
	}
	err = copyDir(temporaryDirectory, path.Join(e.storageDirectory, e.templateFile.Repository.Name))
	if err != nil {
		return nil, err
	}
	return &e.templateFile.Repository, nil
}
