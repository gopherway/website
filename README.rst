$ createdb gopherway
$ psql gopherway < migration/db.sql

$ make install # Install dependencies packages.
$ make # Run server.