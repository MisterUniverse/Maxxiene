package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		img, err := os.Open("./images/red.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer img.Close()
		resp.Header().Set("Content-Type", "image/jpeg")
		io.Copy(resp, img)
	})
	mux.HandleFunc("/blue", func(resp http.ResponseWriter, req *http.Request) {
		img, err := os.Open("./images/blue.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer img.Close()
		resp.Header().Set("Content-Type", "image/jpeg")
		io.Copy(resp, img)
	})
	mux.HandleFunc("/spicy", func(resp http.ResponseWriter, req *http.Request) {
		img, err := os.Open("./images/spicy.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer img.Close()
		resp.Header().Set("Content-Type", "image/jpeg")
		io.Copy(resp, img)
	})
	return serve("192.168.1.179:8080", mux, stop)
}

func Webhandler(w http.ResponseWriter, r *http.Request) {
	//var Path = "/images/red.jpg"
	img, err := os.Open("./images/red.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)
}

func serveDebug(stop <-chan struct{}) error {
	return serve("127.0.0.1:8001", http.DefaultServeMux, stop)
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
