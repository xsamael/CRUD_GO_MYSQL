package repository

import (
	"crud/databases"
	"crud/modelos"
	"fmt"
	"log"
)

func ListarProducto() {
	databases.Conn()
	q := "SELECT * FROM Productos"
	datos, err := databases.Db.Query(q)
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next() {
		var dato modelos.Producto
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.CategoriaId, &dato.Precio, &dato.ProveedorId, &dato.Descripcion, &dato.FechaCreacion)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Identificador: %v | Nombre: %v | Categoria: %v | Precio: %v | Proveedor: %v | Descripcion: %v | Fecha Creacion: %v",
			dato.Id, dato.Nombre, dato.CategoriaId, dato.Precio, dato.ProveedorId, dato.Descripcion, dato.FechaCreacion)
	}
}
