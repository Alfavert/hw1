package api

import (
	"encoding/json"
	"net/http"
)

type Storage interface {
	get(Key string) (TaskStatus, error)
	set(Key string, value TaskStatus) error
	newID() string
}
type Server struct {
	storage Storage
}

func newServer(storage Storage) *Server {
	return &Server{storage: storage}
}
func (s *Server) newTaskHandler(w http.ResponseWriter, r *http.Request) {
	a := s.storage.newID()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"task": a})
}
func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}
	value, err := s.storage.get(key)
	if err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]TaskStatus{"task": value})
}
func CreateAndRunServer(storage Storage, addr string) error {
	server := newServer(storage)
	http.HandleFunc("/task", server.newTaskHandler)
	http.HandleFunc("/status/", server.getHandler)
	return http.ListenAndServe(addr, nil)
}

//func generateUUID64() string {
//	uID := uuid.New().String()
//	return uID
//}
//func createTask(task *task) string {
//	task.ID = generateUUID64()
//	task.TaskStatus = TaskStatus{"in_progress", "not_given"}
//}
//func writeTask(w http.ResponseWriter, r *http.Request) {
//	task := createTask()
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(map[string]string{"task_id": task.ID,
//		"task_status": task.TaskStatus.status,
//		"task_result": task.TaskStatus.result})
//}
//func processTask(u *task) {
//	time.Sleep(5 * time.Second)
//	u.TaskStatus = TaskStatus{"completed", "hello"}
//}
//
//func getStatus(w http.ResponseWriter, r *http.Request, task task) {
//	taskID := r.URL.Path[len("/status/"):]
//	if taskID != task.ID {
//		http.Error(w, "Task not found", http.StatusNotFound)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(map[string]string{"task_id": task.ID,
//		"task_status": task.TaskStatus.status,
//		"task_result": task.TaskStatus.result})
//}
//func getResult(w http.ResponseWriter, r *http.Request, task task) {
//	taskID := r.URL.Path[len("/result/"):]
//	if taskID != task.ID {
//		http.Error(w, "Task not found", http.StatusNotFound)
//		return
//	}
//	if task.TaskStatus.status != "ready" {
//		http.Error(w, "Task not ready", http.StatusAccepted)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(map[string]string{"task_id": task.ID,
//		"task_status": task.TaskStatus.status,
//		"task_result": task.TaskStatus.result})
//}
