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

## [v0.5.16] - 2025-05-20

### Changed

#### Dependency Updates

- (GH-828) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.22.9 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.22.10 in /dependabot/docker/builds/x64
- (GH-825) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.22.9 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.22.10 in /dependabot/docker/builds/x86
- (GH-823) Go Dependency: Bump fyne.io/fyne/v2 from 2.6.0 to 2.6.1
- (GH-818) Go Dependency: Bump golang.org/x/image from 0.26.0 to 0.27.0
- (GH-816) Go Dependency: Bump golang.org/x/net from 0.39.0 to 0.40.0
- (GH-819) Go Dependency: Bump golang.org/x/sys from 0.32.0 to 0.33.0
- (GH-817) Go Dependency: Bump golang.org/x/text from 0.24.0 to 0.25.0
- (GH-821) Go Runtime: Bump golang from 1.23.8 to 1.23.9 in /dependabot/docker/go

## [v0.5.15] - 2025-05-02

### Changed

#### Dependency Updates

- (GH-805) Go Dependency: Bump github.com/yuin/goldmark from 1.7.10 to 1.7.11
- (GH-801) Go Dependency: Bump github.com/yuin/goldmark from 1.7.8 to 1.7.9
- (GH-804) Go Dependency: Bump github.com/yuin/goldmark from 1.7.9 to 1.7.10
- (GH-808) Update `jeandeaual/go-locale` pseudo-version

## [v0.5.14] - 2025-04-14

### Changed

#### Dependency Updates

- (GH-792) Go Dependency: Bump fyne.io/fyne/v2 from 2.5.5 to 2.6.0
- (GH-795) Go Dependency: Bump github.com/hack-pad/safejs from 0.1.0 to 0.1.1

### Fixed

- (GH-797) Fix deprecated func use after fyne v2.6.0 update

## [v0.5.13] - 2025-04-09

### Changed

#### Dependency Updates

- (GH-765) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.22.3 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.22.9 in /dependabot/docker/builds/x64
- (GH-763) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.22.3 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.22.9 in /dependabot/docker/builds/x86
- (GH-770) Disable Dependabot automatic PR rebasing
- (GH-746) Go Dependency: Bump fyne.io/fyne/v2 from 2.5.4 to 2.5.5
- (GH-748) Go Dependency: Bump github.com/BurntSushi/toml from 1.4.0 to 1.5.0
- (GH-781) Go Dependency: Bump github.com/fsnotify/fsnotify from 1.8.0 to 1.9.0
- (GH-714) Go Dependency: Bump github.com/fyne-io/glfw-js from 0.1.0 to 0.2.0
- (GH-768) Go Dependency: Bump github.com/fyne-io/image from 0.1.0 to 0.1.1
- (GH-710) Go Dependency: Bump github.com/go-text/typesetting from 0.2.1 to 0.3.0
- (GH-709) Go Dependency: Bump github.com/google/go-cmp from 0.6.0 to 0.7.0
- (GH-780) Go Dependency: Bump github.com/nicksnyder/go-i18n/v2 from 2.5.1 to 2.6.0
- (GH-716) Go Dependency: Bump github.com/rymdport/portal from 0.4.0 to 0.4.1
- (GH-731) Go Dependency: Bump golang.org/x/image from 0.24.0 to 0.25.0
- (GH-778) Go Dependency: Bump golang.org/x/image from 0.25.0 to 0.26.0
- (GH-755) Go Dependency: Bump golang.org/x/net from 0.35.0 to 0.38.0
- (GH-783) Go Dependency: Bump golang.org/x/net from 0.38.0 to 0.39.0
- (GH-734) Go Dependency: Bump golang.org/x/sys from 0.30.0 to 0.31.0
- (GH-775) Go Dependency: Bump golang.org/x/sys from 0.31.0 to 0.32.0
- (GH-733) Go Dependency: Bump golang.org/x/text from 0.22.0 to 0.23.0
- (GH-779) Go Dependency: Bump golang.org/x/text from 0.23.0 to 0.24.0
- (GH-760) Go Runtime: Bump golang from 1.22.12 to 1.23.8 in /dependabot/docker/go
- (GH-744) go.mod: update minimum Go version to 1.23.0
- (GH-788) Update `github.com/go-gl/glfw/v3.3/glfw`
- (GH-786) Update `golang.org/x/mobile` pseudo-version
- (GH-721) Update project to Go 1.23 series

