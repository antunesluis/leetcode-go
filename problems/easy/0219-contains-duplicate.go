package easy

// Procura dois indices que possuam valores iguais e que subtraidos sejam menores ou iguais a k ([1,2,3,1], k = 3 -> true)
func containsNearbyDuplicate(nums []int, k int) bool {
	// Conjunto para armazenar os números na janela atual (bool indica presença)
	window := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		// Se a janela exceder o tamanho k, remove o elemento mais antigo (contração da janela)
		if i > k {
			delete(window, nums[i-k-1])
		}

		// Se o número atual já existe na janela, encontramos uma duplicata
		if window[nums[i]] {
			return true
		}

		// Adiciona o número atual à janela (expansão da janela)
		window[nums[i]] = true
	}

	return false
}
