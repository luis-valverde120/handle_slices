package main

import (
	"fmt"
)

var MessagesMenu = []string{
	"1. Ingresar un nuevo valor",
	"2. Eliminar un valor",
	"3. Búsqueda binaria",
	"4. Búsqueda lineal",
	"5. Ordenar arreglo",
	"6. Imprimir arreglo",
	"7. Salir",
}

/*
insertValue es una funcion que por medio del puntero al array
agrega un valor a este
*/
func insertValue(arr *[]int) {
	var value int
	fmt.Print("Ingrese un valor: ")
	fmt.Scanln(&value)
	*arr = append(*arr, int(value))
}

func deleteValue(arr *[]int) {
	fmt.Print("Ingrese el indice a eliminar: ")
	var index int
	fmt.Scanln(&index)
	index = int(index)

	if arr == nil {
		fmt.Println("El arreglo está vacío")
		return
	}
	if len(*arr) == 0 {
		fmt.Println("El arreglo está vacío")
		return
	}

	(*arr)[index] = (*arr)[len(*arr)-1]
}

/*
menu es una funcion que imprime el menu de opciones
y valida la entrada del usuario
*/
func menu(op *int) {

	for *op <= 0 || *op >= 8 {

		for _, element := range MessagesMenu {
			fmt.Println(element)
		}

		fmt.Print("Ingrese una opción: ")
		fmt.Scanln(op)
	}
}

func main() {
	fmt.Println("Manejo de Arreglo unidimensional!")
	var op int = 0
	var arr []int

	for {
		menu(&op)

		fmt.Println(op)

		switch op {
		case 1:
			insertValue(&arr)
			op = 0
		case 2:
			deleteValue(&arr)
			op = 0
		case 3:
			value := binarySearch(arr)
			if value == -1 {
				fmt.Println("El valor no se encuentra en el arreglo")
			} else {
				fmt.Println("El valor se encuentra en el índice:", value)
			}
			op = 0
		case 4:
			value := linearSearch(arr)
			if value == -1 {
				fmt.Println("El valor no se encuentra en el arreglo")
			} else {
				fmt.Println("El valor se encuentra en el índice:", value)
			}
			op = 0
		case 5:
			arr = sort(arr, nil)
			for i, value := range arr {
				fmt.Printf("arr[%d] = %d\n", i, value)
			}
			op = 0
		case 6:
			for i, value := range arr {
				fmt.Printf("arr[%d] = %d\n", i, value)
			}
			op = 0
		case 7:
			fmt.Println("Saliendo del programa...")
			return
		}
	}
}
