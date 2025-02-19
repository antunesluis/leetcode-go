package course

// Busca binária corrigida

func binarySearch(arr []int, target, lo, hi int) int {
	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// Busca exponencial corrigida
func exponentialSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	// Verifica o primeiro elemento
	if arr[0] == target {
		return 0
	}

	n := len(arr)
	i := 1

	// Encontra o intervalo potencial
	for i < n && arr[i] < target {
		i *= 2
	}

	// Define os limites para a busca binária
	low := i / 2
	high := i
	if high >= n {
		high = n - 1
	}

	return binarySearch(arr, target, low, high)
}
