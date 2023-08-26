package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Diogo364/ImersaoFullCycle-14/go/internal/routes/entity"
	"github.com/Diogo364/ImersaoFullCycle-14/go/internal/routes/infra/repository"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	routeRepository repository.RouteRepositoryMysql
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/routes?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	routeRepositoryMySql := repository.NewRouteRepositoryMysql(db)

	app := App{
		routeRepository: *routeRepositoryMySql,
	}

	r := chi.NewRouter()
	r.Post("/api/routes", app.createRoute)
	r.Get("/api/routes", app.listRoutes)

	http.ListenAndServe(":8080", r)
}

func (app *App) createRoute(w http.ResponseWriter, r *http.Request) {
	var route entity.Route
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = app.routeRepository.Create(route)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (app *App) listRoutes(w http.ResponseWriter, r *http.Request) {
	routes, err := app.routeRepository.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routes)
}
