package main

import (
	"fmt"
	"red-social-nicho/src/interfaz"
)

func main() {
	fmt.Println("🌐 Red Social de Nicho - Sistema Completo")
	fmt.Println("========================================")
	fmt.Println()
	
	fmt.Println("Seleccione una opción:")
	fmt.Println("1. Ejecutar interfaz CLI interactiva")
	fmt.Println("2. Ejecutar casos de prueba")
	fmt.Println("3. Salir")
	fmt.Print("Opción: ")
	
	var opcion int
	fmt.Scanf("%d", &opcion)
	
	switch opcion {
	case 1:
		cli := interfaz.NuevaCLI()
		cli.Ejecutar()
	case 2:
		ejecutarCasosPruebaSimples()
	case 3:
		fmt.Println("¡Hasta luego!")
	default:
		fmt.Println("Opción no válida")
	}
}