# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [unreleased]
### Added
- Add noscript notice
- Add error message if service is unreachable

### Changed
- Update requirements in readme.md

## [v0.0.3] - 09.10.17
### Added
- Add cleanup call on every request and by configured interval
- Add template inheritance for theming
- Add "assets" configuration for custom public assets

### Changed
- Prevent encryption animation to take too long
- Change "themePath" configuration to "templates"

### Removed
- Remove legacy appearance.css
- Remove CLI call for cleanup

### Fixed
- Fix condition race for activeuntil on show message page
- Fix visible scrollbars if hardware mouse detected #1

## [v0.0.2] - 30.09.17
### Added
- Add build directory to .gitignore
- Add start script for development
- Add styleguide with new appearance
- Add autoprefixer
- Add initialization smoothing to the body

### Changed
- Update README.md information
- Refactor index.js
- Change appearance of the application

### Fixed
- Fix query for message selection

## [v0.0.1] - 28.09.17
### Added
- Add build step with pkg
- Add README, LICENCE and CHANGELOG
- Add roadmap
- Add minification plugin for JS
- Reload page if activeUntil time is reached
- Add destroy button to the show page
- Add basic styling and encryption logic
- Initial port from proof of concept prototype
- Add pre-build executables and usage instructions to readme
- Remove binaries from repository
