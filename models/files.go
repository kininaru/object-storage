package models

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"time"
)

func SaveToLocal(file []byte) string {
	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("%d-%x", timestamp, md5.Sum(file))
	path := fmt.Sprintf("./files/%s", fileName)
	fileObj, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close()

	writer := bufio.NewWriter(fileObj)
	_, err = writer.Write(file)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	return fileName
}
