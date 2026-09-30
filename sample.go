package sample

func Add(a, b int) int {
    return a + b
}

func Sub(a, b int) int {
    return a - b
}

func Mul(a, b int) int {
    return a * b
}

func Div(a, b int) int {
    return a / b
}

func Mod(a, b int) int { return a % b }

func Abs(a int) int { if a < 0 { return -a }; return a }
