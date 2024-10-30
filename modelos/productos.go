package modelos

type Producto struct {
	Id            int64
	Nombre        string
	CategoriaId   int
	Precio        float64
	ProveedorId   int
	Descripcion   string
	FechaCreacion string
}

type Productos []Producto
