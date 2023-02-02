package lists

import (
	"fmt"
	"strings"
)

type PolynomialNode struct {
	coefficient int
	exponent    int
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
	p := l.findPrevious(n.exponent)
	if p.succ != nil && p.succ.exponent == n.exponent {
		p.succ.coefficient += n.coefficient
	} else {
		// allocate new node to be inserted
		n = &PolynomialNode{exponent: n.exponent, coefficient: n.coefficient, succ: p.succ}
		p.succ = n
	}
}

func (l *Polynomial) findPrevious(exponent int) *PolynomialNode {
	p := l.header
	for p.succ != nil && p.succ.exponent > exponent {
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
	exponent := n1.exponent + n2.exponent
	coefficient := n1.coefficient * n2.coefficient
	p := l.findPrevious(exponent)
	if p.succ != nil && p.succ.exponent == exponent {
		p.succ.coefficient += coefficient
	} else {
		// allocate new node to be inserted
		n := &PolynomialNode{exponent: exponent, coefficient: coefficient, succ: p.succ}
		p.succ = n
	}
}

func (l *Polynomial) Print() {
	l.init()
	items := []string{}
	for p := l.header.succ; p != nil; p = p.succ {
		items = append(items, fmt.Sprintf("%d * x^%d", p.coefficient, p.exponent))
	}
	result := strings.Join(items, " + ")
	fmt.Println(result)
}