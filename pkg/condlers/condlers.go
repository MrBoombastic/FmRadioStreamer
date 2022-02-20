// Why 'condlers'? Because 'converters' and 'downloaders'...

package condlers

import (
	"github.com/TheKinrar/goydl"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// DownloadWav download specific video using youtube-dl, extracts audio from it and converts it to wave.
func DownloadWav(URL string) error {
	youtubeDl := goydl.NewYoutubeDl()
	// It is impossible to get final filename and pass it later to WAV converter. Instead, we will mark file to convert by adding "TEMP_" prefix.
	youtubeDl.Options.Output.Value = "music/TEMP_%(title)s.%(ext)s"
	youtubeDl.Options.ExtractAudio.Value = true
	/*
		Hopefully, output will be already in opus codec.
		"Why don't you set wave codec here?" you may ask
		Well, it doesn't always work. :)
		And vorbis option still sometimes gives untouched opus file.
	*/
	// youtubeDl.Options.AudioFormat.Value = "opus"

	// This breaks my RPi, so I commented it out...
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
	err = OpusToWav()
	if err != nil {
		return err
	}
	return nil
}

func OpusToWav() error {
	// Getting list of music files in "music/" directory
	list, err := MusicDir()
	if err != nil {
		return err
	}
	var filename string
	// Finding first unconverted file to convert
	for _, s := range list {
		if strings.HasPrefix(s, "TEMP_") {
			filename = s
			break
		}
	}
	// Creating new filename
	replacer := strings.NewReplacer("TEMP_", "", ".opus", ".wav")
	newFilename := replacer.Replace(filename)
	// Starting the conversion
	cmd := exec.Command("opusdec", "--force-wav", "--rate", "48000", "music/"+filename, "music/"+newFilename)
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	// Deleting temp file
	err = os.Remove("music/" + filename)
	if err != nil {
		return err
	}
	return nil
}

// MusicDir returns slice of all files in the 'music' directory
func MusicDir() ([]string, error) {
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
