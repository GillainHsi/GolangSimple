package rtndata

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//ResponseData ... must return
func ResponseData(respData RespData) (rtnData string) {
	json, err := json.Marshal(respData)
	if err != nil {
		respData = RespData{
			Success: false,
			ErrCode: 500,
			ErrMsg:  "System Fail",
		}
	}

	return string(json)
}

//ConvStrToFloat32 ...
func ConvStrToFloat32(strValue string) (num float32, ok bool) {
	value, err := strconv.ParseFloat(strValue, 32)
	if err != nil {
		return 0.0, false
	}

	return float32(value), true
}

//ConvTimeToStr ...
// func ConvTimeToStr(t time.Time) string {
// 	return t.Format("2006-01-02 15:04:05")
// }

//ConvStrToTime ...
// func ConvStrToTime(strTime string) time.Time {
// 	date := "2006-01-02 15:04:05"
// 	t, _ := time.Parse(strTime, date)
// 	return t
// }

//CheckDataParameter ...
func CheckDataParameter(form url.Values) bool {
	_, ok := form["id"]
	if !ok {
		return false
	}

	lat, ok := form["lat"]
	if !ok {
		return false
	}

	long, ok := form["long"]
	if !ok {
		return false
	}

	_, okLat := ConvStrToFloat32(lat[0])
	if !okLat {
		return false
	}

	_, okLong := ConvStrToFloat32(long[0])
	if !okLong {
		return false
	}

	return true
}
