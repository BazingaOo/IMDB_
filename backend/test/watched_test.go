package Test

import (
	"backend/Models"
	"testing"
)

func TestAddWatched(t *testing.T) {
	var watched = Models.Watch_list{User_id: 2, Movie_id: 3}
	if Models.AddWatched(watched) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestDeleteWatched(t *testing.T) {
	//var review = Models.Review{Review_id: 1, Review_content: "ttttt", User_id: 11, Movie_id: 1}
	//var review_id = 1
	if Models.DeleteWatched(2, 1) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestReadWatched(t *testing.T) {
	//var user_id = 1
	if Models.ReadWatched(2) == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

//
//func TestUpdateReview(t *testing.T) {
//	var review = Models.Review{Review_id: 2, Review_content: "hhhh", User_id: 11, Movie_id: 1}
//	if Models.UpdateReview(review) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}
