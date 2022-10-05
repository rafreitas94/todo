package dal_test

import (
	"testing"
	"todo/dal"
)

func TestReadTaskMethod_NotFound(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	_, err := dalInterface.ReadTask("um id que nao existe")

	if err == nil {
		t.Error("erro era esperado")
		t.FailNow()
	}

	if err.Error() != "não encontrado" {
		t.Error("esperava erro 'não encontrado'")
		t.Fail()
	}
}

func TestReadTaskMethod_Found(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	newTask, err := dalInterface.CreateTask(dal.CreateTaskRequest{
		Subject:     "Assunto da tarefa",
		Description: "descricao",
	})

	if err != nil {
		t.Error("erro não era esperado")
		t.FailNow()
	}

	readTask, err := dalInterface.ReadTask(newTask.ID)

	if err != nil {
		t.Error("erro não era esperado")
		t.FailNow()
	}

	if readTask.Subject != "Assunto da tarefa" {
		t.Error("assunto nao bate")
	}

	if readTask.Description != "descricao" {
		t.Error("descricao nao bate")
	}
}

func TestDeleteTaskMethod_NotFound(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	err := dalInterface.DeleteTask("id que não existe")

	if err == nil {
		t.Error("erro era esperado")
		t.FailNow()
	}

	if err.Error() != "não encontrado" {
		t.Error("esperava erro 'não encontrado'")
		t.Fail()
	}
}

func TestDeleteTaskMethod_Found(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	newTask, err := dalInterface.CreateTask(dal.CreateTaskRequest{
		Subject:     "Assunto da tarefa",
		Description: "descricao",
	})

	if err != nil {
		t.Error("erro não era esperado")
		t.FailNow()
	}

	err = dalInterface.DeleteTask(newTask.ID)

	if err != nil {
		t.Error("erro não era esperado")
		t.FailNow()
	}
}

func TestUpdateTaskMethod_NotFound(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	_, err := dalInterface.UpdateTask("id que não existe", dal.UpdateTaskRequest{
		Subject:     "novo assunto",
		Description: "nova descricao",
	})

	if err == nil {
		t.Error("erro era esperado")
		t.FailNow()
	}

	if err.Error() != "não encontrado" {
		t.Error("esperava erro 'não encontrado'")
		t.Fail()
	}
}

func TestUpdateTaskMethod_Found(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	newTask, err := dalInterface.CreateTask(dal.CreateTaskRequest{
		Subject:     "Assunto da tarefa",
		Description: "descricao",
	})

	if err != nil {
		t.Error("erro não era esperado")
		t.FailNow()
	}

	updatedTask, err := dalInterface.UpdateTask(newTask.ID, dal.UpdateTaskRequest{
		Subject:     "novo assunto",
		Description: "nova descricao",
	})

	if err != nil {
		t.Error("erro não era esperado")
		t.FailNow()
	}

	if updatedTask.Subject != "novo assunto" {
		t.Error("assunto nao bate")
	}

	if updatedTask.Description != "nova descricao" {
		t.Error("descricao nao bate")
	}

	readTask, err := dalInterface.ReadTask(newTask.ID)
	if err != nil {
		t.Error("falhou na leitura da tarefa")
	}

	if readTask.Subject != "novo assunto" {
		t.Error("assunto nao bate")
	}

	if readTask.Description != "nova descricao" {
		t.Error("descricao nao bate")
	}
}

func TestListAllTasks(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	dalInterface.CreateTask(dal.CreateTaskRequest{})
	dalInterface.CreateTask(dal.CreateTaskRequest{})
	dalInterface.CreateTask(dal.CreateTaskRequest{})
	dalInterface.CreateTask(dal.CreateTaskRequest{})

	taskList, err := dalInterface.ListAllTasks(dal.ListTaskRequest{})
	if err != nil {
		t.Error("falhou na listagem de tarefas")
		t.FailNow()
	}

	if len(taskList) != 4 {
		t.Error("nao retornou o numero esperado de tarefas")
	}
}

// TDD
// Test Driven Development
// fail first, fail fast
// iteracao sobre a implementacao ate o teste passar

func TestCreateTaskMethod(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	createdTask, err := dalInterface.CreateTask(dal.CreateTaskRequest{
		Subject:     "Assunto da tarefa",
		Description: "descricao",
	})

	if err != nil {
		t.FailNow()
	}

	if createdTask.Subject != "Assunto da tarefa" {
		t.Error("assunto nao bate")
	}

	if createdTask.Description != "descricao" {
		t.Error("descricao nao bate")
	}

	readTask, err := dalInterface.ReadTask(createdTask.ID)
	if err != nil {
		t.Error("falhou na leitura da tarefa")
	}

	if readTask.Subject != "Assunto da tarefa" {
		t.Error("assunto nao bate")
	}

	if readTask.Description != "descricao" {
		t.Error("descricao nao bate")
	}
}
