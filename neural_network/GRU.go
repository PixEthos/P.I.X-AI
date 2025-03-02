// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// DATE OF FILE CREATION: January 18th, 2025

// GRU.go
package neuralnet

import (
	"log"
	"math/rand/v2"
	mat "pixai/data/matrix"
)

type Layers struct {
	number_of_neurons uint32
	number_of_layers  uint32
}

/* GRU_layering */
/* Explaination:
This is the GRU layering, you _will_ need to call in the output function to be able to get proper outputs
and layers for the GRU layers. All this does is give values to the l.number_of_neurons, and l.number_of_layers
within the algorithm. This makes it a lot easier to process the data.
*/
func (l *Layers) GRU_layering(neurons, layers uint32) (uint32, uint32) {
	l.number_of_neurons = neurons // number of neurons per layer
	l.number_of_layers = layers   // number of layers

	return l.number_of_layers, l.number_of_neurons
}

// neuron encapsulation
func (l *Layers) GRU_encapsulated(count uint32) mat.Matrix32 {
	// encapsuled matrix
	encap := make(mat.Matrix32, count)

	// bias
	var bias, bias1, bias2, bias3, bias4, bias5 float32
	bias = rand.Float32() + 3  // each of these uses its own unique number for the sake of simplicity
	bias1 = rand.Float32() + 1 // there is 6 biases, and 6 weights
	bias2 = rand.Float32() + 4
	bias3 = rand.Float32() + 7
	bias4 = rand.Float32() + 4
	bias5 = rand.Float32() + 5

	// weights
	var weight, weight1, weight2, weight3, weight4, weight5 float32
	weight = rand.Float32() + 9
	weight1 = rand.Float32() + 10
	weight2 = rand.Float32() + 6
	weight3 = rand.Float32() + 2
	weight4 = rand.Float32() + 4
	weight5 = rand.Float32() + 5

	// internals being appended/added to the main matrix
	for i := 0; i < int(count); i++ {
		internals := mat.Matrix32{
			{bias, bias1}, {bias2, bias3}, {bias4, bias5},
			{weight, weight1}, {weight2, weight3}, {weight4, weight5},
		}
		encap = append(encap, internals...)

		if len(encap) > len(internals) {
			break
		}
	}

	if encap != nil {
		return encap
	}

	return nil
}

// processing
func (l *Layers) GRU_processed(input mat.Matrix32, val float32) float64 {
	mat32 := mat.Matrix{}

	// appending the neurons and neuron internals
	neuron := l.GRU_encapsulated(l.number_of_neurons)

	// output for the neurons
	output := mat32.Float32Addition(input, neuron)

	if output != 0 {
		return output
	}

	return 0
}

// processing matrix
func (l *Layers) GRU_processed_matrix(input mat.Matrix32, val float32) mat.Matrix32 {
	mat32 := mat.Matrix{}

	// appending the neurons and neuron internals
	neuron := l.GRU_encapsulated(l.number_of_neurons)

	// output for the neurons
	output := mat32.Matrix32Addition(input, neuron)

	if output != nil {
		return output
	}

	return nil
}

/*
	Explaination:

Bellow here is going to be the layering; each are able to hold,
and process the data for the sake of recurrence.

GRU laying essentially just means I store recurrent data in separated neural_networks
this is a form a deep-learning, but it's just a recurrent neural_network.

The reason I went with a GRU-based system; is I wanted to make sure the generated outputs
were more or less similar to the last output. This is being added --alongside the already 3
neural_networks for multiprocessing-- for the sake of simplicity. I did plan on LGP integration
and I think at this point; if I do add the genetic portion to the model - it would probably have
a bit of a better chance of _learning_ from the recurrent and input methods.

Now, I personally didn't go with transformative methods due to hardware constraints;
I neither have resources for it, nor do I have direct access to a computational card.

Yes, I _do_ have a Radeon RX 6650 XT, and an RX 570 (the 570 I can use for the computational methods)
but that'd require me to rewrite the entire algotihm in a language that supports ROCm; and not many
overall data scientists actually even have access to an ROCm supported card. Let alone most users.

But there's a purpose to the madness; if I create an algorithm, that arguably doesn't need any
external hardware except for a CPU, and available resources; (ideally I'd say 4GiB of RAM)
than I'll build it with the intend of simplicity, not using a thirdparty library. Which is why
--as a result-- I am building the algorithm under a FOSS license. Although there is a wrapper for Go
I'd rather not waste my time learning a new library.
*/

