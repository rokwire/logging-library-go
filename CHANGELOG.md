# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased
### Added
- Expose internal errors [#23](https://github.com/rokwire/logging-library-go/issues/23)

## [2.1.0] - 2022-12-01
### Added
- Regex in HTTPRequestProperties [#20](https://github.com/rokwire/logging-library-go/issues/20)

## [2.0.0] - 2022-11-16
### Added
- Return trace ID in error responses [#13](https://github.com/rokwire/logging-library-go/issues/13)
- Add OpenShift health check request properties builder for log suppression [#14](https://github.com/rokwire/logging-library-go/issues/14)

### Changed
- BREAKING: Restructure logs package [#16](https://github.com/rokwire/logging-library-go/issues/16)

## [1.0.3] - 2021-12-02
### Fixed
- Error JSON encoding incorrect [#9](https://github.com/rokwire/logging-library-go/issues/9)

## [1.0.2] - 2021-11-04
### Fixed
- Error HTTP responses have wrong content type header [#6](https://github.com/rokwire/logging-library-go/issues/6)
- Fix bug on log levels [#5](https://github.com/rokwire/logging-library-go/issues/5)

## [1.0.1] - 2021-11-04
### Added
- Add status to HTTP error responses [#2](https://github.com/rokwire/logging-library-go/issues/2)

## [1.0.0] - 2021-10-01
### Added
- Initial release

[Unreleased]: https://github.com/rokwire/logging-library-go/compare/v2.1.0....HEAD
[2.1.0]: https://github.com/rokwire/logging-library-go/compare/v2.0.0...v2.1.0
[2.0.0]: https://github.com/rokwire/logging-library-go/compare/v1.0.3...v2.0.0
[1.0.3]: https://github.com/rokwire/logging-library-go/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/rokwire/logging-library-go/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/rokwire/logging-library-go/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/rokwire/logging-library-go/tree/v1.0.0