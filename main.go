package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	rootDir string

	//go:embed assets
	staticFiles embed.FS

	assetsContent fs.FS
	mainTemplate  *template.Template

	host   string = "127.0.0.1"
	port   string = "8080"
	scheme string = "http://"
)

func init() {
	var err error
	var staticFS = fs.FS(staticFiles)

	assetsContent, err = fs.Sub(staticFS, "assets")
	if err != nil {
		log.Fatal(err)
	}

	// get current dir
	rootDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// parse template only once
	mainTemplate, err = template.New("main.html").Funcs(template.FuncMap{
		"humanReadable": byteCountBinary,
		"formatDate": func(t time.Time) string {
			return t.Format("2006-01-02 15:04:05")
		},
		"getParentDir": func(path string) string {
			// doing the filepath.Dir function twice because path comes with a "/" at the end
			return filepath.Dir(filepath.Dir(path))
		},
	}).ParseFS(assetsContent, "main.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.FileServer(http.Dir("."))

	http.Handle("/", handlerSelection(http.HandlerFunc(listFilesHandler), http.StripPrefix("/assets/", http.FileServer(http.FS(assetsContent)))))
	log.Println("serving", rootDir)
	http.ListenAndServe(net.JoinHostPort(host, port), nil)
}
