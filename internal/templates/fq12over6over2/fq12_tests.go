package fq12over6over2

// Fq12Tests ...
const Fq12Tests = `

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/commands"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

// ------------------------------------------------------------
// tests

func TestE12ReceiverIsOperand(t *testing.T) {

	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100

	properties := gopter.NewProperties(parameters)

	genA := GenE12()
	genB := GenE12()

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (addition) should output the same result", prop.ForAll(
		func(a, b *e12) bool {
			var c, d e12
			d.Set(a)
			c.Add(a, b)
			a.Add(a, b)
			b.Add(&d, b)
			return a.Equal(b) && a.Equal(&c) && b.Equal(&c)
		},
		genA,
		genB,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (sub) should output the same result", prop.ForAll(
		func(a, b *e12) bool {
			var c, d e12
			d.Set(a)
			c.Sub(a, b)
			a.Sub(a, b)
			b.Sub(&d, b)
			return a.Equal(b) && a.Equal(&c) && b.Equal(&c)
		},
		genA,
		genB,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (mul) should output the same result", prop.ForAll(
		func(a, b *e12) bool {
			var c, d e12
			d.Set(a)
			c.Mul(a, b)
			a.Mul(a, b)
			b.Mul(&d, b)
			return a.Equal(b) && a.Equal(&c) && b.Equal(&c)
		},
		genA,
		genB,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (square) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.Square(a)
			a.Square(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (double) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.Double(a)
			a.Double(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (Inverse) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.Inverse(a)
			a.Inverse(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (Cyclotomic square) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.CyclotomicSquare(a)
			a.CyclotomicSquare(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (Conjugate) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.Conjugate(a)
			a.Conjugate(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (Frobenius) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.Frobenius(a)
			a.Frobenius(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (FrobeniusSquare) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.FrobeniusSquare(a)
			a.FrobeniusSquare(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName}}] Having the receiver as operand (FrobeniusCube) should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.FrobeniusCube(a)
			a.FrobeniusCube(a)
			return a.Equal(&b)
		},
		genA,
	))

	properties.TestingRun(t, gopter.ConsoleReporter(false))
}

func TestE12Ops(t *testing.T) {

	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100

	properties := gopter.NewProperties(parameters)

	genA := GenE12()
	genB := GenE12()

	properties.Property("[{{ toUpper .CurveName }}] sub & add should leave an element invariant", prop.ForAll(
		func(a, b *e12) bool {
			var c e12
			c.Set(a)
			c.Add(&c, b).Sub(&c, b)
			return c.Equal(a)
		},
		genA,
		genB,
	))

	properties.Property("[{{ toUpper .CurveName }}] mul & inverse should leave an element invariant", prop.ForAll(
		func(a, b *e12) bool {
			var c, d e12
			d.Inverse(b)
			c.Set(a)
			c.Mul(&c, b).Mul(&c, &d)
			return c.Equal(a)
		},
		genA,
		genB,
	))

	properties.Property("[{{ toUpper .CurveName }}] inverse twice should leave an element invariant", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.Inverse(a).Inverse(&b)
			return a.Equal(&b)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName }}] square and mul should output the same result", prop.ForAll(
		func(a *e12) bool {
			var b, c e12
			b.Mul(a, a)
			c.Square(a)
			return b.Equal(&c)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName }}] a + pi(a), a-pi(a) should be real", prop.ForAll(
		func(a *e12) bool {
			var b, c, d e12
			var e, f, g e6
			b.Conjugate(a)
			c.Add(a, &b)
			d.Sub(a, &b)
			e.Double(&a.C0)
			f.Double(&a.C1)
			return c.C1.Equal(&g) && d.C0.Equal(&g) && e.Equal(&c.C0) && f.Equal(&d.C1)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName }}] pi**12=id", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.Frobenius(a).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b).
				Frobenius(&b)
			return b.Equal(a)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName }}] (pi**2)**6=id", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.FrobeniusSquare(a).
				FrobeniusSquare(&b).
				FrobeniusSquare(&b).
				FrobeniusSquare(&b).
				FrobeniusSquare(&b).
				FrobeniusSquare(&b)
			return b.Equal(a)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName }}] (pi**3)**4=id", prop.ForAll(
		func(a *e12) bool {
			var b e12
			b.FrobeniusCube(a).
				FrobeniusCube(&b).
				FrobeniusCube(&b).
				FrobeniusCube(&b)
			return b.Equal(a)
		},
		genA,
	))

	properties.Property("[{{ toUpper .CurveName }}] cyclotomic square and square should be the same in the cyclotomic subgroup", prop.ForAll(
		func(a *e12) bool {
			var b, c, d e12
			b.FrobeniusCube(a).
				FrobeniusCube(&b)
			a.Inverse(a)
			b.Mul(&b, a)
			a.Set(&b)
			b.FrobeniusSquare(&b).Mul(&b, a)
			c.Square(&b)
			d.CyclotomicSquare(&b)
			return c.Equal(&d)
		},
		genA,
	))

	properties.TestingRun(t, gopter.ConsoleReporter(false))

}

// ------------------------------------------------------------
// benches

func BenchmarkE12Add(b *testing.B) {
	var a, c e12
	a.SetRandom()
	c.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Add(&a, &c)
	}
}

func BenchmarkE12Sub(b *testing.B) {
	var a, c e12
	a.SetRandom()
	c.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Sub(&a, &c)
	}
}

func BenchmarkE12Mul(b *testing.B) {
	var a, c e12
	a.SetRandom()
	c.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Mul(&a, &c)
	}
}

func BenchmarkE12Cyclosquare(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.CyclotomicSquare(&a)
	}
}

func BenchmarkE12Square(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Square(&a)
	}
}

func BenchmarkE12Inverse(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Inverse(&a)
	}
}

func BenchmarkE12Conjugate(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Conjugate(&a)
	}
}

func BenchmarkE12Frobenius(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Frobenius(&a)
	}
}

func BenchmarkE12FrobeniusSquare(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.FrobeniusSquare(&a)
	}
}

func BenchmarkE12FrobeniusCube(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.FrobeniusCube(&a)
	}
}

func BenchmarkE12Expt(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.Expt(&a)
	}
}

func BenchmarkE12FinalExponentiation(b *testing.B) {
	var a e12
	a.SetRandom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.FinalExponentiation(&a)
	}
}

`
