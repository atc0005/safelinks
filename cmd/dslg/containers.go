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

func newOutputContainer(objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(
		layout.NewVBoxLayout(),
		objects...,
	)
}

func newMainAppContainer(inputField *widget.Entry, buttonContainer *fyne.Container, outputContainer *fyne.Container) *fyne.Container {
	return container.NewBorder(inputField, buttonContainer, nil, nil, outputContainer)
}