### Fixed

- (GH-767) Fix `copyloopvar` linting errors

## [v0.5.12] - 2025-02-20

### Changed

#### Dependency Updates

- (GH-699) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.18 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.22.3 in /dependabot/docker/builds/x64
- (GH-698) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.18 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.22.3 in /dependabot/docker/builds/x86
- (GH-677) Go Dependency: Bump fyne.io/fyne/v2 from 2.5.3 to 2.5.4
- (GH-697) Go Dependency: Bump github.com/fyne-io/gl-js from 0.0.0-20230506162202-1fdaa286a934 to 0.1.0
- (GH-696) Go Dependency: Bump github.com/fyne-io/glfw-js from 0.0.0-20241126112943-313d8a0fe1d0 to 0.1.0
- (GH-695) Go Dependency: Bump github.com/fyne-io/image from 0.0.0-20240417123036-dc0ee9e7c964 to 0.1.0
- (GH-678) Go Dependency: Bump github.com/nicksnyder/go-i18n/v2 from 2.4.1 to 2.5.1
- (GH-702) Go Dependency: Bump github.com/rymdport/portal from 0.3.0 to 0.4.0
- (GH-685) Go Dependency: Bump golang.org/x/image from 0.23.0 to 0.24.0
- (GH-700) Go Dependency: Bump golang.org/x/net from 0.32.0 to 0.35.0
- (GH-684) Go Dependency: Bump golang.org/x/sys from 0.28.0 to 0.30.0
- (GH-683) Go Dependency: Bump golang.org/x/text from 0.21.0 to 0.22.0
- (GH-681) Go Runtime: Bump golang from 1.22.10 to 1.22.12 in /dependabot/docker/go

## [v0.5.11] - 2024-12-18

### Changed

#### Dependency Updates

- (GH-635) Go Dependency: Bump fyne.io/fyne/v2 from 2.5.2 to 2.5.3
- (GH-640) Go Dependency: Bump github.com/go-text/typesetting from 0.2.0 to 0.2.1
- (GH-638) Update `golang.org/x/mobile` pseudo-version
- (GH-642) Update `jeandeaual/go-locale` pseudo-version

## [v0.5.10] - 2024-12-06

### Changed

#### Dependency Updates

- (GH-603) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.15 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.16 in /dependabot/docker/builds/x64
- (GH-607) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.16 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.17 in /dependabot/docker/builds/x64
- (GH-617) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.17 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.18 in /dependabot/docker/builds/x64
- (GH-604) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.15 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.16 in /dependabot/docker/builds/x86
- (GH-609) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.16 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.17 in /dependabot/docker/builds/x86
- (GH-618) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.17 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.18 in /dependabot/docker/builds/x86
- (GH-605) Go Dependency: Bump github.com/stretchr/testify from 1.9.0 to 1.10.0
- (GH-624) Go Dependency: Bump golang.org/x/image from 0.22.0 to 0.23.0
- (GH-622) Go Dependency: Bump golang.org/x/net from 0.31.0 to 0.32.0
- (GH-621) Go Dependency: Bump golang.org/x/sys from 0.27.0 to 0.28.0
- (GH-623) Go Dependency: Bump golang.org/x/text from 0.20.0 to 0.21.0
- (GH-611) Go Runtime: Bump golang from 1.22.9 to 1.22.10 in /dependabot/docker/go
- (GH-629) Update `github.com/fyne-io/glfw-js` pseudo-version
- (GH-631) Update `golang.org/x/mobile` pseudo-version
- (GH-630) Update `jeandeaual/go-locale` pseudo-version

