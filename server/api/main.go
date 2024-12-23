// Run go mod init <app name> to initialise app
// Run go run go build && ./<build file>
// If you need a port, install "go get github.com/lpernett/godotenv", run "go mod vendor" and run "go mod tidy".
// To run a server, install "go get github.com/go-chi/chi" and "go get github.com/go-chi/cors", run "go mod vendor" and run "go mod tidy"

package main

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == ""{
		log.Fatal("PORT not found")
	}

	cfg := &config{}

	app := &application{
		config: *cfg,
	}

	r:= app.mount()

	log.Printf("Server running on port %v", port)

	log.Fatal(app.run(":" + port, r))
}

