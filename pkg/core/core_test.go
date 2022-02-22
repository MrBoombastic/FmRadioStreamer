package core

import (
	"reflect"
	"testing"
)

func TestGenerateOptions(t *testing.T) {
	got := GenerateOptions("")
	wanted := []string{"Config"}
	if reflect.DeepEqual(got, wanted) {
		t.Errorf("got %v, wanted %v", got, wanted)
	}
}
