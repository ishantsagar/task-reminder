package httphandler

import (
	"net/http"

	"tskrm.com/model"
)

// Response is a wrapper response structure
type Response struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(status int, data interface{}) *Response {
	return &Response{
		Status: &model.ResponseMeta{
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
