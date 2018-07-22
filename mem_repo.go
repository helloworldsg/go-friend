package friend

import (
	"errors"
	"sync"
)

type InMemoryRepo struct {
	*sync.RWMutex
	items map[string]User
}

// NewInMemoryRepository returns new in memory object repository
func NewInMemoryRepository() InMemoryRepo {
	return InMemoryRepo{RWMutex: &sync.RWMutex{}, items: make(map[string]User)}
}

// LoadOrStore returns the existing profile for the email if present.
// Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
func (repo InMemoryRepo) LoadOrStore(key string) (User, error) {
	repo.Lock()
	result, ok := repo.items[key]
	if !ok {
		result = NewUser(key)
		repo.items[key] = result
	}
	repo.Unlock()
	return result, nil
}

// Load returns the value stored in the map for a email, or nil if no value is present.
// The ok result indicates whether email was found in the map.
func (repo InMemoryRepo) Load(key string) (User, error) {
	repo.RLock()
	result, ok := repo.items[key]
	repo.RUnlock()

	if !ok {
		return result, errors.New("user not found")
	}
	return result, nil
}

// Store updates the value of a user
func (repo InMemoryRepo) Store(key string, user User) error {
	repo.Lock()
	repo.items[key] = user
	repo.Unlock()
	return nil
}
