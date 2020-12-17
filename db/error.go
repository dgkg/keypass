package db

import "fmt"

type ErrKind int

const (
	ErrUnknown ErrKind = iota
	ErrNotFound
	ErrInternal
)

type ErrDB struct {
	Params string
	Err    error
	Kind   ErrKind
}

func (e ErrDB) Error() string {
	return fmt.Sprintf("params:%v err:%v kind:%v", e.Params, e.Err, e.Kind)
}

func NewErrNotFound(params string, err error) error {
	return &ErrDB{
		Params: params,
		Err:    err,
		Kind:   ErrNotFound,
	}
}
