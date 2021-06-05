package models

import (
	"encoding/base64"
	"strings"
)

func PutFile(path, data, bucket string) string {
	index := strings.Index(data, ",")
	if index >= 0 {
		data = data[index+1:]
	}
	dist, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "base64 error"
	}
	name := SaveToLocal(dist)
	AddToFileRecord(path, name, bucket)
	return ""
}
