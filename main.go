package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
	"github.com/gorilla/handlers"
	"github.com/targodan/alexa-jarvis/ssml"

	"github.com/go-alexa/alexa/events"
	"github.com/go-alexa/alexa/parser"
	"github.com/go-alexa/alexa/response"
	"github.com/go-alexa/alexa/server"
	"github.com/go-alexa/alexa/validations"
)

func main() {
	d, err := bolt.Open("info.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	validations.DB = d
	validations.AppID = "amzn1.ask.skill.f46f9a81-b729-4f82-8247-556f83901793"

	ev := events.New().
		Add("TestIntent",
			func(ev *parser.Event) (*response.Response, error) {
				return response.New().AddSSMLSpeech(
					ssml.Speak().
						Text("hurra, es funktioniert!").
						String(),
				), nil
			})

	server.Events = ev

	mux := http.NewServeMux()
	mux.HandleFunc("/jarvis", server.Handler)

	handler := handlers.LoggingHandler(os.Stdout, handlers.ProxyHeaders(mux))

	err = http.ListenAndServe(server.Host, handler)
	if err != nil {
		fmt.Println(err)
	}
}
