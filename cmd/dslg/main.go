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
	// use io.Discard for normal operation
	// switch to os.Stderr for debugging
	debugLoggingOut := os.Stderr

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(debugLoggingOut)

	a := app.New()
	w := a.NewWindow("Decode Microsoft Defender Safe Links")

	input := NewInputTextField()
	errorOutput := NewErrorOutputTextField()
	output := NewOutputTextLabel()

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

	outputContainer := NewOutputContainer(errorOutput, output)
	mainAppContainer := NewMainAppContainer(input, buttonRowContainer, outputContainer)

	w.SetContent(mainAppContainer)
	w.Resize(fyne.NewSize(windowSizeHeight, windowSizeWidth))
	w.CenterOnScreen()
	w.SetFixedSize(false)

	w.ShowAndRun()
}
