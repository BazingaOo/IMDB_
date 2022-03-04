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

func TestLogIn(t *testing.T) {
	var user = Models.User{Username: "qwe", Password: "123"}
	if Models.LogIn(user) == true {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestSignUp(t *testing.T) {
	var user = Models.User{Username: "zt", Password: "zt123", User_type: 0, Gender: "male", Age: 23}
	if Models.SignUp(user) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateProfile(t *testing.T) {

	var user = Models.User{Username: "zt", Password: "zt123", User_type: 0, Gender: "male", Age: 23}
	if Models.UpdateProfile(user) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}

}

func TestDeleteUser(t *testing.T) {
	//var user = Models.User{User_id: 11}
	if Models.DeleteUser(11) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
