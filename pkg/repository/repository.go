package repository

import (
	track "timeTracker"

	"github.com/jmoiron/sqlx"
)

type TrackPeople interface {
	CreateUser(people track.People) (int, error)
	DeleteUser(userId int) error
	UpdateUser(id int, input track.People) error
	GetPeopleByFilters(filters track.Filter) ([]track.People, error)
}

type TrackTask interface {
	CreateTask(userId int, task track.Task) (int, error)
	DeleteTask(taskId int) error
	StartTask(taskId int) error
	EndTask(taskId int) error
	GetTasksByLaborCosts(userId int) ([]track.Task, error)
}

type Repository struct {
	TrackPeople
	TrackTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TrackPeople: NewTrackPeoplePostgres(db),
		TrackTask:   NewTrackTaskPostgres(db),
	}
}
