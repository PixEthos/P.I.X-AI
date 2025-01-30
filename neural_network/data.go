// I use GPL2

/* Copyright (C) 2024 PixEthos */

// data.go
package neuralnet

import (
	"pixai/data/matrix"
	"math"
	"math/rand"
)

// weights
type Weights struct {
	// float32
	weightdouble_slice float32
	biasdouble_slice   float32
	double             float32

	// int32
	weight_slice int32
	bias_slice   int32
	weighted     int32

	// prediction float64
	prediction float64
	derivative float64
	output     float64
}

// confidence struct
type Confidence struct {
	prediction_accuracy float64
	derivative_accuracy float64
	linear_accuracy     float64
	output_accuracy     float64
	neuron_output       matrix.Matrix32
}

// Variables
type Variables struct {
	input int32
}

// values for the linear algebraic expressions
func (w *Weights) Epsilon() float32 {
	return 0.000001
}

func (w *Weights) Rate() float32 {
	return 0.00001
}

// 1 / 1 + the exponent of either eulers, or any value
func (v *Variables) Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func (v *Variables) SigmoidDerivative(x float64) float64 {
	return 1 - (v.Sigmoid(x)*2/v.Sigmoid(-x) - 2)
}

// the natural log of a constant, each log is equal to e (2.718...)
func (v *Variables) Log(x float64) float64 {
	return math.Log(x)
}

// Weights and biases are random values of 32bit floats subtracting 0.5
func (w *Weights) Weight(i float32) float32 {
	return rand.Float32() + i - 0.5
}

func (w *Weights) Bias(i float32) float32 {
	return rand.Float32() + i - 0.8
}

// Up here are the calculations of weight, bias, sigmoid, epsilon, and rate

// I actually designed the neurnal network based on --once again-- what I seen
// from ProgrammingRainbow; but as usual, I take inspiration, and change aspects
// especially where I see fit.

// Instead of my original Basic_AI, which let's be honest was a complete trashheap
// this one is using slices, and has many differences.

// I reduced the complexity, significantly, started to use Golang for it, and more.

// Now, I learn a lot from mistakes; and that's one of the reasons I am keeping
// Basic_AI up on my Github; because I want to show: "Yes, it's okay to make mistakes"
// I would have kept up with my development of Basic_AI, but I didn't.

// Too many memory errors and more, and honestly? Fuck that noise.

func (v *Variables) Inputs(i int32) int32 {

	v.input = i
	inputs := v.input
	for i = 0; i < inputs; i++ {
		inputs++
	}

	if inputs != 0 {
		return inputs
	}

	return 0
}

// accuracy
func (w *Weights) Accuracy(total float32, correct float64) float64 {
	v := Variables{}

	// fmt.Printf("Correct: %v, total: %v\n", correct, total)

	sum := float64(total) / correct

	return v.Log(sum) // accuracy of the model
}

// Cost
func (w *Weights) Cost(i, j int32) (int32, int32) {
	var resultBias, resultWeight int32

	for bias := range w.bias_slice {
		y := bias * rand.Int31n(5)
		z := y - bias
		resultBias += z * z
	}

	for weight := range w.weight_slice {
		y := weight * rand.Int31n(10)
		z := y - weight
		resultWeight += z * z
	}

	return resultBias, resultWeight
}

// Derivative
func (w *Weights) Derivative(bias_rand, weight_rand int32) (float32, float32) {
	var Bias, Weight float32

	bias, weight := w.Cost(bias_rand, weight_rand)

	w.biasdouble_slice = 10.0
	w.weightdouble_slice = 1.0

	Bias = float32(bias) - w.biasdouble_slice/float32(w.Epsilon())
	Bias *= float32(w.Rate())

	Weight = float32(weight) - w.weightdouble_slice/float32(w.Epsilon())
	Weight *= float32(w.Rate())

	if bias != 0 || weight != 0 {
		return Bias, Weight
	}

	return 0, 0
}
