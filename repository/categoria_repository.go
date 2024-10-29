package repository

import (
	"crud/databases"
	"crud/modelos"
	"fmt"
	"log"
)

func ListarCategoria() {
	databases.Conn()

	query := "SELECT * FROM Categorias;"
	datos, err := databases.Db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next() {
		var dato modelos.Categoria
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Descripcion)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Identificador: %v | Nombre: %v | Descripcion: %v", dato.Id, dato.Nombre, dato.Descripcion)
	}
}
