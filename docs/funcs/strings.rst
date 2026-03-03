Strings
*******

.. _func-default:

trim
====

.. code-block:: go

   trim(value string) string

The ``trim`` function removes space from either side of a ``value``

**Example:**

.. code-block:: html

   This {{ trim "       value      " }} has been truncated on both the left and right sides.

The above returns

.. code-block:: text

   This value has been truncated on both the left and right sides.

----

.. _func-trimAll:

trimAll
====

.. code-block:: go

   trimAll(trimChar string, value string) string

The ``trimAll`` function remove given characters from the front or back of a string

**Example:**

.. code-block:: html

   The price is {{ trimAll "$" "$5.00" }}

The above returns

.. code-block:: text

   The price is 5.00

----

.. _func-trimSuffix:

trimSuffix
====

.. code-block:: go

   trimSuffix(trimChar string, value string) string

The ``trimSuffix`` function trim just the suffix from a string

**Example:**

.. code-block:: html

   The time is {{ trimSuffix "+0000" "23:00:00+0000" }}

The above returns

.. code-block:: text

   The time is 23:00:00

----

.. _func-trimPrefix:

trimPrefix
====

.. code-block:: go

   trimPrefix(trimChar string, value string) string

The ``trimPrefix`` function trim just the prefix from a string

**Example:**

.. code-block:: html

   {{ trimPrefix "--" "--name" }}

The above returns

.. code-block:: text

   name

----

.. _func-upper:

upper
====

.. code-block:: go

   upper(value string) string

The ``upper`` function convert the entire string to uppercase:

**Example:**

.. code-block:: html

   {{ upper "this text is uppercase"}}

The above returns

.. code-block:: text

   THIS TEXT IS UPPERCASE

----

.. _func-lower:

lower
=====

.. code-block:: go

   lower(value string) string

The ``lower`` function convert the entire string to lowercase:

**Example:**

.. code-block:: html

   {{ lower "THIS TEXT IS LOWERCASE"}}

The above returns

.. code-block:: text

   this text is lowercase

----
