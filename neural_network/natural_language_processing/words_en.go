// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// words_en.go
package naturallanguageprocessing

/* This is just conversion area. Reducing the complexity using the conv.go
source file. */

type Words struct {
	Stopwords []string
	GPE       []string
	Nouns     []string
	Verbs     []string
}

// stopwords, essentially these are common use-words
func (p *Predictive) Stopwords() []string {
	c := Conversion{}

	open := "../PixAI/words_en/stopwords.txt"
	create := "../PixAI/words_en/stopwords.csv"

	word := c.Convert(open, create)
	if word != nil {
		return word
	}

	return nil
}

// Geopolitical entities, it encompuses more than what you see here, but this is sample data
func (p *Predictive) GPE() []string {
	c := Conversion{}

	open := "../PixAI/words_en/GPE.txt"
	create := "../PixAI/words_en/GPE.csv"

	word := c.Convert(open, create)
	if word != nil {
		return word
	}

	return nil
}

func (p *Predictive) Nouns() []string {
	c := Conversion{}

	open := "../PixAI/words_en/nouns.txt"
	create := "../PixAI/words_en/nouns.csv"

	word := c.Convert(open, create)
	if word != nil {
		return word
	}

	return nil
}

func (p *Predictive) Verbs() []string {
	c := Conversion{}

	open := "../PixAI/words_en/verbs.txt"
	create := "../PixAI/words_en/verbs.csv"

	word := c.Convert(open, create)
	if word != nil {
		return word
	}

	return nil
}

// Words struct caller
func (w *Words) Words() *Words {
	p := Predictive{}

	stop := p.Stopwords()
	gpe := p.GPE()
	nouns := p.Nouns()
	verbs := p.Verbs()

	words := Words{
		Stopwords: stop,
		GPE:       gpe,
		Nouns:     nouns,
		Verbs:     verbs,
	}

	return &words
}
