// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	debugLoggingOut := os.Stderr // switch to os.Stderr for debugging
	// debugLoggingOut := io.Discard // use io.Discard for normal operation
	log.SetOutput(debugLoggingOut)

	a := app.New()
	w := a.NewWindow("Decode Microsoft Defender Safe Links")

	input := newInputTextField()
	errorOutput := newErrorOutputTextField()
	output := newOutputTextField()

	copyButton := newCopyButton(w, output)
	decodeButton := newDecodeButton(input, copyButton, errorOutput, output)
	resetButton := newResetButton(w, input, copyButton, errorOutput, output)
	aboutButton := newAboutButton(w, input, copyButton, errorOutput, output)

	exitButton := newExitButton(a)

	buttonRowContainer := NewButtonRowContainer(
		decodeButton,
		copyButton,
		resetButton,
		aboutButton,
		exitButton,
	)

	outputContainer := newOutputContainer(errorOutput, output)

	mainAppContainer := newMainAppContainer(input, buttonRowContainer, outputContainer)

	w.SetContent(mainAppContainer)
	w.Resize(fyne.NewSize(windowSizeHeight, windowSizeWidth))
	w.CenterOnScreen()

	// This prevents the UI from being accidentally collapsed to the point
	// that the buttons and text fields are no longer visible. Initial testing
	// showed that enabling this setting did not appear to prevent the user
	// from putting the application into fullscreen mode (on Linux distros at
	// least).
	w.SetFixedSize(true)

	w.ShowAndRun()
}
