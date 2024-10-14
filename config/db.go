package config

import (
	"dungeonSnackBE/helper/atdb"
)

var MongoString string = "mongodb+srv://ukasyzzam:MongoDB12@nano.c3jog.mongodb.net/"

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dsdatabase",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)
