package transaction

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/puoklam/transaction/atomic"
)

var ErrSaveNotFound = errors.New("save not found")

type Transaction struct {
	mu sync.Mutex
	tx *Tx
}

type Tx struct {
	w   atomic.Atomic
	c   *atomic.Commit
	pts []*atomic.SavePoint
}

type TxFunc func(tx *Tx, w atomic.Atomic) error

func (t *Transaction) Begin(fn TxFunc) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	return fn(t.tx, t.tx.W())
}

func (t *Tx) W() atomic.Atomic {
	return t.w
}

func (t *Tx) Commit() error {
	err := t.w.Commit()
	if err != nil {
		return err
	}
	t.c = atomic.NewCommit()
	t.pts = nil
	return nil
}

func (t *Tx) Rollback() error {
	return t.w.Rollback()
}

func (t *Tx) Save(name string) error {
	err := t.w.Save(name)
	if err != nil {
		return err
	}
	sp := atomic.NewSavePoint(name)
	t.pts = append(t.pts, sp)
	return nil
}

func (t *Tx) RollbackSave(id uuid.UUID) error {
	i := t.getSaveIndex(id)
	if i == -1 {
		return ErrSaveNotFound
	}
	err := t.w.RollbackSave()
	if err != nil {
		return err
	}
	t.pts = t.pts[0 : i+1]
	return nil
}

// Recover restores the last commit
func (t *Tx) Recover() error {
	return t.restore(t.c.ID())
}

// resotre restores the commit with id = cid
func (t *Tx) restore(cid uuid.UUID) error {
	return nil
}

func (t *Tx) getSaveIndex(id uuid.UUID) int {
	for i, sp := range t.pts {
		if sp.ID().String() == id.String() {
			return i
		}
	}
	return -1
}

func New[T atomic.Atomic](w T) *Transaction {
	tx := &Tx{
		w: w,
	}
	return &Transaction{
		tx: tx,
	}
}

// Atomic All changes must be successfully or not at all
// Consistent Data must be in a consistent state before and after the transaction
// Isolated No other process can change the while the transaction is running
// Durable The changes made by a transaction must persist
