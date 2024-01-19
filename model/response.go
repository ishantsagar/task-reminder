package model

import (
	"net/http"

	"github.com/go-chi/render"
)

type ResponseMeta struct {
	AppStatusCode int    `json:"code"`
	Message       string `json:"statusType,omitempty"`
	ErrorDetail   string `json:"errorDetail,omitempty"`
	ErrorMessage  string `json:"errorMessage,omitempty"`
	DevMessage    string `json:"devErrorMessage,omitempty"`
}

type ErrResponse struct {
	HTTPStatusCode int          `json:"-"` // http response status code
	Status         ResponseMeta `json:"status"`
	AppCode        int64        `json:"code,omitempty"` // application-specific error code
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
