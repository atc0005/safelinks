// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package safelinks

import "errors"

var (
	// ErrInvalidURL indicates that an invalid URL was provided.
	ErrInvalidURL = errors.New("invalid URL provided")

	// ErrOriginalURLNotResolved indicates that we failed to resolve the
	// original URL using the given Safe Links URL.
	ErrOriginalURLNotResolved = errors.New("unable to resolve original URL")

	// ErrNoURLsFound indicates that an attempt to parse an input string for
	// URLs failed.
	ErrNoURLsFound = errors.New("no URLs found in input")

	// ErrURLNotSafeLinkEncoded indicates that a given URL is not recognized
	// as using Safe Link encoding.
	ErrURLNotSafeLinkEncoded = errors.New("given URL not Safe Link encoded")

	// ErrNoSafeLinkURLsFound indicates that no URLs were found to be encoded
	// as Safe Links.
	ErrNoSafeLinkURLsFound = errors.New("no Safe Link URLs found in input")
)
