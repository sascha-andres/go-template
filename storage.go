package engine

import "github.com/pkg/errors"

// ensureStorage checks storage and asks for creation if required
func (e *Engine) ensureStorage() error {
	e.logger.WithField("method", "ensureStorage").Debugf("checking [%s]", e.storageDirectory)
	return errors.New("no storage implementation done yet")
}
