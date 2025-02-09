// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// encoding.go
package generative

type Encoded struct {
	enStd []byte // uint8
}

var (
	e   = Encoded{}
	std = e.enStd
)

// Byte encoding, basically it's using the hex encoding from the NLP pieces of the algorithm.
func (e *Encoded) Encode(val uint8, input string) []byte {
	std = make([]byte, val)

	// just a for loop to give the inside value length
	for i := range std {
		if i != 0 {
			std = make([]byte, len(input))
		}

		if i == 0 {
			return nil
		}
	}

	if std != nil {
		return std
	}

	return nil
}

func (e *Encoded) Decode(val uint8) []byte {


	return nil
}

func (e *Encoded) Binary(input string) {

}
