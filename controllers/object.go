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
	var params map[string]string
	ok := false

	if params, ok = c.RequireParams("command", "data", "id", "secret", "path"); !ok {
		return
	}

	if !models.CheckUser(params["id"], params["secret"]) {
		c.Response(1, "Secret error")
		return
	}

	data := ""
	index := strings.Index(params["data"], ",")
	if index >= 0 {
		data = params["data"][index+1:]
	}

	switch params["command"] {
	case "put":
		dist, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			fmt.Println(err)
			c.Response(2, "base64 error")
			return
		}
		name := models.SaveToLocal(dist)
		models.AddToFileRecord(params["path"], name)
	}

	logText := fmt.Sprintf("Command received: %s", params["command"])
	c.Response(0)
	logs.Info(logText)
}

func (c *ApiController) GetFile() {
	path := c.Ctx.Input.Param(":splat")
	if len(path) == 0 {
		c.Response(-1, "No such file.")
		return
	}

	name := models.GetFile(path)
	if len(name) == 0 {
		c.Response(-1, "No such file.")
		return
	}
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
