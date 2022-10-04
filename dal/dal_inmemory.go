package dal

import (
	"errors"

	"github.com/google/uuid"
)

// id de uma tarefa -> tarefa
// TAREFA1 -> {}
// TAREFA2 -> {}
// ...
type DataAccessLayerInMemory struct {
	tasksMap map[string]Task // um map eh inicializado sempre como nulo.
}

// - Task
// --- CreateTaskRequest

func (d DataAccessLayerInMemory) CreateTask(req CreateTaskRequest) (Task, error) {
	// uuid
	// unique univesal identifier
	// bbbea3d8-d57c-4f14-bf3e-3888470a8d76

	taskID := uuid.NewString()

	newTask := Task{
		ID:          taskID,
		Subject:     req.Subject,
		Description: req.Description,
	}

	d.tasksMap[taskID] = newTask

	return newTask, nil
}

func (d DataAccessLayerInMemory) ReadTask(taskID string) (Task, error) {
	task, ok := d.tasksMap[taskID]
	// if ok === true, taskID existe
	if !ok {
		return task, errors.New("não encontrado")
	}

	return task, nil
}

func (DataAccessLayerInMemory) UpdateTask(req UpdateTaskRequest) (Task, error) {
	return Task{}, nil
}

func (DataAccessLayerInMemory) DeleteTask(taskID string) error {
	return nil
}

func (DataAccessLayerInMemory) ListAllTasks(req ListTaskRequest) ([]Task, error) {
	// var task []Task
	// task := []Task{}
	// task = nil

	// nil é um slice vazio.
	return nil, nil
}
