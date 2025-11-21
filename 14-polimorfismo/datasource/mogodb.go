package datasource

import "fmt"

var namesMongoDB []string

type MongoDB struct {
	host         string
	organization string
}

func NewMongoDB(host, organization string) MongoDB {
	return MongoDB{
		host:         host,
		organization: organization,
	}
}

// funções que alteram algum dado é legal passa com * para uso de ponteiros
func (m *MongoDB) Save(name string) {
	namesMongoDB = append(namesMongoDB, name)
	fmt.Println("dado persistido no mongodb")
}

func (m MongoDB) GetAll() []string {
	return namesMongoDB
}
