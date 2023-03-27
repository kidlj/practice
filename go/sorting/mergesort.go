package sorting

func MergeSort(s []int) {
	tmp := make([]int, len(s))
	sort(s, tmp)
}

func sort(s, tmp []int) {
	if len(s) > 1 {
		center := len(s) / 2
		sort(s[:center], tmp)
		sort(s[center:], tmp)
		merge(s[:center], s[center:], tmp)
	}
}

func merge(s1, s2, tmp []int) {
	left, right, cursor := 0, 0, 0
	for left < len(s1) && right < len(s2) {
		if s1[left] <= s2[right] {
			tmp[cursor] = s1[left]
			left++
			cursor++
		} else {
			tmp[cursor] = s2[right]
			right++
			cursor++
		}
	}

	for left < len(s1) {
		tmp[cursor] = s1[left]
		left++
		cursor++
	}
	for right < len(s2) {
		tmp[cursor] = s2[right]
		right++
		cursor++
	}

	cursor = 0
	for i := 0; i < len(s1); i++ {
		s1[i] = tmp[cursor]
		cursor++
	}
	for i := 0; i < len(s2); i++ {
		s2[i] = tmp[cursor]
		cursor++
	}
}
