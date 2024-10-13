package config

import (
	"dungeonSnackBE/helper/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dsdatabase",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)
