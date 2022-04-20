package Test

import (
	userController "backend/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//func TestGetList(t *testing.T) {
//
//	router := gin.New()
//	customerGroup := router.Group("/user")
//	{
//		customerGroup.POST("/login", userController.LogIn) // when create sth
//	}
//
//	params := url.Values{}
//	params.Add("username", "qwe")
//	params.Add("password", "123")
//	para1 := params.Encode()
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(para1))
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	router.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//	if Models.GetList("qwe") == nil {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

//func TestDeleteUser(t *testing.T) {
//	if Models.DeleteUser(1) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//
//}

//func TestUpdateProfile(t *testing.T) {
//	var user = Models.User{User_id: 2, Username: "hhhh", Password: "123", User_type: 1, Gender: "s", Age: 2, Email: "123"}
//	if Models.UpdateProfile(user) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

//func TestSignUp(t *testing.T) {
//	var user = Models.User{User_id: 2, Username: "hhhh", Password: "123", User_type: 1, Gender: "s", Age: 2, Email: "123"}
//	//var review = Models.Review{Review_id: 2, Review_content: "lalalal", User_id: 11, Movie_id: 1}
//	if Models.SignUp(user) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

func TestLogIn(t *testing.T) {

	router := gin.New()
	customerGroup := router.Group("/user")
	{
		customerGroup.POST("/login", userController.LogIn)                 // when create sth
		customerGroup.POST("/signUp", userController.SignUp)               // when query sth
		customerGroup.POST("/checkUsername", userController.CheckUsername) // when update sth

	}

	params := url.Values{}
	params.Add("username", "qwe")
	params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	params2 := url.Values{}
	params2.Add("username", "qwe")
	params2.Add("password", "123")
	params2.Add("user_type", "0")
	params2.Add("gender", "female")
	params2.Add("age", "12")
	params2.Add("email", "123@")
	para2 := params2.Encode()
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/user/signUp", strings.NewReader(para2))
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	params3 := url.Values{}
	params3.Add("username", "qwe")
	para3 := params3.Encode()
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("POST", "/user/checkUsername", strings.NewReader(para3))
	req3.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w3, req3)
	assert.Equal(t, 201, w3.Code)

}
