package api

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (t *task) newID() string {
	uID := uuid.New().String()
	fmt.Printf("Generated new ID: %s\n", uID)
	t.data[uID] = TaskStatus{"in_progress", "not_ready"}
	return uID
}
func (t *task) get(Key string) (TaskStatus, error) {
	if _, ok := t.data[Key]; ok {
		return t.data[Key], nil
	}

	return TaskStatus{"", ""}, errors.New("key not found")
}

func (t *task) set(Key string, value TaskStatus) error {
	if _, exists := t.data[Key]; exists {
		return errors.New("key already exists")
	}
	fmt.Printf("Setting task for key %s: %+v\n", Key, value)
	t.data[Key] = value
	return nil
}
func (t *task) processTask(Key string) {
	time.Sleep(60 * time.Second)
	t.data[Key] = TaskStatus{"completed", "hello"}
}
