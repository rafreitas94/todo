package dal

type CreateTaskRequest struct {
	Subject     string
	Description string
}

type UpdateTaskRequest struct {
	Subject     string
	Description string
	Status      string
}

type PatchTaskRequest struct {
	Subject     *string
	Description *string
	Status      *string
}

type ListTaskRequest struct {
}
