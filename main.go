package main

import (
	"fmt"
	"os"

	"github.com/kamogelosekhukhune777/go-todo/cmd"
)

func main() {
	if err := cmd.Todos.Load(cmd.TodoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	cmd.Execute()
}
