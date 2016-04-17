package main

import (
	"flag"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	listen = flag.String("listen", ":8080", "the address to listen to requests on")
)

func main() {
	flag.Parse()

	if *listen == "" {
		log.Fatalln("expected a listen address, found none")
	}

	http.HandleFunc("/", EchoHandler)

	log.WithFields(log.Fields{
		"address": *listen,
		"status":  "listening",
	}).Info("ready")

	log.Fatal(http.ListenAndServe(*listen, nil))
}

// EchoHandler takes a request and echos it back
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := log.WithFields(log.Fields{
		"method":  r.Method,
		"url":     r.URL.String(),
		"headers": r.Header,
	})

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ctx.WithFields(log.Fields{
			"status": 500,
		}).Fatal(err)

		w.WriteHeader(500)
	}

	w.Write(body)

	ctx.WithFields(log.Fields{
		"body":   string(body),
		"status": 200,
	}).Info("OK")
}
