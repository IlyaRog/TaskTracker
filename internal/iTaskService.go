package internal

type TaskService interface {
	CreateTask(title, description string) (*Task, error)
	UpdateTask(id string, status Status) error
	DeleteTask(id string) error
	ListTask(filter Status) ([]Task, error)
}
