package Test

import (
	userController "backend/Controllers"
	"backend/Models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetList(t *testing.T) {
	if Models.GetList("qwe") == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestDeleteUser(t *testing.T) {
	if Models.DeleteUser(1) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}

}

func TestUpdateProfile(t *testing.T) {
	var user = Models.User{User_id: 2, Username: "hhhh", Password: "123", User_type: 1, Gender: "s", Age: 2, Email: "123"}
	if Models.UpdateProfile(user) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestSignUp(t *testing.T) {
	router := gin.New()
	customerGroup := router.Group("/user")
	{
		customerGroup.POST("/login", userController.LogIn) // when create sth
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
}

func TestLogIn(t *testing.T) {
	var user = Models.User{Username: "hhhh", Password: "123"}
	resUser := Models.User{}
	if Models.LogIn(user) == resUser {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestQuery(t *testing.T) {
	user := Models.User{}
	if Models.QueryUserInfoByUserId(2) == user {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

//func UpdateCast(cast Cast) int64 {
//	result := Database.DB.Model(&cast).Updates(&cast)
//	//result := Database.DB.Save(&user)
//	return result.RowsAffected
//}
//
//func DeleteCast(Cast_id int) int64 {
//	result := Database.DB.Delete(&Cast{}, Cast_id)
//	return result.RowsAffected
//}
//
//func SearchCast(castName string) []Cast {
//	var cast []Cast
//	//Database.DB.First(&cast, castName)
//	Database.DB.Where("cast_name LIKE ?", "%"+castName+"%").Find(&cast)
//	return cast
//}
