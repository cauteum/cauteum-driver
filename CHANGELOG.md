# Changelog

## [Unreleased]

## [v0.1.0-alpha.2] - 2026-10-07

### Added

- Support typed Docker and Podman driver configuration, runtime defaults, and provider-specific resource settings.
- Validate driver payloads, mounts, reserved paths, GPU configuration, and SDK-facing container options.
- Configure bounded health, create, and cleanup operations.

### Changed

- Use the Moby API/client modules instead of the deprecated monolithic Docker SDK.
- Update Testcontainers to v0.44.0, Moby API/client to v1.56.1/v0.6.1, and go-archive to v0.3.3.
- Update the core dependency to v0.1.0-alpha.2 and distribute this module under Apache-2.0.

### Fixed

- Keep archive extraction, mount handling, and container cleanup fail-closed on invalid paths or partial failures.
- Update `golang.org/x/crypto` to v0.57.0 to remove reachable SSH vulnerabilities in the Testcontainers test path.

## [v0.0.2-alpha.1] - 2026-09-28

### Security

- `CopyFrom` extracts archives under `os.OpenRoot` with path sanitization and symlink rejection.
