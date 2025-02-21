package algorithms

// ReverseList inverte uma lista encadeada
func ReverseList(head *ListNode) *ListNode {
	// Nova lista invertida (será o next com valor ponteiro nil do novo fim da lista)
	var newList *ListNode

	for head != nil {
		// Armazena o próximo nó
		nextNode := head.Next

		// Redireciona o ponteiro next do nó atual
		head.Next = newList

		// Atualiza newList para o nó atual
		newList = head

		// Move head para o próximo nó
		head = nextNode
	}

	return newList
}
