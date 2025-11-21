package datasource

type Datasource interface {
	Save(name string)
	GetAll() []string
}
