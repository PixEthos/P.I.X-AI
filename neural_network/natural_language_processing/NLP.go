// I use GPL2

/* Copyright (C) 2024 PixEthos */

// NLP.go
package naturallanguageprocessing

import (
	// standard
	"bufio"
	"fmt"
	"os"
	"strconv"

	// AI
	"pixai/data/cache"
)

// Why add natural_language_processing to a games AI?
// Why not?

type NLP struct {
	valStr string
}

type Input struct{}

func NaturalLanguagProcessing() *NLP {
	nlp := &NLP{}
	return nlp
}

/* Machine learning isn't difficult to put into practice
it's only difficult to read the 40 page thesis on what it is
with obfuscated and complicated wording to protect copyright.

Not even kidding.*/

// calling errors up the stack, as usual
func (nlp *NLP) NLPErrors(input, output string, length int32) error {
	var err error

	Valstr, err := strconv.Atoi(nlp.valStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Println(Valstr)
		return err
	}

	return err
}

// init
func (nlp *NLP) NLPinit() (string, error) {
	var err error
	c := Conversion{}
	m := Match{}

	cachestring := cache.RegCache[string, string]()

	sentence := bufio.NewScanner(os.Stdin)
	if sentence.Scan() {
		split := nlp.SplitTokens(sentence.Text())
		in := c.ArraytoString(split)

		sp := nlp.Tokens(len(split))
		m.Matching(in)
		match := m.MatchingLength(in)


		sentence2 := "PLACEHOLDER: I do not know, please understand"
		cachestring.SetReg("s2", sentence2)

		fmt.Printf("input tokens: %d\n", len(sp))
		fmt.Println("matched: ", match)

		if len(in) != 0 {
			return in, nil
		}
	}

	return "", err
}

func (nlp *NLP) close() {
	if nlp != nil {
		fmt.Println("NLP cleared from memory")
		return
	}
}

func (nlp *NLP) Close() {
	nlp.close()
}
