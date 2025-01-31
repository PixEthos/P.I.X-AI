// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// Created: October 28th, 2024

/* Explaination of the algorithm:
This is an algorithm I am basically building to not only see if I can do it but also seeing how I can improve
my understanding of them. Reverse engineering the complexity, reducing the complexity, for an elegent
and simple solution. Not an obscenely complicated algorithm with 50 different dependencies.

I take the thirdparty out, and see what I can do to find a solution.

Sometimes, this results in me taking a lot longer to understand it, but in the end? It's worth it.

This algorithm is basically just --as stated-- to see how I can implement one myself. I was wondering:
"Okay, I see this done a lot. How can I do this myself, without needing to learn dependencies?"

So, I decided: I'll make it in Go. Well, the first was in C, but it was damn garbage. This is because
I am better at experimenting; than I am at making something more functional. I admit that.
However, after deciding the language for my actual AI, Go was something I was considering
simple, concurrent, memorysafe, lightweight, with a powerful standard library.

You're probably wondering whether or not Go is great for machine learning; to put it lightly?
Yes, and no.

This isn't Python with popular and well-known libraries like numpy, tensorflow, scikitlearn, pandas
scapy, nltk, and more. So if you're doing any legitimate machine-learning work; and not just doing this
for fun - look at things like Python. If you're me? Go. I love the language. it's exactly what I needed*/

// main.go
package main

import (
	"fmt"
	"os"

	neuralnet "pixai/neural_network"
)

func main() {
	n := neuralnet.Neurons{}

	defer n.Close()
	if err := n.NeuralNetworkInit(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}
