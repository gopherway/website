Create database and migration:

.. code:: bash
    $ createdb gopherway
    $ psql gopherway < migration/db.sql

Install and run:

.. code:: bash
    $ make install # Install dependencies packages.
    $ make # Run server.