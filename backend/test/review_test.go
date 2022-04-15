package Test

import (
	"backend/Models"
	"testing"
)

func TestAddReview(t *testing.T) {
	var review = Models.Review{Review_id: 1, Review_content: "ldfgdfgl", User_id: 2, Movie_id: 2}
	if Models.AddReview(review) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestDeleteReview(t *testing.T) {
	//var review = Models.Review{Review_id: 1, Review_content: "ttttt", User_id: 11, Movie_id: 1}
	//var review_id = 1
	if Models.DeleteReview(1) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestReadReview(t *testing.T) {
	//var user_id = 1
	if Models.ReadReview(11) == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateReview(t *testing.T) {
	var review = Models.Review{Review_id: 2, Review_content: "hhhh", User_id: 1, Movie_id: 1}
	if Models.UpdateReview(review) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
