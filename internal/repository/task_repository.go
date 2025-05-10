package repository

import "github.com/fardannozami/golang-todolist-cli/internal/model"

type TaskRepository interface {
	AddTask(task model.Task) error
	GetAllTasks() ([]model.Task, error)
	DeleteTask(id int) error
	MarkTaskAsCompleted(id int) error
}

type InMemoryTaskRepository struct {
	tasks []model.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make([]model.Task, 0),
	}
}

func (r *InMemoryTaskRepository) AddTask(task model.Task) error {
	r.tasks = append(r.tasks, task)
	return nil
}
