package user

import (
	"errors"
)

func validateEmail(in JSONInput) error {
	if in.Email == "" {
		return errors.New("empty email address supplied")
	}
	if !ValidateEmail(in.Email) {
		return errors.New("invalid email address: " + in.Email)
	}
	return nil
}

func validateFriends(in JSONInput) error {
	if len(in.Friends) != 2 {
		return errors.New("need at least 2 email addresses in the parameter")
	}

	for _, email := range in.Friends {
		if !ValidateEmail(email) {
			return errors.New("invalid email address: " + email)
		}
	}
	return nil
}

func validateRequester(in JSONInput) error {
	if in.Requester == "" {
		return errors.New("empty requester supplied")
	}

	if !ValidateEmail(in.Requester) {
		return errors.New("invalid email address: " + in.Requester)
	}
	return nil
}

func validateSender(in JSONInput) error {
	if in.Sender == "" {
		return errors.New("empty sender supplied")
	}

	if !ValidateEmail(in.Sender) {
		return errors.New("invalid email address: " + in.Requester)
	}
	return nil
}

func validateTarget(in JSONInput) error {
	if in.Target == "" {
		return errors.New("empty target supplied")
	}

	if !ValidateEmail(in.Target) {
		return errors.New("invalid email address: " + in.Target)
	}
	return nil
}

func validateText(in JSONInput) error {
	if in.Text == "" {
		return errors.New("empty text supplied")
	}
	return nil
}
