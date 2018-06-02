package engine

import (
	"io/ioutil"
	"os"

	"path"

	"github.com/pkg/errors"
	"github.com/sascha-andres/go-template/wrapper"
)

// AddRepository clones a repository and moves it into storage if it is a template repository
func (e *Engine) AddRepository(url string) (*Repository, error) {
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
	wrapper.Git("clone", url, temporaryDirectory)
	if _, err := os.Stat(path.Join(temporaryDirectory, ".go-template.yml")); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("not a go-template repository")
		}
		return nil, err
	}
	templateFile, err := e.readTemplateFile(path.Join(temporaryDirectory, ".go-template.yml"))
	if err != nil {
		return nil, err
	}
	if ok, err := e.exists(templateFile.Repository.Name); ok || err != nil {
		if ok {
			// TODO [resolve conflicting names]
			return nil, errors.New("template already known")
		}
		return nil, err
	}
	err = os.Rename(temporaryDirectory, path.Join(e.storageDirectory, templateFile.Repository.Name))
	if err != nil {
		return nil, err
	}
	return &templateFile.Repository, nil
}
