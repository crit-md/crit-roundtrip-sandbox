package sample

import "errors"

// ErrDivideByZero is returned by Div when b is zero.
var ErrDivideByZero = errors.New("sample: divide by zero")

func Add(a, b int) int {
    return a + b
}

func Sub(a, b int) int {
    return a - b
}

func Mul(a, b int) int {
    return a * b
}

// Div returns a / b, or ErrDivideByZero when b is zero.
func Div(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}
