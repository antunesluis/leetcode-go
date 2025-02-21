package easy

func middleNode(head *ListNode) *ListNode {
	ahead := head // Ponteiro rápido

	// Percorre a lista com os dois ponteiros
	for ahead != nil && ahead.Next != nil {
		ahead = ahead.Next.Next // Avança dois nós
		head = head.Next        // Avança um nó
	}

	// Quando o loop termina, head está no nó do meio
	return head
}
