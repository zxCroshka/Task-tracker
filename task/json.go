package task

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Manager interface {
	Add()
	Delete()
}

type TaskManager struct {
	Tasks []Task `json:"Tasks"`
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make([]Task, 0),
	}
}

func (t *TaskManager) Add(task *Task) {
	if flag, err := exists("tasks.json"); err == nil {
		if !flag {
			file, err := os.Create("tasks.json")
			if err != nil {
				log.Fatal(err)
			}
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		}

	}
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewDecoder(file).Decode(&t); err != nil && err != io.EOF {
		log.Fatal(err)
	}

	t.Tasks = append(t.Tasks, *task)
	n := len(t.Tasks)
	t.Tasks[n-1].Id = n - 1
	file, err = os.OpenFile("tasks.json", os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = json.NewEncoder(file).Encode(t); err != nil {
		log.Fatal(err)
	}
}

func (t *TaskManager) Delete(id int){
	file, err := os.Open("tasks.json")
	if err != nil{
		log.Fatal(err)
	}
	if err := json.NewDecoder(file).Decode(&t); err != nil{
		log.Fatal(err)
	}
	if len(t.Tasks) <= id {
		fmt.Println("there isnt task with this id")
		return
	}
	if err := file.Close(); err != nil{
		log.Fatal(err)
	}
	t.Tasks = append(t.Tasks[:id],t.Tasks[id+1:]...)
	for i:= 0; i < len(t.Tasks);i++{
		t.Tasks[i].Id = i
	}
	file, err = os.OpenFile("tasks.json", os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = json.NewEncoder(file).Encode(t); err != nil {
		log.Fatal(err)
	}
}