package dashboard

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//go:embed sites/*
var sites embed.FS

//go:embed public/*
var public embed.FS

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
func yt(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("yt")
	w.Write([]byte("OK"))
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
func musicList() []string {
	files, err := ioutil.ReadDir("music/")
	if err != nil {
		log.Fatal(err)
	}
	var filesSlice []string
	for _, item := range files {
		filesSlice = append(filesSlice, item.Name())
	}
	return filesSlice
}

func Init() {
	// Handle static files
	var publicFS = http.FS(public)
	fs := http.FileServer(publicFS)
	http.Handle("/public/", fs)
	// Handle else
	http.HandleFunc("/music", music)
	http.HandleFunc("/loudstop", loudstop)
	http.HandleFunc("/superstop", superstop)
	http.HandleFunc("/yt", yt)
	http.HandleFunc("/config", configuration)
	http.HandleFunc("/play", play)
	http.HandleFunc("/save", save)
	http.HandleFunc("/", index)
	// Start!
	port := config.GetPort()
	go http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	fmt.Println(fmt.Sprintf("Dashboard listening at port %v!", port))
}
