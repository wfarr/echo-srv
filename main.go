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
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(body)

	// headers, err := json.Marshal(r.Header)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	log.WithFields(log.Fields{
		"method":  r.Method,
		"url":     r.URL.String(),
		"headers": r.Header,
		"body":    string(body),
		"status":  200,
	}).Info("done with request")
}
