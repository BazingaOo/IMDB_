package Test

import (
	"backend/Models"
	"testing"
)

func TestAddCast(t *testing.T) {
	//vat cast=Models.Cast{}
	var cast = Models.Cast{Cast_id: 4, Cast_name: "lalalal", Cast_description: "11", Cast_image: "123"}
	if Models.AddCast(cast) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestDeleteCast(t *testing.T) {
	//var review = Models.Review{Review_id: 1, Review_content: "ttttt", User_id: 11, Movie_id: 1}
	//var review_id = 1
	if Models.DeleteCast(1) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateCast(t *testing.T) {
	var cast = Models.Cast{Cast_id: 2, Cast_name: "lalalal", Cast_description: "11", Cast_image: "123"}
	if Models.UpdateCast(cast) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestSearchCast(t *testing.T) {
	//var cast = Models.Cast{Cast_id: 2, Cast_name: "lalalal", Cast_description: "11", Cast_image: "123"}
	var castName = "Tao"
	if Models.SearchCastByName(castName) == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestSearchCastById(t *testing.T) {
	var castId = 2
	resCast := Models.Result{}
	if Models.SearchCastById(castId) == resCast {
		t.Error("result is wrong!")
	} else {
		//fmt.Println(Models.SearchMovieByName(name))
		t.Log("result is right")
	}
}
