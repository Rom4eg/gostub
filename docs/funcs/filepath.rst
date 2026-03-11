Filepath Functions
******************

While gostub does not grant access to the filesystem, it does provide functions for working with strings that follow file path conventions.

Paths are separated by the os.PathSeparator variable, are processed by the `path/filepath <https://pkg.go.dev/path/filepath>`__ package.

These are the recommended functions to use when parsing paths of local filesystems, usually when dealing with local files, directories, etc.

**Examples:**

|   Running on Linux or MacOS the filesystem path is separated by the slash character (``/``): ``/home/user/file``, ``/etc/config``.
|   Running on Windows the filesystem path is separated by the backslash character (``\``): ``C:\Users\Username\``, ``C:\Program Files\Application\``.

----

.. _func-osBase:

osBase
======

.. code-block:: go

   osBase(path string) string

The ``osBase`` function returns the last element of a ``path``.

**Example:**

.. code-block:: html

   {{ osBase "/home/userName/docs" }}

The above returns

.. code-block:: text

   docs

----

.. _func-osDir:

osDir
=====

.. code-block:: go

   osDir(path string) string

The ``osDir`` function returns the directory, stripping the last part of the ``path``.

**Example:**

.. code-block:: html

   {{ osDir "/home/userName/docs" }}

The above returns

.. code-block:: text

   /home/userName

----

.. _func-osClean:

osClean
=======

.. code-block:: go

   osClean(path string) string

The ``osClean`` function cleans a ``path``.

**Example:**

.. code-block:: html

   {{ osClean "/home/../etc/docs" }}

The above returns

.. code-block:: text

   /etc/docs

----

.. _func-osExt:

osExt
=====

.. code-block:: go

   osExt(path string) string

The ``osExt`` function returns the file extension.

**Example:**

.. code-block:: html

   {{ osExt "/home/docs/readme.txt" }}

The above returns

.. code-block:: text

   .txt

----

.. _func-osIsAbs:

osIsAbs
=======

.. code-block:: go

   osIsAbs(path string) bool

The ``osIsAbs`` function tests whether a ``path`` is absolute.

**Example:**

.. code-block:: html

   {{ osIsAbs "/home/docs/readme.txt" }}

The above returns

.. code-block:: text

   true

