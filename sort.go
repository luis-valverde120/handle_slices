package main

type sortFunc func(a, b int) bool

func merge(L, R []int, fn sortFunc) []int {
	A := make([]int, len(L)+len(R))
	i, j, k := 0, 0, 0

	// comparar elementos de derecha e izquierda antes de ordenar
	for i < len(L) && j < len(R) {
		if fn(L[i], R[j]) {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
		k++
	}

	// verificar si alguno de los elementos left/rigth
	// se requeirene que alguna seccion
	for i < len(L) {
		A[k] = L[i]
		i++
		k++
	}
	for j < len(R) {
		A[k] = R[j]
		j++
		k++
	}

	return A
}

func sort(A []int, fn sortFunc) []int {
	if fn == nil {
		fn = func(a, b int) bool {
			return a < b
		}
	}

	if len(A) > 1 {
		m := len(A) / 2
		L := A[:m]
		R := A[m:]

		A = merge(sort(L, fn), sort(R, fn), fn)
	}
	return A
}
