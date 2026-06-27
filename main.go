package main

import (
	"example/task-tracker/model"
	"example/task-tracker/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Silahkan masukkan perintah! Contoh: add \"Belajar Go\" atau list")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Deskripsi task tidak boleh kosong! Contoh: add \"Belajar Go\"")
			return
		}

		description := os.Args[2]
		if strings.TrimSpace(description) == "" {
			fmt.Println("Deskripsi task tidak boleh kosong! Contoh: add \"Belajar Go\"")
			return
		}

		newID, err := service.AddTask(description)
		if err != nil {
			fmt.Printf("Gagal menambah task: %v\n", err)
			return
		}

		fmt.Printf("Task Added Successfully (ID: %d)\n", newID)

	case "list":
		statusFilter := ""
		
		if len(os.Args) >= 3 {
			statusFilter = os.Args[2]
		}

		err := service.ListTasks(statusFilter)
		if err != nil {
			fmt.Printf("Gagal menampilkan task: %v\n", err)
			return
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: ID dan deskripsi task tidak boleh kosong! Contoh: update 1 \"Belajar Go Advanced\"")
			return
		}

		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID harus integer! Contoh: update 1 \"Belajar Go Advanced\"")
			return
		}

		description := os.Args[3]
		if strings.TrimSpace(description) == "" {
			fmt.Println("Deskripsi task tidak boleh kosong! Contoh: add \"Belajar Go\"")
			return
		}

		err = service.UpdateTask(i, description)
		if err != nil {
			fmt.Printf("Gagal memperbarui task: %v\n", err)
			return
		}

		fmt.Println("Task updated successfully")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID task tidak boleh kosong! Contoh: delete 1")
			return
		}

		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID harus integer! Contoh: delete 1")
			return
		}

		err = service.DeleteTask(i)
		if err != nil {
			fmt.Printf("Gagal menghapus task: %v\n", err)
			return
		}

		fmt.Println("Task deleted successfully")

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID task tidak boleh kosong! Contoh: mark-in-progress 1")
			return
		}
		
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID harus integer! Contoh: mark-in-progress 1")
			return
		}

		err = service.MarkTaskStatus(i, model.StatusInProgress)
		if err != nil {
			fmt.Printf("Gagal memperbarui status: %v\n", err)
			return
		}

		fmt.Println("Task status updated successfully")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID task tidak boleh kosong! Contoh: mark-done 1")
			return
		}

		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID harus integer! Contoh: mark-done 1")
			return
		}

		err = service.MarkTaskStatus(i, model.StatusDone)
		if err != nil {
			fmt.Printf("Gagal memperbarui status: %v\n", err)
			return
		}

		fmt.Println("Task status updated successfully")

	case "stats":
		if len(os.Args) < 2 {
			fmt.Println("Error: perintah hanya stats saja")
			return
		}

		err := service.TaskStats()
		if err != nil {
			fmt.Printf("Gagal menampilkan stats task: %v\n", err)
			return
		}

	default:
		fmt.Println("Perintah tidak dikenali!")
	}
}