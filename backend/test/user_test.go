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
