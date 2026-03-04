Strings
*******

.. _func-default:

trim
====

.. code-block:: go

   trim(value string) string

The ``trim`` function removes space from either side of a ``value``

**Example:**

.. code-block:: html

   This {{ trim "       value      " }} has been truncated on both the left and right sides.

The above returns

.. code-block:: text

   This value has been truncated on both the left and right sides.

----

.. _func-trimAll:

trimAll
=======

.. code-block:: go

   trimAll(trimChar string, value string) string

The ``trimAll`` function remove given characters from the front or back of a string

**Example:**

.. code-block:: html

   The price is {{ trimAll "$" "$5.00" }}

The above returns

.. code-block:: text

   The price is 5.00

----

.. _func-trimSuffix:

trimSuffix
==========

.. code-block:: go

   trimSuffix(trimChar string, value string) string

The ``trimSuffix`` function trim just the suffix from a string

**Example:**

.. code-block:: html

   The time is {{ trimSuffix "+0000" "23:00:00+0000" }}

The above returns

.. code-block:: text

   The time is 23:00:00

----

.. _func-trimPrefix:

trimPrefix
==========

.. code-block:: go

   trimPrefix(trimChar string, value string) string

The ``trimPrefix`` function trim just the prefix from a string

**Example:**

.. code-block:: html

   {{ trimPrefix "--" "--name" }}

The above returns

.. code-block:: text

   name

----

.. _func-upper:

upper
=====

.. code-block:: go

   upper(value string) string

The ``upper`` function convert the entire string to uppercase

**Example:**

.. code-block:: html

   {{ upper "this text is uppercase"}}

The above returns

.. code-block:: text

   THIS TEXT IS UPPERCASE

----

.. _func-lower:

lower
=====

.. code-block:: go

   lower(value string) string

The ``lower`` function convert the entire string to lowercase

**Example:**

.. code-block:: html

   {{ lower "THIS TEXT IS LOWERCASE"}}

The above returns

.. code-block:: text

   this text is lowercase

----

.. _func-title:

title
=====

.. code-block:: go

   title(value string) string

The ``title`` function convert to title case

**Example:**

.. code-block:: html

   {{ untitle "THIS TEXT IS LOWERCASE" }}

The above returns

.. code-block:: text

   Hello World

----

.. _func-untitle:

untitle
=======

.. code-block:: go

   untitle(value string) string

The ``untitle`` function remove title casing

**Example:**

.. code-block:: html

   {{ untitle "This text Is UTITLED" }}

The above returns

.. code-block:: text

   this text is uTITLED

----

.. _func-repeat:

repeat
======

.. code-block:: go

   repeat(count int, value string) string

The ``repeat`` function repeat a string multiple times

**Example:**

.. code-block:: html

   {{ repeat 3 "-||-" }}

The above returns

.. code-block:: text

   -||--||--||-

----

.. _func-substr:

substr
======

.. code-block:: go

   substr(start int, end int, value string) string

The ``substr`` function get a substring from a string.

**Example:**

.. code-block:: html

   {{ substr 1 4 "Hello" }}

The above returns

.. code-block:: text

   ell

----

.. _func-nospace:

nospace
=======

.. code-block:: go

   nospace(value string) string

The ``nospace`` function remove all whitespace from a string.

**Example:**

.. code-block:: html

   {{ nospace "hello w o r l d" }}

The above returns

.. code-block:: text

   helloworld

----

.. _func-trunc:

trunc
=====

.. code-block:: go

   trunc(length int, value string) string

The ``trunc`` function truncate a string to the specified length

**Example:**

.. code-block:: html

   {{ trunc 4 "some message" }}

The above returns

.. code-block:: text

   some

----

.. _func-abbrev:

abbrev
======

.. code-block:: go

   abbrev(length int, value string) string

The ``abbrev`` function truncate (:ref:`func-trunc`) a string with ellipses

**Example:**

.. code-block:: html

   {{ abbrev 16 "some message with very very very long text" }}

The above returns

.. code-block:: text

   some message ...

----

.. _func-abbrevboth:

abbrevboth
==========

.. code-block:: go

   abbrevboth(offset int, size int, value string) string

