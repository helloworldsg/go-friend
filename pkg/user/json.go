package user

import "errors"

type jsonInput struct {
	Friends   []string `json:"friends"`
	Email     string   `json:"email"`
	Requester string   `json:"requester"`
	Target    string   `json:"target"`
	Sender    string   `json:"sender"`
	Text      string   `json:"text"`
}

type jsonOutput struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type jsonOutputFriends struct {
	jsonOutput
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type jsonOutputRecipients struct {
	jsonOutput
	Recipients []string `json:"recipients"`
}

func validateEmail(in jsonInput) error {
	if in.Email == "" {
		return errors.New("empty email address supplied")
	}
	if !ValidateEmail(in.Email) {
		return errors.New("invalid email address: " + in.Email)
	}
	return nil
}

func validateFriends(in jsonInput) error {
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

func validateRequester(in jsonInput) error {
	if in.Requester == "" {
		return errors.New("empty requester supplied")
	}

	if !ValidateEmail(in.Requester) {
		return errors.New("invalid email address: " + in.Requester)
	}
	return nil
}

func validateSender(in jsonInput) error {
	if in.Sender == "" {
		return errors.New("empty sender supplied")
	}

	if !ValidateEmail(in.Sender) {
		return errors.New("invalid email address: " + in.Requester)
	}
	return nil
}

func validateTarget(in jsonInput) error {
	if in.Target == "" {
		return errors.New("empty target supplied")
	}

	if !ValidateEmail(in.Target) {
		return errors.New("invalid email address: " + in.Target)
	}
	return nil
}

func validateText(in jsonInput) error {
	if in.Text == "" {
		return errors.New("empty text supplied")
	}
	return nil
}
