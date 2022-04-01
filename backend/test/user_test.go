package Test

import (
	"backend/Models"
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
	var user = Models.User{User_id: 2, Username: "hhhh", Password: "123", User_type: 1, Gender: "s", Age: 2, Email: "123"}
	//var review = Models.Review{Review_id: 2, Review_content: "lalalal", User_id: 11, Movie_id: 1}
	if Models.SignUp(user) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
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
