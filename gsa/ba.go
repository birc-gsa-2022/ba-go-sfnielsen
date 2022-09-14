package gsa

// Borderarray computes the border array over the string x. The border
// array ba will at index i have the length of the longest proper border
// of the string x[:i+1], i.e. the longest non-empty string that is both
// a prefix and a suffix of x[:i+1].
func Borderarray(x string) []int {
	ba := make([]int, len(x))

	b := 0
	i := 1
	for i < len(x) {
		if x[b] == x[i] {
			b++
			ba[i] = b
			i++
			continue
		}
		if b > 0 {
			b = ba[b-1]
			continue
		}
		i++
	}

	return ba
}

// StrictBorderarray computes the strict border array over the string x.
// This is almost the same as the border array, but ba[i] will be the
// longest proper border of the string x[:i+1] such that x[ba[i]] != x[i].
func StrictBorderarray(x string) []int {
	ba := make([]int, len(x))

	b := 0
	i := 1
	for i < len(x) {
		if x[b] == x[i] {
			b++
			if i == len(x)-1 {
				ba[i] = b
			}
			i++
			continue
		}
		if ba[i-1] == 0 {
			ba[i-1] = b
		}
		if b > 0 {
			b = ba[b-1]
			continue
		}
		i++
	}
	return ba

	/* **simpler but 2*n solution***

	ba = Borderarray(x)
	for i := 1; i < len(ba); i++ {
		if ba[i-1] == ba[i]-1 {
			ba[i-1] = 0
		}
	}
	return ba
	*/
}
