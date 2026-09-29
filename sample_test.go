package sample

import "testing"

func TestDiv(t *testing.T) {
    if got, err := Div(6, 3); err != nil || got != 2 {
        t.Fatalf("Div(6, 3) = %d, %v", got, err)
    }
    if _, err := Div(1, 0); err != ErrDivideByZero {
        t.Fatalf("Div(1, 0) err = %v, want ErrDivideByZero", err)
    }
}
