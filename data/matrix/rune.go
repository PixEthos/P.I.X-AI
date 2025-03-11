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

// decoding the output
func (m *RuneMatrix) DecodingContext(mat Rune, input string) string {
	conv := nlp.Words{}
	srt := nlp.Conversion{}

	var output string

	// GPE
	if srt.StringCheck(conv.Words().GPE, input) {
		gpe := srt.ArrCheck(conv.Words().GPE, input)

		input = gpe

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

		log.Println("GPE len: ", len(gpe))
	}

	// stopwords
	if srt.StringCheck(conv.Words().Stopwords, input) {
		stop := srt.ArrCheck(conv.Words().Stopwords, input)

		input = stop

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

		log.Println("Stopword len: ", len(stop))
	}

	// verbs
	if srt.StringCheck(conv.Words().Verbs, input) {
		verb := srt.ArrCheck(conv.Words().Verbs, input)

		input = verb

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

		log.Println("Verb len: ", len(verb))
	}

	// nouns
	if srt.StringCheck(conv.Words().Nouns, input) {
		noun := srt.ArrCheck(conv.Words().Nouns, input)

		input = noun

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

		log.Println("Noun len: ", len(noun))

	}

	if len(output) != 0 {
		return output
	}

	return ""
}
