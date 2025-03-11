// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */
package matrix

import (
	"bytes"
	"log"

	nlp "pixai/neural_network/natural_language_processing"
)

type Rune [][]rune

type RuneMatrix struct{}

// rune matrix handling
func (m *RuneMatrix) Rune(mat Matrix32, input string) Rune {
	input_val := []byte(input)
	matrix := make([][]rune, len(mat))

	// splitting the byte
	output := bytes.Split(input_val, []byte("\n"))
	for _, value := range output {
		if len(value) == 0 {
			continue
		}

		// row encoding
		row := make([]rune, 0)
		for _, val := range string(input_val) {
			row = append(row, val)
		}

		// appending to the rune
		matrix = append(matrix, row)
	}

	// nil check
	if matrix != nil {
		return matrix
	} else {
		log.Println("Failure to generate rune")
	}

	return nil
}

// conversion
func (m *RuneMatrix) RuneToMatrix32(mat Rune) Matrix32 {
	// matrix
	mat32 := make(Matrix32, len(mat))
	mat32_1 := make(Matrix32, len(mat32))

	// appending and rune
	for i := range mat32 {
		mat[i] = append(mat[i], rune(len(mat32[i][:])))
		for x := range mat {
			mat32[x] = append(mat32[x], float32(len(mat)))
			if len(mat32) != 0 {
				mat32 = append(mat32, mat32_1[i])
			}

			if len(mat32) == 0 {
				break
			}
		}
	}

	if len(mat32) != 0 {
		return mat32
	}

	return nil
}

// matrix32 to rune
func (m *RuneMatrix) RuneConvert(mat Matrix32) Rune {

	// matrix32 to rune
	mat1 := make(Rune, len(mat))
	for i := range mat {
		for _, x := range mat[i] {
			if len(mat1) == 0 {
				break
			}

			if len(mat1) != 0 {
				mat1[i] = append(mat1[i], rune(x))
			}
		}
	}

	if len(mat1) != 0 {
		return mat1
	}

	return nil
}

// decoding the output
func (m *RuneMatrix) Decoding(mat Rune, input string) string {
	var output string
	input_val := []byte(input)

	// itteration
	for i := range mat {

		// grabbing byte value
		for _, x := range input_val {

			// concatinate
			output += string(rune(x))

			// empty check
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

func (m *RuneMatrix) Context(mat Rune, val []string, input string) string {
	srt := nlp.Conversion{}

	var output string
	if srt.StringCheck(val, input) {
		value := srt.ArrCheck(val, input)

		input = value

		// input val
		input_val := []byte(input)

		// itteration
		for i := range input_val {
			for _, x := range input_val {
				for l := range mat[i] {
					output += string(rune(x))

					// empty check
					if len(output) == 0 {
						log.Println("Decoding failure", l)
						break
					}
				}
			}

			break
		}

		if len(value) != 0 {
			return value
		}

		log.Println("Len: ", len(value))
	}

	return ""
}

// decoding the output
func (m *RuneMatrix) DecodingContext(mat Rune, input string) string {
	conv := nlp.Words{}

	var output string

	// GPE
	val := m.Context(mat, conv.Words().GPE, input)
	if len(val) != 0 {
		output = val
	}

	// stopwords
	val_1 := m.Context(mat, conv.Words().Stopwords, input)
	if len(val_1) != 0 {
		output = val_1
	}

	// verbs
	val_2 := m.Context(mat, conv.Words().Verbs, input)
	if len(val_2) != 0 {
		output = val_2
	}

	// nouns
	val_3 := m.Context(mat, conv.Words().Nouns, input)
	if len(val_3) != 0 {
		output = val_3
	}

	if len(output) != 0 {
		return output
	}

	return ""
}
