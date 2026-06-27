package storage

import (
	"encoding/json"
	"example/task-tracker/model"
	"os"
)

func LoadTasks() ([]model.Task, error) {

	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		return []model.Task{}, nil
	}
	
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return []model.Task{}, err
	}
	
	var task []model.Task
	
	err = json.Unmarshal(file, &task)
	if err != nil {
		return []model.Task{}, err
	}

	return task, nil

}

func SaveTasks(tasks []model.Task) error {

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}