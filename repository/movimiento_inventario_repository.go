package repository

import (
	"crud/databases"
	"crud/modelos"
	"fmt"
	"log"
)

func ListarMovimientosInventario() {
	databases.Conn()
	defer databases.CerarConexion()

	q := "SELECT * FROM MovimientosInventario;"

	datos, err := databases.Db.Query(q)
	if err != nil {
		log.Fatal(err)
	}

	for datos.Next() {
		var dato modelos.MovimientoInventario
		err := datos.Scan(&dato.Id, &dato.ProductoId, &dato.FechaMovimiento, &dato.TipoMovimiento, &dato.Cantidad, &dato.Comentario)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Identificador: %v | Producto Id: %v | Fecha Movimiento: %v | Tipo Movimiento: %v | Cantidad: %v | Comentario: %v",
			dato.Id, dato.ProductoId, dato.FechaMovimiento, dato.TipoMovimiento, dato.Cantidad, dato.Comentario)
	}
}
