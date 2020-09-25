package main

import (
	rtndata "GolangSimple/data"
	"GolangSimple/dbmysql"
	"fmt"
	"net/http"
)

var db *dbmysql.DBOperator
var errDB error

func init() {
	db, _ = dbmysql.InitDB("root", "12345678", "127.0.0.1:3306", "develop")
}

func main() {

	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/data", handleReq)
	server.ListenAndServe()

}

func handleReq(w http.ResponseWriter, r *http.Request) {

	if db == nil {
		db, errDB = dbmysql.InitDB("root", "12345678", "127.0.0.1:3306", "develop")
		if errDB != nil {
			respdata := rtndata.RespData{
				Success: false,
				ErrCode: 500,
				ErrMsg:  "DB Fail",
			}

			w.Write([]byte(rtndata.ResponseData(respdata)))
			return
		}
		fmt.Println("db recover ...")
	}

	switch r.Method {
	case "GET":
		rtndata.HandleDataGet(w, r, db)
	case "POST":
		rtndata.HandleDataPost(w, r, db)
	default:
		rtndata.HandleDataUnknowMethod(w, r, db)
	}
}
