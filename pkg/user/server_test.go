package user_test

import (
	"github.com/hmoniaga/go-friend/pkg/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPHandler_EndToEnd(t *testing.T) {
	successReturn := `{"success":true,"error":""}`
	tests := []struct {
		name     string
		method   string
		path     string
		input    string
		want     string
		wantCode int
	}{
		{"andy and john became friends", http.MethodPost, "/friends/add", `{"friends":["andy@example.com","john@example.com"]}`, successReturn, http.StatusOK},
		{"list all friends", http.MethodPost, "/friends/list", `{"email":"andy@example.com"}`, `{"success":true,"error":"","friends":["john@example.com"],"count":1}`, http.StatusOK},
		{"andy and ben became friends", http.MethodPost, "/friends/add", `{"friends":["andy@example.com","ben@example.com"]}`, successReturn, http.StatusOK},
		{"john and ben became friends", http.MethodPost, "/friends/add", `{"friends":["john@example.com","ben@example.com"]}`, successReturn, http.StatusOK},
		{"confirm ben is a mutual friend", http.MethodPost, "/friends/mutual", `{"friends":["andy@example.com","john@example.com"]}`, `{"success":true,"error":"","friends":["ben@example.com"],"count":1}`, http.StatusOK},
		{"lisa subscribes to john", http.MethodPost, "/follow", `{"requester":"lisa@example.com","target":"john@example.com"}`, successReturn, http.StatusOK},
		{"andy blocks john", http.MethodPost, "/block", `{"requester":"andy@example.com","target":"john@example.com"}`, successReturn, http.StatusOK},
		{"john blocks ben", http.MethodPost, "/block", `{"requester":"john@example.com","target":"ben@example.com"}`, successReturn, http.StatusOK},
		{"john notifies his subscribers", http.MethodPost, "/notify", `{"sender":"john@example.com","text":"Hello World! kate@example.com"}`, `{"success":true,"error":"","recipients":["andy@example.com","kate@example.com","lisa@example.com"]}`, http.StatusOK},
		{"andy notifies his subscribers", http.MethodPost, "/notify", `{"sender":"andy@example.com","text":"Hello World!"}`, `{"success":true,"error":"","recipients":["ben@example.com"]}`, http.StatusOK},
		{"john blocks ashley", http.MethodPost, "/block", `{"requester":"john@example.com","target":"ashley@example.com"}`, successReturn, http.StatusOK},
		{"ashley tried to friend john", http.MethodPost, "/friends/add", `{"friends":["ashley@example.com","john@example.com"]}`, `{"success":false,"error":"john@example.com is blocking: ashley@example.com"}`, http.StatusBadRequest},
		{"ashley tried to subscribe to john", http.MethodPost, "/follow", `{"requester":"ashley@example.com", "target":"john@example.com"}`, `{"success":false,"error":"john@example.com: has blocked: ashley@example.com"}`, http.StatusBadRequest},
	}

	handler := user.NewHTTPHandler()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(tt.method, tt.path, strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status wantCode: got %v want %v",
					status, http.StatusOK)
			}
			if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(tt.want) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.want)
			}
		})
	}
}

