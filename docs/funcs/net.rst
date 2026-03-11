Network Functions
*****************

Gostub provides network manipulation functions.

----

.. _func-getHostByName:

getHostByName
=============

.. code-block:: go

   getHostByName(host string) string

The ``getHostByName`` function receives a domain name and returns the IP address.

**Example:**

.. code-block:: html

   {{ getHostByName "example.com" }}

The above returns

.. code-block:: text

   12.345.678.90

