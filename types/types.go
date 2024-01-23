package types

import (
    "time"
)

type item struct {
	Task string
	Done bool
	Category string
	CreatedOn time.Time
	CompletedOn time.Time
}

type List []item