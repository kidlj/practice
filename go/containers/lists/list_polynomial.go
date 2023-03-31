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

type Polynomial struct {
	header *PolynomialNode
}

func (l *Polynomial) init() {
	if l.header == nil {
		dummy := new(PolynomialNode)
		l.header = dummy
	}
}

func (l *Polynomial) Add(poly *Polynomial) *Polynomial {
	result := new(Polynomial)
	result.add(l)
	result.add(poly)
	return result
}

func (l *Polynomial) add(poly *Polynomial) {
	poly.init()
	for p := poly.header.succ; p != nil; p = p.succ {
		l.Insert(p)
	}
}

func (l *Polynomial) Insert(n *PolynomialNode) {
	l.init()
	p := l.findPrevious(n.Exponent)
	if p.succ != nil && p.succ.Exponent == n.Exponent {
		p.succ.Coefficient += n.Coefficient
	} else {
		// allocate new node to be inserted
		n = &PolynomialNode{Exponent: n.Exponent, Coefficient: n.Coefficient, succ: p.succ}
		p.succ = n
	}
}

func (l *Polynomial) findPrevious(exponent int) *PolynomialNode {
	p := l.header
	for p.succ != nil && p.succ.Exponent > exponent {
		p = p.succ
	}
	return p
}

// O(M^2N^2)
func (l *Polynomial) Multiply(poly *Polynomial) *Polynomial {
	result := new(Polynomial)
	result.init()
	for p := l.header.succ; p != nil; p = p.succ {
		for n := poly.header.succ; n != nil; n = n.succ {
			result.multiply(p, n)
		}
	}
	return result
}

func (l *Polynomial) multiply(n1, n2 *PolynomialNode) {
	n := &PolynomialNode{
		Coefficient: n1.Coefficient * n2.Coefficient,
		Exponent:    n1.Exponent + n2.Exponent,
	}
	l.Insert(n)
}

func (l *Polynomial) String() string {
	l.init()
	items := []string{}
	for p := l.header.succ; p != nil; p = p.succ {
		items = append(items, fmt.Sprintf("%d * x^%d", p.Coefficient, p.Exponent))
	}
	result := strings.Join(items, " + ")
	return result
}
