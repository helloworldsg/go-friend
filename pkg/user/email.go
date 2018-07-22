package user

import (
	"regexp"
	"sort"
)

var emailRe = regexp.MustCompile("[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*")

// FindEmails return unique matches of email from text
func FindEmails(msg string) []string {
	emails := make(map[string]bool)
	for _, m := range emailRe.FindAllStringSubmatch(msg, -1) {
		emails[m[0]] = true
	}
	var result []string
	for email := range emails {
		result = append(result, email)
	}
	sort.Strings(result)
	return result
}

// ValidateEmail returns true if email is valid
func ValidateEmail(email string) bool {
	return emailRe.MatchString(email)
}
