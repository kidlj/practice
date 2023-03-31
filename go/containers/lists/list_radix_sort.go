package lists

func RadixSort(src []int) []int {
	const base = 10
	buckets := [base]List{}
	for i := 0; i < len(buckets); i++ {
		buckets[i] = NewLinkedList()
	}

	finalPass := false
	d := 1
	for !finalPass {
		finalPass = true
		for _, n := range src {
			v := n / d / base
			if v > 0 {
				finalPass = false
			}
			i := (n / d) % base
			l := buckets[i]
			_ = l.Insert(l.Size(), n)
		}

		i := 0
		for _, l := range buckets {
			for j := 0; j < l.Size(); j++ {
				v, _ := l.Get(j)
				src[i] = v.(int)
				i++
			}
			l.Clear()
		}

		d = d * base
	}

	return src
}
