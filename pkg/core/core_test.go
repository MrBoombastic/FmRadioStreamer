package core

import (
	"os"
	"reflect"
	"testing"
)

func TestGenerateOptions(t *testing.T) {
	err := os.Rename("config.json.example", "config.json")
	if err != nil {
		t.Log(err)
	}
	got := GenerateOptions("")
	wanted := []string{"Config"}
	if reflect.DeepEqual(got, wanted) {
		t.Errorf("got %v, wanted %v", got, wanted)
	}
}
