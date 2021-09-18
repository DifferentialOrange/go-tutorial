package solve

import "math"

func float64Equals(v1 float64, v2 float64) (bool) {
    var margin float64 = 1e-9
    return math.Abs(v1 - v2) < margin
}
