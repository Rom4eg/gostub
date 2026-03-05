Encoding Functions
******************

gostub has the following encoding and decoding functions.

----

.. _func-b64enc:

b64enc
======

.. code-block:: go

   b64enc(str string) string

The ``b64enc`` function encodes string to base64 representation.

**Example:**

.. code-block:: html

   {{ b64enc "hello world!" }}

The above returns

.. code-block:: text

   aGVsbG8gd29ybGQh

----

.. _func-b64dec:

b64dec
======

.. code-block:: go

   b64dec(str string) string

The ``b64dec`` function decodes base64 to string.

**Example:**

.. code-block:: html

   {{ b64dec "eyJtZXNzYWdlIjogIkhlbGxvIFdvcmxkISJ9" }}

The above returns

.. code-block:: text

   {"message": "Hello World!"}

----

.. _func-b32enc:

b32enc
======

.. code-block:: go

   b32enc(str string) string

The ``b32enc`` function encodes string to base32 representation.

**Example:**

.. code-block:: html

   {{ b32enc "Example word" }}

The above returns

.. code-block:: text

   IV4GC3LQNRSSA53POJSA====

----

.. _func-b32dec:

b32dec
======

.. code-block:: go

   b32dec(str string) string

The ``b32dec`` function decodes base32 to string.

**Example:**

.. code-block:: html

   {{ b32dec "O5QW43TBEBYGS6T2MEQD6===" }}

The above returns

.. code-block:: text

   wanna pizza ?

