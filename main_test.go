package main

import (
	"fmt"
	"os"
	"testing"

	neuralnet "pixai/neural_network"
	natural "pixai/neural_network/natural_language_processing"
	gen "pixai/neural_network/generative"
)

func TestMain(t *testing.T) {
	n := neuralnet.Neurons{}
	g := gen.Generative{}
	nat := natural.NLP{}

	defer nat.Close()
	in, err := nat.NLPinit()
	if err != nil {
		fmt.Println("NLP returned an error")
	}

	defer g.Close()
	if err := g.GenerativeInit(in) ; err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	defer n.Close()
	if err := n.NeuralNetworkInit(in); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}
