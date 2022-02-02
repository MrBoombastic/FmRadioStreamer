package dashboard

import (
	"encoding/json"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/condlers"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/leds"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/tools"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFS(sites, "sites/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, "")
}

func music(w http.ResponseWriter, _ *http.Request) {
	filesSlice := musicList()
	filesJson, err := json.Marshal(filesSlice)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(filesJson)
}

func loudstop(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("loudstop")
	w.Write([]byte("OK"))
}
func superstop(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("superstop")
	w.Write([]byte("OK"))
}
func yt(w http.ResponseWriter, req *http.Request) {
	onlySearch := req.FormValue("search")
	query := req.FormValue("q")
	result := tools.SearchYouTube(query)
	searchJson, err := json.Marshal(result.Items[0].Snippet)
	if err != nil {
		log.Fatal(err)
	}
	if onlySearch == "true" {
		w.Write(searchJson)
	} else {
		go leds.BlueLedLoopStart()
		err := condlers.DownloadAudioFromYoutube(result.Items[0].ID.VideoID, result.Items[0].Snippet.Title)
		leds.BlueLedLoopStop()
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte("OK"))
	}
}
func play(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("play")
	w.Write([]byte("OK"))
}
func save(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var newConfig config.Config
	json.Unmarshal(body, &newConfig)
	config.Save(newConfig)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
func configuration(w http.ResponseWriter, _ *http.Request) {
	configMap := config.GetMap()
	configJson, err := json.Marshal(configMap)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(configJson)
}
