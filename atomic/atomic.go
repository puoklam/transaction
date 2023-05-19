package atomic

type Atomic interface {
	Commit() error
	Save(name string) error
	Rollback() error
	RollbackSave() error
}
