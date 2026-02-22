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
    "user_id": {{ index $query "id" }},
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


## 🚀 Launch

```bash
./gostub -config /path/to/config.yaml
```
Example request:

```bash
# Request without debug parameter (should return 200)
curl -i "http://localhost:8080?id=123"

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
curl -i "http://localhost:8080?debug=true"

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

## 🎯 Examples of use

Simple template

```go
{{- define "main" -}}
  {{- .SetCode 200 -}}
{
  "message": "Hello, World!",
  "method": {{ .Request.Method }},
  "path": {{ .Request.URL.Path }}
}
{{- end -}}
```

will produce that response

```text
$ curl -i -X OPTIONS "http://localhost:8080"
HTTP/1.1 200 OK
Date: Sun, 22 Feb 2026 13:20:09 GMT
Content-Length: 66
Content-Type: text/plain; charset=utf-8

{
  "message": "Hello, World!",
  "method": OPTIONS,
  "path": /
}
```


Conditional logic

```go
{{- define "main" -}}
  {{- $method := .Request.Method -}}
  {{- if eq $method "GET" -}}
    {{- .SetCode 200 -}}
    { "data": "GET request processed" }
  {{- else if eq $method "POST" -}}
    {{- .SetCode 201 -}}
    { "data": "POST request processed" }
  {{- else -}}
    {{- .SetCode 405 -}}
    { "error": "Method not allowed" }
  {{- end -}}
{{- end -}}
```

will produce response:

for a GET method
```bash
$ curl -i "http://localhost:8080"

HTTP/1.1 200 OK
Date: Sun, 22 Feb 2026 13:21:45 GMT
Content-Length: 35
Content-Type: text/plain; charset=utf-8

{ "data": "GET request processed" }
```

for a POST method
```bash
$ curl -i -X POST "http://localhost:8080"

HTTP/1.1 201 Created
Date: Sun, 22 Feb 2026 13:23:12 GMT
Content-Length: 36
Content-Type: text/plain; charset=utf-8

{ "data": "POST request processed" }
```

for a PUT method
```bash
$ curl -i -X PUT "http://localhost:8080"

HTTP/1.1 405 Method Not Allowed
Date: Sun, 22 Feb 2026 13:17:52 GMT
Content-Length: 33
Content-Type: text/plain; charset=utf-8

{ "error": "Method not allowed" }
```



## 🛠 Use cases

  * Frontend development when the backend isn't ready yet
  * Error handling testing: simulating various HTTP statuses
  * Load testing: fast responses without accessing real services
  * Demos and presentations when a predictable API response is needed

## 📄 License
GNU/GPLv3 License. More details in the LICENSE file

## 🤝 Contribution
PR and ideas are welcome! Create an issue to discuss new features or bugs.

### gostub — made with ❤️ for developers
