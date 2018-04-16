package domain

import (
	"github.com/jinzhu/gorm"
)

// Todo "Object"
type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

// NewTodo create a new Todo by name
func NewTodo(title string) Todo {
	return Todo{
		Title:     title,
		Completed: false,
	}
}