func TestHTTPHandler_Functional(t *testing.T) {
	successReturn := `{"success":true,"error":""}`
	tests := []struct {
		name     string
		method   string
		path     string
		input    string
		want     string
		wantCode int
	}{
		{"handle invalid method", http.MethodGet, "/friends/add", ``, `{"success":false,"error":"method needs to be POST"}`, http.StatusBadRequest},
		{"handle invalid JSON", http.MethodPost, "/friends/add", `asdfasdf`, `{"success":false,"error":"unable to decode json"}`, http.StatusBadRequest},

		{"can create friends", http.MethodPost, "/friends/add", `{"friends":["andy@example.com","john@example.com"]}`, successReturn, http.StatusOK},
		{"can list all friends", http.MethodPost, "/friends/list", `{"email":"andy@example.com"}`, `{"success":true,"error":"","friends":["john@example.com"],"count":1}`, http.StatusOK},
		{"handle list unknown user's friends", http.MethodPost, "/friends/list", `{"email":"unknown@example.com"}`, `{"success":false,"error":"user not found"}`, http.StatusBadRequest},

		{"can create friends", http.MethodPost, "/friends/add", `{"friends":["andy@example.com","common@example.com"]}`, successReturn, http.StatusOK},
		{"can create friends", http.MethodPost, "/friends/add", `{"friends":["john@example.com","common@example.com"]}`, successReturn, http.StatusOK},
		{"can list mutual friends", http.MethodPost, "/friends/mutual", `{"friends":["andy@example.com","john@example.com"]}`, `{"success":true,"error":"","friends":["common@example.com"],"count":1}`, http.StatusOK},
		{"can subscribe updates", http.MethodPost, "/follow", `{"requester":"andy@example.com","target":"john@example.com"}`, successReturn, http.StatusOK},
		{"can block updates", http.MethodPost, "/block", `{"requester":"andy@example.com","target":"john@example.com"}`, successReturn, http.StatusOK},
		{"can notify subscribers", http.MethodPost, "/notify", `{"sender":"john@example.com","text":"Hello World! kate@example.com"}`, `{"success":true,"error":"","recipients":["andy@example.com","common@example.com","kate@example.com"]}`, http.StatusOK},
		{"handles empty input: friends", http.MethodPost, "/friends/add", `{}`, `{"success":false,"error":"need at least 2 email addresses in the parameter"}`, http.StatusBadRequest},
		{"handles empty input: email", http.MethodPost, "/friends/list", `{}`, `{"success":false,"error":"empty email address supplied"}`, http.StatusBadRequest},
		{"handles empty input: requester", http.MethodPost, "/follow", `{}`, `{"success":false,"error":"empty requester supplied"}`, http.StatusBadRequest},
		{"handles empty input: requester", http.MethodPost, "/block", `{}`, `{"success":false,"error":"empty requester supplied"}`, http.StatusBadRequest},
		{"handles empty input: sender", http.MethodPost, "/notify", `{}`, `{"success":false,"error":"empty sender supplied"}`, http.StatusBadRequest},
		{"handles invalid input: friends", http.MethodPost, "/friends/add", `{"friends":["asdf.com","fdsa.com"]}`, `{"success":false,"error":"invalid email address: asdf.com"}`, http.StatusBadRequest},
		{"handles invalid input: email", http.MethodPost, "/friends/add", `{"friends":["asdf.com","fdsa.com"]}`, `{"success":false,"error":"invalid email address: asdf.com"}`, http.StatusBadRequest},
		{"handles invalid input: create friends", http.MethodPost, "/friends/add", `{"friends":["","common@example.com"]}`, `{"success":false,"error":"invalid email address: "}`, http.StatusBadRequest},
		{"handles invalid input: list all friends", http.MethodPost, "/friends/list", `{"email":"asdf"}`, `{"success":false,"error":"invalid email address: asdf"}`, http.StatusBadRequest},
		{"handles invalid input: list mutual friends", http.MethodPost, "/friends/mutual", `{"friends":["","john@example.com"]}`, `{"success":false,"error":"invalid email address: "}`, http.StatusBadRequest},
		{"handles invalid input: subscribe updates", http.MethodPost, "/follow", `{"requester":"fsdadf","target":"john@example.com"}`, `{"success":false,"error":"invalid email address: fsdadf"}`, http.StatusBadRequest},
		{"handles invalid input: block updates", http.MethodPost, "/block", `{"requester":" asdf","target":"john@example.com"}`, `{"success":false,"error":"invalid email address:  asdf"}`, http.StatusBadRequest},
		{"handles invalid input: block updates", http.MethodPost, "/block", `{"requester":"john@example.com","target":"asdf"}`, `{"success":false,"error":"invalid email address: asdf"}`, http.StatusBadRequest},
		{"handles invalid input: block updates", http.MethodPost, "/block", `{"requester":"john@example.com","target":""}`, `{"success":false,"error":"empty target supplied"}`, http.StatusBadRequest},
		{"handles invalid input: notify subscribers", http.MethodPost, "/notify", `{"sender":" asdf","text":"Hello World! kate@example.com"}`, `{"success":false,"error":"invalid email address: "}`, http.StatusBadRequest},
		{"handles invalid input: notify subscribers", http.MethodPost, "/notify", `{"sender":"john@example.com","text":""}`, `{"success":false,"error":"empty text supplied"}`, http.StatusBadRequest},
	}

	handler := user.NewHTTPHandler()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(tt.method, tt.path, strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status wantCode: got %v want %v",
					status, http.StatusOK)
			}
			if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(tt.want) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.want)
			}
		})
	}
}
