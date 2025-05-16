package mul_div

import "errors"

func Mul(a, b float64) float64 {
    return a * b
}

func Div(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}
