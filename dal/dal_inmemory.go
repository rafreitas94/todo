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
	taskMap map[string]Task // Um map eh inicializado sempre como nulo.
}

func (d DataAccessLayerInMemory) CreateTask(req CreateTaskRequest) (Task, error) {
	// UUID
	// unique universal identifier
	// b90d4993-d1f3-41d7-8617-cfcdb9c44c8f

	taskId := uuid.NewString()
	newTask := Task{
		ID:          taskId,
		Subject:     req.Subject,
		Description: req.Description,
	}

	d.taskMap[newTask.ID] = newTask

	return newTask, nil
}

func (d DataAccessLayerInMemory) ReadTask(taskId string) (Task, error) {
	task, ok := d.taskMap[taskId]
	// if ok === true, taskID existe
	if !ok {
		return task, errors.New("não encontrado")
	}
	return task, nil
}

func (DataAccessLayerInMemory) UpdateTask(req UpdateTaskRequest) (Task, error) {
	return Task{}, nil
}

func (DataAccessLayerInMemory) DeleteTask(taskId string) error {
	return nil
}

func (DataAccessLayerInMemory) ListAllTask(req ListTaskRequest) ([]Task, error) {
	// var task []Task
	// task := []Task{}
	// task = nil

	// nil é um slice vazio.
	return nil, nil
}

// func exemplo() DataAcessLayerInterface {
// 	return DataAccessLayerInMemory{}
// }
