package models

import (
	"github.com/astaxie/beego/orm"
//	"database/sql"
	"fmt"
//	"github.com/arnehormann/sqlinternals/mysqlinternals"
//	"strconv"
)

func NewOrm(multiOrm []orm.Ormer, db ...string) (o orm.Ormer) {
	if len(multiOrm) == 0 {
		o = orm.NewOrm()
		if len(db) == 1 {
			o.Using(db[0])
		}
	} else if len(multiOrm) == 1 {
		o = multiOrm[0]
	} else {
		panic("only one Ormer")
	}
	return
}

func NewDBOrm(dbName string) (o orm.Ormer) {
	o = orm.NewOrm()
	o.Using(dbName)
	return
}


// 拼接update的sql语句
func UpdateSql(tableName string, data map[string]interface{}, condition string) (values []interface{}, sql string) {
	field := ""
	for k, v := range data {
		if field == "" {
			field = fmt.Sprintf("`%v`", k) + "=?"
		} else {
			field = field + "," + fmt.Sprintf("`%v`", k) + "=?"
		}
		values = append(values, v)
	}
	sql = "UPDATE " + tableName + " SET " + field + " WHERE " + condition
	return
}

func DeleteSql(tableName string, condition string) string {
	return "DELETE FROM " + tableName + " WHERE " + condition
}

func InsertSql(tableName string, data map[string]interface{}) (values []interface{}, sql string) {
	field := ""
	palceholder := ""
	for k, v := range data {
		if field != "" {
			field = field + "," + fmt.Sprintf("`%v`", k)
			palceholder = palceholder + "," + "?"
		} else {
			field = fmt.Sprintf("`%v`", k)
			palceholder = "?"
		}
		values = append(values, v)
	}
	sql = "INSERT INTO " + tableName + "(" + field + ") VALUES(" + palceholder + ")"
	return
}

func InsertOrUpdateSql(tableName string, data map[string]interface{}) (values []interface{}, sql string) {
	values, sql = InsertSql(tableName, data)
	sql = sql + " ON DUPLICATE KEY UPDATE "

	field := ""
	for k, v := range data {
		if field != "" {
			field = field + ", " + fmt.Sprintf("`%v`", k) + "=?"
		} else {
			field = fmt.Sprintf("`%v`", k) + "=?"
		}
		values = append(values, v)
	}
	sql = sql + field
	return
}

/*
//查询一行数据装到map[string]interface{}里
func QueryRowToMap(db *sql.DB, sqlStr string) (map[string]interface{}, error) {

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := mysqlinternals.Columns(rows)
	if err != nil {
		return nil, err
	}
	columnsLength := len(columns)
	values := make([]sql.RawBytes, columnsLength)

	scanArgs := make([]interface{}, columnsLength)

	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		for i, value := range values {
			if value == nil {
				record[columns[i].Name()] = "NULL"
			} else {
				record[columns[i].Name()] = bytes2RealType(value, columns[i])
			}
		}
	}
	return record, nil
}

func bytes2RealType(src []byte, column mysqlinternals.Column) interface{} {
	srcStr := string(src)
	var result interface{}
	switch column.MysqlType() {
	case "TINYINT":
		fallthrough
	case "SMALLINT":
		fallthrough
	case "INT":
		result, _ = strconv.ParseInt(srcStr, 10, 64)
	case "BIGINT":
		if column.IsUnsigned() {
			result, _ = strconv.ParseUint(srcStr, 10, 64)
		} else {
			result, _ = strconv.ParseInt(srcStr, 10, 64)
		}
	case "CHAR":
		fallthrough
	case "VARCHAR":
		fallthrough
	case "BLOB":
		fallthrough
	case "TIMESTAMP":
		fallthrough
	case "DATE":
		fallthrough
	case "DATETIME":
		fallthrough
	case "TIME":
		result = srcStr
	case "FLOAT":
		fallthrough
	case "DOUBLE":
		fallthrough
	case "DECIMAL":
		result, _ = strconv.ParseFloat(srcStr, 32)
	default:
		result = nil
	}
	return result
}
*/
