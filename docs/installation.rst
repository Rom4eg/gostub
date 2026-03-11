############
Installation
############

gostub can be installed in several ways: via Go, from sources or using Docker.


Installation via Go
*******************

If you have Go version 1.16 or higher installed:

.. code-block:: bash

   go install github.com/Rom4eg/gostub/cmd/gostub@latest

After installation, the ``gostub`` binary will be available in ``$GOPATH/bin``
(usually ``~/go/bin``).

Building from sources
*********************

1. Clone the repository:

   .. code-block:: bash

      git clone https://github.com/rom4eg/gostub.git
      cd gostub

2. Build the project using Make:

   .. code-block:: bash

      make build

   The executable will be created in the ``bin/`` directory.

Using Docker
************

The image is available at `Docker Hub <https://hub.docker.com/r/rom4eg196/gostub>`__:

.. code-block:: bash

   # Download image
   docker pull rom4eg196/gostub:latest

   # Run container
   docker run --rm -p 8080:8080 rom4eg196/gostub

   # Or by mounting a directory with stubs
   docker run --rm \
     -p 8080:8080 \
     -v $(pwd)/config.yaml:/config.yaml \
     -v $(pwd)/stubs:/stubs \
     rom4eg196/gostub


