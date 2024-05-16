package main

import "log"

type header struct {
}

type body struct {
}

type res struct {
	method string
	path   string
	header header
	body   body
}

func main() {
	log.Println("Start App")

	log.Println(("End App"))
}
