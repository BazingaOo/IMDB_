package Controllers

import (
	"backend/Models"
	userService "backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Sign in
func SignIn(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	user := Models.User{
		Username: username,
		Password: password,
	}
	userService.Signin(user)
	//if userService.Signin(user) == 0 {
	//	ctx.JSON(100, gin.H{
	//		"code":     100,
	//		"message":  "error",
	//		"username": "",
	//		"password": "",
	//	})
	//} else {
	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "success",
		"username": username,
		"password": password,
	})
}

//}

//SignUp
func SignUp(ctx *gin.Context) {
	username := ctx.PostForm("username")
	ctx.JSON(200, gin.H{
		"result":    "ok",
		"hello":     username,
		"developer": "Tao!",
	})
	user := Models.User{
		User_id:   3,
		Username:  username,
		Password:  "123456",
		User_type: 0,
		Gender:    0,
		Birthday:  time.Now(),
	}
	userService.AddUser(user)
}
