package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func NewUser(data map[string]interface{}, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm, DBResitc)
	values, sql := InsertSql(TableUser, data)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateUser(data map[string]interface{}, uid int64, multiOrm ...orm.Ormer) (int64, error) {
	o := NewOrm(multiOrm, DBResitc)
	condition := fmt.Sprintf(`id = %v`, uid)
	values, sql := UpdateSql(TableUser, data, condition)
	result, err := o.Raw(sql, values).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}