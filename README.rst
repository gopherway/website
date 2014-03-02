$ cd $GOROOT/src/


$ createdb gopherway
$ psql gopherway < migration/db.sql

$ go run server.go models.go handlers.go