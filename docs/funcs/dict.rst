Dictionaries and Dict Functions
*******************************

Gostub provides a key/value storage type called a dict (short for “dictionary”, as in Python). A dict is an unorder type.

The key to a dictionary must be a string. However, the value can be any type, even another dict or list.

Unlike lists, dicts are not immutable. The :ref:`func-set` and :ref:`func-unset` functions will modify the contents of a dictionary.

----

.. _func-dict:

dict
====

.. code-block:: go

   dict(items []any) map[string]any

The ``dict`` function creates a dictionary.

**Example:**

.. code-block:: html

   {{ dict "name1" "value1" "name2" "value2" "name3" "value 3" | toJson }}

The above returns

.. code-block:: text

   {"name1":"value1","name2":"value2","name3":"value 3"}

----

.. _func-get:

get
===

.. code-block:: go

   get(dict map[string]any, key string) any

The ``get`` function gets the value from the map.
If the key is not found, this operation will return "". No error will be generated.

**Example:**

.. code-block:: html

   {{ $mydict := dict "name1" "value1" "name2" "value2" "name3" "value 3" }}
   {{ get $mydict "name2" }}

The above returns

.. code-block:: text

   value2

----

.. _func-set:

set
===

.. code-block:: go

   set(dict map[string]any, key string, value any) map[string]any

The ``set`` function adds a new key/value pair to a dictionary.
Note that ``set`` returns the dictionary (a requirement of Go template functions), so you may need to trap the value as done below with the ``$_`` assignment.

**Example:**

.. code-block:: html

   {{ $mydict := dict "name1" "value1" "name2" "value2" "name3" "value 3" }}
   {{ $_ := set $mydict "name4" "value4" }}
   {{ toJson $mydict }}

The above returns

.. code-block:: text

   {"name1":"value1","name2":"value2","name3":"value 3","name4":"value4"}

----

.. _func-unset:

unset
=====

.. code-block:: go

   unset(dict map[string]any, key string) map[string]any

The ``unset`` function deletes the key from the map.
Note that if the key is not found, this operation will simply return. No error will be generated.

**Example:**

.. code-block:: html

   {{ $mydict := dict "name1" "value1" "name2" "value2" "name3" "value 3" }}
   {{ $_ := unset $mydict "name1" }}
   {{ toJson $mydict }}

The above returns

.. code-block:: text

   {"name2":"value2","name3":"value 3"}

----

.. _func-hasKey:

hasKey
======

.. code-block:: go

   hasKey(dict map[string]any, key string) bool

The ``hasKey`` function tests if the given dict contains the given key.

**Example:**

.. code-block:: html

   {{ $mydict := dict "name1" "value1" "name2" "value2" "name3" "value 3" }}
   {{ hasKey $mydict "name1" }}

The above returns

.. code-block:: text

   true

----

.. _func-pluck:

pluck
=====

.. code-block:: go

   pluck(key string, dicts ...map[string]any) []any

The ``pluck`` function searches the given key in all given dict and returns a list of coresponds value.

**Example:**

.. code-block:: html

   {{ $dict1 := dict "name1" "value1" "name2" "value2" }}
   {{ $dict2 := dict "name1" "some other value" }}
   {{ $dict3 := dict "name1" "the third value of name1 key" }}
   {{ pluck "name1" $dict1 $dict2 $dict3 | toJson }}

The above returns

.. code-block:: text

   ["value1","some other value","the third value of name1 key"]

----

.. _func-dig:

dig
===

.. code-block:: go

   dig(keys ...string, defaultValue any, dict map[string]any) any

The ``dig`` function traverses a nested set of dicts, selecting keys from a list of values.
It returns a ``defaultValue`` if any of the keys are not found at the associated dict.

**Example:**

.. code-block:: html

   {{ $userJohn := dict "firstName" "John" "lastName" "Doe" }}
   {{ $org := dict "seo" $userJohn }}
   {{ dig "seo" "firstName" "defaultValue" $org }}

The above returns

.. code-block:: text

   John

----

.. _func-merge:

merge
=====

.. code-block:: go

   merge(dicts ...map[string]any) map[string]any

The ``merge`` function merges two or more dictionaries into one, giving precedence to the dest dictionary:
This is a deep merge operation, not a deep copy operation.
Nested objects that are merged are the same instance on both dicts.

**Example:**

.. code-block:: html

   {{ $user := dict "firstName" "John" "lastName" "Doe" "orgName" ""}}
   {{ $defaults := dict "orgName" "MyOrg" }}
   {{ $dst := dict }}
   {{ merge $dst $user $defaults | toJson }}

The above returns

.. code-block:: text

   {"firstName":"John","lastName":"Doe","orgName":"MyOrg"}

----

.. _func-mustMerge:

mustMerge
=========

.. code-block:: go

   mustMerge(dicts ...map[string]any) map[string]any, error

The ``mustMerge`` function is the same as :ref:`func-merge` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ $user := dict "firstName" "John" "lastName" "Doe" "orgName" "" }}
   {{ $defaults := dict "orgName" "MyOrg" }}
   {{ $dst := dict }}
   {{ mustMerge $dst $user $defaults | toJson }}

