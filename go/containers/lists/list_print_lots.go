package lists

import "fmt"

func (l *linkedList) printLots(seqList List) {
	iter := seqList.NewIterator()
	v, ok := iter.Next()
	n := l.head.succ
	seq := 1
	for n != l.head && ok {
		if seq == v {
			fmt.Printf("%v ", n.item)
			v, ok = iter.Next()
		}
		seq++
		n = n.succ
	}
}
