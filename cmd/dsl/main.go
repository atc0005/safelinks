// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

package main

/*
	Misc resources reviewed/consulted while building this:

	https://www.reddit.com/r/golang/comments/fsxkqr/cancelling_blocking_read_from_stdin/
	https://play.golang.com/p/Ff8Grkobd7L
	https://github.com/golang/go/issues/24842#issuecomment-382384215
	https://superuser.com/questions/296969/re-enter-interactive-mode-after-ctrl-z
	https://stackoverflow.com/questions/34481065/break-out-of-input-scan
	https://stackoverflow.com/questions/43947363/detect-if-a-command-is-piped-or-not
	https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang
	https://stackoverflow.com/a/26567513/903870
	https://www.honeybadger.io/blog/a-definitive-guide-to-regular-expressions-in-go/

	https://pkg.go.dev/os/signal#NotifyContext
*/

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	userFeedbackOut := os.Stderr
	// debugLoggingOut := os.Stderr // switch to os.Stderr for debugging
	debugLoggingOut := io.Discard // use io.Discard for normal operation
	log.SetOutput(debugLoggingOut)

	// We wrap the os.Exit() behavior so that we can safely use deferred
	// cleanup functionality.
	var appExitCode int
	defer func(code *int) {
		var exitCode int
		if code != nil {
			exitCode = *code
		}
		os.Exit(exitCode)
	}(&appExitCode)

	ctx, cancel := context.WithCancel(context.Background())

	// OK duration for active input, but too short if switching to another
	// window to copy input text for pasting into this app.
	//
	// timerDuration := 5 * time.Second

	// A better duration but still might be a little too short.
	//
	// timerDuration := 10 * time.Second

	// Potentially too long if entering lines one by one, but a better fit if
	// processing bulk lines by copy/paste operations.
	timerDuration := 15 * time.Second

	// TODO: Consider using signal.NotifyContext in place of current shutdown
	// handling logic as this could simplify the implementation.
	//
	// https://pkg.go.dev/os/signal#NotifyContext
	signalChan, done, timer := setupShutdownHandling(timerDuration)

	defer func() {
		log.Println("Cleaning up signal catch behavior")
		signal.Stop(signalChan)
		cancel()
	}()

	go shutdownListener(cancel, userFeedbackOut, timer, signalChan, done)

	stat, _ := os.Stdin.Stat()
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
			userFeedbackOut,
			"Enter single or multi-line input. Press Ctrl-C to stop "+
				"(or wait %v for timeout).\n\n",
			timerDuration,
		)

		fmt.Fprintf(
			userFeedbackOut,
			"  - Feedback from this app is sent to stderr.\n"+
				"  - Decoding results are sent to stdout.\n"+
				"  - Tip: Redirect stdout to a file for multiple input lines.\n\n",
		)
	}

	resultsChan := make(chan string)
	errChan := make(chan error, 1)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	go pollInputSource(ctx, scanner, timer, timerDuration, resultsChan, errChan, done)

	for {
		select {
		case <-ctx.Done():
			log.Println("Context expired.")
			log.Println("TODO: Raise error condition?")

			return

		case result := <-resultsChan:
			fmt.Fprintln(os.Stdout, result)

		case err := <-errChan:
			fmt.Fprintln(userFeedbackOut, err.Error())
			appExitCode = 1

			return

		case <-done:
			log.Println("done channel activated. Exiting.")

			return
		}
	}
}
