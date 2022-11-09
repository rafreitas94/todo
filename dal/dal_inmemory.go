package dal

import (
	"time"

	"github.com/google/uuid"
)

// id de uma tarefa -> tarefa
// TAREFA1 -> {}
// TAREFA2 -> {}
// ...
type DataAccessLayerInMemory struct {
	tasksMap map[string]Task // um map eh inicializado sempre como nulo.
}

func (DataAccessLayerInMemory) AuthenticateUser(username string, password string) error {
	panic("not implemented")
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
		Status:      "TODO",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	d.tasksMap[taskID] = newTask

	return newTask, nil
}

func (d DataAccessLayerInMemory) ReadTask(taskID string) (Task, error) {
	task, ok := d.tasksMap[taskID]
	// if ok === true, taskID existe
	if !ok {
		return task, ErrNotFound
	}

	return task, nil
}

// PatchTask atualiza parcialmente a tarefa
func (d DataAccessLayerInMemory) PatchTask(taskID string, req PatchTaskRequest) (Task, error) {

	task, err := d.ReadTask(taskID)
	if err != nil { // se a task nao existe
		return task, err
	}

	if req.Subject != nil {
		task.Subject = *req.Subject // * faz o dereference do ponteiro para o valor concreto
	}

	if req.Description != nil {
		task.Description = *req.Description // * faz o dereference do ponteiro para o valor concreto
	}

	if req.Status != nil {
		task.Status = *req.Status // * faz o dereference do ponteiro para o valor concreto
	}

	task.UpdatedAt = time.Now()

	d.tasksMap[taskID] = task

	return task, nil
}

func (d DataAccessLayerInMemory) UpdateTask(taskID string, req UpdateTaskRequest) (Task, error) {

	task, err := d.ReadTask(taskID)
	if err != nil { // se a task nao existe
		return task, err
	}

	task.Subject = req.Subject
	task.Description = req.Description
	task.Status = req.Status
	task.UpdatedAt = time.Now()

	d.tasksMap[taskID] = task

	return task, nil
}

func (d DataAccessLayerInMemory) DeleteTask(taskID string) error {

	_, err := d.ReadTask(taskID) // _ ignorar o retorno da task
	if err != nil {              // se a task nao existe
		return err
	}

	delete(d.tasksMap, taskID)

	return nil
}

func (d DataAccessLayerInMemory) ListAllTasks(req ListTaskRequest) ([]Task, error) {
	var tasks []Task

	// for é o unica keywork para iterar
	// range para maps e slices (Arrays)
	for _, task := range d.tasksMap {
		tasks = append(tasks, task)
	}

	// no caso de um slice
	// for i := 0; i < len(tasks); i++ {
	// 	task := tasks[i]
	// }

	// // equivalente com slice
	// for i, task := range tasks {

	// }

	return tasks, nil
}
