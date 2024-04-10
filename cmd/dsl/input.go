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
	"strings"
	"time"

	"github.com/atc0005/safelinks/internal/safelinks"
)

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
	log.Println("Calling safelinks.DecodeInput")
	modifiedInput, err := safelinks.DecodeInput(txt)

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
		resultsChan <- modifiedInput
	}
}
