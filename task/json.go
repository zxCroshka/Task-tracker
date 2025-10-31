package task

import (
	"log"
	"time"
)

type Manager interface {
	Add()
	Delete()
}

type TaskManager struct {
	Tasks map[int]Task `json:"Tasks"`
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[int]Task, 0),
	}
}

func (t *TaskManager) Add(task *Task) {
	if err := CreateIfNotExists("tasks.json"); err != nil {
		log.Fatal(err)
	}

	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}
	newID := len(t.Tasks)
	for k := range t.Tasks {
		newID = max(newID, k)
	}
	newID += 1
	t.Tasks[newID] = *task
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
	log.Printf("Task added successfully (ID: %d)", newID)
}

func (t *TaskManager) Delete(id int) {
	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}
	if len(t.Tasks) <= id {
		log.Println("there isn't task with this id")
		return
	}
	delete(t.Tasks, id)
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
	log.Printf("Task deleted successfully (ID: %d)", id)
}

func (t *TaskManager) UpdateText(id int, text string) {
	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}
	if len(t.Tasks) <= id {
		log.Println("there isn't task with this id")
		return
	}
	task, ok := t.Tasks[id]
	if !ok {
		log.Fatal("there is not this id")
	}
	task.Text = text
	task.UpdatedAt = time.Now()
	t.Tasks[id] = task
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
}

func (t *TaskManager) UpdateStatus(id int, status string) {
	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}
	if len(t.Tasks) <= id {
		log.Println("there isn't task with this id")
		return
	}
	task := t.Tasks[id]
	task.Status = status
	task.UpdatedAt = time.Now()
	t.Tasks[id] = task
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
}

func (t *TaskManager) List() {
	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}
	for k := range t.Tasks {
		task := t.Tasks[k]
		task.Print()
	}
}

func (t *TaskManager) ListStatus(status string) {
	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(t.Tasks); i++ {
		if t.Tasks[i].Status == status {
			task := t.Tasks[i]
			task.Print()
		}
	}
}
