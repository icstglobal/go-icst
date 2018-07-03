package database

import (
	"database/sql"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
	conf "github.com/icstglobal/go-icst/config"
	"strconv"
)

var (
    // DBCon is the connection handle
    // for the database
    DBCon *sql.DB
)

// DB function
func DB() *sql.DB {
	mysqlConf := conf.Conf.Mysql

	port := strconv.Itoa(mysqlConf.Port)

	DBCon, err := sql.Open("mysql", mysqlConf.User+":"+mysqlConf.Password+"@tcp("+mysqlConf.Host+":"+port+")/"+mysqlConf.Db)
	if err != nil {
		panic(err)
	}
	err = DBCon.Ping()
	if err != nil {
		panic(err)
	}
	return DBCon
}
