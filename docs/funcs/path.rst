Path Functions
**************

While gostub does not grant access to the filesystem, it does provide functions for working with strings that follow file path conventions.

Paths separated by the slash character (/), are processed by the `path <https://pkg.go.dev/path>`__ package.

**Examples:**

|   The Linux and MacOS filesystems: ``/home/user/file``, ``/etc/config``.
|   The path component of URIs: ``https://example.com/some/content/``, ``ftp://example.com/file/``.


----

.. _func-base:

base
====

.. code-block:: go

   base(path string) string

The ``base`` function returns the last element of a ``path``.

**Example:**

.. code-block:: html

   {{ base "/home/userName/docs" }}

The above returns

.. code-block:: text

   docs

----

.. _func-dir:

dir
===

.. code-block:: go

   dir(path string) string

The ``dir`` function returns the directory, stripping the last part of the ``path``.

**Example:**

.. code-block:: html

   {{ dir "/home/userName/docs" }}

The above returns

.. code-block:: text

   /home/userName

----

.. _func-clean:

clean
=====

.. code-block:: go

   clean(path string) string

The ``clean`` function cleans a ``path``.

**Example:**

.. code-block:: html

   {{ clean "/home/../etc/docs" }}

The above returns

.. code-block:: text

   /etc/docs

----

.. _func-ext:

ext
===

.. code-block:: go

   ext(path string) string

The ``ext`` function returns the file extension.

**Example:**

.. code-block:: html

   {{ ext "/home/docs/readme.txt" }}

The above returns

.. code-block:: text

   .txt

----

.. _func-isAbs:

isAbs
=====

.. code-block:: go

   isAbs(path string) bool

The ``isAbs`` function tests whether a ``path`` is absolute.

**Example:**

.. code-block:: html

   {{ isAbs "/home/docs/readme.txt" }}

The above returns

.. code-block:: text

   true

