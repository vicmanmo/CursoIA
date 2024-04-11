package main

import "fmt"

func main() {
	// Imprime "¡Hola Mundo!" en la consola
	fmt.Println("¡Hola Mundo!")

	// Imprime números del 1 al 10 en la consola
	// Se inicializa i en 1, se repite mientras i sea menor o igual a 10, se incrementa i en cada iteración
	for i := 1; i <= 10; i++ {
		fmt.Println(i) // Imprime el valor actual de i
	}

	// Lista de grados en Celsius
	listaGrados := []float64{20.5, 25.3, 30.0, 15.7, 28.9}

	// Llamar a la función para mostrar los grados
	mostrarGradosFahrenheit(listaGrados)
}

// Función para convertir grados Celsius a Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// Función para mostrar los grados en Fahrenheit de una lista
func mostrarGradosFahrenheit(grados []float64) {
	// Iterar sobre la lista de grados
	for i, grado := range grados {
		// Convertir el grado de Celsius a Fahrenheit
		fahrenheit := celsiusToFahrenheit(grado)
		// Mostrar el grado en Fahrenheit y su posición en la lista
		fmt.Printf("Grado %d: %.2f°C = %.2f°F\n", i+1, grado, fahrenheit)
	}
}
