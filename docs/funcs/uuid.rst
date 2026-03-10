UUID Functions
**************

Gostub can generate UUID v4 universally unique IDs.

----

.. _func-uuidv4:

uuidv4
======

.. code-block:: go

   uuidv4() string

The ``uuidv4`` function generates a new UUID.

**Example:**

.. code-block:: html

   {{ uuidv4 }}

The above returns

.. code-block:: text

   bbb2c6da-8485-491b-a3f8-59a7fa7f656c
