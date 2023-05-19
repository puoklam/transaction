package atomic

import "github.com/google/uuid"

type Commit struct {
	id uuid.UUID
}

func (c Commit) ID() uuid.UUID {
	return c.id
}

func NewCommit() *Commit {
	return &Commit{
		id: uuid.New(),
	}
}
