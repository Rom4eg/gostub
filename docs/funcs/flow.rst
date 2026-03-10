Flow Control Functions
**********************

----

.. _func-fail:

fail
====

.. code-block:: go

   fail(text string) string, error

The ``fail`` function unconditionally returns an empty ``string`` and an ``error`` with the specified ``text``.

**Example:**

.. code-block:: html

   {{ fail "Please accept the end user license agreement" }}

The above returns

.. code-block:: text

   template: index:3:6: executing "main" at <fail "Please accept the end user license agreement">: error calling fail: Please accept the end user license agreement
