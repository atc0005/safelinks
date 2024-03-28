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

// NewInputTextField creates a text input field intended for bulk copy/paste
// of encoded content. While we do not intentionally restrict interacting with
// input text the field is sized primarily for reviewing only a portion of it
// before decoding it.
func NewInputTextField() *widget.Entry {
	inputField := widget.NewMultiLineEntry()
	inputField.Wrapping = fyne.TextWrapOff
	inputField.PlaceHolder = inputFieldPlaceholder
	inputField.SetMinRowsVisible(10)

	return inputField
}

// NewErrorOutputTextField creates a text output field to serve as a "bucket"
// for any decoding errors that may occur. We assume that only a limited view
// of this content is needed and so restrict the vertical space to provide
// more space for decoded output.
func NewErrorOutputTextField() *widget.Entry {
	errOutField := widget.NewMultiLineEntry()
	// errOutField.TextStyle = fyne.TextStyle{Monospace: true, Italic: true}
	errOutField.TextStyle = fyne.TextStyle{Monospace: true}
	errOutField.PlaceHolder = errOutPlaceholder
	errOutField.SetMinRowsVisible(5)

	return errOutField
}

// NewOutputTextField creates a text output field to serve as a "bucket" for
// decoded input text. A greater portion of the visible window is dedicated to
// this field for output review purposes.
//
// NOTE: Performance issues have been observed with using this Fyne toolkit
// object type for output text. Due to this, we use a widget.Label instead and
// rely on the copy to clipboard button to retrieve decoded text.
//
// https://stackoverflow.com/questions/75530554
// https://github.com/fyne-io/fyne/issues/4014
// https://github.com/fyne-io/fyne/issues/2969
//
// Deprecated: use NewOutputTextLabel instead.
func NewOutputTextField() *widget.Entry {
	outputField := widget.NewMultiLineEntry()
	outputField.Wrapping = fyne.TextWrapOff
	outputField.PlaceHolder = decodedOutputPlaceholder
	outputField.SetMinRowsVisible(15)

	// This works, but the style mutes the text color significantly when using
	// the dark theme.
	//
	// outputField.Disable()

	return outputField
}

// NewOutputTextLabel creates a text output label to serve as a "bucket" for
// decoded input text. A greater portion of the visible window is dedicated to
// this field for output review purposes.
//
// Due to performance issues observed with using a widget.Entry type we use a
// widget.Label instead and rely solely on the copy to clipboard button to
// retrieve decoded text.
func NewOutputTextLabel() *widget.Label {
	outputLabel := widget.NewLabel(decodedOutputPlaceholder)
	outputLabel.Wrapping = fyne.TextWrapWord

	return outputLabel
}
