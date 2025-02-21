package algorithms

// MiddleNode encontra o nó do meio de uma lista encadeada em um único loop
func MiddleNode(head *ListNode) *ListNode {
	ahead := head

	// Enquanto ahead e ahead.Next não forem nil
	for ahead != nil && ahead.Next != nil {
		// Ponteiro rápido avança 2 posições
		ahead = ahead.Next.Next
		// Ponteiro lento avança 1 posição
		head = head.Next
	}

	return head
}
