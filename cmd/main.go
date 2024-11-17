package main

import (
	"fmt"

	"github.com/4d4mszweda/game-of-life-cli/pkg/todo"
)

func main() {
	fmt.Println("Hello world")

	todos := todo.Init()
	todos.Save()
}
