package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"tskrm.com/model"
	"tskrm.com/store"
)

type Handler struct {
	Store store.TaskOps
}

func (h *Handler) CreateTaskReminder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := task.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id, err := h.Store.CreateTaskReminder(&task); err == nil {
		task.ID = id
		render.Status(r, http.StatusCreated)
		render.Render(w, r, model.NewSuccessResponse(http.StatusCreated, task))
	} else {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

func (h *Handler) UpdateTaskReminder(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if task.ID <= 0 {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.Store.UpdateTaskReminder(&task); err == nil {
		render.Status(r, http.StatusOK)
		render.Render(w, r, model.NewSuccessResponse(http.StatusOK, task))
	} else {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

func (h *Handler) DeleteTaskReminder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.Store.DeleteTaskReminder(id); err == nil {
		render.Status(r, http.StatusOK)
		render.Render(w, r, model.NewSuccessResponse(http.StatusOK, nil))
	} else {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

func (h *Handler) GetAllTasksReminder(w http.ResponseWriter, r *http.Request) {
	if tasks, err := h.Store.GetTaskReminders(); err == nil {
		render.Status(r, http.StatusOK)
		render.Render(w, r, model.NewSuccessResponse(http.StatusCreated, tasks))
	} else {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

func (h *Handler) GetTaskReminder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}
	if task, err := h.Store.GetTaskReminder(id); err == nil {
		render.Status(r, http.StatusOK)
		render.Render(w, r, model.NewSuccessResponse(http.StatusCreated, task))
	} else {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}
