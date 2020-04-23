package controllers

import (
	"encoding/base64"
	"fmt"
	"managerApp/app/models/entities"
	"managerApp/app/models/providers"
	"net/http"
	"strings"
	"time"

	"github.com/revel/revel"
)

type CAuth struct {
	*revel.Controller
	provider *providers.AuthProvider
	//db *sql.DB
}

func (c *CAuth) Init() {
	c.provider = new(providers.AuthProvider)
	c.provider.Init()
}

func cookieHandle(w http.ResponseWriter, name string, value string) {
	expires := time.Now().AddDate(0, 0, 1)
	ck := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
	}
	http.SetCookie(w, &ck)
}

func (c *CAuth) Login(user *entities.User) revel.Result {
	c.Init()
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.Response.Out.Header().Add("WWW-Authenticate", `Basic realm="Please enter your username and password for this site"`)
		c.Response.SetStatus(401)
	}
	// получаем закодированные имя пользователя и пароль
	// убираем подстроку "Basic " и декодируем
	loginAndPassB64 := strings.TrimLeft(authorization, "Basic ")
	bLoginAndPass, err := base64.StdEncoding.DecodeString(loginAndPassB64)
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR decode base64: %v", err))
		return nil
	}
	// конвертируем в string
	loginAndPass := string(bLoginAndPass)
	if len(loginAndPass) != 0 {
		loginAndPassSplitted := strings.Split(loginAndPass, ":")
		userName := loginAndPassSplitted[0]
		password := loginAndPassSplitted[1]
		user, err = c.provider.Login(userName, password)
		if err != nil {
			fmt.Println(err)
			c.Response.Status = 401
			return c.RenderJSON("invalid username or password")
		}
		cookieHandle(c.Response.Out.Server.GetRaw().(http.ResponseWriter), userName, password)
	}
	return c.Render()
}

func (c *CAuth) Logout() revel.Result {
	c.Init()
	return nil
}
