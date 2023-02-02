package tx

import (
	"context"
	"database/sql"
	"fmt"
)

/*
1. 新建时注入db/tx

2. 开始事务

	获取已存在事务
	判定事务传播行为
		REQUIRED 	没有创建，有就加入
		NESTED	 	嵌套子事务 save point / rollback to
		NEW      	总是新建事务
		SUPPORTED   有就加入，没有不理会

3. 事务结束错误回滚无错误提交
*/
type TxNode struct {
	ctx context.Context

	parent *TxNode // 上级事务节点

	tx Tx
	db DB

	subTxPoint string

	committed bool
	rolled    bool

	fc     TxFunc
	txOpts *TxOptions
}

type txContextKey string

var txCtxKey = txContextKey("txContextKey")

func (t *TxNode) withCtx(ctx context.Context) {
	t.ctx = context.WithValue(ctx, txCtxKey, t)
}

func (t *TxNode) useTx() (err error) {
	panicked := true
	defer func() {
		if panicked || err != nil {
			t.rollback()
		}
	}()
	err = t.fc(t.ctx)
	panicked = false
	return err
}

func (t *TxNode) nonTx(ctx context.Context) error {
	return t.fc(ctx)
}

func (t *TxNode) newTx() (err error) {
	panicked := true
	beginTx, err := t.db.BeginTx(t.ctx, &sql.TxOptions{
		Isolation: t.txOpts.Isolation,
		ReadOnly:  t.txOpts.ReadOnly,
	})
	if err != nil {
		return err
	}
	t.tx = beginTx
	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if panicked || err != nil {
			t.rollback()
		}
	}()
	if err = t.fc(t.ctx); err == nil {
		panicked = false
		return t.commit()
	}
	return err
}

func (t *TxNode) subTx() (err error) {
	if committer, ok := t.tx.(SavePointer); ok && committer != nil {
		panicked := true
		// Nested transaction
		subTxPoint := fmt.Sprintf("sp%p", t.fc)
		err = committer.SavePoint(subTxPoint)
		if err != nil {
			return err
		}
		// 嵌套事务
		// 若嵌套事务中使用了Required或Supported事务，这些嵌套内的事务回滚时需回滚到SavePoint:subTxPoint
		t.subTxPoint = subTxPoint

		defer func() {
			// Make sure to rollback when panic, Block error or Commit error
			if panicked || err != nil {
				committer.RollbackTo(subTxPoint)
			}
		}()
		err = t.fc(t.ctx)
		panicked = false
		return err
	} else {
		return t.newTx()
	}
}

func (t *TxNode) setRolled() {
	t.rolled = true

	if t.txOpts.Propagation != Nested && // 嵌套事务
		t.txOpts.Propagation != New_ && // 新事务

		t.parent != nil {
		t.parent.setRolled()
	}
}

func (t *TxNode) rollback() (err error) {
	if t.rolled {
		return err
	}

	t.setRolled()

	if t.subTxPoint == "" {
		return t.tx.Rollback()
	}
	if committer, ok := t.tx.(SavePointer); ok && committer != nil {
		return committer.RollbackTo(t.subTxPoint)
	} else {
		return t.tx.Rollback()
	}
}

func (t *TxNode) commit() (err error) {
	if t.rolled {
		return err
	}
	t.committed = true
	return t.tx.Commit()
}
