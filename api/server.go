package api

import (
	"encoding/json"
	_ "firsthw/docs"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Storage interface {
	get(Key string) (TaskStatus, error)
	set(Key string, value TaskStatus) error
	newID() string
	processTask(Key string)
}
type Server struct {
	storage Storage
}

func newServer(storage Storage) *Server {
	return &Server{storage: storage}
}

// @Summary      Creating new id for the task
// @Tags task
// @Description  Creating new id for the task and returning
// @Success      200 {string} string "Value"
// @Router       /task [get]
func (s *Server) newTaskHandler(w http.ResponseWriter, r *http.Request) {
	//NewTask()
	a := s.storage.newID()
	go s.storage.processTask(a)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"task_id": a})

}

// @Summary      get status_value
// @Description  returns value_status by task id
// @Param        key query string true "Key"
// @Success      200  {string}  string "value1"
// @Failure      400  {string}  string error
// @Router       /status/ [get]
func (s *Server) getHandler_Status(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	fmt.Printf("Received key: %s\n", key)

	if key == "" {
		http.Error(w, "Missing key", http.StatusNotFound)
		return
	}

	value, err := s.storage.get(key)
	if err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	value1 := value.Status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Перенесем установку статуса перед отправкой JSON
	json.NewEncoder(w).Encode(map[string]string{"status": value1})
}

// @Summary      get result_value
// @Description  returns result_status by task id
// @Param        key query string true "Key"
// @Success      200  {string}  string "value1"
// @Failure      400  {string} error
// @Router       /result/ [get]
func (s *Server) getHandler_Result(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	fmt.Printf("Received key: %s\n", key)

	if key == "" {
		http.Error(w, "Missing key", http.StatusNotFound)
		return
	}

	value, err := s.storage.get(key)
	if err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	value1 := value.Result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": value1})
	w.WriteHeader(http.StatusOK)
}

func CreateAndRunServer(storage Storage, addr string) error {
	server := newServer(storage)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/task", server.newTaskHandler)
	http.HandleFunc("/status/", server.getHandler_Status)
	http.HandleFunc("/result/", server.getHandler_Result)
	return http.ListenAndServe(addr, nil)
}

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
