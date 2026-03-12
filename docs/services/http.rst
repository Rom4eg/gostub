************
Http Service
************

The HTTP service listens for incoming HTTP requests and serves stub responses from the file system.

Configuration
=============

When defining an HTTP service in the configuration file, the following parameters are available under the ``options`` field.

.. list-table::
   :header-rows: 1
   :widths: 15 10 30 45

   * - Parameter
     - Required
     - Description
     - Example
   * - ``host``
     - Y
     - Host to listen for incoming connections
     - ``localhost``, ``0.0.0.0``, ``127.0.0.1``
   * - ``port``
     - Y
     - Port for listening for HTTP requests
     - ``8080``, ``80``, ``3000``
   * - ``root``
     - Y
     - The base directory where template files are stored
     - ``/tmp/stubs``, ``./stubs``, ``/var/gostub``

How the path to the stubs is formed
***********************************

When an HTTP request is received, the service constructs a file system path using the following pattern:

.. code-block:: text

   [root]/[name]/[request_path]

* ``root`` — value from configuration
* ``name`` — is the service ID
* ``request_path`` — is the path from the request URL (query is ignored)

Example
-------

.. code-block:: yaml

   services:
     - name: users
       type: http
       options:
         host: localhost
         port: 8080
         root: /var/stubs

Request:

.. code-block:: text

   GET http://localhost:8080/api/v1/profile?id=123

The service will search for:

.. code-block:: text

   /var/stubs/users/api/v1/profile

Functions
=========

Functions available only for ``http`` service

Note: certain headers such as Content-Length can be added automatically and do not appear in the Headers before response is started.

----

.. _func-SetCode:

SetCode
*******

.. code-block:: go

   SetCode(status int)

The ``SetCode`` function sets the HTTP response status code.

**Example:**

.. code-block:: html

   {{- define "main" -}}
       {{ .SetCode 201 }}
       {
           "status": "success"
       }
   {{- end -}}

The above returns

.. code-block:: text

   $ curl -i http://localhost:8081/api

   HTTP/1.1 201 Created
   Date: Thu, 12 Mar 2026 11:27:21 GMT
   Content-Length: 40
   Content-Type: text/plain; charset=utf-8


   {
       "status": "success"
   }

----

.. _func-Code:

Code
****

.. code-block:: go

   Code() int

The ``Code`` function returns the current response code

**Example:**

.. code-block:: html

   {{ .SetCode 201 }}
   The code is {{ .Code }}

The above returns

.. code-block:: text

   The code is 201

----

.. _func-AddHeader:

AddHeader
*********

.. code-block:: go

   AddHeader(key string, value string)

The ``AddHeader`` function adds the ``key``, value pair to the header.
It appends to any existing values associated with ``key``.

The ``key`` is case insensitive

**Example:**

.. code-block:: html

   {{ .AddHeader "X-Test" "Hello world!!!" }}
   Some content

The above returns

.. code-block:: text

   $ curl -i http://localhost:8081/api

   HTTP/1.1 200 OK
   X-Test: Hello world!!!
   Date: Thu, 12 Mar 2026 12:05:23 GMT
   Content-Length: 16
   Content-Type: text/plain; charset=utf-8


   Some content

----

.. _func-SetHeader:

SetHeader
*********

.. code-block:: go

   SetHeader(key string, value string)

The ``SetHeader`` function sets the header entries associated with ``key`` to the single element value.
It replaces any existing values associated with ``key``.

The ``key`` is case insensitive.

**Example:**

.. code-block:: html

   {{ .AddHeader "X-Test" "Hello world!!!" }}
   {{ .SetHeader "X-Test" "Hello gostub!!!" }}
   Some content

The above returns

.. code-block:: text

   $ curl -i http://localhost:8081/api

   HTTP/1.1 200 OK
   X-Test: Hello gostub!!!
   Date: Thu, 12 Mar 2026 12:13:16 GMT
   Content-Length: 20
   Content-Type: text/plain; charset=utf-8


   Some content

----

.. _func-Headers:

Headers
*******

.. code-block:: go

   Headers() http.Header

The ``Headers`` function returns response headers.

Check the official docs for more information about `http.Header <https://pkg.go.dev/net/http#Header>`__

**Example:**

.. code-block:: html

   {{ .AddHeader "X-Test" "Hello world!!!" }}
   {{ .Headers | toPrettyJson }}

The above returns

.. code-block:: json

   {
       "X-Test": [
           "Hello world!!!"
       ]
   }

----

.. _func-DeleteHeader:

DeleteHeader
************

.. code-block:: go

   DeleteHeader(key string)

The ``DeleteHeader`` function deletes the values associated with ``key``.

The ``key`` is case insensitive.

**Example:**

.. code-block:: html

   {{ .AddHeader "X-Test" "Hello world!!!" }}
   {{ .DeleteHeader "X-Test" }}
   {{ .Headers | toPrettyJson }}

The above returns

.. code-block:: text

   $ curl -i http://localhost:8081/api

   HTTP/1.1 200 OK
   Date: Thu, 12 Mar 2026 12:28:27 GMT
   Content-Length: 4
   Content-Type: text/plain; charset=utf-8


   {}
