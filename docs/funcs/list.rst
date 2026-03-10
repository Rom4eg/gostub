Lists and List Functions
************************

----

.. _func-list:

list
====

.. code-block:: go

   list(values ...any) []any

The ``list`` function creates a list.

**Example:**

.. code-block:: html

   {{ list 1 2 3 4 5 }}

The above returns

.. code-block:: text

    [1 2 3 4 5]

----

.. _func-first:

first
=====

.. code-block:: go

   first(list []any) any

The ``first`` function gets the head item of a list.
``first`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | first }}

The above returns

.. code-block:: text

    second

----

.. _func-mustFirst:

mustFirst
=========

.. code-block:: go

   mustFirst(list []any) any, error

The ``mustFirst`` function is the same as :ref:`func-first` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | mustFirst }}

The above returns

.. code-block:: text

    second

----

.. _func-rest:

rest
====

.. code-block:: go

   rest(list []any) []any

The ``rest`` function gets the tail of the ``list`` e.g everything but the first item.
``rest`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | rest }}

The above returns

.. code-block:: text

    [first third]

----

.. _func-mustRest:

mustRest
========

.. code-block:: go

   mustRest(list []any) []any, error

The ``mustRest`` function is the same as :ref:`func-rest` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | mustRest }}

The above returns

.. code-block:: text

    [first third]

----

.. _func-last:

last
====

.. code-block:: go

   last(list []any) any

The ``last`` function gets the last item on a list.
``last`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | last }}

The above returns

.. code-block:: text

    third

----

.. _func-mustLast:

mustLast
========

.. code-block:: go

   mustLast(list []any) any, error

The ``mustLast`` function is the same as :ref:`func-last` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | mustLast }}

The above returns

.. code-block:: text

    third

----

.. _func-initial:

initial
=======

.. code-block:: go

   initial(list []any) []any

The ``initial`` function returns all but the last element.
``initial`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | initial }}

The above returns

.. code-block:: text

    [second first]

----

.. _func-mustInitial:

mustInitial
===========

.. code-block:: go

   mustInitial(list []any) []any, error

The ``mustInitial`` function is the same as :ref:`func-initial` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ (list "second" "first" "third") | mustInitial }}

The above returns

.. code-block:: text

    [second first]

----

.. _func-append:

append
======

.. code-block:: go

   append(list []any, new any) []any

The ``append`` function appends a new item to an existing list, creating a new list.
``append`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ $mylist := (list "second" "first" "third") }}
   this is the new list - {{ append $mylist "fourth" }}
   and this one is old, and not modified - {{ $mylist }}

The above returns

.. code-block:: text

    this is the new list - [second first third fourth]
    and this one is old, and not modified - [second first third]

----

.. _func-mustAppend:

mustAppend
==========

.. code-block:: go

   mustAppend(list []any, new any) []any, error

The ``mustAppend`` function is the same as :ref:`func-append` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ $mylist := (list "second" "first" "third") }}
   this is the new list - {{ mustAppend $mylist "fourth" }}
   and this one is old, and not modified - {{ $mylist }}

The above returns

.. code-block:: text

    this is the new list - [second first third fourth]
    and this one is old, and not modified - [second first third]

----

.. _func-prepend:

prepend
=======

.. code-block:: go

   prepend(list []any, new any) []any

The ``prepend`` function prepends an element to the front of a list, creating a new list.
``prepend`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ $mylist := (list "second" "first" "third") }}
   this is the new list - {{ prepend $mylist "zero" }}
   and this one is old, and not modified - {{ $mylist }}

The above returns

.. code-block:: text

   this is the new list - [zero second first third]
   and this one is old, and not modified - [second first third]

----

.. _func-mustPrepend:

mustPrepend
===========

.. code-block:: go

   mustPrepend(list []any, new any) []any, error

The ``mustPrepend`` function is the same as :ref:`func-prepend` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ $mylist := (list "second" "first" "third") }}
   this is the new list - {{ mustPrepend $mylist "zero" }}
   and this one is old, and not modified - {{ $mylist }}

The above returns

.. code-block:: text

   this is the new list - [zero second first third]
   and this one is old, and not modified - [second first third]

----

.. _func-concat:

concat
======

.. code-block:: go

   concat(first ...[]any) []any

