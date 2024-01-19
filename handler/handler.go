package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"tskrm.com/httphandler"
	"tskrm.com/model"
	"tskrm.com/store"
)

type Handler struct {
	Store *store.Store
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := render.Bind(r, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err := task.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id, err := h.Store.CreateTask(&task); err == nil {
		task.ID = id
		render.Status(r, http.StatusOK)
		render.Render(w, r, httphandler.NewSuccessResponse(http.StatusCreated, task))
	} else {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}
