package main

import (
	"log"

	"noreast_cms/internal/server"
)

func main() {
	log.Println("Starting Noreast CMS server...")
	server.Start()
}