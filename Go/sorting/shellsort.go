package sorting

func ShellSort(s []int) {
	for increment := len(s) / 2; increment > 0; increment /= 2 {
		for i := increment; i < len(s); i++ {
			tmp := s[i]
			var j int
			for j = i; j >= increment; j -= increment {
				if tmp < s[j-increment] {
					s[j] = s[j-increment]
				} else {
					break
				}
			}
			s[j] = tmp
		}
	}
}
