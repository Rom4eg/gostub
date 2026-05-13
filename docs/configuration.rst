#############
Configuration
#############

gostub uses a simple YAML file for configuration. The minimum configuration
requires specifying at least one service.

Configuration lookup order
**************************

By default, gostub looks for config.yaml in the following order:

#. User-specific ``~/.config/gostub/config.yaml``
#. System-wide ``/etc/gostub/config.yaml``
#. Current working directory ``./config.yaml``

The first found file is loaded. If no configuration file is found, gostub will exit with an error.

Configuration format
********************

.. code-block:: yaml
   :caption: config.yaml

   services:
     - name: default # Seervice ID
       type: api     # Type
       options:      # service specific settings

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
     - Service type. Currently only ``api`` is supported
     - ``api``
   * - ``options``
     - N
     - Service specific options
     - ``host: 0.0.0.0``

Service specific configuration parameters described in :doc:`services` section.

