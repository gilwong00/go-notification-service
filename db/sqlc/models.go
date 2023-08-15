// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type State string

const (
	StatePending State = "pending"
	StateSuccess State = "success"
	StateFailed  State = "failed"
)

func (e *State) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = State(s)
	case string:
		*e = State(s)
	default:
		return fmt.Errorf("unsupported scan type for State: %T", src)
	}
	return nil
}

type NullState struct {
	State State
	Valid bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullState) Scan(value interface{}) error {
	if value == nil {
		ns.State, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.State.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.State, nil
}

type Follower struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	FollowerID uuid.UUID
	CreatedAt  time.Time
	DeletedAt  time.Time
}

type NotificationQueue struct {
	ID         uuid.UUID
	Message    string
	FollowerID uuid.UUID
	CreatedAt  time.Time
	StateID    string
	Attempts   sql.NullInt32
	UpdatedAt  time.Time
}

type NotificationState struct {
	ID            uuid.UUID
	State         State
	Message       string
	RequestedAtAt time.Time
	CompletedAt   sql.NullTime
}

type User struct {
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	DeletedAt time.Time
}
