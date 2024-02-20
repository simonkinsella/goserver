package main

import (
	"flag"
	"fmt"
	"github.com/simonkinsella/goserver/internal/handler"
	"github.com/simonkinsella/goserver/internal/templater"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	templatesDir, server := parseArgs()

	_, err := os.Stat(templatesDir)
	if os.IsNotExist(err) {
		slog.Error("Templates dir not found.", "templateDir", templatesDir)
		os.Exit(1)
	}
	slog.Info("Templates dir found.", "templateDir", templatesDir)

	t := templater.New(templatesDir)

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handler.Index(t))

	slog.Info(`Starting webserver. Ctrl-C to end. `, "address", server)
	log.Fatal(http.ListenAndServe(server, mux))
}

func parseArgs() (templatesDir, serverAddr string) {
	flag.StringVar(&templatesDir, "templates", "templates", "path to templates directory")
	host := flag.String("host", "127.0.0.1", "bind on address")
	port := flag.Int("port", 8080, "bind on port")
	flag.Parse()
	return templatesDir, fmt.Sprintf("%s:%d", *host, *port)
}
