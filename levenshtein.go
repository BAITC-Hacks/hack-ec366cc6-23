package main

// Levenshtein calculates the edit distance between two strings
// using a memory-efficient 1D slice approach.
func Levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	lenA, lenB := len(ra), len(rb)
	
	f := make([]int, lenB+1)
	for j := range f {
		f[j] = j
	}
	
	for i := 1; i <= lenA; i++ {
		prev := f[0]
		f[0] = i
		for j := 1; j <= lenB; j++ {
			tmp := f[j]
			cost := 0
			if ra[i-1] != rb[j-1] {
				cost = 1
			}
			f[j] = min(min(f[j]+1, f[j-1]+1), prev+cost)
			prev = tmp
		}
	}
	return f[lenB]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
