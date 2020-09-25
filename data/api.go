package rtndata

import (
	"GolangSimple/dbmysql"
	"fmt"
	"time"
)

// 3. 有二個endpoint
// - GET /data → 回傳所有database裡面的data
// - POST /data → 把上傳的data存入database

//InsertData ... must return anyway
func InsertData(data Data, db *dbmysql.DBOperator) (respData string) {

	sqlCmd := "INSERT INTO `data` (`id`, `lat`, `long`, `dateadded`) VALUES ('%v', %v, %v, '%v');"

	respdata := RespData{
		Success: true,
		ErrCode: -1,
		ErrMsg:  "",
	}

	_, err := db.Exec(sqlCmd, data.ID, data.Location.Lat, data.Location.Long, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Printf("InsertData error: %v\n", err)
		respdata = RespData{
			Success: false,
			ErrCode: 100,
			ErrMsg:  "DB Error",
		}
	}

	return ResponseData(respdata)
}

//QueryDataRow ... must return anyway
func QueryDataRow(db *dbmysql.DBOperator) (respData string) {

	sqlCmd := "SELECT dt.id AS 'id', dt.lat AS 'lat', dt.`long` AS 'long', dt.dateadded AS 'dateadded' FROM `data` dt;"
	resp, err := db.Query(sqlCmd)

	if err != nil {
		respdata := RespData{
			Success: false,
			ErrCode: 100,
			ErrMsg:  "DB Error",
			RtnData: make([]Data, 0, 0),
		}
		return ResponseData(respdata)
	}

	if resp.Length == 0 {
		respdata := RespData{
			Success: false,
			ErrCode: 200,
			ErrMsg:  "Empty Data",
			RtnData: make([]Data, 0, 0),
		}

		return ResponseData(respdata)
	}

	dataList := make([]Data, 0, resp.Length)
	for _, value := range resp.RowsResponse {
		lat, _ := ConvStrToFloat32(value["lat"])
		long, _ := ConvStrToFloat32(value["long"])
		data := Data{
			ID: value["id"],
			Location: struct {
				Lat  float32 `json:"Lat"`
				Long float32 `json:"Long"`
			}{
				Lat:  lat,
				Long: long,
			},
			DateAdded: value["dateadded"],
		}
		dataList = append(dataList, data)
	}

	respdata := RespData{
		Success: true,
		ErrCode: -1,
		ErrMsg:  "",
		RtnData: dataList,
	}

	return ResponseData(respdata)
}
