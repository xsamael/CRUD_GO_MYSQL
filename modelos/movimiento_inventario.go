package modelos

import "time"

type MovimientoInventario struct {
	Id              int
	ProductoId      int
	FechaMovimiento time.Time
	TipoMovimiento  string
	Cantidad        int
	Comentario      string
}
