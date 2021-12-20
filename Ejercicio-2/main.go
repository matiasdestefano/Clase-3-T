package main

import "fmt"

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []producto
}

func main() {
	var usuario = usuario{Nombre: "Matias", Apellido: "De Stefano"}
	nuevoProd := nuevoProducto("Silla PC", 400000)
	agregarProducto(&usuario, nuevoProd, 1)
	fmt.Println(usuario)
	fmt.Println("---------")
	borrarProducto(&usuario)
	fmt.Println(usuario)
}

func nuevoProducto(nombre string, precio float64) producto {
	prod := producto{Nombre: nombre, Precio: precio}
	return prod
}

func agregarProducto(user *usuario, prod producto, cantidad int) {
	prod.Cantidad = cantidad
	user.Productos = append(user.Productos, prod)
}

func borrarProducto(user *usuario) {
	user.Productos = []producto{}
}
