# gostub

gostub - This is a lightweight service for simulating responses to various APIs. During development and debugging, stubs for external services are quickly created when the real service is unavailable or difficult to use.

## 🌟 Possibilities
 * Flexible configuration - support multiple services in a single configuration file
 * Powerful Go templates - use the full power of the templating language to generate dynamic responses
 * Dynamic HTTP statuses - set response codes directly from the template
 * Access to environment variables - retrieve system variable values ​​in templates
 * Easy deployment - a single binary with no external dependencies

## 📦 Installation

From source
```bash
git clone https://github.com/rom4eg/gostub.git
cd gostub
make all
```

Download the binary release
Visit the releases page and download the appropriate version for your OS.

## ⚙️ Configuration
Create a configuration file (e.g. config.yaml):
```yaml
services:
  - name: default
    type: http
    options:
      host: localhost
      port: 8080
      root: /path/to/your/stubs
```

Configuration options
| Parameter | Description                        | Example          |
|-----------|------------------------------------|------------------|
| name      | Service ID                         | default          |
| type      | Service type (currently only http) | http             |
| host      | Host for listening                 | localhost        |
| port      | Listening port                     | 8080             |
| root      | Root directory with templates      | /home/user/stubs |

## 📝 Creating Stubs
Placeholders are Go templates that are rendered when the corresponding URL is accessed.
The request URL is translated into a filesystem path relative to the root directory.

Example template
Save the file, for example, /path/to/your/stubs/api/user.tmpl:
```go

{{ define "main" }}
  {{- $query := .Request.URL.Query -}}
  {{- if eq ($query.Get "debug") "true" -}}
    {{- .SetCode 400 -}}
    {
      "error": "Bad Request",
      "details": "Invalid parameters",
      "reason": {{ printf "Max process limit %s" (env "GOMAXPROCS") | json }}
    }
  {{- else -}}
    {{- .SetCode 200 -}}
    {
      "status": "success",
      "data": {
        "user_id": {{ index $query "id" | first | default "unknown" | json }},
        "gomaxprocs": {{ env "GOMAXPROCS" }}
      }
    }
  {{- end -}}
{{- end -}}
```

Available functions in the template


| Function       | Description                                | Example                       |
|----------------|--------------------------------------------|-------------------------------|
| .SetCode(code) | Sets the HTTP response status              | {{ .SetCode 404 }}            |
| env("VAR")     | Gets the value of an environment variable. | {{ env "HOME" }}              |
| .Request       | Original HTTP request                      | .Request.Method, .Request.URL |


## 🚀 ЗапуLaunchск

```bash
./gostub -config /path/to/config.yaml
```
Example request:

```bash
# Request without debug parameter (should return 200)
curl -ik "http://localhost:8080/api/user?id=123"

# Request with debug parameter (should return 400)
curl -ik "http://localhost:8080/api/user?debug=true&id=123"
```

## 🎯 Examples of use

Simple answer

```go
{{ define "main" }}
  {{ .SetCode 200 }}
  {
    "message": "Hello, World!",
    "method": {{ .Request.Method | json }},
    "path": {{ .Request.URL.Path | json }}
  }
{{ end }}
```

Conditional logic

```go
{{ define "main" }}
  {{- $method := .Request.Method -}}
  {{- if eq $method "GET" -}}
    {{ .SetCode 200 }}
    { "data": "GET request processed" }
  {{- else if eq $method "POST" -}}
    {{ .SetCode 201 }}
    { "data": "POST request processed" }
  {{- else -}}
    {{ .SetCode 405 }}
    { "error": "Method not allowed" }
  {{- end -}}
{{- end -}}
```

## 🛠 Use cases

  * Frontend development when the backend isn't ready yet
  * Error handling testing: simulating various HTTP statuses
  * Load testing: fast responses without accessing real services
  * Demos and presentations when a predictable API response is needed

## 📄 License
GNU/GPLv3 License. More details in the LICENSE file

## 🤝 Contribution to the project
PR and ideas are welcome! Create an issue to discuss new features or bugs.

# gostub — made with ❤️ for developers
