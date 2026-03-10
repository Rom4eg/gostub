Reflection Functions
********************

These help advanced template developers understand the underlying Go type information for a particular value.

----

.. _func-kindOf:

kindOf
======

.. code-block:: go

   kindOf(val any) string

The ``kindOf`` function returns the kind of an object.

**Example:**

.. code-block:: html

   {{ kindOf "hello" }}

The above returns

.. code-block:: text

   string

----

.. _func-kindIs:

kindIs
======

.. code-block:: go

   kindIs(kind string, value any) bool

The ``kindIs`` function verifies that a ``value`` is a particular ``kind``.

**Example:**

.. code-block:: html

   {{ kindIs "int" 123 }}

The above returns

.. code-block:: text

   true

----

.. _func-typeOf:

typeOf
======

.. code-block:: go

   typeOf(val any) string

The ``typeOf`` function returns the underlying type of a value.

**Example:**

.. code-block:: html

   {{ $myInt := semver "1.2.3" }}
   {{ typeOf $myInt }}

The above returns

.. code-block:: text

   *semver.Version

----

.. _func-typeIs:

typeIs
======

.. code-block:: go

   typeIs(type string, value any) bool

The ``typeIs`` function verifies that a ``value`` is an exact ``type``.

**Example:**

.. code-block:: html

   {{ $myInt := semver "1.2.3" }}
   {{ typeIs "*semver.Version" $myInt }}

The above returns

.. code-block:: text

   true

----

.. _func-typeIsLike:

typeIsLike
==========

.. code-block:: go

   typeIsLike(type string, value any) bool

The ``typeIsLike`` function works as :ref:`func-typeIs`, except that it also dereferences pointers.

**Example:**

.. code-block:: html

   {{ $myInt := semver "1.2.3" }}
   {{ typeIsLike "semver.Version" $myInt }}

The above returns

.. code-block:: text

   true

----

.. _func-deepEqual:

deepEqual
=========

.. code-block:: go

   deepEqual(x any, y any) bool

The ``deepEqual`` function reports whether x and y are `deeply equal <https://pkg.go.dev/reflect#DeepEqual>`__.

**Example:**

.. code-block:: html

   {{ $x := list 1 2 3 }}
   {{ $y := list "1" "2" "3" }}
   {{ deepEqual $x $y }}

The above returns

.. code-block:: text

   false

