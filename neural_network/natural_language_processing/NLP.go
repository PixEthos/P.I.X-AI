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
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	// AI
	"pixai/data/cache"
)

// Why add natural_language_processing to a games AI?
// Why not?

type NLP struct {
	valStr string
}

type Input struct{}

func NaturalLanguagProcessing() *NLP {
	nlp := &NLP{}
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
		fmt.Println(Valstr)
		return err
	}

	return err
}

func (nlp *NLP) Scanner(bit []byte) string {
	m := Match{}
	c := Conversion{}

	// buffer io
	reader := bufio.NewReader(os.Stdin)
	val := bufio.NewScanner(reader)

	// scanning inside the loop
	if val.Scan() {

		// tokens splitting
		split := nlp.SplitTokens(val.Text())
		in := c.ArraytoString(split)

		// len values
		sp := nlp.Tokens(len(split))

		// input matching
		m.Matching(in)
		match := m.MatchingLength(in)

		// output
		fmt.Printf("input tokens: %d\n", len(sp))
		fmt.Println("matched: ", match)

		if len(in) > 256 {
			return ""
		}

		if len(in) < 256 {
		// byte encoding
		source := make([]byte, len(in))
		for {
			reader.Read(bit)
			hex.Encode(bit, source)
			hex.Dump(bit)

			break
		}
		}

		if len(in) != 0 {
			return in
		}
	}

	return ""
}

// init
func (nlp *NLP) NLPinit() (string, error) {
	var err error
	cachestring := cache.RegCache[string, string]()
	bit := make([]byte, 256)
	output := nlp.Scanner(bit)

	cachestring.SetReg(output, "output")
	if len(output) != 0 {
		return output, nil
	}

	return "", err
}

func (nlp *NLP) close() {
	if nlp != nil {
		fmt.Println("NLP cleared from memory")
		return
	}
}

func (nlp *NLP) Close() {
	nlp.close()
}
