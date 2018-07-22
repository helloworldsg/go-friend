package user

type JSONInput struct {
	Friends   []string `json:"friends"`
	Email     string   `json:"email"`
	Requester string   `json:"requester"`
	Target    string   `json:"target"`
	Sender    string   `json:"sender"`
	Text      string   `json:"text"`
}

type JSONOutput struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type JSONOutputFriends struct {
	JSONOutput
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

type JSONOutputNotify struct {
	JSONOutput
	Recipients []string `json:"recipients"`
}
