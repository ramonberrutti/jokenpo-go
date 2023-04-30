package main

import "github.com/ramonberrutti/jokenpo-go/internal/server"

func main() {
	srv := server.Server{}

	srv.Run()
}
