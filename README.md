# utils
Provides useful Go utility packages.

#### Error

The `error` package provides error reporting with stack traces and other
captured context (at the point of origin).

Defines and implements interfaces:
- Traced - provides origin stack trace.
- Wrapped - provides Unwrap().
- WithContext - provides additional context properties.
- Snapshot - provides _error Traced, Wrapped, WithContext_.

#### File-Backed

the `filebacked` package provides collections that are backed by the
filesystem rather than in-memory.

#### Logger

The `logr` package provides a rich _logrus_ based Logger that is _Snapshot_
aware when logging at Error().
