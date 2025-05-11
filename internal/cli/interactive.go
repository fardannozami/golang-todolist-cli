package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fardannozami/golang-todolist-cli/internal/model"
	"github.com/fardannozami/golang-todolist-cli/internal/repository"
)

type CLI struct {
	repo         repository.TaskRepository
	PromptRunner func(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error
	ExitFunc     func(code int)
}

func NewCLI(repo repository.TaskRepository) *CLI {
	return &CLI{
		repo:         repo,
		PromptRunner: survey.AskOne,
		ExitFunc:     os.Exit,
	}
}

func (c *CLI) Run() {
	for {
		if !c.RunOnce() {
			break
		}
	}
}

func (c *CLI) RunOnce() bool {
	choice := ""
	prompt := &survey.Select{
		Message: "Apa yang ingin kamu lakukan?",
		Options: []string{"➕ Add Task", "📋 List Tasks", "✅ Complete Task", "🗑️ Delete Task", "🚪 Exit"},
	}
	c.PromptRunner(prompt, &choice)

	switch choice {
	case "➕ Add Task":
		c.addTask()
	case "📋 List Tasks":
		c.listTasks()
	case "✅ Complete Task":
		c.completeTask()
	case "🗑️ Delete Task":
		c.deleteTask()
	case "🚪 Exit":
		fmt.Println("Sampai jumpa 👋")
		c.ExitFunc(0)
		return false
	}
	return true
}

func (c *CLI) addTask() {
	var desc string
	prompt := &survey.Input{Message: "Masukkan deskripsi task:"}
	c.PromptRunner(prompt, &desc)

	if desc == "" {
		fmt.Println("❗ Deskripsi tidak boleh kosong")
		return
	}

	task := model.Task{
		Id:          time.Now().Nanosecond(),
		Description: desc,
		CreatedAt:   time.Now(),
	}
	c.repo.AddTask(task)
	fmt.Println("✅ Task berhasil ditambahkan!")
}

func (c *CLI) listTasks() {
	tasks, _ := c.repo.GetAllTasks()
	if len(tasks) == 0 {
		fmt.Println("📭 Tidak ada task.")
		return
	}
	fmt.Println("📋 Daftar Task:")
	for _, task := range tasks {
		status := "❌"
		if task.CompletedAt != nil {
			status = "✅"
		}
		fmt.Printf("[%s] #%d - %s\n", status, task.Id, task.Description)
	}
}

func (c *CLI) completeTask() {
	tasks, _ := c.repo.GetAllTasks()
	if len(tasks) == 0 {
		fmt.Println("📭 Tidak ada task.")
		return
	}

	options := []string{}
	taskMap := map[string]int{}
	for _, t := range tasks {
		label := fmt.Sprintf("#%d - %s", t.Id, t.Description)
		options = append(options, label)
		taskMap[label] = t.Id
	}

	var selected string
	prompt := &survey.Select{Message: "Pilih task yang ingin diselesaikan:", Options: options}
	c.PromptRunner(prompt, &selected)

	err := c.repo.MarkTaskAsCompleted(taskMap[selected])
	if err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅ Task selesai!")
	}
}

func (c *CLI) deleteTask() {
	tasks, _ := c.repo.GetAllTasks()
	if len(tasks) == 0 {
		fmt.Println("📭 Tidak ada task.")
		return
	}

	options := []string{}
	taskMap := map[string]int{}
	for _, t := range tasks {
		label := fmt.Sprintf("#%d - %s", t.Id, t.Description)
		options = append(options, label)
		taskMap[label] = t.Id
	}

	var selected string
	prompt := &survey.Select{Message: "Pilih task yang ingin dihapus:", Options: options}
	c.PromptRunner(prompt, &selected)

	err := c.repo.DeleteTask(taskMap[selected])
	if err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("🗑️ Task berhasil dihapus.")
	}
}
