package events

import (
	"database/sql"
	"errors"
)

type Event struct {
	EventId    int64
	AppToken   string
	Type       string
	ExternalId string
	HappenedAt int64
	Properties map[string]interface{}
}

func (e *Event) CheckFields() error {
	if e.AppToken == "" {
		return errors.New("Empty AppToken field in event")
	}

	if e.Type == "" {
		return errors.New("Empty Type field in event")
	}
	return nil
}

func (e *Event) Track() error {
	if err := e.CheckFields(); err != nil {
		return err
	}
}

func (e *Event) Delete() error {
}
