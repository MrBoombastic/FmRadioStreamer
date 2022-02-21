package config

import (
	"os"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	err := os.Chdir("../../")
	if err != nil {
		t.Error(err)
		return
	}
	err = os.Rename("config.json.example", "config.json")
	if err != nil {
		t.Error(err)
		return
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
