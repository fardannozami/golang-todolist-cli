package repository

import (
	"testing"
	"time"

	"github.com/fardannozami/golang-todolist-cli/internal/model"
	"github.com/stretchr/testify/assert"
)

var tasks = []model.Task{
	{
		Id:          1,
		Description: "Test Task 1",
		CreatedAt:   time.Now(),
	},
	{
		Id:          2,
		Description: "Test Task 2",
		CreatedAt:   time.Now(),
	},
}

func TestInMemoryTaskRepository_AddTask(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	for _, task := range tasks {
		err := repo.AddTask(task)
		if err != nil {
			t.Errorf("AddTask() error = %v", err)
		}
	}

	assert.Len(t, repo.tasks, 2)
	assert.Equal(t, tasks[0], repo.tasks[0])
}

func TestInMemoryTaskRepository_GetTasks(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	for _, task := range tasks {
		err := repo.AddTask(task)
		if err != nil {
			t.Errorf("AddTask() error = %v", err)
		}
	}

	tasksAll, err := repo.GetAllTasks()
	if err != nil {
		t.Errorf("GetTasks() error = %v", err)
	}

	assert.NoError(t, err)
	assert.Len(t, tasksAll, 2)
}

func TestInMemoryTaskRepository_DeleteTask(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	for _, task := range tasks {
		err := repo.AddTask(task)
		if err != nil {
			t.Errorf("AddTask() error = %v", err)
		}
	}

	err := repo.DeleteTask(1)
	if err != nil {
		t.Errorf("DeleteTask() error = %v", err)
	}
	assert.NoError(t, err)
	assert.Len(t, repo.tasks, 1)
	assert.Equal(t, tasks[1], repo.tasks[0])
}

func TestInMemoryTaskRepository_DeleteTaskNotFound(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	for _, task := range tasks {
		err := repo.AddTask(task)
		if err != nil {
			t.Errorf("AddTask() error = %v", err)
		}
	}

	err := repo.DeleteTask(900)
	assert.Error(t, err)
	assert.Len(t, repo.tasks, 2)
}

func TestInMemoryTaskRepository_MarkTaskAsCompleted(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	for _, task := range tasks {
		err := repo.AddTask(task)
		if err != nil {
			t.Errorf("AddTask() error = %v", err)
		}
	}

	err := repo.MarkTaskAsCompleted(1)
	if err != nil {
		t.Errorf("MarkTaskAsCompleted() error = %v", err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, repo.tasks[0].CompletedAt)
	assert.Nil(t, repo.tasks[1].CompletedAt)
}

func TestInMemoryTaskRepository_MarkTaskAsCompletedNotFound(t *testing.T) {
	repo := NewInMemoryTaskRepository()
	for _, task := range tasks {
		err := repo.AddTask(task)
		if err != nil {
			t.Errorf("AddTask() error = %v", err)
		}
	}

	err := repo.MarkTaskAsCompleted(100)
	assert.Error(t, err)
	assert.Nil(t, repo.tasks[0].CompletedAt)
	assert.Nil(t, repo.tasks[1].CompletedAt)
}