## [v0.5.9] - 2024-11-13

### Changed

#### Dependency Updates

- (GH-591) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.14 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.15 in /dependabot/docker/builds/x64
- (GH-586) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.14 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.15 in /dependabot/docker/builds/x86
- (GH-576) Go Dependency: Bump github.com/fsnotify/fsnotify from 1.7.0 to 1.8.0
- (GH-589) Go Dependency: Bump github.com/rymdport/portal from 0.2.6 to 0.3.0
- (GH-584) Go Dependency: Bump golang.org/x/image from 0.21.0 to 0.22.0
- (GH-590) Go Dependency: Bump golang.org/x/net from 0.30.0 to 0.31.0
- (GH-582) Go Dependency: Bump golang.org/x/sys from 0.26.0 to 0.27.0
- (GH-583) Go Dependency: Bump golang.org/x/text from 0.19.0 to 0.20.0
- (GH-578) Go Runtime: Bump golang from 1.22.8 to 1.22.9 in /dependabot/docker/go
- (GH-595) Update `golang.org/x/mobile` pseudo-version

## [v0.5.8] - 2024-10-18

### Changed

#### Dependency Updates

- (GH-547) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.13 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.14 in /dependabot/docker/builds/x64
- (GH-554) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.13 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.14 in /dependabot/docker/builds/x86
- (GH-559) Go Dependency: Bump fyne.io/fyne/v2 from 2.5.1 to 2.5.2
- (GH-555) Go Dependency: Bump github.com/nicksnyder/go-i18n/v2 from 2.4.0 to 2.4.1
- (GH-566) Go Dependency: Bump github.com/yuin/goldmark from 1.7.4 to 1.7.8
- (GH-546) Go Dependency: Bump golang.org/x/image from 0.20.0 to 0.21.0
- (GH-545) Go Dependency: Bump golang.org/x/net from 0.29.0 to 0.30.0
- (GH-543) Go Dependency: Bump golang.org/x/sys from 0.25.0 to 0.26.0
- (GH-544) Go Dependency: Bump golang.org/x/text from 0.18.0 to 0.19.0
- (GH-542) Go Runtime: Bump golang from 1.22.7 to 1.22.8 in /dependabot/docker/go
- (GH-569) Update `golang.org/x/mobile` pseudo-version
- (GH-571) Update `golang.org/x/mobile` pseudo-version

## [v0.5.7] - 2024-10-16

### Changed

#### Dependency Updates

- (GH-527) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.11 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.12 in /dependabot/docker/builds/x64
- (GH-534) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.12 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.13 in /dependabot/docker/builds/x64
- (GH-511) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.9 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.11 in /dependabot/docker/builds/x64
- (GH-529) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.11 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.12 in /dependabot/docker/builds/x86
- (GH-535) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.12 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.13 in /dependabot/docker/builds/x86
- (GH-514) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.9 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.11 in /dependabot/docker/builds/x86
- (GH-510) Go Dependency: Bump fyne.io/fyne/v2 from 2.5.0 to 2.5.1
- (GH-530) Go Dependency: Bump github.com/go-text/render from 0.1.1-0.20240418202334-dd62631dae9b to 0.1.1
- (GH-518) Go Dependency: Bump golang.org/x/image from 0.19.0 to 0.20.0
- (GH-524) Go Dependency: Bump golang.org/x/net from 0.28.0 to 0.29.0
- (GH-517) Go Dependency: Bump golang.org/x/sys from 0.24.0 to 0.25.0
- (GH-519) Go Dependency: Bump golang.org/x/text from 0.17.0 to 0.18.0
- (GH-525) Go Runtime: Bump golang from 1.22.6 to 1.22.7 in /dependabot/docker/go
- (GH-539) Update `golang.org/x/mobile` pseudo-version
- (GH-538) Update project Go version to 1.22.0

