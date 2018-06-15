package controllers

import (
	"encoding/json"
	"fmt"
	"webgl-cors-cookie-test/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

var loginList = map[string]string{}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {

	username := u.GetString("username")
	password := u.GetString("password")
	random := u.GetString("random")
	deviceType := u.GetString("type")
	osInfo := u.GetString("osInfo")

	beego.Info(fmt.Sprintf("username=%s, password=%s, random=%s, type=%s, osInfo=%s", username, password, random, deviceType, osInfo))
	// all pass

	if models.Login(username, password) {
		// u.Data["json"] = "login success"
		u.Data["json"] = map[string]interface{}{
			"err":  "null",
			"info": "29b6de8a20dae4e25cdf9b5be6d4e1b5",
		}
		var sessionID = "123456" //strconv.FormatInt(time.Now().Unix(), 10)
		loginList[username] = sessionID
		var cookiesecret = beego.AppConfig.String("Cookie")
		// u.SetSecureCookie(cookiesecret, "Cookie", sessionID, nil)
		u.SetSecureCookie(cookiesecret, "sessionID", sessionID, 0, "/", "alex.xinhcao9999.com:8888", true, false)
		beego.Debug(fmt.Sprintf("set session=%s", sessionID))
	} else {
		// u.Data["json"] = "user not exist"
		u.Data["json"] = map[string]interface{}{
			"err":  "paasword",
			"info": "用户名或者密码错误",
		}
	}
	u.ServeJSON()
}

// @Title Info
// @Description Logs user into the system
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /info [post]
func (u *UserController) Info() {

	var cookiesecret = beego.AppConfig.String("Cookie")
	var ok bool = false
	var sessionID string = ""

	sessionID, ok = u.GetSecureCookie(cookiesecret, "Cookie")

	if ok {
		if sessionID == loginList["t001"] {
			beego.Debug(fmt.Sprintf("get session=%s", sessionID))
			u.Data["json"] = map[string]interface{}{
				"Uid":      3,
				"UserName": "t001",
				"State":    0,
				"Grade":    1,
				"Amount":   2548.148,
				"FAmount":  0.000,
				"Percent":  0.9750,
				"NewMsg":   true,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"err":  "timeout",
				"info": "登录超时,请重新登录!",
			}
		}

	} else {
		u.Data["json"] = map[string]interface{}{
			"err":  "timeout",
			"info": "登录超时,请重新登录!",
		}
	}

	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
