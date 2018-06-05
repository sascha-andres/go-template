package engine

import (
	"io/ioutil"
	"strings"
)

func replaceInFile(from, to, path string) error {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(read), from, to, -1)
	return ioutil.WriteFile(path, []byte(newContents), 0)
}
