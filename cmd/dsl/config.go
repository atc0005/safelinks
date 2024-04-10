// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import "time"

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
