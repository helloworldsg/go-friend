package user

import (
	"encoding/json"
	"errors"
	"net/http"
)

// HTTPHandler web service handler for user and friends management
type HTTPHandler struct {
	*http.ServeMux
	service Service
}

// runFn actual processing of input
type runFn func(jsonInput) (interface{}, error)

// preRunFn preRun function types
type preRunFn func(jsonInput) error

func handlePOST(runFn runFn, validFn ...preRunFn) func(w http.ResponseWriter, r *http.Request) {
	writeErr := func(w http.ResponseWriter, err error) {
		w.WriteHeader(400)
		out := jsonOutput{
			Success: false,
			Error:   err.Error(),
		}
		json.NewEncoder(w).Encode(out)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var in jsonInput
		if r.Method != http.MethodPost {
			writeErr(w, errors.New("method needs to be POST"))
			return
		}
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			writeErr(w, errors.New("unable to decode json"))
			return
		}

		// preRun or validations
		for _, fn := range validFn {
			err := fn(in)
			if err != nil {
				writeErr(w, err)
				return
			}
		}

		// actual run
		obj, err := runFn(in)
		if err != nil {
			writeErr(w, err)
			return
		}
		json.NewEncoder(w).Encode(obj)
	}
}

// NewHTTPHandler creates a new HTTP handler
func NewHTTPHandler() HTTPHandler {
	service := NewService(NewInMemoryRepository())
	mux := http.NewServeMux()
	mux.HandleFunc("/friends/add", handlePOST(func(in jsonInput) (interface{}, error) {
		return jsonOutput{Success: true}, service.AddFriend(in.Friends[0], in.Friends[1])
	}, validateFriends))
	mux.HandleFunc("/friends/list", handlePOST(func(in jsonInput) (interface{}, error) {
		result, err := service.ListFriends(in.Email)
		return jsonOutputFriends{
			jsonOutput: jsonOutput{Success: true},
			Friends:    result,
			Count:      len(result),
		}, err
	}, validateEmail))
	mux.HandleFunc("/friends/mutual", handlePOST(func(in jsonInput) (interface{}, error) {
		result, err := service.ListMutualFriends(in.Friends[0], in.Friends[1])
		return jsonOutputFriends{
			jsonOutput: jsonOutput{Success: true},
			Friends:    result,
			Count:      len(result),
		}, err
	}, validateFriends))
	mux.HandleFunc("/follow", handlePOST(func(in jsonInput) (interface{}, error) {
		return jsonOutput{Success: true}, service.AddFollower(in.Target, in.Requester)
	}, validateRequester, validateTarget))
	mux.HandleFunc("/block", handlePOST(func(in jsonInput) (interface{}, error) {
		return jsonOutput{Success: true}, service.AddBlockedUser(in.Requester, in.Target)
	}, validateRequester, validateTarget))
	mux.HandleFunc("/notify", handlePOST(func(in jsonInput) (interface{}, error) {
		result, err := service.Notify(in.Sender, in.Text)
		return jsonOutputRecipients{
			jsonOutput: jsonOutput{Success: true},
			Recipients: result,
		}, err
	}, validateSender, validateText))
	return HTTPHandler{
		ServeMux: mux,
		service:  service,
	}
}
