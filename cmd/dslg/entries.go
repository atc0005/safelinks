// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func newInputTextField() *widget.Entry {
	inputField := widget.NewMultiLineEntry()
	inputField.Wrapping = fyne.TextWrapOff
	inputField.PlaceHolder = "Paste text with encoded URLs here and press Decode."
	inputField.SetMinRowsVisible(10)

	return inputField
}

func newErrorOutputTextField() *widget.Entry {
	errOutField := widget.NewMultiLineEntry()
	// errOutField.TextStyle = fyne.TextStyle{Monospace: true, Italic: true}
	errOutField.TextStyle = fyne.TextStyle{Monospace: true}
	errOutField.PlaceHolder = "Decoding errors (if any) will be logged here. Text pasted here is ignored."
	errOutField.SetMinRowsVisible(3)

	return errOutField
}

func newOutputTextField() *widget.Entry {
	outputField := widget.NewMultiLineEntry()
	outputField.Wrapping = fyne.TextWrapOff
	outputField.PlaceHolder = "Decoded text will be placed here.\n\nChanges are overwritten upon button press."
	outputField.SetMinRowsVisible(15)

	// This works, but the style mutes the text color significantly when using
	// the dark theme.
	//
	// outputField.Disable()

	return outputField
}
