package atomic

import "github.com/google/uuid"

type SavePoint struct {
	id   uuid.UUID
	Name string
}

func (p SavePoint) ID() uuid.UUID {
	return p.id
}

func NewSavePoint(name string) *SavePoint {
	return &SavePoint{
		id:   uuid.New(),
		Name: name,
	}
}
