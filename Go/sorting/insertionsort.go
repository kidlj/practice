package sorting

func InsertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		tmp := s[i]
		var j int
		for j = i; j > 0; j-- {
			if tmp < s[j-1] {
				s[j] = s[j-1]
			} else {
				break
			}
		}
		s[j] = tmp
	}
}
