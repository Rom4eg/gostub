Date Functions
**************

----

.. _func-now:

now
====

.. code-block:: go

   now() time.Time

The ``now`` function returns the current date/time.

**Example:**

.. code-block:: html

   {{ (now).Hour }}

The above returns

.. code-block:: text

   19

----

.. _func-ago:

ago
===

.. code-block:: go

   ago(time time.Time) time.Duration

The ``ago`` function returns the duration from time.Now in seconds resolution.

**Example:**

.. code-block:: html

   {{ $n := (now.AddDate 0 0 -1) }}
   {{ ago $n }}

The above returns

.. code-block:: text

   24h0m0s

----

.. _func-date:

date
====

.. code-block:: go

   date(fmt string, time time.Time) string

The ``date`` function formats a date.

Check out the official `docs <https://pkg.go.dev/time#pkg-constants>`__ for date/time formatting.

**Example:**

.. code-block:: html

   {{ now | date "2006-01-02 15:04:05" }}

The above returns

.. code-block:: text

   2026-03-05 20:08:02

----

.. _func-dateInZone:

dateInZone
==========

.. code-block:: go

   dateInZone(fmt string, time time.Time, tz string) string

The ``dateInZone`` function is the same as :ref:`func-date`, but with a timezone.

**Example:**

.. code-block:: html

   {{ dateInZone "2006-01-02 15:04:05" (now) "UTC" }}

The above returns

.. code-block:: text

   2026-03-05 17:08:02

----

.. _func-duration:

duration
========

.. code-block:: go

   duration(seconds string) time.Duration

The ``duration`` function formats a given number of seconds as a ``time.Duration``.

**Example:**

.. code-block:: html

   {{ duration "321" }}

The above returns

.. code-block:: text

   5m21s

----

.. _func-durationRound:

durationRound
=============

.. code-block:: go

   durationRound(duration any) time.Duration

The ``durationRound`` function rounds a given duration to the most significant unit.
Strings and ``time.Duration`` are parsed as a duration, while a ``time.Time`` is calculated as the duration since.

**Example:**

.. code-block:: html

   {{ durationRound "2h10m5s" }}

The above returns

.. code-block:: text

   2h

----

.. _func-unixEpoch:

unixEpoch
=========

.. code-block:: go

   unixEpoch(d time.Time) string

The ``unixEpoch`` function returns the seconds since the unix epoch for a ``time.Time``.

**Example:**

.. code-block:: html

   {{ now | unixEpoch }}

The above returns

.. code-block:: text

   1772731432

----

.. _func-dateModify:

dateModify
==========

.. code-block:: go

   dateModify(mod string, t time.Time) time.Time

The ``dateModify`` function takes a `modification <https://pkg.go.dev/time#ParseDuration>`__ and a date and returns the modified ``time.Time``.

**Example:**

.. code-block:: html

   {{ now | dateModify "-1.5h" | date "2006-01-02 15:04:05" }}

The above returns

.. code-block:: text

   2026-03-05 19:01:35

----

.. _func-mustDateModify:

mustDateModify
==============

.. code-block:: go

   mustDateModify(mod string, t time.Time) time.Time, error

The ``mustDateModify`` function is the same as :ref:`func-dateModify` but will return an error if the modification format is wrong.

**Example:**

.. code-block:: html

   {{ now | mustDateModify "-1.5h" | date "2006-01-02 15:04:05" }}

The above returns

.. code-block:: text

   2026-03-05 19:01:35

----

.. _func-htmlDate:

htmlDate
========

.. code-block:: go

   htmlDate(t time.Time) string

The ``htmlDate`` function formats a date for inserting into an HTML date picker input field.

**Example:**

.. code-block:: html

   {{ now | htmlDate }}

The above returns

.. code-block:: text

   2026-03-05

----

.. _func-htmlDateInZone:

htmlDateInZone
==============

.. code-block:: go

   htmlDateInZone(t time.Time, tz string) string

The ``htmlDateInZone`` function is the same as :ref:`func-htmlDate`, but with a timezone.

**Example:**

.. code-block:: html

   {{ htmlDateInZone (now) "UTC" }}

The above returns

.. code-block:: text

   2026-03-05

----

.. _func-toDate:

toDate
======

.. code-block:: go

   toDate(fmt string, dateTime string) time.Time

The ``toDate`` function converts a string to a date.
The first argument is the date layout and the second is the date string

Check out the official `docs <https://pkg.go.dev/time#pkg-constants>`__ for date/time formatting.

**Example:**

.. code-block:: html

   {{ toDate "2006-01-02" "2025-11-28" | date "01/02/2006" }}

The above returns

.. code-block:: text

   11/28/2025

----

.. _func-mustToDate:

mustToDate
==========

.. code-block:: go

   mustToDate(fmt string, dateTime string) time.Time, error

The ``mustToDate`` is the same as :ref:`func-toDate` but returns an error in case the string cannot be converted.

**Example:**

.. code-block:: html

   {{ mustToDate "2006-01-02" "2025-11-28" | date "01/02/2006" }}

The above returns

.. code-block:: text

   11/28/2025

