// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"log"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/atc0005/safelinks/internal/safelinks"
)

func newCopyButton(a fyne.App, outputField *widget.Label) *widget.Button {
	copyButton := widget.NewButton("Copy to Clipboard", func() {
		log.Println("Copying decoded text to clipboard")
		a.Clipboard().SetContent(outputField.Text)
	})

	copyButton.Importance = widget.DangerImportance
	copyButton.Disable()

	return copyButton
}

func newEncodeButton(randomEncode bool, inputField *widget.Entry, copyButton *widget.Button, errOutField *widget.Entry, outputField *widget.Label) *widget.Button {
	buttonLabelText := func() string {
		if randomEncode {
			return "Encode Randomly"
		}
		return "Encode All"
	}()

	encodeButton := newProcessInputButton(
		randomEncode,
		buttonLabelText,
		safelinks.EncodeInput,
		inputField,
		copyButton,
		errOutField,
		outputField,
	)

	return encodeButton
}

func newProcessInputButton(
	// TODO: Refactor this to reduce parameters.
	randomEscape bool,
	buttonLabelText string,
	processFunc func(string, bool) (string, error),
	inputField *widget.Entry,
	copyButton *widget.Button,
	errOutField *widget.Entry,
	outputField *widget.Label,
) *widget.Button {

	button := widget.NewButton(buttonLabelText, func() {
		if inputField.Text == "" {
			log.Printf("%s used but no input text provided", buttonLabelText)

			copyButton.Disable()
			errOutField.Text = errOutTryAgain + "\n"
			errOutField.Refresh()

			return
		}

		log.Printf("%s input text", buttonLabelText)

		result, err := processFunc(inputField.Text, randomEscape)
		switch {
		case err != nil:
			errOutField.Append(err.Error() + "\n")
			errOutField.Refresh()

			return

		default:
			errOutField.PlaceHolder = "OK: No errors encountered."
			errOutField.Text = ""
			errOutField.Refresh()

			outputField.Text = result
			outputField.Refresh()

			copyButton.Enable()
		}
	})

	if randomEscape {
		button.Importance = widget.MediumImportance
		button.Icon = theme.QuestionIcon()
	} else {
		button.Importance = widget.HighImportance
	}

	return button
}

func newQueryEscapeButton(randomEscape bool, inputField *widget.Entry, copyButton *widget.Button, errOutField *widget.Entry, outputField *widget.Label) *widget.Button {
	buttonLabelText := func() string {
		if randomEscape {
			return "QueryEscape Randomly"
		}
		return "QueryEscape All"
	}()

	queryEscapeButton := newProcessInputButton(
		randomEscape,
		buttonLabelText,
		safelinks.QueryEscapeInput,
		inputField,
		copyButton,
		errOutField,
		outputField,
	)

	return queryEscapeButton
}

func newResetButton(w fyne.Window, inputField *widget.Entry, copyButton *widget.Button, errOutField *widget.Entry, outputField *widget.Label) *widget.Button {
	resetButton := widget.NewButton("Reset", func() {
		log.Println("Resetting application")
		w.Resize(fyne.NewSize(windowSizeHeight, windowSizeWidth))

		inputField.PlaceHolder = inputFieldPlaceholder
		inputField.Text = ""
		inputField.Refresh()

		errOutField.PlaceHolder = errOutPlaceholder
		errOutField.Text = ""
		errOutField.Refresh()

		outputField.Text = encodedOutputPlaceholder
		outputField.Refresh()

		copyButton.Disable()

		// Force garbage collection to free previously cached text.
		runtime.GC()
	})
	resetButton.Importance = widget.WarningImportance

	return resetButton
}

func newAboutButton(_ fyne.Window, inputField *widget.Entry, copyButton *widget.Button, errOutField *widget.Entry, outputField *widget.Label) *widget.Button {
	aboutButton := widget.NewButton("About", func() {
		log.Println("Displaying About text")
		inputField.PlaceHolder = "Description:\n\n" +
			"GUI tool used to create faux Safe Links encoded URLs for testing purposes"

		inputField.Text = ""
		inputField.Refresh()

		errOutField.Text = "..."
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
