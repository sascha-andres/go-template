package engine

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func replaceInFile(from, to, path string) error {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(read), from, to, -1)
	if strings.Contains(path, "serve") {
		fmt.Println(from)
		fmt.Println(to)
		fmt.Println(string(read))
		fmt.Println(newContents)
	}
	return ioutil.WriteFile(path, []byte(newContents), 0)
}
