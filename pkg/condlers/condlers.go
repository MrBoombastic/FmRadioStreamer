// Why 'condlers'? Because 'converters' and 'downloaders'...

package condlers

import (
	"fmt"
	"github.com/TheKinrar/goydl"
	"strings"
)

// DownloadAudioFromYoutube download specific video from YouTube and converts it to OGG/Vorbis format.
func DownloadAudioFromYoutube(ID string, filename string) error {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Output.Value = fmt.Sprintf("music/%v.wav", strings.ReplaceAll(filename, " ", "-"))
	youtubeDl.Options.ExtractAudio.Value = true
	youtubeDl.Options.AudioFormat.Value = "vorbis" //"wav" does not produce good input for PiFmAdv. It is still OGG/Opus after conversion.
	// As usual, this breaks my Pi, so I commented it out...
	// go io.Copy(os.Stdout, youtubeDl.Stdout)
	// go io.Copy(os.Stderr, youtubeDl.Stderr)
	cmd, err := youtubeDl.Download("https://youtu.be/" + ID)
	if err != nil {
		return err
	}
	cmd.Wait()
	return nil
}
