package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/Diogo364/ImersaoFullCycle-14/go/internal/routes/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) Create(route entity.Route) error {
	sqlString := "INSERT INTO routes (name, source, destination) VALUES (?, ?, ?)"

	routeSourceJSON, err := json.Marshal(route.Source)
	if err != nil {
		return err
	}

	routeDestinationJSON, err := json.Marshal(route.Destination)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sqlString, route.Name, routeSourceJSON, routeDestinationJSON)
	if err != nil {
		return err
	}
	return nil
}

func (r *RouteRepositoryMysql) FindAll() ([]entity.Route, error) {
	sqlString := "SELECT id, name, source, destination FROM routes"
	rows, err := r.db.Query(sqlString)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var routeList []entity.Route

	for rows.Next() {
		var route entity.Route
		err := rows.Scan(
			&route.ID,
			&route.Name,
			&route.Source,
			&route.Destination,
		)
		if err != nil {
			return nil, err
		}
		routeList = append(routeList, route)
	}
	return routeList, nil
}
