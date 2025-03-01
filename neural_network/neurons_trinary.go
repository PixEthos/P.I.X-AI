// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// neurons_trinary.go
package neuralnet

import (
	"math/rand"

	"pixai/data/matrix"
)

// Trinary set

// neuron internal values
func (n *Neurons) values_trinary() *NeuronInternals {
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
func (n *Neurons) encapsulated_trinary(count uint32) matrix.Matrix32 {
	// encapsuled matrix
	encap := make(matrix.Matrix32, count)

	// internals
	ni := n.values_trinary()
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

// gru layering
func (n *Neurons) GRU_trinary(input matrix.Matrix32, x string) matrix.Matrix32 {
	l := Layers{}

	gru_sec := l.GRU_activation(100, 10, input, "float", x)
	if gru_sec != nil {
		return gru_sec
	}

	return nil
}

// processing
func (n *Neurons) processed_trinary(input matrix.Matrix32, count uint32, val float32) matrix.Matrix32 {
	mat32 := matrix.Matrix{}

	// appending the neurons and neuron internals
	neuron_internal := n.encapsulated_trinary(count)
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

// context holder
func (n *Neurons) Gru_processed_trinary(val1 float64, input matrix.Matrix32, con string) (matrix.Matrix32) {
	mat32 := matrix.Matrix32{{float32(val1)}}
	mat32 = append(mat32, input...)
	output := n.GRU_trinary(mat32, con)

	if output != nil {
		return output
	}

	return nil
}
