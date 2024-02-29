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
)
