// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// conv.go
package naturallanguageprocessing

import (
	"bufio"
	"encoding/csv"
	"os"
	"slices"
	"strings"
)

// All this is, is a conversion source file.

/* Explaination:
I needed to reduce complexity, and ensure more reliability
doing what I did here, to convert txt to csv, which while ridiculous
is actually a good idea for me; since it'll reduce complecity even further
especially down the line with more complexity for the chatbot.

This also houses a few string array conversions - mostly to float64 values
this is intentional for easier conversion later down the line.

All they do is convert the length of the slice/array to a float64 value.
Ridiculous? Yeah. But it was needed.*/

type Conversion struct{}

// .txt to .csv conversion.
func (c *Conversion) Convert(open, create string) []string {
	//var embed embed.FS

	// opening the current file
	val, _ := os.Open(open)
	defer val.Close()

	// creating a csv
	create_csv, _ := os.Create(create)
	defer create_csv.Close()

	// scanning the txt file internals
	scanner := bufio.NewScanner(val)

	field := []string{}

	// basically all this does is convert the text to a csv format
	for scanner.Scan() {

		// fields of the text files internals
		fields := strings.Fields(scanner.Text())

		// opening a csv writer
		csv := csv.NewWriter(create_csv)
		defer csv.Flush()

		// writing the fielded
		csv.Write(fields)

		field = append(field, fields...)
	}

	if field != nil {
		return field
	}

	return nil
}

// converting called files to a [][]string format
func (c *Conversion) ConvertToString(open string) [][]string {
	//var embed embed.FS

	// opening the file
	val, _ := os.Open(open)
	defer val.Close()

	v := [][]string{}

	// opening a scanner
	scanner := bufio.NewScanner(val)
	for scanner.Scan() {

		items := strings.Fields(scanner.Text())

		// converting again, into a string
		val := make([][]string, len(items))
		val = append(val, items)

		v = append(v, val...)
	}

	// nil check, I do these a lot
	if len(v) != 0 {
		return v
	}

	return nil
}

// document splitting
func (c *Conversion) DocumentSplitting(loc string) []string {
	//var embed embed.FS
	nlp := NLP{}

	open, _ := os.Open(loc)
	defer open.Close()

	fields := []string{}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		field := strings.Fields(scanner.Text())

		output := nlp.Document(scanner.Text()) // field splitting
		field = append(field, output...)

		fields = append(fields, field...)
	}

	if len(fields) != 0 {
		return fields
	}

	return nil
}

// extracting the string from an array
func (c *Conversion) ArraytoString(val []string) string {
	var v string

	for i := range val {
		v = val[i]
	}

	return v
}

// matching the internal strings within the array
func (c *Conversion) StringCheck(val []string, word string) bool {
	words := make(map[string][]string, len(val))
	for _, x := range val {
		words[x] = []string{x}

		similar, _ := words[word]

		if slices.Contains(similar, x) {
			return true
		}
	}

	return false
}

// Probability filter
func (c *Conversion) Filtration(filter int, element []string) float64 {
	var result float64
	output := []string{}

	// element filtering
	if len(element) != 0 && filter != 0 {
		for _, x := range element {

			// appending
			output = append(output, x)

			// matching
			for i := range output {
				if i != 0 {
					filter = len(output) / len(output)
				}

				// checking
				if filter != 0 {
					result = float64(filter)
					break
				}
			}
		}
	}

	if len(output) != 0 {
		return result
	}

	return 0
}

// string length for float64 calculation
func (c *Conversion) StringLength(val [][]string) float64 {
	var length float64

	x := make([]float64, len(val))
	for i := range x {
		x[i] = length
	}

	return length
}

// same as the first, just with a regular 1D array
func (c *Conversion) ArrayLength(val []string) float64 {
	var length float64

	x := make([]float64, len(val))
	for i := range x {
		x[i] = length
	}

	return length
}
