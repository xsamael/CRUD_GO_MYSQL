package repository

import (
	"crud/databases"
	"crud/modelos"
	"fmt"
	"log"
)

func ListarProveedores() {
	databases.Conn()
	q := "SELECT * FROM Proveedores;"
	datos, err := databases.Db.Query(q)
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next() {
		var dato modelos.Proveedor
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Contacto, &dato.Telefono, &dato.Email, &dato.Direccion)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Identificadoe: %v | Nombre: %v | Contacto: %v | Telefono:%v | Email: %v | Direccion: %v",
			dato.Id, dato.Nombre, dato.Contacto, dato.Telefono, dato.Email, dato.Direccion)
	}
}
