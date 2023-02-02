package std

import (
	"context"
	"database/sql"

	"github.com/fmyxyz/tx"
)

var _ StdSQL = (*stdDB)(nil)

type stdDB struct {
	*sql.DB

	Options
}

type stdTx struct {
	*sql.Tx

	Options
}

func (g *stdTx) SavePoint(name string) error {
	_, err := g.Tx.Exec("SAVEPOINT " + name)
	return err
}

func (g *stdTx) RollbackTo(name string) error {
	_, err := g.Tx.Exec("ROLLBACK TO SAVEPOINT " + name)
	return err
}

func (g *stdTx) Commit() error {
	return g.Tx.Commit()
}

func (g *stdTx) Rollback() error {
	return g.Tx.Rollback()
}

func (g *stdTx) Name() string {
	return "std-" + g.instance
}

func (g *stdDB) Name() string {
	return "std-" + g.instance
}

func (g *stdDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (tx.Tx, error) {
	tx0, err := g.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return warp(tx0), err
}

func warp(tx0 *sql.Tx) *stdTx {
	return &stdTx{Tx: tx0}
}

const defaultInstance = "default"

func Register(db *sql.DB, opts ...Option) {
	o := &Options{instance: defaultInstance}
	stdDB := &stdDB{DB: db}
	for _, opt := range opts {
		opt(o)
	}
	stdDB.instance = o.instance

	tx.Register(stdDB, tx.RegisterDefaultDB(stdDB.instance == defaultInstance))
}

type Options struct {
	instance string
}

type Option func(db *Options)

func Instance(instance string) Option {
	return func(db *Options) {
		db.instance = instance
	}
}

func FromContext(ctx context.Context, opts ...Option) StdSQL {
	o := &Options{instance: defaultInstance}
	std := &stdDB{}
	for _, opt := range opts {
		opt(o)
	}
	std.instance = o.instance

	name := std.Name()
	txManager := tx.GetTxManager(name)
	if txManager == nil {
		panic(name + " not register in txManagers")
	}
	tx0 := txManager.TxFromContext(ctx)
	if tx0 != nil {
		return tx0.(*stdTx)
	}
	db := txManager.DBFromContext(ctx)
	if db != nil {
		return db.(*stdDB)
	}
	return nil
}
