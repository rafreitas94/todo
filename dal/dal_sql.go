package dal

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DataAccessLayerSQL struct {
	db          *sqlx.DB
	redisClient *redis.Client
}

func NewDataAccessLayerSQL(db *sqlx.DB, redisClient *redis.Client) *DataAccessLayerSQL {
	return &DataAccessLayerSQL{
		db:          db,
		redisClient: redisClient,
	}
}

func (d DataAccessLayerSQL) CreateTask(req CreateTaskRequest) (Task, error) {
	// d.db.Exec("UPDATE/INSERT")
	// return d.ReadTask(...)\
	/*
		id text primary key,
		subject text not null,
		description text not null,
		status text not null,
		created_at timestamp not null default now(),
		updated_at timestamp not null default now()
	*/

	task := Task{
		ID:          uuid.NewString(),
		Subject:     req.Subject,
		Description: req.Description,
		Status:      "TODO",
		// Incluir um userID
	}

	err := d.db.Get(&task, `
		INSERT INTO tasks (id, subject, description, status)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`, task.ID, task.Subject, task.Description, task.Status)

	// alternativa
	// d.db.Exec(`
	// 	INSERT INTO tasks (id, subject, description, status)
	// 	VALUES ($1, $2, $3, $4)
	// `, uuid.NewString(), req.Subject, req.Description, "TODO")

	// return d.ReadTask(task.id)
	return task, err
}

func (d DataAccessLayerSQL) ReadTask(taskID string) (Task, error) {
	var task Task

	// SELECT * tasks WHERE id = 'abcd'
	err := d.db.Get(&task, "SELECT * FROM tasks WHERE id = $1", taskID)
	if err != nil {
		return task, err
	}

	return task, nil
}

// PatchTask atualiza parcialmente a tarefa
func (d DataAccessLayerSQL) PatchTask(taskID string, req PatchTaskRequest) (Task, error) {
	// d.db.Exec("UPDATE/INSERT")

	var task Task

	err := d.db.Get(&task, `
		UPDATE tasks 
		SET subject = COALESCE($1, subject)
			description = COALESCE($2, description),
			status = COALESCE($3, status),
			updated_at = now()
		WHERE id = $4
		RETURNING *
	`, req.Subject, req.Description, req.Status, taskID)

	return task, err
}

func (d DataAccessLayerSQL) UpdateTask(taskID string, req UpdateTaskRequest) (Task, error) {
	// d.db.Exec("UPDATE/INSERT")
	//
	task := Task{
		ID:          taskID,
		Subject:     req.Subject,
		Description: req.Description,
		Status:      req.Status,
	}

	err := d.db.Get(&task, `
		UPDATE tasks 
		SET subject = $1,
			description = $2,
			status = $3,
			updated_at = now()
		WHERE id = $4
		RETURNING *
	`, task.Subject, task.Description, task.Status, task.ID)

	return task, err
}

func (d DataAccessLayerSQL) DeleteTask(taskID string) error {
	// d.db.Exec("DELETE")

	redisKey := "operation:list-todos"
	d.redisClient.Del(context.Background(), redisKey)
	_, err := d.db.Exec("DELETE FROM tasks WHERE id = $1", taskID)
	return err
}

func (d DataAccessLayerSQL) ListAllTasks(req ListTaskRequest) ([]Task, error) {
	var listaDeTarefas []Task

	redisKey := "operation:list-todos"
	cachedValue := d.redisClient.Get(context.Background(), redisKey).Val()
	if cachedValue != "" {
		err := json.Unmarshal([]byte(cachedValue), &listaDeTarefas)
		// retornar valor cacheado.
		return listaDeTarefas, err
	}

	err := d.db.Select(&listaDeTarefas, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	// aqui eh para evitar a serializaçao de uma lista nula (vazia) em go
	// como nulo em json.
	if listaDeTarefas == nil {
		listaDeTarefas = []Task{}
	}

	byteValue, err := json.Marshal(listaDeTarefas)
	if err != nil {
		return nil, err
	}
	err = d.redisClient.Set(context.Background(), redisKey, string(byteValue), 30*time.Second).Err()
	if err != nil {
		return nil, err
	}

	return listaDeTarefas, err
}

func (d DataAccessLayerSQL) AuthenticateSession(sessionID string) (string, error) {
	redisKey := "session:" + sessionID
	sessionUsername, err := d.redisClient.Get(context.Background(), redisKey).Result()
	if err != nil {
		return "", err
	}

	return sessionUsername, nil
}

func (d DataAccessLayerSQL) AuthenticateUser(username string, password string) (string, error) {
	// alternativa mais segura
	// buscar o hash de password do banco e fazer a validacao na aplicacao.
	// SELECT password_hash from users where username = $1
	// poderiamos utilizar uma biblioteca como bcrypt para validar
	// a senha na aplicacao sem transmitir pela rede.
	var authenticated bool
	err := d.db.Get(&authenticated, `
		SELECT (password = crypt($1, password)) FROM users where username = $2
	`, password, username)
	if err != nil {
		return "", fmt.Errorf("Usuário ou senha incorretos.")
	}

	if !authenticated {
		return "", fmt.Errorf("Usuário ou senha incorretos.")
	}

	// como utilizar o REDIS para armazenar as sessoes
	sessionID := uuid.NewString()
	redisKey := "session:" + sessionID
	ttl := 2 * time.Minute // time to live
	err = d.redisClient.SetEx(context.Background(), redisKey, username, ttl).Err()
	if err != nil {
		return "", err
	}

	return sessionID, nil
}
