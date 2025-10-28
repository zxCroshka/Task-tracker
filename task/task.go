package task

type Task struct{
	Id int
	Text string`json:"text"`
	Status string`json:"status"`
}

func NewTask(text string) *Task{
	return &Task{
		Id: 0,
		Text : text,
		Status: "todo",
	}
}

func (t *Task) ChangeText(text string){
	t.Text = text
}


func (t *Task) ChangeStatus(status string){
	t.Status = status
}

