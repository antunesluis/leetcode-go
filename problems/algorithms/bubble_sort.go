package algorithms

// BubbleSort ordena um slice de inteiros usando o algoritmo Bubble Sort
func BubbleSort(arr []int) {
	n := len(arr)

	// Flag para otimização: se nenhuma troca for feita em uma passagem,
	// significa que o array já está ordenado
	swapped := true

	// Continue até que nenhuma troca seja necessária
	for swapped {
		swapped = false

		// Percorre o array comparando elementos adjacentes
		for i := 0; i < n-1; i++ {
			// Se o elemento atual é maior que o próximo, troca-os
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		// Otimização: após cada passagem, o último elemento já está na posição correta
		n--
	}
}
