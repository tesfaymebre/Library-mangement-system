package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: uuid.New().String(), Title: "Task 1", Description: "First task", DueDate: time.Now().Format("2006-01-02"), Status: "Pending"},
	{ID: uuid.New().String(), Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"), Status: "In Progress"},
	{ID: uuid.New().String(), Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2).Format("2006-01-02"), Status: "Completed"},
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run("localhost:8080") // Listen and serve on 0.0.0.0:8080
}

// getTasks responds with the list of all tasks as JSON.
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

// geting a specific task
func getTask(c *gin.Context) {
	taskID := c.Param("id")
	for _, task := range tasks {
		if task.ID == taskID {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// create a new task
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	newTask.ID = uuid.New().String()
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// update a specific task
func updateTask(c *gin.Context) {
	// update each attribute of the task if it is provided
	taskID := c.Param("id")

	var updatedTask Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			// update the task
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.DueDate != "" {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			c.IndentedJSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})

}

// delete a specific task by id
func deleteTask(c *gin.Context) {
	taskID := c.Param("id")
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
