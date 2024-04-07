// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package safelinks

import (
	"io"
	"log"
)

func init() {
	// Disable logging output by default.
	log.SetOutput(io.Discard)
}