## [v0.5.6] - 2024-08-22

### Changed

#### Dependency Updates

- (GH-499) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.8 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.9 in /dependabot/docker/builds/x64
- (GH-498) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.8 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.9 in /dependabot/docker/builds/x86
- (GH-501) Go Runtime: Bump golang from 1.21.13 to 1.22.6 in /dependabot/docker/go
- (GH-500) Update project to Go 1.22 series

## [v0.5.5] - 2024-08-13

### Changed

#### Dependency Updates

- (GH-461) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.5 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.6 in /dependabot/docker/builds/x64
- (GH-462) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.6 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.7 in /dependabot/docker/builds/x64
- (GH-483) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.7 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.8 in /dependabot/docker/builds/x64
- (GH-459) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.5 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.6 in /dependabot/docker/builds/x86
- (GH-463) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.6 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.7 in /dependabot/docker/builds/x86
- (GH-479) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.7 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.8 in /dependabot/docker/builds/x86
- (GH-478) Go Dependency: Bump github.com/rymdport/portal from 0.2.3 to 0.2.4
- (GH-487) Go Dependency: Bump github.com/rymdport/portal from 0.2.4 to 0.2.6
- (GH-472) Go Dependency: Bump golang.org/x/image from 0.18.0 to 0.19.0
- (GH-473) Go Dependency: Bump golang.org/x/net from 0.27.0 to 0.28.0
- (GH-469) Go Dependency: Bump golang.org/x/sys from 0.22.0 to 0.23.0
- (GH-485) Go Dependency: Bump golang.org/x/sys from 0.23.0 to 0.24.0
- (GH-471) Go Dependency: Bump golang.org/x/text from 0.16.0 to 0.17.0
- (GH-477) Go Runtime: Bump golang from 1.21.12 to 1.21.13 in /dependabot/docker/go
- (GH-490) Update `golang.org/x/mobile` pseudo-version

#### Other

- (GH-466) Push `REPO_VERSION` var into containers for builds

## [v0.5.4] - 2024-07-17

### Added

- (GH-449) Add cases for single URL with(out) angle brackets

### Changed

#### Dependency Updates

- (GH-445) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.4 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.5 in /dependabot/docker/builds/x64
- (GH-443) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.4 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.5 in /dependabot/docker/builds/x86
- (GH-438) Go Dependency: Bump fyne.io/fyne/v2 from 2.4.5 to 2.5.0
- (GH-447) Go Dependency: Bump github.com/rymdport/portal from 0.2.2 to 0.2.3

### Fixed

- (GH-450) Fix decoding of Markdown formatted URLs

## [v0.5.3] - 2024-07-10

### Changed

#### Dependency Updates

- (GH-395) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.7 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.8 in /dependabot/docker/builds/x64
- (GH-415) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.8 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.2 in /dependabot/docker/builds/x64
- (GH-420) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.2 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.3 in /dependabot/docker/builds/x64
- (GH-425) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.3 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.21.4 in /dependabot/docker/builds/x64
- (GH-394) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.7 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.8 in /dependabot/docker/builds/x86
- (GH-404) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.8 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.2 in /dependabot/docker/builds/x86
- (GH-417) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.2 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.3 in /dependabot/docker/builds/x86
- (GH-422) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.3 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.21.4 in /dependabot/docker/builds/x86
- (GH-400) Go Dependency: Bump fyne.io/systray from 1.10.1-0.20231115130155-104f5ef7839e to 1.11.0
- (GH-399) Go Dependency: Bump github.com/yuin/goldmark from 1.7.1 to 1.7.2
- (GH-401) Go Dependency: Bump github.com/yuin/goldmark from 1.7.2 to 1.7.3
- (GH-411) Go Dependency: Bump github.com/yuin/goldmark from 1.7.3 to 1.7.4
- (GH-410) Go Dependency: Bump golang.org/x/image from 0.17.0 to 0.18.0
- (GH-430) Go Dependency: Bump golang.org/x/net from 0.26.0 to 0.27.0
- (GH-427) Go Dependency: Bump golang.org/x/sys from 0.21.0 to 0.22.0
- (GH-418) Go Runtime: Bump golang from 1.21.11 to 1.21.12 in /dependabot/docker/go
- (GH-433) Update `golang.org/x/mobile` pseudo-version

