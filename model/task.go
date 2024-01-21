package model

import (
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	DateTime    string `json:"datetime"`
}

func (t *Task) Bind(r *http.Request) error {
	return nil
}

func (t *Task) Validate() error {
	if len(t.Title) == 0 {
		return fmt.Errorf("%s", "task should have a title")
	} else if len(t.Description) == 0 {
		return fmt.Errorf("%s", "task should have a description")
	} else if t.Priority < 1 || t.Priority > 3 {
		return fmt.Errorf("%s", "task should have a priority either of 1, 2 or 3 only. (1 = High, 2 = Medium, 3 = Low)")
	} else if _, err := time.Parse(time.RFC3339, t.DateTime); err != nil {
		return fmt.Errorf("%s", "task should have a valid datetime format like '2014-11-12T11:45:26.371Z'")
	}
	return nil
}
