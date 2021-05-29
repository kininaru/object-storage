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

func (c *ApiController) HasParam(key string) (string, bool) {
	param := c.Ctx.Request.Form.Get(key)
	if len(param) == 0 {
		c.Response(-1, "Miss param: " + key)
		return param, false
	}
	return param, true
}

func (c *ApiController) RequireParams(keys ...string) (map[string]string, bool) {
	ret := make(map[string]string)
	var value string
	var has bool
	for _, key := range keys {
		if value, has = c.HasParam(key); !has {
			return ret, false
		}
		ret[key] = value
	}
	return ret, true
}
