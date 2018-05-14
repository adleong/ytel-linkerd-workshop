package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
)

type handler struct {}

func (h *handler) HandleRequest(w http.ResponseWriter, req *http.Request) {

  passwords := [5]string{
  	"password",
  	"admin",
  	"hunter2",
  	"12345",
  	"qwerty",
  }

  pw := passwords[rand.Intn(len(passwords))]
	w.Write([]byte(pw))
}


func main() {
	addr := flag.String("addr", ":8501", "service port to run on")
	flag.Parse()

	fmt.Printf("serving on %s", *addr)

	httpHandler := handler{}
	http.HandleFunc("/", httpHandler.HandleRequest)
	http.ListenAndServe(*addr, nil)
}