The ``abbrevboth`` function abbreviate both sides (:ref:`func-abbr`) a string with ellipses

**Example:**

.. code-block:: html

   {{ abbrevboth 13 20 "some message with very very very long text" }}

The above returns

.. code-block:: text

   ...with very very...

----

.. _func-initials:

initials
========

.. code-block:: go

   initials(value string) string

The ``initials`` function given multiple words, take the first letter of each word and combine

**Example:**

.. code-block:: html

   {{ initials "Roman Paranichev" }}

The above returns

.. code-block:: text

   RP

----

.. _func-randAlphaNum:

randAlphaNum
============

.. code-block:: go

   randAlphaNum(length int) string

The ``randAlphaNum`` generate cryptographically secure (uses crypto/rand) random strings.
Uses 0-9a-zA-Z.

**Example:**

.. code-block:: html

   {{ randAlphaNum 32 }}

The above returns

.. code-block:: text

   TT4NrqsToV5ZB3ePlkmVc9AAJLM85Eu4

----

.. _func-randAlpha:

randAlpha
=========

.. code-block:: go

   randAlpha(length int) string

The ``randAlpha`` is the same as :ref:`func-randAlphaNum` but uses ``a-zA-Z`` character set

**Example:**

.. code-block:: html

   {{ randAlpha 16 }}

The above returns

.. code-block:: text

   JVGBXasLLuIjhADl

----

.. _func-randNumeric:

randNumeric
===========

.. code-block:: go

   randNumeric(length int) string

The ``randNumeric`` is the same as :ref:`func-randAlphaNum` but uses ``0-9`` character set

**Example:**

.. code-block:: html

   {{ randNumeric 64 }}

The above returns

.. code-block:: text

   2969255249870830367702966609597151440068847730779083921815604709

----

.. _func-randAscii:

randAscii
=========

.. code-block:: go

   randAscii(length int) string

The ``randAscii`` is the same as :ref:`func-randAlphaNum` but uses all printable ASCII characters

**Example:**

.. code-block:: html

   {{ randAscii 32 }}

The above returns

