package dashboard

import (
	"embed"
	"fmt"
	"github.com/MrBoombastic/FmRadioStreamer/pkg/config"
	"io/ioutil"
	"log"
	"net/http"
)

//go:embed sites/*
var sites embed.FS

//go:embed public/*
var public embed.FS

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

var httpServer = &http.Server{
	Addr: fmt.Sprintf(":%v", config.GetPort()),
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

	go httpServer.ListenAndServe()

	fmt.Println(fmt.Sprintf("Dashboard listening at port %v!", httpServer.Addr))
}
