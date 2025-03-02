// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// neurons.go
package neuralnet

import (
	"math/rand"

	"pixai/data/matrix"
)

/* Explaination:

The slices and matrix values are basically just calculating the neurons.
This did at one point house really bad practice on my part, where the neurons
themselves were just calling a dual-for loop for either or.

Due to that, and due to me needing to actually implement the NLP eventually
I created myself a matrix handler-file.

This means:
1. I can simplify my development process
2. I can use the values from the matrix.go source
3. If an issue arises, the 2D Arrays can be fixed more easily

Within this source; is just for the neuron handling. ie: matrices, bias, weights, and inputs
this allows for the algorithm itself to be usable and a bit more accurate.

Although this is the 17th of December, 2024 - I've been making solid progress so far on the development
and as such, that means, I will inevitably be able to release this sooner than I thought, so long as
nothing insane happens. */

// structs
type Neurons struct {
	Neuron_Slice  float32
	Neuron_Biases float32
	Neuron_Count  uint32
}

type NeuronInternals struct {
	bias, bias1, bias2,
	bias3, bias4, bias5 float32

	weight, weight1, weight2,
	weight3, weight4, weight5 float32
}

// neuron internal values
func (n *Neurons) values() *NeuronInternals {
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
func (n *Neurons) encapsulated(count uint32) matrix.Matrix32 {
	// encapsuled matrix
	encap := make(matrix.Matrix32, count)

	// internals
	ni := n.values()
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
func (n *Neurons) GRU_primary(input matrix.Matrix32, x string) matrix.Matrix32 {
	l := Layers{}

	gru_sec := l.GRU_activation(200, 30, input, "float", x)
	if gru_sec != nil {
		return gru_sec
	}

	return nil
}

// processing
func (n *Neurons) processed(input matrix.Matrix32, count uint32, val float32) matrix.Matrix32 {
	mat32 := matrix.Matrix{}

	// appending the neurons and neuron internals
	neuron_internal := n.encapsulated(count)
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

// output vals
func (n *Neurons) output(in matrix.Matrix32, count uint32, val float32) ([][]float32, [][]float32, [][]float32) {
	mat32 := matrix.Matrix{}

	// primary neural_network
	input := n.processed(in, count, val)
	result := mat32.Matrix32bit(input) // creating the matrix safely

	// second neural_network
	secondary_input := n.processed_secondary(in, 400, val)
	result_secondary := mat32.Matrix32bit(secondary_input)

	// third neural_network
	trinary_input := n.processed_trinary(in, 400, val)
	result_trinary := mat32.Matrix32bit(trinary_input)

	// nil check for returning matrix results
	if result != nil || result_secondary != nil || result_trinary != nil {
		return result, result_secondary, result_trinary
	}

	return nil, nil, nil
}

// same as the top, but combined
func (n *Neurons) combined_outputs(input matrix.Matrix32, count uint32, val float32) [][]float32 {
	mat32 := matrix.Matrix{}

	primary_input := n.processed(input, count, val)
	primary_result := mat32.Matrix32bit(primary_input)

	secondary_input := n.processed_secondary(input, 400, val)
	result_secondary := mat32.Matrix32bit(secondary_input)

	trinary_input := n.processed_trinary(input, 400, val)
	result_trinary := mat32.Matrix32bit(trinary_input)

	// adding the three together
	second_third := mat32.Matrix32Addition(result_secondary, result_trinary)
	output := mat32.Matrix32Addition(second_third, primary_result)

	if output != nil {
		return output
	}

	return nil
}

/*
The Sigmoid function is basically to give you an accuracy rating; it's weighted with the
logerithm, and derivative functions. Giving you an idea of how the AI is performing
at least, accuracy wise.
*/
func (n *Neurons) neuron_sigmoid(input matrix.Matrix32, count uint32, val float32) (float64, float64, float64) {
	var sigmoid, second, third float64

	v := Variables{}

	// essentially, this is adding the output values to a sigmoid
	output, secondary, trinary := n.output(input, count, val)
	for i := range output {

		// making an output array
		out := make([]float32, len(output))
		for x, values := range output[i] {

			// adding output array value
			out[x] = values + val

			// sigmoid value
			sigmoid = v.Sigmoid(float64(out[x]))
		}
	}

	// second
	for i := range secondary {
		sec := make([]float32, len(secondary))
		for x, values := range secondary[i] {
			sec[x] = values + val
			second = v.Sigmoid(float64(sec[x]))
		}
	}

	// third
	for i := range trinary {
		tri := make([]float32, len(trinary))
		for x, values := range trinary[i] {
			tri[x] = values + val
			third = v.Sigmoid(float64(tri[x]))
		}
	}

	// logerithm of the primary sets of each
	primary := v.Log(sigmoid)
	sec := v.Log(second)
	tri := v.Log(third)

	return primary, sec, tri
}

// input value for the generative pieces
func (n *Neurons) Input(input matrix.Matrix32, count uint32, val float32) ([][]float32, [][]float32, [][]float32) {
	return n.output(input, count, val)
}

// context holder
func (n *Neurons) Gru_processed(input matrix.Matrix32, con string) (matrix.Matrix32) {
	output := n.GRU_primary(input, con)

	if output != nil {
		return output
	}

	return nil
}
