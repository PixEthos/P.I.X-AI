// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// context.go
package generative

import (
	"log"

	matrix "pixai/data/matrix"
)

// gru layering
func (g *Generative) GRU_pri(input matrix.Matrix32, x string) matrix.Matrix32 {

	gru_sec := layer.GRU_activation(200, 30, input, "float64", x)
	if gru_sec != nil {
		return gru_sec
	}

	return nil
}

// context holder
func (g *Generative) Gru_processed(input matrix.Matrix32, con string) matrix.Matrix32 {
	output := g.GRU_pri(input, con)

	if output != nil {
		return output
	}

	return nil
}

// primary context
func (g *Generative) PrimaryContext(input string) (float64, float64) {
	gpe := context.GPEActivator(input)
	stop := context.STOPWORDActivator(input)

	return gpe, stop
}

// primary activator
func (g *Generative) GRU_primary(input matrix.Matrix32, value string) (float64, string) {
	mat32 := matrix.Matrix{}

	// rune
	GRU := layer.GRU_rune_variable(value, input)
	output := mat32.RuneToMatrix32(GRU)
	val := g.GRU_decode(GRU, value)

	// primary
	gpe, stop := g.PrimaryContext(value)
	gru_pri := g.Gru_processed(output, value)

	// accuracy
	primary := layer.GRU_sigmoid(gru_pri, "float64", value)

	// logs
	log.Println("GRU_1 accuracy: ", primary)
	log.Println("GPE match: ", gpe, "Stopword match: ", stop)
	log.Println("GRU_1: ", val)
	log.Println("ASCII: ", GRU)

	return primary, val
}

// primary caller
func (g *Generative) Primary(input matrix.Matrix32, value string) {
	_, char := g.GRU_primary(input, value)
	x := Prefix{char}
	x.Join(char)
	x.Merge(value)
}

// secondary
func (g *Generative) GRU_second(input matrix.Matrix32, x string) matrix.Matrix32 {

	gru_sec := layer.GRU_activation(100, 10, input, "float64", x)
	if gru_sec != nil {
		return gru_sec
	}

	return nil
}

func (g *Generative) Gru_processed_secondary(input matrix.Matrix32, con string) matrix.Matrix32 {
	output := g.GRU_second(input, con)

	if output != nil {
		return output
	}

	return nil
}

func (g *Generative) SecondaryContext(input string) float64 {
	nouns := context.NOUNActivator(input)

	return nouns
}

func (g *Generative) GRU_secondary(input matrix.Matrix32, value string) (float64, string) {
	// rune
	GRU_2 := layer.GRU_rune_variable(value, input)
	output := mat32.RuneToMatrix32(GRU_2)
	val1 := g.GRU_decode(GRU_2, value)

	// secondary
	nouns := g.SecondaryContext(value)
	gru_sec := g.Gru_processed_secondary(output, value)

	// accuracy
	secondary := layer.GRU_sigmoid(gru_sec, "float64", value)

	// logs
	log.Println("GRU_2 accuracy: ", secondary)
	log.Println("Nouns match: ", nouns)
	log.Println("GRU_2: ", val1)
	log.Println("ASCII: ", GRU_2)

	return secondary, val1
}

func (g *Generative) Secondary(input matrix.Matrix32, value string) {
	_, char := g.GRU_secondary(input, value)
	x := Prefix{char}
	x.Join(char)
	x.Merge(value)
}

// trinary
func (g *Generative) GRU_tri(input matrix.Matrix32, x string) matrix.Matrix32 {

	gru_sec := layer.GRU_activation(100, 10, input, "float64", x)
	if gru_sec != nil {
		return gru_sec
	}

	return nil
}

// context holder
func (g *Generative) Gru_processed_trinary(input matrix.Matrix32, con string) matrix.Matrix32 {
	output := g.GRU_tri(input, con)

	if output != nil {
		return output
	}

	return nil
}

func (g *Generative) TrinaryContext(input string) float64 {
	verbs := context.VERBActivator(input)

	return verbs
}

func (g *Generative) GRU_trinary(input matrix.Matrix32, value string) (float64, string) {
	// rune
	GRU_3 := layer.GRU_rune_variable(value, input)
	val2 := g.GRU_decode(GRU_3, value)
	output := mat32.RuneToMatrix32(GRU_3)

	// trinary
	verbs := g.TrinaryContext(value)
	gru_tri := g.Gru_processed_trinary(output, value)

	// accuracy
	trinary := layer.GRU_sigmoid(gru_tri, "float64", value)

	// logs
	log.Println("GRU_3 accuracy: ", trinary)
	log.Println("Verb match: ", verbs)
	log.Println("GRU_3: ", val2)
	log.Println("ASCII: ", GRU_3)

	return trinary, val2
}

func (g *Generative) Trinary(input matrix.Matrix32, value string) {
	_, char := g.GRU_trinary(input, value)
	x := Prefix{char}
	x.Join(char)
	x.Merge(value)
}
