//go:generate go run cmd/gen/main.go -name=String -type=string -output=string.go
package queue

import (
	"errors"
)

var (
	ErrEmptyQueue  = errors.New("queue: the queue is empty and the requested operation could not be performed")
	ErrInvalidType = errors.New("queue: invalid type encountered - this indicates a bug.")
)
