package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const path = "tasks.json"

type Task struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Deadline time.Time `json:"deadline"`
	Done     bool      `json:"done"`
}

func main() {
	t := []Task{{ID: 1, Name: "Test", Deadline: time.Now(), Done: false}, {ID: 2, Name: "Beispiel", Deadline: time.Now().Add(time.Hour * 4), Done: true}}

	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Failed to encode JSON!")
		return
	}

	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		fmt.Println("Failed to write File!")
		return
	}

	data, err = os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read File!")
		return
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Failed to decode JSON!")
		return
	}

	fmt.Println(tasks)
}

func write(data []Task) {
}
