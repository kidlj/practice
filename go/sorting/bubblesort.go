package sorting

func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j, j+1)
			}
		}
	}
}
