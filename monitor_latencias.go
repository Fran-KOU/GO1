package main

import (
	"fmt"
	
)

func main() {
	mostrarEncabezado()
	mostrarLatencia("web-01", 120.0)
}

func mostrarEncabezado() {
	fmt.Println("-- Reportes de Latencias --")
} // Salida Paso 1: -- Reportes de Latencias --

func mostrarLatencia(servidor string, ms float64) {
	fmt.Printf("Servidor: %s - Latencia %.2f ms/n", servidor, ms)
	fmt.Println()
}

