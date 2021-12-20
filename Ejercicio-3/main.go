package main

import "fmt"

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type mantenimiento struct {
	nombre string
	precio float64
}

func main() {
	productos := []producto{{"Silla", 400.0, 2}, {"Televisor", 30000.0, 5}}
	servicios := []servicio{{"Electricidad", 2000, 6400}, {"Reparacion", 4800, 5000}}
	mantenimientos := []mantenimiento{{"Mantenimiento A", 7000.0}}

	canalProductos := make(chan float64)
	canalServicios := make(chan float64)
	canalMantenimientos := make(chan float64)

	go sumarProductos(productos, canalProductos)
	go sumarServicios(servicios, canalServicios)
	go sumarMantenimiento(mantenimientos, canalMantenimientos)

	totalProductos := <-canalProductos
	totalServicios := <-canalServicios
	totalMantenimientos := <-canalMantenimientos

	fmt.Printf("Productos: $%.2f\nServicios: $%.2f\nMantenimientos: $%.2f", totalProductos, totalServicios, totalMantenimientos)
}

func sumarProductos(listaProductos []producto, canal chan float64) {
	var total float64
	for _, prod := range listaProductos {
		total += prod.Precio * float64(prod.Cantidad)
	}

	canal <- total
}

func sumarServicios(listaServicios []servicio, canal chan float64) {
	var total float64
	for _, servicio := range listaServicios {
		cantidadDeMediasHoras := servicio.MinutosTrabajados / 30
		total += float64(cantidadDeMediasHoras) * servicio.Precio
	}
	canal <- total
}

func sumarMantenimiento(listaMantenimiento []mantenimiento, canal chan float64) {
	var total float64
	for _, mantenimiento := range listaMantenimiento {
		total += mantenimiento.precio
	}
	canal <- total
}
