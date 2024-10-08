package controllers

import (
	"net/http"

	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// getTasks responds with the list of all tasks as JSON.
func GetTasks(c *gin.Context) {
	if len(data.GetTasks()) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no tasks found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": data.GetTasks()})
}

// getTask responds with the details of a specific task.
func GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	if task, err := data.GetTaskById(taskID); err == nil {
		c.IndentedJSON(http.StatusOK, task)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// createTask creates a new task.
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	createdTask, err := data.CreateTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to create task"})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdTask)
}

// updateTask updates a specific task.
func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var updatedTask models.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if task, err := data.UpdateTask(taskID, updatedTask); err == nil {
		c.IndentedJSON(http.StatusOK, task)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// deleteTask deletes a specific task by ID.
func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if err := data.DeleteTask(taskID); err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
