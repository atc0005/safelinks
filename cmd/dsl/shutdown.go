// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
