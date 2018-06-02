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

import "github.com/sirupsen/logrus"

type (
	// Engine is the worker implementation
	Engine struct {
		storageDirectory string
		logger           *logrus.Entry
	}

	// TemplateFile contains all relevant information about a template
	TemplateFile struct {
		Repository     Repository      `yaml:"project"`                  // Repository information
		Transformation *Transformation `yaml:"transformation,omitempty"` // Transformation instructions
		InitializeGit  bool            `yaml:"git"`                      // If true new project will be placed in a new git repository
		Arguments      []string        `yaml:"arguments"`                // List of arguments wanted by the template
	}

	// Repository contains information about a repository downloaded
	Repository struct {
		Name        string // Name (given) of repository
		Description string // Description provided by author
		Author      string // Author provided self
		URL         string // URL to remote git repository or homepage
	}

	// Transformation contains information how to transform the template
	Transformation struct {
		ExcludedFiles []string            `yaml:"excluded-files"`         // ExcludedFiles may contain a list of files to not include in the new project
		Renames       []FromToInformation `yaml:"renames,omitempty"`      // Renames is a list of renames in the filesystem, type may be directory or file or nothing
		Replacements  []FromToInformation `yaml:"replacements,omitempty"` // Replacements is a list of replacements within the files, type may be a filename matching regex
		// TODO Templates     []string `yaml:"templates"`      // Templates is a list of files to handle using text/template
	}

	// FromToInformation contains instruction how to change the source
	FromToInformation struct {
		From string  `yaml:"from"`           // From is the part to change
		To   string  `yaml:"to"`             // To is the destination and is treated with text/template
		Type *string `yaml:"type,omitempty"` // Type may be used as an indicator
	}
)

// New creates a new instance to work with go-template's data
func New(storage, logLevel string) (*Engine, error) {
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logrus.SetLevel(lvl)
	entry := logrus.WithField("package", "engine")
	entry.WithField("method", "New").Debugf("called for storage [%s]", storage)
	eng := &Engine{
		storageDirectory: storage,
		logger:           entry,
	}
	err = eng.ensureStorage()
	if err != nil {
		return nil, err
	}
	return eng, nil
}
