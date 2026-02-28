package model

type CreateTodoInput struct {
	Title     string `json:"title" example:"Buy milk"`
	Completed bool   `json:"completed" example:"false"`
}
