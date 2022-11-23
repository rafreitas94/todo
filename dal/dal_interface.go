package dal

// dal -> data access layer - camada de acesso a dados
type DataAccessLayerInterface interface {

	// CreateTask cria uma nova tarefa a partir
	// do parametro req, returna a tarefa criada
	// e um erro.
	CreateTask(req CreateTaskRequest) (Task, error)

	ReadTask(taskID string) (Task, error)

	UpdateTask(taskID string, req UpdateTaskRequest) (Task, error)

	DeleteTask(taskID string) error

	ListAllTasks(req ListTaskRequest) ([]Task, error)

	// AuthenticateUser retorna uma string contendo o id de uma sessao criada pelo usuario
	// autenticado
	AuthenticateUser(username string, password string) (string, error)

	// AuthenticateSession retorna o username de um usuario autenticado pelo sessionID
	AuthenticateSession(sessionID string) (string, error)
}
