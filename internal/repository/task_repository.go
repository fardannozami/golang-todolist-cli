package repository

import (
	"fmt"
	"time"

	"github.com/fardannozami/golang-todolist-cli/internal/model"
)

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

func (r *InMemoryTaskRepository) GetAllTasks() ([]model.Task, error) {
	return r.tasks, nil
}

func (r *InMemoryTaskRepository) DeleteTask(id int) error {
	for i, task := range r.tasks {
		if task.Id == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}

func (r *InMemoryTaskRepository) MarkTaskAsCompleted(id int) error {
	for i, task := range r.tasks {
		if task.Id == id {
			completed := time.Now()
			r.tasks[i].CompletedAt = &completed
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}
