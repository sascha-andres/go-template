package engine

import "github.com/sirupsen/logrus"

type (
	// Engine is the worker implementation
	Engine struct {
		storageDirectory string
		logger           *logrus.Entry
	}

	// Repository contains information about a repository downloaded
	Repository struct {
		Name        string // Name (given) of repository
		Description string // Description provided by author
		Author      string // Author provided self
		URL         string // URL to remote git repository or homepage
	}
)

// New creates a new instance to work with go-template's data
func New(storage, logLevel string) (*Engine, error) {
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logrus.SetLevel(lvl)
	entry := logrus.WithField("package", "engine")
	entry.WithField("method", "New").Debugf("called for storage [%s]", storage)
	eng := &Engine{
		storageDirectory: storage,
		logger:           entry,
	}
	err = eng.ensureStorage()
	if err != nil {
		return nil, err
	}
	return eng, nil
}
