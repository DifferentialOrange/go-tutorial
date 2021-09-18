package solve

import (
    "errors";
    "math"
)

func QuadraticEquation(a float64, b float64, c float64) ([]float64, error) {
    // a x^2 + b x + c = 0
    if math.Abs(a) < 1e-9 {
        return []float64{}, errors.New("Quadratic coefficient is too close to zero")
    }

    D := b * b - 4 * a * c

    if D < 0 {
        return []float64{}, errors.New("Complex roots are not supported yet")
    }

    return []float64{ ( - b + math.Sqrt(D) ) / ( 2 * a ), ( - b - math.Sqrt(D) ) / ( 2 * a ) }, nil
}
