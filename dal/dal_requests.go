package dal

type CreateTaskRequest struct {
	Subject     string
	Description string
}

type UpdateTaskRequest struct {
	Subject     string
	Description string
}

type ListTaskRequest struct {
}
