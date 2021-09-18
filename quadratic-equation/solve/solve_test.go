package solve

import "testing"

func assertTwoRoots(t *testing.T, slice []float64) {
    if len(slice) != 2 {
        t.Fatalf(`Expected slice %v to contain two roots`, slice)
    }
}

func isSquareRootsEquals(t *testing.T, actual []float64, expected []float64) (bool) {
    assertTwoRoots(t, actual)
    assertTwoRoots(t, expected)

    if float64Equals(actual[0], expected[0]) && float64Equals(actual[1], expected[1]) {
        return true
    }

    if float64Equals(actual[0], expected[1]) && float64Equals(actual[1], expected[0]) {
        return true
    }

    return false
}

func isSquareRootsDifferent(t *testing.T, actual []float64, expected []float64) (bool) {
    return !isSquareRootsEquals(t, actual, expected)
}

func TestQuadraticEquationSolverComputesRealRoots(t *testing.T) {
    var a, b, c float64 = 1, 3, 2
    expected_roots := []float64{ -1, -2 }

    actual_roots, err := QuadraticEquation(a, b, c)
    if isSquareRootsDifferent(t, actual_roots, expected_roots) || err != nil {
        t.Fatalf(`Expected QuadraticEquation(%v, %v, %v) = %v %v, got %v %v`, a, b, c, expected_roots, nil, actual_roots, err)
    }
}
