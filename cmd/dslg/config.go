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
	myAppName        string = "usl"
	myAppURL         string = "https://github.com/atc0005/" + myAppName
	myAppDescription string = "Go-based tooling to manipulate (e.g., normalize/decode) Microsoft Office 365 \"Safe Links\" URLs."
)

// GUI app constants.
const (
	windowSizeHeight float32 = 800
	windowSizeWidth  float32 = 650
)

const (
	safeLinksAboutURL string = "https://learn.microsoft.com/en-us/microsoft-365/security/office-365-security/safe-links-about"
)

// Version emits application name, version and repo location.
func Version() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}