// GRU layering; add in the number of neurons per set, and than adde in the number of layers
func (l *Layers) GRU_Layers() mat.Matrix32 {
	if l.number_of_layers != 0 {
		layers := make(mat.Matrix32, l.number_of_neurons)
		for i := 0; i < int(l.number_of_layers); i++ {
			layers = l.GRU_encapsulated(l.number_of_neurons)
		}

		if layers != nil {
			return layers
		}
	}

	return nil
}

// rune decoding
func (l *Layers) GRU_rune_decode(input string, val mat.Rune) string {
	Rune := mat.Matrix{}

	output := Rune.Decoding(val, input)
	if len(output) != 0 {
		return output
	}

	return ""
}

// rune processing
func (l *Layers) GRU_rune_variable(input string, val mat.Matrix32) mat.Rune {
	Rune := mat.Matrix{}

	output := Rune.Rune(val, input)
	if output != nil {
		return output
	}

	return nil
}

// layer processing; which adds in the layering, than the input for the processing
func (l *Layers) GRU_layer_processing(input mat.Matrix32, x string) mat.Matrix32 {
	w := Weights{}
	mat32 := mat.Matrix{}

	// weights
	weight := w.Weight(10)                       // weight randomization for division
	layers := l.GRU_Layers()                     // layers
	processed := l.GRU_processed(layers, weight) // processing of the inputs/values

	// matrix
	process_layers := mat32.Matrix32bit(layers)

	// values
	value := make([][]float32, len(process_layers)+int(processed))
	layers = append(layers, value...)

	// checking for nil
	if layers != nil {
		return layers
	}

	return nil
}

// layer processing; which adds in the layering, than the input for the processing
func (l *Layers) GRU_layer_processing_matrix(input mat.Matrix32, x string) mat.Matrix32 {
	w := Weights{}
	mat32 := mat.Matrix{}

	// weights
	weight := w.Weight(10)                             // weight randomization for division
	layers := l.GRU_Layers()                           // layers
	processed := l.GRU_processed_matrix(input, weight) // processing of the inputs/values

	// matrix
	process_layers := mat32.Matrix32bit(processed)
	input_layer := mat32.Matrix32bit(layers)

	// processing
	processing_layers := make([][]float32, len(process_layers))
	processing_layers = append(processing_layers, process_layers...)
	added_layers := mat32.Matrix32Addition(input_layer, process_layers)

	// checking for nil
	if added_layers != nil {
		return added_layers
	}

	return nil
}

// output layering (combined)

/*
This is the GRU layer process caller; basically just type on the string value either "matrix", or "float"
for singular, or multireturn for debugging.
*/
func (l *Layers) GRU_layer_output(input mat.Matrix32, value, x string) mat.Matrix32 {

	if value == "matrix" {
		output_matrix := l.GRU_layer_processing_matrix(input, x)

		if output_matrix != nil {
			return output_matrix
		}
	} else if value == "float64" {
		output := l.GRU_layer_processing(input, x)

		if output != nil {
			return output
		}
	}

	return nil
}

// output layering (combined)
func (l *Layers) GRU_layer_output_float(input mat.Matrix32, x string) float64 {
	mat32 := mat.Matrix{}
	gru_output := l.GRU_layer_processing(input, x)
	output_mat := mat32.Matrix32bit(gru_output)
	output := mat32.Float32Addition(output_mat, input)

	if output != 0 {
		return output
	}

	return 0
}

// sigmoid
func (l *Layers) GRU_sigmoid(input mat.Matrix32, val, x string) float64 {
	var sigmoid float64

	v := Variables{}

	// essentially, this is adding the output values to a sigmoid
	output := l.GRU_layer_output(input, val, x)
	for i := range output {

		// making an output array
		out := make([]float32, len(output))
		for x, values := range output[i] {

			// adding output array value
			out[x] = values

			// sigmoid value
			sigmoid = v.Sigmoid(float64(out[x]))
		}
	}

	// logerithm of the primary sets of each
	primary := sigmoid

	return primary
}

// activator
/* Expalination:

This is the activator function for the GRU layering.
neurons, and layers first both are uint32 (example: 100, 20)
Input being the values you want processed (needing to be a Matrix32, or [][]float32)
value being the string, this needs to either be "float" or "matrix" for the processed output

Example: l.GRU_activation(100, 10, input, "float")
*/
func (l *Layers) GRU_activation(neurons, layers uint32, input mat.Matrix32, value, x string) mat.Matrix32 {
	l.GRU_layering(neurons, layers)
	output := l.GRU_layer_output(input, value, x)


	if output != nil {
		return output
	}

	return nil
}

func (l *Layers) close() {
	if l != nil {
		log.Println("GRU layer cleared from memory")
		return
	}
}
