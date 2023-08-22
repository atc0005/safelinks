# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/safelinks/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.2.4] - 2023-08-22

### Changed

- Dependencies
  - `Go`
    - `1.19.12` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.13.2` to `go-ci-oldstable-build-v0.13.5`
- (GH-91) Update project to Go 1.20 series

### Fixed

- (GH-85) README: Fix verbose flag description
- (GH-86) README: Add missing coverage for url flag

## [v0.2.3] - 2023-08-09

### Added

- (GH-62) Add initial automated release notes config
- (GH-64) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.19.12`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.4` to `go-ci-oldstable-build-v0.13.2`
- (GH-66) Update Dependabot config to monitor both branches

### Fixed

- (GH-80) Makefile: Fix version output

## [v0.2.2] - 2023-07-14

### Overview

- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.10` to `1.19.11`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.0` to `go-ci-oldstable-build-v0.11.4`

### Fixed

- (GH-55) Update README to link to releases page
- (GH-57) Remove deploy logic from postinstall scripts

## [v0.2.1] - 2023-06-20

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.8` to `1.19.10`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.6` to `go-ci-oldstable-build-v0.11.0`
- (GH-48) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-45) Disable depguard linter
- (GH-49) Restore local CodeQL workflow
- (GH-51) README: adjust section header, note input prompt

## [v0.2.0] - 2023-04-11

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-37) Generate RPM/DEB packages using nFPM
- (GH-36) Add version details to Windows executables

### Changed

- (GH-38) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-35) Makefile: Compress binaries & use fixed filenames
- (GH-34) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-33) Build dev/stable releases using go-ci Docker image

## [v0.1.1] - 2023-04-11

### Overview

- Bug fixes
- GitHub Actions workflow updates
- Dependency updates
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-22) Add Go Module Validation, Dependency Updates jobs

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.8`
- CI
  - (GH-24) Drop `Push Validation` workflow
  - (GH-25) Rework workflow scheduling
  - (GH-27) Remove `Push Validation` workflow status badge

### Fixed

- (GH-29) Update vuln analysis GHAW to use on.push hook

## [v0.1.0] - 2022-12-16

### Overview

- Initial release
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

Initial release!

This release provides an early release version of one new tool:

| Tool Name | Overall Status | Description                                         |
| --------- | -------------- | --------------------------------------------------- |
| `usl`     | Alpha          | Small CLI tool for decoding a given Safe Links URL. |

See the project README for additional details.

[Unreleased]: https://github.com/atc0005/safelinks/compare/v0.2.4...HEAD
[v0.2.4]: https://github.com/atc0005/safelinks/releases/tag/v0.2.4
[v0.2.3]: https://github.com/atc0005/safelinks/releases/tag/v0.2.3
[v0.2.2]: https://github.com/atc0005/safelinks/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/safelinks/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/safelinks/releases/tag/v0.2.0
[v0.1.1]: https://github.com/atc0005/safelinks/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/safelinks/releases/tag/v0.1.0