## [v0.5.2] - 2024-06-07

### Changed

#### Dependency Updates

- (GH-366) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.4 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.5 in /dependabot/docker/builds/x64
- (GH-368) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.5 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.6 in /dependabot/docker/builds/x64
- (GH-387) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.6 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.7 in /dependabot/docker/builds/x64
- (GH-365) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.4 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.5 in /dependabot/docker/builds/x86
- (GH-370) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.5 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.6 in /dependabot/docker/builds/x86
- (GH-388) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.6 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.7 in /dependabot/docker/builds/x86
- (GH-380) Go Dependency: Bump golang.org/x/image from 0.16.0 to 0.17.0
- (GH-377) Go Dependency: Bump golang.org/x/net from 0.25.0 to 0.26.0
- (GH-379) Go Dependency: Bump golang.org/x/sys from 0.20.0 to 0.21.0
- (GH-378) Go Dependency: Bump golang.org/x/text from 0.15.0 to 0.16.0
- (GH-376) Go Runtime: Bump golang from 1.21.10 to 1.21.11 in /dependabot/docker/go
- (GH-386) Update `golang.org/x/mobile` pseudo-version

### Fixed

- (GH-371) Remove inactive maligned linter
- (GH-372) Fix errcheck linting errors

## [v0.5.1] - 2024-05-11

### Changed

#### Dependency Updates

- (GH-335) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.1 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.2 in /dependabot/docker/builds/x64
- (GH-350) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.2 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.3 in /dependabot/docker/builds/x64
- (GH-352) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.3 to go-ci-oldstable-cgo-mingw-w64-buildx64-v0.20.4 in /dependabot/docker/builds/x64
- (GH-334) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.1 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.2 in /dependabot/docker/builds/x86
- (GH-348) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.2 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.3 in /dependabot/docker/builds/x86
- (GH-353) Build image: Bump atc0005/go-ci from go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.3 to go-ci-oldstable-cgo-mingw-w64-buildx86-v0.20.4 in /dependabot/docker/builds/x86
- (GH-318) Go Dependency: Bump fyne.io/fyne/v2 from 2.4.4 to 2.4.5
- (GH-316) Go Dependency: Bump github.com/go-text/render from 0.0.0-20240410160112-301cb7dc78d6 to 0.1.0
- (GH-338) Go Dependency: Bump golang.org/x/image from 0.15.0 to 0.16.0
- (GH-345) Go Dependency: Bump golang.org/x/net from 0.24.0 to 0.25.0
- (GH-336) Go Dependency: Bump golang.org/x/sys from 0.19.0 to 0.20.0
- (GH-337) Go Dependency: Bump golang.org/x/text from 0.14.0 to 0.15.0
- (GH-347) Go Runtime: Bump golang from 1.21.9 to 1.21.10 in /dependabot/docker/go
- (GH-326) Update `github.com/fyne-io/image` pseudo-version
- (GH-356) Update `github.com/go-gl/glfw/v3.3/glfw`
- (GH-358) Update `golang.org/x/mobile` pseudo-version

### Fixed

- (GH-322) Fix `packages` Makefile recipe
- (GH-328) Enable i386 arch for Makefile `all` builds
- (GH-330) Rework support for i386 arch for Makefile `all`

## [v0.5.0] - 2024-04-10

