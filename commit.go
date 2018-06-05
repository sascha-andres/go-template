package engine

import "github.com/sascha-andres/go-template/wrapper"

// commit adds all file to the stage and commits them
func commit(workingDirectory, message string) {
	wrapper.Git("-C", workingDirectory, "add", "--all", ":/")
	wrapper.Git("-C", workingDirectory, "commit", "-m", message)
}
