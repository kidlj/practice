package lists

import (
	"fmt"
	"strings"
)

type PolynomialNode struct {
	Coefficient int
	Exponent    int
	succ        *PolynomialNode
}

// If you need to init() your data structure, make it unexported.
type polynomial struct {
	header *PolynomialNode
}

// And export New... func to init() once.
func NewPolynomial() *polynomial {
	p := new(polynomial)
	p.init()
	return p
}

func (l *polynomial) init() {
	if l.header == nil {
		dummy := new(PolynomialNode)
		l.header = dummy
	}
}

func (l *polynomial) Add(poly *polynomial) *polynomial {
	result := NewPolynomial()
	result.add(l)
	result.add(poly)
	return result
}

func (l *polynomial) add(poly *polynomial) {
	for p := poly.header.succ; p != nil; p = p.succ {
		l.Insert(p)
	}
}

func (l *polynomial) Insert(n *PolynomialNode) {
	p := l.findPrevious(n.Exponent)
	if p.succ != nil && p.succ.Exponent == n.Exponent {
		p.succ.Coefficient += n.Coefficient
	} else {
		// allocate new node to be inserted
		n = &PolynomialNode{Exponent: n.Exponent, Coefficient: n.Coefficient, succ: p.succ}
		p.succ = n
	}
}

func (l *polynomial) findPrevious(exponent int) *PolynomialNode {
	p := l.header
	for p.succ != nil && p.succ.Exponent > exponent {
		p = p.succ
	}
	return p
}

// O(M^2N^2)
func (l *polynomial) Multiply(poly *polynomial) *polynomial {
	result := NewPolynomial()
	for p := l.header.succ; p != nil; p = p.succ {
		for n := poly.header.succ; n != nil; n = n.succ {
			result.multiply(p, n)
		}
	}
	return result
}

func (l *polynomial) multiply(n1, n2 *PolynomialNode) {
	n := &PolynomialNode{
		Coefficient: n1.Coefficient * n2.Coefficient,
		Exponent:    n1.Exponent + n2.Exponent,
	}
	l.Insert(n)
}

func (l *polynomial) String() string {
	items := []string{}
	for p := l.header.succ; p != nil; p = p.succ {
		items = append(items, fmt.Sprintf("%d * x^%d", p.Coefficient, p.Exponent))
	}
	result := strings.Join(items, " + ")
	return result
}
