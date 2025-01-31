// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

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
