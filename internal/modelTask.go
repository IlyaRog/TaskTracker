package internal

import (
	"fmt"
	"time"

	uuid "github.com/google/uuid"
)

type Task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updateAt"`
}

func newTask(title, description string) (*Task, error) {
	return &Task{
		Id:          uuid.New().String(),
		Title:       title,
		Description: description,
		Status:      StatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}, nil
}

func (task Task) PrintTask() {
	fmt.Printf("\nDate of creation: %s\nTitle: %s\nStatus: %s\nDescription: %s\n", task.CreatedAt.Format("02.01.2006 15:04"), task.Title, task.Status, task.Description)
}
