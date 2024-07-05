package service

import (
	track "timeTracker"
	"timeTracker/pkg/repository"
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

type Service struct {
	TrackPeople
	TrackTask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TrackPeople: NewTrackPeopleService(repos.TrackPeople),
		TrackTask:   NewTrackTaskService(repos.TrackTask),
	}
}
