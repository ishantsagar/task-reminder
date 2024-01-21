package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"tskrm.com/model"
	"tskrm.com/store"
)

func TestCreateTask(t *testing.T) {
	task := &model.Task{
		Title:       "Lunch",
		Description: "Lunch with Jhonson",
		Priority:    2,
		DateTime:    "2024-02-12T11:45:26.371Z",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(task)
	if err != nil {
		t.Errorf(err.Error())
	}
	req := httptest.NewRequest(http.MethodPost, "/task", &buf)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	h := Handler{Store: &store.MockStore{}}
	h.CreateTaskReminder(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestUpdateTask(t *testing.T) {
	task := &model.Task{
		ID:          1,
		Title:       "Lunch",
		Description: "Lunch with Jhonson",
		Priority:    2,
		DateTime:    "2024-02-12T11:45:26.371Z",
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(task)
	if err != nil {
		t.Errorf(err.Error())
	}
	req := httptest.NewRequest(http.MethodPut, "/task", &buf)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	h := Handler{Store: &store.MockStore{}}
	h.UpdateTaskReminder(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetTaskById(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	h := Handler{Store: &store.MockStore{}}
	h.GetTaskReminder(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestDeleteTaskById(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/task/1", nil)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	h := Handler{Store: &store.MockStore{}}
	h.DeleteTaskReminder(w, req)
	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
