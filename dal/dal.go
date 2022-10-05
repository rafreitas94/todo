package dal

// NewDataAccessLayer constróis uma nova camada de banco de dados
func NewDataAccessLayer() DataAcessLayerInterface {
	return DataAccessLayerInMemory{
		taskMap: map[string]Task{},
	}
}
