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
