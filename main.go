package main

import (
	"log"
)

func main() {
	log.Fatal(NewServer(Serve{}).ListenAndServe())
}
