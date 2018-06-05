package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// handleReplacements changes the contents of files in the template by replacements
func (e *Engine) handleReplacements(templateFile *TemplateFile, replacements FromToInformation, workingDirectory, name string, arguments map[string]string) {
	if e.err != nil {
		return
	}
	localTo, err := applyVariables(replacements.To, name, arguments)
	if err != nil {
		e.err = err
		return
	}
	err = filepath.Walk(workingDirectory, func(path string, info os.FileInfo, err error) error {
		if nil != templateFile.Transformation.Templates && stringInSlice(strings.Replace(path, workingDirectory+"/", "", 1), templateFile.Transformation.Templates) {
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
