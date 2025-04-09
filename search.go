package main

import (
	"fmt"
)

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
