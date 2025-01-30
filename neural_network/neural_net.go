// I use GPL2

/* Copyright (C) 2024 PixEthos */

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

// learning
func (n *Neurons) NetworkLearning() (float64, error) {
	// structs
	w := Weights{}
	v := Variables{}
	c := Confidence{}
	nr := NotRecognized{}
	l := Layers{}

	// cache types
	f64cache := cache.Newf64[string, float64]()
	matrixCache := cache.RegCache[string, matrix.Matrix32]()

	defer nat.Close()
	in, err := nat.NLPinit()
	if err != nil {
		fmt.Println("NLP returned an error")
	}

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

	// neuron calcs
	inputs := mat32.Matrix32bit(input_array)
	neurons := n.encapsulated(n.Neuron_Count)
	secondary_neurons := n.encapsulated_secondary(400)
	trinary_neurons := n.encapsulated_trinary(400)
	primary, secondary, trinary := n.output(inputs, n.Neuron_Count, total)
	neuron_output := n.combined_outputs(inputs, n.Neuron_Count, total)
	neuron_accuracy, sec, tri := n.neuron_sigmoid(input_array, n.Neuron_Count, total)

	// GRU
	fmt.Println("\n---[DEBUGGING OUTPUTS:]--")

	gru_pri := n.gru_processed(inputs, in)
	pri_der := l.GRU_sigmoid(gru_pri, "float")

	gru_sec := n.gru_processed_secondary(inputs, in)
	sec_der := l.GRU_sigmoid(gru_sec, "float")

	gru_tri := n.gru_processed_trinary(inputs, in)
	tri_der := l.GRU_sigmoid(gru_tri, "float")

	// calculations
	w.derivative = v.SigmoidDerivative(neuron_accuracy) // derivative of the curve
	sec_derivative := v.SigmoidDerivative(sec)
	tri_derivative := v.SigmoidDerivative(tri)
	w.prediction = correct / float64(total)          // predictive
	accuracy = v.Sigmoid(w.Accuracy(total, correct)) // weighted prediction
	w.output = float64(total) + correct              // neuron output

	// confidence logs
	c.output_accuracy = w.output
	c.derivative_accuracy = w.derivative
	c.prediction_accuracy = w.prediction
	c.neuron_output = neuron_output

	// confidence caching
	f64cache.Setf64bit("confidence accuracy", accuracy)
	f64cache.Setf64bit("derivative accuracy", c.derivative_accuracy)
	f64cache.Setf64bit("prediction accuracy", c.prediction_accuracy)
	matrixCache.SetReg("neuron output", c.neuron_output)
	f64cache.Setf64bit("output total", c.output_accuracy)

	// output
	fmt.Println("\n------[NEURONS:]------")
	fmt.Println("combined number of neurons: ", len(neurons)+len(secondary_neurons)+len(trinary_neurons))
	fmt.Println("primary set: ", len(neurons))
	fmt.Println("secondary set: ", len(secondary_neurons))
	fmt.Println("trinary set: ", len(trinary_neurons))

	fmt.Println("\n------[ACCURACY:]-----")
	fmt.Println("primary_accuracy: ", c.derivative_accuracy, pri_der)
	fmt.Println("secondary_accuracy: ", sec_derivative, sec_der)
	fmt.Println("trinary_accuracy: ", tri_derivative, tri_der)
	fmt.Println("prediction: ", c.prediction_accuracy)
	fmt.Println("accuracy: ", accuracy)

	fmt.Println("\n------[OUTPUTS:]------")
	fmt.Println("output total: ", c.output_accuracy)
	fmt.Println("linear: ", c.linear_accuracy)
	fmt.Println("primary set: ", primary)
	fmt.Println("secondary set: ", secondary)
	fmt.Println("trinary set: ", trinary)
	fmt.Println("combined: ", c.neuron_output)

	// NaN checking
	if math.IsNaN(w.derivative) || math.IsNaN(w.prediction) || math.IsNaN(w.output) || math.IsNaN(w.Accuracy(total, correct)) {
		return 0, nr.Values()
	}

	// output
	out = w.output + float64(output) + correct + float64(total) // total outputs

	return out, nil
}

// network
func (n *Neurons) NeuralNetwork() error {

	output, err := n.NetworkLearning()
	if err != nil {
		return err
	}

	fmt.Println("output:", output)

	return err
}

// init
func (n *Neurons) NeuralNetworkInit() error {
	var err error

	if err := n.NeuralNetwork(); err != nil {
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
