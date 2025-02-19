package main

import (
	"fmt"

	"github.com/seu-usuario/leetcode-go/problems/course"
)

func main() {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	testCases := []int{1, 7, 19, 10}

	for _, target := range testCases {
		result := course.BinarySearch(nums, target)
		if result != -1 {
			fmt.Printf("Número %d encontrado no índice %d\n", target, result)
		} else {
			fmt.Printf("Número %d não encontrado no array\n", target)
		}
	}
}
