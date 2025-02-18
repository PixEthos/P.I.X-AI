// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// encoding.go
package encode

import (
	"bufio"
	"encoding/hex"
	"log"
	"os"
)

type Encoded struct {
	enStd []byte // uint8
	input string
}

var (
	e   = Encoded{}
	std = e.enStd
	key = e.input
)

// key for encoding
func (e *Encoded) Key() string {

	// scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		// key
		key = scanner.Text()
	}

	// 0 check
	if len(key) != 0 {
		return key
	}

	return ""
}

// string encoding
func (e *Encoded) StringEncode(bit []byte) string {
	return hex.EncodeToString(bit)
}

// encoding
func (e *Encoded) Encode(input string, bit []byte) {
	std = make([]byte, len(input))
	hex.Encode(bit, std)
	e.Decode(input, bit)
}

// decoding
func (e *Encoded) Decode(input string, bit []byte) []byte {
	encoded := e.StringEncode(bit)
	decoded, err := hex.DecodeString(encoded) // decoding
	if err != nil {
		log.Fatal("Error decoding: ", err)
	}

	if decoded != nil {
		return decoded
	}

	return nil
}
