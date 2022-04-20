package Test

import (
	"backend/Models"
	"testing"
)

func TestAddRating(t *testing.T) {
	var rating = Models.Rating{Movie_id: 1, User_id: 2, Score: 1}
	if Models.AddRating(rating) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestDeleteRating(t *testing.T) {

	var rating = Models.Rating{User_id: 2, Movie_id: 3}
	if Models.DeleteRating(rating) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestReadRating(t *testing.T) {
	//var user_id = 1
	var rating = Models.Rating{}
	if Models.ReadRating(2, 3) == rating {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateRating(t *testing.T) {
	var rating = Models.Rating{User_id: 1, Movie_id: 3, Score: 4}
	if Models.UpdateRating(rating) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateGrade(t *testing.T) {
	if Models.UpdateGrade(3) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
