package endpoint

import (
	"encoding/json"
	"net/http"
	"s4m/internal/app/model"
)

type Service interface {
	RegisterEvent(event model.Event) error
	RegisterEvents(event []model.Event) error
	GetEvents(eventType, startTime, finishTime string) ([]model.Event, error)
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}

// CreateEvent is endpoint for register new event
func (e *Endpoint) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = e.s.RegisterEvent(event)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	w.Write([]byte("Success"))
}

// CreateEvents is endpoint for create new test events
func (e *Endpoint) CreateEvents(w http.ResponseWriter, r *http.Request) {
	var events []model.Event
	err := json.NewDecoder(r.Body).Decode(&events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = e.s.RegisterEvents(events)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	w.Write([]byte("Success"))
}

// GetEvents is endpoint for getn events by eventType, start event time and finish event time
func (e *Endpoint) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := e.s.GetEvents(r.URL.Query().Get("eventType"), r.URL.Query().Get("startTime"), r.URL.Query().Get("finishTime"))
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	body, err := json.Marshal(&events)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.Write(body)
}
