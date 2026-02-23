# gostub

[![Go Version](https://img.shields.io/github/go-mod/go-version/rom4eg/gostub)](https://golang.org/)
[![License](https://img.shields.io/github/license/rom4eg/gostub)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/rom4eg/gostub/go.yml)](https://github.com/rom4eg/gostub/actions)

gostub - This is a lightweight service for simulating responses to various APIs. During development and debugging, stubs for external services are quickly created when the real service is unavailable or difficult to use.

## 🌟 Possibilities
 * Powerful Go templates - use the full power of the templating language to generate dynamic responses
 * Dynamic HTTP statuses - set response codes directly from the template
 * Flexible configuration - support multiple services in a single configuration file
 * Easy deployment - a single binary with no external dependencies
 * Access to environment variables - retrieve system variable values ​​in templates

## 🛠 Use cases

  * Demos and presentations when a predictable API response is needed
  * Load testing: fast responses without accessing real services
  * Error handling testing: simulating various HTTP statuses
  * Frontend development when the backend isn't ready yet

## 📦 Installation

### From source
```bash
git clone https://github.com/rom4eg/gostub.git
cd gostub
make all
```

### Docker
```
docker pull rom4eg196/gostub
docker run --rm -p 8080:8080 rom4eg196/gostub
```

## ⚙️ Configuration

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

#### Note: each service will append it's name to the root directory parameter.
In the example above the service will look for stubs in /tmp/stubs/default

## 📝 Creating Stubs
Stubs are Go templates that are rendered when the corresponding URL is accessed.
The request URL is translated into a filesystem path relative to the root directory and service name.

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

### Available functions in the template


| Function       | Description                                | Example                       |
|----------------|--------------------------------------------|-------------------------------|
| .SetCode(code) | Sets the HTTP response status              | {{ .SetCode 404 }}            |
| env("VAR")     | Gets the value of an environment variable. | {{ env "HOME" }}              |
| .Request       | Original HTTP request                      | .Request.Method, .Request.URL |


## 🚀 Launch

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

## 🤝 Contribution
PR and ideas are welcome! Create an issue to discuss new features or bugs.

