package engine

// ListRepositories lists all repositories downloaded
func (e *Engine) ListRepositories() ([]*Repository, error) {
	e.logger.Info(e.storageDirectory)
	return nil, nil
}
