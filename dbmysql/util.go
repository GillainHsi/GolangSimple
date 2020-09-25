package dbmysql

import (
	"database/sql"
)

func innerQueryParsing(sqlCmd string, db *sql.DB) ([]map[string]string, error) {
	rows, err := db.Query(sqlCmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return queryRtnParsing(rows), nil
}

func queryRtnParsing(rows *sql.Rows) []map[string]string {
	columns, _ := rows.Columns() //return []string
	columnsCnt := len(columns)
	columnsPtr := make([]interface{}, columnsCnt)
	columnsVal := make([]interface{}, columnsCnt)

	for index := range columnsVal {
		columnsPtr[index] = &columnsVal[index]
	}

	rtnRowsMap := make([]map[string]string, 0) //slice input map
	for rows.Next() {
		oneRowMap := make(map[string]string)
		rows.Scan(columnsPtr...) //for columns
		for index, column := range columnsVal {
			if column == nil {
				continue
			}
			oneRowMap[columns[index]] = string(column.([]byte))
		}
		rtnRowsMap = append(rtnRowsMap, oneRowMap)
	}

	return rtnRowsMap
}
