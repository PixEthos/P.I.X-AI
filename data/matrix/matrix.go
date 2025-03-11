// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

/* This is just the matrix handling of the neural_network - I did this as a way to handle the data more
carefully, and accurately.

A lot of the information I read online is genuinely pretty okay - but the way they handled it is _years_
out of date. Leading me to have to build this alone. This is just trial and error regardless for the
most part. I also didn't want to rely on preexisting tooling.

Just like the rest of this algorithm - I am not going to use a thirdparty library for it.

They often add a fuckload of overhead, and what can happen is --especially if it isn't implemented right--
is access violations - and in a lot of cases, specifically with NLP, Matrix, or even standard handling
you can end up coming across a literal issue with the framework - causing you to need to work around
it, leading to endless hours of debugging, or waiting.

This inevitably lead me to just working on things alone. Not working with a framework, which who the hell
knows whether or not the genuine code itself is updated to the current standard/version. */
// matrix.go
package matrix

import (
	"log"
)

type Matrix32 [][]float32

type Matrix struct{}

// creating 32bit 2D arrays
func (m *Matrix) Matrix32bit(mat Matrix32) Matrix32 {

	// matrix for adding the values safely
	for i := range len(mat) {

		// redeclaring
		additiveMat := make([][]float32, len(mat[i]))
		for x := range mat[i] {
			// value adding
			if len(mat[i]) <= 0 {
				additiveMat[i] = append(additiveMat[i], mat[i][x])
				additiveMat[i] = append(additiveMat[i], mat[i]...)
			}

			for l := range additiveMat {
				for k := range additiveMat[l] {

					// checks in place for memory safety
					if len(additiveMat[l]) == 0 {
						additiveMat[l] = append(additiveMat[l], mat[0][0])
					}

					// checking to add values
					if len(additiveMat[l]) <= 0 {
						mat[l] = append(mat[l], float32(len(additiveMat[:l][:k])))
					}
					// adding values
					mat[l] = append(mat[l], additiveMat[l][k])

					if len(additiveMat[l]) <= i {
						// adding to the called matrix
						mat[l] = append(mat[l], additiveMat[l]...)
					}
				}
			}
		}
	}

	if len(mat) != 0 {
		return mat
	} else {
		log.Println("Failure to create matrix")
	}

	return nil
}

// Adding 32bit 2D arrays
func (m *Matrix) Matrix32Addition(mat, mat1 Matrix32) Matrix32 {
	if mat == nil || mat1 == nil {
		return nil
	}

	output := make([][]float32, len(mat))
	for i := range mat {
		for x := range mat[i] {
			if len(mat1) > i && len(mat1[i]) > x {
				output[i] = append(output[i], mat1[i][x]+mat[i][x])
			} else {
				output[i] = append(output[i], mat[i][x])
			}
		}
	}

	if len(output) != 0 {
		return output
	} else {
		log.Println("Failure to add matrix")
	}

	return nil
}

// Adding 32bit 2D arrays
func (m *Matrix) Float32Addition(mat, mat1 Matrix32) float64 {
	if mat == nil || mat1 == nil {
		return 0
	}

	var output float64
	for i := range mat {
		for x := range mat[i] {
			if len(mat1) > i && len(mat1[i]) > x {
				output += float64(mat1[i][x] + mat[i][x])
			} else {
				output += float64(mat[i][x])
			}
		}
	}

	if output != 0 {
		return output
	} else {
		log.Println("Failure to add matrix")
	}

	return 0
}

// subtracting 32bit 2D arrays
func (m *Matrix) Matrix32Subtraction(mat, mat1 Matrix32) Matrix32 {
	if mat == nil || mat1 == nil {
		return nil
	}

	output := make([][]float32, len(mat))
	for i := range mat {
		for x := range mat[0] {
			if len(mat) > len(mat1) {
				output[i] = append(output[i], mat1[i][x]-mat[i][x])
			}

			if len(mat1) > len(mat) {
				output[i] = append(output[i], mat[i][x]-mat1[i][x])
			}
		}
	}

	if len(output) != 0 {
		return output
	} else {
		log.Println("Failure to subtract matrix")
	}

	return nil
}

// dividing 32bit 2D arrays
func (m *Matrix) Matrix32Divide(mat, mat1 Matrix32) Matrix32 {
	if mat == nil || mat1 == nil {
		return nil
	}

	output := make([][]float32, len(mat1))
	for i := range mat {
		output = make([][]float32, len(mat))
		for j := range mat1[i] {
			for k := range mat1 {
				if len(mat1) > i && len(mat1[i]) > j {
					output[i][j] += mat[i][k] / mat1[k][j]
				} else if len(mat) > i && len(mat[i]) > j {
					output[i][j] += mat1[i][k] / mat[k][j]
				}
			}
		}
	}

	if len(output) != 0 {
		return output
	} else {
		log.Println("Failure to divide matrix")
	}

	return nil
}

// multiplying 32bit 2D arrays
func (m *Matrix) Matrix32Multiply(mat, mat1 Matrix32) Matrix32 {
	if mat == nil || mat1 == nil {
		return nil
	}

	output := make([][]float32, len(mat1))
	for i := range mat {
		output = make([][]float32, len(mat))
		for j := range mat1[i] {
			for k := range mat1 {
				if len(mat1) > i && len(mat1[i]) > j {
					output[i][j] += mat[i][k] * mat1[k][j]
				} else {
					output[i][j] += mat[i][k]
				}
			}
		}
	}

	if len(output) != 0 {
		return output
	} else {
		log.Println("Failure to multiply matrix")
	}

	return nil
}
