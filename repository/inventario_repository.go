package repository

import (
	"crud/databases"
	"crud/modelos"
	"fmt"
	"log"
)

func ListarInventario() {
	databases.Conn()
	q := "SELECT * FROM Inventario;"

	datos, err := databases.Db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	for datos.Next() {
		var dato modelos.Inventario

		err := datos.Scan(&dato.Id, &dato.ProductoId, &dato.Cantidad, &dato.Ubicacion)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Identificador: %v | Producto Id: %v | Cantidad: %v | Ubicacion: %v",
			dato.Id, dato.ProductoId, dato.Cantidad, dato.Ubicacion)
	}
}
