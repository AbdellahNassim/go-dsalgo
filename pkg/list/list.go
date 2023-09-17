package list

type List interface {
	Len() int
	Append(value interface{})
	Remove(index int) error
	Insert(index int, value interface{}) error
	Get(index int) interface{}
	Empty() bool
	Clear()
	Values() []interface{}
	Shift() interface{}
	Pop() interface{}
	Unshift(value interface{})
}

// SliceBasedList is a slice-based implementation of the List interface
type SliceBasedList struct {
	data []interface{}
}

// NewSliceBasedList creates a new SliceBasedList

func NewSliceBasedList() *SliceBasedList {
	return &SliceBasedList{}
}

// Len returns the length of the list
func (l *SliceBasedList) Len() int {
	return len(l.data)
}

// Append adds a value to the end of the list
func (l *SliceBasedList) Append(value interface{}) {
	l.data = append(l.data, value)
}

// Remove removes a value from the list by index

func (l *SliceBasedList) Remove(index int) {
	if index < 0 || index >= len(l.data) {
		panic("Index out of bounds")
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
}

// Values (): Returns the values of the list as a slice

func (l *SliceBasedList) Values() []interface{} {
	return l.data
}

// Insert(): Inserts a value at a given index

func (l *SliceBasedList) Insert(index int, value interface{}) {
	if index < 0 || index >= len(l.data) {
		panic("Index out of bounds")
	}
	l.data = append(l.data[:index], append([]interface{}{value}, l.data[index:]...)...)
}

// Get(): Returns the value at a given index

func (l *SliceBasedList) Get(index int) interface{} {
	if index < 0 || index >= len(l.data) {
		panic("Index out of bounds")
	}
	return l.data[index]
}

// Empty(): Returns true if the list is empty

func (l *SliceBasedList) Empty() bool {
	return len(l.data) == 0
}

// Clear(): Clears the list

func (l *SliceBasedList) Clear() {
	l.data = []interface{}{}
}

// Shift(): Removes the first element from the list and returns it

func (l *SliceBasedList) Shift() interface{} {
	if len(l.data) == 0 {
		panic("Index out of bounds")
	}
	value := l.data[0]
	l.data = l.data[1:]
	return value
}

// Pop(): Removes the last element from the list and returns it

func (l *SliceBasedList) Pop() interface{} {
	if len(l.data) == 0 {
		panic("Index out of bounds")
	}
	value := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-1]
	return value
}

// Unshift(): Adds a value to the beginning of the list

func (l *SliceBasedList) Unshift(value interface{}) {
	l.data = append([]interface{}{value}, l.data...)

}
