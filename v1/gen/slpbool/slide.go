package main

// Cut removes a piece of the array
func Cut(s []*bool, i, j int) []*bool {
	return append(s[:i], s[j:]...)
}

// Copy copies the provided array in a new one with the same capacity
func Copy(s []*bool) []*bool {
	b := make([]*bool, len(s), cap(s))
	copy(b, s)
	return b
}

// CopyFit copies the provided array in a new one with capacity set as the length of the input one
func CopyFit(s []*bool) []*bool {
	b := make([]*bool, len(s), len(s))
	copy(b, s)
	return b
}

// Delete removes an entry from the array
func Delete(s []*bool, i int) []*bool {
	return append(s[:i], s[i+1:]...)
}

// DeleteWithoutOrder removes an entry from the array
func DeleteWithoutOrder(s []*bool, i int) []*bool {
	lp := len(s) - 1
	s[i] = s[lp]
	return s[:lp]
}

// FilterInPlace filters the provided array using the the filter function
func FilterInPlace(s []*bool, filter func(int, *bool) bool) []*bool {
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
func Filter(s []*bool, filter func(int, *bool) bool) []*bool {
	scp := CopyFit(s)
	return FilterInPlace(scp, filter)
}

// PopLast pops the last element from the provided array
func PopLast(s []*bool) (*bool, []*bool) {
	ep := len(s) - 1
	return s[ep], s[:ep]
}

// PopFirst pops the first element from the provided array
func PopFirst(s []*bool) (*bool, []*bool) {
	return s[0], s[1:]
}

// InsertAtPos inserts the i slide into the s one at the pos position
func InsertAtPos(s []*bool, pos int, i ...*bool) []*bool {
	return append(append(s[pos:], i...), s[:pos]...)
}

// PushFront adds a new element in front of the slice
func PushFront(s []*bool, e ...*bool) []*bool {
	return append(e, s...)
}

// Extend adds new entries (with default values) at the provided slice at given position
func Extend(s []*bool, pos, size int) []*bool {
	return append(s[:pos], append(make([]*bool, size), s[pos:]...)...)
}

// ExtendEnd adds new entries (with default values) at the provided slice at the end
func ExtendEnd(s []*bool, size int) []*bool {
	return append(s, make([]*bool, size)...)
}
