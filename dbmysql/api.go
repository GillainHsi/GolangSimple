package dbmysql

import (
	"errors"
	"fmt"
	"strings"
)

func (ref *DBOperator) Exec(sqlCmd string, args ...interface{}) (int64, error) {
	if args != nil && len(args) != 0 {
		sqlCmd = fmt.Sprintf(sqlCmd, args...) // auto trans %v => input value
	}

	ref.sqlCmd = sqlCmd
	if len(sqlCmd) == 0 {
		ref.errMsg = "Error: empty sql cmd"
		return 0, errors.New(ref.errMsg)
	}

	if ref.conn == nil {
		ref.errMsg = "Error: SQL connection is nil"
		return 0, errors.New(ref.errMsg)
	}

	resp, err := ref.conn.Exec(sqlCmd)
	if err != nil {
		ref.errMsg = fmt.Sprintf("Excution fail! error: %v\n", err)
		return 0, err
	}

	if strings.HasPrefix(strings.ToLower(sqlCmd), "insert") {
		return resp.LastInsertId()
	}

	return resp.RowsAffected()
}

func (ref *DBOperator) Query(sqlCmd string, args ...interface{}) (*DBResponse, error) {
	if args != nil && len(args) != 0 {
		sqlCmd = fmt.Sprintf(sqlCmd, args...) // auto trans %v => input value
	}

	ref.sqlCmd = sqlCmd
	if len(sqlCmd) == 0 {
		ref.errMsg = "Error: empty sql cmd"
		return nil, errors.New(ref.errMsg)
	}

	if ref.conn == nil {
		ref.errMsg = "Error: SQL connection is nil"
		return nil, errors.New(ref.errMsg)
	}

	rtnRowsMap, err := innerQueryParsing(sqlCmd, ref.conn)
	if err != nil {
		ref.sqlCmd = fmt.Sprintf(sqlCmd, args...)
		ref.errMsg = fmt.Sprintf("Query error: %v", err)
		return nil, err
	}

	var response DBResponse
	response.RowsResponse = make([]map[string]string, 0) // map slice
	for _, valueMap := range rtnRowsMap {
		response.RowsResponse = append(response.RowsResponse, valueMap)
	}
	response.Length = uint32(len(response.RowsResponse))
	return &response, nil
}
