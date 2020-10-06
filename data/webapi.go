package rtndata

import (
	"GolangSimple/dbmysql"
	"fmt"
	"net/http"
)

func HandleDataPost(w http.ResponseWriter, r *http.Request, db *dbmysql.DBOperator) {

	r.ParseForm()
	ok := CheckDataParameter(r.PostForm)
	if !ok {
		respdata := RespData{
			Success: false,
			ErrCode: 555,
			ErrMsg:  "Parameter check fail",
		}
		w.Write([]byte(ResponseData(respdata)))
		return
	}

	id := string(r.PostForm["id"][0])
	lat, _ := ConvStrToFloat32(r.PostForm["lat"][0])
	long, _ := ConvStrToFloat32(r.PostForm["long"][0])
	data := Data{
		ID: id,
		Location: struct {
			Lat  float32 `json:"Lat"`
			Long float32 `json:"Long"`
		}{
			Lat:  lat,
			Long: long,
		},
	}

	w.Write([]byte(InsertData(data, db)))

	return
}

func HandleDataGet(w http.ResponseWriter, r *http.Request, db *dbmysql.DBOperator) {
	w.Write([]byte(QueryDataRow(db)))
}

func HandleDataUnknowMethod(w http.ResponseWriter, r *http.Request, db *dbmysql.DBOperator) {
	respdata := RespData{
		Success: false,
		ErrCode: 999,
		ErrMsg:  fmt.Sprintf("Method %v Not Support.", r.Method),
	}

	w.Write([]byte(ResponseData(respdata)))
}
