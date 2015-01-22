package ints

// Count counts the number of non-overlapping instances of t in s
func Count(s []int, t int) (count int) {
	for _, j := range s {
		if j == t {
			count++
		}
	}
	return count
}

// Contains returns true if t is in s
func Contains(s []int, t int) bool {
	return Index(s, t) >= 0
}

// Index returns the index of the first instance of t in s or -1 if t is not present in s.
func Index(s []int, t int) int {
	for i, j := range s {
		if j == t {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of t in s or -1 of t is not present in s.
func LastIndex(s []int, t int) int {
	found_index := -1
	for i, j := range s {
		if j == t {
			found_index = i
		}
	}
	return found_index
}

//Removes all instances of t if it exists in s
func RemoveAll(s []int, t int) []int {
	for i := Index(s, t); i >= 0; i = Index(s, t) {
		s = RemoveAt(s, i)
	}
	return s
}

//Removes element at i from s
func RemoveAt(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}
	res := make([]int, len(s)-1)
	copy(res[0:i], s[0:i])
	copy(res[i:], s[i+1:])
	return res
}

//TODO Replace, Repeat, Split, Map
