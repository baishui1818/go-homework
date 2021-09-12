package main

import (
	"database/sql"
	"go-homework/week02/common"
	"imooc-product/datamodels"
	"strconv"
)

type UserMangerRepository struct {
	table     string
	mysqlConn *sql.DB
}

func (u *UserMangerRepository) Conn() error {
	if u.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		u.mysqlConn = mysql
	}

	if u.table == "" {
		u.table = "user"
	}
	return nil
}

func (u *UserMangerRepository) SelectByKey(userId int64) (user *datamodels.User, err error) {
	if errConn := u.Conn(); errConn != nil {
		return &datamodels.User{}, errConn
	}

	sqlStr := "select * from " + u.table + " where ID=" + strconv.FormatInt(userId, 10)
	row := u.mysqlConn.QueryRow(sqlStr)
	if row.Err() == sql.ErrNoRows {
		return &datamodels.User{}, err
	}

	return
}

func main() {

}
