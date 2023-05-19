package main

import "errors"

type Counter struct {
	counter int
	value   int
	pts     []int
}

func (w *Counter) Commit() error {
	w.value = w.counter
	return nil
}

func (w *Counter) Rollback() error {
	w.counter = w.value
	return nil
}

func (w *Counter) Save(name string) error {
	w.pts = append(w.pts, w.counter)
	return nil
}

func (w *Counter) RollbackSave() error {
	return nil
}

func (w *Counter) Add(n int) error {
	if n > 10 {
		return errors.New("n > 10")
	}
	w.counter += n
	return nil
}

func (w *Counter) Sub(n int) error {
	if n > 10 {
		return errors.New("n > 10")
	}
	w.counter -= n
	return nil
}

func (w Counter) Count() int {
	return w.counter
}
