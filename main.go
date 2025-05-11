package main

import (
	"github.com/fardannozami/golang-todolist-cli/internal/cli"
	"github.com/fardannozami/golang-todolist-cli/internal/repository"
)

func main() {
	repo := repository.NewInMemoryTaskRepository()
	cli := cli.NewCLI(repo)
	cli.Run()
}
