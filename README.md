# go-template

Create go projects from templates

## Concept

go-template stores templates locally by cloning them from a git repository. While not necessary, this is required to use management commands.

Each template contains a `.go-template.yml` which contains information like the name of the template, homepage, etc and metadata required to apply the template.

Cloning of the repositories is done by calls to git, so in order to have the management command working, your git installation has to be configured correctly.

## Configuration

TODO

## .go-template.yml

See go-template.md

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

## History

|Version|Description|
|---|---|
|TBD|initial version|