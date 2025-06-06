package main

import (
	"api/asr"
	"flag"
	"fmt"
	"golib/server"
	"golib/server/http"
	"log"
	"time"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "0.0.0.0:3000", "Addr")
	flag.Parse()

	if err := asr.Initialize(0x92, "model.bin", "tokens.json", 4); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serve HTTP on", addr)
	log.Fatal(server.ServeHTTP(addr, http.NewServer(), 5*time.Second))
}
