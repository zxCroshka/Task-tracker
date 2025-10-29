package main

import (
	"log"
	"os"
	"strconv"
	"github.com/zxCroshka/Task-tracker/task"

)


func main() {
	tm := task.NewTaskManager()
	cmd := os.Args[1]
	if cmd == "add" {
		t := task.NewTask(os.Args[2])
		tm.Add(t)
	}
	if cmd == "update" {
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		text := os.Args[3]
		tm.UpdateText(id, text)

	}
	if cmd == "delete" {
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		tm.Delete(id)
	}
	if cmd == "mark" {
		status := os.Args[2]
		marks := [3]string{"todo","in-progress","done"}
		flag := false
		for i:= 0; i < 3;i++{
			if status == marks[i]{
				flag = true
				break
			}
		}
		if flag{
			id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		tm.UpdateStatus(id, status)
		}else{
			log.Println("Invalid status of task")
		}
		
	}
	if cmd == "list" {
		if len(os.Args) == 2 {
			tm.List()
		} else {
			status := os.Args[2]
			tm.ListStatus(status)
		}
	}
	// t := task.NewTask("uga")
	// t1 := task.NewTask("buga")

	// tm.Add(t)
	// tm.Add(t1)
	// tm.List()
	// tm.Delete(5)
	// tm.Delete(0)
	// t2 := task.NewTask("croha")
	// t3 := task.NewTask("ashot")
	// tm.Add(t2)
	// tm.Add(t3)
	// tm.UpdateText(0, "buga updated")
	// tm.UpdateStatus(1, "in-progress")
	// tm.List()
	// tm.ListStatus("in-progress")

}
