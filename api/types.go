package api

type TaskStatus struct {
	Status string `json:"status"`
	Result string `json:"result"`
}
type task struct {
	data map[string]TaskStatus `json:"data"`
}

func NewTask() *task {
	return &task{
		data: make(map[string]TaskStatus),
	}
}
func (t *task) NewTask() *task {
	return &task{
		data: make(map[string]TaskStatus),
	}
}
