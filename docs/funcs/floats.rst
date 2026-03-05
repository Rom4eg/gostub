Float Math Functions
********************

All math functions operate on ``float64`` values.

----

.. _func-addf:

addf
====

.. code-block:: go

   addf(numbers ...float64) float64

The ``addf`` function sums numbers.

**Example:**

.. code-block:: html

   {{ addf 2.2 2 3 }}

The above returns

.. code-block:: text

   7.2

----

.. _func-add1f:

add1f
=====

.. code-block:: go

   add1f(num float64) float64

The ``add1f`` function increments ``num`` by 1.

**Example:**

.. code-block:: html

   {{ add1f 1.3 }}

The above returns

.. code-block:: text

   2.3

----

.. _func-subf:

subf
====

.. code-block:: go

   subf(a float64, nums ...float64) float64

The ``subf`` function subtracts all ``nums`` from ``a``.

**Example:**

.. code-block:: html

   {{ subf 7.5 2 1 }}

The above returns

.. code-block:: text

   4.5

----

.. _func-divf:

divf
====

.. code-block:: go

   divf(a float64, nums ...float64) float64

The ``divf`` function divides ``a`` by each of ``nums``.

**Example:**

.. code-block:: html

   {{ divf 10 2 4 }}

The above returns

.. code-block:: text

   1.25

----

.. _func-mulf:

mulf
====

.. code-block:: go

   mulf(a float64, nums ...float64) float64

The ``mulf`` function multiplies ``a`` on each of ``nums``.

**Example:**

.. code-block:: html

   {{ mulf 1.5 2 2 }}

The above returns

.. code-block:: text

   6

----

.. _func-maxf:

maxf
====

.. code-block:: go

   maxf(nums ...float64) float64

The ``maxf`` function returns the largest of a series of floats.

**Example:**

.. code-block:: html

   {{ maxf 1 2.5 3 }}

The above returns

.. code-block:: text

   3

----

.. _func-minf:

minf
====

.. code-block:: go

   minf(nums ...float64) float64

The ``minf`` function returns the smallest of a series of floats.

**Example:**

.. code-block:: html

   {{ minf 1.5 2 3 }}

The above returns

.. code-block:: text

   1.5

