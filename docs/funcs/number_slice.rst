Integer Slice Functions
***********************

.. _func-until:

until
=====

.. code-block:: go

   until(n int) []int

The ``until`` function builds a range of integers.

**Example:**

.. code-block:: html

   {{ until 5 }}

The above returns

.. code-block:: text

   [0 1 2 3 4]

----

.. _func-untilStep:

untilStep
=========

.. code-block:: go

   untilStep(from int, to int, step int) []int

The ``untilStep`` function works like :ref:`func-until`, but it allows you to define a start, stop, and step.

**Example:**

.. code-block:: html

   {{ untilStep 3 6 2 }}

The above returns

.. code-block:: text

   [3 5]

----

.. _func-seq:

seq
===

.. code-block:: go

   seq(first int, second int, third int) []int

The ``seq`` function works like the bash seq command.

* 1 parameter (end) - will generate all counting integers between 1 and end inclusive.
* 2 parameters (start, end) - will generate all counting integers between start and end inclusive incrementing or decrementing by 1.
* 3 parameters (start, step, end) - will generate all counting integers between start and end inclusive incrementing or decrementing by step.

**Example:**

.. code-block:: html

   {{ seq 5 }}
   {{ seq -3 }}
   {{ seq 0 2 }}
   {{ seq 2 -2 }}
   {{ seq 0 2 10 }}
   {{ seq 0 -2 -5 }}

The above returns

.. code-block:: text

   [1 2 3 4 5]
   [1 0 -1 -2 -3]
   [0 1 2]
   [2 1 0 -1 -2]
   [0 2 4 6 8 10]
   [0 -2 -4]

