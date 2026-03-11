Cryptographic Functions
************************************

Gostub provides a couple of advanced cryptographic functions.

----

.. _struct-certificate:

Certificate object
==================

:Cert: PEM-encoded certificate
:Key: PEM-encoded private key

----

.. _func-sha1sum:

sha1sum
=======

.. code-block:: go

   sha1sum(val string) string

The ``sha1sum`` function computes its SHA1 digest.

**Example:**

.. code-block:: html

   {{ sha1sum "Hello world!!!" }}

The above returns

.. code-block:: text

   6555aa9d245f6dc2b57aa13366cc6c6fcccab6ad

----

.. _func-sha256sum:

sha256sum
=========

.. code-block:: go

   sha256sum(val string) string

The ``sha256sum`` function computes its SHA256 digest.

**Example:**

.. code-block:: html

   {{ sha256sum "Hello world!!!" }}

The above returns

.. code-block:: text

   4354dfda70c8f0d3991b9de3d56dcb6e9f2fc6c0316d235b63afeb388471ada4

----

.. _func-sha512sum:

sha512sum
=========

.. code-block:: go

   sha512sum(val string) string

The ``sha512sum`` function computes its SHA512 digest.

**Example:**

.. code-block:: html

   {{ sha512sum "Hello world!!!" }}

The above returns

.. code-block:: text

   833ec109e1c1b84bdfe3e05aa00b395bb47f913dc6cc4288066edc79e968f7212cec0bf652b8b3c89405a30223c71462bc03e77cb038781ae01d6ba2e04bfe2b

----

.. _func-adler32sum:

adler32sum
==========

.. code-block:: go

   adler32sum(val string) string

The ``adler32sum`` function computes its Adler-32 digest.

**Example:**

.. code-block:: html

   {{ adler32sum "Hello world!!!" }}

The above returns

.. code-block:: text

   640156832

----

.. _func-bcrypt:

bcrypt
======

.. code-block:: go

   bcrypt(val string) string

The ``bcrypt`` function generates its bcrypt hash.

**Example:**

.. code-block:: html

   {{ bcrypt "Hello world!!!" }}

The above returns

.. code-block:: text

   $2a$10$838K8mNOx.WBF20/4P.mheEZFEslen1.8DIKGKixLftVGSGCkuL6q

----

.. _func-htpasswd:

htpasswd
========

.. code-block:: go

   htpasswd(username string, password string) string

The ``htpasswd`` function generates an Apache .htpasswd auth config row e.g., username:password.

**Example:**

.. code-block:: html

   {{ htpasswd "myUser" "SeCrEt!1" }}

The above returns

.. code-block:: text

   myUser:$2a$10$rsu.w7gNQGD4pA.7sA/5juR9G39fKAMsRQ.WEBM17ek0i5NqcDw1G

----

.. _func-randBytes:

randBytes
=========

.. code-block:: go

   randBytes(len int) string

The ``randBytes`` function generates a cryptographically secure random sequence of ``len`` bytes.
The sequence is returned as a base64 encoded string.

**Example:**

.. code-block:: html

   {{ randBytes 32 }}

The above returns

.. code-block:: text

   enm6C5P/nYWhRm8V/Zks+2UiBn+vePKY+qSpxP9kZVg=

----

.. _func-derivePassword:

derivePassword
==============

.. code-block:: go

   derivePassword(counter int, type string, password string, user string, site string) string

The ``derivePassword`` function generates a `master password <https://spectre.app/spectre-algorithm.pdf>`__.
**Example:**

.. code-block:: html

   {{ derivePassword 1 "long" "password" "user" "example.com" }}

The above returns

.. code-block:: text

   ZedaFaxcZaso9*

----

.. _func-genPrivateKey:

genPrivateKey
=============

.. code-block:: go

   genPrivateKey(typ string) string

The ``genPrivateKey`` function generates a new private key encoded into a PEM block.

It takes one of the values for its first param:

:ecdsa: Generate an elliptic curve DSA key (P256)
:dsa: Generate a DSA key (L2048N256)
:rsa: Generate an RSA 4096 key
:ed25519: Generate an Ed25519 key

**Example:**

.. code-block:: html

   {{ genPrivateKey "ecdsa" }}

The above returns

