package engine

import (
	"html/template"
	"io/ioutil"
	"os"
	"path"

	"bytes"

	"fmt"

	"path/filepath"

	"strings"

	"github.com/pkg/errors"
	"github.com/sascha-andres/go-template/wrapper"
)

func (e *Engine) New(name, templateName string, arguments map[string]string) error {
	logger := e.logger.WithField("method", "New")
	currentDirectory, err := os.Getwd()
	if err != nil {
		return err
	}
	projectDirectory := path.Join(currentDirectory, name)
	logger.Debugf("working for new project in %s", projectDirectory)
	if _, err := os.Stat(projectDirectory); err != nil && !os.IsNotExist(err) {
		return err
	}
	templateDirectory := path.Join(e.storageDirectory, templateName)
	logger.Debugf("working with templateName in [%s]", templateDirectory)
	if _, err := os.Stat(templateDirectory); err != nil && !os.IsExist(err) {
		return errors.New("no such templateName")
	}
	temporaryDirectory, err := ioutil.TempDir("", "go-templateName")
	if err != nil {
		return err
	}
	defer func() {
		logger.Debugf("removing [%s]", temporaryDirectory)
		if err := os.RemoveAll(temporaryDirectory); err != nil {
			logger.Errorf("unable to clean up: [%s]", err.Error())
		}
	}()
	workingDirectory := path.Join(temporaryDirectory, "work")
	logger.Debugf("creating working directory [%s]", workingDirectory)
	err = copyDir(templateDirectory, workingDirectory)
	if err != nil {
		return err
	}
	logger.Debugf("remove [%s] folder", path.Join(workingDirectory, ".git"))
	err = os.RemoveAll(path.Join(workingDirectory, ".git"))
	if err != nil {
		return err
	}
	templateFile, err := e.readTemplateFile(path.Join(e.storageDirectory, templateName, ".go-template.yml"))
	if err != nil {
		return err
	}
	if templateFile.InitializeGit {
		wrapper.Git("-C", workingDirectory, "init")
		wrapper.Git("-C", workingDirectory, "add", "--all", ":/")
		wrapper.Git("-C", workingDirectory, "commit", "-m", "\"feat: initial commit\"")
		wrapper.Git("-C", workingDirectory, "checkout", "-b", "develop")
	}
	for _, arg := range templateFile.Arguments {
		if _, ok := arguments[arg]; !ok {
			return fmt.Errorf("argument not provided: [%s]", arg)
		}
	}
	for _, excluded := range templateFile.Transformation.ExcludedFiles {
		err = os.RemoveAll(path.Join(workingDirectory, excluded))
		if err != nil {
			logger.Warnf("could not remove [%s]: %s", excluded, err.Error())
		} else {
			logger.Debugf("removed [%s]", excluded)
		}
	}
	if templateFile.InitializeGit {
		wrapper.Git("-C", workingDirectory, "add", "--all", ":/")
		wrapper.Git("-C", workingDirectory, "commit", "-m", "\"feat: removed excluded files from template\"")
	}
	for _, rename := range templateFile.Transformation.Renames {
		textTemplate, err := template.New("rename").Parse(rename.To)
		if err != nil {
			return err
		}
		var result []byte
		buff := bytes.NewBuffer(result)
		textTemplate.Execute(buff, struct {
			Name      string
			Arguments map[string]string
		}{
			Name:      name,
			Arguments: arguments,
		})
		err = os.Rename(path.Join(workingDirectory, rename.From), path.Join(workingDirectory, buff.String()))
		if err != nil {
			logger.Warnf("could not rename [%s]: %s", rename.From, err.Error())
		} else {
			logger.Debugf("renamed [%s]", rename.From)
		}
	}
	if templateFile.InitializeGit {
		wrapper.Git("-C", workingDirectory, "add", "--all", ":/")
		wrapper.Git("-C", workingDirectory, "commit", "-m", "\"feat: rename transformations\"")
	}
	for _, replacements := range templateFile.Transformation.Replacements {
		textTemplate, err := template.New("replacement").Parse(replacements.To)
		if err != nil {
			return err
		}
		var result []byte
		buff := bytes.NewBuffer(result)
		textTemplate.Execute(buff, struct {
			Name      string
			Arguments map[string]string
		}{
			Name:      name,
			Arguments: arguments,
		})
		localFrom, localTo := replacements.From, buff.String()
		err = filepath.Walk(workingDirectory, func(path string, info os.FileInfo, err error) error {
			if nil != templateFile.Transformation.Templates && stringInSlice(strings.Replace(path, workingDirectory+"/", "", 1), templateFile.Transformation.Templates) {
				logger.Debugf("[%s] is a template", path)
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
				return replaceInFile(localFrom, localTo, path)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	if templateFile.InitializeGit {
		wrapper.Git("-C", workingDirectory, "add", "--all", ":/")
		wrapper.Git("-C", workingDirectory, "commit", "-m", "\"feat: replacements in files\"")
	}
	return copyDir(workingDirectory, projectDirectory)
}
