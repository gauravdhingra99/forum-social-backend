package middleware

import (
	"encoding/json"
	"net/http"
)

// UnmarshalResponse unmarshals the JSON response.
func UnmarshalResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal marshals the JSON response.
func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Response is the structure for all API responses.
type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Errors  []Error     `json:"errors"`
}

// Error is the structure for all errors returned via the API.
type Error struct {
	Code         string `json:"code"`
	Message      string `json:"message"`
	MessageTitle string `json:"message_title"`
}

// NewErrorResponse creates a new error to be sent as the response.
func NewErrorResponse(code, message, title string) *Response {
	r := &Response{
		Data:    struct{}{},
		Success: false,
		Errors:  []Error{{code, message, title}},
	}
	return r
}

func (r *Response) Write(w http.ResponseWriter, status int) {
	body, err := r.Marshal()
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write(body); err != nil {
	}
}
