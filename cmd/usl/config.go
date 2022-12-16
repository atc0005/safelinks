// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/safelinks
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"flag"
	"fmt"
	"os"
)

// Updated via Makefile builds. Setting placeholder value here so that
// something resembling a version string will be provided for non-Makefile
// builds.
var version = "x.y.z"

const myAppName string = "usl"
const myAppURL string = "https://github.com/atc0005/" + myAppName

type config struct {
	URL     string
	Verbose bool
	Version bool
}

func NewConfig() *config {
	var cfg config
	setupFlags(&cfg)

	return &cfg
}

// Version emits application name, version and repo location.
func Version() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}

// usage is a custom override for the default Help text provided by the flag
// package. Here we prepend some additional metadata to the existing output.
func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), "\n"+Version()+"\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func setupFlags(c *config) {
	flag.StringVar(&c.URL, "url", "", "Safe Links URL to decode")
	flag.StringVar(&c.URL, "u", "", "Safe Links URL to decode"+" (shorthand)")
	flag.BoolVar(&c.Verbose, "verbose", false, "Display additional information about a given Safe Links URL")
	flag.BoolVar(&c.Verbose, "v", false, "Display additional information about a given Safe Links URL"+" (shorthand)")
	flag.BoolVar(&c.Version, "version", false, "Display version information and immediately exit")
	flag.Usage = usage
	flag.Parse()
}
