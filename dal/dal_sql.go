package dal

import "github.com/jmoiron/sqlx"

type DataAccessLayerSQL struct {
	db *sqlx.DB
}

func NewDataAccessLayerSQL(db *sqlx.DB) *DataAccessLayerSQL {
	return &DataAccessLayerSQL{
		db: db,
	}
}

func (d DataAccessLayerSQL) CreateTask(req CreateTaskRequest) (Task, error) {
	// d.db.Exec("UPDATE/INSERT")
	// return d.ReadTask(...)
	panic("not implemented")
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
	panic("not implemented")
}

func (d DataAccessLayerSQL) UpdateTask(taskID string, req UpdateTaskRequest) (Task, error) {
	// d.db.Exec("UPDATE/INSERT")
	panic("not implemented")
}

func (d DataAccessLayerSQL) DeleteTask(taskID string) error {
	// d.db.Exec("DELETE")
	panic("not implemented")
}

func (d DataAccessLayerSQL) ListAllTasks(req ListTaskRequest) ([]Task, error) {

	// usar select em lista d.db.Select()
	// var tasks []Task
	// d.db.Select(&tasks)
	panic("not implemented")
}
