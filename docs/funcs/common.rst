Common Functions
****************

gostub provides tools for setting default values for templates.

----

.. _func-default:

default
=======

.. code-block:: go

   default(defaultValue any, value any) any

If ``value`` evaluates to a non-empty value, it will be used. But if it is empty, ``defaultValue`` will be returned instead.

The definition of “empty” depends on the type:

* Numeric: ``0``
* String: ``""``
* Lists: ``[]``
* Dicts: ``{}``
* Boolean: ``false``
* nil (aka null)

For structs, there is no definition of empty, so a struct will never return the default.

**Example:**

.. code-block:: html

   {{ $name := .Request.URL.Query.Get "name" | default "Guest" }}
   Hello, {{ $name }}!

The above returns

.. code-block:: text

    Hello, Guest!

----

.. _func-empty:

empty
=====

.. code-block:: go

   empty(value any) bool

The ``empty`` function returns ``true`` if the given value is considered empty and false otherwise.
The ``empty`` values are listed in the :ref:`func-default` section.

**Example:**

.. code-block:: html

   {{ $query := .Request.URL.Query }}
   {{ $id := index $query "id" | first }}

   {{ if empty $id }}
       Error: ID is required
   {{ else }}
       Processing ID: {{ $id }}
   {{ end }}

The above returns

.. code-block:: text

    Error: ID is required

----

.. _func-coalesce:

coalesce
========

.. code-block:: go

   coalesce(values ...any) any

The ``coalesce`` function takes multiple values and returns the first non-empty one.

**Example:**

.. code-block:: html

   First non-empty {{ coalesce 0 "" "value" }}

The above returns

.. code-block:: text

    First non-empty value

----

.. _func-all:

all
===

.. code-block:: go

   all(values ...any) bool

The ``all`` function takes a list of values and returns ``true`` if all values are non-empty.

**Example:**

.. code-block:: html

   all non-empty {{ all "" "some-value" }}

The above returns

.. code-block:: text

    all non-empty false

----

.. _func-any:

any
===

.. code-block:: go

   any(values ...any) bool

The ``any``  function takes a list of values and returns ``true`` if any value is non-empty.

**Example:**

.. code-block:: html

   has non-empty {{ any "" "value" }}

The above returns

.. code-block:: text

    has non-empty true

----

.. _func-fromJson:
.. _func-mustFromJson:

fromJson, mustFromJson
======================

.. code-block:: go

   fromJson(value string) interface{}
   mustFromJson(value string) interface{}, error

The ``fromJson`` function decodes a JSON document into a structure.
If the input cannot be decoded as JSON, the function will return an empty string.
``mustFromJson`` will return an error in case the JSON is invalid.

**Example:**

.. code-block:: html

   The value of 'foo' is {{ (fromJson "{\"foo\": 333}").foo }}

The above returns

.. code-block:: text

    The value of 'foo' is 333

----

.. _func-toJson:
.. _func-mustToJson:

toJson, mustToJson
==================

.. code-block:: go

   toJson(value any) string
   mustToJson(value any) string, error

The ``toJson`` function encodes an item into a JSON string.
If the item cannot be converted to JSON the function will return an empty string.
``mustToJson`` will return an error in case the item cannot be encoded in JSON.

**Example:**

.. code-block:: html

   JSON Request {{ toJson .Request.URL }}

The above returns

.. code-block:: text

    JSON Request {"Scheme":"","Opaque":"","User":null,"Host":"","Path":"/","RawPath":"","OmitHost":false,"ForceQuery":false,"RawQuery":"","Fragment":"","RawFragment":""}

----

.. _func-toPrettyJson:
.. _func-mustToPrettyJson:

toPrettyJson, mustToPrettyJson
==============================

.. code-block:: go

   toPrettyJson(value any) string
   mustToPrettyJson(value any) string, error

The ``toPrettyJson`` function encodes an item into a pretty (indented) JSON string.
If the input cannot be decoded as JSON, the function will return an empty string.
``mustToPrettyJson`` will return an error in case the JSON is invalid.

**Example:**

.. code-block:: html

   JSON Request {{ toPrettyJson .Request.URL }}

The above returns

.. code-block:: text

   JSON Request {
      "Scheme": "",
      "Opaque": "",
      "User": null,
      "Host": "",
      "Path": "/",
      "RawPath": "",
      "OmitHost": false,
      "ForceQuery": false,
      "RawQuery": "days=1",
      "Fragment": "",
      "RawFragment": ""
   }

----

.. _func-toRawJson:
.. _func-mustToRawJson:

toRawJson, mustToRawJson
========================

.. code-block:: go

   toRawJson(value any) string
   mustToRawJson(value any) string, error

The ``toRawJson`` function encodes an item into JSON string with HTML characters unescaped.
If the input cannot be decoded as JSON, the function will return an empty string.
``mustToRawJson`` will return an error in case the JSON is invalid.

**Example:**

.. code-block:: html

   JSON Request {{ toRawJson .Request.URL }}

The above returns

.. code-block:: text

   JSON Request {"Scheme":"","Opaque":"","User":null,"Host":"","Path":"/","RawPath":"","OmitHost":false,"ForceQuery":false,"RawQuery":"days=1","Fragment":"","RawFragment":""}

----

.. _func-ternary:

ternary
=======

.. code-block:: go

   ternary(trueValue any, falseValue any, testValue bool) any

The ``ternary`` function takes two values, and a test value.
If the ``testValue`` is true, the ``trueValue`` will be returned.
If the ``testValue`` is false, the ``falseValue`` will be returned.
This is similar to the *C* ternary operator.

**Example:**

.. code-block:: html

   {
      "debug": {{ ternary "true" "false" (eq (.Request.URL.Query.Get "mode") "debug") }}
   }

The above returns

.. code-block:: text

   curl "http://localhost:8080?mode=debug"
   {
      "debug": true
   }

----