.. code-block:: text

   xR}l{}B>6:8{Y'n8r&{N/h$/'oXdZ& Q

----

.. _func-wrap:

wrap
====

.. code-block:: go

   wrap(length int, value string) string

The ``wrap`` function wrap text at a given column count

**Example:**

.. code-block:: html

   {{ wrap 10 "some very very long text" }}

The above returns

.. code-block:: text

   some very
   very long
   text

----

.. _func-wrapWith:

wrapWith
========

.. code-block:: go

   wrapWith(length int, wrapChar string, value string) string

The ``wrapWith`` function works as :ref:`func_wrap`, but lets you specify the string to wrap with.
(:ref:`func_wrap` uses \n)

**Example:**

.. code-block:: html

   {{ wrapWith 10 "%" "some very very long text" }}

The above returns

.. code-block:: text

   some very%very long%text

----

.. _func-contains:

contains
========

.. code-block:: go

   contains(substr string, value string) bool

The ``contains`` function test if one string is contained inside of another

**Example:**

.. code-block:: html

   {{ contains "cat" "catch" }}

The above returns

.. code-block:: text

   true

----

.. _func-hasPrefix:

hasPrefix
=========

.. code-block:: go

   hasPrefix(substr string, value string) bool

The ``hasPrefix`` function test whether a string has a given prefix

**Example:**

.. code-block:: html

   {{ hasPrefix "cat" "catch" }}

The above returns

.. code-block:: text

   true

----

.. _func-hasSuffix:

hasSuffix
=========

.. code-block:: go

   hasSuffix(substr string, value string) bool

The ``hasSuffix`` function test whether a string has a given suffix

**Example:**

.. code-block:: html

   {{ hasSuffix "cat" "catch" }}

The above returns

.. code-block:: text

   false

----

.. _func-quote:

quote
=====

.. code-block:: go

   quote(value string) string

The ``quote`` function wrap a string in double quotes

**Example:**

.. code-block:: html

   {{ quote "double quotes" }}

The above returns

.. code-block:: text

   "double quotes"

----

.. _func-squote:

squote
======

.. code-block:: go

   squote(value string) string

The ``squote`` function wrap a string in single quotes

**Example:**

.. code-block:: html

   {{ squote "single quotes" }}

The above returns

.. code-block:: text

   'single quotes'

----

.. _func-cat:

cat
===

.. code-block:: go

   cat(value ...string) string

The ``cat`` function concatenates multiple strings together into one, separating them with spaces

**Example:**

.. code-block:: html

   {{ cat "hello" "beautiful" "world" }}

The above returns

.. code-block:: text

   hello beautiful world

----

.. _func-indent:

indent
======

.. code-block:: go

   indent(width int, value string) string

The ``indent`` function indents every line in a given string to the specified indent width.
This is useful when aligning multi-line strings

**Example:**

.. code-block:: html

   {{ indent 6 "first\nsecond\nthird" }}

----

.. _func-nindent:

nindent
=======

.. code-block:: go

   nindent(width int, value string) string

The ``nindent`` function is the same as the :ref:`func_indent` function, but prepends a new line to the beginning of the string.

**Example:**

.. code-block:: html

   {{ nindent 6 "first\nsecond\nthird" }}

----

.. _func-replace:

replace
=======

.. code-block:: go

   replace(replace string, replaceWith string, source string) string

The ``nindent`` function perform simple string replacement

**Example:**

.. code-block:: html

   {{ "I Am Henry VIII" | replace " " "-" }}

The above returns

.. code-block:: text

   I-Am-Henry-VIII

----

.. _func-plural:

plural
======

.. code-block:: go

   plural(singular string, plural string, length int) string

The ``plural`` function pluralize a string.

It does not currently support languages with more complex pluralization rules.
And 0 is considered a plural because the English language treats it as such “zero days”.

**Example:**

.. code-block:: html

   {{ $days := (.Request.URL.Query.Get "days") | default 0}}
   this task takes up to {{ $days }} {{ plural "day" "days" (int $days) }}

The above returns

.. code-block:: text

   curl "http://localhost:8080?days=1"

.. code-block:: text

   this task takes up to 1 day

----

.. _func-snakecase:

snakecase
=========

.. code-block:: go

   snakecase(value string) string

The ``snakecase`` function convert string from camelCase to snake_case.

**Example:**

.. code-block:: html

   {{ snakecase "FirstName" }}

The above returns

.. code-block:: text

   first_name

----

.. _func-camelcase:

camelcase
=========

.. code-block:: go

   camelcase(value string) string

The ``camelcase`` function convert string from snake_case to camelCase.

**Example:**

.. code-block:: html

   {{ camelcase "first_name" }}

The above returns

.. code-block:: text

   FirstName


----

.. _func-kebabcase:

kebabcase
=========

.. code-block:: go

   kebabcase(value string) string

The ``kebabcase`` function convert string from camelCase to kebab-case.

**Example:**

.. code-block:: html

   {{ kebabcase "FirstName" }}

The above returns

.. code-block:: text

   first-name

----

.. _func-swapcase:

swapcase
========

.. code-block:: go

   swapcase(value string) string

The ``swapcase`` function swap the case of a string using a word based algorithm.

Conversion algorithm:

* Upper case character converts to Lower case
* Title case character converts to Lower case
* Lower case character after Whitespace or at start converts to Title case
* Other Lower case character converts to Upper case
* Whitespace is defined by unicode.IsSpace(char)

**Example:**

.. code-block:: html

   {{ swapcase "This Is A.Test" }}

The above returns

.. code-block:: text

   tHIS iS a.tEST

----

.. _func-shuffle:

shuffle
=======

.. code-block:: go

   shuffle(value string) string

The ``shuffle`` function shuffle a string

**Example:**

.. code-block:: html

   {{ shuffle "Hello world !" }}

The above returns

.. code-block:: text

   w!l lole Hdor

----

.. _func-regexMatch:
.. _func-mustRegexMatch:

regexMatch, mustRegexMatch
==========================

.. code-block:: go

   regexMatch(re string, value string) bool
   mustRegexMatch(re string, value string) bool, error

The ``regexMatch`` function returns true if the input string contains any match of the regular expression.
``regexMatch`` panics if there is a problem and ``mustRegexMatch`` returns an error to the template engine if there is a problem.

**Example:**

.. code-block:: html

   {{ regexMatch "^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}$" "test@acme.com" }}

The above returns

.. code-block:: text

   true

----

.. _func-regexFindAll:
.. _func-mustRegexFindAll:

regexFindAll, mustRegexFindAll
==============================

.. code-block:: go

   regexFindAll(re string, value string, count int) []string
   mustRegexFindAll(re string, value string, count int) []string, error

The ``regexFindAll`` function returns a slice of all matches of the regular expression in the input string.
The last parameter n determines the number of substrings to return, where -1 means return all matches

``regexFindAll`` panics if there is a problem and ``mustRegexFindAll`` returns an error to the template engine if there is a problem.

**Example:**

.. code-block:: html

   {{ regexFindAll "[2,4,6,8]" "123456789" -1 }}

The above returns

.. code-block:: text

   [2 4 6 8]

----

.. _func-regexFind:
.. _func-mustRegexFind:

regexFind, mustRegexFind
========================

.. code-block:: go

   regexFind(re string, value string) string
   mustRegexFind(re string, value string) string, error

The ``regexFind`` function return the first (left most) match of the regular expression in the input string

``regexFind`` panics if there is a problem and ``mustRegexFind`` returns an error to the template engine if there is a problem.

**Example:**

.. code-block:: html

   {{ regexFind "[a-zA-Z][1-9]" "abcd1234" }}

The above returns

.. code-block:: text

   d1

----

.. _func-regexReplaceAll:
.. _func-mustRegexReplaceAll:

regexReplaceAll, mustRegexReplaceAll
====================================

.. code-block:: go

   regexReplaceAll(re string, value string, replace string) string
   mustRegexReplaceAll(re string, value string, replace string) string, error

The ``regexReplaceAll`` function returns a copy of the input string, replacing matches of the Regexp with the replacement string replacement.
Inside string replacement, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch

``regexReplaceAll`` panics if there is a problem and ``mustRegexReplaceAll`` returns an error to the template engine if there is a problem.

**Example:**

.. code-block:: html

   {{ regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W" }}

The above returns

.. code-block:: text

   -W-xxW-

----

.. _func-regexReplaceAllLiteral:
.. _func-mustRegexReplaceAllLiteral:

regexReplaceAllLiteral, mustRegexReplaceAllLiteral
==================================================

.. code-block:: go

   regexReplaceAllLiteral(re string, value string, replace string) string
   mustRegexReplaceAllLiteral(re string, value string, replace string) string, error

The ``regexReplaceAllLiteral`` function returns a copy of the input string, replacing matches of the Regexp with the replacement string replacement.
The replacement string is substituted directly, without using Expand

``regexReplaceAllLiteral`` panics if there is a problem and ``mustRegexReplaceAllLiteral`` returns an error to the template engine if there is a problem.

**Example:**

.. code-block:: html

   {{ regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}" }}

The above returns

.. code-block:: text

   -${1}-${1}-

----

.. _func-regexSplit:
.. _func-mustRegexSplit:

regexSplit, mustRegexSplit
==========================

.. code-block:: go

   regexSplit(re string, value string, count int) []string
   mustRegexSplit(re string, value string, count int) []string, error

The ``regexSplit`` function returns a copy of the input string, replacing matches of the Regexp with the replacement string replacement.
The replacement string is substituted directly, without using Expand

``regexSplit`` panics if there is a problem and ``mustRegexSplit`` returns an error to the template engine if there is a problem.

**Example:**

.. code-block:: html

   {{ regexSplit "z+" "pizza" -1 }}

The above returns

.. code-block:: text

   [pi a]

----

.. _func-regexQuoteMeta:

regexQuoteMeta
==============

.. code-block:: go

   regexQuoteMeta(value string) string

The ``regexQuoteMeta`` function returns a string that escapes all regular expression metacharacters inside the argument text.
The returned string is a regular expression matching the literal text.

**Example:**

.. code-block:: html

   {{ regexQuoteMeta "1.2.3" }}

The above returns

.. code-block:: text

   1\.2\.3




