package main

import (
	"crypto/tls"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	app.Get("*", Forward())

	// Create tls certificate
	cer, err := tls.LoadX509KeyPair("certs/ssl.cert", "certs/ssl.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates:       []tls.Certificate{cer},
		InsecureSkipVerify: true}

	// Create custom listener
	ln, err := tls.Listen("tcp", ":9999", config)
	if err != nil {
		panic(err)
	}

	// Start server with https/ssl enabled on http://localhost:9999
	log.Fatal(app.Listener(ln))
}
