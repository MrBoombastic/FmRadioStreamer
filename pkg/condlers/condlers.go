// Why 'condlers'? Because 'converters' and 'downloaders'...

package condlers

import (
	"context"
	"github.com/lrstanley/go-ytdlp"
	"os"
)

func CheckYtdlp() (*ytdlp.ResolvedInstall, error) {
	return ytdlp.Install(context.TODO(), nil)
}

// Download downloads specific video (from given URL) using yt-dlp, extracts audio from it and converts it to chosen format.
func Download(URL string, format string) (err error) {
	dl := ytdlp.New().FormatSort("ba").ExtractAudio().AudioQuality("0").RecodeVideo(format).Output("music/%(title)s.%(ext)s")
	// .Proggres()
	_, err = dl.Run(context.TODO(), URL)
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
