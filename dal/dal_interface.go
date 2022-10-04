package dal

// dal -> data access layer - camada de acesso a dados
type DataAccessLayerInterface interface {

	// CreateTask cria uma nova tarefa a partir
	// do parametro req, returna a tarefa criada
	// e um erro.
	CreateTask(req CreateTaskRequest) (Task, error)

	ReadTask(taskID string) (Task, error)

	UpdateTask(req UpdateTaskRequest) (Task, error)

	DeleteTask(taskID string) error

	ListAllTasks(req ListTaskRequest) ([]Task, error)
}
