package Test

import (
	"backend/Models"
	"fmt"
	"testing"
)

func TestSearchMovieByCast(t *testing.T) {
	var castName = ""
	if Models.SearchMovieByCast(castName) == nil {
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

func TestSearchMovieByYear(t *testing.T) {
	var year = 2021
	//fmt.Println(Models.SearchMovieByMovieId(5))
	//movie := Models.Movie{}
	if Models.SearchMovieByYear(year) == nil {
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

func TestAddMovieCast(t *testing.T) {
	var movieCast = Models.Movie_cast{Cast_id: 2, Movie_id: 3}
	if Models.AddMovieCast(movieCast) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
