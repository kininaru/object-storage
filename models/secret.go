package models

type User struct {
	Id string `xorm:"varchar(100) notnull pk"`
	Secret string `xorm:"varchar(100) notnull"`
}

func CheckUser(key, secret string) bool {
	user := User{key, secret}
	has, _ := database.Get(&user)
	return has
}

func UpdateUser(id, secret string) bool {
	user := User{Id: id}
	has, _ := database.Get(&user)
	user.Secret = secret
	var lines int64
	if has {
		lines, _ = database.ID(user.Id).AllCols().Update(&user)
	} else {
		lines, _ = database.Insert(&user)
	}
	return lines != 0
}
