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

func TestDownloadWav(t *testing.T) {
	/*err := os.Chdir("../../")
	if err != nil {
		t.Error(err)
		return
	}*/
	err := DownloadWav("https://www.youtube.com/watch?v=lp7zvP4GxQA")
	if err != nil {
		t.Error(err)
		return
	}
	got, err := MusicDir()
	if err != nil {
		t.Error(err)
		return
	}
	//
	wanted := []string{"Jarre arr. by Rob Hubbard - Zoolook (Oscilloscope View).wav", "piano-kozco-com.wav"}
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("got %v, wanted %v", got, wanted)
	}
}
