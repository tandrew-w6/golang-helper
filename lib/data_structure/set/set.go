package set

// Set data structure implementation in Go
type Set struct {
	container map[interface{}]struct{}
}

// NewSet creates a new Set
func NewSet(keys ...interface{}) *Set {
	s := &Set{
		container: make(map[interface{}]struct{}),
	}
	s.Insert(keys...)
	return s
}

// Insert keys into the Set
func (s *Set) Insert(keys ...interface{}) {
	for _, key := range keys {
		s.container[key] = struct{}{}
	}
}

// Erase will remove keys from a Set
func (s *Set) Erase(keys ...interface{}) {
	for _, key := range keys {
		delete(s.container, key)
	}
}

// IsContain will check wheter a key is present within a Set
func (s *Set) IsContain(key interface{}) bool {
	_, ok := s.container[key]
	return ok
}

// IsEmpty will examine wheter a Set is empty or not
func (s *Set) IsEmpty() bool {
	return len(s.container) == 0
}

// Clear will empty the Set
func (s *Set) Clear() {
	s.container = make(map[interface{}]struct{})
}

// Length returns the set length
func (s *Set) Length() int {
	return len(s.container)
}

// ToSlice converts a set into a slice
func (s *Set) ToSlice() []interface{} {
	var (
		ret = make([]interface{}, 0)
	)
	for key := range s.container {
		ret = append(ret, key)
	}
	return ret
}
