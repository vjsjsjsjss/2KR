package main

import (
    "testing"
    "2KR/add_sub"
    "2KR/mul_div"
)

func TestAdd(t *testing.T) {
    res := add_sub.Add(2, 3)
    if res != 5 {
        t.Errorf("Ожидалось 5, получено %f", res)
    }
}

func TestSub(t *testing.T) {
    res := add_sub.Sub(5, 2)
    if res != 3 {
        t.Errorf("Ожидалось 3, получено %f", res)
    }
}

func TestMul(t *testing.T) {
    res := mul_div.Mul(3, 4)
    if res != 12 {
        t.Errorf("Ожидалось 12, получено %f", res)
    }
}

func TestDivNormal(t *testing.T) {
    res, err := mul_div.Div(10, 2)
    if err != nil {
        t.Errorf("Ошибка при делении: %v", err)
    }
    if res != 5 {
        t.Errorf("Ожидалось 5, получено %f", res)
    }
}

func TestDivByZero(t *testing.T) {
    _, err := mul_div.Div(5, 0)
    if err == nil {
        t.Error("Ожидалась ошибка при делении на ноль, но её не было")
    }
}
