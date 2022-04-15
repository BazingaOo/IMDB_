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
func TestSearchMovieByName(t *testing.T) {
	var name = "2"
	if Models.SearchMovieByName(name) == nil {
		t.Error("result is wrong!")
	} else {
		fmt.Println(Models.SearchMovieByName(name))
		t.Log("11")
	}
}

func TestAddMovie(t *testing.T) {
	movie := Models.Movie{
		Movie_id:    4,
		Movie_name:  "test",
		Year:        1998,
		Description: "awesome!"}
	if Models.AddMovie(movie) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestUpdateMovie(t *testing.T) {
	movie := Models.Movie{
		Movie_id:    4,
		Movie_name:  "test",
		Year:        2001,
		Description: "awesome!"}
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

func TestDeleteMovie(t *testing.T) {
	if Models.DeleteMovie(1) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestTopMovie(t *testing.T) {
	if Models.TopMovie() == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
