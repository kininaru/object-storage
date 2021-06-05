package models

type User struct {
	Id string `xorm:"varchar(100) notnull pk"`
	Bucket string `xorm:"varchar(100)"`
	Secret string `xorm:"varchar(100) notnull"`
}

func CheckUser(id, secret, bucket string) bool {
	user := User{Id: id, Secret: secret}
	has, _ := database.Get(&user)
	if !has {
		return false
	}
	if bucket == user.Bucket {
		return true
	}
	if user.Id == "root" {
		return true
	}
	return false
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
