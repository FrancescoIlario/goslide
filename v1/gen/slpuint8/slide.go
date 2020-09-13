package main

// Cut removes a piece of the array
func Cut(s []*uint8, i, j int) []*uint8 {
	return append(s[:i], s[j:]...)
}

// Copy copies the provided array in a new one with the same capacity
func Copy(s []*uint8) []*uint8 {
	b := make([]*uint8, len(s), cap(s))
	copy(b, s)
	return b
}

// CopyFit copies the provided array in a new one with capacity set as the length of the input one
func CopyFit(s []*uint8) []*uint8 {
	b := make([]*uint8, len(s), len(s))
	copy(b, s)
	return b
}

// Delete removes an entry from the array
func Delete(s []*uint8, i int) []*uint8 {
	return append(s[:i], s[i+1:]...)
}

// DeleteWithoutOrder removes an entry from the array
func DeleteWithoutOrder(s []*uint8, i int) []*uint8 {
	lp := len(s) - 1
	s[i] = s[lp]
	return s[:lp]
}

// FilterInPlace filters the provided array using the the filter function
func FilterInPlace(s []*uint8, filter func(int, *uint8) bool) []*uint8 {
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
func Filter(s []*uint8, filter func(int, *uint8) bool) []*uint8 {
	scp := CopyFit(s)
	return FilterInPlace(scp, filter)
}

// PopLast pops the last element from the provided array
func PopLast(s []*uint8) (*uint8, []*uint8) {
	ep := len(s) - 1
	return s[ep], s[:ep]
}

// PopFirst pops the first element from the provided array
func PopFirst(s []*uint8) (*uint8, []*uint8) {
	return s[0], s[1:]
}

// InsertAtPos inserts the i slide into the s one at the pos position
func InsertAtPos(s []*uint8, pos int, i ...*uint8) []*uint8 {
	return append(append(s[pos:], i...), s[:pos]...)
}

// PushFront adds a new element in front of the slice
func PushFront(s []*uint8, e ...*uint8) []*uint8 {
	return append(e, s...)
}

// Extend adds new entries (with default values) at the provided slice at given position
func Extend(s []*uint8, pos, size int) []*uint8 {
	return append(s[:pos], append(make([]*uint8, size), s[pos:]...)...)
}

// ExtendEnd adds new entries (with default values) at the provided slice at the end
func ExtendEnd(s []*uint8, size int) []*uint8 {
	return append(s, make([]*uint8, size)...)
}
