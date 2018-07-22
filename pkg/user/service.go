// Package user features user and friends connection management
package user

import (
	"errors"
	"sort"
)

// IRepository interface of repository, it should be able to do load, load+store, and store
type IRepository interface {
	Load(string) (User, error)
	LoadOrStore(string) (User, error)
	Store(string, User) error
}

// Service represents service layer
type Service struct {
	userRepo IRepository
}

// NewService returns a new friend service
func NewService(userRepo IRepository) Service {
	return Service{
		userRepo: userRepo,
	}
}

// AddFriend add a friend, it also subscribes each other's updates.
// It will create a user if not found
func (s *Service) AddFriend(email, email2 string) error {
	p, _ := s.userRepo.LoadOrStore(email)
	p2, _ := s.userRepo.LoadOrStore(email2)
	if p.Blocked[email2] {
		return errors.New(email + " is blocking: " + email2)
	}
	if p2.Blocked[email] {
		return errors.New(email2 + " is blocking: " + email)
	}
	p.Friends[email2] = true
	p.Subscribers[email2] = true
	p2.Friends[email] = true
	p2.Subscribers[email] = true
	s.userRepo.Store(email, p)
	s.userRepo.Store(email2, p2)
	return nil
}

// ListFriends list friends of a user
func (s *Service) ListFriends(email string) ([]string, error) {
	var emails []string
	p, err := s.userRepo.Load(email)
	if err != nil {
		return emails, err
	}
	for email := range p.Friends {
		emails = append(emails, email)
	}
	sort.Strings(emails)
	return emails, err
}

// ListMutualFriends list mutual friends of a user
func (s *Service) ListMutualFriends(email, email2 string) ([]string, error) {
	var mutual []string
	p1, err := s.userRepo.LoadOrStore(email)
	if err != nil {
		return mutual, err
	}
	p2, err := s.userRepo.LoadOrStore(email2)
	if err != nil {
		return mutual, err
	}
	for email := range p1.Friends {
		_, ok := p2.Friends[email]
		if ok {
			mutual = append(mutual, email)
		}
	}
	sort.Strings(mutual)
	return mutual, err
}

// AddFollower add follower into a user, if possible
// If blocked by user, it will throw an error
func (s *Service) AddFollower(email, follower string) error {
	p, err := s.userRepo.LoadOrStore(email)
	if err != nil {
		return err
	}
	if p.Blocked[follower] {
		return errors.New(email + ": has blocked: " + follower)
	}
	p.Subscribers[follower] = true
	return err
}

// AddBlockedUser add user into the block list
func (s *Service) AddBlockedUser(email, blocked string) error {
	p, err := s.userRepo.LoadOrStore(email)
	if err != nil {
		return err
	}
	p.Blocked[blocked] = true
	if p.Subscribers[blocked] {
		delete(p.Subscribers, blocked)
	}
	return err
}

// Notify based on the text given, it will return the emails of active followers
func (s *Service) Notify(email, text string) ([]string, error) {
	var emails []string
	p, err := s.userRepo.Load(email)
	if err != nil {
		return emails, err
	}
	matches := FindEmails(text) // it's possible to have duplicate matches
	for email := range p.Subscribers {
		emails = append(emails, email)
	}
	for _, email := range matches {
		if !p.Subscribers[email] {
			emails = append(emails, email)
		}
	}
	sort.Strings(emails)
	return emails, err
}
