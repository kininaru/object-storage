package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"object-storage/models"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Command() {
	command := c.Ctx.Request.Form.Get("command")
	data := c.Ctx.Request.Form.Get("data")
	id := c.Ctx.Request.Form.Get("id")
	secret := c.Ctx.Request.Form.Get("secret")
	ext := c.Ctx.Request.Form.Get("extension")
	path := c.Ctx.Request.Form.Get("path")

	if models.CheckUser(id, secret) {
		c.Response(1, "Secret error")
		return
	}

	index := strings.Index(data, ",")
	if index >= 0 {
		data = data[index+1:]
	}

	switch command {
	case "put":
		dist, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			fmt.Println(err)
			c.Response(2, "base64 error")
			return
		}
		name := models.SaveToLocal(dist, ext)
		models.AddToFileRecord(path, name)
	}

	logText := fmt.Sprintf("Command received: %s", command)
	c.Response(0)
	logs.Info(logText)
}

func (c *ApiController) GetFile() {
	path := c.Ctx.Input.Param(":splat")

	name := models.GetFile(path)
	name = fmt.Sprintf("./files/%s", name)

	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		c.Response(1)
		return
	}
	defer file.Close()

	c.Ctx.ResponseWriter.WriteHeader(200)
	fileByte, err := ioutil.ReadAll(file)
	_, err = c.Ctx.ResponseWriter.Write(fileByte)
	if err != nil {
		c.Response(2)
		return
	}
}
