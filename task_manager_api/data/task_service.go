package data

import (
	"errors"
	"task_manager/models"
	"time"

	"github.com/google/uuid"
)

// Mock data for tasks
var Tasks = []models.Task{
	{ID: uuid.New().String(), Title: "Task 1", Description: "First task", DueDate: time.Now().Format("2006-01-02"), Status: "Pending"},
	{ID: uuid.New().String(), Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1).Format("2006-01-02"), Status: "In Progress"},
	{ID: uuid.New().String(), Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2).Format("2006-01-02"), Status: "Completed"},
}

// getTasks responds with the list of all tasks as JSON.
func GetTasks() []models.Task {
	return Tasks
}

// getTask responds with the details of a specific task.
func GetTaskById(taskID string) (models.Task, error) {
	for _, task := range Tasks {
		if task.ID == taskID {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

// createTask creates a new task.
func CreateTask(newTask models.Task) models.Task {
	newTask.ID = uuid.New().String()
	Tasks = append(Tasks, newTask)
	return newTask
}

// updateTask updates a specific task.
func UpdateTask(taskID string, updatedTask models.Task) (models.Task, error) {
	for i, task := range Tasks {
		if task.ID == taskID {
			if updatedTask.Title != "" {
				Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				Tasks[i].Description = updatedTask.Description
			}
			if updatedTask.DueDate != "" {
				Tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				Tasks[i].Status = updatedTask.Status
			}
			return Tasks[i], nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

// deleteTask deletes a specific task by ID.
func DeleteTask(taskID string) error {
	for i, task := range Tasks {
		if task.ID == taskID {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
