// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// NewButtonColumnContainer creates a container for given buttons using a
// vertical layout.
func NewButtonColumnContainer(buttons ...fyne.CanvasObject) *fyne.Container {
	return container.New(
		layout.NewVBoxLayout(),
		buttons...,
	)
}

// NewButtonRowContainer creates a container for given buttons using a
// horizontal layout.
func NewButtonRowContainer(buttons ...fyne.CanvasObject) *fyne.Container {
	return container.New(
		layout.NewHBoxLayout(),
		buttons...,
	)
}

// NewOutputContainer creates a container holding error output at the top and
// decoded output at the bottom with a spacer between them to prevent the
// decoded output from expanding vertically beyond the current application
// window size.
//
// The resulting "output" container is intended for display in the center of
// another container.
func NewOutputContainer(errorOutput fyne.CanvasObject, decodedOutput fyne.CanvasObject) *fyne.Container {
	outputLabelContainer := container.NewVScroll(decodedOutput)
	outputContainer := container.NewBorder(errorOutput, layout.NewSpacer(), nil, nil, outputLabelContainer)

	return outputContainer
}

// NewMainAppContainer creates an container that places an input field at the
// top, a button row container across the bottom and an "output" container in
// the center holding error output and decoded output. This container is
// intended to flex giving precedence to decoded output but without allowing
// it to expand beyond the current window size.
func NewMainAppContainer(inputField *widget.Entry, buttonContainer *fyne.Container, outputContainer *fyne.Container) *fyne.Container {
	return container.NewBorder(inputField, buttonContainer, nil, nil, outputContainer)
}
