package dal

type DataAccessLayerSQL struct {
}

func (d DataAccessLayerSQL) CreateTask(req CreateTaskRequest) (Task, error) {
	panic("not implemented")
}

func (d DataAccessLayerSQL) ReadTask(taskID string) (Task, error) {
	panic("not implemented")
}

// PatchTask atualiza parcialmente a tarefa
func (d DataAccessLayerSQL) PatchTask(taskID string, req PatchTaskRequest) (Task, error) {
	panic("not implemented")
}

func (d DataAccessLayerSQL) UpdateTask(taskID string, req UpdateTaskRequest) (Task, error) {
	panic("not implemented")
}

func (d DataAccessLayerSQL) DeleteTask(taskID string) error {
	panic("not implemented")
}

func (d DataAccessLayerSQL) ListAllTasks(req ListTaskRequest) ([]Task, error) {
	panic("not implemented")
}
