package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type TaskCLI struct {
	Repo TaskRepository
}

func CreateCLI(newRepo TaskRepository) *TaskCLI { return &TaskCLI{Repo: newRepo} }

func (cli TaskCLI) CreateTask(title, description string) (*Task, error) {
	newTask, err := newTask(title, description)
	if err != nil {
		log.Fatal(err)
	}

	if err = cli.Repo.Save(newTask); err != nil {
		log.Fatal(err)
	}
	return newTask, err
}

func (cli TaskCLI) UpdateTask(id string, key, value string) error {

	err := cli.Repo.Update(id, key, value)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (cli TaskCLI) DeleteTask(id string) error {
	return cli.Repo.Delete(id)
}

func (cli TaskCLI) ListTask(filter Status) ([]Task, error) {
	Tasks, err := cli.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	var temp []Task

	for _, task := range Tasks {
		if task.Status == filter {
			temp = append(temp, task)
		}
	}

	return temp, nil
}

func (cli TaskCLI) Menu() {
	var input string
	cli.clearScreen()
	fmt.Printf("\t\t Добро пожаловать в Task Tracker!")
	for {
		fmt.Printf("\nВыберите нужный пункт меню:\n 1. Создать новую задачу\n 2. Обновить задачу\n 3. Удалить задачу\n 4. Вывести список задач по статусу\n exit - чтобы завершить работу*\n")

		fmt.Scan(&input)

		switch input {
		case "1":
			var title, description string
			input := bufio.NewReader(os.Stdin)

			fmt.Printf("\n\nСоздание задачи...\n")

			input.ReadString('\n') //"очистка" stdin

			fmt.Println("Введите название задачи:")
			title, _ = input.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Println("Введите описание задачи:")
			description, _ = input.ReadString('\n')
			description = strings.TrimSpace(description)

			if _, err := cli.CreateTask(title, description); err != nil {
				fmt.Println("Задача успешно создана!")
			}

		case "2":
			var (
				id,
				num,
				key,
				value string
			)
			fmt.Printf("\n\nВедите id задачи:\n")
			fmt.Scan(&id)
			fmt.Printf("Выбирите нужное поле для изменения:\n1 - Название\n2 - Статус\n3 - Описание\n")
			fmt.Scan(&num)
			switch num {
			case "1":
				key = "title"
			case "2":
				key = "status"
			case "3":
				key = "description"
			default:
				fmt.Printf("Данный вариант отсутствует, возвращение в меню")
				return
			}
			fmt.Printf("\nВедите значение:\n")
			input := bufio.NewReader(os.Stdin)
			input.ReadString('\n')

			value, _ = input.ReadString('\n')
			value = strings.TrimSpace(value)

			if err := cli.UpdateTask(id, key, value); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Задача успешно обновлена!")

		case "3":
			var id string
			fmt.Printf("\n\nВедите id задачи:\n")
			fmt.Scan(&id)
			if err := cli.DeleteTask(id); err != nil {
				log.Fatal(err)
			}
		case "4":
			var num string
			status := StatusToDo
			fmt.Printf("Выбирите статус задач:\n1 - %s\n2 - %s\n3 - %s\n", StatusToDo, StatusInProgress, StatusComplite)
			fmt.Scan(&num)

			switch num {
			case "1":
				status = StatusToDo
			case "2":
				status = StatusInProgress
			case "3":
				status = StatusComplite
			default:
				fmt.Printf("Данный вариант отсутствует, выбран статус по умолчанию - %s", StatusToDo)
			}

			listTask, err := cli.ListTask(status)

			if err != nil {
				log.Fatal(err)
			}

			for _, task := range listTask {
				task.PrintTask()
			}

		case "exit":
			os.Exit(0)
		default:
			cli.clearScreen()
			fmt.Println("Такой вариант отсутствует...")
		}
	}
}

func (cli TaskCLI) clearScreen() {
	fmt.Print("\x1b[40m")
	fmt.Print("\x1b[2J")
}
