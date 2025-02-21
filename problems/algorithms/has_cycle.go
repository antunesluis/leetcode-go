package algorithms

// HasCycle verifica se a lista encadeada possui um ciclo
func HasCycle(head *ListNode) bool {
	slow := head // Ponteiro lento
	fast := head // Ponteiro rápido

	// Percorre a lista com os dois ponteiros
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // Avança um nó
		fast = fast.Next.Next // Avança dois nós

		// Se os ponteiros se encontrarem, há um ciclo
		if slow == fast {
			return true
		}
	}

	// Se o loop terminar, não há ciclo
	return false
}
