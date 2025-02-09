// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// predictive.go
package naturallanguageprocessing

/* Simplicity should be an absolute priority when making anything.

I find myself reading and reading, and the more I do; the more frustrated
I grow with the overcomplicating terms, and long ideals.

If one cannot simplify, than one does not understand.*/

type PredictiveOutput struct {
	gpe  float64
	stop float64
	noun float64
	verb float64
}

/* Explaination:
What you see within each of these functions, is basically the same functions
per value, or word, or match. This is intentional, as it reduces the need for
consistently writing the same thing over and over elsewhere.

Adding to this; using goroutines for this purpose allows me to use Go's concurrency
for effective matching, and predictive values.*/

func (p *Predictive) GPE_predictive(input string) ([][]string, []string, []string) {
	// calling variables, and structs
	nlp := NLP{}
	c := Conversion{}
	doc := nlp.SplitTokens(input)       // Document
	words := p.Words().GPE              // words
	open := "../PixAI/words_en/GPE.csv" // file location
	word := c.ConvertToString(open)     // converting to string

	if doc != nil || words != nil || len(word) != 0 {
		return word, words, doc
	}

	return nil, nil, nil
}

// GPE prediction
func (p *Predictive) PredictGPE(input string, val chan float64) {
	c := Conversion{}

	_, words, doc := p.GPE_predictive(input)
	//fmt.Println("known_matches:", len(word), "number_known:", len(words), "input_len:", len(doc))

	// comparing the lengths and values with the other arrays
	v := 0
	x := len(doc)
	for _, word_s := range doc {
		found := c.StringCheck(words, word_s)
		for found {
			if true {
				v = v + 1
				if v > x {
					v = v - x
				}
			}
			break
		}

		//fmt.Println("GPE:", found)
	}

	// appending the document fielding
	doc = append(doc, words...)

	filter := c.Filtration(v, doc)

	output := p.Probability(filter, float64(v))

	// sending the data through a channel
	val <- output

	close(val)
}

// Caller functions
func (p *Predictive) GPEVal(input string, val chan float64) float64 {
	// calling the function with a goroutine
	go p.PredictGPE(input, val)

	// checking the internals to return a value
	for i := range val {
		if i != 0 {
			return i
		}
	}

	return 0
}

// Activator functions
func (p *Predictive) GPEActivator(input string) float64 {
	// creating a channel for each variable
	GPE := make(chan float64)
	gpe := p.GPEVal(input, GPE) // function

	if gpe != 0 {
		return gpe
	}

	return 0
}

func (p *Predictive) STOPWORDS_predictive(input string) ([][]string, []string, []string) {
	// calling variables, and structs
	nlp := NLP{}
	c := Conversion{}
	doc := nlp.SplitTokens(input) // Document
	words := p.Words().Stopwords
	open := "../PixAI/words_en/stopwords.csv"
	word := c.ConvertToString(open) // converting to string

	if doc != nil || words != nil || len(word) != 0 {
		return word, words, doc
	}

	return nil, nil, nil
}

// stopword predictions
func (p *Predictive) PredictStopwords(input string, val chan float64) {
	c := Conversion{}

	_, words, doc := p.STOPWORDS_predictive(input)
	//fmt.Println("known_matches:", len(word), "number_known:", len(words), "input_len:", len(doc))

	v := 0
	x := len(doc)
	for _, word_s := range doc {
		found := c.StringCheck(words, word_s)
		for found {
			if true {
				v = v + 1
				if v > x {
					v = v - x
				}
			}
			break
		}

		//fmt.Println("StopWords:", found)
	}

	doc = append(doc, words...)

	filter := c.Filtration(v, doc)

	output := p.Probability(filter, float64(v))

	val <- output

	close(val)
}

func (p *Predictive) STOPWORDSVal(input string, val chan float64) float64 {

	go p.PredictStopwords(input, val)

	for i := range val {
		if i != 0 {
			return i
		}
	}

	return 0
}