### Added

- (GH-268) Add .gitattributes file to ignore merge conflicts
- (GH-308) Evaluate URLs within angle brackets
- (GH-309) Add GUI apps for encoding & decoding input text
- (GH-310) Add tests using testdata input files

### Changed

#### Dependency Updates

- (GH-219) Build image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.4 to go-ci-oldstable-build-v0.16.0 in /dependabot/docker/builds
- (GH-222) Build image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.0 to go-ci-oldstable-build-v0.16.1 in /dependabot/docker/builds
- (GH-224) Build image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.1 to go-ci-oldstable-build-v0.19.0 in /dependabot/docker/builds
- (GH-265) Build image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.19.0 to go-ci-oldstable-build-v0.20.0 in /dependabot/docker/builds
- (GH-287) Build image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.0 to go-ci-oldstable-build-v0.20.1 in /dependabot/docker/builds
- (GH-281) Go Runtime: Bump golang from 1.21.8 to 1.21.9 in /dependabot/docker/go
- (GH-313) Update `github.com/go-text/render` pseudo-version

#### Other Changes

- (GH-269) Disable potential auto EOL conversion for testdata
- (GH-299) Update workflows to specify OS deps for GUI apps
- (GH-304) Disable internal/safelinks logging by default
- (GH-305) Update cmd/usl to use userFeedbackOut output sink
- (GH-307) Minor refactoring and cleanup

### Fixed

- (GH-264) Update Dependabot build Dockerfile paths
- (GH-300) Fix project URL and app name value
- (GH-301) Fix GetURLPatternsUsingIndex URL end pos matching
- (GH-311) Fix Dockerfile paths

## [v0.4.0] - 2024-03-15

### Added

- (GH-216) Add initial support for decoding text streams

## [v0.3.5] - 2024-03-07

### Changed

#### Dependency Updates

- (GH-188) Update Dependabot PR prefixes
- (GH-187) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.15.0 to go-ci-oldstable-build-v0.15.2 in /dependabot/docker/builds
- (GH-189) Update Dependabot PR prefixes (redux)
- (GH-191) Build image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.2 to go-ci-oldstable-build-v0.15.3 in /dependabot/docker/builds
- (GH-204) Update project to Go 1.21 series
- (GH-203) Add todo/release label to "Go Runtime" PRs
- (GH-205) Go Runtime: Bump golang from 1.21.6 to 1.21.8 in /dependabot/docker/go
- (GH-202) Build image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.3 to go-ci-oldstable-build-v0.15.4 in /dependabot/docker/builds

#### Other Changes

- (GH-192) Move shared functionality to internal/safelinks

## [v0.3.4] - 2024-02-16

### Changed

#### Dependency Updates

- (GH-175) canary: bump golang from 1.20.13 to 1.20.14 in /dependabot/docker/go
- (GH-162) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.3 to go-ci-oldstable-build-v0.14.4 in /dependabot/docker/builds
- (GH-168) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.4 to go-ci-oldstable-build-v0.14.5 in /dependabot/docker/builds
- (GH-169) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.5 to go-ci-oldstable-build-v0.14.6 in /dependabot/docker/builds
- (GH-177) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.6 to go-ci-oldstable-build-v0.14.9 in /dependabot/docker/builds
- (GH-181) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.9 to go-ci-oldstable-build-v0.15.0 in /dependabot/docker/builds

## [v0.3.3] - 2024-01-19

### Changed

#### Dependency Updates

- (GH-153) canary: bump golang from 1.20.11 to 1.20.12 in /dependabot/docker/go
- (GH-158) canary: bump golang from 1.20.12 to 1.20.13 in /dependabot/docker/go
- (GH-155) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.1 to go-ci-oldstable-build-v0.14.2 in /dependabot/docker/builds
- (GH-161) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.2 to go-ci-oldstable-build-v0.14.3 in /dependabot/docker/builds
- (GH-156) ghaw: bump github/codeql-action from 2 to 3

