package Test

import (
	userController "backend/Controllers"
	jwt "backend/Middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUpdateUserProfile(t *testing.T) {
	router := gin.New()
	userGroup := router.Group("/user")
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.POST("/UpdateProfile", userController.UpdateUserProfile)
		userGroup.POST("/DeleteUser", userController.DeleteUser)
		userGroup.POST("/queryUserInfoById", userController.QueryUserInfoById)
	}

	params := url.Values{}
	params.Add("userId", "2")
	params.Add("username", "qwe2")
	params.Add("password", "123456")
	params.Add("gender", "female")
	params.Add("age", "2")
	params.Add("email", "2")

	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/UpdateProfile", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestDeleteUser(t *testing.T) {
	router := gin.New()
	userGroup := router.Group("/user")
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.POST("/UpdateProfile", userController.UpdateUserProfile)
		userGroup.POST("/DeleteUser", userController.DeleteUser)
		userGroup.POST("/queryUserInfoById", userController.QueryUserInfoById)
	}

	params := url.Values{}
	params.Add("userId", "2")

	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/DeleteUser", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestQueryUserInfoById(t *testing.T) {
	router := gin.New()
	userGroup := router.Group("/user")
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.POST("/UpdateProfile", userController.UpdateUserProfile)
		userGroup.POST("/DeleteUser", userController.DeleteUser)
		userGroup.POST("/queryUserInfoById", userController.QueryUserInfoById)
	}

	params := url.Values{}
	params.Add("userId", "2")

	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/queryUserInfoById", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
