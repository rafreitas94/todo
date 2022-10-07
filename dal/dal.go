package dal

// NewDataAccessLayer constroi uma nova camada de acesso de dados
func NewDataAccessLayer() DataAccessLayerInterface {
	return DataAccessLayerInMemory{
		tasksMap: map[string]Task{},
	}
}
