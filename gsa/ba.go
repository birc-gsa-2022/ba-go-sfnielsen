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

		//we have a border
		if x[i] == x[b] {
			b++
			i++
			//case where last index element is a border
			if i == len(x) {
				ba[len(x)-1] = b
				break
			}
			//found border is not strict
			if x[i] == x[b] {
				if b-2 >= 0 {
					b_2 := ba[b-1]
					//while loop to find greatest strict border (if exists)
					for b_2 > 0 {
						if x[i-1] == x[b_2-1] {
							if x[i] != x[b_2] {
								ba[i-1] = b_2
								break
							}
						}
						b_2 = ba[b_2-1]
					}
				}
				//case where current index was a strict border
			} else {
				ba[i-1] = b
				b = 0
			}
			//case where current index is not a border
		} else {
			i++
			b = 0
		}
	}
	return ba

}