.. code-block:: text

   -----BEGIN EC PRIVATE KEY-----
   MHcCAQEEINQ0wXP/+rRJNjqVfiVt6sPzhk007IVDq67Gc3DkcrFFoAoGCCqGSM49
   AwEHoUQDQgAEzj2dRUPnK9fCtlV7boBDFj9o7ga0bt17AaUOy7FPDzeW4JUi1eoD
   XNgsxnN0/b33swK7WC2zRFaCXI5VLswweA==
   -----END EC PRIVATE KEY-----

----

.. _func-buildCustomCert:

buildCustomCert
===============

.. code-block:: go

   buildCustomCert(cert string, key string) certificate, error

The ``buildCustomCert`` function allows customizing the certificate.

It takes the following string parameters:

* A base64 encoded PEM format certificate
* A base64 encoded PEM format private key

It returns a :ref:`struct-certificate`

**Example:**

.. code-block:: html

   {{ $ca := genCA "foo-ca" 365 }}
   {{ $cert := genSignedCert "foo.com" (list "10.0.0.1" "10.0.0.2") (list "bar.com" "bat.com") 365 $ca }}

   {{ $key := genPrivateKey "rsa" }}

   {{ $crt := buildCustomCert (b64enc $cert.Cert) (b64enc $key) }}
   {{ toJson $crt }}

The above returns

.. code-block:: text

   {"Cert":"-----BEGIN CERTIFICATE-----\n .......... -----END CERTIFICATE-----\n","Key":"-----BEGIN RSA PRIVATE KEY-----\n .......... \n-----END RSA PRIVATE KEY-----\n"}

----

.. _func-genCA:

genCA
=====

.. code-block:: go

   genCA(cn string, days int) certificate, error

The ``genCA`` function generates a new, self-signed x509 certificate authority using a 2048-bit RSA private key.

It takes the following parameters:

* Subject’s common name (cn)
* Cert validity duration in days

It returns a :ref:`struct-certificate`

**Example:**

.. code-block:: html

   {{ $ca := genCA "foo-ca" 365 }}
   {{ toJson $ca }}

The above returns

.. code-block:: text

   {"Cert":"-----BEGIN CERTIFICATE-----\n .......... -----END CERTIFICATE-----\n","Key":"-----BEGIN RSA PRIVATE KEY-----\n .......... \n-----END RSA PRIVATE KEY-----\n"}

----

.. _func-genCAWithKey:

genCAWithKey
============

.. code-block:: go

   genCAWithKey(cn string, days int, pk string) certificate, error

The ``genCAWithKey`` function generates a new, self-signed x509 certificate authority using a given private key.

It takes the following parameters:

* Subject’s common name (cn)
* Cert validity duration in days
* Private key (PEM-encoded); DSA keys are not supported

It returns a :ref:`struct-certificate`

**Example:**

.. code-block:: html

   {{ $ca := genCAWithKey "foo-ca" 365 (genPrivateKey "rsa") }}
   {{ toJson $ca }}

The above returns

.. code-block:: text

   {"Cert":"-----BEGIN CERTIFICATE-----\n .......... -----END CERTIFICATE-----\n","Key":"-----BEGIN RSA PRIVATE KEY-----\n .......... \n-----END RSA PRIVATE KEY-----\n"}

----

.. _func-genSelfSignedCert:

genSelfSignedCert
=================

.. code-block:: go

   genSelfSignedCert(cn string, ips []string, names []string, days int) certificate, error

The ``genSelfSignedCert`` function generates a new, self-signed x509 certificate using a 2048-bit RSA private key.

It takes the following parameters:

* Subject’s common name (cn)
* Optional list of IPs; may be nil
* Optional list of alternate DNS names; may be nil
* Cert validity duration in days

It returns a :ref:`struct-certificate`

**Example:**

.. code-block:: html

   {{ $ca := genSelfSignedCert "foo.com" (list "10.0.0.1" "10.0.0.2") (list "bar.com" "baz.com") 365 }}
   {{ toJson $ca }}

The above returns

.. code-block:: text

   {"Cert":"-----BEGIN CERTIFICATE-----\n .......... -----END CERTIFICATE-----\n","Key":"-----BEGIN RSA PRIVATE KEY-----\n .......... \n-----END RSA PRIVATE KEY-----\n"}

----

.. _func-genSelfSignedCertWithKey:

genSelfSignedCertWithKey
========================

.. code-block:: go

   genSelfSignedCertWithKey(cn string, ips []string, names []string, days int, pk string) certificate, error

