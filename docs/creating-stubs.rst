##############
Creating stubs
##############

Gostub uses the built-in `Go template package <https://golang.org/pkg/text/template/>`__
(text/template) to generate dynamic responses.
This provides a powerful and flexible tool for creating any response type.

Basic principles
****************

1. Each request URL corresponds to a file on disk.
2. The file must contain a Go template with the required ``define "main"`` block.
3. The result of executing the template is sent to the client.
4. The HTTP response code can be dynamically set using the ``.SetCode`` function.

Required template structure
***************************

Each stub file must **required** to contain a template definition named ``main``:

.. code-block:: go

   {{ define "main" }}
       {{ .SetCode 200 }}
       The content of the response is located here
   {{ end }}

This block is what will be executed when processing the request.
Anything outside this block is ignored (but can be used to define additional templates).

Available objects and functions
*******************************

* ``.Request`` — HTTP Request (``*http.Request``)
* ``.SetCode`` — function for setting the HTTP response status

Working with query parameters
-----------------------------

Example of obtaining query parameters:

.. code-block:: go

   {{ define "main" }}
       {{- $query := .Request.URL.Query -}}
       {{- $id := index $query "id" | first -}}
       {{- $debug := $query.Get "debug" -}}

       User ID: {{ $id }}, Debug: {{ $debug }}
   {{ end }}

Conditional constructions
-------------------------

.. code-block:: go

   {{ define "main" }}
       {{- $query := .Request.URL.Query -}}

       {{- if eq ($query.Get "debug") "true" -}}
           {{ .SetCode 400 }}
           {
             "error": "Debug mode activated",
             "status": "bad_request"
           }
       {{- else if ne ($query.Get "id") "" -}}
           {{ .SetCode 200 }}
           {
             "user_id": {{ index $query "id" | first }},
             "status": "found"
           }
       {{- else -}}
           {{ .SetCode 404 }}
           {
             "error": "User not found"
           }
       {{- end -}}
   {{ end }}

Loops (iterating through collections)
-------------------------------------

.. code-block:: go

   {{ define "main" }}
       {{ .SetCode 200 }}
       {
         "users": [
           {{- $query := .Request.URL.Query -}}
           {{- $ids := index $query "id" -}}
           {{- range $index, $id := $ids -}}
             {{- if $index -}}, {{- end -}}
             {
               "id": {{ $id }},
               "position": {{ $index }}
             }
           {{- end -}}
         ]
       }
   {{ end }}

Pipelines
---------

Go templates support the concept of "pipes" similar to Unix:

.. code-block:: go

   {{ define "main" }}
       {{ .SetCode 200 }}
       {{- $query := .Request.URL.Query -}}
       {{- $id := index $query "id" | first | default "0" -}}

       Processed ID: {{ $id | printf "ID-%s" }}
   {{ end }}

Here the value is passed through a chain of functions:
1. ``index $query "id"`` — get all values of the id parameter
2. ``first`` — get the first element
3. ``default "0"`` — if the value is empty, substitute "0"

Auxiliary templates
*******************

Files beginning with the underscore character ``_`` are considered auxiliary files and are not processed as standalone stubs.
They are used to define common templates that can be reused.

Example directory structure:

.. code-block:: text

   stubs/
   └── users/
       ├── _helpers
       ├── profile
       └── settings

Content ``_helpers``:

.. code-block:: go

   {{ define "format.user" -}}
       "user_{{ . }}"
   {{- end }}

   {{ define "error.response" -}}
       {
         "error": "{{ .message }}",
         "code": {{ .code }}
       }
   {{- end }}

   {{ define "isAdmin" -}}
       {{- eq . "admin" -}}
   {{- end }}

Now the main ``profile`` template can use these definitions:

.. code-block:: go

   {{ define "main" }}
       {{- $role := "admin" -}}

       {{- if template "isAdmin" $role -}}
           {{ .SetCode 200 }}
           {
             "user": {{ template "format.user" "admin_user" }}
           }
       {{- else -}}
           {{ .SetCode 403 }}
           {{ template "error.response" dict "message" "Forbidden" "code" 403 }}
       {{- end -}}
   {{ end }}

Variables
*********

Templates support the creation and use of variables:

.. code-block:: go

   {{ define "main" }}
       {{- $query := .Request.URL.Query -}}
       {{- $debug := $query.Get "debug" -}}
       {{- $id := index $query "id" | first -}}

       {{- $user := dict "id" $id "role" "user" "active" true -}}

       {{ .SetCode 200 }}
       {
         "debug_mode": {{ $debug }},
         "user_info": {
           "id": {{ $user.id }},
           "role": {{ $user.role }},
           "active": {{ $user.active }}
         }
       }
   {{ end }}

Comments
********

.. code-block:: go

   {{ define "main" }}
       {{- /* This is a comment, it will not be included in the reply */ -}}

       {{ .SetCode 200 }}
       Hello World  {{- /* Another comment */ -}}
   {{ end }}

Spaces and formatting
*********************

Go templates have an important feature: the ``-`` characters in curly braces
control whitespace removal:

* ``{{-`` - removes spaces from the left
* ``-}}`` - removes spaces from the right

This helps create clean JSON formatting:

.. code-block:: go

   {{ define "main" }}
       {{ .SetCode 200 -}}
       {
         "data": [
           {{- range $i, $v := slice "a" "b" "c" -}}
             {{- if $i -}}, {{- end -}}
             "{{ $v }}"
           {{- end -}}
         ]
       }
   {{ end }}

The result will be formatted without extra spaces and line breaks.

Full template example
*********************

.. code-block:: go

   {{ define "main" }}
       {{- $query := .Request.URL.Query -}}
       {{- $id := index $query "id" | first -}}
       {{- $format := $query.Get "format" | default "json" -}}

       {{- if not $id -}}
           {{ .SetCode 400 }}
           {{ template "error" "Missing id parameter" }}
       {{- else -}}
           {{ .SetCode 200 }}
           {{ template "response" dict "id" $id "format" $format }}
       {{- end -}}
   {{ end }}

   {{ define "error" }}
       {
         "error": "{{ . }}",
         "status": "failed"
       }
   {{ end }}

   {{ define "response" }}
       {{- if eq .format "json" -}}
           {
             "user_id": {{ .id }},
             "format": "json",
             "timestamp": "{{ now }}"
           }
       {{- else -}}
           User ID: {{ .id }}
       {{- end -}}
   {{ end }}

Notes
*****

* Always use ``{{ define "main" }}`` — without it, the template will not execute.
* Auxiliary files (with ``_``) are automatically available in all templates in the directory.
* For complex logic, move repeating parts into auxiliary templates.
* Be careful with whitespace when generating JSON; use ``{{-`` and ``-}}``.
* Default values can be implemented using the ``default`` function.
