package client

import (
	"time"

	"github.com/cppforlife/turbulence/incident"
	"github.com/cppforlife/turbulence/incident/reporter"
	"github.com/cppforlife/turbulence/tasks"
)

type Turbulence interface {
	CreateIncident(incident.Request) (Incident, error)
}

type Incident interface {
	ID() string
	Wait() error // todo add timeout?

	// EventsOfType returns list events that match particular options type
	// Example: incident.EventsOfType(tasks.KillOptions{})
	EventsOfType(tasks.Options) []reporter.EventResponse
	HasEventErrors() bool

	// ExecutionStartedAt is expected to always return time,
	// unlike ExecutionCompletedAt which may return nil
	// when execution is not yet finished
	ExecutionStartedAt() time.Time
	ExecutionCompletedAt() *time.Time
}
