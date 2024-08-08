package api

import (
	"errors"
	"github.com/google/uuid"
)

type TaskStatus struct {
	status string
	result string
}

type task struct {
	data map[string]TaskStatus
}

func NewTask() *task {
	return &task{
		data: make(map[string]TaskStatus),
	}
}
func (t *task) newID() string {
	uID := uuid.New().String()
	t.data[uID] = TaskStatus{"in_progress", "in_progress"}
	return uID
}
func (t *task) get(Key string) (TaskStatus, error) {
	if value, ok := t.data[Key]; ok {
		return value, nil
	}
	return TaskStatus{"", ""}, errors.New("key not found")
}
func (t *task) set(Key string, value TaskStatus) error {
	if _, exists := t.data[Key]; exists {
		return errors.New("key already exists")
	}
	t.data[Key] = value
	return nil
}
