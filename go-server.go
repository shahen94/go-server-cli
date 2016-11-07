package main

import (
	"fmt"
	"flag"
	"net/http"
	"log"
)
func main() {
	port := flag.Int("port", 3000, "port on server")
	dir := flag.String("dir", "./", "directory of web files")
	flag.Parse()

	fs := http.Dir(*dir)
	handler := http.FileServer(fs)
	http.Handle("/", handler)
	log.Printf("Running on port %d\n", *port)
	addr := fmt.Sprintf("127.0.0.1:%d", *port)

	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}