#############
Configuration
#############

gostub uses a simple YAML file for configuration. The minimum configuration
requires specifying at least one service.

Configuration format
********************

.. code-block:: yaml
   :caption: config.yaml

   services:
     - name: default               # Seervice ID
       type: http                   # Type
       options:
         host: localhost            # host
         port: 8080                  # port
         root: /tmp/stubs            # stubs root directory

Parameters
**********

.. list-table::
   :header-rows: 1
   :widths: 15 10 30 45

   * - Parametr
     - Required
     - Description
     - Example
   * - ``name``
     - Y
     - Service ID. Used to form the path to directories with stubs.
     - ``default``, ``api``, ``auth``
   * - ``type``
     - Y
     - Service type. Currently only ``http`` is supported
     - ``http``
   * - ``host``
     - Y
     - Host to listen for incoming connections
     - ``localhost``, ``0.0.0.0``, ``127.0.0.1``
   * - ``port``
     - Y
     - Port for listening for HTTP requests
     - ``8080``, ``80``, ``3000``
   * - ``root``
     - Y
     - The base directory where template files are stored.
     - ``/tmp/stubs``, ``./stubs``, ``/var/gostub``

How the path to the stubs is formed
***********************************

gostub uses the following scheme to search for template files:

.. code-block:: text

   {root}/{name}/{request_path}

* ``root`` — value from configuration
* ``name`` — service ID
* ``request_path`` — path from the request URL (query is ignored)

Example
------

.. code-block:: yaml

   services:
     - name: users
       type: http
       options:
         host: localhost
         port: 8080
         root: /var/stubs

Request:

.. code-block:: text

   GET http://localhost:8080/api/v1/profile?id=123

The service will search for:

.. code-block:: text

   /var/stubs/users/api/v1/profile
