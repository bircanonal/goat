package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"

    //"goat/config"
    "goat/internal/core/jwt"
    "goat/internal/service"
)


//GET /api/profile
func (ctr *userController) getProfile(c *gin.Context) {
    user, err := ctr.uServ.GetProfile(jwt.GetUserId(c))

    if err != nil {
        c.JSON(500, gin.H{"error": http.StatusText(500)})
        c.Abort()
        return
    }

    c.JSON(200, user)
}


//PUT[POST] /api/password
func (ctr *userController) changePassword(c *gin.Context) {
    userId := jwt.GetUserId(c)

    m := map[string]string{}
    c.BindJSON(&m)
    pass := m["password"]

    if ctr.uServ.ChangePassword(userId, pass) != service.CHANGE_PASSWORD_SUCCESS_INT {
        c.JSON(500, gin.H{"error": "登録に失敗しました。"})
        c.Abort()
        return
    }

    c.JSON(200, gin.H{})
}


//PUT[POST] /api/username
func (ctr *userController) changeUsername(c *gin.Context) {
    userId := jwt.GetUserId(c)

    m := map[string]string{}
    c.BindJSON(&m)
    name := m["user_name"]

    if ctr.uServ.ChangeUsername(userId, name) != service.CHANGE_USERNAME_SUCCESS_INT {
        c.JSON(500, gin.H{"error": "登録に失敗しました。"})
        c.Abort()
        return
    }

    c.JSON(200, gin.H{})
}


//DELETE /api/account
func (ctr *userController) deleteUser(c *gin.Context) {
    userId := jwt.GetUserId(c)

    if ctr.uServ.DeleteUser(userId) != service.DELETE_USER_SUCCESS_INT {
        c.JSON(500, gin.H{"error": "削除に失敗しました。"})
        c.Abort()
        return
    }

    c.JSON(200, gin.H{})
}


/* ----------------------------------------------------------*/
// Use either user.go(signup, login, logout) 
//  or user.api.go(signup, login, logout)
/* ----------------------------------------------------------*/

/*
//POST /api/signup
func (ctr *userController) signup(c *gin.Context) {
    m := map[string]string{}
    c.BindJSON(&m)
    name := m["user_name"]
    pass := m["password"]

    result := ctr.uServ.Signup(name, pass)

    if result == service.SIGNUP_SUCCESS_INT {
        c.JSON(200, gin.H{})

    } else if result == service.SIGNUP_CONFLICT_INT {
        c.JSON(409, gin.H{
            "error": "Nameが既に使われています。",
        })

    } else {
        c.JSON(500, gin.H{
            "error": "登録に失敗しました。",
        })
    }
}


//POST /api/login
func (ctr *userController) login(c *gin.Context) {
    m := map[string]string{}
    c.BindJSON(&m)
    name := m["user_name"]
    pass := m["password"]

    userId := ctr.uServ.Login(name, pass)

    if userId == service.LOGIN_FAILURE_INT {
        c.JSON(401, gin.H{
            "error": "NameまたはPasswordが異なります。",
        })
        c.Abort()
        return
    }

    jwtStr := ctr.uServ.GenerateJWT(userId)

    if jwtStr == service.GENERATE_JWT_FAILURE_STR {
        c.JSON(500, gin.H{
            "error": "ログインに失敗しました。",
        })
        c.Abort()
        return
    }
    cf := config.GetConfig()
    c.SetCookie(jwt.COOKIE_KEY_JWT, jwtStr, int(jwt.JWT_EXPIRES), "/", cf.AppHost, false, true)
    c.JSON(200, gin.H{})

    //when "jwt" store in localStorage
    //c.JSON(200, gin.H{"access_token": jwtStr}))
}


//GET /api/logout
func (ctr *userController) logout(c *gin.Context) {
    cf := config.GetConfig()
    c.SetCookie(jwt.COOKIE_KEY_JWT, "", 0, "/", cf.AppHost, false, true)
    c.JSON(200, gin.H{})
}
*/