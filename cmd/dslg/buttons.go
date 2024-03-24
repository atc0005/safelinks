// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func newCopyButton(w fyne.Window, outputField *widget.Entry) *widget.Button {
	copyButton := widget.NewButton("Copy to Clipboard", func() {
		log.Println("Copying decoded text to clipboard")
		w.Clipboard().SetContent(outputField.Text)
	})

	copyButton.Importance = widget.DangerImportance
	copyButton.Disable()

	return copyButton
}

func newDecodeButton(inputField *widget.Entry, copyButton *widget.Button, errOutField *widget.Entry, outputField *widget.Entry) *widget.Button {
	decodeButton := widget.NewButton("Decode", func() {
		if inputField.Text == "" {
			log.Println("Decoding requested but no input text provided")

			copyButton.Disable()
			errOutField.SetText("Please insert text to decode and try again.")
			errOutField.Refresh()

			return
		}

		log.Println("Decoding provided input text")

		result, err := processInput(inputField.Text)
		switch {
		case err != nil:
			errOutField.Append(err.Error() + "\n")
			errOutField.Refresh()

			return

		default:
			errOutField.SetText("OK: No errors encountered.")
			errOutField.Refresh()

			outputField.Text = result
			outputField.Refresh()

			copyButton.Enable()
		}
	})

	decodeButton.Importance = widget.HighImportance

	return decodeButton
}

func newResetButton(w fyne.Window, inputField *widget.Entry, copyButton *widget.Button, errOutField *widget.Entry, outputField *widget.Entry) *widget.Button {
	resetButton := widget.NewButton("Reset", func() {
		log.Println("Resetting application")
		w.Resize(fyne.NewSize(windowSizeHeight, windowSizeWidth))

		inputField.Text = ""
		inputField.Refresh()

		errOutField.Text = ""
		errOutField.Refresh()

		outputField.Text = ""
		outputField.Refresh()

		copyButton.Disable()
	})
	resetButton.Importance = widget.WarningImportance

	return resetButton
}

func newAboutButton(_ fyne.Window, inputField *widget.Entry, copyButton *widget.Button, errOutField *widget.Entry, outputField *widget.Entry) *widget.Button {
	aboutButton := widget.NewButton("About", func() {
		log.Println("Displaying About text")
		// inputField.Text = fmt.Sprintf(
		// 	"Description:\n\n%s\n\nLearn more about Safe Links:\n\n%s",
		// 	myAppDescription,
		// 	safeLinksAboutURL,
		// )
		inputField.Text = fmt.Sprintf("Description:\n\n%s", myAppDescription)
		inputField.Refresh()

		errOutField.Text = fmt.Sprintf(
			"Safe Links overview:\n\n%s",
			safeLinksAboutURL,
		)
		errOutField.Refresh()

		outputField.Text = "Current version:\n\n" + Version()
		outputField.Refresh()

		copyButton.Disable()
	})

	return aboutButton
}

func newExitButton(a fyne.App) *widget.Button {
	exitButton := widget.NewButton("Quit", func() {
		log.Println("Quit button called, exiting application")
		a.Quit()
	})
	exitButton.Importance = widget.WarningImportance

	return exitButton
}
