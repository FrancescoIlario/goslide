package main

// Cut removes a piece of the array
func Cut(s []int16, i, j int) []int16 {
	return append(s[:i], s[j:]...)
}

// Copy copies the provided array in a new one with the same capacity
func Copy(s []int16) []int16 {
	b := make([]int16, len(s), cap(s))
	copy(b, s)
	return b
}

// CopyFit copies the provided array in a new one with capacity set as the length of the input one
func CopyFit(s []int16) []int16 {
	b := make([]int16, len(s), len(s))
	copy(b, s)
	return b
}

// Delete removes an entry from the array
func Delete(s []int16, i int) []int16 {
	return append(s[:i], s[i+1:]...)
}

// DeleteWithoutOrder removes an entry from the array
func DeleteWithoutOrder(s []int16, i int) []int16 {
	lp := len(s) - 1
	s[i] = s[lp]
	return s[:lp]
}

// FilterInPlace filters the provided array using the the filter function
func FilterInPlace(s []int16, filter func(int, int16) bool) []int16 {
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
func Filter(s []int16, filter func(int, int16) bool) []int16 {
	scp := CopyFit(s)
	return FilterInPlace(scp, filter)
}

// PopLast pops the last element from the provided array
func PopLast(s []int16) (int16, []int16) {
	ep := len(s) - 1
	return s[ep], s[:ep]
}

// PopFirst pops the first element from the provided array
func PopFirst(s []int16) (int16, []int16) {
	return s[0], s[1:]
}

// InsertAtPos inserts the i slide into the s one at the pos position
func InsertAtPos(s []int16, pos int, i ...int16) []int16 {
	return append(append(s[pos:], i...), s[:pos]...)
}

// PushFront adds a new element in front of the slice
func PushFront(s []int16, e ...int16) []int16 {
	return append(e, s...)
}

// Extend adds new entries (with default values) at the provided slice at given position
func Extend(s []int16, pos, size int) []int16 {
	return append(s[:pos], append(make([]int16, size), s[pos:]...)...)
}

// ExtendEnd adds new entries (with default values) at the provided slice at the end
func ExtendEnd(s []int16, size int) []int16 {
	return append(s, make([]int16, size)...)
}
