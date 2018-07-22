// Package user features user and friends connection management
package user

// User represents a registered user based on unique identifier: email
type User struct {
	Email       string
	Friends     map[string]bool
	Blocked     map[string]bool
	Followers   map[string]bool
	Subscribers map[string]bool
}

// NewUser creates a new user based on unique email
func NewUser(email string) User {
	return User{
		Email:       email,
		Friends:     make(map[string]bool),
		Blocked:     make(map[string]bool),
		Followers:   make(map[string]bool),
		Subscribers: make(map[string]bool),
	}
}
