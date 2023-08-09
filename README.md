<!-- omit in toc -->
# safelinks

Go-based tooling to manipulate (e.g., normalize/decode) Microsoft Office 365
"Safe Links" URLs.

[![Latest Release](https://img.shields.io/github/release/atc0005/safelinks.svg?style=flat-square)](https://github.com/atc0005/safelinks/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/safelinks.svg)](https://pkg.go.dev/github.com/atc0005/safelinks)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/safelinks)](https://github.com/atc0005/safelinks)
[![Lint and Build](https://github.com/atc0005/safelinks/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/safelinks/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/safelinks/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/safelinks/actions/workflows/project-analysis.yml)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
  - [`usl` CLI tool](#usl-cli-tool)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
  - [Deployment](#deployment)
- [Configuration](#configuration)
  - [Command-line arguments](#command-line-arguments)
    - [`usl`](#usl)
      - [Flags](#flags)
      - [Positional Argument](#positional-argument)
      - [Just Hit Enter](#just-hit-enter)
- [Examples](#examples)
  - [Using positional argument](#using-positional-argument)
  - [Using flag](#using-flag)
  - [Using input prompt](#using-input-prompt)
  - [Verbose output](#verbose-output)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

This repo is intended to provide various tools used to monitor processes.

| Tool Name | Overall Status | Description                                         |
| --------- | -------------- | --------------------------------------------------- |
| `usl`     | Alpha          | Small CLI tool for decoding a given Safe Links URL. |

## Features

### `usl` CLI tool

Small CLI tool for decoding a given Safe Links URL.

- Specify Safe Links URL via CLI argument or flag

- Optional verbose listing of query parameter values within a given Safe Links
  URL.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Microsoft Windows 10
- Ubuntu 20.04

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/safelinks`
   1. `cd safelinks`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - manually, explicitly specifying target OS and architecture
     - `GOOS=linux GOARCH=amd64 go build -mod=vendor ./cmd/usl/`
       - most likely this is what you want (if building manually)
       - substitute `amd64` with the appropriate architecture if using
         different hardware (e.g., `arm64` or `386`)
   - using Makefile `all` recipe
     - `make all`
       - generates x86 and x64 binaries
   - using Makefile `release-build` recipe
     - `make release-build`
       - generates the same release assets as provided by this project's
         releases
1. Locate generated binaries
   - if using `Makefile`
     - look in `/tmp/safelinks/release_assets/usl/`
   - if using `go build`
     - look in `/tmp/safelinks/`
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**: Depending on which `Makefile` recipe you use the generated binary
may be compressed and have an `xz` extension. If so, you should decompress the
binary first before deploying it (e.g., `xz -d usl-linux-amd64.xz`).

### Using release binaries

1. Download the [latest release][releases-url] binaries
1. Decompress binaries
   - e.g., `xz -d usl-linux-amd64.xz`
   - 7-Zip also works well for this on Windows systems (e.g., for systems
     without Git for Windows or WSL)
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

### Deployment

1. Place `usl` in a location where it can be easily accessed

## Configuration

### Command-line arguments

- Use the `-h` or `--help` flag to display current usage information.
- Flags marked as **`required`** must be set via CLI flag.
- Flags *not* marked as required are for settings where a useful default is
  already defined, but may be overridden if desired.

#### `usl`

##### Flags

| Flag           | Required | Default | Repeat | Possible       | Description                                                                   |
| -------------- | -------- | ------- | ------ | -------------- | ----------------------------------------------------------------------------- |
| `h`, `help`    | No       | `false` | No     | `h`, `help`    | Show Help text along with the list of supported flags.                        |
| `version`      | No       | `false` | No     | `version`      | Whether to display application version and then immediately exit application. |
| `v`, `verbose` | No       | `false` | No     | `v`, `verbose` | Display additional information about a given Safe Links URL.                  |
| `u`, `url`     |

##### Positional Argument

A URL pattern is accepted as a single positional argument in place of the `u`
or `url` flag. It is recommended that you quote the URL pattern to help
prevent some of the characters from being interpreted as shell commands (e.g.,
`&` as an attempt to background a command).

##### Just Hit Enter

The `usl` tool can also be called without any flags or positional argument. In
this scenario it will prompt you to insert/paste the URL pattern (quoted or
otherwise).

## Examples

Though probably not required for all terminals, we quote the Safe Links URL to
prevent unintended interpretation of characters in the URL.

### Using positional argument

```console
$ ./usl 'SafeLinksURLHere'

Original URL:

https://go.dev/dl/
```

### Using flag

```console
$ ./usl --url 'SafeLinksURLHere'

Original URL:

https://go.dev/dl/
```

### Using input prompt

In this example we just press enter so that we will be prompted for the input
URL pattern.

```console
$ ./usl
Enter URL: SafeLinksURLHere

Original URL:

https://go.dev/dl/
```

### Verbose output

```console
$ ./usl --verbose --url 'SafeLinksURLHere'

Expanded values from the given link:

  data      : PLACEHOLDER
  host      : nam99.safelinks.protection.outlook.com
  reserved  : 0
  sdata     : PLACEHOLDER
  url       : https://go.dev/dl/
```

## License

See the [LICENSE](LICENSE) file for details.

## References

- <https://learn.microsoft.com/en-us/microsoft-365/security/office-365-security/safe-links-about>
- <https://security.stackexchange.com/questions/230309/is-a-safelinks-protection-outlook-com-link-phishing>
- <https://techcommunity.microsoft.com/t5/security-compliance-and-identity/data-sdata-and-reserved-parameters-in-office-atp-safelinks/td-p/1637050>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/safelinks>  "This project's GitHub repo"

[releases-url]: <https://github.com/atc0005/safelinks/releases> "This project's releases"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"
