package lists

func ExamplePolynomial_Add() {
	p1 := new(Polynomial)
	p1.Insert(&PolynomialNode{exponent: 100, coefficient: 100})
	p1.Insert(&PolynomialNode{exponent: 1, coefficient: 3})
	p1.Insert(&PolynomialNode{exponent: 10, coefficient: 2})
	p1.Insert(&PolynomialNode{exponent: 10, coefficient: 3})
	p1.Print()

	p2 := new(Polynomial)
	p2.Insert(&PolynomialNode{exponent: 100, coefficient: 100})
	p2.Insert(&PolynomialNode{exponent: 10, coefficient: 3})
	p2.Insert(&PolynomialNode{exponent: 50, coefficient: 50})
	p2.Insert(&PolynomialNode{exponent: 5, coefficient: 5})
	p2.Insert(&PolynomialNode{exponent: 1, coefficient: 4})
	p2.Print()

	result := p1.Add(p2)
	result.Print()

	// Output:
	// 100 * x^100 + 5 * x^10 + 3 * x^1
	// 100 * x^100 + 50 * x^50 + 3 * x^10 + 5 * x^5 + 4 * x^1
	// 200 * x^100 + 50 * x^50 + 8 * x^10 + 5 * x^5 + 7 * x^1
}

func ExamplePolynomial_Multiply() {
	p1 := new(Polynomial)
	p1.Insert(&PolynomialNode{exponent: 100, coefficient: 100})
	p1.Insert(&PolynomialNode{exponent: 10, coefficient: 2})
	p1.Print()

	p2 := new(Polynomial)
	p2.Insert(&PolynomialNode{exponent: 100, coefficient: 100})
	p2.Insert(&PolynomialNode{exponent: 10, coefficient: 5})
	p2.Print()

	result := p1.Multiply(p2)
	result.Print()
	// Output:
	// 100 * x^100 + 2 * x^10
	// 100 * x^100 + 5 * x^10
	// 10000 * x^200 + 700 * x^110 + 10 * x^20
}
