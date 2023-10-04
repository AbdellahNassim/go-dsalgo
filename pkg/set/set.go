package set

type Set interface {
	// Add adds elements to the set.
	Add(elements ...interface{})
	// Remove removes elements from the set.
	Remove(elements ...interface{})
	// Contains checks if the set contains the elements.
	Contains(elements ...interface{}) bool
	// Len returns the number of elements in the set.
	Len() int
	// Clear removes all elements from the set.
	Clear()
	// Values returns all elements in the set.
	Values() []interface{}
	// Empty checks if the set is empty.
	Empty() bool
	// String returns a string representation of the set.
	String() string
	// Equal checks if the set is equal to another set.
	Equal(other *Set) bool
	// Intersect returns a new set with elements that are in both sets.
	Intersect(other *Set) *Set
	// Union returns a new set with elements from both sets.
	Union(other *Set) *Set
	// Difference returns a new set with elements in the set that are not in the other set.
	Difference(other *Set) *Set
}