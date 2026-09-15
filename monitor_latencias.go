package main

import (
	"fmt"
	"errors"
)

func main() {
	//-------------Funcion encabezado--------------
	mostrarEncabezado()

	// -----------------Función mostrar latencia--------------
	mostrarLatencia("web-01", 120.0)
	estado := clasificar(120.0, 300.0)

	// ------------------Función clasificar ms y umbral-------------------
	fmt.Println("Estado de latencia: ", estado)

	// -----------------------Esta parte está hecha con IA, no pude entender del todo las funciones variádicas.--------------------
	prom, max := resumen(120.0, 340.0, 95.0) 
	fmt.Printf("Resumen - Promedio: %.2f ms | Máximo: %.2f ms\n", prom, max)

	// ----------------------------Esta función también esta hecha con IA.-----------------------
	// Creamos una lista vacía

	muestraVacia := []float64{}


	// -----------Función ERROR-------------------
	// Llamamos a la función y atrapamos el error en la variable 'err'

	err := validar(muestraVacia)

	// Así es como Go comprueba si hubo un error: "¿es el error distinto de nulo?"

	if err != nil {
    fmt.Println("Alerta de auditoría:", err)


	// ---------------------Función Closures (Hecha por IA)-------------------

	// Creamos una instancia de nuestro contador. 

	// "nuevaAlerta" ahora ES la función interna que retornamos.

	nuevaAlerta := contadorAlertas()


	// Si la llamamos varias veces, verás cómo "recuerda" el valor anterior y lo suma

	fmt.Println("Alerta N°:", nuevaAlerta()) 

	fmt.Println("Alerta N°:", nuevaAlerta())
}

	// ---------------------------Función Strucs y metodos------------------------
	// Creamos nuestro servidor con los datos iniciales
	miServidor := Servidor{
    	Nombre:    "web-01",
    	Latencias: []float64{120.0, 340.0, 95.0},
	}

	// Usamos el método de puntero para registrar el 410.0
	miServidor.Registrar(410.0)

	// Usamos el método de valor para ver el estado con umbral 300.0
	fmt.Printf("Servidor: %s | Datos finales: %v | Estado: %s\n", 
    	miServidor.Nombre, miServidor.Latencias, miServidor.Estado(300.0))
	}





func mostrarEncabezado() {
	fmt.Println("-- Reportes de Latencias --")
	} // Salida Paso 1: -- Reportes de Latencias --

func mostrarLatencia(servidor string, ms float64) {
	fmt.Printf("Servidor: %s - Latencia %.2f ms\n", servidor, ms)
}

func clasificar (ms, umbral float64) string {
	if ms > umbral {
		return "Lenta"
	}
	return "OK"
}

// Esta función esta hecha con IA, las funciones verádicas fue lo que menos entendí de la clase, a falta de tiempo lo hice con IA.

func resumen(latencias ...float64) (promedio, maximo float64) {
    // Si no enviaron ningún dato, terminamos aquí para evitar dividir por cero
    if len(latencias) == 0 {
        return // Como usamos retorno desnudo, esto devolverá 0 y 0
    }

    var suma float64

    // En Go, usamos 'range' para recorrer listas. 
    // Nos da el índice (que ignoramos con '_') y el valor (la latencia)
    for _, latencia := range latencias {
        suma += latencia // Sumamos cada latencia
        
        // Lógica para encontrar el máximo:
        if latencia > maximo {
            maximo = latencia
        }
    }

    // Calculamos el promedio. len() nos da la cantidad de elementos.
    // Usamos float64() para convertir la cantidad (que es entera) a decimal.
    promedio = suma / float64(len(latencias))

    return // Retorno desnudo: devuelve 'promedio' y 'maximo' automáticamente
}

// Para esta función también tuve que usar IA, no entendí las listas, pero entiendo el uso de "len", se usa para saber el número
// de elementos detro de la lista.

func validar(latencias []float64) error {
    // Si la lista no tiene elementos, creamos y devolvemos un error
    if len(latencias) == 0 {
        return errors.New("rechazado: la muestra está vacía")
    }
    // Si devolvemos nil, significa que no hubo error.
    return nil 
}

// Esta función también está hecha por la IA, la doble función es confusa y no puedo entenderla en poco tiempo.

func contadorAlertas() func() int {
    // Esta variable 'conteo' queda "atrapada" o protegida aquí adentro
    conteo := 0 
    
    // Retornamos una función anónima (sin nombre)
    return func() int {
        conteo++ // Le sumamos 1 al conteo cada vez que la llamen
        return conteo
    }
}

// Función STRUC y metodos. Hecha por IA.

// 1. Definimos nuestra estructura de datos
type Servidor struct {
	Nombre    string
	Latencias []float64
}

// 2. Método con Receptor de Valor (solo lee, no modifica)
// Fíjate en la (s Servidor) antes del nombre de la función
func (s Servidor) Estado(umbral float64) string {
	if len(s.Latencias) == 0 {
		return "SIN DATOS"
	}
	// Revisamos la última latencia registrada
	ultima := s.Latencias[len(s.Latencias)-1]
	if ultima > umbral {
		return "LENTA"
	}
	return "OK"
}

// 3. Método con Receptor de Puntero (modifica los datos originales)
// Fíjate en el asterisco (*Servidor). El * es el puntero.
func (s *Servidor) Registrar(ms float64) {
    // append sirve para agregar un nuevo dato al final de la lista
	s.Latencias = append(s.Latencias, ms) 
}






// ========== Declaración de IA ===============

// Usé IA para gran maypría del trabajo, aqui esta el prompt.

// Necesito que me ayudes a realizar la actividad, no me des haciendo la actividad, 
// guíame desde lo más básico hasta poder hacer por mí mismo el código. 
// Actúa como un profesor de programación, explícame las cosas detalladamente para entender 
// la forma en la que GO actúa.

// Modelo: GEMINI PRO

// 15/09/2026 12:58