####
Home
####

**gostub** is a lightweight service for simulating responses to various APIs.
During development and debugging, stubs for external services are quickly created
when the real service is unavailable or difficult to use.

The idea is simple: the requested URL is converted into a path to a file on disk,
and the file's contents are processed as a `Go template <https://golang.org/pkg/text/template/>`_,
which allows you to create dynamic responses of any complexity.

Check out the :doc:`usage` section for further information.

.. toctree::
   :maxdepth: 2
   :name: mastertoc

   installation
   configuration
   creating-stubs
   functions
