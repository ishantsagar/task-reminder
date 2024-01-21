package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"tskrm.com/handler"
	"tskrm.com/store"
)

var db *sql.DB
var err error

func Setup() {
	db, err = store.Conn()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	h := handler.Handler{Store: &store.Store{}}

	r.Post("/task", h.CreateTaskReminder) //Create Task

	r.Put("/task", h.UpdateTaskReminder) //Update Task

	r.Delete("/task/{id}", h.DeleteTaskReminder) //Delete Task

	r.Get("/task/{id}", h.GetTaskReminder) //Get Task By ID

	r.Get("/tasks", h.GetAllTasksReminder) //Get All Tasks

	fmt.Println("server listening on port 3000...")
	http.ListenAndServe(":3000", r)
}
