package internal

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type JsonTaskRepository struct {
	FilePath string
}

func (repo *JsonTaskRepository) GetAll() (list []Task, err error) {
	files, err := os.ReadDir(repo.FilePath)
	var temp Task
	for _, entity := range files {
		fileData, err := os.ReadFile(repo.FilePath + entity.Name())
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(fileData, &temp); err != nil {
			return nil, err
		}
		list = append(list, temp)
	}

	return list, err
}

func (repo *JsonTaskRepository) GetByID(id string) (task *Task, err error) {
	fileName := repo.FilePath + id
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileData, task)

	return task, err
}

func (repo *JsonTaskRepository) Save(task *Task) error {
	jsonData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	filename := repo.FilePath + string(task.Id)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}

	return nil
}

func (repo *JsonTaskRepository) Update(id string, key, value string) error {
	var task Task

	fileName := repo.FilePath + id
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, &task)
	if err != nil {
		return err
	}

	switch {
	case key == "title":
		task.Title = value
	case key == "description":
		task.Description = value
	case key == "status":
		task.Status = Status(value)
	default:
		err = errors.New("Update Task: key is not found")
		return err
	}

	task.UpdatedAt = time.Now()

	result, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, result, 0644)
}

func (repo *JsonTaskRepository) Delete(id string) error {
	fileName := repo.FilePath + id
	err := os.Remove(fileName)
	return err
}
