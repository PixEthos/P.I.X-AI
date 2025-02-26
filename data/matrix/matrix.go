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
	"bytes"
	"log"
)

type Matrix32 [][]float32
type Rune [][]rune

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
	for i := range len(mat) {
		for x := range mat[i] {
			if len(mat1) > i && len(mat1[i]) > x {
				output[i] = append(output[i], mat1[i][x]-mat[i][x])
			} else if len(mat) > i && len(mat[i]) > x {
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

// rune matrix handling
func (m *Matrix) Rune(mat Matrix32, input string) Rune {
	input_val := []byte(input)
	matrix := make([][]rune, len(mat))
	for _, value := range bytes.Split(input_val, []byte("\n")) {
		if len(value) == 0 {
			continue
		}

		row := make([]rune, 0)
		for _, val := range string(input_val) {
			row = append(row, val)
		}

		matrix = append(matrix, row)
	}

	if matrix != nil {
		return matrix
	} else {
		log.Println("Failure to generate rune")
	}

	return nil
}

func (m *Matrix) Decoding(mat Rune, input string) string {
	var output string
	input_val := []byte(input)
	for i := range mat {
		for _, x := range input_val {
			output += string(rune(x))
			if len(output) == 0 {
				log.Println("Decoding failure", i)
				break
			}
		}
	}

	if len(output) != 0 {
		return output
	}

	return ""
}

/* The 32bit, and 64bit are split for good reasons - one of them being safety.

 I decided to use append in this case - especially for a heap of the framework I am building for this
 which as a result had lead me down rabbitholes of poorly implemented code. Or even reading downright
 hilarious implementation: "I need to use goroutines for matrix integration, for school"

 The fuck? Whomever is the actual professor that requires that; ask them why simplicity isn't
 the goal, and why complexity is the ideal.

 Regardless; the implementation I kept finding was the following:
 'mat[][] = mat1[][] + mat2[][]'
 For base arrays, this is fine depending on the context. But these are often not even allocated correctly
 you actually get basic results where the arrays are already standard within the algorithms provided.

 So what happens is you are ending up with memory allocation issues, and sometimes irritating troubleshooting
 and debugging - often leading to irritation. So, how about append?

 Well, I decided to start using append for more complex work within the algorithm - this is the result
 of me requiring dynamic allocation. Yes, append leads to more memory being allocated, but that isn't the sole
 reason - it's a bit safer actually. Now, let's say I implement this with goroutines - well, I can.

 But where exactly is the line between 'can' and 'should'? Well, I 'should' do so eventually.

 The main reason is for concurrency. So something like this:
 func (m* Matrix_Struct) Matrix32bit(mat Matrix32, output chan [][]float32)
 Than the caller:
 func (m* Matrix_Struct) Matrix32bitCaller(mat Matrix32, output chan [][]float32) Matrix32 {
	go m.Matrix32bit(mat, output)
	// than add in the activation here
 }

 Though, why don't I do this? Because it's easier for me to develop with simplicity, 'til I need the
 complexity.*/
