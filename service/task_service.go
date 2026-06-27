package service

import (
	"errors"
	"example/task-tracker/model"
	"example/task-tracker/storage"
	"fmt"
	"strings"
	"time"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidStatus = errors.New("invalid status")
	ErrEmptyDescription = errors.New("description cannot be empty")
)

func isValidStatus(status string) bool {
	return status == model.StatusTodo || status == model.StatusInProgress || status == model.StatusDone
}

func findTaskIndexByID(tasks []model.Task, id int) int {
	for i := range tasks {
		if tasks[i].ID == id {
			return i
		}
	}

	return -1
}

func AddTask(description string) (int, error) {
	if strings.TrimSpace(description) == "" {
		return 0, ErrEmptyDescription
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	nextID := maxID + 1

	now := time.Now()
	newTask := model.Task {
		ID: nextID,
		Description: description,
		Status: model.StatusTodo,
		CreatedAt: now,
		UpdatedAt: now,
	}

	tasks = append(tasks, newTask)

	err = storage.SaveTasks(tasks)
	if err != nil {
		return 0, err
	}

	return nextID, nil
}

func ListTasks(status string) error {

	if status != "" && !isValidStatus(status) {
		return ErrInvalidStatus
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}
	
	if len(tasks) == 0 {
		fmt.Println("Tidak ada task yang ditemukan.")
		return nil
	}

	hasDisplayed := false
	for _, task := range tasks {
		if status != "" && task.Status != status {
			continue
		}

		fmt.Println("======================")
		fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n", 
				task.ID, 
				task.Description, 
				task.Status, 
				task.CreatedAt.Format("2006-01-02 15:04"), 
				task.UpdatedAt.Format("2006-01-02 15:04"),
		)
		hasDisplayed = true
	}

	if !hasDisplayed {
		fmt.Printf("Tidak ada task dengan status '%s'.\n", status)
	}

	return nil
}

func UpdateTask(id int, description string) error {
	if strings.TrimSpace(description) == "" {
		return ErrEmptyDescription
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	idx := findTaskIndexByID(tasks, id)
	if idx == -1 {
		return ErrTaskNotFound
	}

	tasks[idx].Description = description
	tasks[idx].UpdatedAt = time.Now()
	
	return storage.SaveTasks(tasks)
}

func DeleteTask(id int) error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	idx := findTaskIndexByID(tasks, id)
	if idx == -1 {
		return ErrTaskNotFound
	}

	tasks = append(tasks[:idx], tasks[idx+1:]... )

	return storage.SaveTasks(tasks)
}

func MarkTaskStatus(id int, status string) error {
	if !isValidStatus(status) {
		return ErrInvalidStatus
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	idx := findTaskIndexByID(tasks, id)
	if idx == -1 {
		return ErrTaskNotFound
	}

	tasks[idx].Status = status
	tasks[idx].UpdatedAt = time.Now()

	return storage.SaveTasks(tasks)
}

func TaskStats() error {
	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}
	
	totalTask := len(tasks)
	todoCount := 0
	inProgressCount := 0
	doneCount := 0

	for _, task := range tasks {
		switch task.Status {
		case model.StatusTodo:
			todoCount++
		case model.StatusInProgress:
			inProgressCount++
		case model.StatusDone:
			doneCount++
		}
	}

	fmt.Println("===== Task Statistics =====")
	fmt.Printf("Total Tasks: %d\n", totalTask)
	fmt.Printf("Todo: %d\n", todoCount)
	fmt.Printf("In Progress: %d\n", inProgressCount)
	fmt.Printf("Done: %d\n", doneCount)

	return nil
}