# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Endpoints (Baseball only):
  - GetStandings
  - GetTeams
  - GetYears
- Travis CI
- Unit Tests
### Changed
- Endpoints (Baseball only):
  - Converted GetSport to internal call (removed from router)
### Removed
- Endpoints (Baseball only):
  - GetDummy

## [0.1.0] - 2019.05.26
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

[Unreleased]: https://github.com/mtslzr/srapi/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/mtslzr/srapi/releases/tag/v0.1.0