package task

import (
	"log"
	"time"
)

type Task struct {
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTask(text string) *Task {
	return &Task{
		Text:      text,
		Status:    "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *Task) ChangeText(text string) {
	t.Text = text
}

func (t *Task) ChangeStatus(status string) {
	t.Status = status
}
func (t *Task) Print() {
	log.Printf("{\n	Text: %s,\n	Status: %s,\n	CreatedAt: %s,\n	UpdatedAt: %s\n}\n", t.Text, t.Status, t.CreatedAt, t.UpdatedAt)
}
