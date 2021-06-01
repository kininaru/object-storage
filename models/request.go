package models

import "encoding/base64"

func PutFile(path, data string) string {
	dist, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return"base64 error"
	}
	name := SaveToLocal(dist)
	AddToFileRecord(path, name)
	return ""
}
