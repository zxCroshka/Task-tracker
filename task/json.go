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
	Tasks []Task `json:"Tasks"`
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make([]Task, 0),
	}
}

func (t *TaskManager) Add(task *Task) {
	if err := CreateIfNotExists("tasks.json"); err != nil {
		log.Fatal(err)
	}

	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}

	t.Tasks = append(t.Tasks, *task)
	n := len(t.Tasks)
	t.Tasks[n-1].Id = n - 1
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
	log.Printf("Task added successfully (ID: %d)",n-1)
}

func (t *TaskManager) Delete(id int) {
	if err := JSONtoStruct(&t); err != nil {
		log.Fatal(err)
	}
	if len(t.Tasks) <= id {
		log.Println("there isn't task with this id")
		return
	}
	t.Tasks = append(t.Tasks[:id], t.Tasks[id+1:]...)
	for i := 0; i < len(t.Tasks); i++ {
		t.Tasks[i].Id = i
	}
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
	log.Printf("Task deleted successfully (ID: %d)",id)
}

func (t *TaskManager) UpdateText(id int, text string){
	if err := JSONtoStruct(&t); err != nil{
		log.Fatal(err)
	}
	if len(t.Tasks) <= id {
		log.Println("there isn't task with this id")
		return
	}
	t.Tasks[id].Text = text
	t.Tasks[id].UpdatedAt = time.Now()
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
}


func (t *TaskManager) UpdateStatus(id int, status string){
	if err := JSONtoStruct(&t); err != nil{
		log.Fatal(err)
	}
	if len(t.Tasks) <= id {
		log.Println("there isn't task with this id")
		return
	}
	t.Tasks[id].Status = status
	t.Tasks[id].UpdatedAt = time.Now()
	if err := StructToJSON(t); err != nil {
		log.Fatal(err)
	}
}

func (t *TaskManager) List(){
	if err := JSONtoStruct(&t); err != nil{
		log.Fatal(err)
	}
	for i:= 0; i < len(t.Tasks);i++{
		log.Println(t.Tasks[i])
	}
}

func (t *TaskManager) ListStatus(status string){
	if err := JSONtoStruct(&t); err != nil{
		log.Fatal(err)
	}
	for i:= 0; i < len(t.Tasks);i++{
		if t.Tasks[i].Status == status{
			log.Println(t.Tasks[i])
		}
	}
}