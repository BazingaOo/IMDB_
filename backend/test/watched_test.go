package Test

import (
	"backend/Models"
	"testing"
)

func TestAddWatched(t *testing.T) {
	var watched = Models.Watched_list{User_id: 14, Movie_id: 1}
	if Models.AddWatched(watched) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestDeleteWatched(t *testing.T) {
	//var watched = Models.Watched_list{User_id: 14, Movie_id: 1}
	if Models.DeleteWatched(14, 1) == 0 {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}

func TestReadWatched(t *testing.T) {
	if Models.ReadWatched(13) == nil {
		t.Error("result is wrong!")
	} else {
		t.Log("result is right")
	}
}
