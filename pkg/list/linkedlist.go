package list

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Gets the length of the linked list
func (l *LinkedList) Len() int {
	return l.length
}

// Appends a new value to the linked list
func (l *LinkedList) Append(value interface{}) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	l.length++
}

// Removes a value from the linked list by index

func (l *LinkedList) Remove(index int) {
	if index < 0 || index >= l.length {
		panic("Index out of bounds")
	}
	if index == 0 {
		l.head = l.head.next
	} else {
		node := l.head
		// Iterate to the node before the one we want to remove
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		node.next = node.next.next
	}
	l.length--
}

// Values (): Returns the values of the linked list as a slice

func (l *LinkedList) Values() []interface{} {
	values := make([]interface{}, l.length)
	node := l.head
	for i := 0; i < l.length; i++ {
		values[i] = node.value
		node = node.next
	}
	return values
}

// Insert(): Inserts a value at a given index

func (l *LinkedList) Insert(index int, value interface{}) {
	if index < 0 || index >= l.length {
		panic("Index out of bounds")
	}
	node := &Node{value: value}
	if index == 0 {
		node.next = l.head
		l.head = node
	} else {
		// Iterate to the node before the one we want to insert
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		node.next = prev.next
		prev.next = node
	}
	l.length++
}

// Get(): Returns the value at a given index
func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.length {
		panic("Index out of bounds")
	}
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value
}

// Empty(): Returns true if the linked list is empty
func (l *LinkedList) Empty() bool {
	return l.length == 0
}

// Clear(): Clears the linked list
func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Shift(): Removes the first element from the linked list and returns it
func (l *LinkedList) Shift() interface{} {
	if l.length == 0 {
		panic("Index out of bounds")
	}
	value := l.head.value
	l.head = l.head.next
	l.length--
	return value
}

// Pop(): Removes the last element from the linked list and returns it
func (l *LinkedList) Pop() interface{} {
	if l.length == 0 {
		panic("Index out of bounds")
	}
	value := l.tail.value
	node := l.head
	// Iterate to the node before the tail
	for i := 0; i < l.length-2; i++ {
		node = node.next
	}
	l.tail = node
	l.tail.next = nil
	l.length--
	return value
}

// Unshift(): Adds a new element to the beginning of the linked list

func (l *LinkedList) Unshift(value interface{}) {
	node := &Node{value: value}
	node.next = l.head
	l.head = node
	if l.length == 0 {
		l.tail = node
	}
	l.length++
}