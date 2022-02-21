package condlers

import (
	"os"
	"reflect"
	"testing"
)

func TestMusicDir(t *testing.T) {
	err := os.Chdir("../../")
	if err != nil {
		t.Error(err)
		return
	}
	got, err := MusicDir()
	if err != nil {
		t.Error(err)
		return
	}
	wanted := []string{"piano-kozco-com.wav"}
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("got %v, wanted %v", got, wanted)
	}
}
