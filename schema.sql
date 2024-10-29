CREATE DATABASE inventario;

USE inventario;

CREATE TABLE Categorias (
    categoria_id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100) NOT NULL,
    descripcion TEXT
);

CREATE TABLE Proveedores (
    proveedor_id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(255) NOT NULL,
    contacto VARCHAR(100),
    telefono VARCHAR(15),
    email VARCHAR(100),
    direccion TEXT
);

CREATE TABLE Productos (
    producto_id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(255) NOT NULL,
    categoria_id INT,
    precio DECIMAL(10, 2) NOT NULL,
    proveedor_id INT,
    descripcion TEXT,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (categoria_id) REFERENCES Categorias(categoria_id),
    FOREIGN KEY (proveedor_id) REFERENCES Proveedores(proveedor_id)
);

CREATE TABLE Inventario (
    inventario_id INT PRIMARY KEY AUTO_INCREMENT,
    producto_id INT,
    cantidad INT DEFAULT 0,
    ubicacion VARCHAR(255),
    FOREIGN KEY (producto_id) REFERENCES Productos(producto_id)
);

CREATE TABLE MovimientosInventario (
    movimiento_id INT PRIMARY KEY AUTO_INCREMENT,
    producto_id INT,
    fecha_movimiento TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tipo_movimiento ENUM('entrada', 'salida') NOT NULL,
    cantidad INT NOT NULL,
    comentario TEXT,
    FOREIGN KEY (producto_id) REFERENCES Productos(producto_id)
);