### Fixed

- (GH-164) Update flag help text

## [v0.3.2] - 2023-11-16

### Changed

#### Dependency Updates

- (GH-143) canary: bump golang from 1.20.10 to 1.20.11 in /dependabot/docker/go
- (GH-135) canary: bump golang from 1.20.8 to 1.20.10 in /dependabot/docker/go
- (GH-140) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.10 to go-ci-oldstable-build-v0.13.12 in /dependabot/docker/builds
- (GH-146) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.12 to go-ci-oldstable-build-v0.14.1 in /dependabot/docker/builds

### Fixed

- (GH-148) Fix goconst linting errors

## [v0.3.1] - 2023-10-10

### Changed

## Dependency Updates

- (GH-116) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-118) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-125) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-128) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.9 to go-ci-oldstable-build-v0.13.10 in /dependabot/docker/builds
- (GH-115) ghaw: bump actions/checkout from 3 to 4

## [v0.3.0] - 2023-08-25

### Added

- (GH-78) Add support for processing multiple input lines

### Changed

- Dependencies
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.13.5` to `go-ci-oldstable-build-v0.13.7`

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

[Unreleased]: https://github.com/atc0005/safelinks/compare/v0.5.16...HEAD
[v0.5.16]: https://github.com/atc0005/safelinks/releases/tag/v0.5.16
[v0.5.15]: https://github.com/atc0005/safelinks/releases/tag/v0.5.15
[v0.5.14]: https://github.com/atc0005/safelinks/releases/tag/v0.5.14
[v0.5.13]: https://github.com/atc0005/safelinks/releases/tag/v0.5.13
[v0.5.12]: https://github.com/atc0005/safelinks/releases/tag/v0.5.12
[v0.5.11]: https://github.com/atc0005/safelinks/releases/tag/v0.5.11
[v0.5.10]: https://github.com/atc0005/safelinks/releases/tag/v0.5.10
[v0.5.9]: https://github.com/atc0005/safelinks/releases/tag/v0.5.9
[v0.5.8]: https://github.com/atc0005/safelinks/releases/tag/v0.5.8
[v0.5.7]: https://github.com/atc0005/safelinks/releases/tag/v0.5.7
[v0.5.6]: https://github.com/atc0005/safelinks/releases/tag/v0.5.6
[v0.5.5]: https://github.com/atc0005/safelinks/releases/tag/v0.5.5
[v0.5.4]: https://github.com/atc0005/safelinks/releases/tag/v0.5.4
[v0.5.3]: https://github.com/atc0005/safelinks/releases/tag/v0.5.3
[v0.5.2]: https://github.com/atc0005/safelinks/releases/tag/v0.5.2
[v0.5.1]: https://github.com/atc0005/safelinks/releases/tag/v0.5.1
[v0.5.0]: https://github.com/atc0005/safelinks/releases/tag/v0.5.0
[v0.4.0]: https://github.com/atc0005/safelinks/releases/tag/v0.4.0
[v0.3.5]: https://github.com/atc0005/safelinks/releases/tag/v0.3.5
[v0.3.4]: https://github.com/atc0005/safelinks/releases/tag/v0.3.4
[v0.3.3]: https://github.com/atc0005/safelinks/releases/tag/v0.3.3
[v0.3.2]: https://github.com/atc0005/safelinks/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/safelinks/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/safelinks/releases/tag/v0.3.0
[v0.2.4]: https://github.com/atc0005/safelinks/releases/tag/v0.2.4
[v0.2.3]: https://github.com/atc0005/safelinks/releases/tag/v0.2.3
[v0.2.2]: https://github.com/atc0005/safelinks/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/safelinks/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/safelinks/releases/tag/v0.2.0
[v0.1.1]: https://github.com/atc0005/safelinks/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/safelinks/releases/tag/v0.1.0
