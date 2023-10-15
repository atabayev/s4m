package service

import (
	"fmt"
	"log"
	"s4m/internal/app/model"
	"time"
)

type Database interface {
	InsertOne(event model.Event) error
	InsertMore(events []model.Event) error
	SelectEvent(eventType string, startDate, finishDate time.Time) ([]model.Event, error)
}

type Service struct {
	db Database
}

func New(db Database) *Service {
	return &Service{db: db}
}

func (s *Service) RegisterEvent(event model.Event) error {
	err := s.db.InsertOne(event)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s *Service) RegisterEvents(events []model.Event) error {
	err := s.db.InsertMore(events)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (s *Service) GetEvents(eventType, startTime, finishTime string) ([]model.Event, error) {
	layout := "02.01.2006 15:04" // RFC3339 format

	st, err := time.Parse(layout, startTime)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	ft, err := time.Parse(layout, finishTime)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	events, err := s.db.SelectEvent(eventType, st, ft)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return events, err
}
