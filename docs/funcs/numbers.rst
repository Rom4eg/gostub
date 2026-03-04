Integer Math Functions
**********************

The following math functions operate on ``int64`` values.

----

.. _func-add:

add
===

.. code-block:: go

   add(numbers ...int64) int64

The ``add`` function sums numbers.

**Example:**

.. code-block:: html

   {{ add 1 2 3 }}

The above returns

.. code-block:: text

   6
----

.. _func-add1:

add1
====

.. code-block:: go

   add1(a int64) int64

The ``add1`` function increments ``a`` by 1.

**Example:**

.. code-block:: html

   {{ add1 3 }}

The above returns

.. code-block:: text

   4

----

.. _func-sub:

sub
===

.. code-block:: go

   sub(a int64, b int64) int64

The ``sub`` function subtracts ``b`` from ``a``.

**Example:**

.. code-block:: html

   {{ sub 5 3 }}

The above returns

.. code-block:: text

   2

----

.. _func-div:

div
===

.. code-block:: go

   div(a int64, b int64) int64

The ``div`` function divides ``a`` on ``b``.

**Example:**

.. code-block:: html

   {{ div 10 2 }}

The above returns

.. code-block:: text

   5

----

.. _func-mod:

mod
===

.. code-block:: go

   mod(a int64, b int64) int64

The ``mod`` function computes the remainder after dividing ``a`` by ``b``.

**Example:**

.. code-block:: html

   {{ mod 10 3 }}

The above returns

.. code-block:: text

   1

----

.. _func-mul:

mul
===

.. code-block:: go

   mul(a int64, b int64) int64

The ``mul`` function multiplies ``a`` by ``b``.

**Example:**

.. code-block:: html

   {{ mul 4 2 }}

The above returns

.. code-block:: text

   8

----

.. _func-max:

max
===

.. code-block:: go

   max(numbers ...int64) int64

The ``max`` function returns the largest of a series of integers.

**Example:**

.. code-block:: html

   {{ max 1 3 5 11 8 4 }}

The above returns

.. code-block:: text

   11

----

.. _func-min:

min
===

.. code-block:: go

   min(numbers ...int64) int64

The ``min`` function returns the smallest of a series of integers.

**Example:**

.. code-block:: html

   {{ min 1 3 5 11 8 4 }}

The above returns

.. code-block:: text

   1

----

.. _func-floor:

floor
=====

.. code-block:: go

   floor(a float64) float64

The ``floor`` function returns the greatest float value less than or equal to the input value

**Example:**

.. code-block:: html

   {{ floor 123.9999 }}

The above returns

.. code-block:: text

   123

----

.. _func-ceil:

ceil
====

.. code-block:: go

   ceil(a float64) float64

The ``ceil`` function returns the greatest float value greater than or equal to the input value.

**Example:**

.. code-block:: html

   {{ ceil 123.001 }}

The above returns

.. code-block:: text

   124

----

.. _func-round:

round
=====

.. code-block:: go

   round(a float64, n int) float64

The ``round`` function returns a float value rounded to the given number of digits after the decimal point.

**Example:**

.. code-block:: html

   {{ round 123.555555 3 }}

The above returns

.. code-block:: text

   123.556

----

.. _func-randInt:

randInt
=======

.. code-block:: go

   randInt(min int, max int) int

The ``randInt`` function returns a random integer value from min (inclusive) to max (exclusive).
The ``randInt`` function will produce a random number in the range [min,max).

**Example:**

.. code-block:: html

   {{ randInt 12 30 }}

The above returns

.. code-block:: text

   20

