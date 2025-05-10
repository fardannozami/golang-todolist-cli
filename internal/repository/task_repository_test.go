package repository

import (
	"testing"
	"time"

	"github.com/fardannozami/golang-todolist-cli/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryTaskRepository_AddTask(t *testing.T) {
	task := model.Task{
		Id:          1,
		Description: "Test Task",
		CreatedAt:   time.Now(),
	}

	repo := NewInMemoryTaskRepository()
	err := repo.AddTask(task)
	if err != nil {
		t.Errorf("AddTask() error = %v", err)
	}

	assert.Len(t, repo.tasks, 1)
}
