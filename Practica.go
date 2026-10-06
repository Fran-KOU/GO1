package main

import "fmt"

func main() {
	array()
	servidores()
	latencias()
	rango()
	copia()
}

func array() {
	latencias := [5]int{12, 32, 20, 240, 9}
	fmt.Println(latencias[2])
	latencias[2] = 12
	fmt.Println(latencias[2])
	fmt.Println(latencias[len(latencias)-1])
}

func servidores() {
	servidores := [3]string{"web01", "web02", "db01"}
	fmt.Println((servidores))
}

func latencias() {
	latencia := [5]int{120, 25, 240, 95, 60}

	for i := 0; i < len(latencia); i++ {
		fmt.Println(i, latencia[i])
	}
}

func rango() {
	latenciaR := [5]int{120, 85, 240, 95, 60}
	for i, ms := range latenciaR {
		fmt.Println(i, ms)
	}
}

func copia() {
	latenciaP := [5]int{120, 85, 240, 95, 60}
	copia := latenciaP
	copia[0] = 999
	fmt.Println(latenciaP)
	fmt.Println(copia)
}