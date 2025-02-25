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

	// std
	"fmt"
	"log"
	"math/rand"
	"strings"

	// local
	matrix "pixai/data/matrix"
	information "pixai/neural_network"
	natural "pixai/neural_network/natural_language_processing"
)

// It's time. We'll start over from zero with this V2, and entrust the future to the next generation.

// naming vars
var (
	// natural language processing
	tokens = natural.NLP{}        // tokens
	enum   = natural.Enumerate{}  // enumeration
	conv   = natural.Conversion{} // Conversion variables

	// data processing
	variables = information.Variables{} // Variables for the neural_network
	weights   = information.Weights{}   // Randomness and weights
	neurons   = information.Neurons{}   // Neuron groups

	// matrix handling
	mat32 = matrix.Matrix{} // Matrix
)

// struct for organization and global control
type Generative struct {

	// Markov
	frequency map[string][]string
	order     int
}

// prefix
type Prefix []string

// prefix joining
func (p Prefix) Join(input string) string {
	return strings.Join(p, input)
}

// prefix merging
func (p Prefix) Merge(input string) {

	// for loop
	for i := range p {
		if p != nil {

			// appending for a change
			p = make([]string, len(p[i:]))

			// merging inputs
			p[len(p)-1] = input
		}

		// nil check
		if i == 0 {
			p[i] = p[0]
		}
	}
}

// chain
func (Generative) Chain(order int) *Generative {
	chain := Generative{order: order, frequency: make(map[string][]string)}
	return &chain
}

// GRU
func (Generative) GRU_layers(length int, input string) float64 {

	// matrix
	matrix := matrix.Matrix32{{float32(length)}}
	variable := mat32.Matrix32bit(matrix)

	// GRU activation layers
	primary, secondary, trinary := neurons.GRUActivation(variable, input)
	endpoint := primary + secondary + trinary

	log.Println("GRU_primary: ", primary)
	log.Println("GRU_seconary: ", secondary)
	log.Println("GRU_trinary: ", trinary)

	return endpoint
}

// building
func (chain *Generative) Build(input string) {
	p := make(Prefix, chain.order)

	for l := range input {
		if l != 0 {
			key := p.Join(input)
			chain.frequency[key] = append(chain.frequency[key], input)
			p.Merge(input)
		}

		if l == 0 {
			chain := chain.frequency[""]
			chain[l] = chain[0]
		}
	}
}

// chain array
func (g *Generative) ChainArray(val string, num int) []string {
	value := make([]string, num)
	for x := range value {
		value[x] = val
	}

	if len(value) != 0 {
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
	var_to_string := conv.ArraytoString(split)

	if len(var_to_string) != 0 {
		return var_to_string
	}

	return ""
}

func (g *Generative) Enum(input string) []string {
	enumerate := enum.Enumeration(input)
	fmt.Println(enumerate)
	if len(enumerate) != 0 {
		return enumerate
	}

	return nil
}

// adding the markov chains
//                       int, string, []string
/* Example: chain.Adding(2, input, extracted_input) */
func (chain *Generative) Adding(n int, input string) []string {
	order := rand.Intn(n)

	// arrays
	s_tokens := chain.ChainArray(input, order)
	e_tokens := chain.ChainArray(input, order)

	// combination
	combined := chain.Concatinate(s_tokens, e_tokens)
	if combined != nil {
		return combined
	}

	return nil
}

func (chain *Generative) MarkovChains(input string) []string {
	split := chain.Splitting(input)
	enums := chain.Enum(split)
	enum := conv.ArraytoString(enums)
	chain.GRU_layers(len(enum), split)
	add := chain.Adding(10, enum)

	for i := range add {
		if add != nil {
			add = append(add, enum)
		}

		if i == 0 {
			add[i] = add[0]
		}
	}

	if add != nil {
		return add
	}

	return nil
}

func (g *Generative) Convert(input string) string {
	val := g.MarkovChains(input)
	conv := conv.ArraytoString(val)

	if len(conv) != 0 {
		return conv
	}

	return ""
}

// initializing
func (g *Generative) GenerativeInit(val string) (string, error) {
	p := Prefix{}
	value := g.Convert(val)
	join := p.Join(value)
	fmt.Println("joining: ", join)

	return value, nil
}

// closing
func (g *Generative) close() {
	if g != nil {
		log.Println("Generation cleared from memory.")
	}
}

// close caller
func (g *Generative) Close() {
	g.close()
}
