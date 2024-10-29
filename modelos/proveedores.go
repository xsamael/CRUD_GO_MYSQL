package modelos

type Proveedor struct {
	Id        int
	Nombre    string
	Contacto  string
	Telefono  string
	Email     string
	Direccion string
}

type Proveedores []Proveedor
