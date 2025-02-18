// I use GPL2

/* Copyright (C) 2024, 2025 PixEthos */

/* This file is part of PixAI.

PixAI is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 2 of the License, or (at your option) any later version.

PixAI is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with PixAI. If not, see <https://www.gnu.org/licenses/>. */

// ui.go
package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	output string
}

// caller
func (UserInterface) OutputCaller(input string) UserInterface {
	user := UserInterface{
		output: input,
	}

	return user
}

// input
func (UserInterface) ApplicationInput() string {
	input.SetPlaceHolder("Input the data/input here")

	if len(input.Text) != 0 {
		return input.Text
	}

	return ""
}

// Window
func (ui *UserInterface) ApplicationWindow() {

	// Fyne callers
	runner := app.New()
	window := runner.NewWindow("data") // title

	// resize
	window.Resize(fyne.NewSize(width, height))

	// functions
	ui.ApplicationInput()

	// labels
	draw_input := widget.NewLabel("Input: ")
	draw_output := widget.NewLabel(ui.output)

	// container
	window.SetContent(container.NewVBox(
		input,
		draw_input,
		widget.NewButton("send", func() {
			draw_input.SetText(input.Text)
		}),
		draw_output,
	))

	log.Println(ui.output)

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
		log.Println("UI clear from memory")
	}
}

// close caller
func (ui *UserInterface) Close() {
	ui.close()
}
