package modelos

type MovimientoInventario struct {
	Id              int
	ProductoId      int
	FechaMovimiento string
	TipoMovimiento  string
	Cantidad        int
	Comentario      string
}

type MovimientosInventario []MovimientoInventario
