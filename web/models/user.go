package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id int64 `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreateTime int64 `json:"create_time,omitempty"`
}

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
func UserByEmail(email string ,multiOrm ...orm.Ormer)(){
	sql := fmt.Sprintf(`
	SELECT
		password
	FROM
		%v
	WHERE
		email = %v
	`,TableUser,email)
	o := NewOrm(multiOrm,DBResitc)
	o.Raw(sql)
}