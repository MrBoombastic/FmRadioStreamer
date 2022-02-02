// Why 'condlers'? Because 'converters' and 'downloaders'...

package condlers

import (
	"fmt"
	"github.com/TheKinrar/goydl"
	"strings"
)

func DownloadAudioFromYoutube(ID string, filename string) error {
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Output.Value = fmt.Sprintf("music/%v.wav", strings.ReplaceAll(filename, " ", "-"))
	youtubeDl.Options.ExtractAudio.Value = true
	youtubeDl.Options.AudioFormat.Value = "wav"
	//	go io.Copy(os.Stdout, youtubeDl.Stdout)
	//	go io.Copy(os.Stderr, youtubeDl.Stderr)
	cmd, err := youtubeDl.Download("https://youtu.be/" + ID)
	fmt.Printf("Title: %s\n", youtubeDl.Info.Title)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cmd.Wait()
	return nil
}
