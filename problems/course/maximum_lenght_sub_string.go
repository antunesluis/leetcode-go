package course

// Implementação exemplo para exercicios de Sliding Window
func maximumLengthSubstring(s string) int {
	l, r := 0, 0
	_max := 1
	counter := make(map[byte]int)

	// Inicializamos o primeiro caractere
	counter[s[0]] = 1

	for r < len(s)-1 {
		r++
		// Incrementamos o contador para o caractere atual
		counter[s[r]]++

		// Enquanto tivermos um caractere com frequência 3 movemos a janela da esquerda
		for counter[s[r]] == 3 {
			counter[s[l]]--
			l++
		}

		// Atualizamos o máximo
		if r-l+1 > _max {
			_max = r - l + 1
		}
	}

	return _max
}
