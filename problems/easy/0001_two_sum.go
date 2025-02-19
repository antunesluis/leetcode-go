package easy

func twoSum(nums []int, target int) []int {
	visitados := make(map[int]int)

	for i, num := range nums {
		complemento := target - num
		// Verifica primeiro se o complemento existe no map
		if j, existe := visitados[complemento]; existe && j != i {
			return []int{j, i}
		}
		visitados[num] = i
	}
	return []int{}
}
