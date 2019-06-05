# go-template

[![Go Report Card](https://goreportcard.com/badge/github.com/sascha-andres/go-template)](https://goreportcard.com/report/github.com/sascha-andres/go-template) [![codebeat badge](https://codebeat.co/badges/530c668a-42cb-41f5-a8e5-5284f606aa5e)](https://codebeat.co/projects/github-com-sascha-andres-go-template-master) [![Maintainability](https://api.codeclimate.com/v1/badges/f3fd3deb5a43e8acf536/maintainability)](https://codeclimate.com/github/sascha-andres/go-template/maintainability) [![CodeFactor](https://www.codefactor.io/repository/github/sascha-andres/go-template/badge)](https://www.codefactor.io/repository/github/sascha-andres/go-template)

Create go projects from templates

## Concept

go-template stores templates locally by cloning them from a git repository. While not necessary, this is required to use management commands.

Each template contains a `.go-template.yml` which contains information like the name of the template, homepage, etc and metadata required to apply the template.

Cloning of the repositories is done by calls to git, so in order to have the management command working, your git installation has to be configured correctly.

## Installation

### Source code

    go get livingit.de/code/go-template

## Configuration

There is one global configuration, the storage. It is a directorty where the templates are stored and defaults to `/.go-template`. You can override it using the commandline, a `~/.go-template.yml` containing the storage setting.

## .go-template.yml

Placed in the root of the template repository this file enables go-template to use the repository as a template

    project:
      name: "go-rest"                                                    # Name your template
      description: "rest service implementation inspired by Mat Ryer"    # A longer description
      author: "Sascha Andres"                                            # Your name or handle
      repository: "https://github.com/sascha-andres/go-rest"             # homepage/repository
    
    transformation:
      templates:                                                         # List all files that are based on Go templates
        - README.md
      excluded-files:                                                    # List all files that shall not be part of new project
        - ".go-template.yml"
      renames:                                                           # rename files/folders
        - from: "go-rest"
          to: "{{ .Name }}"
      replacements:                                                      # Replace content in files, templates will not be treated
        - from: "go_rest"
          to: "{{ .Name }}"
        - from: "github.com/sascha-andres/go-rest/go-rest/cmd"
          to: "{{ index .Arguments \"Namespace\" }}/{{ .Name }}/{{ .Name }}/cmd"
        - from: "github.com/sascha-andres/go-rest"
          to: "{{ index .Arguments \"Namespace\" }}/{{ .Name }}"
        - from: "go-rest"
          to: "{{ .Name }}"
    
    git: true                                                            # Initialize a git repository for new project
    
    arguments:                                                           # List of arguments required to run template
      - Namespace

## Usage

### Repository management

#### Add a repository

    go-template repo add --url <git-url>

#### Update all repositories

Iterate over all local repositories and issue a `git pull`

    go-template repo update

#### Update a repository

Issue a `git pull` for a specific repository

    go-template repo update --limit-to <name>

#### List all local repositories

Print out a tabular list of repositories installed locally

    go-template repo list

#### Info for a repository

Print out information for a specific repository. Useful to learn about arguments required to run the template

    go-template repo info --name <name>

### Apply

    go-template new --template <name> --name <project-name> --arguments [key=value],[key=value],...

Providing an argument more than one time will be considered an error

## Template repositories

### https://github.com/sascha-andres/go-rest

A rest service based on gin with prometheus metrics

## Acknowledgements

Using code from https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

## History

|Version|Description|
|---|---|
|1.0.2|- remove quotes around commit messages|
||- version info|
|1.0.1|add version command|
|1.0.0|switch to using go modules|
|0.3.4|code improvements|
|0.3.3|license information|
|0.3.2|add goreleaser configuration|
|0.3.1|remove some old url references|
| |update readme for binary distribution without go get|
|0.3.0|add vanity url|
|0.2.1|add badges to README|
|0.2.0|add vendoring|
| |add explicit templates|
| |add .go-template.yml description|
|0.1.0|initial version|
