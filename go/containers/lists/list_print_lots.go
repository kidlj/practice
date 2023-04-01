package lists

func (l *linkedList) printLots(seq List) []any {
	result := make([]any, 0, seq.Size())
	iter := seq.NewIterator()
	v, ok := iter.Next()
	n := l.head.succ
	i := 1
	for n != nil && ok {
		if i == v {
			result = append(result, n.item)
			v, ok = iter.Next()
		}
		i++
		n = n.succ
	}
	return result
}

func (l *linkedList) printLotsRaw(seq *linkedList) []any {
	result := make([]any, 0, seq.Size())
	n := l.head.succ
	c := seq.head.succ

	i := 1
	for n != nil && c != nil {
		if i == c.item {
			result = append(result, n.item)
			c = c.succ
		}
		i++
		n = n.succ
	}
	return result
}
