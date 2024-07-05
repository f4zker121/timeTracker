package repository

import (
	"fmt"
	"strconv"
	"strings"
	track "timeTracker"

	"github.com/jmoiron/sqlx"
)

type TrackPeoplePostgres struct {
	db *sqlx.DB
}

func NewTrackPeoplePostgres(db *sqlx.DB) *TrackPeoplePostgres {
	return &TrackPeoplePostgres{db: db}
}

func (t *TrackPeoplePostgres) CreateUser(people track.People) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (passport_serie, passport_number, surname, name, patronymic, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", peopleTable)
	row := t.db.QueryRow(query, people.PassportSerie, people.PassportNumber, people.Surname, people.Name, people.Patronymic, people.Address)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TrackPeoplePostgres) DeleteUser(userId int) error {
	tx, err := t.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	deleteQueryUser := fmt.Sprintf("DELETE FROM %s WHERE id = $1", peopleTable)
	_, err = t.db.Exec(deleteQueryUser, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteQueryTask := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", taskTable)
	_, err = t.db.Exec(deleteQueryTask, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

func (t *TrackPeoplePostgres) UpdateUser(id int, input track.People) error {
	query := fmt.Sprintf("UPDATE %s SET passport_serie = $1, passport_number = $2, surname = $3, name = $4, patronymic = $5, address = $6 WHERE id = $7", peopleTable)

	_, err := t.db.Exec(query, input.PassportSerie, input.PassportNumber, input.Surname, input.Name, input.Patronymic, input.Address, id)
	return err
}

func (t *TrackPeoplePostgres) GetPeopleByFilters(filters track.Filter) ([]track.People, error) {
	var people []track.People
	var args []interface{}
	query := "SELECT * FROM people WHERE"
	conditions := []string{}

	if filters.Id != 0 {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(args)+1))
		args = append(args, filters.Id)
	}
	if filters.PassportSerie != 0 {
		conditions = append(conditions, "passport_serie = $"+strconv.Itoa(len(args)+1))
		args = append(args, filters.PassportSerie)
	}
	if filters.PassportNumber != 0 {
		conditions = append(conditions, "passport_number = $"+strconv.Itoa(len(args)+1))
		args = append(args, filters.PassportNumber)
	}
	if filters.Surname != "" {
		conditions = append(conditions, "surname = $"+strconv.Itoa(len(args)+1))
		args = append(args, filters.Surname)
	}
	if filters.Name != "" {
		conditions = append(conditions, "name = $"+strconv.Itoa(len(args)+1))
		args = append(args, filters.Name)
	}
	if filters.Patronymic != "" {
		conditions = append(conditions, "patronymic = $"+strconv.Itoa(len(args)+1))
		args = append(args, filters.Patronymic)
	}
	if filters.Address != "" {
		conditions = append(conditions, "address = $"+strconv.Itoa(len(args)+1))
		args = append(args, filters.Address)
	}

	query += " " + strings.Join(conditions, " AND ")

	if filters.Pelingation > 0 {
		query += " LIMIT " + strconv.Itoa(filters.Pelingation)
	}

	err := t.db.Select(&people, query, args...)
	if err != nil {
		return nil, err
	}

	return people, nil
}
