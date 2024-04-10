// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
)

// Updated via Makefile builds. Setting placeholder value here so that
// something resembling a version string will be provided for non-Makefile
// builds.
var version = "x.y.z"

// Application metadata for easy reference.
const (
	myAppName        string = "eslg"
	myAppURL         string = "https://github.com/atc0005/safelinks"
	myAppDescription string = "GUI tool for encoding normal URLs within input text for testing purposes."
)

// GUI app constants.
const (
	windowSizeHeight float32 = 800
	windowSizeWidth  float32 = 650
)

// Constants used for field and button text.
const (
	inputFieldPlaceholder    string = "Paste text with unencoded URLs here and press Encode."
	errOutPlaceholder        string = "Encoding errors (if any) will be logged here. Text pasted here is ignored."
	errOutTryAgain           string = "Please insert text to encode and try again."
	encodedOutputPlaceholder string = "Encoded text will be placed here.\n\nChanges are overwritten upon button press."
)

// Version emits application name, version and repo location.
func Version() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}
