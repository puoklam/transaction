package main

import (
	"fmt"

	"github.com/puoklam/transaction"
	"github.com/puoklam/transaction/atomic"
)

func main() {
	c := &Counter{}
	tx := transaction.New(c)
	fn := func(tx *transaction.Tx, w atomic.Atomic) error {
		c := w.(*Counter)
		if c.Add(5) != nil {
			tx.Rollback()
			return nil
		}
		tx.Commit()
		return nil
	}
	tx.Begin(fn)
	fmt.Println(c.Count())
}
