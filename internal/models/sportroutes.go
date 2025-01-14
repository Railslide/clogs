package models

import (
	"database/sql"
	"errors"
)

type Gym struct {
	ID       int
	Name     string
	Location string
	Country  string
}

// Using the prefix sport to separate it from web application routes
type SportRoute struct {
	ID        int
	Name      string
	Grade     string // this should likely be an enum value
	RouteType string // this should also be an enum
	Sent      bool
	Archived  bool
	Notes     string
	Gym       Gym
}

type SportRouteModel struct {
	DB *sql.DB
}

func (m *SportRouteModel) Insert(name, grade, routeType string, gymID int, sent, archived bool, notes string) (int, error) {
	stmt := `INSERT INTO routes (name, grade, route_type, sent, archived, notes, gym )
        VALUES(?, ?, ?, ?, ?, ?, ?)`
	result, err := m.DB.Exec(stmt, name, grade, routeType, sent, archived, notes, gymID)
	if err != nil {
		return 0, err
	}

	// Not all sql drivers support LastInsertId!
	// Check that the new driver does when changing it
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SportRouteModel) Get(id int) (SportRoute, error) {
	stmt := `SELECT
	     r.id, r.name, r.grade, r.route_type, r.sent, r.archived, r.notes,
	     g.id, g.name, g.location, g.country
	     FROM routes r
	     JOIN gyms g on r.gym = g.id
	     WHERE r.id = ?
	     `
	row := m.DB.QueryRow(stmt, id)

	var s SportRoute
	err := row.Scan(
		&s.ID, &s.Name, &s.Grade, &s.RouteType, &s.Sent, &s.Archived, &s.Notes,
		&s.Gym.ID, &s.Gym.Name, &s.Gym.Location, &s.Gym.Country,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return SportRoute{}, ErrNoRecord
		} else {
			return SportRoute{}, err
		}
	}

	return s, nil
}

func (m *SportRoute) Ongoing() ([]SportRoute, error) {
	return nil, nil
}
