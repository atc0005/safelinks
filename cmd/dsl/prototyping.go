// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/atc0005/safelinks/internal/safelinks"
)

// getAppTimeout is a small helper function that provides the application
// timeout value. This is mostly a placeholder for future logic to return the
// default timeout value if the user does not specify one of their own.
func getAppTimeout() time.Duration {
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

	return timerDuration
}

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

// shutdownListener listens to the given input sources for an indication that
// it is time to shutdown the application:
//
//   - listen to the given quit channel for a known termination signal from
//     the user or the OS
//   - watch the given timer for inactivity expiration
//
// When a timeout is reached or a signal is received the given
// context.CancelFunc is called and a done signal is returned to commence
// application termination.
func shutdownListener(
	cancel context.CancelFunc,
	w io.Writer,
	timer *time.Timer,
	quit <-chan os.Signal,
	done chan<- bool,
) {

	var gracefulShutdownSignal bool

	gracefulExitMsg := "Exiting as requested."
	failureExitMsg := "Unexpected exit signal. Aborting application."

	select {
	case <-timer.C:
		fmt.Fprintln(w, "Timeout reached. Exiting application.")

		log.Println("Calling cancel func")
		cancel()

		return

	case sig := <-quit:
		log.Println("Signal received:", sig)

		switch sig {
		case syscall.SIGINT:
			gracefulShutdownSignal = true
		case syscall.SIGTERM:
			gracefulShutdownSignal = true
		case syscall.SIGHUP:
			gracefulShutdownSignal = true
		case syscall.SIGQUIT:
			gracefulShutdownSignal = true
		case syscall.SIGPIPE:
			gracefulShutdownSignal = false
		default:
			gracefulShutdownSignal = false
		}

		if gracefulShutdownSignal {
			fmt.Fprintln(w, gracefulExitMsg)
		} else {
			fmt.Fprintln(w, failureExitMsg)
		}

		done <- true

		return
	}

}

// setupShutdownHandling uses the given timeout duration to initialize the
// timer and channels used to control application shutdown behavior.
func setupShutdownHandling(timeout time.Duration) (chan os.Signal, chan bool, *time.Timer) {
	// Override default Go handling of specified signals in order to customize
	// the shutdown process.
	signalChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	timer := time.NewTimer(timeout)

	// Asynchronous signals: triggered from the kernel or another app.
	//
	// https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
	// https://en.wikipedia.org/wiki/Signal.h
	signal.Notify(signalChan,
		// syscall.SIGKILL,  // unable to catch this signal
		// syscall.SIGSTOP,  // unable to catch this signal

		// The "program interrupt" or "stop" signal is sent when the user
		// types the INTR character (normally ctrl-c).
		syscall.SIGINT,

		// The SIGTERM signal is a generic signal used to cause program
		// termination. Unlike SIGKILL, this signal can be blocked, handled,
		// and ignored. It is the normal way to politely ask a program to
		// terminate.
		//
		// The shell command kill generates SIGTERM by default.
		syscall.SIGTERM,

		// The SIGHUP signal is often used to soft "reload" an application
		// (e.g., asking it to close open files and reload its configuration).
		syscall.SIGHUP,

		// The SIGQUIT signal is similar to SIGINT, except that it's
		// controlled by the QUIT character, usually Ctrl-\ and produces a
		// core dump when it terminates the process. You can think of this as
		// a program error condition "detected" by the user.
		syscall.SIGQUIT,

		// The SIGPIPE signal is is raised by the kernel when standard output
		// from a program is written to a broken pipe. This allows the
		// application to output an error regarding the broken pipe.
		syscall.SIGPIPE,
	)

	return signalChan, done, timer
}

// pollInputSource reads from the given io.Reader until the input is exhausted
// or the given inactivity timer expires.
func pollInputSource(
	ctx context.Context,
	r io.Reader,
	inactivityTimer *time.Timer,
	timerDuration time.Duration,
	resultsChan chan<- string,
	errChan chan<- error,
	done chan<- bool,
) {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// Ctrl-D (UNIX) or Ctrl-Z (Windows) will send EOF which will abort the
	// scanner.
	// https://stackoverflow.com/questions/34481065/break-out-of-input-scan
	for scanner.Scan() {
		if !inactivityTimer.Stop() {
			<-inactivityTimer.C
		}

		// Reset inactivity timer on input read.
		inactivityTimer.Reset(timerDuration)

		log.Println("New scanner loop")

		select {
		case <-ctx.Done():
			log.Println("Context expired. Aborting.")
			return
		case <-inactivityTimer.C:
			log.Println("Timer expired. Aborting.")
			return
		default:
		}

		txt := scanner.Text()

		// Skip text parsing if nothing present.
		if strings.TrimSpace(txt) == "" {
			log.Println("Skipping text parsing for empty or whitespace only input")
			resultsChan <- txt
			continue
		}

		processInput(txt, resultsChan, errChan)

	}

	log.Println("Checking scanner.Err()")
	switch err := scanner.Err(); {
	case err != nil:
		errChan <- fmt.Errorf(
			"error reading URLs: %w",
			err,
		)

		return

	default:
		done <- true
	}
}

// processInput processes given input replacing any Safe Links encoded URL
// with the original value. Other input is returned unmodified.
func processInput(txt string, resultsChan chan<- string, errChan chan<- error) {
	log.Println("Calling safelinks.SafeLinkURLs(txt)")
	safeLinks, err := safelinks.SafeLinkURLs(txt)

	// Failing to find a URL in the input is considered OK. Other errors
	// result in aborting the decode attempt.
	switch {
	case errors.Is(err, safelinks.ErrNoURLsFound):
		resultsChan <- txt
	case errors.Is(err, safelinks.ErrNoSafeLinkURLsFound):
		resultsChan <- txt
	case err != nil:
		errChan <- err

		return
	default:
		modifiedInput := txt

		for _, sl := range safeLinks {
			modifiedInput = strings.Replace(modifiedInput, sl.EncodedURL, sl.DecodedURL, 1)
		}

		resultsChan <- modifiedInput
	}
}
