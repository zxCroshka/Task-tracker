package taskstructure

type Task struct{
	id int
	text string
	status string
}

func NewTask(id int, text string, status string) *Task{
	return &Task{
		id : id,
		text : text,
		status: status,
	}
}

