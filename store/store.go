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

type Store struct{}

func Conn() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	return db, err
}

func (s *Store) CreateTask(task *model.Task) (int64, error) {
	insertStmt := `insert into "Task"("Title", "Description", "Priority", "DateTime") values($1,$2,$3,$4)`
	result, err := db.Exec(insertStmt, task.Title, task.Description, task.Priority, task.DateTime)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}
