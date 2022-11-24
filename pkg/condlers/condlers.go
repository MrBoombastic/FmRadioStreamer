// Why 'condlers'? Because 'converters' and 'downloaders'...

package condlers

import (
	"github.com/TheKinrar/goydl"
	"os"
)

// Download downloads specific video (from given URL) using youtube-dl, extracts audio from it and converts it to chosen format.
func Download(URL string, format string) error {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Output.Value = "music/%(title)s.%(ext)s"
	youtubeDl.Options.Format.Value = "bestaudio" //may break at some point, change to "best" if needed
	youtubeDl.Options.ExtractAudio.Value = true
	youtubeDl.Options.AudioFormat.Value = format
	youtubeDl.Options.AudioQuality.Value = "0" //best quality

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
	return nil
}

// MusicDir returns slice of all files in the 'music' directory.
func MusicDir() ([]string, error) {
	files, err := os.ReadDir("music/")
	if err != nil {
		return nil, err
	}
	var filesSlice []string
	for _, item := range files {
		filesSlice = append(filesSlice, item.Name())
	}
	return filesSlice, nil
}
