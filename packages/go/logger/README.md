# BeatFlow Shared Logger Package

This package provides shared structured logging utilities for BeatFlow Go services.

It uses Zap to provide a consistent logging standard across services.

## Purpose

Each BeatFlow service should use the same logging style.

This package creates reusable Zap loggers with environment-aware configuration:

* development logger for local development
* production logger for structured JSON logs
* service name field support
* request ID and correlation ID field pattern

The package returns a `*zap.Logger` directly so services can keep using Zap's production-ready API without unnecessary custom abstraction.

## Logger Options

```go
type Options struct {
	Environment string
	ServiceName string
	Level       string
}
```

### `Environment`

Controls which logger configuration is used.

Supported development values:

```txt
development
dev
local
```

These use Zap's development logger configuration.

Any other value, including an empty value, uses the production logger configuration.

This means production-style JSON logging is the safe default.

### `ServiceName`

Adds a persistent `service_name` field to every log produced by the logger.

Example:

```json
{
  "level": "info",
  "service_name": "auth-service",
  "message": "server started"
}
```

### `Level`

Controls the minimum log level.

Supported values:

```txt
debug
info
warn
warning
error
```

Unknown or empty values fall back to `info`.

## Development Logger

Development logger is intended for local development.

Example:

```go
log, err := logger.New(logger.Options{
	Environment: "development",
	ServiceName: "auth-service",
	Level:       "debug",
})
```

Development logs are more human-readable and useful while building locally.

## Production Logger

Production logger is intended for production-like environments.

Example:

```go
log, err := logger.New(logger.Options{
	Environment: "production",
	ServiceName: "auth-service",
	Level:       "info",
})
```

Production logs use JSON encoding and write to standard output.

The production encoder uses these field names:

```txt
timestamp
level
logger
caller
message
stacktrace
```

This makes logs easier to collect and query later with log aggregation tools.

## Example Usage

```go
package main

import (
	"log"

	"github.com/dogukanttopcuoglu/beatflow/packages/go/logger"
	"go.uber.org/zap"
)

func main() {
	appLogger, err := logger.New(logger.Options{
		Environment: "development",
		ServiceName: "auth-service",
		Level:       "debug",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer appLogger.Sync()

	appLogger.Info("service started",
		zap.String("address", ":8080"),
	)
}
```

## Request ID and Correlation ID Pattern

The shared logger package creates the base service logger.

Request-specific fields should be added by HTTP middleware.

The logger package does not generate request IDs or correlation IDs by itself.

Example request-scoped logger pattern:

```go
requestLogger := baseLogger.With(
	zap.String("request_id", requestID),
	zap.String("correlation_id", correlationID),
)
```

Then handlers can use the request-scoped logger:

```go
requestLogger.Info("request handled",
	zap.String("method", "GET"),
	zap.String("path", "/healthz"),
)
```

## Why Request Fields Are Added Outside This Package

`service_name` is stable for the whole service, so it is added when the base logger is created.

`request_id` and `correlation_id` are different for every HTTP request, so they should be added inside middleware.

Typical request flow:

```txt
HTTP request enters service
middleware reads X-Request-ID and X-Correlation-ID
middleware generates missing IDs if needed
middleware creates request-scoped logger with log.With(...)
handler uses request-scoped logger
```

## Design Rule

This package should only provide shared structured logging behavior.

It should not contain HTTP middleware, request ID generation, tracing logic, or business-specific logging rules.

Those concerns should live inside the owning service or a dedicated shared middleware package later.

## Current Scope

This initial version supports:

* reusable Zap logger creation
* development logger configuration
* production JSON logger configuration
* log level parsing
* safe default log level
* persistent `service_name` field
* request/correlation field pattern through `log.With(...)`

Future improvements may include:

* shared HTTP logging middleware
* context-aware logger extraction
* OpenTelemetry trace ID field support
* log sampling configuration
* standardized error logging helpers
