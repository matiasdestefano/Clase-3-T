package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canalInsercion := make(chan float64)
	canalBurbujeo := make(chan float64)
	canalSeleccion := make(chan float64)

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	go ordenamientoInsercion(variable1, canalInsercion)
	go ordenamientoBurbujeo(variable1, canalBurbujeo)
	go ordenamientoSeleccion(variable1, canalSeleccion)

	fmt.Printf("Para 100 numeros aleatorios: \n")
	fmt.Printf("Burbujeo: %f\n", <-canalBurbujeo)
	fmt.Printf("Seleccion: %f\n", <-canalSeleccion)
	fmt.Printf("Insercion: %f\n\n", <-canalInsercion)

	go ordenamientoInsercion(variable2, canalInsercion)
	go ordenamientoBurbujeo(variable2, canalBurbujeo)
	go ordenamientoSeleccion(variable2, canalSeleccion)

	fmt.Printf("Para 1000 numeros aleatorios: \n")
	fmt.Printf("Burbujeo: %f\n", <-canalBurbujeo)
	fmt.Printf("Seleccion: %f\n", <-canalSeleccion)
	fmt.Printf("Insercion: %f\n\n", <-canalInsercion)

	go ordenamientoInsercion(variable3, canalInsercion)
	go ordenamientoBurbujeo(variable3, canalBurbujeo)
	go ordenamientoSeleccion(variable3, canalSeleccion)

	fmt.Printf("Para 10000 numeros aleatorios: \n")
	fmt.Printf("Burbujeo: %f\n", <-canalBurbujeo)
	fmt.Printf("Seleccion: %f\n", <-canalSeleccion)
	fmt.Printf("Insercion: %f\n", <-canalInsercion)
}

func ordenamientoInsercion(lista []int, canal chan float64) {
	start := time.Now()
	var n = len(lista)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if lista[j-1] > lista[j] {
				lista[j-1], lista[j] = lista[j], lista[j-1]
			}
			j = j - 1
		}
	}
	canal <- time.Since(start).Seconds()
}

func ordenamientoBurbujeo(lista []int, canal chan float64) {
	start := time.Now()
	for i := len(lista); i > 0; i-- {
		for j := 1; j < i; j++ {
			if lista[j-1] > lista[j] {
				intermedio := lista[j]
				lista[j] = lista[j-1]
				lista[j-1] = intermedio
			}
		}
	}
	canal <- time.Since(start).Seconds()
}

func ordenamientoSeleccion(lista []int, canal chan float64) {
	start := time.Now()
	var n = len(lista)
	for i := 0; i < n; i++ {
		var indiceMinimo = i
		for j := i; j < n; j++ {
			if lista[j] < lista[indiceMinimo] {
				indiceMinimo = j
			}
		}
		lista[i], lista[indiceMinimo] = lista[indiceMinimo], lista[i]
	}
	canal <- time.Since(start).Seconds()
}
