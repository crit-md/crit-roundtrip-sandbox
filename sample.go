package sample

import (
    "errors"
    "math"
)

// ErrOverflow is returned when a result does not fit in an int.
var ErrOverflow = errors.New("sample: integer overflow")

// Add returns a + b, or ErrOverflow if the sum overflows.
func Add(a, b int) (int, error) {
    if (b > 0 && a > math.MaxInt-b) || (b < 0 && a < math.MinInt-b) {
        return 0, ErrOverflow
    }
    return a + b, nil
}

func Sub(a, b int) int {
    return a - b
}

// Mul returns a * b, or ErrOverflow if the product overflows.
func Mul(a, b int) (int, error) {
    if a == 0 || b == 0 {
        return 0, nil
    }
    c := a * b
    if c/b != a {
        return 0, ErrOverflow
    }
    return c, nil
}

func Div(a, b int) int {
    return a / b
}
