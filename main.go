package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sebastianvera/songify/spotify"
)

var (
	currentSong spotify.Track
	upgrader    = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	port = flag.Int("port", 1616, "HTTP Server port")
)

const (
	addressTmpl = ":%d"

	sleepDuration = 2 * time.Second
)

func main() {
	flag.Parse()
	indexFile, err := os.Open("html/index.html")
	if err != nil {
		fmt.Println(err)
	}

	index, err := ioutil.ReadAll(indexFile)
	if err != nil {
		fmt.Println(err)
	}

	go h.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(index))
	})

	http.Handle("/current-track", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&currentSong); err != nil {
			panic(err)
		}
	}))

	http.HandleFunc("/ws", serveWs)

	go watcher()
	address := fmt.Sprintf(addressTmpl, *port)
	fmt.Println("Listening on", address)
	fmt.Println(http.ListenAndServe(address, nil))
}

func watcher() {
	for {
		track, err := spotify.GetCurrentTrack()
		if err != nil {
			os.Exit(1)
		}

		if *track != currentSong {
			csp, _ := json.Marshal(track)
			h.broadcast <- csp
		}

		currentSong = *track
		time.Sleep(sleepDuration)
	}
}
