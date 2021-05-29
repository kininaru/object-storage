package models

type FileRecord struct {
	Path string `xorm:"varchar(100) notnull pk"`
	Name string `xorm:"varchar(100) notnull"`
}

func AddToFileRecord(path, name string) bool {
	record := FileRecord{Path: path}
	has, err := database.Get(&record)
	if err != nil {
		panic(err)
	}
	record.Name = name
	var lines int64
	if has {
		lines, err = database.ID(record.Path).AllCols().Update(&record)
	} else {
		lines, err = database.Insert(&record)
	}
	if err != nil {
		panic(err)
	}
	return lines != 0
}

func GetFile(path string) string {
	record := FileRecord{Path: path}
	has, _ := database.Get(&record)
	if !has {
		return ""
	}
	return record.Name
}
