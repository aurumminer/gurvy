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

package bw761

import (
	"math/big"

	"github.com/consensys/gurvy"
	"github.com/consensys/gurvy/bw761/fp"
	"github.com/consensys/gurvy/utils"
)

// https://eprint.iacr.org/2020/351.pdf

// E: y**2=x**3-1
// Etwist: y**2 = x**3+4
// Tower: Fp->Fp6, u**6=-4
// Generator (same as BLS377): x=9586122913090633729
// optimal Ate loops: x+1, x**2-x-1
// Fp: p=6891450384315732539396789682275657542479668912536150109513790160209623422243491736087683183289411687640864567753786613451161759120554247759349511699125301598951605099378508850372543631423596795951899700429969112842764913119068299
// Fr: r=258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177

// ID bls377 ID
var ID = gurvy.BW761

// B b coeff of the curve
var B fp.Element

// generators of the r-torsion group, resp. in ker(pi-id), ker(Tr)
var g1Gen G1Jac
var g2Gen G2Jac

// point at infinity
var g1Infinity G1Jac
var g2Infinity G2Jac

// optimal Ate loop counters
// Miller loop 1: f(P), div(f) = (x+1)(Q)-([x+1]Q)-x(O)
// Miller loop 2: f(P), div(f) = (x**3-x**2-x)(Q) -([x**3-x**2-x]Q)-(x**3-x**2-x-1)(O)
var loopCounter1 [64]int8
var loopCounter2 [127]int8

func init() {

	B.SetOne().Neg(&B)

	g1Gen.X.SetString("5492337019202608651620810666633622531924946248948182754748114963334556774714407693672822645637243083342924475378144397780999025266189779523629084326871556483802038026432771927197170911996417793635501066231650458516636932478125208")
	g1Gen.Y.SetString("4874298780810344118673004453041997030286302865034758641338313952140849332867290574388366379298818956144982860224857872858166812124104845663394852158352478303048122861831479086904887356602146134586313962565783961814162269209043907")
	g1Gen.Z.SetString("1")

	g2Gen.X.SetString("5779457169892140542970811884673908634889239063901429247094594197042136765689827803062459420720318762253427359282239252479201196985966853806926626938528693270647807548111019296972244105103687281416386903420911111573334083829048020")
	g2Gen.Y.SetString("2945005085389580383802706904000483833228424888054664780252599806365093320701303614818391222418768857269542753796449953578553937529004880983494788715529986360817835802796138196037201453469654110552028363169895102423753717534586247")
	g2Gen.Z.SetString("1")

	//binary decomposition of 9586122913090633729, little endian
	loopCounter1 = [64]int8{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1}

	T, _ := new(big.Int).SetString("91893752504881257691937156713741811711", 10)
	utils.NafDecomposition(T, loopCounter2[:])

	g1Infinity.X.SetOne()
	g1Infinity.Y.SetOne()

	g2Infinity.X.SetOne()
	g2Infinity.Y.SetOne()

}