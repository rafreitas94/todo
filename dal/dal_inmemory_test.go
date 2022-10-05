package dal_test

import (
	"testing"
	"todo/dal"
)

func TestReadTaskMethod_NotFound(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	_, err := dalInterface.ReadTask("12345")

	if err == nil {
		t.Fail()
	}

	// fmt.Printf("%+v\n", task)
	// fmt.Println(err)
}

// TDD
// Test Driven Development

func TestCreateTaskMethod(t *testing.T) {
	dalInterface := dal.NewDataAccessLayer()

	createdTask, err := dalInterface.CreateTask(dal.CreateTaskRequest{
		Subject:     "Assunto da tarefa",
		Description: "descrição",
	})

	if err != nil {
		t.Fail()
	}

	if createdTask.Subject != "Assunto da tarefa" {
		t.Error("assunto nao bate")
	}

	if createdTask.Description != "descricao" {
		t.Error("descricao nao bate")
	}

	readTask, err := dalInterface.ReadTask(createdTask.ID)
	if err != nil {
		t.Error("Falhou na leitura da tarefa")
	}

	if readTask.Subject != "Assunto da tarefa" {
		t.Error("assunto nao bate")
	}

	if readTask.Description != "descricao" {
		t.Error("descricao nao bate")
	}
}
