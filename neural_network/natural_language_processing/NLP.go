// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// NLP.go
package naturallanguageprocessing

import (
	// standard

	"fmt"
	"log"
	"os"
	"strconv"

	// AI

	encode "pixai/data/encoding"
	// ui
)

// Why add natural_language_processing to a games AI?
// Why not?

type NLP struct {
	valStr string
	Input  string
}

type Input struct{}

func NaturalLanguagProcessing(input string) *NLP {
	nlp := &NLP{
		Input: input,
	}
	return nlp
}

/* Machine learning isn't difficult to put into practice
it's only difficult to read the 40 page thesis on what it is
with obfuscated and complicated wording to protect copyright.

Not even kidding.*/

// calling errors up the stack, as usual
func (nlp *NLP) NLPErrors(input, output string, length int32) error {
	var err error

	Valstr, err := strconv.Atoi(nlp.valStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		log.Println(err)
		fmt.Println(Valstr)
		return err
	}

	return err
}

// scanning input
func (nlp *NLP) Scanner(val string) string {
	m := Match{}
	c := Conversion{}
	encode := encode.Encoded{}

	value := NaturalLanguagProcessing(val)

	// tokens splitting
	split := nlp.SplitTokens(val)
	in := c.ArraytoString(split)

	// len values
	sp := nlp.Tokens(len(split))

	// input matching
	m.Matching(in)
	match := m.MatchingLength(in)

	// output
	fmt.Printf("input tokens: %d\n", len(sp))
	fmt.Println("matched: ", match)

	if len(in) < 128 {
		value.Input = in
		log.Println("Input length: ", len(in))

		// byte encoding
		bit := make([]byte, 256)
		encode.Encode(in, bit)

		if len(in) != 0 {
			return in
		}
	} else {
		log.Println("Data length is more than 256")
		return ""
	}

	return ""
}

// init
func (nlp *NLP) NLPinit(val string) (string, error) {
	var err error

	output := nlp.Scanner(val)

	if len(output) != 0 {
		return output, nil
	}

	if len(output) == 0 {
		log.Println(err)
		return "", err
	}

	return "", err
}

func (nlp *NLP) close() {
	if nlp != nil {
		log.Println("NLP cleared from memory")
		return
	}
}

func (nlp *NLP) Close() {
	nlp.close()
}