The above returns

.. code-block:: text

   {"firstName":"John","lastName":"Doe","orgName":"MyOrg"}

----

.. _func-mergeOverwrite:

mergeOverwrite
==============

.. code-block:: go

   mergeOverwrite(dest map[string]any, sources ...map[string]any) map[string]any

The ``mergeOverwrite`` function merges two or more dictionaries into one,
giving precedence from **right to left**, effectively overwriting values in the ``dest`` dictionary.

This is a deep merge operation, not a deep copy operation.
Nested objects that are merged are the same instance on both dicts.

**Example:**

.. code-block:: html

   {{ $dst := dict "overwrite" "initialValue" "otherField" "value1" }}
   {{ $overwrite := dict "overwrite" "newValue" }}
   {{ mergeOverwrite $dst $overwrite | toJson }}

The above returns

.. code-block:: text

   {"otherField":"value1","overwrite":"newValue"}

----

.. _func-mustMergeOverwrite:

mustMergeOverwrite
==================

.. code-block:: go

   mustMergeOverwrite(dest map[string]any, sources ...map[string]any) map[string]any, error

The ``mustMergeOverwrite`` function is the same as :ref:`func-mergeOverwrite` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ $dst := dict "overwrite" "initialValue" "otherField" "value1" }}
   {{ $overwrite := dict "overwrite" "newValue" }}
   {{ mustMergeOverwrite $dst $overwrite | toJson }}

The above returns

.. code-block:: text

   {"otherField":"value1","overwrite":"newValue"}

----

.. _func-keys:

keys
====

.. code-block:: go

   keys(dicts ...map[string]any) []string

The ``keys`` function returns a list of all of the keys in one or more ``dicts``.
Since a dictionary is unordered, the keys will not be in a predictable order.

**Example:**

.. code-block:: html

   {{ dict "name1" "value1" "name2" "value2" "name3" "value 3" | keys }}

The above returns

.. code-block:: text

   [name1 name3 name2]

----

.. _func-pick:

pick
====

.. code-block:: go

   pick(dict map[string]any, keys ...string) map[string]any

The ``pick`` function selects just the given keys out of a dictionary, creating a new dict.

**Example:**

.. code-block:: html

   {{ $user := dict "firstName" "John" "lastName" "Doe" "password" "sEcReT1!" }}
   {{ pick $user "firstName" "lastName" | toJson }}

The above returns

.. code-block:: text

   {"firstName":"John","lastName":"Doe"}

----

.. _func-omit:

omit
====

.. code-block:: go

   omit(dict map[string]any, keys ...string) map[string]any

The ``omit`` function is similar to :ref:`func-pick`, except it returns a new dict with all the keys that *do not match* the given keys.

**Example:**

.. code-block:: html

   {{ $user := dict "firstName" "John" "lastName" "Doe" "password" "sEcReT1!" }}
   {{ omit $user "password" | toJson }}

The above returns

.. code-block:: text

   {"firstName":"John","lastName":"Doe"}

----

.. _func-values:

values
======

.. code-block:: go

   values(dict map[string]any) []any

The ``values`` function returns a new ``list`` with all the values of the source ``dict``.
Note that the ``values`` function gives no guarantees about the result ordering.

**Example:**

.. code-block:: html

   {{ dict "name1" "value1" "name2" "value2" "name3" "value 3" | values }}

The above returns

.. code-block:: text

   [value2 value 3 value1]

----

.. _func-deepCopy:

deepCopy
========

.. code-block:: go

   deepCopy(dict map[string]any) map[string]any

The ``deepCopy`` function makes a deep copy of dicts and other structures.
``deepCopy`` panics when there is a problem.

**Example:**

.. code-block:: html

   {{ $src := dict "name1" "value1" "name2" "value2" }}
   {{ $dst := deepCopy $src }}
   {{ $result := set $dst "name3" "value3" }}

   The original dict is not modified - {{ toJson $src }}
   The new dict - {{ toJson $result }}

The above returns

.. code-block:: text

   The original dict is not modified - {"name1":"value1","name2":"value2"}
   The new dict - {"name1":"value1","name2":"value2","name3":"value3"}

----

.. _func-mustDeepCopy:

mustDeepCopy
============

.. code-block:: go

   mustDeepCopy(dict map[string]any) map[string]any, error

The ``mustDeepCopy`` function is the same as :ref:`func-deepCopy` but returns an error if there is a problem.

**Example:**

.. code-block:: html

   {{ $src := dict "name1" "value1" "name2" "value2" }}
   {{ $dst := mustDeepCopy $src }}
   {{ $result := set $dst "name3" "value3" }}

   The original dict is not modified - {{ toJson $src }}
   The new dict - {{ toJson $result }}

The above returns

.. code-block:: text

   The original dict is not modified - {"name1":"value1","name2":"value2"}
   The new dict - {"name1":"value1","name2":"value2","name3":"value3"}

