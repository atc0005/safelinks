// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"log"
	"strings"

	"github.com/atc0005/safelinks/internal/safelinks"
)

// decodeInput processes given input replacing any Safe Links encoded URL
// with the original decoded value. Other input is returned unmodified.
func decodeInput(txt string) (string, error) {
	log.Println("Calling safelinks.SafeLinkURLs(txt)")

	safeLinks, err := safelinks.SafeLinkURLs(txt)
	if err != nil {
		return "", err
	}

	modifiedInput := txt
	for _, sl := range safeLinks {
		modifiedInput = strings.Replace(modifiedInput, sl.EncodedURL, sl.DecodedURL, 1)
	}

	return modifiedInput, nil

}
