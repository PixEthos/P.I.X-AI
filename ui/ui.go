// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// ui.go
package ui

import (
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	neural "pixai/neural_network"
	gen "pixai/neural_network/generative"
	nlp "pixai/neural_network/natural_language_processing"
)

// constants
const (
	width  = 500
	height = 500
)

// variable constants
var (
	input = widget.NewEntry()
)

type UserInterface struct {
	Output string
}

func (UserInterface) OutputCaller(input string) UserInterface {
	user := UserInterface{
		Output: input,
	}

	return user
}

func ApplicationInit(input string) string {
	natural := nlp.NLP{}
	n := neural.Neurons{}
	output := gen.Generative{}
	ui := UserInterface{}

	defer natural.Close()
	in, err := natural.NLPinit(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	defer output.Close()
	if information, err := output.GenerativeInit(in); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		ui.OutputCaller(information)

		defer n.Close()
		if err := n.NeuralNetworkInit(in); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		if len(information) > 0 {
			return information
		}
	}

	return ""
}

// input
func (ui *UserInterface) ApplicationInput() string {
	input.SetPlaceHolder("Input the data/input here")
	if len(input.Text) != 0 {
		return input.Text
	}

	return ""
}

// Window
func (ui *UserInterface) ApplicationWindow() {
	log.Println("Application started")

	// Fyne callers
	runner := app.New()
	window := runner.NewWindow("data") // title

	// resize
	window.Resize(fyne.NewSize(width, height))

	// functions
	ui.ApplicationInput()

	// labels
	draw_input := widget.NewLabel("Input: ")
	draw_output := widget.NewLabel(ui.Output)

	// container
	window.SetContent(container.NewVBox(
		input,
		draw_input,
		draw_output,
		widget.NewButton("send", func() {
			draw_input.SetText(input.Text)
			output := ApplicationInit(input.Text)
			draw_output.SetText(output)

		}),
	))

	log.Println(ui.Output)

	// runner
	window.ShowAndRun()
}

// UI caller
func (ui *UserInterface) ApplicationUI() {
	ui.ApplicationWindow()
}

// closer
func (ui *UserInterface) close() {
	if ui != nil {
		log.Println("Closing application")
	}
}

// close caller
func (ui *UserInterface) Close() {
	ui.close()
}
