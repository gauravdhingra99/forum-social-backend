package models

import (
	"encoding/json"
	"net/http"
)

const (
	ErrInvalidRequestBody  string = "forum:invalid_request_body"
	ErrInternalServerError string = "forum:internal_server_error"
	TooManyRequestError    string = "forum:too_many_request"
)

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

// NewSuccessResponse creates a new success response to be send via the API.
func NewSuccessResponse(data interface{}) *Response {
	r := &Response{
		Data:    data,
		Success: true,
		Errors:  []Error{},
	}
	return r
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
	_, err = w.Write(body)
	if err != nil {
		return
	}
}

// Marshal marshals the JSON response.
func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// CallbackServiceResponse is the structure for the response from CallbackService
type CallbackServiceResponse struct {
	Success    bool
	ErrorCode  int
	Error      error
	ErrorTitle string
}
