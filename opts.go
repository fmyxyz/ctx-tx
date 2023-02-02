package tx

import (
	"context"
	"database/sql"
)

type TxOption func(opt *TxOptions)

type TxOptions struct {
	Isolation   sql.IsolationLevel
	ReadOnly    bool
	Propagation int8
	Name        string
}

func IsolationLevel(level sql.IsolationLevel) TxOption {
	return func(opt *TxOptions) {
		opt.Isolation = level
	}
}

func ReadOnly(readOnly bool) TxOption {
	return func(opt *TxOptions) {
		opt.ReadOnly = readOnly
	}
}

func Name(name string) TxOption {
	return func(opt *TxOptions) {
		opt.Name = name
	}
}

const (
	Required  = 1
	Supported = 2
	Nested    = 3
	New_      = 4
)

func PropagationRequired() TxOption {
	return func(opt *TxOptions) {
		opt.Propagation = Required
	}
}

func PropagationSupported() TxOption {
	return func(opt *TxOptions) {
		opt.Propagation = Supported
	}
}

func PropagationNested() TxOption {
	return func(opt *TxOptions) {
		opt.Propagation = Nested
	}
}

func PropagationNew() TxOption {
	return func(opt *TxOptions) {
		opt.Propagation = New_
	}
}

// tx beginner
type TxBeginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

// tx committer
type TxCommitter interface {
	Commit() error
	Rollback() error
}

// save pointer interface
type SavePointer interface {
	SavePoint(name string) error
	RollbackTo(name string) error
}

type Tx interface {
	TxCommitter
}

type DB interface {
	TxBeginner

	Name() string
}
