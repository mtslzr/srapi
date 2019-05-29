# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.2] / 2019.05.28
### Added
- Vendor folder for dependencies
### Changed
- Converted to using Go modules
- Updated Travis CI to include lint/build and module support

## [0.2.1] / 2019.05.27
### Fixed
- Cleaned up long lines of code to be consistent

## [0.2.0] / 2019.05.27
### Added
- Endpoints (Baseball only):
  - GetStandings
  - GetTeams
  - GetYears
- Travis CI (with Slack notifications)
- Unit Tests
### Changed
- Endpoints (Baseball only):
  - Converted GetSport to internal call (removed from router)
- Various minor tweaks to clean up redundant code
### Removed
- Endpoints (Baseball only):
  - GetDummy

## [0.1.0] / 2019.05.26
### Added
- Base project files:
  - Changelog
  - Contributors file
  - gitignore
  - MIT license
  - README
- Endpoints (Baseball only):
  - GetDummy
  - GetSport
- Local SQLite database

[Unreleased]: https://github.com/mtslzr/srapi/compare/v0.2.2...HEAD
[0.2.2]: https://github.com/mtslzr/srapi/compare/v0.2.1...v0.2.2
[0.2.1]: https://github.com/mtslzr/srapi/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/mtslzr/srapi/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/mtslzr/srapi/releases/tag/v0.1.0