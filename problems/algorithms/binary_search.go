package algorithms

import "fmt"

func BinarySearch(nums []int, n int) int {
	low := 0
	high := len(nums) - 1 // começa do último índice válido
	steps := 0

	for low <= high {
		steps += 1
		mid := low + (high-low)/2 // Calcula o ponto médio

		if nums[mid] == n { // Encontrou o número
			fmt.Println(steps)
			return mid
		} else if nums[mid] < n { // Número está na metade direita
			low = mid + 1
		} else { // Número está na metade esquerda
			high = mid - 1 // reduz o high para evitar loop infinito
		}
	}

	return -1
}
