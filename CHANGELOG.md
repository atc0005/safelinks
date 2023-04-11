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

[Unreleased]: https://github.com/atc0005/safelinks/compare/v0.1.1...HEAD
[v0.1.1]: https://github.com/atc0005/safelinks/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/safelinks/releases/tag/v0.1.0
