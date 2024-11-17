package todo

import (
	"fmt"
	"sync"
)

var manager *todoManager

type todoManager struct {
	todoLists *[]todoList
	mutex     sync.Mutex
}

type todoList struct {
	name  string
	todos []todo
}

type todo struct {
	desc  string
	check bool
}

func Init() *todoManager {
	if manager != nil {
		return manager
	}

	return &todoManager{}
}

func (m *todoManager) Save() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	fmt.Println("Zapisano")
	return nil
}
