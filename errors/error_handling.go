// I use GLP2

/* Copyright (C) 2024 PixEthos */

// error_handling.go
package error_handling

import (
	"math"
	"errors"
)

type Math_Errors struct {
	Inf float64
	NaN float64
	Zero int
	Val int
}

func Maths() *Math_Errors {
	m := &Math_Errors{}
	return m
}

// common math errors
func (m* Math_Errors) Common() error {

	if m.Inf == math.Inf(1) {
		return errors.New("Infinite")
	}

	if m.NaN == math.NaN() {
		return errors.New("NaN values")
	}

	if m.Zero != 0 {
		return errors.New("Divided by zero")
	}

	if m.Val < 0 {
		return errors.New("Negative value")
	}

	if math.IsNaN(m.NaN) == true {
		return errors.New("Values are NaN")
	}

	if math.IsInf(m.Inf, 0) == true {
		return errors.New("Values are infinite")
	}

	return nil
}
