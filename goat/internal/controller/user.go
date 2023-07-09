package controller

import (
	"github.com/gin-gonic/gin"

	"goat/config"
	"goat/internal/core/jwt"
	"goat/internal/service"
	"goat/internal/model/entity"
)


type UserService interface {
	Signup(username, password string) int
	Login(username, password string) int
	GenerateJWT(userId int) string
	GetProfile(userId int) (entity.User, error)
	ChangeUsername(userId int, username string) int
	ChangePassword(userId int, password string) int
	DeleteUser(userId int) int
}

type userController struct {
	uServ UserService
}


func NewUserController() *userController {
	uServ := service.NewUserService()
	return &userController{uServ}
}


//GET /signup
func (ctr *userController) SignupPage(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

//GET /login
func (ctr *userController) LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}


//POST /signup
func (ctr *userController) Signup(c *gin.Context) {
	name := c.PostForm("user_name")
	pass := c.PostForm("password")

	result := ctr.uServ.Signup(name, pass)

	if result == service.SIGNUP_SUCCESS_INT {
		c.Redirect(303, "/login")

	} else if result == service.SIGNUP_CONFLICT_INT {
		c.HTML(409, "signup.html", gin.H{
			"error": "Usernameが既に使われています。",
		})

	} else {
		c.HTML(500, "signup.html", gin.H{
			"error": "登録に失敗しました。",
		})
	}
}


//POST /login
func (ctr *userController) Login(c *gin.Context) {
	name := c.PostForm("user_name")
	pass := c.PostForm("password")

	userId := ctr.uServ.Login(name, pass)

	if userId == service.LOGIN_FAILURE_INT {
		c.HTML(401, "login.html", gin.H{
			"error": "UserNameまたはPasswordが異なります。",
		})
		c.Abort()
		return
	}

	jwtStr := ctr.uServ.GenerateJWT(userId)

	if jwtStr == service.GENERATE_JWT_FAILURE_STR {
		c.HTML(500, "login.html", gin.H{
			"error": "ログインに失敗しました。",
		})
		c.Abort()
		return
	}

	cf := config.GetConfig()
	c.SetCookie(jwt.COOKIE_KEY_JWT, jwtStr, int(jwt.JWT_EXPIRES), "/", cf.AppHost, false, true)
	c.Redirect(303, "/")
}


//GET /logout
func (ctr *userController) Logout(c *gin.Context) {
	cf := config.GetConfig()
	c.SetCookie(jwt.COOKIE_KEY_JWT, "", 0, "/", cf.AppHost, false, true)
	c.Redirect(303, "/login")
}
