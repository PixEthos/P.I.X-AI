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
	"strconv"
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
	return 1.0 / (1.0 + math.Exp(-probability))
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

func (e *Enumerate) Enumerate(input []string, val int) *Enumerate {
	enum := Enumerate{
		num: val,
		val: input,
	}

	return &enum
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
func (e *Enumerate) EnumerationDocument(input string) []Enumerate {
	var count int
	nlp := NLP{}
	split := nlp.Document(input)
	enum := make([]Enumerate, len(split))

	for count = 0; count < len(split); count++ {
		caller := e.Enumerate(split, count)
		enum[count] = Enumerate{
			num: caller.num,
			val: caller.val,
		}
	}

	return enum
}

// enumerate split tokens
func (e *Enumerate) EnumerationSplit(input string) []Enumerate {
	var count int
	nlp := NLP{}
	split := nlp.SplitTokens(input)
	enum := make([]Enumerate, len(split))

	for count = 0; count < len(split); count++ {
		caller := e.Enumerate(split, count)
		enum[count] = Enumerate{
			num: caller.num,
			val: caller.val,
		}
	}

	return enum
}

// converting to string
func (e *Enumerate) EnumerationItoa(input string) string {
	nlp := NLP{}
	split := nlp.SplitTokens(input)
	caller := e.Enumerate(split, len(split))
	value := caller.num
	word := caller.val

	var val string
	for i := range word {
		word = append(word, input)

		val = word[i]

		for i := 0; i < len(word); i++ {
			x := strconv.Itoa(value)
			val = x
		}
	}

	if len(val) != 0 {
		return val
	}

	return ""
}

// enumeration
func (e *Enumerate) Enumeration(input string) []string {
	nlp := NLP{}
	split := nlp.SplitTokens(input)
	word := e.EnumerationItoa(input)
	caller := e.Enumerate(split, len(input))

	enum := make([]string, len(caller.val))
	if len(word) > 0 {
		value := caller.num
		for value = 0; value < len(caller.val); value++ {
			enum[value] = caller.val[value]
		}

		if len(enum) != 0 {
			return enum
		}
	}

	return nil
}
