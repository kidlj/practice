package stacks

type Calculator interface {
	Evaluate(postfix string) (float64, error)
	Reset()
}
