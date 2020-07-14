package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var (
	UICDB   *sqlx.DB
	MYSQLDB *sqlx.DB
)

const (
	TBL_SIDECAR_LIST = "sidecar_list"
	TBL_KUBE         = "kube"
)

func USE_MYSQL_DB() *sqlx.DB {
	return MYSQLDB
}

func USE_UIC_DB() *sqlx.DB {
	return UICDB
}

func connectDatabase(host, user, password, dbname string, port int) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local&parseTime=true", user, password, host, port, dbname))
}

func ConfigureUicDatabase(host string, port int, user, password, dbname string) error {
	var err error
	UICDB, err = connectDatabase(host, user, password, dbname, port)
	return err
}

func ConfigureMysqlDatabase(host string, port int, user, password, dbname string) error {
	var err error
	MYSQLDB, err = connectDatabase(host, user, password, dbname, port)
	return err
}
