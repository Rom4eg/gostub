Type Conversion Functions
*************************

The following type conversion functions are provided.

----

.. _func-atoi:

atoi
====

.. code-block:: go

   atoi(val string) int

The ``atoi`` function converts a string to an integer.

**Example:**

.. code-block:: html

   {{ atoi "5" }}

The above returns

.. code-block:: text

   5

----

.. _func-float64:

float64
=======

.. code-block:: go

   float64(val any) float64

The ``float64`` function converts a ``val`` to an float64.

**Example:**

.. code-block:: html

   {{ float64 "5" }}

The above returns

.. code-block:: text

   5

----

.. _func-int:

int
===

.. code-block:: go

   int(val any) int

The ``int`` function converts a ``val`` to an integer.

**Example:**

.. code-block:: html

   {{ int "5" }}

The above returns

.. code-block:: text

   5

----

.. _func-int64:

int64
=====

.. code-block:: go

   int64(val any) int64

The ``int64`` function converts a ``val`` to an int64 type.

**Example:**

.. code-block:: html

   {{ int64 "5" }}

The above returns

.. code-block:: text

   5

----

.. _func-toDecimal:

toDecimal
=========

.. code-block:: go

   toDecimal(val any) int64

The ``toDecimal`` function converts a unix octal to a int64.

**Example:**

.. code-block:: html

   {{ toDecimal "0777" }}

The above returns

.. code-block:: text

   511

----

.. _func-toString:

toString
========

.. code-block:: go

   toString(val any) string

The ``toString`` function converts a ``val`` to string.

**Example:**

.. code-block:: html

   {{ toString 123 }}

The above returns

.. code-block:: text

   123

----

.. _func-toStrings:

toStrings
=========

.. code-block:: go

   toStrings(val []any) []string

The ``toStrings`` function converts a list, slice, or array to a list of strings.

**Example:**

.. code-block:: html

   {{ list 1 2 3 | toStrings }}

The above returns

.. code-block:: text

   [1 2 3]

