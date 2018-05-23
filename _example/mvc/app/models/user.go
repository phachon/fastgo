package models

import (
	"github.com/snail007/go-activerecord/mysql"
)

var User = UserModel{}

const Table_User_Name = "user"

type UserModel struct {
	BaseModel
}

// get user by user_id
func (u *UserModel) GetUserByUserId(userId string) (user map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_User_Name).Where(map[string]interface{}{
		"user_id":  userId,
	}))
	if err != nil {
		return
	}
	user = rs.Row()
	return
}