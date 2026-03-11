URL Functions
*************

----

.. _func-urlParse:

urlParse
========

.. code-block:: go

   urlParse(url string) map[string]any

The ``urlParse`` function parses a string as a URL and produces a dict with the URL parts.

**Example:**

.. code-block:: html

   {{ urlParse "http://admin:secret@server.com:8080/api?list=false#anchor" | toPrettyJson }}

The above returns

.. code-block:: text

   {
     "fragment": "anchor",
     "host": "server.com:8080",
     "hostname": "server.com",
     "opaque": "",
     "path": "/api",
     "query": "list=false",
     "scheme": "http",
     "userinfo": "admin:secret"
   }

----

.. _func-urlJoin:

urlJoin
=======

.. code-block:: go

   urlJoin(parts  map[string]any) string

The ``urlJoin`` function joins map (produced by :ref:`func-urlParse`) to produce URL string.

**Example:**

.. code-block:: html

   {{ $url := urlParse "http://admin:secret@server.com:8080/api?list=false#anchor" }}
   {{ set $url "scheme" "ftp" | urlJoin }}

The above returns

.. code-block:: text

   ftp://admin:secret@server.com:8080/api?list=false#anchor

