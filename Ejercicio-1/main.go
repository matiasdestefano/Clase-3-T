package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func main() {
	user := usuario{"Matias", "De Stefano", 24, "mdestefano@mail.com", "abc123"}
	fmt.Println(user)
	fmt.Println("----------------")
	user.cambiarNombre("Mati", "DS")
	user.cambiarEdad(25)
	user.cambiarContrasena("123ABC")
	user.cambiarCorreo("m.destefano@mercadolibre.com")
	fmt.Println(user)
}

func (user *usuario) cambiarNombre(nuevoNombre, nuevoApellido string) {
	user.Nombre = nuevoNombre
	user.Apellido = nuevoApellido
}

func (user *usuario) cambiarEdad(nuevaEdad int) {
	user.Edad = nuevaEdad
}

func (user *usuario) cambiarContrasena(nuevaContrasena string) {
	user.Contrasena = nuevaContrasena
}

func (user *usuario) cambiarCorreo(nuevoCorreo string) {
	user.Correo = nuevoCorreo
}
