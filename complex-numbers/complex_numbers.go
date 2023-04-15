package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imaginary + n2.imaginary}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imaginary - n2.imaginary}
}

func (n1 Number) Multiply(n2 Number) Number {
	realResult := n1.real*n2.real - n1.imaginary*n2.imaginary
	imaginaryResult := n1.imaginary*n2.real + n1.real*n2.imaginary
	return Number{realResult, imaginaryResult}
}

func (n Number) Times(factor float64) Number {
	return Number{n.real * factor, n.imaginary * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	powerN2 := (n2.real*n2.real + n2.imaginary*n2.imaginary)
	realResult := (n1.real*n2.real + n1.imaginary*n2.imaginary) / powerN2
	imaginaryResult := (n1.imaginary*n2.real - n1.real*n2.imaginary) / powerN2
	return Number{realResult, float64(imaginaryResult)}
}

func (n Number) Conjugate() Number {
	return Number{n.real, -n.imaginary}
}

func (n Number) Abs() float64 {
	squares := n.real*n.real + n.imaginary*n.imaginary
	return math.Sqrt(squares)
}

func (n Number) Exp() Number {
	expA := math.Exp(n.real)
	realResult := expA * math.Cos(n.imaginary)
	imaginaryResult := expA * math.Sin(n.imaginary)
	return Number{realResult, imaginaryResult}
}
