// I use GPL2

/* Copyright (C) 2024 PixEthos */

// matching.go
package naturallanguageprocessing

import (
	"regexp"
	"sync"
)

// matching struct, helps with punctuation, whitespace, values
type Match struct {
	punct      string
	alpha      string
	alphabetic string
	digit      string
	white      string
	word       string
}

// Array struct, for adding in each matched value
type MatchArray struct {
	punct_a      []string
	alpha_a      []string
	alphabetic_a []string
	digit_a      []string
	white_a      []string
	word_a       []string
}

// Input/matching struct
type IvyMike struct {
	input    string
	matching string
}

func (i *IvyMike) Input(input, matching string) *IvyMike {
	ivy := IvyMike{
		input:    input,
		matching: matching,
	}

	return &ivy
}

// These are the regexp patterns for whitespace, punctuation, alphanumeric, alphabetic, and digit
func (m *Match) Values() *Match {

	match := Match{
		punct:      `\p{P}\s`,
		white:      `\s`,
		alpha:      `\S`,
		alphabetic: `\a`,
		digit:      `\d`,
		word:       `\w`,
	}

	return &match
}

// Punctuation matching
func (m *Match) Puntcuation(input string, val chan []string) {

	p := m.Values().punct

	reg := regexp.MustCompile(p)
	expression := reg.Split(input, -1)
	f := reg.Match([]byte(input))

	if f {
		for _, x := range expression {
			expression = append(expression, x)
			val <- expression
		}
	}

	close(val)
}

// Punct len
func (m *Match) Punct(input string, val chan []string) []string {
	var wg sync.WaitGroup

	go m.Puntcuation(input, val)
	wg.Add(5)
	defer wg.Done()

	for i := range val {
		if len(i) == 0 {
			return nil
		}

		if len(i) != 0 {
			return i
		}
	}

	return nil
}

// Whitespace matching
func (m *Match) Whitespace(input string, val chan []string) {

	w := m.Values().white

	reg := regexp.MustCompile(w)
	space := reg.Split(input, -1)
	f := reg.Match([]byte(input))

	if f {
		for _, x := range space {
			space = append(space, x)
			val <- space
		}
	}

	close(val)
}

// Whitespace len
func (m *Match) White(input string, val chan []string) []string {
	var wg sync.WaitGroup

	go m.Whitespace(input, val)
	wg.Add(5)
	defer wg.Done()

	for i := range val {
		if len(i) == 0 {
			return nil
		}

		if len(i) != 0 {
			return i
		}
	}

	return nil
}

// Character match
func (m *Match) Characters(input string, val chan []string) {

	char := m.Values().alpha

	reg := regexp.MustCompile(char)
	alpha := reg.Split(input, -1)
	f := reg.Match([]byte(input))

	if f {
		for _, x := range alpha {
			alpha = append(alpha, x)
			val <- alpha
		}
	}

	close(val)
}

// Characters len
func (m *Match) Char(input string, val chan []string) []string {
	var wg sync.WaitGroup

	go m.Characters(input, val)
	wg.Add(5)
	defer wg.Done()

	for i := range val {
		if len(i) == 0 {
			return nil
		}

		if len(i) != 0 {
			return i
		}
	}

	return nil
}

// Alphabetic match
func (m *Match) Alphabetic(input string, val chan []string) {

	a := m.Values().alphabetic

	reg := regexp.MustCompile(a)
	alphabet := reg.Split(input, -1)
	f := reg.Match([]byte(input))

	if f {
		for _, x := range alphabet {
			alphabet = append(alphabet, x)
			val <- alphabet
		}
	}

	close(val)
}

// Alphabet len
func (m *Match) Alpha(input string, val chan []string) []string {
	var wg sync.WaitGroup

	go m.Alphabetic(input, val)
	wg.Add(5)
	defer wg.Done()

	for i := range val {
		if len(i) == 0 {
			return nil
		}

		if len(i) != 0 {
			return i
		}
	}

	return nil
}

// Digits match
func (m *Match) Digits(input string, val chan []string) {

	d := m.Values().digit

	reg := regexp.MustCompile(d)
	digit := reg.Split(input, -1)
	f := reg.Match([]byte(input))

	if f {
		for _, x := range digit {
			digit = append(digit, x)
			val <- digit
		}
	}

	close(val)
}

// Digits len
func (m *Match) Digit(input string, val chan []string) []string {
	var wg sync.WaitGroup

	go m.Digits(input, val)
	wg.Add(5)
	defer wg.Done()

	for i := range val {
		if len(i) == 0 {
			return nil
		}

		if len(i) != 0 {
			return i
		}
	}

	return nil
}

// Word matching
func (m *Match) Word(input string, val chan []string) {

	w := m.Values().word

	reg := regexp.MustCompile(w)
	word := reg.Split(input, -1)
	f := reg.Match([]byte(input))

	if f {
		for _, x := range word {
			word = append(word, x)
			val <- word
		}
	}

	close(val)
}

// Word len
func (m *Match) Words(input string, val chan []string) []string {
	var wg sync.WaitGroup

	go m.Word(input, val)
	wg.Add(5)
	defer wg.Done()

	for i := range val {
		if len(i) == 0 {
			return nil
		}

		if len(i) != 0 {
			return i
		}
	}

	return nil
}

// Matching values. Essentially this is just here for all the available numeric/alphanumeric setup
func (m *Match) Matching(input string) *MatchArray {

	word := make(chan []string)
	words := m.Words(input, word)

	punc := make(chan []string)
	punct := m.Punct(input, punc)

	digit := make(chan []string)
	digits := m.Digit(input, digit)

	white := make(chan []string)
	whitespace := m.White(input, white)

	char := make(chan []string)
	characters := m.Char(input, char)

	alpha := make(chan []string)
	alphabetic := m.Alpha(input, alpha)

	ma := MatchArray{
		alpha_a:      characters,
		alphabetic_a: alphabetic,
		punct_a:      punct,
		digit_a:      digits,
		word_a:       words,
		white_a:      whitespace,
	}

	return &ma
}

// Length of matched values, it is just used to give you an idea how many values match
func (m *Match) MatchingLength(input string) int {
	length := m.Matching(input)
	alpha := len(length.alpha_a)
	alphabetic := len(length.alphabetic_a)
	whitespace := len(length.white_a)
	punctuation := len(length.punct_a)
	digits := len(length.digit_a)
	words := len(length.word_a)

	numeric := digits + whitespace
	alphanumeric := words + punctuation + alpha + alphabetic

	added_length := numeric + alphanumeric

	return added_length
}
