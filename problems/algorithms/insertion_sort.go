package algorithms

// InsertionSort ordena um slice de inteiros usando o algoritmo Insertion Sort
func InsertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		// Elemento atual que será inserido na posição correta
		chave := arr[i]

		// Índice para comparação com elementos anteriores
		j := i - 1

		// Move elementos maiores que o elemento atual uma posição à frente, para quando
		// encontrar um elemento menor que a chave e adiciona a chave na posição seguinte
		for j >= 0 && arr[j] > chave {
			arr[j+1] = arr[j]
			j--
		}

		// Insere o elemento atual na posição correta, seguinte ao elemento maior ou na posição inicial
		arr[j+1] = chave
	}
}
