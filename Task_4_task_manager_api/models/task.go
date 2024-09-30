package models

type Task struct {
	ID          string `json:"id"` // Unique identifier for the task
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	DueDate     string `json:"due_date" binding:"required"`
	Status      string `json:"status"`
}
