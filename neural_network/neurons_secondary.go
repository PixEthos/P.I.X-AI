// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// neurons_secondary.go
package neuralnet

import (
	"math/rand"

	"pixai/data/matrix"
)

// secondary set

// neuron internal values
func (n *Neurons) values_secondary() *NeuronInternals {
	// internals of each neuron
	Encapsuled := NeuronInternals{
		// bias
		bias:  rand.Float32() + 10,
		bias1: rand.Float32() + 5,
		bias2: rand.Float32() + 2,
		bias3: rand.Float32() + 9,
		bias4: rand.Float32() + 7,
		bias5: rand.Float32() + 1,

		// weights
		weight:  rand.Float32() + 10,
		weight1: rand.Float32() + 3,
		weight2: rand.Float32() + 5,
		weight3: rand.Float32() + 10,
		weight4: rand.Float32() + 8,
		weight5: rand.Float32() + 4,
	}

	return &Encapsuled
}

// neuron encapsulation
func (n *Neurons) encapsulated_secondary(count uint32) matrix.Matrix32 {
	// encapsuled matrix
	encap := make(matrix.Matrix32, count)

	// internals
	ni := n.values_secondary()
	//mat32 := matrix.Matrix{}

	// internals being appended/added to the main matrix
	for i := 0; i < int(count); i++ {
		internals := matrix.Matrix32{
			{ni.bias, ni.bias1}, {ni.bias2, ni.bias3}, {ni.bias4, ni.bias5},
			{ni.weight, ni.weight1}, {ni.weight2, ni.weight3}, {ni.weight4, ni.weight5},
		}
		encap = append(encap, internals...)

		if len(internals) >= len(encap) {
			break
		}
	}

	if encap != nil {
		return encap
	}

	return nil
}

// processing
func (n *Neurons) processed_secondary(input matrix.Matrix32, count uint32, val float32) matrix.Matrix32 {
	mat32 := matrix.Matrix{}

	// appending the neurons and neuron internals
	neuron_internal := n.encapsulated_secondary(count)
	neurons := mat32.Matrix32bit(neuron_internal)
	neuron_output := mat32.Matrix32bit(input)

	// output for the neurons
	output := mat32.Matrix32Addition(neuron_output, neurons)
	for i := range output {
		out := make([]float32, len(output))
		for x, diviser := range output[i] {
			out[x] = diviser / val
			output = append(output, out)
		}
	}

	if output != nil {
		return output
	}

	return nil
}

// context holding pieces
/* Explaination:
These functions are simple; all they do, is match a variable.

You take the input matrix, and than you have the input string
what happens next? It will be matched with a map.

See? That simple
*/
