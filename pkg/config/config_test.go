package config

import (
	"os"
	"reflect"
	"testing"
)

var isPathFixed = false

func fixPath() error {
	if isPathFixed == true {
		return nil
	}
	err := os.Chdir("../../")
	if err != nil {
		return err
	} else {
		isPathFixed = true
		return nil
	}
}
func TestGet(t *testing.T) {
	err := fixPath()
	if err != nil {
		t.Error(err)
		return
	}
	err = os.Rename("config.json.example", "config.json")
	if err != nil {
		t.Log(err)
	}
	get := Get()
	if err != nil {
		t.Error(err)
		return
	}
	got := reflect.TypeOf(get).Name()
	wanted := "Config"
	if got != wanted {
		t.Errorf("got %v, wanted %v", got, wanted)
	}
}
