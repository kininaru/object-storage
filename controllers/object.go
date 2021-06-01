package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"object-storage/models"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Command() {
	if !c.CheckPostBody("command", "data", "id", "secret", "path") {
		return
	}

	if !models.CheckUser(c.Ctx.Request.Form.Get("id"), c.Ctx.Request.Form.Get("secret")) {
		c.Response(1, "Secret error")
		return
	}

	data := ""
	index := strings.Index(c.Ctx.Request.Form.Get("data"), ",")
	if index >= 0 {
		data = c.Ctx.Request.Form.Get("data")[index+1:]
	}

	switch c.Ctx.Request.Form.Get("command") {
	case "put":
		dist, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			fmt.Println(err)
			c.Response(2, "base64 error")
			return
		}
		name := models.SaveToLocal(dist)
		models.AddToFileRecord(c.Ctx.Request.Form.Get("path"), name)
	}

	c.Response(0)
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
