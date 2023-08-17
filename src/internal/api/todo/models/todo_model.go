package models

import "sync"

type Todo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete string `json:"is_complete"`
}

type TodoManager struct {
	todos []Todo
	m     sync.Mutex // We Want All of Our Operations to be Atomic
}

func NewTodoManager() TodoManager {
	return TodoManager{
		todos: make([]Todo, 0),
		m:     sync.Mutex{},
	}
}
