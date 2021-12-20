package models

type Todo struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"title"`
	Body  string `json:"body" gorm:"body"`
}

type CreateTodoStruct struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type UpdateTodo struct {
	Title string `json:"title"`
	Body  string `json:"body" `
}
