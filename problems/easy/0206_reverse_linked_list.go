package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newList *ListNode

	for head != nil {
		// salvo o próximo nó para não perder a referência
		nextNode := head.Next

		// Corrijo o ponteiro next do nó atual para apontar para o novo nó
		head.Next = newList

		// Atualizo a nova lista para o nó atual
		newList = head

		// Avanço para o próximo nó
		head = nextNode
	}

	return newList
}
