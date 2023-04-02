package algorithms

func printLots(l *listNode, s *listNode) []int {
	result := make([]int, 0)

	i := 1
	for l != nil && s != nil {
		if i == s.val {
			result = append(result, l.val)
			s = s.next
		}
		i++
		l = l.next
	}
	return result
}
