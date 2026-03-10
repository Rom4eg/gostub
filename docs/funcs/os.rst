OS Functions
************

----

.. _func-env:

env
===

.. code-block:: go

   env(name string) string

The ``env`` function reads an environment variable.

**Example:**

.. code-block:: html

   {{ env "GO111MODULE" }}

The above returns

.. code-block:: text

   on

----

.. _func-expandenv:

expandenv
=========

.. code-block:: go

   expandenv(text string) string

The ``expandenv`` function substitutes environment variables in a ``text``.

**Example:**

.. code-block:: html

   {{ expandenv "Force to use modern Go modules is - $GO111MODULE" }}

The above returns

.. code-block:: text

   Force to use modern Go modules is - on

