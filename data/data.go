package rtndata

//Data ...  orm
type Data struct {
	ID       string `json:"ID"`
	Location struct {
		Lat  float32 `json:"Lat"`
		Long float32 `json:"Long"`
	} `json:"Location"`
	DateAdded interface{} `json:"DateAdded"`
}

//RespData ... return json
type RespData struct {
	Success bool   `json:"success"`
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`

	RtnData []Data `json:"data"`
}
