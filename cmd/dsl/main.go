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
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
)

func main() {
	userFeedbackOut := os.Stderr

	// use io.Discard for normal operation
	// switch to os.Stderr for debugging
	debugLoggingOut := io.Discard

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
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

	timerDuration := getAppTimeout()

	// TODO: Consider using signal.NotifyContext in place of current shutdown
	// handling logic as this could simplify the implementation.
	//
	// https://pkg.go.dev/os/signal#NotifyContext
	signalChan, done, timer := setupShutdownHandling(timerDuration)

	defer func() {
		log.Println("Cleaning up signal catch behavior")
		signal.Stop(signalChan)

		log.Println("Cancelling base context")
		cancel()
	}()

	go shutdownListener(cancel, userFeedbackOut, timer, signalChan, done)

	resultsChan := make(chan string)
	errChan := make(chan error, 1)

	go pollInputSource(ctx, os.Stdin, timer, timerDuration, resultsChan, errChan, done)

	showAppUsageInfo(os.Stdin, timerDuration, userFeedbackOut)

	for {
		select {
		case <-ctx.Done():
			log.Println("Context expired.")

			return

		case result := <-resultsChan:
			_, _ = fmt.Fprintln(os.Stdout, result)

		case err := <-errChan:
			_, _ = fmt.Fprintln(userFeedbackOut, err.Error())
			appExitCode = 1

			return

		case <-done:
			log.Println("done channel activated. Exiting.")

			return
		}
	}
}
