// Why 'condlers'? Because 'converters' and 'downloaders'...

package condlers

import (
	"fmt"
	"github.com/TheKinrar/goydl"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// DownloadWav download specific video using youtube-dl, extracts audio and converts it to wave
func DownloadWav(URL string) error {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Output.Value = "music/TEMP_%(title)s.%(ext)s"
	youtubeDl.Options.ExtractAudio.Value = true
	// Hopefully, output will be already in opus codec
	// youtubeDl.Options.AudioQuality.Value = "opus"

	// As usual, this breaks my Pi, so I commented it out...
	// go io.Copy(os.Stdout, youtubeDl.Stdout)
	// go io.Copy(os.Stderr, youtubeDl.Stderr)
	cmd, err := youtubeDl.Download(URL)
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	OpusToWav()
	return nil
}

func OpusToWav() {
	list, err := MusicList()
	if err != nil {
		log.Println(err)
		return
	}
	var filename string
	for _, s := range list {
		if strings.HasPrefix(s, "TEMP_") {
			filename = s
			break
		}
	}
	//opusdec --rate 48000 input.opus output.wav
	replacer := strings.NewReplacer("TEMP_", "", ".opus", ".wav")
	newFilename := replacer.Replace(filename)
	cmd := exec.Command("opusdec", "--force-wav", "--rate", "48000", "music/"+filename, "music/"+newFilename)
	cmd.Start()
	cmd.Wait()
	fmt.Println("Done")
	os.Remove("music/" + filename)
}

// MusicList returns slice of all files in 'music' directory
func MusicList() ([]string, error) {
	files, err := ioutil.ReadDir("music/")
	if err != nil {
		return nil, err
	}
	var filesSlice []string
	for _, item := range files {
		filesSlice = append(filesSlice, item.Name())
	}
	return filesSlice, nil
}
