/*
Copyright © 2020 ConsenSys

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fields

import (
	"github.com/consensys/gnark/frontend"
)

// Extension stores the non residue elmt for an extension of type Fp->Fp2->Fp6->Fp12 (Fp2 = Fp(u), Fp6 = Fp2(v), Fp12 = Fp6(w))
type Extension struct {

	// generators of each sub field
	uSquare interface{}
	vCube   *Fp2Elmt
	wSquare *Fp6Elmt

	// frobenius applied to generators
	frobv   interface{} // v**p  = (v**6)**(p-1/6)*v, frobv=(v**6)**(p-1/6), belongs to Fp)
	frobv2  interface{} // frobv2 = (v**6)**(p-1/3)
	frobw   interface{} // frobw = (w**12)**(p-1/12)
	frobvw  interface{} // frobvw = (v**6)**(p-1/6)*(w*12)**(p-1/12)
	frobv2w interface{} // frobv2w = (v**6)**(2*(p-1)/6)*(w*12)**(p-1/12)

	// frobenius square applied to generators
	frob2v   interface{} // v**(p**2)  = (v**6)**(p**2-1/6)*v, frobv=(v**6)**(p**2-1/6), belongs to Fp)
	frob2v2  interface{} // frobv2 = (v**6)**(2*(p**2-1)/6)
	frob2w   interface{} // frobw = (w**12)**(p**2-1/12)
	frob2vw  interface{} // frobvw = (v**6)**(p**2-1/6)*(w*12)**(p**2-1/12)
	frob2v2w interface{} // frobv2w = (v**6)**(2*(p**2-1)/6)*(w*12)**(p**2-1/12)

	// frobenius cube applied to generators
	frob3v   interface{} // v**(p**3)  = (v**6)**(p**3-1/6)*v, frobv=(v**6)**(p**3-1/6), belongs to Fp)
	frob3v2  interface{} // frobv2 = (v**6)**(2*(p**3-1)/6)
	frob3w   interface{} // frobw = (w**12)**(p**3-1/12)
	frob3vw  interface{} // frobvw = (v**6)**(p**3-1/6)*(w*12)**(p**3-1/12)
	frob3v2w interface{} // frobv2w = (v**6)**(2*(p**3-1)/6)*(w*12)**(p**3-1/12)

}

// Fp12Elmt element in a quadratic extension
type Fp12Elmt struct {
	C0, C1 Fp6Elmt
}

// GetBLS377ExtensionFp12 get extension field parameters for bls377
func GetBLS377ExtensionFp12(circuit *frontend.CS) Extension {
	res := Extension{}
	res.uSquare = 5
	res.vCube = &Fp2Elmt{X: circuit.ALLOCATE(0), Y: circuit.ALLOCATE(1)}
	res.wSquare = &Fp6Elmt{
		B0: NewFp2Zero(circuit),
		B1: NewFp2Elmt(circuit, circuit.ALLOCATE(1), circuit.ALLOCATE(0)),
		B2: NewFp2Zero(circuit),
	}

	res.frobv = "80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410946"
	res.frobv2 = "80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410945"
	res.frobw = "92949345220277864758624960506473182677953048909283248980960104381795901929519566951595905490535835115111760994353"
	res.frobvw = "216465761340224619389371505802605247630151569547285782856803747159100223055385581585702401816380679166954762214499"
	res.frobv2w = "123516416119946754630746545296132064952198520638002533875843642777304321125866014634106496325844844051843001220146"

	res.frob2v = "80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410945"
	res.frob2v2 = "258664426012969093929703085429980814127835149614277183275038967946009968870203535512256352201271898244626862047231"
	res.frob2w = "80949648264912719408558363140637477264845294720710499478137287262712535938301461879813459410946"
	res.frob2vw = "258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458176"
	res.frob2v2w = "258664426012969093929703085429980814127835149614277183275038967946009968870203535512256352201271898244626862047232"

	res.frob3v = "258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458176"
	res.frob3v2 = "1"
	res.frob3w = "216465761340224619389371505802605247630151569547285782856803747159100223055385581585702401816380679166954762214499"
	res.frob3vw = "42198664672744474621281227892288285906241943207628877683080515507620245292955241189266486323192680957485559243678"
	res.frob3v2w = "216465761340224619389371505802605247630151569547285782856803747159100223055385581585702401816380679166954762214499"

	return res
}

// NewFp12Elmt creates a fp6elmt from fp elmts
func NewFp12Elmt(circuit *frontend.CS, a, b, c, d, e, f, g, h, i, j, k, l interface{}) Fp12Elmt {

	var res Fp12Elmt

	res.C0.B0.X = circuit.ALLOCATE(a)
	res.C0.B0.Y = circuit.ALLOCATE(b)
	res.C0.B1.X = circuit.ALLOCATE(c)
	res.C0.B1.Y = circuit.ALLOCATE(d)
	res.C0.B2.X = circuit.ALLOCATE(e)
	res.C0.B2.Y = circuit.ALLOCATE(f)
	res.C1.B0.X = circuit.ALLOCATE(g)
	res.C1.B0.Y = circuit.ALLOCATE(h)
	res.C1.B1.X = circuit.ALLOCATE(i)
	res.C1.B1.Y = circuit.ALLOCATE(j)
	res.C1.B2.X = circuit.ALLOCATE(k)
	res.C1.B2.Y = circuit.ALLOCATE(l)

	return res
}

// NewFp12ElmtNil creates a fp6elmt from fp elmts
func NewFp12ElmtNil(circuit *frontend.CS) Fp12Elmt {

	a := NewFp6Elmt(circuit, nil, nil, nil, nil, nil, nil)
	b := NewFp6Elmt(circuit, nil, nil, nil, nil, nil, nil)

	res := Fp12Elmt{
		C0: a,
		C1: b,
	}
	return res
}

// SetOne returns a newly allocated element equal to 1
func (e *Fp12Elmt) SetOne(circuit *frontend.CS) *Fp12Elmt {
	e.C0.B0.X = circuit.ALLOCATE(1)
	e.C0.B0.Y = circuit.ALLOCATE(0)
	e.C0.B1.X = circuit.ALLOCATE(0)
	e.C0.B1.Y = circuit.ALLOCATE(0)
	e.C0.B2.X = circuit.ALLOCATE(0)
	e.C0.B2.Y = circuit.ALLOCATE(0)
	e.C1.B0.X = circuit.ALLOCATE(0)
	e.C1.B0.Y = circuit.ALLOCATE(0)
	e.C1.B1.X = circuit.ALLOCATE(0)
	e.C1.B1.Y = circuit.ALLOCATE(0)
	e.C1.B2.X = circuit.ALLOCATE(0)
	e.C1.B2.Y = circuit.ALLOCATE(0)
	return e
}

// Assign assigne e to e1
func (e *Fp12Elmt) Assign(circuit *frontend.CS, e1 *Fp12Elmt) *Fp12Elmt {
	e.C0.B0.X = circuit.ALLOCATE(e1.C0.B0.X)
	e.C0.B0.Y = circuit.ALLOCATE(e1.C0.B0.Y)
	e.C0.B1.X = circuit.ALLOCATE(e1.C0.B1.X)
	e.C0.B1.Y = circuit.ALLOCATE(e1.C0.B1.Y)
	e.C0.B2.X = circuit.ALLOCATE(e1.C0.B2.X)
	e.C0.B2.Y = circuit.ALLOCATE(e1.C0.B2.Y)
	e.C1.B0.X = circuit.ALLOCATE(e1.C1.B0.X)
	e.C1.B0.Y = circuit.ALLOCATE(e1.C1.B0.Y)
	e.C1.B1.X = circuit.ALLOCATE(e1.C1.B1.X)
	e.C1.B1.Y = circuit.ALLOCATE(e1.C1.B1.Y)
	e.C1.B2.X = circuit.ALLOCATE(e1.C1.B2.X)
	e.C1.B2.Y = circuit.ALLOCATE(e1.C1.B2.Y)
	return e
}

// Add adds 2 elmts in Fp12
func (e *Fp12Elmt) Add(circuit *frontend.CS, e1, e2 *Fp12Elmt) *Fp12Elmt {
	e.C0.Add(circuit, &e1.C0, &e2.C0)
	e.C1.Add(circuit, &e1.C1, &e2.C1)
	return e
}

// Sub substracts 2 elmts in Fp12
func (e *Fp12Elmt) Sub(circuit *frontend.CS, e1, e2 *Fp12Elmt) *Fp12Elmt {
	e.C0.Sub(circuit, &e1.C0, &e2.C0)
	e.C1.Sub(circuit, &e1.C1, &e2.C1)
	return e
}

// Neg negates an Fp6elmt
func (e *Fp12Elmt) Neg(circuit *frontend.CS, e1 *Fp12Elmt) *Fp12Elmt {
	e.C0.Neg(circuit, &e1.C0)
	e.C1.Neg(circuit, &e1.C1)
	return e
}

// Mul multiplies 2 elmts in Fp12
func (e *Fp12Elmt) Mul(circuit *frontend.CS, e1, e2 *Fp12Elmt, ext Extension) *Fp12Elmt {
	a := NewFp6Elmt(circuit, nil, nil, nil, nil, nil, nil)
	b := NewFp6Elmt(circuit, nil, nil, nil, nil, nil, nil)
	c := NewFp6Elmt(circuit, nil, nil, nil, nil, nil, nil)
	d := NewFp6Elmt(circuit, nil, nil, nil, nil, nil, nil)
	a.Mul(circuit, &e1.C0, &e2.C0, ext)
	b.Mul(circuit, &e1.C1, &e2.C1, ext).
		Mul(circuit, &b, ext.wSquare, ext)
	c.Mul(circuit, &e1.C0, &e2.C1, ext)
	d.Mul(circuit, &e1.C1, &e2.C0, ext)
	e.C0.Add(circuit, &a, &b)
	e.C1.Add(circuit, &c, &d)
	return e
}

// Conjugate applies Frob**6 (conjugation over Fp6)
func (e *Fp12Elmt) Conjugate(circuit *frontend.CS, e1 *Fp12Elmt) *Fp12Elmt {
	zero := NewFp6Zero(circuit)
	e.C1.Sub(circuit, &zero, &e1.C1)
	e.C0 = e1.C0
	return e
}

// MulByVW multiplies an e12 elmt by an elmt of the form a*VW (Fp6=Fp2(V), Fp12 = Fp6(W))
func (e *Fp12Elmt) MulByVW(circuit *frontend.CS, e1 *Fp12Elmt, e2 *Fp2Elmt, ext Extension) *Fp12Elmt {

	tmp := NewFp2Elmt(circuit, nil, nil)
	tmp.MulByIm(circuit, e2, ext)

	res := NewFp12ElmtNil(circuit)

	res.C0.B0.Mul(circuit, &e1.C1.B1, &tmp, ext)
	res.C0.B1.Mul(circuit, &e1.C1.B2, &tmp, ext)
	res.C0.B2.Mul(circuit, &e1.C1.B0, e2, ext)
	res.C1.B0.Mul(circuit, &e1.C0.B2, &tmp, ext)
	res.C1.B1.Mul(circuit, &e1.C0.B0, e2, ext)
	res.C1.B2.Mul(circuit, &e1.C0.B1, e2, ext)

	e.C0 = res.C0
	e.C1 = res.C1

	return e
}

// MulByV multiplies an e12 elmt by an elmt of the form a*V (Fp6=Fp2(V), Fp12 = Fp6(W))
func (e *Fp12Elmt) MulByV(circuit *frontend.CS, e1 *Fp12Elmt, e2 *Fp2Elmt, ext Extension) *Fp12Elmt {

	tmp := NewFp2Elmt(circuit, nil, nil)
	tmp.MulByIm(circuit, e2, ext)

	res := NewFp12ElmtNil(circuit)

	res.C0.B0.Mul(circuit, &e1.C0.B2, &tmp, ext)
	res.C0.B1.Mul(circuit, &e1.C0.B0, e2, ext)
	res.C0.B2.Mul(circuit, &e1.C0.B1, e2, ext)
	res.C1.B0.Mul(circuit, &e1.C1.B2, &tmp, ext)
	res.C1.B1.Mul(circuit, &e1.C1.B0, e2, ext)
	res.C1.B2.Mul(circuit, &e1.C1.B1, e2, ext)

	e.C0 = res.C0
	e.C1 = res.C1

	return e
}

// MulByV2W multiplies an e12 elmt by an elmt of the form a*V**2W (Fp6=Fp2(V), Fp12 = Fp6(W))
func (e *Fp12Elmt) MulByV2W(circuit *frontend.CS, e1 *Fp12Elmt, e2 *Fp2Elmt, ext Extension) *Fp12Elmt {

	tmp := NewFp2Elmt(circuit, nil, nil)
	tmp.MulByIm(circuit, e2, ext)

	res := NewFp12ElmtNil(circuit)

	res.C0.B0.Mul(circuit, &e1.C1.B0, &tmp, ext)
	res.C0.B1.Mul(circuit, &e1.C1.B1, &tmp, ext)
	res.C0.B2.Mul(circuit, &e1.C1.B2, &tmp, ext)
	res.C1.B0.Mul(circuit, &e1.C0.B1, &tmp, ext)
	res.C1.B1.Mul(circuit, &e1.C0.B2, &tmp, ext)
	res.C1.B2.Mul(circuit, &e1.C0.B0, e2, ext)

	e.C0 = res.C0
	e.C1 = res.C1

	return e
}

// Frobenius applies frob to an fp12 elmt
func (e *Fp12Elmt) Frobenius(circuit *frontend.CS, e1 *Fp12Elmt, ext Extension) *Fp12Elmt {

	e.C0.B0.Conjugate(circuit, &e1.C0.B0)
	e.C0.B1.Conjugate(circuit, &e1.C0.B1).MulByFp(circuit, &e.C0.B1, ext.frobv)
	e.C0.B2.Conjugate(circuit, &e1.C0.B2).MulByFp(circuit, &e.C0.B2, ext.frobv2)
	e.C1.B0.Conjugate(circuit, &e1.C1.B0).MulByFp(circuit, &e.C1.B0, ext.frobw)
	e.C1.B1.Conjugate(circuit, &e1.C1.B1).MulByFp(circuit, &e.C1.B1, ext.frobvw)
	e.C1.B2.Conjugate(circuit, &e1.C1.B2).MulByFp(circuit, &e.C1.B2, ext.frobv2w)

	return e

}

// FrobeniusSquare applies frob**2 to an fp12 elmt
func (e *Fp12Elmt) FrobeniusSquare(circuit *frontend.CS, e1 *Fp12Elmt, ext Extension) *Fp12Elmt {

	e.C0.B0 = e1.C0.B0
	e.C0.B1.MulByFp(circuit, &e1.C0.B1, ext.frob2v)
	e.C0.B2.MulByFp(circuit, &e1.C0.B2, ext.frob2v2)
	e.C1.B0.MulByFp(circuit, &e1.C1.B0, ext.frob2w)
	e.C1.B1.MulByFp(circuit, &e1.C1.B1, ext.frob2vw)
	e.C1.B2.MulByFp(circuit, &e1.C1.B2, ext.frob2v2w)

	return e
}

// FrobeniusCube applies frob**2 to an fp12 elmt
func (e *Fp12Elmt) FrobeniusCube(circuit *frontend.CS, e1 *Fp12Elmt, ext Extension) *Fp12Elmt {

	e.C0.B0.Conjugate(circuit, &e1.C0.B0)
	e.C0.B1.Conjugate(circuit, &e1.C0.B1).MulByFp(circuit, &e.C0.B1, ext.frob3v)
	e.C0.B2.Conjugate(circuit, &e1.C0.B2).MulByFp(circuit, &e.C0.B2, ext.frob3v2)
	e.C1.B0.Conjugate(circuit, &e1.C1.B0).MulByFp(circuit, &e.C1.B0, ext.frob3w)
	e.C1.B1.Conjugate(circuit, &e1.C1.B1).MulByFp(circuit, &e.C1.B1, ext.frob3vw)
	e.C1.B2.Conjugate(circuit, &e1.C1.B2).MulByFp(circuit, &e.C1.B2, ext.frob3v2w)

	return e
}

// Inverse inverse an elmt in Fp12
func (e *Fp12Elmt) Inverse(circuit *frontend.CS, e1 *Fp12Elmt, ext Extension) *Fp12Elmt {

	var t [2]Fp6Elmt
	var buf Fp6Elmt

	t[0].Mul(circuit, &e1.C0, &e1.C0, ext)
	t[1].Mul(circuit, &e1.C1, &e1.C1, ext)

	buf.MulByV(circuit, &t[1], ext)
	t[0].Sub(circuit, &t[0], &buf)

	t[1].Inverse(circuit, &t[0], ext)
	e.C0.Mul(circuit, &e1.C0, &t[1], ext)
	e.C1.Mul(circuit, &e1.C1, &t[1], ext).Neg(circuit, &e.C1)

	return e
}

// ConjugateFp12 conjugates an Fp12 elmt (applies Frob**6)
func (e *Fp12Elmt) ConjugateFp12(circuit *frontend.CS, e1 *Fp12Elmt) *Fp12Elmt {
	e.C0 = e1.C0
	e.C1.Neg(circuit, &e1.C1)
	return e
}

// Select sets e to r1 if b=1, r2 otherwise
func (e *Fp12Elmt) Select(circuit *frontend.CS, b *frontend.Constraint, r1, r2 *Fp12Elmt) *Fp12Elmt {

	e.C0.B0.X = circuit.SELECT(b, r1.C0.B0.X, r2.C0.B0.X)
	e.C0.B0.Y = circuit.SELECT(b, r1.C0.B0.Y, r2.C0.B0.Y)
	e.C0.B1.X = circuit.SELECT(b, r1.C0.B1.X, r2.C0.B1.X)
	e.C0.B1.Y = circuit.SELECT(b, r1.C0.B1.Y, r2.C0.B1.Y)
	e.C0.B2.X = circuit.SELECT(b, r1.C0.B2.X, r2.C0.B2.X)
	e.C0.B2.Y = circuit.SELECT(b, r1.C0.B2.Y, r2.C0.B2.Y)
	e.C1.B0.X = circuit.SELECT(b, r1.C1.B0.X, r2.C1.B0.X)
	e.C1.B0.Y = circuit.SELECT(b, r1.C1.B0.Y, r2.C1.B0.Y)
	e.C1.B1.X = circuit.SELECT(b, r1.C1.B1.X, r2.C1.B1.X)
	e.C1.B1.Y = circuit.SELECT(b, r1.C1.B1.Y, r2.C1.B1.Y)
	e.C1.B2.X = circuit.SELECT(b, r1.C1.B2.X, r2.C1.B2.X)
	e.C1.B2.Y = circuit.SELECT(b, r1.C1.B2.Y, r2.C1.B2.Y)

	return e
}

// FixedExponentiation compute e1**exponent, where the exponent is hardcoded
// This function is only used for the final expo of the pairing for bls377, so the exponent is supposed to be hardcoded
// and on 64 bits.
func (e *Fp12Elmt) FixedExponentiation(circuit *frontend.CS, e1 *Fp12Elmt, exponent uint64, ext Extension) *Fp12Elmt {

	var expoBin [64]uint8
	for i := 0; i < 64; i++ {
		expoBin[i] = uint8((exponent >> (63 - i))) & 1
	}

	res := NewFp12Elmt(circuit, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	for i := 0; i < 64; i++ {
		res.Mul(circuit, &res, &res, ext)
		if expoBin[i] == 1 {
			res.Mul(circuit, &res, e1, ext)
		}
	}
	e.Assign(circuit, &res)

	return e
}

// FinalExpoBLS final  exponentation for curves of the bls family (t is the parameter used to generate the curve)
func (e *Fp12Elmt) FinalExpoBLS(circuit *frontend.CS, e1 *Fp12Elmt, genT uint64, ext Extension) *Fp12Elmt {

	var res Fp12Elmt
	res.Assign(circuit, e1)

	var t [6]Fp12Elmt

	t[0].FrobeniusCube(circuit, e1, ext).FrobeniusCube(circuit, &t[0], ext)

	res.Inverse(circuit, &res, ext)
	t[0].Mul(circuit, &t[0], &res, ext)

	res.FrobeniusSquare(circuit, &t[0], ext).Mul(circuit, &res, &t[0], ext)

	t[0].ConjugateFp12(circuit, &res).Mul(circuit, &t[0], &t[0], ext)
	t[5].FixedExponentiation(circuit, &res, genT, ext)
	t[1].Mul(circuit, &t[5], &t[5], ext)
	t[3].Mul(circuit, &t[0], &t[5], ext)

	t[0].FixedExponentiation(circuit, &t[3], genT, ext)
	t[2].FixedExponentiation(circuit, &t[0], genT, ext)
	t[4].FixedExponentiation(circuit, &t[2], genT, ext)

	t[4].Mul(circuit, &t[1], &t[4], ext)
	t[1].FixedExponentiation(circuit, &t[4], genT, ext)
	t[3].Conjugate(circuit, &t[3])
	t[1].Mul(circuit, &t[3], &t[1], ext)
	t[1].Mul(circuit, &t[1], &res, ext)

	t[0].Mul(circuit, &t[0], &res, ext)
	t[0].FrobeniusCube(circuit, &t[0], ext)

	t[3].Conjugate(circuit, &res)
	t[4].Mul(circuit, &t[3], &t[4], ext)
	t[4].Frobenius(circuit, &t[4], ext)

	t[5].Mul(circuit, &t[2], &t[5], ext)
	t[5].FrobeniusSquare(circuit, &t[5], ext)

	t[5].Mul(circuit, &t[5], &t[0], ext)
	t[5].Mul(circuit, &t[5], &t[4], ext)
	t[5].Mul(circuit, &t[5], &t[1], ext)

	e.Assign(circuit, &t[5])

	return e
}