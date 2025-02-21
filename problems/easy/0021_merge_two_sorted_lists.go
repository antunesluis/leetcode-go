package easy

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Nó dummy para não ser necessário lidar com edge cases de primeiros elementos
	dummy := &ListNode{}
	// Ponteiro para o nó atual da nova lista
	current := dummy
	// Percorre as duas listas comparando os valores dos nós e selecionando o menor para o current
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {

			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next // Atualizamos o ponteiro da nova lista
	}

	// Anexa no final os nós restantes da lista que não foi completamente percorrida
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Retorna a nova lista ignorando o dummy
	return dummy.Next
}
