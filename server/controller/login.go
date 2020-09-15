package controller

import (
	"encoding/json"
	"errors"
	m "ims/model"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Res(c, -1, nil, err)
		return
	}

	args := make(map[string]string)
	err = json.Unmarshal(body, &args)
	if err != nil {
		Res(c, -1, nil, err)
		return
	}

	user, ok := args["user"]
	if !ok || user == "" {
		Res(c, -1, nil, errors.New("用户名不能为空！"))
		return
	}

	password, ok := args["password"]
	if !ok || password == "" {
		Res(c, -1, nil, errors.New("密码不能为空！"))
		return
	}

	var data map[string]interface{}

	sel := m.Select().
		Fields("password", "name").
		From("user").
		Where(m.IS("user", user))

	var rows m.Rows
	if rows, err = sel.Rows(sel.Str()); err != nil {
		Res(c, -1, nil, err)
		return
	}

	if len(rows) == 0 {
		err = errors.New("当前用户不存在！")
		Res(c, -1, nil, err)
		return
	}

	if rows[0][0].(string) != password {
		err = errors.New("用户密码错误！")
		Res(c, -1, nil, err)
		return
	}

	name := rows[0][1]
	menus := []string{"depAcct"}
	if user == "admin" {
		menus = append(menus, "depInput", "depUpdate")
	}

	data = map[string]interface{}{
		"user":  user,
		"name":  name,
		"menus": menus,
	}

	t, _ := buildToken(user)

	c.Header("Token", t)
	Res(c, 0, data, err)
}
