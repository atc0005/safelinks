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

func pollInputSource(
	ctx context.Context,
	scanner *bufio.Scanner,
	timer *time.Timer,
	timerDuration time.Duration,
	resultsChan chan<- string,
	errChan chan<- error,
	done chan<- bool,
) {

	// Ctrl-D (UNIX) or Ctrl-Z (Windows) will send EOF which will abort the
	// scanner.
	// https://stackoverflow.com/questions/34481065/break-out-of-input-scan
	for scanner.Scan() {
		// Reset timer on activity.
		if !timer.Stop() {
			<-timer.C
		}
		timer.Reset(timerDuration)

		log.Println("New scanner loop")

		select {
		case <-ctx.Done():
			log.Println("Context expired. Aborting.")
			return
		case <-timer.C:
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

		log.Println("Calling safelinks.SafeLinkURLs(txt)")
		safeLinks, err := safelinks.SafeLinkURLs(txt)

		// Failing to find a URL in the input is considered OK. Other errors
		// result in aborting the decode attempt.
		//
		// TODO: This behavior needs further testing.
		//
		// It's likely that we will wish to continue processing further lines
		// and not abort early.
		//
		switch {
		case errors.Is(err, safelinks.ErrNoURLsFound):
			resultsChan <- txt
		case errors.Is(err, safelinks.ErrNoSafeLinkURLsFound):
			resultsChan <- txt
		case err != nil:
			errChan <- err

			return
		default:
			// TESTING
			// fmt.Printf("%d Safe Links:\n", len(safeLinks))
			// for _, sl := range safeLinks {
			// 	fmt.Printf("\tOriginal: %s\n\tDecoded: %s\n\n", sl.EncodedURL, sl.DecodedURL)
			// }

			modifiedInput := txt

			for _, sl := range safeLinks {
				modifiedInput = strings.Replace(modifiedInput, sl.EncodedURL, sl.DecodedURL, 1)
			}

			resultsChan <- modifiedInput
		}

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
