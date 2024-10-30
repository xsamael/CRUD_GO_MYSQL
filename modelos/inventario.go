package modelos

type Inventario struct {
	Id         int
	ProductoId int
	Cantidad   int
	Ubicacion  string
}

type Inventarios []Inventario
