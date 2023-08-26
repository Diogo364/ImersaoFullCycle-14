package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONString map[string]interface{}

func (j JSONString) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONString) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Failed to access bytes")
	}
	var v map[string]interface{}
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		return err
	}
	*j = v
	return nil
}

type Route struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Source      JSONString `json:"source"`
	Destination JSONString `json:"destination"`
}

type RouteRepository interface {
	Create(route *Route) error
	FindAll() ([]Route, error)
}

func NewRoute(id string, name string, source JSONString, destination JSONString) *Route {
	return &Route{
		ID:          id,
		Name:        name,
		Source:      source,
		Destination: destination,
	}
}
