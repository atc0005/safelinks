// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package safelinks

import "errors"

var (
	// ErrMissingValue indicates that an expected value was missing.
	ErrMissingValue = errors.New("missing expected value")

	// ErrInvalidURL indicates that an invalid URL was provided.
	ErrInvalidURL = errors.New("invalid URL provided")

	// ErrOriginalURLNotResolved indicates that we failed to resolve the
	// original URL using the given Safe Links URL.
	ErrOriginalURLNotResolved = errors.New("unable to resolve original URL")

	// ErrNoURLsFound indicates that an attempt to parse an input string for
	// URLs failed.
	ErrNoURLsFound = errors.New("no URLs matching requirements found in input")

	// ErrURLNotSafeLinkEncoded indicates that a given URL is not recognized
	// as using Safe Link encoding.
	ErrURLNotSafeLinkEncoded = errors.New("given URL not Safe Link encoded")

	// ErrNoSafeLinkURLsFound indicates that no URLs were found to be encoded
	// as Safe Links.
	ErrNoSafeLinkURLsFound = errors.New("no Safe Link URLs found in input")

	// ErrNoNonSafeLinkURLsFound indicates that no URLs were found to not
	// already be encoded as Safe Links.
	ErrNoNonSafeLinkURLsFound = errors.New("no non-Safe Link URLs found in input")

	// ErrQueryEscapingUnsuccessful indicates that an attempt to query escape
	// input was unsuccessful.
	ErrQueryEscapingUnsuccessful = errors.New("failed to query escape input")

	// ErrEncodingUnsuccessful indicates that an attempt to encode input was
	// unsuccessful.
	ErrEncodingUnsuccessful = errors.New("failed to encode input")
)
