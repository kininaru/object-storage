package models

import "xorm.io/core"

type FileRecord struct {
	Path   string `xorm:"varchar(100) notnull pk"`
	Bucket string `xorm:"varchar(100) notnull pk"`
	Name   string `xorm:"varchar(100) notnull"`
}

func AddToFileRecord(path, name, bucket string) bool {
	record := FileRecord{Path: path, Bucket: bucket}
	has, err := database.Get(&record)
	if err != nil {
		panic(err)
	}
	record.Name = name
	var lines int64
	if has {
		lines, err = database.ID(core.PK{record.Path, record.Bucket}).AllCols().Update(&record)
	} else {
		lines, err = database.Insert(&record)
	}
	if err != nil {
		panic(err)
	}
	return lines != 0
}

func GetFile(path, bucket string) string {
	record := FileRecord{Path: path, Bucket: bucket}
	has, _ := database.Get(&record)
	if !has {
		return ""
	}
	return record.Name
}
