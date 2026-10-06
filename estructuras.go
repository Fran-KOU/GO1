// Estudiante: Francis Acuña

package main

import (
	"fmt"
)

func main() {
	ejercicioArrayPrimos()
	arrayPromedio()
	slicePrimos()
	votacionMapa()
}

// Resuelva los 4 ejercicios en GoLand y suba al Canvas el archivo estructuras.go:

/* Ejercicio 1 · Array
Los números primos son la base del cifrado RSA. Guarde 2, 3, 5, 7 y 11 en un array y muestre el primero, el del centro, el último y cuántos hay.
Use: [5]int, índices y len. */

func ejercicioArrayPrimos() {
	primos := [5]int {2, 3, 5, 7, 11}
	fmt.Println("El primer numero es:", primos[0])
	fmt.Println("El numero de en medio es:",primos[2])
	fmt.Println("El último número es:", primos[4])
}

/* Ejercicio 2 · Array
Un sensor tomó tres muestras: 71.8, 56.2 y 89.5. Guárdelas en un array y muestre su promedio con 2 decimales.
Use: [3]float64, for…range, len y Printf.*/

func arrayPromedio() {
	muestras := [3]float64 {71.8, 56.2, 89.5}
	suma := 0.0

	for _, valor := range muestras {
		suma += valor
	}

	promedio := suma / float64(len(muestras))
	fmt.Printf("El promedio es igual a:  %.2f\n", promedio)
}

/* Ejercicio 3 · Slice
Cree un slice con los primos 2, 3 y 5. Agréguele el 7 y el 11. Muestre el slice y el corte que va del índice 1 al índice 3.
Use: []int, append y [ini:fin].*/

func slicePrimos () {
	primos := []int {2, 3, 5}
	primos = append(primos, 7, 11)
	fmt.Println("Primos", primos, "Corte:", primos[1:4])
}

/* Ejercicio 4 · Map
Un sistema de votación recibió estos votos: Amber, Brian, Amber, Brian, Amber. Cuente con un map cuántos votos tiene cada candidato y muestre el resultado.
Use: map[string]int, make, for…range y ++. */

func votacionMapa () {
	votos := []string{"Amber", "Brian", "Amber", "Brian", "Amber"}
	conteo := make(map[string]int)

	for _, v := range votos {
		conteo[v] ++
	}

	fmt.Println("Votos", conteo)
}