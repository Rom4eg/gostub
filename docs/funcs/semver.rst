Semantic Version Functions
**************************

Some version schemes are easily parseable and comparable. gostub provides functions for working with `SemVer 2 <http://semver.org/>`__ versions.

gostub uses the `Masterminds semver package <https://github.com/Masterminds/semver>`__


.. _struct-version:

Version object
==============

:Major: The major number (1 above)
:Minor: The minor number (2 above)
:Patch: The patch number (3 above)
:Prerelease: The prerelease (alpha.1 above)
:Metadata: The build metadata (123 above)
:Original: The original version as a string

----

.. _func-semver:

semver
======

.. code-block:: go

   semver(ver string) *Version

The ``semver`` function parses a string into a :ref:`struct-version`.

**Example:**

.. code-block:: html

   {{ $ver := semver "1.4.2" }}
   {{ $ver.Minor }}

The above returns

.. code-block:: text

   4

----

.. _func-semverCompare:

semverCompare
=============

.. code-block:: go

   semverCompare(first string, second string) bool

The ``semverCompare`` function returns ``true`` if the constraint matches, or ``false`` if it does not match.

**Example:**

.. code-block:: html

   {{ semverCompare "^1.2.0" "1.2.3" }}

The above returns

.. code-block:: text

   true

