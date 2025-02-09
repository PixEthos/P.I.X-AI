// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// neural_net.go
package neuralnet

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	cache "pixai/data/cache"
	matrix "pixai/data/matrix"
	error_handling "pixai/errors"
	lang "pixai/neural_network/natural_language_processing"
)

var (
	predict = lang.Predictive{}
	nat     = lang.NLP{}
)

/* AI documentation: "Take the dinglehopper, and add it to the whoopdiedoo"
Me: "Uhhhhhhhhhhhh how does that translate into code? 'Cause I do not understands"
AI documentation: "It doesn't, it's just nonsense" */

type NotRecognized struct {
	NaN  float64
	Inf  float64
	Zero int
}

func MathErrors() error_handling.Math_Errors {
	m := error_handling.Math_Errors{}
	return m
}

// error caller
func (nr *NotRecognized) Values() error {
	var err error

	w := Weights{}
	v := Variables{}
	me := error_handling.Math_Errors{}

	var j, k int32
	var x float64

	if math.IsNaN(float64(w.Weight(float32(j)))) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return me.Common()
	}

	if math.IsNaN(float64(w.Bias(float32(k)))) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return me.Common()
	}

	if math.IsNaN(float64(w.biasdouble_slice)) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return me.Common()
	}

	if math.IsNaN(float64(w.weightdouble_slice)) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return me.Common()
	}

	if math.IsNaN(v.Sigmoid(x)) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return me.Common()
	}

	if math.IsNaN(v.Log(x)) {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return me.Common()
	}

	return err
}

func (n *Neurons) GRUActivation(input matrix.Matrix32, in string) {
	l := Layers{}

	gru_pri := n.gru_processed(input, in)
	primary := l.GRU_sigmoid(gru_pri, "float")
	fmt.Println("GRU_primary_set: ", primary)

	gru_sec := n.gru_processed_secondary(input, in)
	secondary := l.GRU_sigmoid(gru_sec, "float")
	fmt.Println("GRU_secondary_set: ", secondary)

	gru_tri := n.gru_processed_trinary(input, in)
	trinary := l.GRU_sigmoid(gru_tri, "float")
	fmt.Println("GRU_trinary_set: ", trinary)
}

func (n *Neurons) NeuronActivation(input matrix.Matrix32, total float32) {
	v := Variables{}

	n.encapsulated(n.Neuron_Count)
	n.encapsulated_secondary(400)
	n.encapsulated_trinary(400)
	output_1, output_2, output_3 := n.output(input, n.Neuron_Count, total)
	n.combined_outputs(input, n.Neuron_Count, total)
	accuracy, sec, tri := n.neuron_sigmoid(input, n.Neuron_Count, total)

	primary := v.SigmoidDerivative(accuracy) // derivative of the curve
	secondary := v.SigmoidDerivative(sec)
	trinary := v.SigmoidDerivative(tri)

	fmt.Println("Accuracy outputs:")
	fmt.Println("Primary accuracy: ", primary,
		"\nSecondary accuracy: ", secondary,
		"\nTrinary accuracy: ", trinary)

	fmt.Println("Output_1: ", output_1,
		"\nOutput_2: ", output_2,
		"\nOutput_3: ", output_3)
}

// learning
func (n *Neurons) NetworkLearning(in string) (float64, error) {
	// structs
	w := Weights{}
	v := Variables{}
	c := Confidence{}
	nr := NotRecognized{}

	// cache types
	f64cache := cache.Newf64[string, float64]()
	matrixCache := cache.RegCache[string, matrix.Matrix32]()

	// rand
	i := rand.Int31n(10000) + 190/5

	// vars
	var x, weight, bias int32
	var vals, output float32
	var out, accuracy float64

	// declarations
	n.Neuron_Count = 2000
	bias_slice := w.Bias(2)
	weight_slice := w.Weight(10)

	// redefines
	weight = int32(bias_slice)
	bias = int32(weight_slice)
	w.weightdouble_slice = rand.Float32()
	w.biasdouble_slice = rand.Float32()

	// in.eachNeuron(*values, uint32(n.Neuron_Biases), n.Neuron_Slice)

	// weighted learning
	for x = 0; x < i; x++ {
		w.weighted = bias + weight
		w.double = w.biasdouble_slice * w.weightdouble_slice
		vals = (w.double) / float32(w.weighted)
		c.linear_accuracy = float64(vals)
	}

	// testing values
	input := rand.Float32() + 200 + float32(i)/2

	Biasn := rand.Int31n(10)
	Weightn := rand.Int31n(2)

	// derivative
	Bias, Weight := w.Derivative(Biasn, Weightn)

	// values for the corrections
	var weighted, total float32
	var correct, sigmoid float64

	// learning loop
	for x = 0; x < i; x++ {
		weighted = input / (Weight + Bias + vals)
		sigmoid = v.Sigmoid(float64(weighted))
		output = float32(sigmoid) * float32(c.linear_accuracy)
	}

	// outputs
	correct = float64(output) + sigmoid
	total = output + weighted + float32(sigmoid) + vals

	mat32 := matrix.Matrix{}
	input_array := matrix.Matrix32{{total + float32(correct) + float32(w.output) + output}}
	inputs := mat32.Matrix32bit(input_array)

	// neuron calcs
	n.NeuronActivation(inputs, total)

	// GRU
	n.GRUActivation(inputs, in)

	// calculations
	w.prediction = correct / float64(total)          // predictive
	accuracy = v.Sigmoid(w.Accuracy(total, correct)) // weighted prediction
	w.output = float64(total) + correct              // neuron output

	// confidence logs
	c.output_accuracy = w.output
	c.derivative_accuracy = w.derivative
	c.prediction_accuracy = w.prediction

	// confidence caching
	f64cache.Setf64bit("confidence accuracy", accuracy)
	f64cache.Setf64bit("derivative accuracy", c.derivative_accuracy)
	f64cache.Setf64bit("prediction accuracy", c.prediction_accuracy)
	matrixCache.SetReg("neuron output", c.neuron_output)
	f64cache.Setf64bit("output total", c.output_accuracy)

	// NaN checking
	if math.IsNaN(w.derivative) || math.IsNaN(w.prediction) || math.IsNaN(w.output) || math.IsNaN(w.Accuracy(total, correct)) {
		return 0, nr.Values()
	}

	// output
	out = float64(total)

	return out, nil
}

// network
func (n *Neurons) NeuralNetwork(input string) error {

	output, err := n.NetworkLearning(input)
	if err != nil {
		return err
	}

	fmt.Println("output:", output)

	return err
}

// init
func (n *Neurons) NeuralNetworkInit(input string) error {
	var err error

	if err := n.NeuralNetwork(input); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}

	return err
}

// close
func (n *Neurons) close() {
	if n != nil {
		fmt.Println("NeuralNetwork cleared from memory")
		return
	}
}

// this is for being called outside the base function
func (n *Neurons) Close() {
	l := Layers{}

	l.close()
	n.close()
}
