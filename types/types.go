package types

import (
	"time"
	"github.com/google/uuid"

)

type TaskItem struct {
	ID uuid.UUID
	Task string
	Done bool
	Category string
	CreatedOn time.Time
	CompletedOn *time.Time
}

type Tasks []TaskItem
