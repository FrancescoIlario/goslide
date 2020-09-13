package main

// Cut removes a piece of the array
func Cut(s []interface{}, i, j int) []interface{} {
	return append(s[:i], s[j:]...)
}

// Copy copies the provided array in a new one with the same capacity
func Copy(s []interface{}) []interface{} {
	b := make([]interface{}, len(s), cap(s))
	copy(b, s)
	return b
}

// CopyFit copies the provided array in a new one with capacity set as the length of the input one
func CopyFit(s []interface{}) []interface{} {
	b := make([]interface{}, len(s), len(s))
	copy(b, s)
	return b
}

// Delete removes an entry from the array
func Delete(s []interface{}, i int) []interface{} {
	return append(s[:i], s[i+1:]...)
}

// DeleteWithoutOrder removes an entry from the array
func DeleteWithoutOrder(s []interface{}, i int) []interface{} {
	lp := len(s) - 1
	s[i] = s[lp]
	return s[:lp]
}

// FilterInPlace filters the provided array using the the filter function
func FilterInPlace(s []interface{}, filter func(int, interface{}) bool) []interface{} {
	c := 0
	for p, e := range s {
		if filter(p, e) {
			s[c] = e
			c++
		}
	}
	return s[:c]
}

// Filter filters the provided array using the the filter function
func Filter(s []interface{}, filter func(int, interface{}) bool) []interface{} {
	scp := CopyFit(s)
	return FilterInPlace(scp, filter)
}

// PopLast pops the last element from the provided array
func PopLast(s []interface{}) (interface{}, []interface{}) {
	ep := len(s) - 1
	return s[ep], s[:ep]
}

// PopFirst pops the first element from the provided array
func PopFirst(s []interface{}) (interface{}, []interface{}) {
	return s[0], s[1:]
}

// InsertAtPos inserts the i slide into the s one at the pos position
func InsertAtPos(s []interface{}, pos int, i ...interface{}) []interface{} {
	return append(append(s[pos:], i...), s[:pos]...)
}

// PushFront adds a new element in front of the slice
func PushFront(s []interface{}, e ...interface{}) []interface{} {
	return append(e, s...)
}

// Extend adds new entries (with default values) at the provided slice at given position
func Extend(s []interface{}, pos, size int) []interface{} {
	return append(s[:pos], append(make([]interface{}, size), s[pos:]...)...)
}

// ExtendEnd adds new entries (with default values) at the provided slice at the end
func ExtendEnd(s []interface{}, size int) []interface{} {
	return append(s, make([]interface{}, size)...)
}
