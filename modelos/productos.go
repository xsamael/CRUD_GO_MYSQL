package modelos

import "time"

type Producto struct {
	Id            int64
	Nombre        string
	CategoriaId   int
	Precio        float64
	Descripcion   string
	FechaCreacion time.Time
}

type Productos []Producto
