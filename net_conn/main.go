package main

import "flag"

var server *bool
var client *bool

func init() {
	server = flag.Bool("s", false, "server mode")
	client = flag.Bool("c", false, "client mode")
	flag.Parse()
}

func main() {
	if *server {
		Server()
	} else if *client {
		Client()
	} else {
		flag.Usage()
	}
}
