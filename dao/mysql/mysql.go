package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Init(user, pass, dbname, host string, port int) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=true",
		user, pass, host, port, dbname)
	db, err = sqlx.Connect("mysql", dsn)
	return
}
