// Why 'condlers'? Because 'converters' and 'downloaders'...

package condlers

import (
	"fmt"
	"github.com/TheKinrar/goydl"
	"strings"
)

// DownloadAudio download specific video using youtube-dl and converts it to OGG/Vorbis format.
func DownloadAudio(URL string, filename string) error {
	youtubeDl := goydl.NewYoutubeDl()
	if filename != "" {
		youtubeDl.Options.Output.Value = fmt.Sprintf("music/%v.ogg", strings.ReplaceAll(filename, " ", "-"))
	} else {
		youtubeDl.Options.Output.Value = "music/%(title)s.%(ext)s"
	}
	youtubeDl.Options.ExtractAudio.Value = true
	youtubeDl.Options.AudioFormat.Value = "vorbis" //"wav" does not produce good input for PiFmAdv. It is still OGG/Opus after conversion.
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
	return nil
}
