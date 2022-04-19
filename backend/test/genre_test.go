package Test

import (
	"backend/Models"
	"testing"
)

func TestAddGenre(t *testing.T) {
	var genre = Models.Genre{Genre_id: 2, Genre_name: "3"}
	if Models.AddGenre(genre) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateGenre(t *testing.T) {
	var genre = Models.Genre{Genre_id: 1, Genre_name: "13"}
	if Models.AddGenre(genre) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestDeleteGenre(t *testing.T) {
	var genreId = 1
	if Models.DeleteGenre(genreId) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestSearchGenreName(t *testing.T) {
	var genreName = "3"
	if Models.SearchGenreName(genreName) == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestSearchGenre(t *testing.T) {
	var genreName = "123"
	if Models.SearchGenre(genreName) == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

//func TestDeleteRating(t *testing.T) {
//
//	var rating = Models.Rating{User_id: 2, Movie_id: 3}
//	if Models.DeleteRating(rating) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}
//
//func TestReadRating(t *testing.T) {
//	//var user_id = 1
//	if Models.ReadRating(2) == nil {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}
//
//func TestUpdateRating(t *testing.T) {
//	var rating = Models.Rating{User_id: 1, Movie_id: 3, Score: 4}
//	if Models.UpdateRating(rating) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}
//
//func TestUpdateGrade(t *testing.T) {
//	if Models.UpdateGrade(3) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}
