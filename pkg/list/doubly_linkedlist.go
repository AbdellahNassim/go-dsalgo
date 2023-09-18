package list

// Node is a node in a doubly linked list

type DoubleNode struct {
	value interface{}
	next  *DoubleNode
	prev  *DoubleNode
}

// DoublyLinkedList is a doubly linked list

type DoublyLinkedList struct {
	head   *DoubleNode
	tail   *DoubleNode
	length int
}

// NewDoublyLinkedList creates a new doubly linked list

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Len returns the length of the doubly linked list
func (l *DoublyLinkedList) Len() int {
	return l.length
}

// Append adds a value to the end of the doubly linked list
func (l *DoublyLinkedList) Append(value interface{}) {
	node := &DoubleNode{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
	l.length++
}

// Remove removes a value from the doubly linked list by index
func (l *DoublyLinkedList) Remove(index int) {
	if index < 0 || index >= l.length {
		panic("Index out of bounds")
	}
	if index == 0 {
		l.head = l.head.next
		l.head.prev = nil
	} else {
		node := l.head
		// Iterate to the node before the one we want to remove
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		node.next = node.next.next
		node.next.prev = node
	}
	l.length--
}

// Values returns the values of the doubly linked list as a slice
func (l *DoublyLinkedList) Values() []interface{} {
	values := make([]interface{}, l.length)
	node := l.head
	for i := 0; i < l.length; i++ {
		values[i] = node.value
		node = node.next
	}
	return values
}

// Insert inserts a value at a given index
func (l *DoublyLinkedList) Insert(index int, value interface{}) {
	if index < 0 || index >= l.length {
		panic("Index out of bounds")
	}
	node := &DoubleNode{value: value}
	if index == 0 {
		node.next = l.head
		l.head.prev = node
		l.head = node
	} else {
		// Iterate to the node before the one we want to insert
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		prev.next.prev = node
		node.next = prev.next
		node.prev = prev
		prev.next = node
	}
	l.length++
}

// Get returns the value at a given index
func (l *DoublyLinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.length {
		panic("Index out of bounds")
	}
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value
}

// Empty returns true if the doubly linked list is empty
func (l *DoublyLinkedList) Empty() bool {
	return l.length == 0
}

// Clear clears the doubly linked list

func (l *DoublyLinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Shift removes the first element from the doubly linked list and returns it
func (l *DoublyLinkedList) Shift() interface{} {
	if l.length == 0 {
		panic("Index out of bounds")
	}
	value := l.head.value
	l.head = l.head.next
	l.head.prev = nil
	l.length--
	return value	
}

// Pop removes the last element from the doubly linked list and returns it
func (l *DoublyLinkedList) Pop() interface{} {
	if l.length == 0 {
		panic("Index out of bounds")
	}
	value := l.tail.value
	l.tail = l.tail.prev
	l.tail.next = nil
	l.length--
	return value
}

// Unshift adds a value to the beginning of the doubly linked list
func (l *DoublyLinkedList) Unshift(value interface{}) {
	node := &DoubleNode{value: value}
	node.next = l.head
	l.head.prev = node
	l.head = node
	l.length++
}