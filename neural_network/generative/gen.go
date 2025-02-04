// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// gen.go
package generative

// calling variables
import (
	"fmt"
	"math/rand"
	information "pixai/neural_network"
	natural "pixai/neural_network/natural_language_processing"
)

// naming vars
var (
	tokens    = natural.NLP{}           // tokens
	enum      = natural.Enumerate{}     // enumeration
	conv      = natural.Conversion{}    // Conversion variables
	variables = information.Variables{} // Variables for the neural_network
	weights   = information.Weights{}   // Randomness and weights
	neurons   = information.Layers{}    // Input
)

// array
type StringArray map[string]int

// struct for organization and global control
type Generative struct {

	// Markov
	frequency map[string]StringArray
	order     int
}

func (Generative) NewChain(order int) *Generative {
	chain := Generative{order: order, frequency: make(map[string]StringArray)}
	return &chain
}

// chain array
func (g *Generative) ChainArray(val string, num int) []string {
	value := make([]string, num)
	for x := range value {
		value[x] = val
	}

	if value != nil {
		return value
	}

	return nil
}

// concatinating (ie: Combining)
func (g *Generative) Concatinate(input1, input2 []string) []string {

	combined := make([]string, len(input1), len(input2))
	for i := range combined {
		if i > 0 {
			combined = append(combined, append(input1, input2...)...)
		}

		if i == 0 {
			combined[i] = combined[0]
		}
	}

	if combined != nil {
		return combined
	}

	return nil
}

// split
func (g *Generative) Splitting(input string) string {
	split := tokens.Document(input)

	words := make(map[string]int)
	for _, x := range split {
		_, word := words[x]
		for word {
			if len(x) == 0 {
				break
			}

			if len(x) > 0 {
				return x
			}
		}
	}

	return ""
}

func (g *Generative) Enum(input string) string {
	enumerate := enum.Enumeration(input)
	var x string
	for _, i := range enumerate {
		x = i
	}

	if len(x) != 0 {
		return x
	}

	return ""
}

// adding the markov chains
//                       int, string, []string
/* Example: chain.Adding(2, input, extracted_input) */
func (chain *Generative) Adding(n int, input string) {
	order := rand.Intn(n)

	// arrays
	s_tokens := chain.ChainArray(input, order)
	e_tokens := chain.ChainArray(input, order)

	// combination
	combined := chain.Concatinate(s_tokens, e_tokens)
	sequence := make([]string, len(combined))
	for i := range sequence {
		if sequence != nil {
			sequence = append(sequence, combined...)
		}

		if sequence == nil {
			sequence[i] = sequence[0]
		}
	}
}

func (g *Generative) MarkovChains(input string) []string {
	g.Adding(10, input)

	return nil
}

// initializing
func (g *Generative) GenerativeInit() error {
	return nil
}

// closing
func (g *Generative) close() {
	if g != nil {
		fmt.Println("Generation cleared from memory.")
	}
}

// close caller
func (g *Generative) Close() {
	g.close()
}
