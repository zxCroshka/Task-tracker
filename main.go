package main

import "tasktracker/task"

func main() {
	t := task.NewTask("croha1")
	t1 := task.NewTask("salyam2")
	tm := task.NewTaskManager()
	tm.Add(t)
	tm.Add(t1)
	tm.Delete(5)
	tm.Delete(0)
}
