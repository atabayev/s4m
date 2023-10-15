package app

import (
	"net/http"
	"s4m/internal/app/database"
	"s4m/internal/app/endpoint"
	"s4m/internal/app/service"

	"github.com/gorilla/mux"
)

type App struct {
	db  *database.Database
	svc *service.Service
	r   *mux.Router
	e   *endpoint.Endpoint
}

func New() (*App, error) {
	a := &App{}
	var err error
	a.db, err = database.New("clickhouse", "http://127.0.0.1:8123/test")
	if err != nil {
		return nil, err
	}
	a.svc = service.New(a.db)
	a.e = endpoint.New(a.svc)
	a.r = mux.NewRouter()
	a.r.HandleFunc("/api/event", a.e.CreateEvent).Methods("POST")
	a.r.HandleFunc("/api/set/events", a.e.CreateEvents).Methods("POST")
	a.r.HandleFunc("/api/get/events", a.e.GetEvents).Methods("GET")
	return a, nil
}

func (a *App) Run() error {
	return http.ListenAndServe(":8080", a.r)
}
