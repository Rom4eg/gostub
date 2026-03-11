# gostub

[![Go Version](https://img.shields.io/github/go-mod/go-version/rom4eg/gostub)](https://golang.org/)
[![License](https://img.shields.io/github/license/rom4eg/gostub)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/rom4eg/gostub/go.yml)](https://github.com/rom4eg/gostub/actions)

gostub is an HTTP server that simulates API responses. It is used during development and testing to replace external services that are unavailable, unreliable, or impractical to use.

## Features
  * __Go templates:__ Supports Go's templating engine for dynamic response generation.
  * __Configurable HTTP status codes:__ Response status can be set within the template logic.
  * __Multi-service configuration:__ Multiple mock services can be defined in a single YAML file.
  * __Static binary:__ Distributed as a single executable file with no runtime dependencies.
  * __Environment variable access:__ Templates can read and use system environment variables.

## Use cases

  * __Demos and prototyping:__ Provide a consistent and predictable API response for demonstrations or prototypes without relying on live systems.
  * __Load testing:__ Isolate the system under test by mocking backend dependencies, ensuring load tests target only the intended component.
  * __Testing error handling:__ Simulate various HTTP error statuses (4xx, 5xx) to verify application resilience.
  * __Frontend development:__ Frontend teams can develop against a predefined API contract without waiting for the backend implementation to be complete.

## Documentation

Complete documentation is available at [Read the Docs](https://gostub.readthedocs.io/).

## Installation

```bash
go install github.com/Rom4eg/gostub/cmd/gostub@latest
```

### From source
```bash
git clone https://github.com/rom4eg/gostub.git
cd gostub
make build
```

### Docker
```
docker pull rom4eg196/gostub:latest
docker run --rm -p 8080:8080 rom4eg196/gostub
```

## Configuration

Create a configuration file (e.g. config.yaml):
```yaml
services:
  - name: default
    type: http
    options:
      host: localhost
      port: 8080
      root: /tmp/stubs
```

### Configuration
| Parameter | Description                        | Example          |
|-----------|------------------------------------|------------------|
| name      | Service ID                         | default          |
| type      | Service type (currently only http) | http             |
| host      | Host for listening                 | localhost        |
| port      | Listening port                     | 8080             |
| root      | Root directory with templates      | /tmp/stubs |

#### Note: each service appends its name to the root directory path.
For example, with root: /tmp/stubs and name: default, the service looks for stub files in /tmp/stubs/default.

## Creating Stubs
Stub files are written as Go templates.
When a request is received, gostub maps the URL path to a file on disk.
The mapping format is: [root]/[service_name]/[URL_path].

For example, a request to http://localhost:8080/api/user with the configuration above will render the file /tmp/stubs/default/api/user.

### Example template
Save the file, for example, /tmp/stubs/default/api/user:
```go

{{ define "main" }}
  {{- $query := .Request.URL.Query -}}
  {{- if eq ($query.Get "debug") "true" -}}
    {{- .SetCode 400 }}
{
  "error": "Bad Request",
  "details": "Invalid parameters",
  "reason": {{ printf "Max process limit %s" (env "GOMAXPROCS") }}
}
  {{- else -}}
    {{- .SetCode 200 -}}
{
  "status": "success",
  "data": {
    "user_id": {{ index $query "id" | first }},
    "gomaxprocs": {{ env "GOMAXPROCS" }}
  }
}
  {{- end -}}
{{- end -}}
```

## Usage

```bash
./gostub -config /path/to/config.yaml
```
Example request:

```bash
# Request without debug parameter (should return 200)
curl -i "http://localhost:8080/api/user?id=123"

HTTP/1.1 200 OK
Date: Sun, 22 Feb 2026 13:32:05 GMT
Content-Length: 85
Content-Type: text/plain; charset=utf-8

{
  "status": "success",
  "data": {
    "user_id": [123],
    "gomaxprocs": 10
  }
}

# Request with debug parameter (will return 400)
curl -i "http://localhost:8080/api/user?debug=true&id=123"

HTTP/1.1 400 Bad Request
Date: Sun, 22 Feb 2026 13:30:57 GMT
Content-Length: 98
Content-Type: text/plain; charset=utf-8

{
  "error": "Bad Request",
  "details": "Invalid parameters",
  "reason": Max process limit 10
}
```

## Contribution
Pull requests and feature requests are welcome. Please open an issue to discuss significant changes before submitting a pull request.

