package controllers

import (
	"fmt"
	"io/ioutil"
	"os"

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

	var respMsg string
	switch c.Ctx.Request.Form.Get("command") {
	case "put":
		path := c.Ctx.Request.Form.Get("path")
		data := c.Ctx.Request.Form.Get("data")
		respMsg = models.PutFile(path, data)
	default:
		respMsg = "Unknown command type."
	}

	if len(respMsg) != 0 {
		c.Response(2, respMsg)
		return
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
