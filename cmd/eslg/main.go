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

	// Help this tool stand out from the dslg app.
	if err := os.Setenv("FYNE_THEME", "light"); err != nil {
		log.Println("Failed to set fyne toolkit theme")
	}

	// NOTE: This is deprecated and set to be removed in v3.0.
	// fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())

	a := app.New()
	w := a.NewWindow("Create faux encoded Microsoft Defender Safe Links")

	input := NewInputTextField()
	errorOutput := NewErrorOutputTextField()
	output := NewOutputTextLabel()

	copyButton := newCopyButton(w, output)
	encodeButton := newEncodeButton(input, copyButton, errorOutput, output)
	resetButton := newResetButton(w, input, copyButton, errorOutput, output)
	aboutButton := newAboutButton(w, input, copyButton, errorOutput, output)

	exitButton := newExitButton(a)

	buttonRowContainer := NewButtonRowContainer(
		encodeButton,
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
