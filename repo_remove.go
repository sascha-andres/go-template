package engine

import (
	"fmt"
	"os"
	"path"
)

// RemoveRepository deletes a repository from the local machine
func (e *Engine) RemoveRepository(name string) error {
	if ok, _ := e.exists(name); ok {
		return os.RemoveAll(path.Join(e.storageDirectory, name))
	} else {
		return fmt.Errorf("no such repository: %s", name)
	}
}
