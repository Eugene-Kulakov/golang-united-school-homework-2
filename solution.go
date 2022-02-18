package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type SidesCount_t int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:

const SidesTriangle SidesCount_t = 3
const SidesSquare SidesCount_t = 4
const SidesCircle SidesCount_t = 0

// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum SidesCount_t) float64 {
	switch sidesNum {
	case SidesTriangle:
		return sideLen * sideLen * math.Sqrt(3.0) / 4.0
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	default:
		return 0
	}
}
