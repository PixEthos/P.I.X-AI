// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// ui.go
package ui

import (

	// std
	"fmt"
	"log"
	"os"
	"strconv"

	// fyne
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// internal
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

// OutputCaller
func (UserInterface) OutputCaller(input string) UserInterface {
	user := UserInterface{
		Output: input,
	}

	return user
}

// Machine learning init
func ApplicationInit(input string) string {

	// structs
	natural := nlp.NLP{}
	n := neural.Neurons{}
	output := gen.Generative{}
	ui := UserInterface{}

	// NLP
	defer natural.Close()
	in, err := natural.NLPinit(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	// generative
	defer output.Close()
	if information, err := output.GenerativeInit(in); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		ui.OutputCaller(information)

		// neural_network/GRU
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

// input length
func InputLenght(input string) int {
	var i int
	for i = range len(input) {
		i = len(input)
	}

	if i != 0 {
		return i
	}

	return 0
}

// input
func (ui *UserInterface) ApplicationInput() string {
	input.SetPlaceHolder("Input the data/input here")
	if len(input.Text) != 0 || len(input.Text) != 127 {
		return input.Text
	}

	if len(input.Text) > 128 {
		log.Println("127 is the limit of characters")
	}

	return ""
}

// Window
func (ui *UserInterface) ApplicationWindow() {
	log.Println("Application started")

	os.Setenv("FYNE_SCALE", "1.5")

	// Fyne callers
	runner := app.New()
	window := runner.NewWindow("PixAI: Prealpha (0.39-1)") // title

	// resize
	window.Resize(fyne.NewSize(width, height))
	settings := fyne.CurrentApp().Settings()
	settings.Scale()

	// functions
	ui.ApplicationInput()

	limit := widget.NewLabel("The Input is limited to 128 (characters)")
	description := widget.NewLabel("This is a passion project, still in early alpha. Just a basic match/predict as of now")
	github := widget.NewLabel("Check here, for updates: https://github.com/PixEthos/PixAI")

	// input labels
	input_length := widget.NewLabel("")
	draw_input := widget.NewLabel("")

	// output labels
	true_output := widget.NewLabel("Output: ")
	draw_output := widget.NewLabel("")

	// output style
	draw_output.TextStyle.Bold = true
	draw_output.TextStyle.Symbol = true

	// scrolling area
	horizontal := container.NewScroll(draw_output)
	holding := container.NewStack(horizontal)
	grid := container.NewGridWithRows(2, holding)

	// container
	content := container.NewVBox(
		description,
		github,
		limit,
		input_length,
		input,
		widget.NewButton("send", func() {
			// length
			x := InputLenght(input.Text)
			input_length.SetText(strconv.Itoa(x))

			// input
			draw_input.SetText(input.Text)

			// output
			output := ApplicationInit(input.Text)
			draw_output.SetText(output)
		}),
		true_output,
		grid,
	)

	window.SetContent(content)

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
