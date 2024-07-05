package service

import (
	track "timeTracker"
	"timeTracker/pkg/repository"
)

type TrackPeopleService struct {
	repo repository.TrackPeople
}

func NewTrackPeopleService(repo repository.TrackPeople) *TrackPeopleService {
	return &TrackPeopleService{repo: repo}
}

func (t *TrackPeopleService) CreateUser(people track.People) (int, error) {
	return t.repo.CreateUser(people)
}

func (t *TrackPeopleService) DeleteUser(userId int) error {
	return t.repo.DeleteUser(userId)
}

func (t *TrackPeopleService) UpdateUser(id int, input track.People) error {
	return t.repo.UpdateUser(id, input)
}

func (t *TrackPeopleService) GetPeopleByFilters(filters track.Filter) ([]track.People, error) {
	return t.repo.GetPeopleByFilters(filters)
}
