// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bls381

import "github.com/consensys/gurvy/bls381/fp"

// Mul sets z to the e2-product of x,y, returns z
func (z *e2) Mul(x, y *e2) *e2 {
	var a, b, c fp.Element
	a.Add(&x.A0, &x.A1)
	b.Add(&y.A0, &y.A1)
	a.Mul(&a, &b)
	b.Mul(&x.A0, &y.A0)
	c.Mul(&x.A1, &y.A1)
	z.A1.Sub(&a, &b).Sub(&z.A1, &c)
	z.A0.Sub(&b, &c) //z.A0.MulByNonResidue(&c).Add(&z.A0, &b)
	return z
}

// Square sets z to the e2-product of x,x returns z
func (z *e2) Square(x *e2) *e2 {
	// algo 22 https://eprint.iacr.org/2010/354.pdf
	var a, b fp.Element
	a.Add(&x.A0, &x.A1)
	b.Sub(&x.A0, &x.A1)
	a.Mul(&a, &b)
	b.Mul(&x.A0, &x.A1).Double(&b)
	z.A0.Set(&a)
	z.A1.Set(&b)
	return z
}

// MulByNonResidue multiplies a e2 by (1,1)
func (z *e2) MulByNonResidue(x *e2) *e2 {
	var a fp.Element
	a.Sub(&x.A0, &x.A1)
	z.A1.Add(&x.A0, &x.A1)
	z.A0.Set(&a)
	return z
}

// MulByNonResidueInv multiplies a e2 by (1,1)^{-1}
func (z *e2) MulByNonResidueInv(x *e2) *e2 {

	twoinv := fp.Element{
		1730508156817200468,
		9606178027640717313,
		7150789853162776431,
		7936136305760253186,
		15245073033536294050,
		1728177566264616342,
	}

	var tmp fp.Element
	tmp.Add(&x.A0, &x.A1)
	z.A1.Sub(&x.A1, &x.A0).Mul(&z.A1, &twoinv)
	z.A0.Set(&tmp).Mul(&z.A0, &twoinv)

	return z
}

// Inverse sets z to the e2-inverse of x, returns z
func (z *e2) Inverse(x *e2) *e2 {
	// Algorithm 8 from https://eprint.iacr.org/2010/354.pdf
	var t0, t1 fp.Element
	t0.Square(&x.A0)
	t1.Square(&x.A1)
	t0.Add(&t0, &t1)
	t1.Inverse(&t0)
	z.A0.Mul(&x.A0, &t1)
	z.A1.Mul(&x.A1, &t1).Neg(&z.A1)

	return z
}

// norm sets x to the norm of z
func (z *e2) norm(x *fp.Element) {
	var tmp fp.Element
	x.Square(&z.A0)
	tmp.Square(&z.A1)
	x.Add(x, &tmp)
}
