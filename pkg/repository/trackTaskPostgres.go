package repository

import (
	"fmt"
	track "timeTracker"

	"github.com/jmoiron/sqlx"
)

type TrackTaskPostgres struct {
	db *sqlx.DB
}

func NewTrackTaskPostgres(db *sqlx.DB) *TrackTaskPostgres {
	return &TrackTaskPostgres{db: db}
}

func (t *TrackTaskPostgres) CreateTask(userId int, task track.Task) (int, error) {
	var taskId int
	query := fmt.Sprintf("INSERT INTO %s (user_id, name, description) values ($1, $2, $3) RETURNING id", taskTable)

	row := t.db.QueryRow(query, userId, task.Name, task.Description)
	err := row.Scan(&taskId)

	return taskId, err
}

func (t *TrackTaskPostgres) DeleteTask(taskId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", taskTable)

	_, err := t.db.Exec(query, taskId)
	return err
}

func (t *TrackTaskPostgres) StartTask(taskId int) error {
	query := fmt.Sprintf("UPDATE %s SET start_time = CURRENT_TIME AT TIME ZONE 'Europe/Moscow' WHERE id = $1", taskTable)

	_, err := t.db.Exec(query, taskId)
	return err
}

func (t *TrackTaskPostgres) EndTask(taskId int) error {
	query := fmt.Sprintf("UPDATE %s SET end_time = CURRENT_TIME AT TIME ZONE 'Europe/Moscow', done = TRUE, diff_time = (CURRENT_TIME AT TIME ZONE 'Europe/Moscow' - start_time) WHERE id = $1", taskTable)

	_, err := t.db.Exec(query, taskId)
	return err
}

func (t *TrackTaskPostgres) GetTasksByLaborCosts(userId int) ([]track.Task, error) {
	var tasks []track.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND done = TRUE ORDER BY diff_time DESC", taskTable)

	err := t.db.Select(&tasks, query, userId)

	return tasks, err
}
