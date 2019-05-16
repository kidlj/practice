package iterator

type Ints []int

type Iterator struct {
	index int
	data  Ints
}

// Classic
func (i Ints) IterClassic() *Iterator {
	return &Iterator{
		data: i,
	}
}

func (it *Iterator) HasNext() bool {
	return it.index < len(it.data)
}

func (it *Iterator) Next() int {
	v := it.data[it.index]
	it.index++
	return v
}

// Closure
func (i Ints) IterClosure() func() (int, bool) {
	index := 0
	return func() (int, bool) {
		if index >= len(i) {
			return 0, false
		}

		val, ok := i[index], true
		index++
		return val, ok
	}
}

// Channel
func (i Ints) IterChannel() <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range i {
			ch <- val
		}
		close(ch)
	}()
	return ch
}

// Do
func (i Ints) IterDo(fn func(int)) {
	for _, val := range i {
		fn(val)
	}
}
