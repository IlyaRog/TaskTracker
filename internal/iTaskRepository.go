package internal

type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Save(task *Task) error
	Update(id string, key, value string) error
	Delete(id string) error
}
