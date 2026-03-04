String Slice Functions
**********************

These functions operate on or generate slices of strings. In Go, a slice is a growable array. In Gostub, it’s a special case of a ``list``.

----

.. _func-join:

join
====

.. code-block:: go

   join(sep string, value ...string) string

The ``join`` function joins a list of strings into a single string, with the given separator.

**Example:**

.. code-block:: html

   {{ list "hello" "world" | join "_" }}

The above returns

.. code-block:: text

   hello_world

----

.. _func-splitList:

splitList
=========

.. code-block:: go

   splitList(sep string, value string) []string

The ``splitList`` function splits a string into a list of strings.

**Example:**

.. code-block:: html

   {{ splitList "$" "foo$bar$baz" }}

The above returns

.. code-block:: text

   [foo bar baz]

----

.. _func-split:

split
=====

.. code-block:: go

   split(sep string, value string) map[string]string

The ``split`` function splits a string into a dict.
The keys of the dict are zero-based position numbers from the source string, prefixed with an underscore (e.g., _0, _1, _2) and the values are actual parts of source string.
It is designed to make it easy to use template dot notation for accessing members.

**Example:**

.. code-block:: html

   {{ $a := split "$" "foo$bar$baz" }}
   {{ $a }}
   The first value is {{ $a._0 }}

The above returns

.. code-block:: text

   map[_0:foo _1:bar _2:baz]
   The first value is foo

----

.. _func-splitn:

splitn
======

.. code-block:: go

   splitn(sep string, count int, value string) map[string]string

The ``splitn`` function splits a string into a dict.
The keys of the dict are zero-based position numbers from the source string, prefixed with an underscore (e.g., _0, _1, _2) and the values are actual parts of source string.
It is designed to make it easy to use template dot notation for accessing members.

**Example:**

.. code-block:: html

   {{ $a := splitn "$" 2 "foo$bar$baz" }}
   {{ $a }}
   The first value is {{ $a._0 }}

The above returns

.. code-block:: text

   map[_0:foo _1:bar$baz]
   The first value is foo

----

.. _func-sortAlpha:

sortAlpha
=========

.. code-block:: go

   sortAlpha(value []string) []string

The ``sortAlpha`` function sorts a list of strings into alphabetical (lexicographical) order.
It does not sort in place, but returns a sorted copy of the list, in keeping with the immutability of lists.

**Example:**

.. code-block:: html

   {{ sortAlpha (list "d" "a" "v" "x" "b") }}

The above returns

.. code-block:: text

   [a b d v x]


