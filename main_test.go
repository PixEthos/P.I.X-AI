package main

import (
	"fmt"
	"os"
	"testing"

	neuralnet "pixai/neural_network"
)

func TestMain(t *testing.T) {
	n := neuralnet.Neurons{}

	defer n.Close()
	if err := n.NeuralNetworkInit(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}
