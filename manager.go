package tx

import (
	"context"
	"sync/atomic"
)

type txManager struct {
	db DB
}

func newTxManager(db DB) *txManager {
	t := &txManager{db: db}
	return t
}

func existTx(ctx context.Context) bool {
	if ctx == nil {
		ctx = context.Background()
	}
	_, ok := ctx.Value(txCtxKey).(*TxNode)
	return ok
}

func (t *txManager) NewTxNode(ctx context.Context, fc TxFunc, txOpts *TxOptions) *TxNode {
	var node *TxNode

	if ctx == nil {
		ctx = context.Background()
	}
	pNode, ok := ctx.Value(txCtxKey).(*TxNode)
	if ok {
		node = &TxNode{
			parent: pNode,
			tx:     pNode.tx,
			db:     pNode.db,
			fc:     fc,
			txOpts: txOpts,
		}
	} else {
		node = &TxNode{
			db:     t.db,
			fc:     fc,
			txOpts: txOpts,
		}
	}

	node.withCtx(ctx)
	return node
}

func (t *txManager) WithTx(ctx context.Context, fc TxFunc, opts ...TxOption) (err error) {
	txOpts := &TxOptions{Propagation: Required}
	for _, opt := range opts {
		opt(txOpts)
	}
	txNode := t.NewTxNode(ctx, fc, txOpts)
	switch txOpts.Propagation {
	case Required:
		if !existTx(ctx) { // 不存在就创建
			return txNode.newTx()
		} else {
			return txNode.useTx()
		}
	case Supported:
		if existTx(ctx) {
			return txNode.useTx()
		} else {
			return txNode.nonTx(ctx)
		}
	case Nested:
		if !existTx(ctx) { // 不存在就创建
			return txNode.newTx()
		} else { // 存在就嵌套
			return txNode.subTx()
		}
	case New_:
		return txNode.newTx()
	}

	return
}

func WithTx(ctx context.Context, fc TxFunc, opts ...TxOption) (err error) {
	name := defaultKey.Load().(string)
	options := &TxOptions{Name: name}
	for _, opt := range opts {
		opt(options)
	}
	return GetTxManager(options.Name).WithTx(ctx, fc, opts...)
}

func (t *txManager) TxFromContext(ctx context.Context) (tx Tx) {
	if ctx == nil {
		return nil
	}
	node, ok := ctx.Value(txCtxKey).(*TxNode)
	if !ok {
		return nil
	}
	return node.tx
}

func (t *txManager) DBFromContext(ctx context.Context) (db DB) {
	if ctx == nil {
		return t.db
	}
	node, ok := ctx.Value(txCtxKey).(*TxNode)
	if !ok {
		return t.db
	}
	return node.db
}

type TxFunc func(ctx context.Context) error

var txManagers = map[string]*txManager{}

func GetTxManager(name string) *txManager {
	return txManagers[name]
}

var defaultKey atomic.Value

func init() {
	defaultKey.Store("")
}

func Register(db DB, opts ...RegisterDBOption) {
	o := &RegisterDBOptions{}
	for _, opt := range opts {
		opt(o)
	}
	name := db.Name()
	if o.Default {
		defaultKey.Store(name)
	}
	txManagers[name] = newTxManager(db)
}

type RegisterDBOptions struct {
	Default bool
}

type RegisterDBOption func(opt *RegisterDBOptions)

func RegisterDefaultDB(def bool) RegisterDBOption {
	return func(opt *RegisterDBOptions) {
		opt.Default = def
	}
}