The ``genSelfSignedCertWithKey`` function generates a new, self-signed x509 certificate using a given private key.

It takes the following parameters:

* Subject’s common name (cn)
* Optional list of IPs; may be nil
* Optional list of alternate DNS names; may be nil
* Cert validity duration in days
* Private key (PEM-encoded); DSA keys are not supported

It returns a :ref:`struct-certificate`

**Example:**

.. code-block:: html

   {{ $ca := genSelfSignedCertWithKey "foo.com" (list "10.0.0.1" "10.0.0.2") (list "bar.com" "baz.com") 365 (genPrivateKey "ecdsa")}}
   {{ toJson $ca }}

The above returns

.. code-block:: text

   {"Cert":"-----BEGIN CERTIFICATE-----\n .......... -----END CERTIFICATE-----\n","Key":"-----BEGIN RSA PRIVATE KEY-----\n .......... \n-----END RSA PRIVATE KEY-----\n"}

----

.. _func-genSignedCert:

genSignedCert
=============

.. code-block:: go

   genSignedCert(cn string, ips []string, names []string, days int, ca certificate) certificate, error

The ``genSignedCert`` function generates a new, x509 certificate signed by the specified CA, using a 2048-bit RSA private key.

It takes the following parameters:

* Subject’s common name (cn)
* Optional list of IPs; may be nil
* Optional list of alternate DNS names; may be nil
* Cert validity duration in days
* CA (see :ref:`func-genCA`)

It returns a :ref:`struct-certificate`

**Example:**

.. code-block:: html

   {{ $ca := genCA "foo-ca" 365 }}
   {{ $cert := genSignedCert "foo.com" (list "10.0.0.1" "10.0.0.2") (list "bar.com" "bat.com") 365 $ca }}
   {{ toJson $ca }}

The above returns

.. code-block:: text

   {"Cert":"-----BEGIN CERTIFICATE-----\n .......... -----END CERTIFICATE-----\n","Key":"-----BEGIN RSA PRIVATE KEY-----\n .......... \n-----END RSA PRIVATE KEY-----\n"}

----

.. _func-genSignedCertWithKey:

genSignedCertWithKey
====================

.. code-block:: go

   genSignedCertWithKey(cn string, ips []string, names []string, days int, ca certificate, pk string) certificate, error

The ``genSignedCertWithKey`` function generates a new, x509 certificate signed by the specified CA, using a given private key.

It takes the following parameters:

* Subject’s common name (cn)
* Optional list of IPs; may be nil
* Optional list of alternate DNS names; may be nil
* Cert validity duration in days
* CA (see :ref:`func-genCA`)
* Private key (PEM-encoded); DSA keys are not supported

It returns a :ref:`struct-certificate`

**Example:**

.. code-block:: html

   {{ $ca := genCA "foo-ca" 365 }}
   {{ $cert := genSignedCert "foo.com" (list "10.0.0.1" "10.0.0.2") (list "bar.com" "bat.com") 365 $ca (genPrivateKey "ed25519") }}
   {{ toJson $ca }}

The above returns

.. code-block:: text

   {"Cert":"-----BEGIN CERTIFICATE-----\n .......... -----END CERTIFICATE-----\n","Key":"-----BEGIN RSA PRIVATE KEY-----\n .......... \n-----END RSA PRIVATE KEY-----\n"}

----

.. _func-encryptAES:

encryptAES
==========

.. code-block:: go

   encryptAES(password string, plaintext string) string, error

The ``encryptAES`` function encrypts text with AES-256 CBC and returns a base64 encoded string.

**Example:**

.. code-block:: html

   {{ encryptAES "SeCrEt!1" "Hello world!" }}

The above returns

.. code-block:: text

   qSXClQAznrZDigJJbPaGCxG5KpWWYFe54QrY7DDFRLM=

----

.. _func-decryptAES:

decryptAES
==========

.. code-block:: go

   decryptAES(password string, ciphertext string) string, error

The ``decryptAES`` function receives a base64 string encoded by the AES-256 CBC algorithm and returns the decoded text.

**Example:**

.. code-block:: html

   {{ "qSXClQAznrZDigJJbPaGCxG5KpWWYFe54QrY7DDFRLM=" | decryptAES "SeCrEt!1" }}

The above returns

.. code-block:: text

   Hello world!

