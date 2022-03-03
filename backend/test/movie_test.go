package Test

import (
	"backend/Models"
	"fmt"
	"testing"
)

func TestSearchMovieByCast(t *testing.T) {
	var review = ""
	if Models.SearchMovieByCast(review) == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestAddMovie(t *testing.T) {
	movie := Models.Movie{
		Movie_id:   4,
		Movie_name: "test",
		Year:       1998,
		Desciption: "awesome!"}
	if Models.AddMovie(movie) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateMovie(t *testing.T) {
	movie := Models.Movie{
		Movie_id:   4,
		Movie_name: "test",
		Year:       2001,
		Desciption: "awesome!"}
	if Models.UpdateMovie(movie) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestSelectMovie(t *testing.T) {
	fmt.Println(Models.SearchMovieByMovieId(5))
	movie := Models.Movie{}
	if Models.SearchMovieByMovieId(4) == movie {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
