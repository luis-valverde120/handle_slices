package main

import (
	"fmt"
)

/*
binarySearch es una funcion que busca un valor en un arreglo
ordenado usando el algoritmo de busqueda binaria
*/
func binarySearch(arr []int) int {
	var target int
	fmt.Print("Ingrese el valor a buscar: ")
	fmt.Scanf("%d", &target)
	target = int(target)

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

/*
linearSearch es una funcion que realiza la busqueda de un valor en la lista
mediante una busqueda lineal
*/
func linearSearch(arr []int) int {
	var target int
	fmt.Print("Ingrese el valor a buscar: ")
	fmt.Scanf("%d", &target)
	target = int(target)

	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
