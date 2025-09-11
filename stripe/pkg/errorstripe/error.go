package errorstripe

import "fmt"

// ErrNotFound error type for objects not found
type ErrNotFound struct {
	ID   string
	Type string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("object %v of type: %v, not found", e.ID, e.Type)
}

type ErrIsEmpty struct {
	ID   string
	Type string
}

func (e *ErrIsEmpty) Error() string {
	return fmt.Sprintf("object %v of type: %v, is empty", e.ID, e.Type)
}

type ErrHttpServer struct {
	Type string
	Err  error
}

func (e *ErrHttpServer) Error() string {
	return fmt.Sprintf("Error from http %v: %v", e.Type, e.Err)
}
