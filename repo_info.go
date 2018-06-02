package engine

import "path"

func (e *Engine) InfoRepository(name string) (*Repository, []string, error) {
	if ok, err := e.exists(path.Join(name)); err != nil || !ok {
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, nil
	}
	templateFile, err := e.readTemplateFile(path.Join(e.storageDirectory, name, ".go-template.yml"))
	if err != nil {
		return nil, nil, err
	}
	return &templateFile.Repository, templateFile.Arguments, nil
}
