package main

import (
	"fmt"

	"github.com/Paulinhoh/golang_pratico/14-polimorfismo/datasource"
)

func main() {
	// datasource := datasource.NewMongoDB("localhost", "teste-polimorfismo")
	datasource := datasource.NewPostgreSQL("localhost", "pg-polimosfismo", 5432)
	saveNewName(&datasource, "paulinho")
	saveNewName(&datasource, "biel")
	saveNewName(&datasource, "aninha")
	getNames(&datasource)
}

func saveNewName(d datasource.Datasource, name string) {
	d.Save(name)
}

func getNames(d datasource.Datasource) {
	fmt.Println(d.GetAll())
}
