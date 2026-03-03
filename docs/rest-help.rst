=====================
reStructuredText help
=====================


*italic text*

**bold text**

``inline code``

* non ordered
* list

  * nested list
  * items

1. ordered
2. list


#. auto ordered
#. list

Definition
  some description
  of the definition

next term
   Description.

Quoted paragraph

| These lines are
| broken exactly like in
|
| the source file.

Literal block::

    block content


Doctest block

>>> 1+1

Table

+------------------------+------------+----------+----------+
| Header row, column 1   | Header 2   | Header 3 | Header 4 |
| (header rows optional) |            |          |          |
+========================+============+==========+==========+
| body row 1, column 1   | column 2   | column 3 | column 4 |
+------------------------+------------+----------+----------+
| body row 2             | ...        | ...      |          |
+------------------------+------------+----------+----------+

Simple Table

=====  =====  =======
A      B      A and B
=====  =====  =======
False  False  False
True   False  False
False  True   False
True   True   True
=====  =====  =======

External link `a link`_

.. _a link: https://github.com/Rom4eg/gostub

`Inline link <https://github.com/Rom4eg/gostub>`__

####
Part
####

*******
Chapter
*******

=======
Section
=======


    :param my_arg: The first of my arguments.
    :param my_other_arg: The second of my arguments.

    :returns: A message (just for me, of course).
    :orphan:
