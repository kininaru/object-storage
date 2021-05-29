package controllers

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *ApiController) Response(code int, args ...interface{}) {
	var msg string
	var data interface{}
	switch len(args) {
	case 2:
		data = args[1]
		fallthrough
	case 1:
		msg, _ = args[0].(string)
	}
	c.Data["json"] = Response{code, msg, data}
	err := c.ServeJSON()
	if err != nil {
		panic(err)
	}
}
