// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// showAppUsageInfo provides usage information to the user if input is not
// received via a pipe. The given timer duration is listed so that the user is
// aware of the timeout behavior.
func showAppUsageInfo(appInput *os.File, timerDuration time.Duration, appOutput io.Writer) {
	stat, _ := appInput.Stat()
	switch {
	case (stat.Mode() & os.ModeCharDevice) == 0:
		// Input was from a pipe, so do not provide usage information.

	default:
		// Ctrl-D (UNIX) or Ctrl-Z (Windows) also trigger shutdown behavior
		// but we do not advertise those keystrokes for simplicity.
		//
		// For example, unintentionally using the Ctrl-Z keystroke on a Linux
		// distro will put the process into the background in what seems like
		// a "hung" state. Running "fg" will return the process to an
		// interactive state.
		//
		// To keep things simple, it is best to only advertise Ctrl-C or
		// waiting for the configured timeout to stop input processing.
		fmt.Fprintf(
			appOutput,
			"Enter single or multi-line input. Press Ctrl-C to stop "+
				"(or wait %v for timeout).\n\n",
			timerDuration,
		)

		fmt.Fprintf(
			appOutput,
			"  - Feedback from this app is sent to stderr.\n"+
				"  - Decoding results are sent to stdout.\n"+
				"  - Tip: Redirect stdout to a file for multiple input lines.\n\n",
		)
	}
}
