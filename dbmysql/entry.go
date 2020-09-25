package dbmysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//DBOperator ...
type DBOperator struct {
	conn   *sql.DB
	dbname string
	sqlCmd string
	isConn bool
	errMsg string
}

//DBResponse ...
type DBResponse struct {
	RowsResponse []map[string]string
	Length       uint32
}

//InitDB ...
func InitDB(user, pass, host, dbname string) (*DBOperator, error) {
	sqlDB, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+")/"+dbname+"?charset=utf8")
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	dbOperator := DBOperator{
		conn:   sqlDB,
		dbname: dbname,
		isConn: true,
	}

	return &dbOperator, nil
}