The ``concat`` function concatenates an arbitrary number of lists into one.
``concat`` creates a new list reather modifying the existing one.

**Example:**

.. code-block:: html

   {{ $mylist := (list "second" "first") }}
   {{ $mylist2 := (list "third") }}
   {{ concat $mylist $mylist2 }}

The above returns

.. code-block:: text

   [second first third]

----

.. _func-reverse:

reverse
=======

.. code-block:: go

   reverse(item []any) []any

The ``reverse`` function produces a new list with the reversed elements of the given list.
``reverse`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ list "first" "second" "third" | reverse }}

The above returns

.. code-block:: text

   [third second first]

----

.. _func-mustReverse:

mustReverse
===========

.. code-block:: go

   mustReverse(item []any) []any, error

The ``mustReverse`` function is the same as :ref:`func-reverse` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ list "first" "second" "third" | mustReverse }}

The above returns

.. code-block:: text

   [third second first]

----

.. _func-uniq:

uniq
====

.. code-block:: go

   uniq(item []any) []any

The ``uniq`` function generates a new list with all of the duplicates removed.
``uniq`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ list 1 1 1 2 3 | uniq }}

The above returns

.. code-block:: text

   [1 2 3]

----

.. _func-mustUniq:

mustUniq
========

.. code-block:: go

   mustUniq(item []any) []any, error

The ``mustUniq`` function is the same as :ref:`func-uniq` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ list 1 1 1 2 3 | mustUniq }}

The above returns

.. code-block:: text

   [1 2 3]

----

.. _func-without:

without
=======

.. code-block:: go

   without(list []any, filter ...any) []any

The ``without`` function filters items out of a list.
``without`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ $mylist := list "first" "second" "third" }}
   {{ without $mylist "second" }}

The above returns

.. code-block:: text

   [first third]

----

.. _func-mustWithout:

mustWithout
===========

.. code-block:: go

   mustWithout(list []any, filter ...any) []any, error

The ``mustWithout`` function is the same as :ref:`func-without` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ $mylist := list "first" "second" "third" }}
   {{ mustWithout $mylist "second" }}

The above returns

.. code-block:: text

   [first third]

----

.. _func-has:

has
===

.. code-block:: go

   has(item any, list []any) bool

The ``has`` function tests to see if a list has a particular element.
``has`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ list "first" "second" "third" | has "third" }}

The above returns

.. code-block:: text

   true

----

.. _func-mustHas:

mustHas
=======

.. code-block:: go

   mustHas(item any, list []any) bool, error

The ``mustHas`` function is the same as :ref:`func-has` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ list "first" "second" "third" | mustHas "third" }}

The above returns

.. code-block:: text

   true

----

.. _func-compact:

compact
=======

.. code-block:: go

   compact(list []any) []any

The ``compact`` function removes entries with empty values.
``compact`` panics if there is a problem.

**Example:**

.. code-block:: html

   {{ list "first" "second" "" "third" | compact }}

The above returns

.. code-block:: text

   [first second third]

----

.. _func-mustCompact:

mustCompact
===========

.. code-block:: go

   mustCompact(list []any) []any, error

The ``mustCompact`` function is the same as :ref:`func-compact` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ list "first" "second" "" "third" | mustCompact }}

The above returns

.. code-block:: text

   [first second third]

----

.. _func-slice:

slice
=====

.. code-block:: go

   slice(list []any, start int, end int) []any

The ``slice`` function gets partial elements (e.g. the slice) of a list.

**Example:**

.. code-block:: html

   {{ slice (list "first" "second" "third" "fourth") 1 3 }}

The above returns

.. code-block:: text

   [second third]

----

.. _func-mustSlice:

mustSlice
=========

.. code-block:: go

   mustSlice(list []any, start int, end int) []any, error

The ``mustSlice`` function is the same as :ref:`func-slice` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ mustSlice (list "first" "second" "third" "fourth") 1 3 }}

The above returns

.. code-block:: text

   [second third]

----

.. _func-chunk:

chunk
=====

.. code-block:: go

   chunk(size int, list []any) [][]any

The ``chunk`` function splits a list into chunks of given size.

**Example:**

.. code-block:: html

   {{ chunk 3 (list 1 2 3 4 5 6 7 8) }}

The above returns

.. code-block:: text

   [[1 2 3] [4 5 6] [7 8]]

