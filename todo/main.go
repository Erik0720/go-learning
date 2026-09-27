package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Task struct {
	ID       int
	Name     string
	Deadline time.Time
	Done     bool
}

func main() {
	t := Task{ID: 1, Name: "Test", Deadline: time.Now(), Done: false}
	data, _ := json.Marshal(t)
	fmt.Println(string(data))
}
