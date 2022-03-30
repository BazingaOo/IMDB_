package Controllers

import (
	myjwt "backend/Middleware"
	"backend/Models"
	"backend/Services"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// LogIn Sign in
func LogIn(c *gin.Context) {
	var user Models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	if Models.LogIn(user).User_id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "username or password is wrong",
		})
	} else {
		generateToken(c, Models.LogIn(user))
	}

}

//SignUp for User
func SignUp(c *gin.Context) {
	var user Models.User
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	user = Models.User{
		Username:  c.PostForm("username"),
		Password:  c.PostForm("password"),
		User_type: 0,
		Gender:    c.PostForm("gender"),
		Age:       user.Age,
		Email:     c.PostForm("email"),
	}
	//Models.SignUp(user)
	if Models.SignUp(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    user,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    user,
		})
	}
}

// CheckUsername When sign up, check if the username is unique
func CheckUsername(c *gin.Context) {
	username := c.PostForm("username")
	isTure := Services.CheckUsername(username)
	if isTure == true {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(201, gin.H{
			"message": "error",
		})
	}
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	Models.User
}

// 生成Token
func generateToken(c *gin.Context, user Models.User) {
	j := &myjwt.JWT{
		[]byte("Tao"),
	}
	claims := myjwt.CustomClaims{
		user.User_id,
		user.Username,
		user.Password,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "Tao",                           //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	log.Println(token)
	//data := LoginResult{
	//	User:  user,
	//	Token: token,
	//}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "login successfully",
		"status":  0,
		"user":    user,
		"token":   token,
	})
	return
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "valid token",
			"data":   claims,
		})
	}
}

//Update user profile

func UpdateUserProfile(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		var user Models.User
		user.User_id, _ = strconv.Atoi(c.PostForm("userId"))
		user.Age, _ = strconv.Atoi(c.PostForm("age"))
		user = Models.User{
			User_id:   user.User_id,
			Username:  c.PostForm("username"),
			Password:  c.PostForm("password"),
			User_type: 0,
			Gender:    c.PostForm("gender"),
			Age:       user.Age,
			Email:     c.PostForm("email"),
		}
		if Models.UpdateProfile(user) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "update error",
				"user":    user,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "update success",
				"user":    user,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 0,
			"msg":    "invalid token",
			"data":   claims,
		})
	}
}

//DeleteUser

func DeleteUser(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		var userId, _ = strconv.Atoi(c.PostForm("userId"))
		if Models.DeleteUser(userId) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "delete error",
				"userId":  userId,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete success",
				"userId":  userId,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 0,
			"msg":    "invalid token",
			"data":   claims,
		})
	}
}

func QueryUserInfoById(c *gin.Context) {
	var userId, _ = strconv.Atoi(c.PostForm("userId"))
	Models.QueryUserInfoByUserId(userId)
	c.JSON(http.StatusOK, gin.H{
		"userInfo": Models.QueryUserInfoByUserId(userId),
	})
}