func (p *Predictive) STOPWORDActivator(input string) float64 {

	// creating a channel for each variable
	STOP := make(chan float64)
	stop := p.STOPWORDSVal(input, STOP)
	if stop != 0 {
		return stop
	}

	return 0
}

func (p *Predictive) NOUNS_predictive(input string) ([][]string, []string, []string) {
	// calling variables, and structs
	nlp := NLP{}
	c := Conversion{}
	doc := nlp.SplitTokens(input) // Document
	words := p.Words().Nouns
	open := "../PixAI/words_en/nouns.csv"
	word := c.ConvertToString(open) // converting to string

	if doc != nil || words != nil || len(word) != 0 {
		return word, words, doc
	}

	return nil, nil, nil
}

// noun predictions
func (p *Predictive) PredictNouns(input string, val chan float64) {
	c := Conversion{}

	_, words, doc := p.NOUNS_predictive(input)
	//fmt.Println("known_matches:", len(word), "number_known:", len(words), "input_len:", len(doc))

	v := 0
	x := len(doc)
	for _, word_s := range doc {
		found := c.StringCheck(words, word_s)
		for found {
			if true {
				v = v + 1
				if v > x {
					v = v - x
				}
			}
			break
		}

		//fmt.Println("Nouns:", found)
	}

	doc = append(doc, words...)

	filter := c.Filtration(v, doc)

	output := p.Probability(filter, float64(v))

	val <- output

	close(val)
}

func (p *Predictive) NOUNSVal(input string, val chan float64) float64 {

	go p.PredictNouns(input, val)

	for i := range val {
		if i != 0 {
			return i
		}
	}

	return 0
}

func (p *Predictive) NOUNActivator(input string) float64 {
	// creating a channel for each variable
	NOUN := make(chan float64)
	noun := p.NOUNSVal(input, NOUN)
	if noun != 0 {
		return noun
	}

	return 0
}

func (p *Predictive) VERBS_predictive(input string) ([][]string, []string, []string) {
	// calling variables, and structs
	nlp := NLP{}
	c := Conversion{}
	doc := nlp.SplitTokens(input) // Document
	words := p.Words().Verbs
	open := "../PixAI/words_en/verbs.csv"
	word := c.ConvertToString(open) // converting to string

	if doc != nil || words != nil || len(word) != 0 {
		return word, words, doc
	}

	return nil, nil, nil
}

// verb predictions
func (p *Predictive) PredictVerbs(input string, val chan float64) {
	c := Conversion{}

	_, words, doc := p.VERBS_predictive(input)
	//fmt.Println("known_matches:", len(word), "number_known:", len(words), "input_len:", len(doc))

	v := 0
	x := len(doc)
	for _, word_s := range doc {
		found := c.StringCheck(words, word_s)
		for found {
			if true {
				v = v + 1
				if v > x {
					v = v - x
				}
			}
			break
		}

		//fmt.Println("Verbs:", found)
	}

	doc = append(doc, words...)

	filter := c.Filtration(v, doc)

	output := p.Probability(filter, float64(v))

	val <- output

	close(val)
}

func (p *Predictive) VERBSVal(input string, val chan float64) float64 {

	go p.PredictVerbs(input, val)

	for i := range val {
		if i != 0 {
			return i
		}
	}

	return 0
}

func (p *Predictive) VERBActivator(input string) float64 {

	VERB := make(chan float64)
	verb := p.VERBSVal(input, VERB)
	if verb != 0 {
		return verb
	}

	return 0
}

// Prediction caller
func (p *Predictive) Predict(input string) *PredictiveOutput {

	gpe := p.GPEActivator(input) // function
	stop := p.STOPWORDActivator(input)
	noun := p.NOUNActivator(input)
	verb := p.VERBActivator(input)

	// struct caller for outputs
	output := PredictiveOutput{
		noun: noun,
		verb: verb,
		stop: stop,
		gpe:  gpe,
	}

	return &output
}
