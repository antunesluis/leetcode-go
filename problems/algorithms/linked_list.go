package algorithms

// Node representa um nó na lista duplamente encadeada
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

// DoublyLinkedList representa a lista duplamente encadeada
type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// NewDoublyLinkedList cria uma nova lista duplamente encadeada vazia
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// InsertFront insere um novo elemento no início da lista
func (dll *DoublyLinkedList) InsertFront(data interface{}) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: dll.head,
	}

	if dll.head == nil {
		// Se a lista está vazia, head e tail apontam para o novo nó
		dll.tail = newNode
	} else {
		// Caso contrário, o nó anterior ao head atual será o novo nó
		dll.head.prev = newNode
	}
	dll.head = newNode
	dll.length++
}

// InsertBack insere um novo elemento no final da lista
func (dll *DoublyLinkedList) InsertBack(data interface{}) {
	newNode := &Node{
		data: data,
		prev: dll.tail,
		next: nil,
	}

	if dll.head == nil {
		// Se a lista está vazia, head e tail apontam para o novo nó
		dll.head = newNode
		dll.tail = newNode
	} else {
		// Adiciona o novo nó depois do tail e atualiza o tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
	dll.length++
}

// RemoveFront remove o primeiro elemento da lista
func (dll *DoublyLinkedList) RemoveFront() interface{} {
	if dll.head == nil {
		return nil
	}

	data := dll.head.data
	dll.head = dll.head.next
	if dll.head != nil {
		// Se a lista não ficou vazia, atualiza o ponteiro prev do novo head para nil
		dll.head.prev = nil
	} else {
		// Se a lista ficou vazia, atualiza o tail para nil
		dll.tail = nil
	}
	dll.length--
	return data
}

// RemoveBack remove o último elemento da lista
func (dll *DoublyLinkedList) RemoveBack() interface{} {
	if dll.tail == nil {
		return nil
	}

	data := dll.tail.data
	dll.tail = dll.tail.prev
	if dll.tail != nil {
		// Se a lista não ficou vazia, atualiza o ponteiro next do novo tail para nil
		dll.tail.next = nil
	} else {
		// Se a lista ficou vazia, atualiza o head para nil
		dll.head = nil
	}
	dll.length--
	return data
}

// Search procura um elemento na lista e retorna true se encontrado
func (dll *DoublyLinkedList) Search(data interface{}) bool {
	current := dll.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}
