package store

import (
	"database/sql"
	"fmt"

	"tskrm.com/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "TaskDB"
)

var db *sql.DB

type TaskOps interface {
	CreateTaskReminder(task *model.Task) (int64, error)
	UpdateTaskReminder(task *model.Task) error
	DeleteTaskReminder(id string) error
	GetTaskReminder(id string) (*model.Task, error)
	GetTaskReminders() ([]*model.Task, error)
}

type Store struct{}

func Conn() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	return db, err
}

func (s *Store) CreateTaskReminder(task *model.Task) (int64, error) {
	insertStmt := `insert into "Task"("Title", "Description", "Priority", "DateTime") values($1,$2,$3,$4)`
	result, err := db.Exec(insertStmt, task.Title, task.Description, task.Priority, task.DateTime)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *Store) UpdateTaskReminder(task *model.Task) error {
	updateStmt := `update "Task" set "Title"=$1, "Description"=$2, "Priority"=$3 where "id"=$4`
	_, err := db.Exec(updateStmt, task.Title, task.Description, task.Priority, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteTaskReminder(id string) error {
	deleteStmt := `delete from "Task" where id=$1`
	_, err := db.Exec(deleteStmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetTaskReminder(id string) (*model.Task, error) {
	var t model.Task
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM Task where id=%s", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&t)
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}

func (s *Store) GetTaskReminders() ([]*model.Task, error) {
	var t model.Task
	var tasks []*model.Task
	rows, err := db.Query(`SELECT * FROM "Task"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&t)
		if err != nil {
			return nil, err
		} else {
			tasks = append(tasks, &t)
		}
	}
	return tasks, nil
}
