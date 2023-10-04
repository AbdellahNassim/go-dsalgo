package set

import "fmt"

type HashSet struct {
	data map[interface{}]struct{}
}

var itemExists = struct{}{}

// New creates a new set.
func NewHashSet(values ...interface{}) *HashSet {
	set := &HashSet{
		data: make(map[interface{}]struct{}),
	}
	if (len(values) > 0) {
		set.Add(values...)
	}
	return set
}

// Add adds elements to the set.
func (s *HashSet) Add(elements ...interface{}) {
	for _, element := range elements {
		s.data[element] = itemExists
	}
}
// Remove removes elements from the set.
func (s *HashSet) Remove(elements ...interface{}) {
	for _, element := range elements {
		delete(s.data, element)
	}
}
// Contains checks if the set contains the elements.
func (s *HashSet) Contains(elements ...interface{}) bool {
	for _, element := range elements {
		if _, ok := s.data[element]; !ok {
			return false
		}
	}
	return true
}
// Len returns the number of elements in the set.
func (s *HashSet) Len() int {
	return len(s.data)
}
// Clear removes all elements from the set.
func (s *HashSet) Clear() {
	s.data = make(map[interface{}]struct{})
}
// Values returns all elements in the set.
func (s *HashSet) Values() []interface{} {
	values := make([]interface{}, s.Len())
	i := 0
	for k := range s.data {
		values[i] = k
		i++
	}
	return values
}
// Empty checks if the set is empty.
func (s *HashSet) Empty() bool {
	return s.Len() == 0
}
// String returns a string representation of the set.
func (s *HashSet) String() string {
	return fmt.Sprintf("%v", s.Values())
}
// Equal checks if the set is equal to another set.
func (s *HashSet) Equal(other *HashSet) bool {
	if s.Len() != other.Len() {
		return false
	}
	for _, v := range s.Values() {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}
// Intersect returns a new set with elements that are in both sets.
func (s *HashSet) Intersect(other *HashSet) *HashSet {
	intersection := NewHashSet()
	for _, v := range s.Values() {
		if other.Contains(v) {
			intersection.Add(v)
		}
	}
	return intersection
}
// Union returns a new set with elements from both sets.
func (s *HashSet) Union(other *HashSet) *HashSet {
	union := NewHashSet()
	for _, v := range s.Values() {
		union.Add(v)
	}
	for _, v := range other.Values() {
		union.Add(v)
	}
	return union
}
// Difference returns a new set with elements in the set that are not in the other set.
func (s *HashSet) Difference(other *HashSet) *HashSet {
	difference := NewHashSet()
	for _, v := range s.Values() {
		if !other.Contains(v) {
			difference.Add(v)
		}
	}
	return difference
}