package model

import (
	"net/http"
)

// Response is a wrapper response structure
type Response struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(status int, data interface{}) *Response {
	return &Response{
		Status: &ResponseMeta{
			AppStatusCode: status,
			Message:       "reminder successfully added",
		},
		Data: data,
	}
}

// Render for All Responses
func (rd *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
