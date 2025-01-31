// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// tokens.go
package naturallanguageprocessing

import (
	"math"
	"regexp"
	"strings"
)

type Enumerate struct {
	num int
	val []string
}

type Predictive struct{}

// Probability
func (pre *Predictive) Probability(frequency, total float64) float64 {
	probability := total / frequency
	return probability
}

// Entropy
func (pre *Predictive) Entropy(x, total float64) float64 {
	p := pre.Probability(x, total)
	return p * math.Log(p)
}

// Information
func (pre *Predictive) Information(x, total float64) float64 {
	information := 1 / pre.Entropy(x, total)
	return math.Log1p(information)
}

/* This is the tokenizing file, handling sentences, fielding, and conversions */

func (nlp *NLP) Trimming(input string) string {
	for i := range len(input) {
		x := strings.ToLower(input)
		v := strings.TrimLeft(input, ",!.")

		lower := make([]string, len(input))
		trim := make([]string, len(input))
		if i >= len(input) {
			lower = append(lower, x)
			trim = append(trim, v)
			lower[i] = trim[i]
			input = lower[i]

			return input
		}
	}

	return ""
}

func (nlp *NLP) SplitTokens(input string) []string {

	puncuation := `[\p{P}\d\W]`

	doc := strings.Split(input, `,?!`)
	reg := regexp.MustCompile(puncuation)
	split := reg.Split(input, -1)

	split = append(split, doc...)

	// fmt.Printf("%s\n", doc)

	//fmt.Println("Split tokens:", tokens)
	if split != nil {
		return split
	}

	return nil
}

// Document tokenizing
//
// What this does is allow me to split sentences and documents by punctuation. Think of how
// NLPTK splits these values. This is arguably a lot more powerful than working with a framework
// prebuilt with it.
//
// I was looking online for something akin to this, mostly due to the fact that looking at
// framework code all day was a frustrating experience. So, I basically decided to experiment.
//
// All this does is take a string input, and splits the values
func (nlp *NLP) Document(input string) []string {

	puncuation := `\p{P}\s`

	fielding := strings.Fields(input)
	reg := regexp.MustCompile(puncuation)
	count := make([]int, len(input))

	tokens := reg.Split(input, -1)

	// tokens
	for x := range fielding {
		for i := range count {
			if len(count) == 0 {
				tokens = append(tokens, tokens[x])
			}

			if len(count) <= i {
				tokens[x] = tokens[x+1]
			}
		}
	}

	if tokens != nil {
		return tokens
	}

	return nil
}

// token splitting values
func (nlp *NLP) Tokens(input int) []int {
	var num []int

	for n := 0; n < input; n++ {
		num = append(num, input)
		for i := range len(num) {
			num[i] = input
		}
	}

	if num != nil {
		return num
	}

	return nil
}

// enumerate doument
func (e *Enumerate) EnumerationDocument(input string) *Enumerate {
	var count int
	val := []string{}
	nlp := NLP{}
	doc := nlp.Document(input)

	for i, x := range doc {
		val = append(val, x)
		for count = range i {
			count++
		}
	}

	enumerate := Enumerate{
		num: count,
		val: val,
	}

	return &enumerate
}

// enumerate split tokens
func (e *Enumerate) EnumerationSplit(input string) *Enumerate {
	var count int
	val := []string{}
	nlp := NLP{}
	split := nlp.SplitTokens(input)

	for i, x := range split {
		val = append(val, x)
		for count = range i {
			count++
		}
	}

	enumerate := Enumerate{
		num: count,
		val: val,
	}

	return &enumerate
}
