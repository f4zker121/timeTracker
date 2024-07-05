package service

import (
	track "timeTracker"
	"timeTracker/pkg/repository"
)

type TrackTaskService struct {
	repo repository.TrackTask
}

func NewTrackTaskService(repo repository.TrackTask) *TrackTaskService {
	return &TrackTaskService{repo: repo}
}

func (t *TrackTaskService) CreateTask(userId int, task track.Task) (int, error) {
	return t.repo.CreateTask(userId, task)
}

func (t *TrackTaskService) DeleteTask(taskId int) error {
	return t.repo.DeleteTask(taskId)
}

func (t *TrackTaskService) StartTask(taskId int) error {
	return t.repo.StartTask(taskId)
}

func (t *TrackTaskService) EndTask(taskId int) error {
	return t.repo.EndTask(taskId)
}

func (t *TrackTaskService) GetTasksByLaborCosts(userId int) ([]track.Task, error) {
	return t.repo.GetTasksByLaborCosts(userId)
}
