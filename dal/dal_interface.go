package dal

// dal -> data acess layer - camada de acesso a dados
type DataAcessLayerInterface interface {

	// CreateTask cria uma nova tarefa a partir
	// do parametro req, retorna a tarefa criada
	// e um erro
	CreateTask(req CreateTaskRequest) (Task, error)

	ReadTask(taskId string) (Task, error)

	UpdateTask(req UpdateTaskRequest) (Task, error)

	DeleteTask(taskId string) error

	ListAllTask(req ListTaskRequest) ([]Task, error)
}
