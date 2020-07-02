package main

import "github.com/insomnia-dreams-official/service-catalog/internal/server"

// Create and run server.
func main() {
	s := server.New()
	s.Run()
}
