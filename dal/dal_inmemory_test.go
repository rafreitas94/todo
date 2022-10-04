package dal_test

import (
	"testing"
	"todo/dal"
)

func TestReadTaskMethod_NotFound(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	_, err := dalInterface.ReadTask("um id que isso nao existe")

	if err == nil {
		t.FailNow()
	}

	if err.Error() != "não encontrado" {
		t.Fail()
	}
}

// TDD
// Test Driven Development
// fail first, fail fast
// iteracao sobre a implementacao ate o teste falhar

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
